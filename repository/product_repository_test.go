package repository

import (
	"context"
	"testing"
	"fmt"

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

	id := uuid.New()
	product := &domain.Product{
		ID: id,
		// Add other fields from your Product struct as needed.
	}
	repo.Store(context.Background(), product)

	foundProduct, err := repo.FindByID(context.Background(), id)
	assert.NoError(t, err)
	assert.Equal(t, id, foundProduct.ID)

	_, err = repo.FindByID(context.Background(), uuid.New())
	assert.Error(t, err)
}

func TestUpdateProduct(t *testing.T) {
	repo := NewProductRepository(db, "products")

	product := &domain.Product{
		ID:   uuid.New(),
		Name: "Updated Name",
		Type: "maritime",
	}
	repo.Store(context.Background(), product)

	// Update fields for product and save updated values
	// For example: product.Name = "Updated Name"
	_, err := repo.Update(context.Background(), product)
	assert.NoError(t, err)

	updatedProduct, _ := repo.FindByID(context.Background(), product.ID)
	//Assert updated fields, for example:
	assert.Equal(t, "Updated Name", updatedProduct.Name)
}


func TestDeleteProduct(t *testing.T) {
	repo := NewProductRepository(db, "products")

	product := &domain.Product{
		ID: uuid.New(),
		// Add other fields from your Product struct as needed.
	}

	repo.Store(context.Background(), product)

	err := repo.Delete(context.Background(), product.ID)
	assert.NoError(t, err)

	_, err = repo.FindByID(context.Background(), product.ID)
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

	products, err := repo.FindMany(context.Background(), pagination)
	assert.NoError(t, err)
	assert.Len(t, products, 2) // Debería devolver solo 2 productos debido a la limitación
}

func TestUpdateProductNotFound(t *testing.T) {
	repo := NewProductRepository(db, "products")

	product := &domain.Product{
		ID:   uuid.New(), // Usar un ID que no existe en la base de datos
		Name: "Non-existent Product",
		Type: "land",
	}

	_, err := repo.Update(context.Background(), product)
	assert.Error(t, err)
}

func TestDeleteProductNotFound(t *testing.T) {
	repo := NewProductRepository(db, "products")

	err := repo.Delete(context.Background(), uuid.New()) // Usar un ID que no existe en la base de datos
	assert.Error(t, err)
}
