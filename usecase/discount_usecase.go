package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jordanlanch/docucenter-test/domain"
)

type discountUsecase struct {
	discountRepository domain.DiscountRepository
	contextTimeout     time.Duration
}

func NewDiscountUsecase(discountRepository domain.DiscountRepository, timeout time.Duration) domain.DiscountUsecase {
	return &discountUsecase{
		discountRepository: discountRepository,
		contextTimeout:     timeout,
	}
}

func (du *discountUsecase) GetMany(pagination *domain.Pagination, filters map[string]interface{}) ([]*domain.Discount, error) {
	ctx, cancel := context.WithTimeout(context.Background(), du.contextTimeout)
	defer cancel()

	return du.discountRepository.FindMany(ctx, pagination, filters)
}

func (du *discountUsecase) GetByTypeAndQuantity(dtype domain.LogisticsType, quantity int) (*domain.Discount, error) {
	ctx, cancel := context.WithTimeout(context.Background(), du.contextTimeout)
	defer cancel()

	return du.discountRepository.FindByTypeAndQuantity(ctx, dtype, quantity)
}

func (du *discountUsecase) Create(d *domain.Discount) (*domain.Discount, error) {
	ctx, cancel := context.WithTimeout(context.Background(), du.contextTimeout)
	defer cancel()

	d.ID = uuid.New() // Assign a new UUID to the discount
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()

	return du.discountRepository.Store(ctx, d)
}

func (du *discountUsecase) Modify(d *domain.Discount, id uuid.UUID) (*domain.Discount, error) {
	ctx, cancel := context.WithTimeout(context.Background(), du.contextTimeout)
	defer cancel()

	existingDiscount, err := du.discountRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update fields
	existingDiscount.Type = d.Type
	existingDiscount.QuantityFrom = d.QuantityFrom
	existingDiscount.QuantityTo = d.QuantityTo
	existingDiscount.Percentage = d.Percentage
	existingDiscount.UpdatedAt = time.Now()

	return du.discountRepository.Update(ctx, existingDiscount, id)
}

func (du *discountUsecase) Remove(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), du.contextTimeout)
	defer cancel()

	return du.discountRepository.Delete(ctx, id)
}
