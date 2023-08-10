package repository

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jordanlanch/docucenter-test/domain"
	"github.com/stretchr/testify/assert"
)

func TestCustomerRepositoryStore(t *testing.T) {
	repo := NewCustomerRepository(db, domain.CustomerTable)

	// Test storing a valid customer
	customer := &domain.Customer{
		ID:    uuid.New(),
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}

	_, err := repo.Store(context.Background(), customer)
	assert.NoError(t, err)

	// Test storing a customer with an existing ID
	existingID := customer.ID
	customer2 := &domain.Customer{
		ID:    existingID,
		Name:  "Jane Doe",
		Email: "janedoe@example.com",
	}

	_, err = repo.Store(context.Background(), customer2)
	assert.Error(t, err)
}

func TestCustomerRepositoryFindByID(t *testing.T) {
	repo := NewCustomerRepository(db, domain.CustomerTable)

	// Test finding a customer by valid ID
	customer := &domain.Customer{
		ID:    uuid.New(),
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}

	newCustomer, err := repo.Store(context.Background(), customer)
	assert.NoError(t, err)

	foundCustomer, err := repo.FindByID(context.Background(), newCustomer.ID)
	assert.NoError(t, err)
	assert.Equal(t, newCustomer.ID, foundCustomer.ID)

	// Test finding a customer by non-existent ID
	_, err = repo.FindByID(context.Background(), uuid.New())
	assert.Error(t, err)
}

func TestCustomerRepositoryUpdate(t *testing.T) {
	repo := NewCustomerRepository(db, domain.CustomerTable)

	// Test updating a valid customer
	customer := &domain.Customer{
		ID:    uuid.New(),
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}
	newCustomer, err := repo.Store(context.Background(), customer)
	assert.NoError(t, err)

	customer.Email = "updated@example.com"
	_, err = repo.Update(context.Background(), customer, newCustomer.ID)
	assert.NoError(t, err)

	updatedCustomer, _ := repo.FindByID(context.Background(), newCustomer.ID)
	assert.Equal(t, "updated@example.com", updatedCustomer.Email)
}

func TestCustomerRepositoryDelete(t *testing.T) {
	repo := NewCustomerRepository(db, domain.CustomerTable)

	// Test deleting a valid customer
	customer := &domain.Customer{
		ID:    uuid.New(),
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}

	newCustomer, err := repo.Store(context.Background(), customer)
	assert.NoError(t, err)

	err = repo.Delete(context.Background(), newCustomer.ID)
	assert.NoError(t, err)

	_, err = repo.FindByID(context.Background(), newCustomer.ID)
	assert.Error(t, err)

	// Test deleting a non-existent customer
	err = repo.Delete(context.Background(), uuid.New())
	assert.Error(t, err)
}
