package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jordanlanch/docucenter-test/domain"
	"gorm.io/gorm"
)

type LogisticsRepository struct {
	db    *gorm.DB
	table string
}

func NewLogisticsRepository(db *gorm.DB, table string) domain.LogisticsRepository {
	return &LogisticsRepository{db, table}
}

func (r *LogisticsRepository) FindMany(ctx context.Context, pagination *domain.Pagination) ([]*domain.Logistics, error) {
	var logistics []*domain.Logistics
	db := r.db.WithContext(ctx)
	db = applyPagination(db, pagination)

	result := db.Find(&logistics)
	if result.Error != nil {
		return nil, result.Error
	}
	return logistics, nil
}

func (r *LogisticsRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Logistics, error) {
	var logistics domain.Logistics

	result := r.db.WithContext(ctx).First(&logistics, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &logistics, nil
}

func (r *LogisticsRepository) Store(ctx context.Context, ll *domain.Logistics) (*domain.Logistics, error) {
	result := r.db.WithContext(ctx).Create(ll)
	if result.Error != nil {
		return nil, result.Error
	}
	return ll, nil
}

func (r *LogisticsRepository) Update(ctx context.Context, ll *domain.Logistics) (*domain.Logistics, error) {
	result := r.db.WithContext(ctx).Model(&domain.Logistics{}).Where("id = ?", ll.ID).Updates(ll)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("no logistics record found to update")
	}
	return ll, nil
}

func (r *LogisticsRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&domain.Logistics{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no logistics record found to delete")
	}
	return nil
}
