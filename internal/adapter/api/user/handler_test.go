package userhandler_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	userhandler "github.com/joaofilippe/subclub/internal/adapter/api/user"
	"github.com/joaofilippe/subclub/internal/domain/user/model"
	"github.com/labstack/echo/v4"
)

type mockUserService struct {
	createFn func(ctx context.Context, input model.CreateUserInput) (string, error)
}

func (m *mockUserService) Create(ctx context.Context, input model.CreateUserInput) (string, error) {
	return m.createFn(ctx, input)
}
func (m *mockUserService) GetByID(ctx context.Context, id string) (*model.User, error) {
	return nil, nil
}
func (m *mockUserService) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return nil, nil
}
func (m *mockUserService) GetByRole(ctx context.Context, role model.UserRole) ([]*model.User, error) {
	return nil, nil
}
func (m *mockUserService) GetByType(ctx context.Context, userType model.UserType) ([]*model.User, error) {
	return nil, nil
}
func (m *mockUserService) Update(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	return nil, nil
}
func (m *mockUserService) Delete(ctx context.Context, id string) error    { return nil }
func (m *mockUserService) List(ctx context.Context) ([]*model.User, error) { return nil, nil }

func newEchoContext(method, body string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			_ = c.JSON(he.Code, he.Message)
		}
	}
	req := httptest.NewRequest(method, "/users", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec), rec
}

func TestUserHandler_Create_Success(t *testing.T) {
	svc := &mockUserService{
		createFn: func(_ context.Context, _ model.CreateUserInput) (string, error) {
			return "generated-uuid", nil
		},
	}
	h := userhandler.NewUserHandler(svc)
	c, rec := newEchoContext(http.MethodPost, `{"email":"test@test.com","type":"individual","role":"admin"}`)

	if err := h.Create(c); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rec.Code != http.StatusCreated {
		t.Errorf("expected status 201, got %d — body: %s", rec.Code, rec.Body.String())
	}
}

func TestUserHandler_Create_MissingFields(t *testing.T) {
	h := userhandler.NewUserHandler(&mockUserService{})
	c, rec := newEchoContext(http.MethodPost, `{"email":"only@email.com"}`)

	if err := h.Create(c); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", rec.Code)
	}
}

func TestUserHandler_Create_ServiceError(t *testing.T) {
	svc := &mockUserService{
		createFn: func(_ context.Context, _ model.CreateUserInput) (string, error) {
			return "", errors.New("db error")
		},
	}
	h := userhandler.NewUserHandler(svc)
	c, rec := newEchoContext(http.MethodPost, `{"email":"err@test.com","type":"individual","role":"admin"}`)

	if err := h.Create(c); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", rec.Code)
	}
}
