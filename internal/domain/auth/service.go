package auth

import (
	"context"

	authmodel "github.com/joaofilippe/subclub/internal/domain/auth/model"
)

type Service interface {
	Login(ctx context.Context, input authmodel.LoginInput) (*authmodel.TokenOutput, error)
	Lookup(ctx context.Context, input authmodel.LookupInput) ([]authmodel.AccountInfo, error)
}
