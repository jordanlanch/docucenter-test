package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const CustomerTable = "customers"

type Customer struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CustomerRepository interface {
	FindMany(ctx context.Context, pagination *Pagination) ([]*Customer, error)
	FindByID(ctx context.Context, id uuid.UUID) (*Customer, error)
	Store(ctx context.Context, c *Customer) (*Customer, error)
	Update(ctx context.Context, c *Customer) (*Customer, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type CustomerUsecase interface {
	GetMany(pagination *Pagination) ([]*Customer, error)
	GetByID(id uuid.UUID) (*Customer, error)
	Create(c *Customer) (*Customer, error)
	Modify(c *Customer) (*Customer, error)
	Remove(id uuid.UUID) error
}
