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
	ID         uuid.UUID     `json:"id"`
	Type       LogisticsType `json:"type"` // land or maritime
	Quantity   int           `json:"quantity"`
	Percentage float64       `json:"percentage"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at"`
}

type DiscountRepository interface {
	FindByType(ctx context.Context, dtype LogisticsType) (*Discount, error)
	Store(ctx context.Context, d *Discount) error
	Update(ctx context.Context, d *Discount) error
	Delete(ctx context.Context, dtype LogisticsType) error
}

type DiscountUsecase interface {
	GetByType(dtype LogisticsType) (*Discount, error)
	Create(d *Discount) error
	Modify(d *Discount) error
	Remove(dtype LogisticsType) error
}
