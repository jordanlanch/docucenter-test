package repository

import (
	"context"
	"errors"

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

func (r *warehousePortRepository) FindMany(ctx context.Context, pagination *domain.Pagination, filters map[string]interface{}) ([]*domain.WarehousesAndPorts, error) {
	var warehousePort []*domain.WarehousesAndPorts
	db := r.db.WithContext(ctx)
	db = applyPagination(db, pagination)

	for k, v := range filters {
		db = db.Where(k+" = ?", v)
	}

	result := db.Find(&warehousePort)
	if result.Error != nil {
		return nil, result.Error
	}
	return warehousePort, nil
}

func (r *warehousePortRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.WarehousesAndPorts, error) {
	var warehousePort domain.WarehousesAndPorts

	result := r.db.WithContext(ctx).First(&warehousePort, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &warehousePort, nil
}

func (r *warehousePortRepository) Store(ctx context.Context, wp *domain.WarehousesAndPorts) (*domain.WarehousesAndPorts, error) {
	result := r.db.WithContext(ctx).Create(wp)
	if result.Error != nil {
		return nil, result.Error
	}
	return wp, nil
}

func (r *warehousePortRepository) Update(ctx context.Context, wp *domain.WarehousesAndPorts, id uuid.UUID) (*domain.WarehousesAndPorts, error) {
	result := r.db.WithContext(ctx).Model(&domain.WarehousesAndPorts{}).Where("id = ?", id).Updates(wp)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("no warehouses_and_ports record found to update")
	}
	return wp, nil
}

func (r *warehousePortRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&domain.WarehousesAndPorts{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no warehouses_and_ports record found to delete")
	}
	return nil
}
