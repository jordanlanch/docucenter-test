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

func TestProductGetByID(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	productID := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockProduct := &domain.Product{
			ID: productID,
		}

		mockProductRepo.On("FindByID", mock.Anything, productID).Return(mockProduct, nil).Once()

		u := NewProductUsecase(mockProductRepo, time.Second*2)

		product, err := u.GetByID(productID)

		assert.NoError(t, err)
		assert.NotNil(t, product)
		assert.Equal(t, productID, product.ID)

		mockProductRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockProductRepo.On("FindByID", mock.Anything, productID).Return(nil, errors.New("Not found")).Once()

		u := NewProductUsecase(mockProductRepo, time.Second*2)

		product, err := u.GetByID(productID)

		assert.Error(t, err)
		assert.Nil(t, product)

		mockProductRepo.AssertExpectations(t)
	})
}

func TestProductCreate(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)

	mockProduct := &domain.Product{
		Name: "TestProduct",
		Type: "TestType",
	}

	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("Store", mock.Anything, mockProduct).Return(mockProduct, nil).Once()

		u := NewProductUsecase(mockProductRepo, time.Second*2)

		_, err := u.Create(mockProduct)

		assert.NoError(t, err)
		assert.NotNil(t, mockProduct.ID)

		mockProductRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockProductRepo.On("Store", mock.Anything, mockProduct).Return(nil, errors.New("Error storing")).Once()

		u := NewProductUsecase(mockProductRepo, time.Second*2)

		_, err := u.Create(mockProduct)

		assert.Error(t, err)

		mockProductRepo.AssertExpectations(t)
	})
}

func TestProductLandModify(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	productID := uuid.New()

	mockProduct := &domain.Product{
		ID:   productID,
		Name: "TestProductLand",
		Type: "land",
	}

	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("FindByID", mock.Anything, productID).Return(mockProduct, nil).Once()
		mockProductRepo.On("Update", mock.Anything, mockProduct).Return(mockProduct, nil).Once()

		u := NewProductUsecase(mockProductRepo, time.Second*2)

		_, err := u.Modify(mockProduct)

		assert.NoError(t, err)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockProductRepo.On("FindByID", mock.Anything, productID).Return(nil, errors.New("Not found")).Once()

		u := NewProductUsecase(mockProductRepo, time.Second*2)

		_, err := u.Modify(mockProduct)

		assert.Error(t, err)
		mockProductRepo.AssertExpectations(t)
	})
}

func TestProductMaritimeModify(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	productID := uuid.New()

	mockProduct := &domain.Product{
		ID:   productID,
		Name: "TestProductMaritime",
		Type: "maritime",
	}

	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("FindByID", mock.Anything, productID).Return(mockProduct, nil).Once()
		mockProductRepo.On("Update", mock.Anything, mockProduct).Return(mockProduct, nil).Once()

		u := NewProductUsecase(mockProductRepo, time.Second*2)

		_, err := u.Modify(mockProduct)

		assert.NoError(t, err)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockProductRepo.On("FindByID", mock.Anything, productID).Return(nil, errors.New("Not found")).Once()

		u := NewProductUsecase(mockProductRepo, time.Second*2)

		_, err := u.Modify(mockProduct)

		assert.Error(t, err)
		mockProductRepo.AssertExpectations(t)
	})
}

func TestProductRemove(t *testing.T) {
	mockProductRepo := new(mocks.ProductRepository)
	productID := uuid.New()

	t.Run("success", func(t *testing.T) {
		mockProductRepo.On("Delete", mock.Anything, productID).Return(nil).Once()

		u := NewProductUsecase(mockProductRepo, time.Second*2)

		err := u.Remove(productID)

		assert.NoError(t, err)
		mockProductRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockProductRepo.On("Delete", mock.Anything, productID).Return(errors.New("Error deleting")).Once()

		u := NewProductUsecase(mockProductRepo, time.Second*2)

		err := u.Remove(productID)

		assert.Error(t, err)
		mockProductRepo.AssertExpectations(t)
	})
}
