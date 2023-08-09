package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jordanlanch/docucenter-test/domain"
)

type landLogisticsUsecase struct {
	landLogisticsRepository domain.LogisticsRepository
	discountRepository      domain.DiscountRepository
	contextTimeout          time.Duration
}

func NewLogisticsUsecase(landLogisticsRepository domain.LogisticsRepository, discountRepository domain.DiscountRepository, timeout time.Duration) domain.LogisticsUsecase {
	return &landLogisticsUsecase{
		landLogisticsRepository: landLogisticsRepository,
		discountRepository:      discountRepository,
		contextTimeout:          timeout,
	}
}

func (llu *landLogisticsUsecase) GetByID(id uuid.UUID) (*domain.Logistics, error) {
	ctx, cancel := context.WithTimeout(context.Background(), llu.contextTimeout)
	defer cancel()

	return llu.landLogisticsRepository.FindByID(ctx, id)
}

func (llu *landLogisticsUsecase) Create(ll *domain.Logistics) error {
	ctx, cancel := context.WithTimeout(context.Background(), llu.contextTimeout)
	defer cancel()

	ll.ID = uuid.New() // Assign a new UUID to the land logistics entry
	ll.CreatedAt = time.Now()
	ll.UpdatedAt = time.Now()

	// Calculate discount based on quantity
	discount, err := llu.discountRepository.FindByType(ctx, ll.Type)
	if err != nil {
		return err
	}

	if ll.Quantity > discount.Quantity {
		ll.DiscountShippingPrice = discount.Percentage / 100 * ll.ShippingPrice
		ll.DiscountPercent = discount.Percentage
	}

	return llu.landLogisticsRepository.Store(ctx, ll)
}

func (llu *landLogisticsUsecase) Modify(ll *domain.Logistics) error {
	ctx, cancel := context.WithTimeout(context.Background(), llu.contextTimeout)
	defer cancel()

	existingLogistics, err := llu.landLogisticsRepository.FindByID(ctx, ll.ID)
	if err != nil {
		return err
	}

	// Update fields
	existingLogistics.ProductID = ll.ProductID
	existingLogistics.Quantity = ll.Quantity
	existingLogistics.RegistrationDate = ll.RegistrationDate
	existingLogistics.DeliveryDate = ll.DeliveryDate
	existingLogistics.WarehousePortID = ll.WarehousePortID
	existingLogistics.ShippingPrice = ll.ShippingPrice
	existingLogistics.VehiclePlate = ll.VehiclePlate
	existingLogistics.GuideNumber = ll.GuideNumber
	existingLogistics.UpdatedAt = time.Now()

	// Calculate discount based on quantity
	discount, err := llu.discountRepository.FindByType(ctx, ll.Type)
	if err != nil {
		return err
	}
	if existingLogistics.Quantity > discount.Quantity {
		existingLogistics.DiscountShippingPrice = discount.Percentage / 100 * existingLogistics.ShippingPrice
		existingLogistics.DiscountPercent = discount.Percentage
	}

	return llu.landLogisticsRepository.Update(ctx, existingLogistics)
}

func (llu *landLogisticsUsecase) Remove(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), llu.contextTimeout)
	defer cancel()

	return llu.landLogisticsRepository.Delete(ctx, id)
}
