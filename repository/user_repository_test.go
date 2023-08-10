package repository

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jordanlanch/docucenter-test/domain"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	repo := NewUserRepository(db, "users")

	user := &domain.User{
		ID:    uuid.New(),
		Email: "test@test.com",
	}

	createdUser, err := repo.Create(context.Background(), user)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, createdUser.ID)
}

func TestFindByID(t *testing.T) {
	repo := NewUserRepository(db, "users")

	id := uuid.New()
	user := &domain.User{
		ID:    id,
		Email: "test@test.com",
	}
	db.Create(user)

	foundUser, err := repo.FindByID(context.Background(), id)
	assert.NoError(t, err)
	assert.Equal(t, id, foundUser.ID)

	_, err = repo.FindByID(context.Background(), uuid.New())
	assert.Error(t, err)
}

func TestFindByEmail(t *testing.T) {
	repo := NewUserRepository(db, "users")

	email := "test@test.com"
	user := &domain.User{
		ID:    uuid.New(),
		Email: email,
	}
	db.Create(user)

	foundUser, err := repo.FindByEmail(context.Background(), email)
	assert.NoError(t, err)
	assert.Equal(t, email, foundUser.Email)

	_, err = repo.FindByEmail(context.Background(), "email@email.com")
	assert.Error(t, err)
}

func TestUpdate(t *testing.T) {
	repo := NewUserRepository(db, "users")

	user := &domain.User{
		ID:    uuid.New(),
		Email: "test@test.com",
	}
	db.Create(user)

	user.Email = "updated@test.com"
	updatedUser, err := repo.Update(context.Background(), user)
	assert.NoError(t, err)
	assert.Equal(t, "updated@test.com", updatedUser.Email)
}
