package repository

// import (
// 	"context"
// 	"testing"

// 	"github.com/google/uuid"
// 	"github.com/jordanlanch/docucenter-test/domain"
// 	"github.com/stretchr/testify/assert"
// 	"k8s.io/utils/pointer"
// )

// func TestFindMeterByID(t *testing.T) {
// 	repo := NewMeterRepository(db, "meters")

// 	id := uuid.New()
// 	meter := &domain.Meter{
// 		ID:      id,
// 		Address: "Mock Address",
// 	}
// 	db.Create(meter)

// 	foundMeter, err := repo.FindByID(context.Background(), id)
// 	assert.NoError(t, err)
// 	assert.Equal(t, id, foundMeter.ID)
// 	_, err = repo.FindByID(context.Background(), uuid.New())
// 	assert.Error(t, err)
// }

// func TestFindManyMeters(t *testing.T) {
// 	repo := NewMeterRepository(db, "meters")

// 	meters := make([]*domain.Meter, 10)
// 	for i := range meters {
// 		meters[i] = &domain.Meter{
// 			ID:      uuid.New(),
// 			Address: "Mock Address",
// 		}
// 		db.Create(meters[i])
// 	}

// 	pagination := &domain.Pagination{Limit: pointer.Int(5), Offset: pointer.Int(0)}
// 	foundMeters, err := repo.FindMany(context.Background(), pagination)
// 	assert.NoError(t, err)
// 	assert.Len(t, foundMeters, 5)
// }
