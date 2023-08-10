package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/jordanlanch/docucenter-test/domain"
	"github.com/stretchr/testify/assert"
)

func TestStoreProduct(t *testing.T) {
	repo := NewProductRepository(db, "products")

	product := &domain.Product{
		ID:   uuid.New(),
		Name: "Product Name",
		Type: "land",
	}

	_, err := repo.Store(context.Background(), product)
	assert.NoError(t, err)
}

func TestFindProductByID(t *testing.T) {
	repo := NewProductRepository(db, "products")

	product := &domain.Product{
		ID:   uuid.New(),
		Name: "Name Product",
		Type: "land",
	}
	newProduct, _ := repo.Store(context.Background(), product)

	foundProduct, err := repo.FindByID(context.Background(), newProduct.ID)
	assert.NoError(t, err)
	assert.Equal(t, newProduct.ID, foundProduct.ID)
	_, err = repo.FindByID(context.Background(), uuid.New())
	assert.Error(t, err)
}

func TestUpdateProduct(t *testing.T) {
	repo := NewProductRepository(db, "products")

	product := &domain.Product{
		ID:   uuid.New(),
		Name: "New Name",
		Type: "maritime",
	}
	newProduct, _ := repo.Store(context.Background(), product)

	// Update fields for product and save updated values
	product.Name = "Updated Name"
	_, err := repo.Update(context.Background(), product, newProduct.ID)
	assert.NoError(t, err)

	updatedProduct, _ := repo.FindByID(context.Background(), product.ID)
	//Assert updated fields, for example:
	assert.Equal(t, "Updated Name", updatedProduct.Name)
}

func TestDeleteProduct(t *testing.T) {
	repo := NewProductRepository(db, "products")

	product := &domain.Product{
		ID:   uuid.New(),
		Name: "Name Product delete",
		Type: "land",
	}

	newProduct, _ := repo.Store(context.Background(), product)

	err := repo.Delete(context.Background(), newProduct.ID)
	assert.NoError(t, err)

	_, err = repo.FindByID(context.Background(), newProduct.ID)
	assert.Error(t, err)
}

func TestFindManyProducts(t *testing.T) {
	repo := NewProductRepository(db, "products")

	// Crear algunos productos para probar la paginación
	for i := 0; i < 5; i++ {
		product := &domain.Product{
			ID:   uuid.New(),
			Name: "Product " + fmt.Sprintf("%d", i),
			Type: "land",
		}
		repo.Store(context.Background(), product)
	}

	// Prueba de paginación
	limit := 2
	offset := 0
	pagination := &domain.Pagination{Limit: &limit, Offset: &offset}
	filter := map[string]interface{}{}
	products, err := repo.FindMany(context.Background(), pagination, filter)
	assert.NoError(t, err)
	assert.Len(t, products, 2) // Debería devolver solo 2 productos debido a la limitación
}

func TestFindManyProductsWithFilter(t *testing.T) {
	repo := NewProductRepository(db, "products")

	// Crear algunos productos para probar la paginación
	for i := 0; i < 5; i++ {
		product := &domain.Product{
			ID:   uuid.New(),
			Name: "Products " + fmt.Sprintf("%d", i),
			Type: "land",
		}
		repo.Store(context.Background(), product)
	}

	// Prueba de paginación
	limit := 2
	offset := 0
	pagination := &domain.Pagination{Limit: &limit, Offset: &offset}
	filter := map[string]interface{}{"name": "Product 1"}
	products, err := repo.FindMany(context.Background(), pagination, filter)
	assert.NoError(t, err)
	assert.Len(t, products, 1) // Debería devolver solo 1 productos debido al filtro
}

func TestUpdateProductNotFound(t *testing.T) {
	repo := NewProductRepository(db, "products")

	product := &domain.Product{
		ID:   uuid.New(),
		Name: "Non-existent Product",
		Type: "land",
	}

	_, err := repo.Update(context.Background(), product, uuid.New())
	assert.Error(t, err)
}

func TestDeleteProductNotFound(t *testing.T) {
	repo := NewProductRepository(db, "products")

	err := repo.Delete(context.Background(), uuid.New()) // Usar un ID que no existe en la base de datos
	assert.Error(t, err)
}
