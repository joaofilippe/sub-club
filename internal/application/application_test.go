package application

import (
	"context"
	"testing"

	"github.com/joaofilippe/subclub/ent/enttest"
	_ "modernc.org/sqlite"
)

const testDSN = "file::memory:?_pragma=foreign_keys(1)"

func TestApplication_InitServices_PopulatesAllServices(t *testing.T) {
	client := enttest.Open(t, "sqlite3", testDSN)
	t.Cleanup(func() { client.Close() })

	app := New(nil)
	if err := app.initServices(context.Background(), client); err != nil {
		t.Fatalf("initServices returned error: %v", err)
	}

	if app.UserService == nil {
		t.Error("UserService is nil after initServices")
	}
	if app.ClientService == nil {
		t.Error("ClientService is nil after initServices")
	}
	if app.PlanService == nil {
		t.Error("PlanService is nil after initServices")
	}
	if app.ProductService == nil {
		t.Error("ProductService is nil after initServices")
	}
	if app.SubscriptionService == nil {
		t.Error("SubscriptionService is nil after initServices")
	}
}
