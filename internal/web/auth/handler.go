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
// @Router       /api/v1/auth/login [post]
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

// Lookup godoc
// @Summary      Lookup accounts by email
// @Description  Returns all accounts that have a user registered with the given email. Uses domain-based lookup first; falls back to a parallel tenant scan for generic domains (gmail, outlook, etc.)
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body      LookupRequestDTO             true  "Email to look up"
// @Success      200   {object}  common.Response{data=[]AccountInfoDTO}
// @Failure      400   {object}  common.Response
// @Failure      500   {object}  common.Response
// @Router       /api/v1/auth/lookup [post]
func (h *AuthHandler) Lookup(c echo.Context) error {
	var req LookupRequestDTO
	if err := c.Bind(&req); err != nil {
		return common.Error(c, http.StatusBadRequest, "Invalid request payload")
	}
	if req.Email == "" {
		return common.Error(c, http.StatusBadRequest, "Email is required")
	}

	accounts, err := h.service.Lookup(c.Request().Context(), authmodel.LookupInput{Email: req.Email})
	if err != nil {
		return common.Error(c, http.StatusInternalServerError, "Lookup failed")
	}

	result := make([]AccountInfoDTO, len(accounts))
	for i, a := range accounts {
		result[i] = AccountInfoDTO{Slug: a.Slug, Name: a.Name}
	}
	return common.Success(c, http.StatusOK, "Lookup successful", result)
}

// TenantLogin godoc
// @Summary      Authenticate tenant user
// @Description  Validates credentials against a specific tenant schema and returns a signed JWT
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body      TenantLoginRequestDTO        true  "Tenant login credentials"
// @Success      200   {object}  common.Response{data=TokenResponseDTO}
// @Failure      400   {object}  common.Response
// @Failure      401   {object}  common.Response
// @Failure      500   {object}  common.Response
// @Router       /api/v1/auth/tenant-login [post]
func (h *AuthHandler) TenantLogin(c echo.Context) error {
	var req TenantLoginRequestDTO
	if err := c.Bind(&req); err != nil {
		return common.Error(c, http.StatusBadRequest, "Invalid request payload")
	}
	if req.Email == "" || req.Password == "" || req.AccountSlug == "" {
		return common.Error(c, http.StatusBadRequest, "Email, password and account_slug are required")
	}

	out, err := h.service.TenantLogin(c.Request().Context(), authmodel.TenantLoginInput{
		Email:       req.Email,
		Password:    req.Password,
		AccountSlug: req.AccountSlug,
	})
	if err != nil {
		if errors.Is(err, authmodel.ErrInvalidCredentials) {
			return common.Error(c, http.StatusUnauthorized, "Invalid credentials")
		}
		return common.Error(c, http.StatusInternalServerError, "Authentication failed")
	}

	return common.Success(c, http.StatusOK, "Login successful", TokenResponseDTO{Token: out.Token})
}
