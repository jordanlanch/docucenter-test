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

func TestWarehousePortGetByID(t *testing.T) {
	mockWarehousePortRepo := new(mocks.WarehousePortRepository)
	warehousePortID := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockWarehousePort := &domain.WarehousesAndPorts{
			ID: warehousePortID,
		}

		mockWarehousePortRepo.On("FindByID", mock.Anything, warehousePortID).Return(mockWarehousePort, nil).Once()

		u := NewWarehousePortUsecase(mockWarehousePortRepo, time.Second*2)

		wp, err := u.GetByID(warehousePortID)

		assert.NoError(t, err)
		assert.NotNil(t, wp)
		assert.Equal(t, warehousePortID, wp.ID)

		mockWarehousePortRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockWarehousePortRepo.On("FindByID", mock.Anything, warehousePortID).Return(nil, errors.New("Not found")).Once()

		u := NewWarehousePortUsecase(mockWarehousePortRepo, time.Second*2)

		wp, err := u.GetByID(warehousePortID)

		assert.Error(t, err)
		assert.Nil(t, wp)

		mockWarehousePortRepo.AssertExpectations(t)
	})
}

func TestWarehousePortCreate(t *testing.T) {
	mockWarehousePortRepo := new(mocks.WarehousePortRepository)

	mockWarehousePort := &domain.WarehousesAndPorts{
		Name:     "TestWarehousePort",
		Type:     "land",
		Location: "TestLocation",
	}

	t.Run("success", func(t *testing.T) {
		mockWarehousePortRepo.On("Store", mock.Anything, mockWarehousePort).Return(mockWarehousePort, nil).Once()

		u := NewWarehousePortUsecase(mockWarehousePortRepo, time.Second*2)

		_, err := u.Create(mockWarehousePort)

		assert.NoError(t, err)
		assert.NotNil(t, mockWarehousePort.ID)

		mockWarehousePortRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockWarehousePortRepo.On("Store", mock.Anything, mockWarehousePort).Return(nil, errors.New("Error storing")).Once()

		u := NewWarehousePortUsecase(mockWarehousePortRepo, time.Second*2)

		_, err := u.Create(mockWarehousePort)

		assert.Error(t, err)

		mockWarehousePortRepo.AssertExpectations(t)
	})
}

func TestWarehousePortModify(t *testing.T) {
	mockWarehousePortRepo := new(mocks.WarehousePortRepository)
	warehousePortID := uuid.New()

	mockWarehousePort := &domain.WarehousesAndPorts{
		ID:       warehousePortID,
		Name:     "TestWarehousePort",
		Type:     "maritime",
		Location: "TestLocation",
	}

	t.Run("success", func(t *testing.T) {
		mockWarehousePortRepo.On("FindByID", mock.Anything, warehousePortID).Return(mockWarehousePort, nil).Once()
		mockWarehousePortRepo.On("Update", mock.Anything, mockWarehousePort).Return(mockWarehousePort, nil).Once()

		u := NewWarehousePortUsecase(mockWarehousePortRepo, time.Second*2)

		_, err := u.Modify(mockWarehousePort)

		assert.NoError(t, err)
		mockWarehousePortRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockWarehousePortRepo.On("FindByID", mock.Anything, warehousePortID).Return(nil, errors.New("Not found")).Once()

		u := NewWarehousePortUsecase(mockWarehousePortRepo, time.Second*2)

		_, err := u.Modify(mockWarehousePort)

		assert.Error(t, err)
		mockWarehousePortRepo.AssertExpectations(t)
	})
}

func TestWarehousePortRemove(t *testing.T) {
	mockWarehousePortRepo := new(mocks.WarehousePortRepository)
	warehousePortID := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockWarehousePortRepo.On("Delete", mock.Anything, warehousePortID).Return(nil).Once()

		u := NewWarehousePortUsecase(mockWarehousePortRepo, time.Second*2)

		err := u.Remove(warehousePortID)

		assert.NoError(t, err)
		mockWarehousePortRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockWarehousePortRepo.On("Delete", mock.Anything, warehousePortID).Return(errors.New("Error deleting")).Once()

		u := NewWarehousePortUsecase(mockWarehousePortRepo, time.Second*2)

		err := u.Remove(warehousePortID)

		assert.Error(t, err)
		mockWarehousePortRepo.AssertExpectations(t)
	})
}
