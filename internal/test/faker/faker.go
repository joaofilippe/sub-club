package faker

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	accountmodel "github.com/joaofilippe/subclub/internal/domain/account/model"
	accountplanmodel "github.com/joaofilippe/subclub/internal/domain/accountplan/model"
	customermodel "github.com/joaofilippe/subclub/internal/domain/customer/model"
	modulemodel "github.com/joaofilippe/subclub/internal/domain/module/model"
	planmodel "github.com/joaofilippe/subclub/internal/domain/plan/model"
	productmodel "github.com/joaofilippe/subclub/internal/domain/product/model"
	tenantuserdomain "github.com/joaofilippe/subclub/internal/domain/tenantuser/model"
	usermodel "github.com/joaofilippe/subclub/internal/domain/user/model"
)

// FakeModule generates a fake module model
func FakeModule() *modulemodel.Module {
	return &modulemodel.Module{
		ID:        uuid.New().String(),
		Name:      gofakeit.AppName(),
		Active:    true,
		CreatedAt: time.Now(),
	}
}

// FakeAccountPlan generates a fake account plan model
func FakeAccountPlan() *accountplanmodel.AccountPlan {
	return &accountplanmodel.AccountPlan{
		ID:           uuid.New().String(),
		Name:         gofakeit.JobTitle(),
		Description:  gofakeit.Sentence(8),
		Price:        gofakeit.Price(50, 500),
		MaxCustomers: gofakeit.Number(10, 1000),
		MaxPlans:     gofakeit.Number(1, 20),
		MaxProducts:  gofakeit.Number(1, 50),
		Active:       true,
		CreatedAt:    time.Now(),
	}
}

// FakeAccount generates a fake account model
func FakeAccount() *accountmodel.Account {
	return &accountmodel.Account{
		ID:                 uuid.New().String(),
		Name:               gofakeit.Company(),
		Email:              gofakeit.Email(),
		Document:           gofakeit.DigitN(14),
		Slug:               gofakeit.Lexify("????????"),
		SubscriptionStatus: accountmodel.SubscriptionStatusTrial,
		Active:             true,
		CreatedAt:          time.Now(),
	}
}

