package repository

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jordanlanch/docucenter-test/domain"
	"github.com/stretchr/testify/assert"
)

func TestDiscountRepositoryStore(t *testing.T) {
	repo := NewDiscountRepository(db, domain.DiscountTable)

	discount := &domain.Discount{
		ID:           uuid.New(),
		Type:         domain.Land,
		QuantityFrom: 5,
		QuantityTo:   10,
		Percentage:   10.5,
	}

	_, err := repo.Store(context.Background(), discount)
	assert.NoError(t, err)
}

func TestDiscountRepositoryFindMany(t *testing.T) {
	repo := NewDiscountRepository(db, domain.DiscountTable)
	filter := map[string]interface{}{}

	discounts, err := repo.FindMany(context.Background(), nil, filter)
	assert.NoError(t, err)
	assert.NotNil(t, discounts)
}

func TestDiscountRepositoryFindByID(t *testing.T) {
	repo := NewDiscountRepository(db, domain.DiscountTable)

	id := uuid.New()
	discount := &domain.Discount{
		ID:           id,
		Type:         domain.Land,
		QuantityFrom: 5,
		QuantityTo:   10,
		Percentage:   10.5,
	}

	newDiscount, err := repo.Store(context.Background(), discount)
	assert.NoError(t, err)

	foundDiscount, err := repo.FindByID(context.Background(), newDiscount.ID)
	assert.NoError(t, err)
	assert.Equal(t, id, foundDiscount.ID)
}

func TestDiscountRepositoryFindByType(t *testing.T) {
	repo := NewDiscountRepository(db, domain.DiscountTable)

	discount := &domain.Discount{
		ID:           uuid.New(),
		Type:         domain.Land,
		QuantityFrom: 11,
		QuantityTo:   16,
		Percentage:   12,
	}

	_, err := repo.Store(context.Background(), discount)
	assert.NoError(t, err)

	foundDiscount, err := repo.FindByTypeAndQuantity(context.Background(), domain.Land, 13)
	assert.NoError(t, err)
	assert.Equal(t, domain.Land, foundDiscount.Type)

	_, err = repo.FindByTypeAndQuantity(context.Background(), domain.Maritime, 13)
	assert.Error(t, err)
}

func TestDiscountRepositoryUpdate(t *testing.T) {
	repo := NewDiscountRepository(db, domain.DiscountTable)

	discount := &domain.Discount{
		ID:           uuid.New(),
		Type:         domain.Land,
		QuantityFrom: 17,
		QuantityTo:   21,
		Percentage:   13,
	}
	newDiscount, err := repo.Store(context.Background(), discount)
	assert.NoError(t, err)

	discount.Percentage = 15.0
	_, err = repo.Update(context.Background(), discount, newDiscount.ID)
	assert.NoError(t, err)

	updatedDiscount, _ := repo.FindByTypeAndQuantity(context.Background(), domain.Land, 19)
	assert.Equal(t, 15.0, updatedDiscount.Percentage)
}

func TestDiscountRepositoryDelete(t *testing.T) {
	repo := NewDiscountRepository(db, domain.DiscountTable)

	discount := &domain.Discount{
		Type:         domain.Land,
		QuantityFrom: 22,
		QuantityTo:   29,
		Percentage:   16,
	}

	newDiscount, err := repo.Store(context.Background(), discount)
	assert.NoError(t, err)

	err = repo.Delete(context.Background(), newDiscount.ID)
	assert.NoError(t, err)

	_, err = repo.FindByTypeAndQuantity(context.Background(), domain.Land, 23)
	assert.Error(t, err)
}

func TestDiscountRepositoryUpdateNonExistent(t *testing.T) {
	repo := NewDiscountRepository(db, domain.DiscountTable)

	discount := &domain.Discount{
		ID:           uuid.New(),
		Type:         domain.Land,
		QuantityFrom: 30,
		QuantityTo:   35,
		Percentage:   17,
	}

	_, err := repo.Update(context.Background(), discount, discount.ID)
	assert.Error(t, err)
}

func TestDiscountRepositoryDeleteNonExistent(t *testing.T) {
	repo := NewDiscountRepository(db, domain.DiscountTable)

	err := repo.Delete(context.Background(), uuid.New())
	assert.Error(t, err)
}
