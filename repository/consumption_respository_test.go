package repository

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"
// 	"testing"

// 	"github.com/google/uuid"
// 	"github.com/jordanlanch/docucenter-test/domain"
// 	"github.com/ory/dockertest/v3"
// 	"github.com/stretchr/testify/assert"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"k8s.io/utils/pointer"
// )

// var db *gorm.DB

// func splitHostPort(hostAndPort string) (string, string) {
// 	parts := strings.Split(hostAndPort, ":")
// 	return parts[0], parts[1]
// }

// func TestMain(m *testing.M) {
// 	var err error

// 	pool, err := dockertest.NewPool("")
// 	if err != nil {
// 		log.Fatalf("Could not connect to docker: %s", err)
// 	}

// 	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
// 		Repository: "postgres",
// 		Tag:        "13-alpine",
// 		Env: []string{
// 			"POSTGRES_USER=test",
// 			"POSTGRES_PASSWORD=test",
// 			"POSTGRES_DB=test",
// 		},
// 	})

// 	if err != nil {
// 		log.Fatalf("Could not start resource: %s", err)
// 	}

// 	// Get the host and port the container is running at
// 	hostAndPort := resource.GetHostPort("5432/tcp")

// 	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
// 	if err := pool.Retry(func() error {
// 		var err error
// 		host, port := splitHostPort(hostAndPort)
// 		db, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=test password=test dbname=test sslmode=disable", host, port)), &gorm.Config{})
// 		if err != nil {
// 			return err
// 		}

// 		sqlDB, err := db.DB()
// 		if err != nil {
// 			return err
// 		}

// 		return sqlDB.Ping()
// 	}); err != nil {
// 		log.Fatalf("Could not connect to docker: %s", err)
// 	}

// 	// Run migrations
// 	db.AutoMigrate(&domain.User{})
// 	db.AutoMigrate(&domain.Meter{})
// 	db.AutoMigrate(&domain.Consumption{})

// 	code := m.Run()

// 	// You can't defer this because os.Exit doesn't care for defer
// 	if err := pool.Purge(resource); err != nil {
// 		log.Fatalf("Could not purge resource: %s", err)
// 	}

// 	os.Exit(code)
// }

// func TestNewConsumptionRepository(t *testing.T) {
// 	repo := NewConsumptionRepository(db, "consumptions")
// 	assert.NotNil(t, repo)
// }
// func TestIsValidPeriodType(t *testing.T) {
// 	assert.True(t, isValidPeriodType("daily"))
// 	assert.False(t, isValidPeriodType("notvalid"))
// }

// func TestGetDatePeriodFormat(t *testing.T) {
// 	format, sql := getDatePeriodFormat("daily")
// 	assert.Equal(t, "TO_CHAR(timestamp, 'Mon DD') AS period", format)
// 	assert.Equal(t, "TO_CHAR(timestamp, 'Mon DD')", sql)

// 	format, sql = getDatePeriodFormat("notvalid")
// 	assert.Equal(t, "", format)
// 	assert.Equal(t, "", sql)
// }

// func TestGetUniqueSortedPeriods(t *testing.T) {
// 	meterDataMap := make(map[int][]MeterData)
// 	meterDataMap[1] = []MeterData{
// 		{Period: "Jul 01"},
// 		{Period: "Aug 01"},
// 	}
// 	meterDataMap[2] = []MeterData{
// 		{Period: "Aug 01"},
// 		{Period: "Sep 01"},
// 	}

// 	periods := getUniqueSortedPeriods(meterDataMap)
// 	expected := []string{"Jul 01", "Aug 01", "Sep 01"}
// 	assert.Equal(t, expected, periods)
// }

// func TestFormatConsumptionData(t *testing.T) {
// 	consumptions := []*domain.Consumption{
// 		{
// 			MeterID:            123,
// 			Active:             100.5,
// 			ReactiveInductive:  200.5,
// 			ReactiveCapacitive: 300.5,
// 			Exported:           400.5,
// 			Period:             "Jul 2023",
// 		},
// 	}

