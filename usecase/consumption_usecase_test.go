package usecase

// import (
// 	"errors"
// 	"testing"
// 	"time"

// 	"github.com/google/uuid"
// 	"github.com/jordanlanch/docucenter-test/domain"
// 	"github.com/jordanlanch/docucenter-test/domain/mocks"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestGetConsumptionByID(t *testing.T) {
// 	mockConsumptionRepo := new(mocks.ConsumptionRepository)
// 	consumptionId := uuid.New()

// 	t.Run("success", func(t *testing.T) {
// 		mockConsumption := &domain.Consumption{
// 			ID: consumptionId,
// 		}

// 		mockConsumptionRepo.On("FindByID", mock.Anything, consumptionId).Return(mockConsumption, nil).Once()

// 		u := NewConsumptionUsecase(mockConsumptionRepo, time.Second*2)

// 		consumption, err := u.GetConsumptionByID(consumptionId)

// 		assert.NoError(t, err)
// 		assert.NotNil(t, consumption)
// 		assert.Equal(t, consumptionId, consumption.ID)

// 		mockConsumptionRepo.AssertExpectations(t)
// 	})

// 	t.Run("error", func(t *testing.T) {
// 		mockConsumptionRepo.On("FindByID", mock.Anything, consumptionId).Return(nil, errors.New("Unexpected")).Once()

// 		u := NewConsumptionUsecase(mockConsumptionRepo, time.Second*2)

// 		consumption, err := u.GetConsumptionByID(consumptionId)

// 		assert.Error(t, err)
// 		assert.Nil(t, consumption)

// 		mockConsumptionRepo.AssertExpectations(t)
// 	})
// }

// func TestGetConsumptionsByPeriod(t *testing.T) {
// 	mockConsumptionRepo := new(mocks.ConsumptionRepository)
// 	periodType := "monthly"
// 	start := "Jan 2021"
// 	end := "Dec 2021"
// 	meterIds := []int{1, 2, 3}
// 	limit := 10
// 	offset := 0
// 	pagination := &domain.Pagination{Limit: &limit, Offset: &offset}

// 	t.Run("success", func(t *testing.T) {
// 		mockResponse := &domain.Response{}

// 		mockConsumptionRepo.On("FindByPeriod", mock.Anything, periodType, start, end, meterIds, pagination).Return(mockResponse, nil).Once()

// 		u := NewConsumptionUsecase(mockConsumptionRepo, time.Second*2)

// 		response, err := u.GetConsumptionsByPeriod(periodType, start, end, meterIds, pagination)

// 		assert.NoError(t, err)
// 		assert.NotNil(t, response)

// 		mockConsumptionRepo.AssertExpectations(t)
// 	})

// 	t.Run("error", func(t *testing.T) {
// 		mockConsumptionRepo.On("FindByPeriod", mock.Anything, periodType, start, end, meterIds, pagination).Return(nil, errors.New("Unexpected")).Once()

// 		u := NewConsumptionUsecase(mockConsumptionRepo, time.Second*2)

// 		response, err := u.GetConsumptionsByPeriod(periodType, start, end, meterIds, pagination)

// 		assert.Error(t, err)
// 		assert.Nil(t, response)

// 		mockConsumptionRepo.AssertExpectations(t)
// 	})
// }
