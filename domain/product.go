package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const ProductTable = "products"

type Product struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name" jsonschema:"required,description=Name of the product,title=Name"`
	Type      string    `json:"type" jsonschema:"required,description=Type of product (land or maritime),title=Type"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type ProductRepository interface {
	FindMany(ctx context.Context, pagination *Pagination, filters map[string]interface{}) ([]*Product, error)
	FindByID(ctx context.Context, id uuid.UUID) (*Product, error)
	Store(ctx context.Context, p *Product) (*Product, error)
	Update(ctx context.Context, p *Product, id uuid.UUID) (*Product, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type ProductUsecase interface {
	GetMany(pagination *Pagination, filters map[string]interface{}) ([]*Product, error)
	GetByID(id uuid.UUID) (*Product, error)
	Create(p *Product) (*Product, error)
	Modify(p *Product, id uuid.UUID) (*Product, error)
	Remove(id uuid.UUID) error
}
