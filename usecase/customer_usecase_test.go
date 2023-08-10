package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jordanlanch/docucenter-test/domain"
	"github.com/jordanlanch/docucenter-test/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCustomerUseCase_GetMany(t *testing.T) {
	mockCustomerRepo := new(mocks.CustomerRepository)
	limit := 10
	offset := 0
	pagination := &domain.Pagination{Limit: &limit, Offset: &offset}

	t.Run("success", func(t *testing.T) {
		mockCustomers := []*domain.Customer{
			{
				ID:    uuid.New(),
				Name:  "John Doe",
				Email: "johndoe@example.com",
			},
			{
				ID:    uuid.New(),
				Name:  "Jane Smith",
				Email: "janesmith@example.com",
			},
		}

		mockCustomerRepo.On("FindMany", mock.Anything, pagination).Return(mockCustomers, nil).Once()

		u := NewCustomerUsecase(mockCustomerRepo, time.Second*2)

		customers, err := u.GetMany(pagination)

		assert.NoError(t, err)
		assert.NotNil(t, customers)
		assert.Equal(t, 2, len(customers))

		mockCustomerRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockCustomerRepo.On("FindMany", mock.Anything, pagination).Return(nil, errors.New("Error fetching")).Once()

		u := NewCustomerUsecase(mockCustomerRepo, time.Second*2)

		customers, err := u.GetMany(pagination)

		assert.Error(t, err)
		assert.Nil(t, customers)

		mockCustomerRepo.AssertExpectations(t)
	})
}

func TestCustomerUseCase_GetByID(t *testing.T) {
	mockCustomerRepo := new(mocks.CustomerRepository)

	id := uuid.New()
	t.Run("success", func(t *testing.T) {
		mockCustomer := &domain.Customer{
			ID:    id,
			Name:  "John Doe",
			Email: "johndoe@example.com",
		}

		mockCustomerRepo.On("FindByID", mock.Anything, id).Return(mockCustomer, nil).Once()

		u := NewCustomerUsecase(mockCustomerRepo, time.Second*2)

		customer, err := u.GetByID(id)

		assert.NoError(t, err)
		assert.NotNil(t, customer)
		assert.Equal(t, id, customer.ID)

		mockCustomerRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockCustomerRepo.On("FindByID", mock.Anything, id).Return(nil, errors.New("Not found")).Once()

		u := NewCustomerUsecase(mockCustomerRepo, time.Second*2)

		customer, err := u.GetByID(id)

		assert.Error(t, err)
		assert.Nil(t, customer)

		mockCustomerRepo.AssertExpectations(t)
	})
}

func TestCustomerUseCase_Create(t *testing.T) {
	mockCustomerRepo := new(mocks.CustomerRepository)

	mockCustomer := &domain.Customer{
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}

	t.Run("success", func(t *testing.T) {
		mockCustomerRepo.On("Store", mock.Anything, mockCustomer).Return(mockCustomer, nil).Once()

		u := NewCustomerUsecase(mockCustomerRepo, time.Second*2)

		_, err := u.Create(mockCustomer)

		assert.NoError(t, err)
		assert.NotNil(t, mockCustomer.ID)

		mockCustomerRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockCustomerRepo.On("Store", mock.Anything, mockCustomer).Return(nil, errors.New("Error storing")).Once()

		u := NewCustomerUsecase(mockCustomerRepo, time.Second*2)

		_, err := u.Create(mockCustomer)

		assert.Error(t, err)

		mockCustomerRepo.AssertExpectations(t)
	})
}

func TestCustomerUseCase_Modify(t *testing.T) {
	mockCustomerRepo := new(mocks.CustomerRepository)

	customerID := uuid.New()

	mockCustomer := &domain.Customer{
		ID:    customerID,
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}

	t.Run("success", func(t *testing.T) {
		mockCustomerRepo.On("FindByID", mock.Anything, customerID).Return(mockCustomer, nil).Once()
		mockCustomerRepo.On("Update", mock.Anything, mockCustomer, customerID).Return(mockCustomer, nil).Once()

		u := NewCustomerUsecase(mockCustomerRepo, time.Second*2)

		_, err := u.Modify(mockCustomer, mockCustomer.ID)

		assert.NoError(t, err)
		mockCustomerRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockCustomerRepo.On("FindByID", mock.Anything, customerID).Return(nil, errors.New("Not found")).Once()

		u := NewCustomerUsecase(mockCustomerRepo, time.Second*2)

		_, err := u.Modify(mockCustomer, mockCustomer.ID)

		assert.Error(t, err)
		mockCustomerRepo.AssertExpectations(t)
	})
}

func TestCustomerUseCase_Remove(t *testing.T) {
	mockCustomerRepo := new(mocks.CustomerRepository)

	customerID := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockCustomerRepo.On("Delete", mock.Anything, customerID).Return(nil).Once()

		u := NewCustomerUsecase(mockCustomerRepo, time.Second*2)

		err := u.Remove(customerID)

		assert.NoError(t, err)
		mockCustomerRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockCustomerRepo.On("Delete", mock.Anything, customerID).Return(errors.New("Error deleting")).Once()

		u := NewCustomerUsecase(mockCustomerRepo, time.Second*2)

		err := u.Remove(customerID)

		assert.Error(t, err)
		mockCustomerRepo.AssertExpectations(t)
	})
}
