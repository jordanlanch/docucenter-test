package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const ProductTable = "products"

type Product struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"` // Terrestre o Marítima
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductRepository interface {
	FindMany(ctx context.Context, pagination *Pagination) ([]*Product, error)
	FindByID(ctx context.Context, id uuid.UUID) (*Product, error)
	Store(ctx context.Context, p *Product) (*Product, error)
	Update(ctx context.Context, p *Product) (*Product, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type ProductUsecase interface {
	GetMany(pagination *Pagination) ([]*Product, error)
	GetByID(id uuid.UUID) (*Product, error)
	Create(p *Product) (*Product, error)
	Modify(p *Product) (*Product, error)
	Remove(id uuid.UUID) error
}
