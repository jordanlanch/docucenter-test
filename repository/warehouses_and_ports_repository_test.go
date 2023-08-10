package repository

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jordanlanch/docucenter-test/domain"
	"github.com/stretchr/testify/assert"
	"k8s.io/utils/pointer"
)

func TestStoreWarehousePort(t *testing.T) {
	repo := NewWarehousePortRepository(db, domain.WarehousePortTable)

	warehousePort := &domain.WarehousesAndPorts{
		ID:       uuid.New(),
		Type:     "land",
		Name:     "TestWarehouse",
		Location: "TestLocation",
	}

	_, err := repo.Store(context.Background(), warehousePort)
	assert.NoError(t, err)
}

func TestFindWarehousePortByID(t *testing.T) {
	repo := NewWarehousePortRepository(db, domain.WarehousePortTable)

	warehousePort := &domain.WarehousesAndPorts{
		ID:       uuid.New(),
		Type:     "land",
		Name:     "TestWarehouse",
		Location: "TestLocation",
	}

	newWarehousePort, err := repo.Store(context.Background(), warehousePort)
	assert.NoError(t, err)

	foundWarehousePort, err := repo.FindByID(context.Background(), newWarehousePort.ID)
	assert.NoError(t, err)
	assert.Equal(t, newWarehousePort.ID, foundWarehousePort.ID)
	_, err = repo.FindByID(context.Background(), uuid.New())
	assert.Error(t, err)
}

func TestUpdateWarehousePort(t *testing.T) {
	repo := NewWarehousePortRepository(db, domain.WarehousePortTable)

	warehousePort := &domain.WarehousesAndPorts{
		ID:       uuid.New(),
		Type:     "land",
		Name:     "TestWarehouse",
		Location: "TestLocation",
	}

	newWarehouse, err := repo.Store(context.Background(), warehousePort)
	assert.NoError(t, err)

	warehousePort.Name = "UpdatedWarehouse"
	_, err = repo.Update(context.Background(), warehousePort, newWarehouse.ID)
	assert.NoError(t, err)

	updatedWarehousePort, _ := repo.FindByID(context.Background(), warehousePort.ID)
	assert.Equal(t, "UpdatedWarehouse", updatedWarehousePort.Name)
}

func TestFindManyMeters(t *testing.T) {
	repo := NewWarehousePortRepository(db, domain.WarehousePortTable)

	warehousesPorts := make([]*domain.WarehousesAndPorts, 10)
	for i := range warehousesPorts {
		warehousesPorts[i] = &domain.WarehousesAndPorts{
			ID:       uuid.New(),
			Type:     "land",
			Name:     "TestWarehouse",
			Location: "TestLocation",
		}
		db.Create(warehousesPorts[i])
	}

	pagination := &domain.Pagination{Limit: pointer.Int(5), Offset: pointer.Int(0)}
	foundWarehousesPorts, err := repo.FindMany(context.Background(), pagination)
	assert.NoError(t, err)
	assert.Len(t, foundWarehousesPorts, 5)
}

func TestDeleteWarehousePort(t *testing.T) {
	repo := NewWarehousePortRepository(db, domain.WarehousePortTable)

	warehousePort := &domain.WarehousesAndPorts{
		ID:       uuid.New(),
		Type:     "land",
		Name:     "TestWarehouse",
		Location: "TestLocation",
	}

	newWarehouse, err := repo.Store(context.Background(), warehousePort)
	assert.NoError(t, err)

	err = repo.Delete(context.Background(), newWarehouse.ID)
	assert.NoError(t, err)

	_, err = repo.FindByID(context.Background(), newWarehouse.ID)
	assert.Error(t, err)
}

func TestDeleteWarehousePortError(t *testing.T) {
	repo := NewWarehousePortRepository(db, domain.WarehousePortTable)

	// Usamos un ID que no existe para forzar un error
	err := repo.Delete(context.Background(), uuid.New())
	assert.Error(t, err)
}
