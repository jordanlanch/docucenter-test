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

func (du *discountUsecase) GetByType(dtype domain.LogisticsType) (*domain.Discount, error) {
	ctx, cancel := context.WithTimeout(context.Background(), du.contextTimeout)
	defer cancel()

	return du.discountRepository.FindByType(ctx, dtype)
}

func (du *discountUsecase) Create(d *domain.Discount) error {
	ctx, cancel := context.WithTimeout(context.Background(), du.contextTimeout)
	defer cancel()

	d.ID = uuid.New() // Assign a new UUID to the discount
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()

	return du.discountRepository.Store(ctx, d)
}

func (du *discountUsecase) Modify(d *domain.Discount) error {
	ctx, cancel := context.WithTimeout(context.Background(), du.contextTimeout)
	defer cancel()

	existingDiscount, err := du.discountRepository.FindByType(ctx, d.Type)
	if err != nil {
		return err
	}

	// Update fields
	existingDiscount.Type = d.Type
	existingDiscount.Quantity = d.Quantity
	existingDiscount.Percentage = d.Percentage
	existingDiscount.UpdatedAt = time.Now()

	return du.discountRepository.Update(ctx, existingDiscount)
}

func (du *discountUsecase) Remove(dtype domain.LogisticsType) error {
	ctx, cancel := context.WithTimeout(context.Background(), du.contextTimeout)
	defer cancel()

	return du.discountRepository.Delete(ctx, dtype)
}
