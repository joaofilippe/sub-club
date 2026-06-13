package usecase

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/joaofilippe/subclub/internal/domain/account"
	authdomain "github.com/joaofilippe/subclub/internal/domain/auth"
	authmodel "github.com/joaofilippe/subclub/internal/domain/auth/model"
	"github.com/joaofilippe/subclub/internal/domain/user"
)

type LoginUseCase struct {
	userRepo    user.Repository
	accountRepo account.Repository
	tenantRepo  authdomain.TenantRepository
	jwtSecret   []byte
}

func NewLoginUseCase(
	userRepo user.Repository,
	accountRepo account.Repository,
	tenantRepo authdomain.TenantRepository,
	jwtSecret []byte,
) *LoginUseCase {
	return &LoginUseCase{
		userRepo:    userRepo,
		accountRepo: accountRepo,
		tenantRepo:  tenantRepo,
		jwtSecret:   jwtSecret,
	}
}

func (uc *LoginUseCase) Execute(ctx context.Context, input authmodel.LoginInput) (*authmodel.TokenOutput, error) {
	if input.AccountSlug != "" {
		return uc.tenantLogin(ctx, input)
	}
	return uc.systemLogin(ctx, input)
}

func (uc *LoginUseCase) systemLogin(ctx context.Context, input authmodel.LoginInput) (*authmodel.TokenOutput, error) {
	u, err := uc.userRepo.GetByEmail(ctx, input.Email)
	if err != nil {
		return nil, authmodel.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(input.Password)); err != nil {
		return nil, authmodel.ErrInvalidCredentials
	}

	accountSlug := ""
	if u.AccountID != nil {
		acc, err := uc.accountRepo.GetByID(ctx, *u.AccountID)
		if err == nil {
			accountSlug = acc.Slug
		}
	}

	return signToken(uc.jwtSecret, u.ID, accountSlug, string(u.Role))
}

func (uc *LoginUseCase) tenantLogin(ctx context.Context, input authmodel.LoginInput) (*authmodel.TokenOutput, error) {
	u, err := uc.tenantRepo.FindUserByEmailAndSlug(ctx, input.AccountSlug, input.Email)
	if err != nil {
		return nil, authmodel.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(input.Password)); err != nil {
		return nil, authmodel.ErrInvalidCredentials
	}

	return signToken(uc.jwtSecret, u.ID, input.AccountSlug, u.Role)
}
