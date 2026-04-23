package service

import (
	"context"

	"github.com/joaofilippe/subclub/internal/domain/product"
	"github.com/joaofilippe/subclub/internal/domain/product/model"
	"github.com/joaofilippe/subclub/internal/domain/product/usecase"
)

type ProductService struct {
	createUseCase  *usecase.CreateProductUseCase
	getByIDUseCase *usecase.GetProductByIDUseCase
	updateUseCase  *usecase.UpdateProductUseCase
	deleteUseCase  *usecase.DeleteProductUseCase
	listUseCase    *usecase.ListProductsUseCase
}

func NewProductService(repo product.Repository) *ProductService {
	return &ProductService{
		createUseCase:  usecase.NewCreateProductUseCase(repo),
		getByIDUseCase: usecase.NewGetProductByIDUseCase(repo),
		updateUseCase:  usecase.NewUpdateProductUseCase(repo),
		deleteUseCase:  usecase.NewDeleteProductUseCase(repo),
		listUseCase:    usecase.NewListProductsUseCase(repo),
	}
}

func (s *ProductService) Create(ctx context.Context, input model.CreateProductInput) (*model.Product, error) {
	return s.createUseCase.Execute(ctx, input)
}

func (s *ProductService) GetByID(ctx context.Context, id string) (*model.Product, error) {
	return s.getByIDUseCase.Execute(ctx, id)
}

func (s *ProductService) Update(ctx context.Context, input model.UpdateProductInput) (*model.Product, error) {
	return s.updateUseCase.Execute(ctx, input)
}

func (s *ProductService) Delete(ctx context.Context, id string) error {
	return s.deleteUseCase.Execute(ctx, id)
}

func (s *ProductService) List(ctx context.Context, filter model.Filter) (*model.PaginatedList, error) {
	return s.listUseCase.Execute(ctx, filter)
}
