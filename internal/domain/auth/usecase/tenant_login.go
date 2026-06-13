package usecase

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	authdomain "github.com/joaofilippe/subclub/internal/domain/auth"
	authmodel "github.com/joaofilippe/subclub/internal/domain/auth/model"
)

type TenantLoginUseCase struct {
	tenantRepo authdomain.TenantRepository
	jwtSecret  []byte
}

func NewTenantLoginUseCase(
	tenantRepo authdomain.TenantRepository,
	jwtSecret []byte,
) *TenantLoginUseCase {
	return &TenantLoginUseCase{
		tenantRepo: tenantRepo,
		jwtSecret:  jwtSecret,
	}
}

func (uc *TenantLoginUseCase) Execute(ctx context.Context, input authmodel.TenantLoginInput) (*authmodel.TokenOutput, error) {
	u, err := uc.tenantRepo.FindUserByEmailAndSlug(ctx, input.AccountSlug, input.Email)
	if err != nil {
		return nil, authmodel.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(input.Password)); err != nil {
		return nil, authmodel.ErrInvalidCredentials
	}

	return signToken(uc.jwtSecret, u.ID, input.AccountSlug, u.Role)
}
