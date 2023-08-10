package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const WarehousePortTable = "warehouses_and_ports"

type WarehousesAndPorts struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name" jsonschema:"required,description=Name of the warehouse or port,title=Name"`
	Type      string    `json:"type" jsonschema:"required,description=Type of the location (land or maritime),title=Type"`
	Location  string    `json:"location" jsonschema:"required,description=Location of the warehouse or port,title=Location"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type WarehousePortRepository interface {
	FindMany(ctx context.Context, pagination *Pagination) ([]*WarehousesAndPorts, error)
	FindByID(ctx context.Context, id uuid.UUID) (*WarehousesAndPorts, error)
	Store(ctx context.Context, wp *WarehousesAndPorts) (*WarehousesAndPorts, error)
	Update(ctx context.Context, wp *WarehousesAndPorts, id uuid.UUID) (*WarehousesAndPorts, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type WarehousePortUsecase interface {
	GetMany(pagination *Pagination) ([]*WarehousesAndPorts, error)
	GetByID(id uuid.UUID) (*WarehousesAndPorts, error)
	Create(wp *WarehousesAndPorts) (*WarehousesAndPorts, error)
	Modify(wp *WarehousesAndPorts, id uuid.UUID) (*WarehousesAndPorts, error)
	Remove(id uuid.UUID) error
}
