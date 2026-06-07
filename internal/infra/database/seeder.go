package database

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/ent"
	"github.com/joaofilippe/subclub/ent/systemuser"
	"github.com/joaofilippe/subclub/internal/config"
	"github.com/joaofilippe/subclub/internal/test/faker"
	"golang.org/x/crypto/bcrypt"
)

func SeedAll(ctx context.Context, client *ent.Client, cfg *config.Config, manager *TenantClientManager) {
	if !cfg.IsDevelopment() {
		log.Printf("[Seeder] Skipping: current environment is %q\n", cfg.AppEnv)
		return
	}

	log.Println("[Seeder] Checking public schema state...")

	count, err := client.SystemUser.Query().Count(ctx)
	if err != nil {
		log.Printf("[Seeder] Error counting system users: %v\n", err)
		return
	}

	if count > 0 {
		log.Println("[Seeder] Public schema already populated. Skipping seeds.")
		return
	}

	log.Println("[Seeder] Seeding public schema...")

	// System admin
	hash, err := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[Seeder] Error hashing password: %v\n", err)
		return
	}
	_, err = client.SystemUser.Create().
		SetID(uuid.New()).
		SetName("Admin").
		SetEmail("adm@adm.com").
		SetPassword(string(hash)).
		SetType(systemuser.TypeSystem).
		SetRole(systemuser.RoleAdmin).
		Save(ctx)
	if err != nil {
		log.Printf("[Seeder] Failed to seed system user: %v\n", err)
	}

	// Demo AccountPlan
	plan, err := client.AccountPlan.Create().
		SetID(uuid.New()).
		SetName("Demo").
		SetDescription("Plano demonstração para avaliação da plataforma.").
		SetPrice(0).
		SetMaxCustomers(100).
		SetMaxPlans(5).
		SetMaxProducts(20).
		SetActive(true).
		Save(ctx)
	if err != nil {
		log.Printf("[Seeder] Failed to seed account plan: %v\n", err)
		return
	}

	// Demo Account
	const demoSlug = "demo"
	_, err = client.Account.Create().
		SetID(uuid.New()).
		SetName("Demo").
		SetEmail("demo@subclub.com").
		SetDocument("00.000.000/0001-00").
		SetSlug(demoSlug).
		SetAccountPlanID(plan.ID).
		SetSubscriptionStatus("trial").
		SetActive(true).
		Save(ctx)
	if err != nil {
		log.Printf("[Seeder] Failed to seed demo account: %v\n", err)
		return
	}

	// Provision tenant schema + seed tenant data
	if err := manager.CreateTenantSchema(ctx, demoSlug); err != nil {
		log.Printf("[Seeder] Failed to provision demo tenant schema: %v\n", err)
		return
	}

	if err := manager.CreateTenantOwner(ctx, demoSlug, "demo@subclub.com", "Demo", "12345678"); err != nil {
		log.Printf("[Seeder] Failed to create demo tenant owner: %v\n", err)
	}

	log.Println("[Seeder] Done! Admin: adm@adm.com / 12345678 | Tenant: demo@subclub.com / 12345678")
}

// SeedTenant seeds a newly provisioned tenant schema with demo data.
// It is called automatically on development environments when a new account is created.
func SeedTenant(ctx context.Context, client *ent.Client, slug string) {
	log.Printf("[Seeder] Seeding tenant schema for %q...\n", slug)

	// 1. Products
	log.Println("[Seeder] Seeding coffee products (10)...")
	coffeeProducts := faker.CoffeeProducts()
	products := make([]*ent.Product, len(coffeeProducts))
	for i, f := range coffeeProducts {
		products[i], _ = client.Product.Create().
			SetID(uuid.New()).
			SetCode(f.Code).
			SetName(f.Name).
			SetDescription(f.Description).
			SetCostPrice(f.CostPrice).
			SetActive(true).
			Save(ctx)
	}

	// 2. Plans
	log.Println("[Seeder] Seeding plans (Básico, Intermediário, Avançado)...")
	fixedPlans := faker.FixedPlans()
	plans := make([]*ent.Plan, len(fixedPlans))
	for i, f := range fixedPlans {
		plans[i], _ = client.Plan.Create().
			SetID(uuid.New()).
			SetCode(f.Code).
			SetName(f.Name).
			SetDescription(f.Description).
			SetProductValue(f.ProductValue).
			SetDiscountValue(f.DiscountValue).
			SetPrice(f.Price).
			SetIntervalDays(f.IntervalDays).
			SetActive(true).
			Save(ctx)
	}

	// 3. Customers
	log.Println("[Seeder] Seeding fake customers (50)...")
	customers := make([]*ent.Customer, 50)
	for i := 0; i < 50; i++ {
		f := faker.FakeCustomer()
		customers[i], _ = client.Customer.Create().
			SetID(uuid.New()).
			SetName(f.Name).
			SetEmail(f.Email).
			SetPhone(f.Phone).
			SetDocument(f.Document).
			SetActive(true).
			Save(ctx)
	}

	// 4. Subscriptions (first 25 customers)
	log.Println("[Seeder] Seeding fake subscriptions (25)...")
	for i, customer := range customers {
		if customer == nil || i >= 25 {
			break
		}
		plan := plans[i%len(plans)]
		if plan != nil {
			_, _ = client.Subscription.Create().
				SetID(uuid.New()).
				SetCustomerID(customer.ID).
				SetPlanID(plan.ID).
				SetStatus("ACTIVE").
				Save(ctx)
		}
	}

	log.Printf("[Seeder] Tenant %q seeded with demo data.\n", slug)
}
