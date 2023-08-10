package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type LogisticsType string

const (
	Land     LogisticsType = "land"
	Maritime LogisticsType = "maritime"
)

const DiscountTable = "discounts"

type Discount struct {
	ID           uuid.UUID     `json:"id,omitempty"`
	Type         LogisticsType `json:"type" jsonschema:"required,description=Type of logistics,title=Type"`
	QuantityFrom int           `json:"quantity_from" jsonschema:"required,description=Starting quantity for discount,title=Quantity From"`
	QuantityTo   int           `json:"quantity_to" jsonschema:"required,description=Ending quantity for discount,title=Quantity To"`
	Percentage   float64       `json:"percentage" jsonschema:"required,description=Discount percentage,title=Percentage"`
	CreatedAt    time.Time     `json:"created_at,omitempty"`
	UpdatedAt    time.Time     `json:"updated_at,omitempty"`
}

type DiscountRepository interface {
	FindMany(ctx context.Context, pagination *Pagination) ([]*Discount, error)
	FindByID(ctx context.Context, id uuid.UUID) (*Discount, error)
	FindByTypeAndQuantity(ctx context.Context, dtype LogisticsType, quantity int) (*Discount, error)
	Store(ctx context.Context, d *Discount) (*Discount, error)
	Update(ctx context.Context, d *Discount, id uuid.UUID) (*Discount, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type DiscountUsecase interface {
	GetMany(pagination *Pagination) ([]*Discount, error)
	GetByTypeAndQuantity(dtype LogisticsType, quantity int) (*Discount, error)
	Create(d *Discount) (*Discount, error)
	Modify(d *Discount, id uuid.UUID) (*Discount, error)
	Remove(id uuid.UUID) error
}
