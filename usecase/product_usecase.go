package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jordanlanch/docucenter-test/domain"
)

type productUsecase struct {
	productRepository domain.ProductRepository
	contextTimeout    time.Duration
}

func NewProductUsecase(productRepository domain.ProductRepository, timeout time.Duration) domain.ProductUsecase {
	return &productUsecase{
		productRepository: productRepository,
		contextTimeout:    timeout,
	}
}

func (pu *productUsecase) GetMany(pagination *domain.Pagination) ([]*domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), pu.contextTimeout)
	defer cancel()

	return pu.productRepository.FindMany(ctx, pagination)
}

func (pu *productUsecase) GetByID(id uuid.UUID) (*domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), pu.contextTimeout)
	defer cancel()

	return pu.productRepository.FindByID(ctx, id)
}

func (pu *productUsecase) Create(p *domain.Product) (*domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), pu.contextTimeout)
	defer cancel()

	p.ID = uuid.New() // Assign a new UUID to the product
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	return pu.productRepository.Store(ctx, p)
}

func (pu *productUsecase) Modify(p *domain.Product, id uuid.UUID) (*domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), pu.contextTimeout)
	defer cancel()

	existingProduct, err := pu.productRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update fields
	existingProduct.Name = p.Name
	existingProduct.Type = p.Type
	existingProduct.UpdatedAt = time.Now()

	return pu.productRepository.Update(ctx, existingProduct, id)
}

func (pu *productUsecase) Remove(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), pu.contextTimeout)
	defer cancel()

	return pu.productRepository.Delete(ctx, id)
}
