package service

import (
	"context"

	authmodel "github.com/joaofilippe/subclub/internal/domain/auth/model"
	"github.com/joaofilippe/subclub/internal/domain/auth/usecase"
)

type AuthService struct {
	loginUC  *usecase.LoginUseCase
	lookupUC *usecase.LookupUseCase
}

func NewAuthService(
	loginUC *usecase.LoginUseCase,
	lookupUC *usecase.LookupUseCase,
) *AuthService {
	return &AuthService{
		loginUC:  loginUC,
		lookupUC: lookupUC,
	}
}

func (s *AuthService) Login(ctx context.Context, input authmodel.LoginInput) (*authmodel.TokenOutput, error) {
	return s.loginUC.Execute(ctx, input)
}

func (s *AuthService) Lookup(ctx context.Context, input authmodel.LookupInput) ([]authmodel.AccountInfo, error) {
	return s.lookupUC.Execute(ctx, input)
}
