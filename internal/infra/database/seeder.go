package database

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/joaofilippe/subclub/ent"
	"github.com/joaofilippe/subclub/ent/user"
	"golang.org/x/crypto/bcrypt"
)

// SeedAll populates the database with initial mock data if the database is empty.
func SeedAll(ctx context.Context, client *ent.Client) {
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
	hash, err := bcrypt.GenerateFromPassword([]byte("root"), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[Seeder] Error hashing password: %v\n", err)
		return
	}

	_, err = client.User.Create().
		SetID(uuid.New()).
		SetName("Admin Supreme").
		SetEmail("admin@subclub.com").
		SetPassword(string(hash)).
		SetType(user.TypeSystem).
		SetRole(user.RoleAdmin).
		Save(ctx)

	if err != nil {
		log.Printf("[Seeder] Failed to seed user: %v\n", err)
	}

	// 2. Seed Products
	prodBase, _ := client.Product.Create().
		SetID(uuid.New()).
		SetCode("COFFEE-BASE").
		SetName("Café Torrado Base").
		SetDescription("Café diário de excelente custo-benefício.").
		SetCostPrice(25.90).
		SetActive(true).
		Save(ctx)

	prodPremium, _ := client.Product.Create().
		SetID(uuid.New()).
		SetCode("COFFEE-PREMIUM").
		SetName("Café Especial Altitude").
		SetDescription("Graos 100% arábica do sul de minas.").
		SetCostPrice(55.00).
		SetActive(true).
		Save(ctx)

	// 3. Seed Plans
	planBasic, _ := client.Plan.Create().
		SetID(uuid.New()).
		SetCode("PLAN-BASIC").
		SetName("Assinatura Mensal Diária").
		SetDescription("Receba nosso café diário em sua casa todos os meses.").
		SetProductValue(25.90).
		SetDiscountValue(1.90).
		SetPrice(24.00).
		SetIntervalDays(30).
		SetActive(true).
		Save(ctx)

	planPro, _ := client.Plan.Create().
		SetID(uuid.New()).
		SetCode("PLAN-PRO").
		SetName("Assinatura Premium Trimestral").
		SetDescription("Cafés premium a cada 3 meses com desconto agressivo.").
		SetProductValue(165.00).
		SetDiscountValue(15.00).
		SetPrice(150.00).
		SetIntervalDays(90).
		SetActive(true).
		Save(ctx)

	// 4. Seed Customers (Clients)
	customer1, _ := client.Customer.Create().
		SetID(uuid.New()).
		SetName("João da Silva").
		SetEmail("joaocustomer@example.com").
		SetPhone("11999999999").
		SetDocument("12345678909").
		SetActive(true).
		Save(ctx)

	customer2, _ := client.Customer.Create().
		SetID(uuid.New()).
		SetName("Maria Joaquina").
		SetEmail("maria.vip@example.com").
		SetPhone("21988888888").
		SetDocument("98765432100").
		SetActive(true).
		Save(ctx)

	// 5. Seed Subscriptions
	if customer1 != nil && planBasic != nil {
		_, _ = client.Subscription.Create().
			SetID(uuid.New()).
			SetClientID(customer1.ID).
			SetPlanID(planBasic.ID).
			SetStatus("ACTIVE").
			Save(ctx)
	}

	if customer2 != nil && planPro != nil {
		_, _ = client.Subscription.Create().
			SetID(uuid.New()).
			SetClientID(customer2.ID).
			SetPlanID(planPro.ID).
			SetStatus("ACTIVE").
			Save(ctx)
	}

	log.Println("[Seeder] Mocks successfully inserted! Root Admin: admin@subclub.com / root")
	// Evitar erro compiler "declare and not use" se falhar silenciosamente
	_ = prodBase
	_ = prodPremium
}
