package service

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/joaofilippe/subclub/internal/domain/account"
	authmodel "github.com/joaofilippe/subclub/internal/domain/auth/model"
	"github.com/joaofilippe/subclub/internal/domain/user"
)

// Claims are the JWT payload fields embedded in every token.
type Claims struct {
	UserID      string `json:"user_id"`
	AccountSlug string `json:"account_slug"`
	Role        string `json:"role"`
	jwt.RegisteredClaims
}

type AuthService struct {
	userRepo    user.Repository
	accountRepo account.Repository
	jwtSecret   []byte
}

func NewAuthService(userRepo user.Repository, accountRepo account.Repository, jwtSecret []byte) *AuthService {
	return &AuthService{
		userRepo:    userRepo,
		accountRepo: accountRepo,
		jwtSecret:   jwtSecret,
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

	claims := Claims{
		UserID:      u.ID,
		AccountSlug: accountSlug,
		Role:        string(u.Role),
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