// 	response := formatConsumptionData(consumptions)
// 	assert.NotNil(t, response)
// 	assert.Equal(t, 1, len(response.DataGraph))
// 	assert.Equal(t, 123, response.DataGraph[0]["meter_id"])
// }

// func TestParsePeriodDate(t *testing.T) {
// 	// Test when period is in "Mon DD" format
// 	date := parsePeriodDate("Jul 01")
// 	assert.Equal(t, "01 Jul 00 00:00 UTC", date.Format("02 Jan 06 15:04 MST"))

// 	// Test when period is in "Mon YYYY" format
// 	date = parsePeriodDate("Jul 2023")
// 	assert.Equal(t, "01 Jul 23 00:00 UTC", date.Format("02 Jan 06 15:04 MST"))

// 	// Test when period is in "Mon DD - Mon DD" format
// 	date = parsePeriodDate("Jul 01 - Aug 01")
// 	assert.Equal(t, "01 Jul 00 00:00 UTC", date.Format("02 Jan 06 15:04 MST"))

// 	// Test when period is in an unrecognized format
// 	date = parsePeriodDate("Invalid Date")
// 	assert.Equal(t, "01 Jan 01 00:00 UTC", date.Format("02 Jan 06 15:04 MST"))
// }

// func TestSave(t *testing.T) {
// 	repo := NewConsumptionRepository(db, "consumptions")

// 	consumption := &domain.Consumption{
// 		ID:                 uuid.New(),
// 		MeterID:            123,
// 		Active:             100.5,
// 		ReactiveInductive:  200.5,
// 		ReactiveCapacitive: 300.5,
// 		Exported:           400.5,
// 		Period:             "Jul 2023",
// 	}

// 	err := repo.Save(context.Background(), consumption)
// 	assert.NoError(t, err)
// }

// func TestFindByIDConsumption(t *testing.T) {
// 	repo := NewConsumptionRepository(db, "consumptions")

// 	id := uuid.New()
// 	consumption := &domain.Consumption{
// 		ID:                 id,
// 		MeterID:            123,
// 		Active:             100.5,
// 		ReactiveInductive:  200.5,
// 		ReactiveCapacitive: 300.5,
// 		Exported:           400.5,
// 		Period:             "Jul 2023",
// 	}
// 	db.Create(consumption)

// 	foundConsumption, err := repo.FindByID(context.Background(), id)
// 	assert.NoError(t, err)
// 	assert.Equal(t, id, foundConsumption.ID)
// 	_, err = repo.FindByID(context.Background(), uuid.New())
// 	assert.Error(t, err)
// }

// func TestFindByPeriodInvalidPeriod(t *testing.T) {
// 	repo := NewConsumptionRepository(db, "consumptions")

// 	consumption1 := &domain.Consumption{
// 		ID:                 uuid.New(),
// 		MeterID:            123,
// 		Active:             100.5,
// 		ReactiveInductive:  200.5,
// 		ReactiveCapacitive: 300.5,
// 		Exported:           400.5,
// 		Period:             "Jan 01",
// 	}

// 	consumption2 := &domain.Consumption{
// 		ID:                 uuid.New(),
// 		MeterID:            124,
// 		Active:             150.5,
// 		ReactiveInductive:  250.5,
// 		ReactiveCapacitive: 350.5,
// 		Exported:           450.5,
// 		Period:             "Feb 09",
// 	}

// 	db.Create(consumption1)
// 	db.Create(consumption2)

// 	periodType := "invalidPeriod"
// 	start := "2023-01-01"
// 	end := "2023-02-02"
// 	meterIDs := []int{123, 124}
// 	pagination := &domain.Pagination{Limit: pointer.Int(5), Offset: pointer.Int(0)}

