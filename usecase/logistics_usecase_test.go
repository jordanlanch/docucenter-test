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

func TestLogisticsUsecase_GetMany(t *testing.T) {
	mockLogisticsRepo := new(mocks.LogisticsRepository)
	mockDiscountRepo := new(mocks.DiscountRepository)

	limit := 10
	offset := 0
	pagination := &domain.Pagination{Limit: &limit, Offset: &offset}

	mockLogistics := []*domain.Logistics{
		{ID: uuid.New()},
		{ID: uuid.New()},
	}

	t.Run("success", func(t *testing.T) {
		mockLogisticsRepo.On("FindMany", mock.AnythingOfType("*context.timerCtx"), pagination).Return(mockLogistics, nil).Once()

		u := NewLogisticsUsecase(mockLogisticsRepo, mockDiscountRepo, time.Second*2)
		result, err := u.GetMany(pagination)

		assert.NoError(t, err)
		assert.Len(t, result, len(mockLogistics))

		mockLogisticsRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockLogisticsRepo.On("FindMany", mock.AnythingOfType("*context.timerCtx"), pagination).Return(nil, errors.New("Unexpected error")).Once()

		u := NewLogisticsUsecase(mockLogisticsRepo, mockDiscountRepo, time.Second*2)
		result, err := u.GetMany(pagination)

		assert.Error(t, err)
		assert.Nil(t, result)

		mockLogisticsRepo.AssertExpectations(t)
	})
}

func TestLogisticsUsecase_GetByID(t *testing.T) {
	mockLogisticsRepo := new(mocks.LogisticsRepository)

	logisticID := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockLogistic := &domain.Logistics{
			ID: logisticID,
		}

		mockLogisticsRepo.On("FindByID", mock.AnythingOfType("*context.timerCtx"), logisticID).Return(mockLogistic, nil).Once()

		u := NewLogisticsUsecase(mockLogisticsRepo, nil, time.Second*2)

		logistic, err := u.GetByID(logisticID)

		assert.NoError(t, err)
		assert.NotNil(t, logistic)
		assert.Equal(t, logisticID, logistic.ID)

		mockLogisticsRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockLogisticsRepo.On("FindByID", mock.AnythingOfType("*context.timerCtx"), logisticID).Return(nil, errors.New("Not found")).Once()

		u := NewLogisticsUsecase(mockLogisticsRepo, nil, time.Second*2)

		logistic, err := u.GetByID(logisticID)

		assert.Error(t, err)
		assert.Nil(t, logistic)

		mockLogisticsRepo.AssertExpectations(t)
	})
}

func TestLogisticsUsecase_Create(t *testing.T) {
	mockLogisticsRepo := new(mocks.LogisticsRepository)
	mockDiscountRepo := new(mocks.DiscountRepository)

	mockLogistic := &domain.Logistics{
		ID:            uuid.New(),
		Type:          "land",
		ShippingPrice: 1000,
		Quantity:      18,
	}
	mockDiscount := &domain.Discount{
		Type:         "maritime",
		QuantityFrom: 17,
		QuantityTo:   25,
		Percentage:   12,
	}

	t.Run("success", func(t *testing.T) {
		mockDiscountRepo.On("FindByTypeAndQuantity", mock.Anything, mockLogistic.Type, mockLogistic.Quantity).Return(mockDiscount, nil).Once()
		mockLogisticsRepo.On("Store", mock.Anything, mockLogistic).Return(mockLogistic, nil).Once()

		u := NewLogisticsUsecase(mockLogisticsRepo, mockDiscountRepo, time.Second*2)

		result, err := u.Create(mockLogistic)

		assert.NoError(t, err)
		assert.NotNil(t, result)

		mockLogisticsRepo.AssertExpectations(t)
		mockDiscountRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockDiscountRepo.On("FindByTypeAndQuantity", mock.Anything, mockLogistic.Type, mockLogistic.Quantity).Return(nil, errors.New("Unexpected error")).Once()

		u := NewLogisticsUsecase(mockLogisticsRepo, mockDiscountRepo, time.Second*2)

		result, err := u.Create(mockLogistic)

		assert.Error(t, err)
		assert.Nil(t, result)

		mockDiscountRepo.AssertExpectations(t)
	})
}

func TestLogisticsUsecase_Modify(t *testing.T) {
	mockLogisticsRepo := new(mocks.LogisticsRepository)
	mockDiscountRepo := new(mocks.DiscountRepository)

	mockLogistic := &domain.Logistics{
		ID:            uuid.New(),
		Type:          "land",
		ShippingPrice: 1000,
		Quantity:      18,
	}
	mockDiscount := &domain.Discount{
		Type:         "maritime",
		QuantityFrom: 17,
		QuantityTo:   25,
		Percentage:   12,
	}

	t.Run("success", func(t *testing.T) {
		mockLogisticsRepo.On("FindByID", mock.Anything, mockLogistic.ID).Return(mockLogistic, nil).Once()
		mockDiscountRepo.On("FindByTypeAndQuantity", mock.Anything, mockLogistic.Type, mockLogistic.Quantity).Return(mockDiscount, nil).Once()
		mockLogisticsRepo.On("Update", mock.Anything, mockLogistic, mockLogistic.ID).Return(mockLogistic, nil).Once()

		u := NewLogisticsUsecase(mockLogisticsRepo, mockDiscountRepo, time.Second*2)

		result, err := u.Modify(mockLogistic, mockLogistic.ID)

		assert.NoError(t, err)
		assert.NotNil(t, result)

		mockLogisticsRepo.AssertExpectations(t)
		mockDiscountRepo.AssertExpectations(t)
	})

	t.Run("error-not-found", func(t *testing.T) {
		mockLogisticsRepo.On("FindByID", mock.Anything, mockLogistic.ID).Return(nil, errors.New("Not found")).Once()

		u := NewLogisticsUsecase(mockLogisticsRepo, mockDiscountRepo, time.Second*2)

		result, err := u.Modify(mockLogistic, mockLogistic.ID)

		assert.Error(t, err)
		assert.Nil(t, result)

		mockLogisticsRepo.AssertExpectations(t)
	})
}

func TestLogisticsUsecase_Remove(t *testing.T) {
	mockLogisticsRepo := new(mocks.LogisticsRepository)

	logisticID := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockLogisticsRepo.On("Delete", mock.Anything, logisticID).Return(nil).Once()

		u := NewLogisticsUsecase(mockLogisticsRepo, nil, time.Second*2)

		err := u.Remove(logisticID)

		assert.NoError(t, err)

		mockLogisticsRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockLogisticsRepo.On("Delete", mock.Anything, logisticID).Return(errors.New("Unexpected error")).Once()

		u := NewLogisticsUsecase(mockLogisticsRepo, nil, time.Second*2)

		err := u.Remove(logisticID)

		assert.Error(t, err)

		mockLogisticsRepo.AssertExpectations(t)
	})
}
