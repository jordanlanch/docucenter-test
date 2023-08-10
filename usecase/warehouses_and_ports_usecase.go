package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jordanlanch/docucenter-test/domain"
)

type warehousePortUsecase struct {
	warehousePortRepository domain.WarehousePortRepository
	contextTimeout          time.Duration
}

func NewWarehousePortUsecase(warehousePortRepository domain.WarehousePortRepository, timeout time.Duration) domain.WarehousePortUsecase {
	return &warehousePortUsecase{
		warehousePortRepository: warehousePortRepository,
		contextTimeout:          timeout,
	}
}

func (wpu *warehousePortUsecase) GetMany(pagination *domain.Pagination) ([]*domain.WarehousesAndPorts, error) {
	ctx, cancel := context.WithTimeout(context.Background(), wpu.contextTimeout)
	defer cancel()

	return wpu.warehousePortRepository.FindMany(ctx, pagination)
}

func (wpu *warehousePortUsecase) GetByID(id uuid.UUID) (*domain.WarehousesAndPorts, error) {
	ctx, cancel := context.WithTimeout(context.Background(), wpu.contextTimeout)
	defer cancel()

	return wpu.warehousePortRepository.FindByID(ctx, id)
}

func (wpu *warehousePortUsecase) Create(wp *domain.WarehousesAndPorts) (*domain.WarehousesAndPorts, error) {
	ctx, cancel := context.WithTimeout(context.Background(), wpu.contextTimeout)
	defer cancel()

	wp.ID = uuid.New() // Assign a new UUID to the warehouse port
	wp.CreatedAt = time.Now()
	wp.UpdatedAt = time.Now()

	return wpu.warehousePortRepository.Store(ctx, wp)
}

func (wpu *warehousePortUsecase) Modify(wp *domain.WarehousesAndPorts, id uuid.UUID) (*domain.WarehousesAndPorts, error) {
	ctx, cancel := context.WithTimeout(context.Background(), wpu.contextTimeout)
	defer cancel()

	existingWarehousePort, err := wpu.warehousePortRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update fields
	existingWarehousePort.Name = wp.Name
	existingWarehousePort.Type = wp.Type
	existingWarehousePort.Location = wp.Location
	existingWarehousePort.UpdatedAt = time.Now()

	return wpu.warehousePortRepository.Update(ctx, existingWarehousePort, id)
}

func (wpu *warehousePortUsecase) Remove(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), wpu.contextTimeout)
	defer cancel()

	return wpu.warehousePortRepository.Delete(ctx, id)
}
