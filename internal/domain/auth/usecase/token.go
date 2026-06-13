package usecase

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	authmodel "github.com/joaofilippe/subclub/internal/domain/auth/model"
)

type Claims struct {
	UserID      string `json:"user_id"`
	AccountSlug string `json:"account_slug"`
	Role        string `json:"role"`
	jwt.RegisteredClaims
}

func signToken(jwtSecret []byte, userID, accountSlug, role string) (*authmodel.TokenOutput, error) {
	claims := Claims{
		UserID:      userID,
		AccountSlug: accountSlug,
		Role:        role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(jwtSecret)
	if err != nil {
		return nil, fmt.Errorf("signing token: %w", err)
	}
	return &authmodel.TokenOutput{Token: signed}, nil
}
