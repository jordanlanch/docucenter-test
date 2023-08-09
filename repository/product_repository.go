package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jordanlanch/docucenter-test/domain"
	"gorm.io/gorm"
)

type productRepository struct {
	db    *gorm.DB
	table string
}

func NewProductRepository(db *gorm.DB, table string) domain.ProductRepository {
	return &productRepository{db, table}
}

func (r *productRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
	var product domain.Product

	result := r.db.WithContext(ctx).First(&product, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (r *productRepository) Store(ctx context.Context, p *domain.Product) error {
	result := r.db.WithContext(ctx).Create(p)
	return result.Error
}

func (r *productRepository) Update(ctx context.Context, p *domain.Product) error {
	result := r.db.WithContext(ctx).Model(&domain.Product{}).Where("id = ?", p.ID).Updates(p)
	return result.Error
}

func (r *productRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&domain.Product{}, id)
	return result.Error
}
