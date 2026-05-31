package faker

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	accountmodel "github.com/joaofilippe/subclub/internal/domain/account/model"
	accountplanmodel "github.com/joaofilippe/subclub/internal/domain/accountplan/model"
	customermodel "github.com/joaofilippe/subclub/internal/domain/customer/model"
	planmodel "github.com/joaofilippe/subclub/internal/domain/plan/model"
	productmodel "github.com/joaofilippe/subclub/internal/domain/product/model"
	usermodel "github.com/joaofilippe/subclub/internal/domain/user/model"
)

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

// FakeUser generates a fake user model
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
