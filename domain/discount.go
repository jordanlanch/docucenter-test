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
	ID           uuid.UUID     `json:"id"`
	Type         LogisticsType `json:"type"`
	QuantityFrom int           `json:"quantity_from"`
	QuantityTo   int           `json:"quantity_to"`
	Percentage   float64       `json:"percentage"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}

type DiscountRepository interface {
	FindMany(ctx context.Context, pagination *Pagination) ([]*Discount, error)
	FindByID(ctx context.Context, id uuid.UUID) (*Discount, error)
	FindByTypeAndQuantity(ctx context.Context, dtype LogisticsType, quantity int) (*Discount, error)
	Store(ctx context.Context, d *Discount) (*Discount, error)
	Update(ctx context.Context, d *Discount) (*Discount, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type DiscountUsecase interface {
	GetMany(pagination *Pagination) ([]*Discount, error)
	GetByTypeAndQuantity(dtype LogisticsType, quantity int) (*Discount, error)
	Create(d *Discount) (*Discount, error)
	Modify(d *Discount) (*Discount, error)
	Remove(id uuid.UUID) error
}
