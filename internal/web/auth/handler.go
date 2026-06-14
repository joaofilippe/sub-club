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
// @Summary      Authenticate by email
// @Description  Validates email + password and returns a signed JWT. If the email belongs to a single tenant the user is logged into that tenant automatically. If it matches multiple tenants a 409 is returned instructing the client to use the username + account_slug login instead.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body      LoginRequestDTO                      true  "Login credentials"
// @Success      200   {object}  common.Response{data=TokenResponseDTO}
// @Failure      400   {object}  common.Response
// @Failure      401   {object}  common.Response
// @Failure      409   {object}  common.Response
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
		if errors.Is(err, authmodel.ErrMultipleAccountsFound) {
			return common.Error(c, http.StatusConflict, err.Error())
		}
		if errors.Is(err, authmodel.ErrInvalidCredentials) {
			return common.Error(c, http.StatusUnauthorized, "Invalid email or password")
		}
		return common.Error(c, http.StatusInternalServerError, "Authentication failed")
	}

	return common.Success(c, http.StatusOK, "Login successful", TokenResponseDTO{Token: out.Token})
}

// LoginByUsername godoc
// @Summary      Authenticate by username and account slug
// @Description  Validates username + account_slug + password for a tenant user. Use this when the email-based login returns 409 (multiple accounts found).
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body      UsernameLoginRequestDTO              true  "Username login credentials"
// @Success      200   {object}  common.Response{data=TokenResponseDTO}
// @Failure      400   {object}  common.Response
// @Failure      401   {object}  common.Response
// @Failure      500   {object}  common.Response
// @Router       /api/v1/auth/login/username [post]
func (h *AuthHandler) LoginByUsername(c echo.Context) error {
	var req UsernameLoginRequestDTO
	if err := c.Bind(&req); err != nil {
		return common.Error(c, http.StatusBadRequest, "Invalid request payload")
	}
	if req.Username == "" || req.AccountSlug == "" || req.Password == "" {
		return common.Error(c, http.StatusBadRequest, "Username, account_slug and password are required")
	}

	out, err := h.service.LoginByUsername(c.Request().Context(), authmodel.UsernameLoginInput{
		Username:    req.Username,
		AccountSlug: req.AccountSlug,
		Password:    req.Password,
	})
	if err != nil {
		if errors.Is(err, authmodel.ErrInvalidCredentials) {
			return common.Error(c, http.StatusUnauthorized, "Invalid username, account or password")
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
