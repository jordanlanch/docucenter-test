package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const WarehousePortTable = "warehouses_and_ports"

type WarehousesAndPorts struct {
	ID        uuid.UUID `json:"id"`
	Type      string    `json:"type"` // land or maritime
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WarehousePortRepository interface {
	FindMany(ctx context.Context, pagination *Pagination) ([]*WarehousesAndPorts, error)
	FindByID(ctx context.Context, id uuid.UUID) (*WarehousesAndPorts, error)
	Store(ctx context.Context, wp *WarehousesAndPorts) (*WarehousesAndPorts, error)
	Update(ctx context.Context, wp *WarehousesAndPorts) (*WarehousesAndPorts, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type WarehousePortUsecase interface {
	GetMany(pagination *Pagination) ([]*WarehousesAndPorts, error)
	GetByID(id uuid.UUID) (*WarehousesAndPorts, error)
	Create(wp *WarehousesAndPorts) (*WarehousesAndPorts, error)
	Modify(wp *WarehousesAndPorts) (*WarehousesAndPorts, error)
	Remove(id uuid.UUID) error
}
