package usecase

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	authdomain "github.com/joaofilippe/subclub/internal/domain/auth"
	authmodel "github.com/joaofilippe/subclub/internal/domain/auth/model"
)

type LoginByUsernameUseCase struct {
	tenantRepo authdomain.TenantRepository
	jwtSecret  []byte
}

func NewLoginByUsernameUseCase(tenantRepo authdomain.TenantRepository, jwtSecret []byte) *LoginByUsernameUseCase {
	return &LoginByUsernameUseCase{tenantRepo: tenantRepo, jwtSecret: jwtSecret}
}

func (uc *LoginByUsernameUseCase) Execute(ctx context.Context, input authmodel.UsernameLoginInput) (*authmodel.TokenOutput, error) {
	u, err := uc.tenantRepo.FindUserByUsernameAndSlug(ctx, input.AccountSlug, input.Username)
	if err != nil {
		return nil, authmodel.ErrInvalidCredentials
	}
	if bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(input.Password)) != nil {
		return nil, authmodel.ErrInvalidCredentials
	}
	return signToken(uc.jwtSecret, u.ID, input.AccountSlug, u.Role)
}
