package service

import (
	"context"

	authdomain "github.com/joaofilippe/subclub/internal/domain/auth"
	authmodel "github.com/joaofilippe/subclub/internal/domain/auth/model"
)

type AuthService struct {
	loginUC       authdomain.LoginUseCase
	tenantLoginUC authdomain.TenantLoginUseCase
	lookupUC      authdomain.LookupUseCase
}

func NewAuthService(
	loginUC authdomain.LoginUseCase,
	tenantLoginUC authdomain.TenantLoginUseCase,
	lookupUC authdomain.LookupUseCase,
) *AuthService {
	return &AuthService{
		loginUC:       loginUC,
		tenantLoginUC: tenantLoginUC,
		lookupUC:      lookupUC,
	}
}

func (s *AuthService) Login(ctx context.Context, input authmodel.LoginInput) (*authmodel.TokenOutput, error) {
	return s.loginUC.Execute(ctx, input)
}

func (s *AuthService) TenantLogin(ctx context.Context, input authmodel.TenantLoginInput) (*authmodel.TokenOutput, error) {
	return s.tenantLoginUC.Execute(ctx, input)
}

func (s *AuthService) Lookup(ctx context.Context, input authmodel.LookupInput) ([]authmodel.AccountInfo, error) {
	return s.lookupUC.Execute(ctx, input)
}
