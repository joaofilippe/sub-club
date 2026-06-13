package authusecase

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	entuser "github.com/joaofilippe/subclub/ent/user"
	authmodel "github.com/joaofilippe/subclub/internal/domain/auth/model"
	"github.com/joaofilippe/subclub/internal/infra/database"

	"github.com/joaofilippe/subclub/ent"
)

type TenantLoginUseCase struct {
	tenantManager *database.TenantClientManager
	jwtSecret     []byte
}

func NewTenantLoginUseCase(
	tenantManager *database.TenantClientManager,
	jwtSecret []byte,
) *TenantLoginUseCase {
	return &TenantLoginUseCase{
		tenantManager: tenantManager,
		jwtSecret:     jwtSecret,
	}
}

func (uc *TenantLoginUseCase) Execute(ctx context.Context, input authmodel.TenantLoginInput) (*authmodel.TokenOutput, error) {
	client, err := uc.tenantManager.GetOrCreate(input.AccountSlug)
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

	return signToken(uc.jwtSecret, u.ID.String(), input.AccountSlug, string(u.Role))
}

func findTenantUserByEmail(ctx context.Context, client *ent.Client, email string) (*ent.User, error) {
	return client.User.Query().
		Where(entuser.EmailEQ(email), entuser.DeletedAtIsNil()).
		Only(ctx)
}
