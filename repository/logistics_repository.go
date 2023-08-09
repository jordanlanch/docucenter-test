package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jordanlanch/docucenter-test/domain"
	"gorm.io/gorm"
)

type landLogisticsRepository struct {
	db    *gorm.DB
	table string
}

func NewLogisticsRepository(db *gorm.DB, table string) domain.LogisticsRepository {
	return &landLogisticsRepository{db, table}
}

func (r *landLogisticsRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Logistics, error) {
	var logistics domain.Logistics

	result := r.db.WithContext(ctx).First(&logistics, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &logistics, nil
}

func (r *landLogisticsRepository) Store(ctx context.Context, ll *domain.Logistics) error {
	result := r.db.WithContext(ctx).Create(ll)
	return result.Error
}

func (r *landLogisticsRepository) Update(ctx context.Context, ll *domain.Logistics) error {
	result := r.db.WithContext(ctx).Model(&domain.Logistics{}).Where("id = ?", ll.ID).Updates(ll)
	return result.Error
}

func (r *landLogisticsRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&domain.Logistics{}, id)
	return result.Error
}
