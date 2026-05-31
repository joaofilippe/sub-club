package database

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/ent"
	"github.com/joaofilippe/subclub/ent/user"
	"github.com/joaofilippe/subclub/internal/config"
	"github.com/joaofilippe/subclub/internal/test/faker"
	"golang.org/x/crypto/bcrypt"
)

func SeedAll(ctx context.Context, client *ent.Client, cfg *config.Config) {
	if !cfg.IsDevelopment() {
		log.Printf("[Seeder] Skipping: current environment is %q\n", cfg.AppEnv)
		return
	}

	log.Println("[Seeder] Checking database state...")

	// Verify if there are any users in the database
	count, err := client.User.Query().Count(ctx)
	if err != nil {
		log.Printf("[Seeder] Error counting users: %v\n", err)
		return
	}

	if count > 0 {
		log.Println("[Seeder] Database already populated. Skipping seeds.")
		return
	}

	log.Println("[Seeder] Database is empty. Starting seed process...")

	// 1. Seed Admin User
	hash, err := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[Seeder] Error hashing password: %v\n", err)
		return
	}

	_, err = client.User.Create().
		SetID(uuid.New()).
		SetName("Admin").
		SetEmail("adm@adm.com").
		SetPassword(string(hash)).
		SetType(user.TypeSystem).
		SetRole(user.RoleAdmin).
		Save(ctx)

	if err != nil {
		log.Printf("[Seeder] Failed to seed user: %v\n", err)
	}

	// 2. Seed Coffee Products
	log.Println("[Seeder] Seeding coffee products (10)...")
	coffeeProducts := faker.CoffeeProducts()
	products := make([]*ent.Product, len(coffeeProducts))
	for i, f := range coffeeProducts {
		products[i], _ = client.Product.Create().
			SetID(uuid.MustParse(f.ID)).
			SetCode(f.Code).
			SetName(f.Name).
			SetDescription(f.Description).
			SetCostPrice(f.CostPrice).
			SetActive(true).
			Save(ctx)
	}

	// 3. Seed Fixed Plans (Básico, Intermediário, Avançado)
	log.Println("[Seeder] Seeding plans (Básico, Intermediário, Avançado)...")
	fixedPlans := faker.FixedPlans()
	plans := make([]*ent.Plan, len(fixedPlans))
	for i, f := range fixedPlans {
		plans[i], _ = client.Plan.Create().
			SetID(uuid.MustParse(f.ID)).
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

	// 4. Seed Fake Customers (Clients)
	log.Println("[Seeder] Seeding fake customers (50)...")
	customers := make([]*ent.Customer, 50)
	for i := 0; i < 50; i++ {
		f := faker.FakeCustomer()
		customers[i], _ = client.Customer.Create().
			SetID(uuid.MustParse(f.ID)).
			SetName(f.Name).
			SetEmail(f.Email).
			SetPhone(f.Phone).
			SetDocument(f.Document).
			SetActive(true).
			Save(ctx)
	}

	// 5. Seed Subscriptions for some customers
	log.Println("[Seeder] Seeding fake subscriptions (25)...")
	for i, customer := range customers {
		if customer == nil {
			continue
		}
		// Subscribe first 25 customers to random plans
		if i < 25 {
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
	}

	log.Println("[Seeder] Mocks successfully inserted! Root Admin: adm@adm.com / 12345678")
}
