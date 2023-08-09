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

// func TestGetMeterById(t *testing.T) {
// 	mockMeterRepository := new(mocks.MeterRepository)
// 	meterID := uuid.New()

// 	t.Run("success", func(t *testing.T) {
// 		mockMeter := &domain.Meter{
// 			ID:      meterID,
// 			Address: "Mock Address",
// 		}

// 		mockMeterRepository.On("FindByID", mock.Anything, meterID).Return(mockMeter, nil).Once()

// 		u := NewMeterUsecase(mockMeterRepository, time.Second*2)

// 		meter, err := u.GetMeterById(meterID)

// 		assert.NoError(t, err)
// 		assert.NotNil(t, meter)
// 		assert.Equal(t, meterID, meter.ID)

// 		mockMeterRepository.AssertExpectations(t)
// 	})

// 	t.Run("error", func(t *testing.T) {
// 		mockMeterRepository.On("FindByID", mock.Anything, meterID).Return(nil, errors.New("Unexpected")).Once()

// 		u := NewMeterUsecase(mockMeterRepository, time.Second*2)

// 		meter, err := u.GetMeterById(meterID)

// 		assert.Error(t, err)
// 		assert.Nil(t, meter)

// 		mockMeterRepository.AssertExpectations(t)
// 	})
// }

// func TestGetManyMeters(t *testing.T) {
// 	mockMeterRepository := new(mocks.MeterRepository)
// 	limit := 10
// 	offset := 0
// 	pagination := &domain.Pagination{Limit: &limit, Offset: &offset}

// 	t.Run("success", func(t *testing.T) {
// 		mockMeters := []*domain.Meter{
// 			{
// 				ID:      uuid.New(),
// 				Address: "Mock Address",
// 			},
// 			{
// 				ID:      uuid.New(),
// 				Address: "Mock Address 2",
// 			},
// 		}

// 		mockMeterRepository.On("FindMany", mock.Anything, pagination).Return(mockMeters, nil).Once()

// 		u := NewMeterUsecase(mockMeterRepository, time.Second*2)

// 		meters, err := u.GetManyMeters(pagination)

// 		assert.NoError(t, err)
// 		assert.NotNil(t, meters)
// 		assert.Len(t, meters, len(mockMeters))

// 		mockMeterRepository.AssertExpectations(t)
// 	})

// 	t.Run("error", func(t *testing.T) {
// 		mockMeterRepository.On("FindMany", mock.Anything, pagination).Return(nil, errors.New("Unexpected")).Once()

// 		u := NewMeterUsecase(mockMeterRepository, time.Second*2)

// 		meters, err := u.GetManyMeters(pagination)

// 		assert.Error(t, err)
// 		assert.Nil(t, meters)

// 		mockMeterRepository.AssertExpectations(t)
// 	})
// }

// func TestNewMeterUsecase(t *testing.T) {
// 	mockMeterRepository := new(mocks.MeterRepository)
// 	timeout := time.Second * 2

// 	u := NewMeterUsecase(mockMeterRepository, timeout)

// 	assert.NotNil(t, u)
// 	assert.Equal(t, mockMeterRepository, u.(*MeterUsecase).meterRepository)
// 	assert.Equal(t, timeout, u.(*MeterUsecase).contextTimeout)
// }
