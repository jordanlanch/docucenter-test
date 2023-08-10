package repository

import (
	"context"
	"errors"

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

func (r *productRepository) FindMany(ctx context.Context, pagination *domain.Pagination, filters map[string]interface{}) ([]*domain.Product, error) {
	var product []*domain.Product
	db := r.db.WithContext(ctx)
	db = applyPagination(db, pagination)

	for k, v := range filters {
		db = db.Where(k+" = ?", v)
	}

	result := db.Find(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (r *productRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Product, error) {
	var product domain.Product

	result := r.db.WithContext(ctx).First(&product, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (r *productRepository) Store(ctx context.Context, p *domain.Product) (*domain.Product, error) {
	result := r.db.WithContext(ctx).Create(p)
	if result.Error != nil {
		return nil, result.Error
	}
	return p, nil
}

func (r *productRepository) Update(ctx context.Context, p *domain.Product, id uuid.UUID) (*domain.Product, error) {
	result := r.db.WithContext(ctx).Model(&domain.Product{}).Where("id = ?", id).Updates(p)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("no product record found to update")
	}
	return p, nil
}

func (r *productRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&domain.Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no product record found to delete")
	}
	return nil
}
