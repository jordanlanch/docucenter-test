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
	ID                    uuid.UUID     `json:"id"`
	ProductID             uuid.UUID     `json:"product_id"`
	WarehousePortID       uuid.UUID     `json:"warehouse_port_id"`
	Type                  LogisticsType `json:"type"`
	Quantity              int           `json:"quantity"`
	RegistrationDate      time.Time     `json:"registration_date"`
	DeliveryDate          time.Time     `json:"delivery_date"`
	ShippingPrice         float64       `json:"shipping_price"`
	DiscountShippingPrice float64       `json:"discount_shipping_price"`
	VehiclePlate          string        `json:"vehicle_plate,omitempty"`
	FleetNumber           string        `json:"fleet_number,omitempty"`
	GuideNumber           string        `json:"guide_number"`
	DiscountPercent       float64       `json:"discount_percent"`
	CreatedAt             time.Time     `json:"created_at"`
	UpdatedAt             time.Time     `json:"updated_at"`
}

func (ll *Logistics) Validate() error {
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
	FindByID(ctx context.Context, id uuid.UUID) (*Logistics, error)
	Store(ctx context.Context, ll *Logistics) error
	Update(ctx context.Context, ll *Logistics) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type LogisticsUsecase interface {
	GetByID(id uuid.UUID) (*Logistics, error)
	Create(ll *Logistics) error
	Modify(ll *Logistics) error
	Remove(id uuid.UUID) error
}
