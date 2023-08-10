package repository

import (
	"context"
	"errors"

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

// applyPagination aplica la paginación a la consulta de base de datos.
func applyPagination(db *gorm.DB, pagination *domain.Pagination) *gorm.DB {
	if pagination != nil {
		if pagination.Limit != nil {
			db = db.Limit(*pagination.Limit)
		}
		if pagination.Offset != nil {
			db = db.Offset(*pagination.Offset)
		}
	}
	return db
}

func (r *customerRepository) FindMany(ctx context.Context, pagination *domain.Pagination) ([]*domain.Customer, error) {
	var customer []*domain.Customer
	db := r.db.WithContext(ctx)
	db = applyPagination(db, pagination)

	result := db.Find(&customer)
	if result.Error != nil {
		return nil, result.Error
	}
	return customer, nil
}

func (r *customerRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Customer, error) {
	var customer domain.Customer

	result := r.db.WithContext(ctx).First(&customer, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &customer, nil
}

func (r *customerRepository) Store(ctx context.Context, c *domain.Customer) (*domain.Customer, error) {
	result := r.db.WithContext(ctx).Create(c)
	if result.Error != nil {
		return nil, result.Error
	}
	return c, nil
}

func (r *customerRepository) Update(ctx context.Context, c *domain.Customer, id uuid.UUID) (*domain.Customer, error) {
	result := r.db.WithContext(ctx).Model(&domain.Customer{}).Where("id = ?", c.ID).Updates(c)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("no customer record found to update")
	}
	return c, nil
}

func (r *customerRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&domain.Customer{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no customer record found to delete")
	}
	return nil
}