// FakeUser generates a fake system user model
func FakeUser() *usermodel.User {
	return &usermodel.User{
		ID:        uuid.New().String(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Password:  gofakeit.Password(true, true, true, true, false, 12),
		Type:      usermodel.UserTypeIndividual,
		Role:      usermodel.UserRoleOperations,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// FakeTenantUser generates a fake tenant user model
func FakeTenantUser() *tenantuserdomain.TenantUser {
	return &tenantuserdomain.TenantUser{
		ID:        uuid.New().String(),
		Name:      gofakeit.Name(),
		Username:  gofakeit.Username(),
		Email:     gofakeit.Email(),
		Role:      tenantuserdomain.RoleOperations,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// FakeCustomer generates a fake customer model
func FakeCustomer() *customermodel.Customer {
	return &customermodel.Customer{
		ID:       uuid.New().String(),
		Name:     gofakeit.Name(),
		Email:    gofakeit.Email(),
		Phone:    gofakeit.Phone(),
		Document: gofakeit.DigitN(11),
		Active:   true,
		Address: &customermodel.Address{
			ZipCode:      gofakeit.Zip(),
			Street:       gofakeit.Street(),
			Number:       gofakeit.DigitN(3),
			Neighborhood: gofakeit.City(),
			City:         gofakeit.City(),
			State:        gofakeit.State(),
		},
		CreatedAt: time.Now(),
	}
}

// FakeProduct generates a fake product model
func FakeProduct() *productmodel.Product {
	return &productmodel.Product{
		ID:          uuid.New().String(),
		Code:        gofakeit.AppName(),
		Name:        gofakeit.ProductName(),
		Description: gofakeit.Sentence(10),
		CostPrice:   gofakeit.Price(10, 200),
		Category:    gofakeit.ProductCategory(),
		Active:      true,
		CreatedAt:   time.Now(),
	}
}

// CoffeeProducts returns 10 fixed coffee products for seeding.
func CoffeeProducts() []*productmodel.Product {
	now := time.Now()
	items := []struct {
		code, name, description string
		costPrice               float64
	}{
		{"CAFE-EXPRESSO", "Café Expresso", "Grão arábica tostado escuro para extração concentrada em 25 ml.", 8.50},
		{"CAFE-AMERICANO", "Café Americano", "Espresso diluído em água quente, suave e levemente encorpado.", 7.50},
		{"CAFE-LEITE", "Café com Leite", "Blend suave de espresso com leite vaporizado na proporção 1:2.", 9.00},
		{"CAPPUCCINO", "Cappuccino", "Espresso com leite vaporizado e camada generosa de espuma cremosa.", 12.00},
		{"MACCHIATO", "Macchiato", "Espresso marcado com um toque de leite vaporizado.", 10.50},
		{"FLAT-WHITE", "Flat White", "Espresso duplo com leite vaporizado sedoso, sem espuma.", 13.00},
		{"CAFE-GELADO", "Café Gelado", "Espresso resfriado servido sobre gelo com leite opcional.", 11.00},
		{"MOCHA", "Mocha", "Espresso com calda de chocolate amargo e leite vaporizado.", 14.00},
		{"COLD-BREW", "Cold Brew", "Café extraído a frio por 12 horas, resultado suave e adocicado.", 15.00},
		{"CAFE-COADO", "Café Coado", "Método tradicional filtrado, torra média para xícara limpa e aromática.", 6.50},
	}

	products := make([]*productmodel.Product, len(items))
	for i, item := range items {
		products[i] = &productmodel.Product{
			ID:          uuid.New().String(),
			Code:        item.code,
			Name:        item.name,
			Description: item.description,
			CostPrice:   item.costPrice,
			Category:    "Café",
			Active:      true,
			CreatedAt:   now,
		}
	}
	return products
}

// FakePlan generates a fake plan model
func FakePlan() *planmodel.Plan {
	productValue := gofakeit.Price(50, 300)
	discount := gofakeit.Price(5, 50)

	return &planmodel.Plan{
		ID:            uuid.New().String(),
		Code:          gofakeit.UUID(),
		Name:          gofakeit.JobTitle(),
		Description:   gofakeit.Phrase(),
		ProductValue:  productValue,
		DiscountValue: discount,
		Price:         productValue - discount,
		IntervalDays:  30,
		Active:        true,
		CreatedAt:     time.Now(),
	}
}

// FixedPlans returns the 3 subscription tiers (Básico, Intermediário, Avançado) for seeding.
func FixedPlans() []*planmodel.Plan {
	now := time.Now()
	items := []struct {
		code, name, description string
		productValue, discount  float64
	}{
		{
			"PLANO-BASICO", "Básico",
			"Ideal para quem quer experimentar o clube. Receba 2 produtos por mês.",
			49.90, 5.00,
		},
		{
			"PLANO-INTERMEDIARIO", "Intermediário",
			"Para os apreciadores frequentes. Receba 4 produtos por mês com desconto especial.",
			89.90, 10.00,
		},
		{
			"PLANO-AVANCADO", "Avançado",
			"Experiência completa para entusiastas. Receba 8 produtos por mês e acesso a lançamentos.",
			149.90, 20.00,
		},
	}

	plans := make([]*planmodel.Plan, len(items))
	for i, item := range items {
		plans[i] = &planmodel.Plan{
			ID:            uuid.New().String(),
			Code:          item.code,
			Name:          item.name,
			Description:   item.description,
			ProductValue:  item.productValue,
			DiscountValue: item.discount,
			Price:         item.productValue - item.discount,
			IntervalDays:  30,
			Active:        true,
			CreatedAt:     now,
		}
	}
	return plans
}
