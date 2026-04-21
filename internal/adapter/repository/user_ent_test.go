package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/joaofilippe/subclub/ent/enttest"
	"github.com/joaofilippe/subclub/internal/adapter/repository"
	domain "github.com/joaofilippe/subclub/internal/domain/user"
	_ "modernc.org/sqlite" // registers "sqlite" driver; "sqlite3" alias set in testsetup_test.go
)

const testDSN = "file::memory:?_pragma=foreign_keys(1)"

func newTestRepo(t *testing.T) domain.Repository {
	t.Helper()
	client := enttest.Open(t, "sqlite3", testDSN)
	t.Cleanup(func() { client.Close() })
	return repository.NewUserEntRepository(client)
}

func makeUser(email string) *domain.User {
	now := time.Now()
	return &domain.User{
		ID:        "00000000-0000-0000-0000-000000000001",
		Name:      "Test User",
		Email:     email,
		Password:  "hashed_password",
		Type:      domain.UserTypeIndividual,
		Role:      domain.UserRoleAdmin,
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
	u.Role = domain.UserRoleOperations
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
	if got.Role != domain.UserRoleOperations {
		t.Errorf("expected role %q, got %q", domain.UserRoleOperations, got.Role)
	}
}

func TestUserRepository_GetByRole(t *testing.T) {
	repo := newTestRepo(t)
	ctx := context.Background()

	u1 := makeUser("admin@test.com")
	u1.Role = domain.UserRoleAdmin
	u2 := makeUser("ops@test.com")
	u2.ID = "00000000-0000-0000-0000-000000000002"
	u2.Role = domain.UserRoleOperations

	_ = repo.Create(ctx, u1)
	_ = repo.Create(ctx, u2)

	admins, err := repo.GetByRole(ctx, domain.UserRoleAdmin)
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
	u1.Type = domain.UserTypeCorporate
	u2 := makeUser("indiv@test.com")
	u2.ID = "00000000-0000-0000-0000-000000000002"
	u2.Type = domain.UserTypeIndividual

	_ = repo.Create(ctx, u1)
	_ = repo.Create(ctx, u2)

	corps, err := repo.GetByType(ctx, domain.UserTypeCorporate)
	if err != nil {
		t.Fatalf("GetByType failed: %v", err)
	}
	if len(corps) != 1 {
		t.Errorf("expected 1 corporate user, got %d", len(corps))
	}
}
