package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jordanlanch/docucenter-test/domain"
	"gorm.io/gorm"
)

type warehousePortRepository struct {
	db    *gorm.DB
	table string
}

func NewWarehousePortRepository(db *gorm.DB, table string) domain.WarehousePortRepository {
	return &warehousePortRepository{db, table}
}

func (r *warehousePortRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.WarehousesAndPorts, error) {
	var warehousePort domain.WarehousesAndPorts

	result := r.db.WithContext(ctx).First(&warehousePort, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &warehousePort, nil
}

func (r *warehousePortRepository) Store(ctx context.Context, wp *domain.WarehousesAndPorts) error {
	result := r.db.WithContext(ctx).Create(wp)
	return result.Error
}

func (r *warehousePortRepository) Update(ctx context.Context, wp *domain.WarehousesAndPorts) error {
	result := r.db.WithContext(ctx).Model(&domain.WarehousesAndPorts{}).Where("id = ?", wp.ID).Updates(wp)
	return result.Error
}

func (r *warehousePortRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&domain.WarehousesAndPorts{}, id)
	return result.Error
}
