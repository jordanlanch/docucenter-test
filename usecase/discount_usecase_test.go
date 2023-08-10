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

func TestDiscountUseCase_GetMany(t *testing.T) {
	mockDiscountRepo := new(mocks.DiscountRepository)

	limit := 10
	offset := 0
	pagination := &domain.Pagination{Limit: &limit, Offset: &offset}

	t.Run("success", func(t *testing.T) {
		mockDiscounts := []*domain.Discount{
			{
				ID:           uuid.New(),
				Type:         domain.LogisticsType("land"),
				QuantityFrom: 10,
				QuantityTo:   20,
				Percentage:   10,
			},
		}

		mockDiscountRepo.On("FindMany", mock.Anything, pagination).Return(mockDiscounts, nil).Once()

		u := NewDiscountUsecase(mockDiscountRepo, time.Second*2)

		discounts, err := u.GetMany(pagination)

		assert.NoError(t, err)
		assert.NotNil(t, discounts)
		assert.Equal(t, 1, len(discounts))

		mockDiscountRepo.AssertExpectations(t)
	})
}

func TestDiscountUseCase_GetByType(t *testing.T) {
	mockDiscountRepo := new(mocks.DiscountRepository)
	dtype := domain.LogisticsType("land")

	t.Run("success", func(t *testing.T) {
		mockDiscount := &domain.Discount{
			ID:           uuid.New(),
			Type:         dtype,
			QuantityFrom: 11,
			QuantityTo:   16,
			Percentage:   10,
		}

		mockDiscountRepo.On("FindByTypeAndQuantity", mock.Anything, dtype, 15).Return(mockDiscount, nil).Once()

		u := NewDiscountUsecase(mockDiscountRepo, time.Second*2)

		discount, err := u.GetByTypeAndQuantity(dtype, 15)

		assert.NoError(t, err)
		assert.NotNil(t, discount)
		assert.Equal(t, dtype, discount.Type)

		mockDiscountRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockDiscountRepo.On("FindByTypeAndQuantity", mock.Anything, dtype, 19).Return(nil, errors.New("Not found")).Once()

		u := NewDiscountUsecase(mockDiscountRepo, time.Second*2)

		discount, err := u.GetByTypeAndQuantity(dtype, 19)

		assert.Error(t, err)
		assert.Nil(t, discount)

		mockDiscountRepo.AssertExpectations(t)
	})
}

func TestDiscountUseCase_Create(t *testing.T) {
	mockDiscountRepo := new(mocks.DiscountRepository)

	mockDiscount := &domain.Discount{
		Type:         domain.LogisticsType("maritime"),
		QuantityFrom: 17,
		QuantityTo:   25,
		Percentage:   12,
	}

	t.Run("success", func(t *testing.T) {
		mockDiscountRepo.On("Store", mock.Anything, mockDiscount).Return(mockDiscount, nil).Once()

		u := NewDiscountUsecase(mockDiscountRepo, time.Second*2)

		_, err := u.Create(mockDiscount)

		assert.NoError(t, err)
		assert.NotNil(t, mockDiscount.ID)

		mockDiscountRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockDiscountRepo.On("Store", mock.Anything, mockDiscount).Return(nil, errors.New("Error storing")).Once()

		u := NewDiscountUsecase(mockDiscountRepo, time.Second*2)

		_, err := u.Create(mockDiscount)

		assert.Error(t, err)

		mockDiscountRepo.AssertExpectations(t)
	})
}

func TestDiscountUseCase_Modify(t *testing.T) {
	mockDiscountRepo := new(mocks.DiscountRepository)

	mockDiscount := &domain.Discount{
		ID:           uuid.New(),
		Type:         domain.LogisticsType("land"),
		QuantityFrom: 26,
		QuantityTo:   32,
		Percentage:   19,
	}

	t.Run("success", func(t *testing.T) {
		mockDiscountRepo.On("FindByID", mock.Anything, mockDiscount.ID).Return(mockDiscount, nil).Once()
		mockDiscountRepo.On("Update", mock.Anything, mockDiscount, mockDiscount.ID).Return(mockDiscount, nil).Once()

		u := NewDiscountUsecase(mockDiscountRepo, time.Second*2)

		_, err := u.Modify(mockDiscount, mockDiscount.ID)

		assert.NoError(t, err)
		mockDiscountRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockDiscountRepo.On("FindByID", mock.Anything, mockDiscount.ID).Return(nil, errors.New("Not found")).Once()

		u := NewDiscountUsecase(mockDiscountRepo, time.Second*2)

		_, err := u.Modify(mockDiscount, mockDiscount.ID)

		assert.Error(t, err)
		mockDiscountRepo.AssertExpectations(t)
	})
}

func TestDiscountUseCase_Remove(t *testing.T) {
	mockDiscountRepo := new(mocks.DiscountRepository)

	discountID := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockDiscountRepo.On("Delete", mock.Anything, discountID).Return(nil).Once()

		u := NewDiscountUsecase(mockDiscountRepo, time.Second*2)

		err := u.Remove(discountID)

		assert.NoError(t, err)
		mockDiscountRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockDiscountRepo.On("Delete", mock.Anything, discountID).Return(errors.New("Error deleting")).Once()

		u := NewDiscountUsecase(mockDiscountRepo, time.Second*2)

		err := u.Remove(discountID)

		assert.Error(t, err)
		mockDiscountRepo.AssertExpectations(t)
	})
}
