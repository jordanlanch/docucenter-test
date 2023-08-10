package repository

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jordanlanch/docucenter-test/domain"
	"github.com/stretchr/testify/assert"
)

func createSampleProduct() *domain.Product {
	return &domain.Product{
		ID:   uuid.New(),
		Name: "Sample Product",
		Type: "Terrestre",
	}
}

func createSampleCustomer() *domain.Customer {
	return &domain.Customer{
		ID:    uuid.New(),
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}
}

func createSampleWarehousePort() *domain.WarehousesAndPorts {
	return &domain.WarehousesAndPorts{
		ID:       uuid.New(),
		Type:     "land",
		Name:     "Sample Warehouse",
		Location: "Sample Location",
	}
}

func createSampleLogistics(productID, warehousePortID, customerID uuid.UUID) *domain.Logistics {
	return &domain.Logistics{
		ID:                    uuid.New(),
		CustomerID:            customerID.String(),
		ProductID:             productID.String(),
		WarehousePortID:       warehousePortID.String(),
		Type:                  "land",
		Quantity:              10,
		RegistrationDate:      time.Now(),
		DeliveryDate:          time.Now().Add(24 * time.Hour),
		ShippingPrice:         100.0,
		DiscountShippingPrice: 90.0,
		VehiclePlate:          "ABC123",
		GuideNumber:           "12345",
		DiscountPercent:       10.0,
	}
}

func TestLogisticsRepository(t *testing.T) {
	productRepo := NewProductRepository(db, domain.ProductTable)
	customerRepo := NewCustomerRepository(db, domain.CustomerTable)
	warehousePortRepo := NewWarehousePortRepository(db, domain.WarehousePortTable)
	logisticsRepo := NewLogisticsRepository(db, domain.LogisticsTable)

	// Create sample product and warehouse_port for reference
	product := createSampleProduct()
	_, err := productRepo.Store(context.Background(), product)
	assert.NoError(t, err)

	customer := createSampleCustomer()
	_, err = customerRepo.Store(context.Background(), customer)
	assert.NoError(t, err)

	warehousePort := createSampleWarehousePort()
	_, err = warehousePortRepo.Store(context.Background(), warehousePort)
	assert.NoError(t, err)

	t.Run("FindMany - success", func(t *testing.T) {
		logistics1 := createSampleLogistics(product.ID, warehousePort.ID, customer.ID)
		_, _ = logisticsRepo.Store(context.Background(), logistics1)

		logistics2 := createSampleLogistics(product.ID, warehousePort.ID, customer.ID)
		_, _ = logisticsRepo.Store(context.Background(), logistics2)

		limit := 10
		offset := 0
		pagination := &domain.Pagination{Limit: &limit, Offset: &offset}
		filter := map[string]interface{}{}
		logisticsList, err := logisticsRepo.FindMany(context.Background(), pagination, filter)
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(logisticsList), 2)
	})

	t.Run("FindMany - success with filter", func(t *testing.T) {
		logistics1 := createSampleLogistics(product.ID, warehousePort.ID, customer.ID)
		_, _ = logisticsRepo.Store(context.Background(), logistics1)

		logistics2 := createSampleLogistics(product.ID, warehousePort.ID, customer.ID)
		_, _ = logisticsRepo.Store(context.Background(), logistics2)

		limit := 10
		offset := 0
		pagination := &domain.Pagination{Limit: &limit, Offset: &offset}
		filter := map[string]interface{}{"product_id": product.ID}
		logisticsList, err := logisticsRepo.FindMany(context.Background(), pagination, filter)
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(logisticsList), 1)
	})

	t.Run("FindMany - empty", func(t *testing.T) {
		limit := 10
		offset := 100
		pagination := &domain.Pagination{Limit: &limit, Offset: &offset}
		filter := map[string]interface{}{}
		logisticsList, err := logisticsRepo.FindMany(context.Background(), pagination, filter)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(logisticsList))
	})

	// Casos de prueba para FindByID
	t.Run("FindByID - success", func(t *testing.T) {
		logistics := createSampleLogistics(product.ID, warehousePort.ID, customer.ID)
		_, _ = logisticsRepo.Store(context.Background(), logistics)

		found, err := logisticsRepo.FindByID(context.Background(), logistics.ID)
		assert.NoError(t, err)
		assert.Equal(t, logistics.ID, found.ID)
	})

	t.Run("FindByID - not found", func(t *testing.T) {
		_, err := logisticsRepo.FindByID(context.Background(), uuid.New())
		assert.Error(t, err)
	})

	// Casos de prueba para Store
	t.Run("Store - success", func(t *testing.T) {
		logistics := createSampleLogistics(product.ID, warehousePort.ID, customer.ID)
		_, err := logisticsRepo.Store(context.Background(), logistics)
		assert.NoError(t, err)
	})

	// Casos de prueba para Update
	t.Run("Update - success", func(t *testing.T) {
		logistics := createSampleLogistics(product.ID, warehousePort.ID, customer.ID)
		newLogistic, _ := logisticsRepo.Store(context.Background(), logistics)

		logistics.Quantity = 50
		_, err := logisticsRepo.Update(context.Background(), logistics, newLogistic.ID)
		assert.NoError(t, err)

		updated, _ := logisticsRepo.FindByID(context.Background(), logistics.ID)
		assert.Equal(t, 50, updated.Quantity)
	})

	t.Run("Update - not found", func(t *testing.T) {
		logistics := &domain.Logistics{
			ID: uuid.New(),
		}
		_, err := logisticsRepo.Update(context.Background(), logistics, logistics.ID)
		assert.Error(t, err)
	})

	// Casos de prueba para Delete
	t.Run("Delete - success", func(t *testing.T) {
		logistics := createSampleLogistics(product.ID, warehousePort.ID, customer.ID)
		_, _ = logisticsRepo.Store(context.Background(), logistics)

		err := logisticsRepo.Delete(context.Background(), logistics.ID)
		assert.NoError(t, err)
	})

	t.Run("Delete - not found", func(t *testing.T) {
		err := logisticsRepo.Delete(context.Background(), uuid.New())
		assert.Error(t, err)
	})
}
