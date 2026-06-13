package auth

import (
	"context"

	authmodel "github.com/joaofilippe/subclub/internal/domain/auth/model"
)

type LoginUseCase interface {
	Execute(ctx context.Context, input authmodel.LoginInput) (*authmodel.TokenOutput, error)
}

type TenantLoginUseCase interface {
	Execute(ctx context.Context, input authmodel.TenantLoginInput) (*authmodel.TokenOutput, error)
}

type LookupUseCase interface {
	Execute(ctx context.Context, input authmodel.LookupInput) ([]authmodel.AccountInfo, error)
}
