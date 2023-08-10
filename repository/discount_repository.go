package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jordanlanch/docucenter-test/domain"
	"gorm.io/gorm"
)

type discountRepository struct {
	db    *gorm.DB
	table string
}

func NewDiscountRepository(db *gorm.DB, table string) domain.DiscountRepository {
	return &discountRepository{db, table}
}

func (r *discountRepository) FindMany(ctx context.Context, pagination *domain.Pagination) ([]*domain.Discount, error) {
	var discount []*domain.Discount
	db := r.db.WithContext(ctx)
	db = applyPagination(db, pagination)

	result := db.Find(&discount)
	if result.Error != nil {
		return nil, result.Error
	}
	return discount, nil
}

func (r *discountRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Discount, error) {
	var discount domain.Discount
	result := r.db.WithContext(ctx).Where("id = ?", id).First(&discount)
	if result.Error != nil {
		return nil, result.Error
	}
	return &discount, nil
}

func (r *discountRepository) FindByTypeAndQuantity(ctx context.Context, dtype domain.LogisticsType, quantity int) (*domain.Discount, error) {
	var discount domain.Discount
	result := r.db.WithContext(ctx).Where("type = ? AND quantity_from <= ? AND quantity_to >= ?", dtype, quantity, quantity).First(&discount)
	if result.Error != nil {
		return nil, result.Error
	}
	return &discount, nil
}

func (r *discountRepository) Store(ctx context.Context, d *domain.Discount) (*domain.Discount, error) {
	result := r.db.WithContext(ctx).Create(d)
	if result.Error != nil {
		return nil, result.Error
	}
	return d, nil
}

func (r *discountRepository) Update(ctx context.Context, d *domain.Discount, id uuid.UUID) (*domain.Discount, error) {
	result := r.db.WithContext(ctx).Model(&domain.Discount{}).Where("id = ?", id).Updates(d)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("no discounts record found to update")
	}
	return d, nil
}

func (r *discountRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&domain.Discount{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no discounts record found to delete")
	}
	return nil
}
