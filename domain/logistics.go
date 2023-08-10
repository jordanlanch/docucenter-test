package domain

import (
	"context"
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
)

const LogisticsTable = "logistics"

type Logistics struct {
	ID                    uuid.UUID     `json:"id,omitempty"`
	CustomerID            string        `json:"customer_id" jsonschema:"required,description=ID of the customer,title=Customer ID,type=string"`
	ProductID             string        `json:"product_id" jsonschema:"required,description=ID of the product,title=Product ID,type=string"`
	WarehousePortID       string        `json:"warehouse_port_id" jsonschema:"required,description=ID of the warehouse port,title=Warehouse Port ID,type=string"`
	Type                  LogisticsType `json:"type" jsonschema:"required,description=Type of logistics,title=Type"`
	Quantity              int           `json:"quantity" jsonschema:"required,description=Quantity of items,title=Quantity"`
	RegistrationDate      time.Time     `json:"registration_date" jsonschema:"required,description=Date of registration,title=Registration Date"`
	DeliveryDate          time.Time     `json:"delivery_date" jsonschema:"required,description=Date of delivery,title=Delivery Date"`
	ShippingPrice         float64       `json:"shipping_price" jsonschema:"required,description=Price of shipping,title=Shipping Price"`
	DiscountShippingPrice float64       `json:"discount_shipping_price,omitempty"`
	VehiclePlate          string        `json:"vehicle_plate,omitempty"`
	FleetNumber           string        `json:"fleet_number,omitempty"`
	GuideNumber           string        `json:"guide_number" jsonschema:"required,description=Guide number,title=Guide Number"`
	DiscountPercent       float64       `json:"discount_percent,omitempty"`
	CreatedAt             time.Time     `json:"created_at,omitempty"`
	UpdatedAt             time.Time     `json:"updated_at,omitempty"`
}

func (ll *Logistics) Validate() error {
	// Validación para campos UUID vacíos
	if ll.CustomerID == "" {
		return errors.New("customer_id is required")
	}
	if ll.ProductID == "" {
		return errors.New("product_id is required")
	}
	if ll.WarehousePortID == "" {
		return errors.New("warehouse_port_id is required")
	}

	if ll.Type == "land" {
		if matched, _ := regexp.MatchString(`^[A-Za-z]{3}\d{3}$`, ll.VehiclePlate); !matched {
			return errors.New("invalid vehicle plate format. Expected format: XXX123")
		}
	}
	if ll.Type == "maritime" {
		if matched, _ := regexp.MatchString(`^[A-Za-z]{3}\d{4}[A-Za-z]$`, ll.FleetNumber); !matched {
			return errors.New("invalid fleet number format. Expected format: XXX1234X")
		}
	}
	return nil
}

type LogisticsRepository interface {
	FindMany(ctx context.Context, pagination *Pagination) ([]*Logistics, error)
	FindByID(ctx context.Context, id uuid.UUID) (*Logistics, error)
	Store(ctx context.Context, ll *Logistics) (*Logistics, error)
	Update(ctx context.Context, ll *Logistics, id uuid.UUID) (*Logistics, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type LogisticsUsecase interface {
	GetMany(pagination *Pagination) ([]*Logistics, error)
	GetByID(id uuid.UUID) (*Logistics, error)
	Create(ll *Logistics) (*Logistics, error)
	Modify(ll *Logistics, id uuid.UUID) (*Logistics, error)
	Remove(id uuid.UUID) error
}
