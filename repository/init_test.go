package repository

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/jordanlanch/docucenter-test/domain"
	"github.com/ory/dockertest/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func splitHostPort(hostAndPort string) (string, string) {
	parts := strings.Split(hostAndPort, ":")
	return parts[0], parts[1]
}

func TestMain(m *testing.M) {
	var err error

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "13-alpine",
		Env: []string{
			"POSTGRES_USER=test",
			"POSTGRES_PASSWORD=test",
			"POSTGRES_DB=docucenter-test",
		},
	})

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// Get the host and port the container is running at
	hostAndPort := resource.GetHostPort("5432/tcp")

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		host, port := splitHostPort(hostAndPort)
		db, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=test password=test dbname=docucenter-test sslmode=disable", host, port)), &gorm.Config{})
		if err != nil {
			return err
		}

		sqlDB, err := db.DB()
		if err != nil {
			return err
		}

		return sqlDB.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// Run migrations
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Customer{})
	db.AutoMigrate(&domain.Discount{})
	db.AutoMigrate(&domain.Logistics{})
	db.AutoMigrate(&domain.Product{})
	db.AutoMigrate(&domain.WarehousesAndPorts{})

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}
