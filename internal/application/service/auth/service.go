package service

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/joaofilippe/subclub/ent"
	entuser "github.com/joaofilippe/subclub/ent/user"
	"github.com/joaofilippe/subclub/internal/domain/account"
	"github.com/joaofilippe/subclub/internal/domain/accountdomain"
	authmodel "github.com/joaofilippe/subclub/internal/domain/auth/model"
	"github.com/joaofilippe/subclub/internal/domain/user"
	"github.com/joaofilippe/subclub/internal/infra/database"
)

// Claims are the JWT payload fields embedded in every token.
type Claims struct {
	UserID      string `json:"user_id"`
	AccountSlug string `json:"account_slug"`
	Role        string `json:"role"`
	jwt.RegisteredClaims
}

type AuthService struct {
	userRepo          user.Repository
	accountRepo       account.Repository
	accountDomainRepo accountdomain.Repository
	tenantManager     *database.TenantClientManager
	jwtSecret         []byte
}

func NewAuthService(
	userRepo user.Repository,
	accountRepo account.Repository,
	accountDomainRepo accountdomain.Repository,
	tenantManager *database.TenantClientManager,
	jwtSecret []byte,
) *AuthService {
	return &AuthService{
		userRepo:          userRepo,
		accountRepo:       accountRepo,
		accountDomainRepo: accountDomainRepo,
		tenantManager:     tenantManager,
		jwtSecret:         jwtSecret,
	}
}

func (s *AuthService) Login(ctx context.Context, input authmodel.LoginInput) (*authmodel.TokenOutput, error) {
	u, err := s.userRepo.GetByEmail(ctx, input.Email)
	if err != nil {
		return nil, authmodel.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(input.Password)); err != nil {
		return nil, authmodel.ErrInvalidCredentials
	}

	accountSlug := ""
	if u.AccountID != nil {
		acc, err := s.accountRepo.GetByID(ctx, *u.AccountID)
		if err == nil {
			accountSlug = acc.Slug
		}
	}

	return s.signToken(u.ID, accountSlug, string(u.Role))
}

// Lookup finds all accounts that have a user with the given email.
// It first tries a fast domain-based lookup; if the domain is generic (gmail, etc.)
// it falls back to a parallel scan across all active tenant schemas.
func (s *AuthService) Lookup(ctx context.Context, input authmodel.LookupInput) ([]authmodel.AccountInfo, error) {
	parts := strings.SplitN(input.Email, "@", 2)
	if len(parts) == 2 {
		info, err := s.accountDomainRepo.GetByDomain(ctx, parts[1])
		if err == nil {
			return []authmodel.AccountInfo{{Slug: info.Slug, Name: info.Name}}, nil
		}
	}

	accounts, err := s.accountRepo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("listing accounts: %w", err)
	}

	type hit struct {
		info  authmodel.AccountInfo
		found bool
	}

	results := make(chan hit, len(accounts))
	var wg sync.WaitGroup

	for _, acc := range accounts {
		wg.Add(1)
		go func(slug, name string) {
			defer wg.Done()
			client, err := s.tenantManager.GetOrCreate(slug)
			if err != nil {
				results <- hit{}
				return
			}
			exists, err := client.User.Query().
				Where(entuser.EmailEQ(input.Email), entuser.DeletedAtIsNil()).
				Exist(ctx)
			if err != nil || !exists {
				results <- hit{}
				return
			}
			results <- hit{info: authmodel.AccountInfo{Slug: slug, Name: name}, found: true}
		}(acc.Slug, acc.Name)
	}

	wg.Wait()
	close(results)

	var matches []authmodel.AccountInfo
	for r := range results {
		if r.found {
			matches = append(matches, r.info)
		}
	}
	return matches, nil
}

// TenantLogin authenticates a user against a specific tenant schema and returns a JWT.
func (s *AuthService) TenantLogin(ctx context.Context, input authmodel.TenantLoginInput) (*authmodel.TokenOutput, error) {
	client, err := s.tenantManager.GetOrCreate(input.AccountSlug)
	if err != nil {
		return nil, authmodel.ErrInvalidCredentials
	}

	u, err := findTenantUserByEmail(ctx, client, input.Email)
	if err != nil {
		return nil, authmodel.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(input.Password)); err != nil {
		return nil, authmodel.ErrInvalidCredentials
	}

	return s.signToken(u.ID.String(), input.AccountSlug, string(u.Role))
}

func (s *AuthService) signToken(userID, accountSlug, role string) (*authmodel.TokenOutput, error) {
	claims := Claims{
		UserID:      userID,
		AccountSlug: accountSlug,
		Role:        role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return nil, fmt.Errorf("signing token: %w", err)
	}
	return &authmodel.TokenOutput{Token: signed}, nil
}

func findTenantUserByEmail(ctx context.Context, client *ent.Client, email string) (*ent.User, error) {
	return client.User.Query().
		Where(entuser.EmailEQ(email), entuser.DeletedAtIsNil()).
		Only(ctx)
}
