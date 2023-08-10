package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jordanlanch/docucenter-test/domain"
	"gorm.io/gorm"
)

type LogisticsUsecase struct {
	LogisticsRepository domain.LogisticsRepository
	discountRepository  domain.DiscountRepository
	contextTimeout      time.Duration
}

func NewLogisticsUsecase(LogisticsRepository domain.LogisticsRepository, discountRepository domain.DiscountRepository, timeout time.Duration) domain.LogisticsUsecase {
	return &LogisticsUsecase{
		LogisticsRepository: LogisticsRepository,
		discountRepository:  discountRepository,
		contextTimeout:      timeout,
	}
}

func (llu *LogisticsUsecase) GetMany(pagination *domain.Pagination) ([]*domain.Logistics, error) {
	ctx, cancel := context.WithTimeout(context.Background(), llu.contextTimeout)
	defer cancel()

	return llu.LogisticsRepository.FindMany(ctx, pagination)
}

func (llu *LogisticsUsecase) GetByID(id uuid.UUID) (*domain.Logistics, error) {
	ctx, cancel := context.WithTimeout(context.Background(), llu.contextTimeout)
	defer cancel()

	return llu.LogisticsRepository.FindByID(ctx, id)
}

func (llu *LogisticsUsecase) Create(ll *domain.Logistics) (*domain.Logistics, error) {
	ctx, cancel := context.WithTimeout(context.Background(), llu.contextTimeout)
	defer cancel()

	ll.ID = uuid.New() // Assign a new UUID to the land logistics entry
	ll.CreatedAt = time.Now()
	ll.UpdatedAt = time.Now()

	// Calculate discount based on quantity
	discount, err := llu.discountRepository.FindByTypeAndQuantity(ctx, ll.Type, ll.Quantity)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err // if it's another error, return it
		}

		// Set default values when there's no discount found
		ll.DiscountShippingPrice = 0
		ll.DiscountPercent = 0
	} else {
		ll.DiscountShippingPrice = discount.Percentage / 100 * ll.ShippingPrice
		ll.DiscountPercent = discount.Percentage
	}

	return llu.LogisticsRepository.Store(ctx, ll)
}

func (llu *LogisticsUsecase) Modify(ll *domain.Logistics, id uuid.UUID) (*domain.Logistics, error) {
	ctx, cancel := context.WithTimeout(context.Background(), llu.contextTimeout)
	defer cancel()

	existingLogistics, err := llu.LogisticsRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
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
	discount, err := llu.discountRepository.FindByTypeAndQuantity(ctx, ll.Type, ll.Quantity)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err // if it's another error, return it
		}

		// Set default values when there's no discount found
		existingLogistics.DiscountShippingPrice = 0
		existingLogistics.DiscountPercent = 0
	} else {
		existingLogistics.DiscountShippingPrice = discount.Percentage / 100 * existingLogistics.ShippingPrice
		existingLogistics.DiscountPercent = discount.Percentage
	}

	return llu.LogisticsRepository.Update(ctx, existingLogistics, id)
}

func (llu *LogisticsUsecase) Remove(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), llu.contextTimeout)
	defer cancel()

	return llu.LogisticsRepository.Delete(ctx, id)
}
