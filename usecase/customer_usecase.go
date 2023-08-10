package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jordanlanch/docucenter-test/domain"
)

type customerUsecase struct {
	customerRepository domain.CustomerRepository
	contextTimeout     time.Duration
}

func NewCustomerUsecase(customerRepository domain.CustomerRepository, timeout time.Duration) domain.CustomerUsecase {
	return &customerUsecase{
		customerRepository: customerRepository,
		contextTimeout:     timeout,
	}
}

func (cu *customerUsecase) GetMany(pagination *domain.Pagination) ([]*domain.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cu.contextTimeout)
	defer cancel()

	return cu.customerRepository.FindMany(ctx, pagination)
}

func (cu *customerUsecase) GetByID(id uuid.UUID) (*domain.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cu.contextTimeout)
	defer cancel()

	return cu.customerRepository.FindByID(ctx, id)
}

func (cu *customerUsecase) Create(c *domain.Customer) (*domain.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cu.contextTimeout)
	defer cancel()

	c.ID = uuid.New() // Assign a new UUID to the customer
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()

	return cu.customerRepository.Store(ctx, c)
}

func (cu *customerUsecase) Modify(c *domain.Customer, id uuid.UUID) (*domain.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cu.contextTimeout)
	defer cancel()

	existingCustomer, err := cu.customerRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update fields
	existingCustomer.Name = c.Name
	existingCustomer.Email = c.Email
	existingCustomer.UpdatedAt = time.Now()

	return cu.customerRepository.Update(ctx, existingCustomer, id)
}

func (cu *customerUsecase) Remove(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), cu.contextTimeout)
	defer cancel()

	return cu.customerRepository.Delete(ctx, id)
}
