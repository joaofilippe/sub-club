package faker

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	clientmodel "github.com/joaofilippe/subclub/internal/domain/client/model"
	planmodel "github.com/joaofilippe/subclub/internal/domain/plan/model"
	productmodel "github.com/joaofilippe/subclub/internal/domain/product/model"
	usermodel "github.com/joaofilippe/subclub/internal/domain/user/model"
)

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

// FakeClient generates a fake client (customer) model
func FakeClient() *clientmodel.Client {
	return &clientmodel.Client{
		ID:       uuid.New().String(),
		Name:     gofakeit.Name(),
		Email:    gofakeit.Email(),
		Phone:    gofakeit.Phone(),
		Document: gofakeit.DigitN(11), // Simple CPF-like
		Active:   true,
		Address: &clientmodel.Address{
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
