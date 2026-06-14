package usecase

import (
	"context"
	"strings"
	"sync"

	"golang.org/x/crypto/bcrypt"

	"github.com/joaofilippe/subclub/internal/domain/account"
	"github.com/joaofilippe/subclub/internal/domain/accountdomain"
	authdomain "github.com/joaofilippe/subclub/internal/domain/auth"
	authmodel "github.com/joaofilippe/subclub/internal/domain/auth/model"
	"github.com/joaofilippe/subclub/internal/domain/user"
)

type LoginUseCase struct {
	userRepo          user.Repository
	accountRepo       account.Repository
	accountDomainRepo accountdomain.Repository
	tenantRepo        authdomain.TenantRepository
	jwtSecret         []byte
}

func NewLoginUseCase(
	userRepo user.Repository,
	accountRepo account.Repository,
	accountDomainRepo accountdomain.Repository,
	tenantRepo authdomain.TenantRepository,
	jwtSecret []byte,
) *LoginUseCase {
	return &LoginUseCase{
		userRepo:          userRepo,
		accountRepo:       accountRepo,
		accountDomainRepo: accountDomainRepo,
		tenantRepo:        tenantRepo,
		jwtSecret:         jwtSecret,
	}
}

func (uc *LoginUseCase) Execute(ctx context.Context, input authmodel.LoginInput) (*authmodel.TokenOutput, error) {
	// Try system login first (public schema).
	u, err := uc.userRepo.GetByEmail(ctx, input.Email)
	if err == nil {
		if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(input.Password)) != nil {
			return nil, authmodel.ErrInvalidCredentials
		}
		accountSlug := ""
		if u.AccountID != nil {
			if acc, err := uc.accountRepo.GetByID(ctx, *u.AccountID); err == nil {
				accountSlug = acc.Slug
			}
		}
		return signToken(uc.jwtSecret, u.ID, accountSlug, string(u.Role))
	}

	// Tenant lookup by email.
	accounts, err := uc.findTenantAccountsForEmail(ctx, input.Email)
	if err != nil {
		return nil, authmodel.ErrInvalidCredentials
	}
	switch len(accounts) {
	case 0:
		return nil, authmodel.ErrInvalidCredentials
	case 1:
		return uc.tenantLogin(ctx, accounts[0].Slug, input.Email, input.Password)
	default:
		return nil, authmodel.ErrMultipleAccountsFound
	}
}

func (uc *LoginUseCase) tenantLogin(ctx context.Context, slug, email, password string) (*authmodel.TokenOutput, error) {
	u, err := uc.tenantRepo.FindUserByEmailAndSlug(ctx, slug, email)
	if err != nil {
		return nil, authmodel.ErrInvalidCredentials
	}
	if bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)) != nil {
		return nil, authmodel.ErrInvalidCredentials
	}
	return signToken(uc.jwtSecret, u.ID, slug, u.Role)
}

func (uc *LoginUseCase) findTenantAccountsForEmail(ctx context.Context, email string) ([]authmodel.AccountInfo, error) {
	parts := strings.SplitN(email, "@", 2)
	if len(parts) == 2 {
		if info, err := uc.accountDomainRepo.GetByDomain(ctx, parts[1]); err == nil {
			return []authmodel.AccountInfo{{Slug: info.Slug, Name: info.Name}}, nil
		}
	}

	allAccounts, err := uc.accountRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	type hit struct {
		info  authmodel.AccountInfo
		found bool
	}
	results := make(chan hit, len(allAccounts))
	var wg sync.WaitGroup
	for _, acc := range allAccounts {
		wg.Add(1)
		go func(slug, name string) {
			defer wg.Done()
			exists, err := uc.tenantRepo.UserExistsByEmailAndSlug(ctx, slug, email)
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
