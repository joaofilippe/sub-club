package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/joaofilippe/subclub/ent/enttest"
	repository "github.com/joaofilippe/subclub/internal/application/repository/user"
	userdomain "github.com/joaofilippe/subclub/internal/domain/user"
	"github.com/joaofilippe/subclub/internal/domain/user/model"
	_ "modernc.org/sqlite"
)

const testDSN = "file::memory:?_pragma=foreign_keys(1)"

func newTestRepo(t *testing.T) userdomain.Repository {
	t.Helper()
	client := enttest.Open(t, "sqlite3", testDSN)
	t.Cleanup(func() { client.Close() })
	return repository.NewUserEntRepository(client)
}

func makeUser(email string) *model.User {
	now := time.Now()
	return &model.User{
		ID:        "00000000-0000-0000-0000-000000000001",
		Name:      "Test User",
		Email:     email,
		Password:  "hashed_password",
		Type:      model.UserTypeIndividual,
		Role:      model.UserRoleAdmin,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func TestUserRepository_Create_Success(t *testing.T) {
	repo := newTestRepo(t)
	ctx := context.Background()

	u := makeUser("create@test.com")
	if err := repo.Create(ctx, u); err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

func TestUserRepository_Create_DuplicateEmail(t *testing.T) {
	repo := newTestRepo(t)
	ctx := context.Background()

	u1 := makeUser("dup@test.com")
	if err := repo.Create(ctx, u1); err != nil {
		t.Fatalf("first create failed: %v", err)
	}

	u2 := makeUser("dup@test.com")
	u2.ID = "00000000-0000-0000-0000-000000000002"
	if err := repo.Create(ctx, u2); err == nil {
		t.Fatal("expected error for duplicate email, got nil")
	}
}

func TestUserRepository_GetByID_Found(t *testing.T) {
	repo := newTestRepo(t)
	ctx := context.Background()

	u := makeUser("getbyid@test.com")
	if err := repo.Create(ctx, u); err != nil {
		t.Fatalf("create failed: %v", err)
	}

	got, err := repo.GetByID(ctx, u.ID)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if got.Email != u.Email {
		t.Errorf("expected email %q, got %q", u.Email, got.Email)
	}
}

func TestUserRepository_GetByID_NotFound(t *testing.T) {
	repo := newTestRepo(t)
	ctx := context.Background()

	_, err := repo.GetByID(ctx, "00000000-0000-0000-0000-000000000099")
	if err == nil {
		t.Fatal("expected error for missing ID, got nil")
	}
}

func TestUserRepository_GetByID_InvalidUUID(t *testing.T) {
	repo := newTestRepo(t)
	ctx := context.Background()

	_, err := repo.GetByID(ctx, "not-a-uuid")
	if err == nil {
		t.Fatal("expected error for invalid UUID, got nil")
	}
}

func TestUserRepository_GetByEmail_Found(t *testing.T) {
	repo := newTestRepo(t)
	ctx := context.Background()

	u := makeUser("byemail@test.com")
	if err := repo.Create(ctx, u); err != nil {
		t.Fatalf("create failed: %v", err)
	}

	got, err := repo.GetByEmail(ctx, u.Email)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if got.ID != u.ID {
		t.Errorf("expected ID %q, got %q", u.ID, got.ID)
	}
}

func TestUserRepository_List_ReturnsPersisted(t *testing.T) {
	repo := newTestRepo(t)
	ctx := context.Background()

	u1 := makeUser("list1@test.com")
	u2 := makeUser("list2@test.com")
	u2.ID = "00000000-0000-0000-0000-000000000002"

	_ = repo.Create(ctx, u1)
	_ = repo.Create(ctx, u2)

	users, err := repo.List(ctx)
	if err != nil {
		t.Fatalf("list failed: %v", err)
	}
	if len(users) != 2 {
		t.Errorf("expected 2 users, got %d", len(users))
	}
}

func TestUserRepository_List_IgnoresSoftDeleted(t *testing.T) {
	repo := newTestRepo(t)
	ctx := context.Background()

	u := makeUser("softdel@test.com")
	if err := repo.Create(ctx, u); err != nil {
		t.Fatalf("create failed: %v", err)
	}
	if err := repo.Delete(ctx, u.ID); err != nil {
		t.Fatalf("delete failed: %v", err)
	}

	users, err := repo.List(ctx)
	if err != nil {
		t.Fatalf("list failed: %v", err)
	}
	if len(users) != 0 {
		t.Errorf("expected 0 active users after soft delete, got %d", len(users))
	}
}

func TestUserRepository_Delete_SoftDelete(t *testing.T) {
	repo := newTestRepo(t)
	ctx := context.Background()

	u := makeUser("delete@test.com")
	if err := repo.Create(ctx, u); err != nil {
		t.Fatalf("create failed: %v", err)
	}
	if err := repo.Delete(ctx, u.ID); err != nil {
		t.Fatalf("delete failed: %v", err)
	}

	_, err := repo.GetByID(ctx, u.ID)
	if err == nil {
		t.Fatal("expected error after soft delete, got nil")
	}
}

func TestUserRepository_Update(t *testing.T) {
	repo := newTestRepo(t)
	ctx := context.Background()

	u := makeUser("update@test.com")
	if err := repo.Create(ctx, u); err != nil {
		t.Fatalf("create failed: %v", err)
	}

	u.Name = "Updated Name"
	u.Role = model.UserRoleOperations
	u.UpdatedAt = time.Now()

	if err := repo.Update(ctx, u); err != nil {
		t.Fatalf("update failed: %v", err)
	}

	got, err := repo.GetByID(ctx, u.ID)
	if err != nil {
		t.Fatalf("get after update failed: %v", err)
	}
	if got.Name != "Updated Name" {
		t.Errorf("expected name %q, got %q", "Updated Name", got.Name)
	}
	if got.Role != model.UserRoleOperations {
		t.Errorf("expected role %q, got %q", model.UserRoleOperations, got.Role)
	}
}

func TestUserRepository_GetByRole(t *testing.T) {
	repo := newTestRepo(t)
	ctx := context.Background()

	u1 := makeUser("admin@test.com")
	u1.Role = model.UserRoleAdmin
	u2 := makeUser("ops@test.com")
	u2.ID = "00000000-0000-0000-0000-000000000002"
	u2.Role = model.UserRoleOperations

	_ = repo.Create(ctx, u1)
	_ = repo.Create(ctx, u2)

	admins, err := repo.GetByRole(ctx, model.UserRoleAdmin)
	if err != nil {
		t.Fatalf("GetByRole failed: %v", err)
	}
	if len(admins) != 1 {
		t.Errorf("expected 1 admin, got %d", len(admins))
	}
}

func TestUserRepository_GetByType(t *testing.T) {
	repo := newTestRepo(t)
	ctx := context.Background()

	u1 := makeUser("corp@test.com")
	u1.Type = model.UserTypeCorporate
	u2 := makeUser("indiv@test.com")
	u2.ID = "00000000-0000-0000-0000-000000000002"
	u2.Type = model.UserTypeIndividual

	_ = repo.Create(ctx, u1)
	_ = repo.Create(ctx, u2)

	corps, err := repo.GetByType(ctx, model.UserTypeCorporate)
	if err != nil {
		t.Fatalf("GetByType failed: %v", err)
	}
	if len(corps) != 1 {
		t.Errorf("expected 1 corporate user, got %d", len(corps))
	}
}
