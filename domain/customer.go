package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const CustomerTable = "customers"

type Customer struct {
	ID        uuid.UUID `json:"id,omitempty" jsonschema:"description=ID,title=ID"`
	Name      string    `json:"name" jsonschema:"required,description=Name of the customer,title=Name"`
	Email     string    `json:"email" jsonschema:"required,description=Email of the customer,title=Email"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type CustomerRepository interface {
	FindMany(ctx context.Context, pagination *Pagination, filters map[string]interface{}) ([]*Customer, error)
	FindByID(ctx context.Context, id uuid.UUID) (*Customer, error)
	Store(ctx context.Context, c *Customer) (*Customer, error)
	Update(ctx context.Context, c *Customer, id uuid.UUID) (*Customer, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type CustomerUsecase interface {
	GetMany(pagination *Pagination, filters map[string]interface{}) ([]*Customer, error)
	GetByID(id uuid.UUID) (*Customer, error)
	Create(c *Customer) (*Customer, error)
	Modify(c *Customer, id uuid.UUID) (*Customer, error)
	Remove(id uuid.UUID) error
}
