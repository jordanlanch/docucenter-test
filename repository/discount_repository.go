package repository

import (
	"context"

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

func (r *discountRepository) FindByType(ctx context.Context, dtype domain.LogisticsType) (*domain.Discount, error) {
	var discount domain.Discount

	result := r.db.WithContext(ctx).First(&discount, "type = ?", dtype)
	if result.Error != nil {
		return nil, result.Error
	}
	return &discount, nil
}

func (r *discountRepository) Store(ctx context.Context, d *domain.Discount) error {
	result := r.db.WithContext(ctx).Create(d)
	return result.Error
}

func (r *discountRepository) Update(ctx context.Context, d *domain.Discount) error {
	result := r.db.WithContext(ctx).Model(&domain.Discount{}).Where("id = ?", d.ID).Updates(d)
	return result.Error
}

func (r *discountRepository) Delete(ctx context.Context, dtype domain.LogisticsType) error {
	result := r.db.WithContext(ctx).Delete(&domain.Discount{}, "type = ?", dtype)
	return result.Error
}
