package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jordanlanch/docucenter-test/domain"
	"gorm.io/gorm"
)

type customerRepository struct {
	db    *gorm.DB
	table string
}

func NewCustomerRepository(db *gorm.DB, table string) domain.CustomerRepository {
	return &customerRepository{db, table}
}

func (r *customerRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Customer, error) {
	var customer domain.Customer

	result := r.db.WithContext(ctx).First(&customer, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}

func (r *customerRepository) Store(ctx context.Context, c *domain.Customer) error {
	result := r.db.WithContext(ctx).Create(c)
	return result.Error
}

func (r *customerRepository) Update(ctx context.Context, c *domain.Customer) error {
	result := r.db.WithContext(ctx).Model(&domain.Customer{}).Where("id = ?", c.ID).Updates(c)
	return result.Error
}

func (r *customerRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&domain.Customer{}, id)
	return result.Error
}