// 	_, err := repo.FindByPeriod(context.Background(), periodType, start, end, meterIDs, pagination)
// 	assert.Error(t, err)
// }

// func TestFindByPeriodDaily(t *testing.T) {
// 	repo := NewConsumptionRepository(db, "consumptions")

// 	consumption1 := &domain.Consumption{
// 		ID:                 uuid.New(),
// 		MeterID:            123,
// 		Active:             100.5,
// 		ReactiveInductive:  200.5,
// 		ReactiveCapacitive: 300.5,
// 		Exported:           400.5,
// 		Period:             "Jan 01",
// 	}

// 	consumption2 := &domain.Consumption{
// 		ID:                 uuid.New(),
// 		MeterID:            124,
// 		Active:             150.5,
// 		ReactiveInductive:  250.5,
// 		ReactiveCapacitive: 350.5,
// 		Exported:           450.5,
// 		Period:             "Feb 09",
// 	}

// 	db.Create(consumption1)
// 	db.Create(consumption2)

// 	periodType := "daily"
// 	start := "2023-01-01"
// 	end := "2023-02-02"
// 	meterIDs := []int{123, 124}
// 	pagination := &domain.Pagination{Limit: pointer.Int(5), Offset: pointer.Int(0)}

// 	response, err := repo.FindByPeriod(context.Background(), periodType, start, end, meterIDs, pagination)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, response)
// }

// func TestFindByPeriodWeekly(t *testing.T) {
// 	repo := NewConsumptionRepository(db, "consumptions")

// 	consumption1 := &domain.Consumption{
// 		ID:                 uuid.New(),
// 		MeterID:            123,
// 		Active:             100.5,
// 		ReactiveInductive:  200.5,
// 		ReactiveCapacitive: 300.5,
// 		Exported:           400.5,
// 		Period:             "Jan 02 - Jan 08",
// 	}

// 	consumption2 := &domain.Consumption{
// 		ID:                 uuid.New(),
// 		MeterID:            124,
// 		Active:             150.5,
// 		ReactiveInductive:  250.5,
// 		ReactiveCapacitive: 350.5,
// 		Exported:           450.5,
// 		Period:             "Feb 06 - Feb 12",
// 	}

// 	db.Create(consumption1)
// 	db.Create(consumption2)

// 	periodType := "weekly"
// 	start := "2023-01-01"
// 	end := "2023-02-02"
// 	meterIDs := []int{123, 124}
// 	pagination := &domain.Pagination{Limit: pointer.Int(5), Offset: pointer.Int(0)}

// 	response, err := repo.FindByPeriod(context.Background(), periodType, start, end, meterIDs, pagination)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, response)
// }

// func TestFindByPeriodMonthly(t *testing.T) {
// 	repo := NewConsumptionRepository(db, "consumptions")

// 	consumption1 := &domain.Consumption{
// 		ID:                 uuid.New(),
// 		MeterID:            123,
// 		Active:             100.5,
// 		ReactiveInductive:  200.5,
// 		ReactiveCapacitive: 300.5,
// 		Exported:           400.5,
// 		Period:             "Jan 2023",
// 	}

// 	consumption2 := &domain.Consumption{
// 		ID:                 uuid.New(),
// 		MeterID:            124,
// 		Active:             150.5,
// 		ReactiveInductive:  250.5,
// 		ReactiveCapacitive: 350.5,
// 		Exported:           450.5,
// 		Period:             "Feb 2023",
// 	}

// 	db.Create(consumption1)
// 	db.Create(consumption2)

// 	periodType := "monthly"
// 	start := "2023-01-01"
// 	end := "2023-02-02"
// 	meterIDs := []int{123, 124}
// 	pagination := &domain.Pagination{Limit: pointer.Int(5), Offset: pointer.Int(0)}

// 	response, err := repo.FindByPeriod(context.Background(), periodType, start, end, meterIDs, pagination)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, response)
// }
