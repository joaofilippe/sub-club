package auth

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	authdomain "github.com/joaofilippe/subclub/internal/domain/auth"
	authmodel "github.com/joaofilippe/subclub/internal/domain/auth/model"
	"github.com/joaofilippe/subclub/internal/web/common"
)

type AuthHandler struct {
	service authdomain.Service
}

func NewAuthHandler(service authdomain.Service) *AuthHandler {
	return &AuthHandler{service: service}
}

// Login godoc
// @Summary      Authenticate
// @Description  Validates credentials and returns a signed JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body      LoginRequestDTO          true  "Login credentials"
// @Success      200   {object}  common.Response{data=TokenResponseDTO}
// @Failure      400   {object}  common.Response
// @Failure      401   {object}  common.Response
// @Failure      500   {object}  common.Response
// @Router       /auth/login [post]
func (h *AuthHandler) Login(c echo.Context) error {
	var req LoginRequestDTO
	if err := c.Bind(&req); err != nil {
		return common.Error(c, http.StatusBadRequest, "Invalid request payload")
	}
	if req.Email == "" || req.Password == "" {
		return common.Error(c, http.StatusBadRequest, "Email and password are required")
	}

	out, err := h.service.Login(c.Request().Context(), authmodel.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		if errors.Is(err, authmodel.ErrInvalidCredentials) {
			return common.Error(c, http.StatusUnauthorized, "Invalid email or password")
		}
		return common.Error(c, http.StatusInternalServerError, "Authentication failed")
	}

	return common.Success(c, http.StatusOK, "Login successful", TokenResponseDTO{Token: out.Token})
}
