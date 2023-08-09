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
	FindByID(ctx context.Context, id uuid.UUID) (*Product, error)
	Store(ctx context.Context, p *Product) error
	Update(ctx context.Context, p *Product) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type ProductUsecase interface {
	GetByID(id uuid.UUID) (*Product, error)
	Create(p *Product) error
	Modify(p *Product) error
	Remove(id uuid.UUID) error
}
