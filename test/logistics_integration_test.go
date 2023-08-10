package test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/jordanlanch/docucenter-test/domain"
	"github.com/jordanlanch/docucenter-test/test/e2e"
)

func TestLogistics(t *testing.T) {
	expect, teardown := e2e.Setup(t, "fixtures/logistics_test")
	defer teardown()

	var accessToken string
	var logisticsID string
	var customerID string
	var productID string
	var warehousePortID string

	t.Run("Login success", func(t *testing.T) {
		t.Log("Iniciando prueba:", t.Name())

		in := domain.LoginRequest{
			Email:    "jordan.dev93@gmail.com",
			Password: "pass1234*",
		}

		obj := expect.POST("/login").
			WithJSON(in).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		accessToken = obj.Value("accessToken").String().Raw()
		obj.Value("accessToken").String().NotEqual("")
		t.Log("Finalizando prueba:", t.Name())
	})

	t.Run("Create customer", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		customer := map[string]interface{}{
			"name":  "Jordan",
			"email": "jordan.doe@example.com",
		}
		resp := expect.POST("/customers").
			WithJSON(customer).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		customerID = resp.Value("id").String().Raw()
		resp.Value("name").String().IsEqual("Jordan")
		resp.Value("email").String().IsEqual("jordan.doe@example.com")

		t.Log("Ending test:", t.Name())
	})

	t.Run("Create product", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		customer := map[string]interface{}{
			"name": "product 1",
			"type": "land",
		}
		resp := expect.POST("/products").
			WithJSON(customer).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		productID = resp.Value("id").String().Raw()
		resp.Value("name").String().IsEqual("product 1")
		resp.Value("type").String().IsEqual("land")

		t.Log("Ending test:", t.Name())
	})

	t.Run("Create warehouse", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		customer := map[string]interface{}{
			"name":     "land 1",
			"type":     "land",
			"location": "calle 1 N 1 1",
		}
		resp := expect.POST("/warehouse_ports").
			WithJSON(customer).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		warehousePortID = resp.Value("id").String().Raw()
		resp.Value("name").String().IsEqual("land 1")
		resp.Value("type").String().IsEqual("land")
		resp.Value("location").String().IsEqual("calle 1 N 1 1")

		t.Log("Ending test:", t.Name())
	})

	t.Run("Create logistics Land", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		logistics := map[string]interface{}{
			"customer_id":       customerID,
			"product_id":        productID,
			"warehouse_port_id": warehousePortID,
			"type":              domain.Land,
			"quantity":          10,
			"registration_date": time.Now(),
			"delivery_date":     time.Now(),
			"shipping_price":    100.00,
			"vehicle_plate":     "ABC123",
			"guide_number":      "1234567890",
			// Add other necessary fields
		}

		resp := expect.POST("/logistics").
			WithJSON(logistics).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		logisticsID = resp.Value("id").String().Raw()
		resp.Value("customer_id").String().IsEqual(logistics["customer_id"].(string))

		t.Log("Ending test:", t.Name())
	})

	t.Run("Create logistics", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		logistics := map[string]interface{}{
			"customer_id":       customerID,
			"product_id":        productID,
			"warehouse_port_id": warehousePortID,
			"type":              domain.Maritime,
			"quantity":          15,
			"registration_date": time.Now(),
			"delivery_date":     time.Now(),
			"shipping_price":    100.00,
			"fleet_number":      "DEF1234H",
			"guide_number":      "9876543210",
			// Add other necessary fields
		}

		resp := expect.POST("/logistics").
			WithJSON(logistics).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		logisticsID = resp.Value("id").String().Raw()
		resp.Value("customer_id").String().IsEqual(logistics["customer_id"].(string))

		t.Log("Ending test:", t.Name())
	})

	// Logistics - Fetch
	t.Run("Fetch logistics", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		resp := expect.GET("/logistics").
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			WithQueryString("limit=10&offset=0").
			Expect().
			Status(http.StatusOK).
			JSON().Array()

		firstLogistic := resp.Value(3).Object()
		firstLogistic.Value("id").String().IsEqual(logisticsID)

		t.Log("Ending test:", t.Name())
	})

	t.Run("Fetch logistics With Filter", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		resp := expect.GET("/logistics").
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			WithQueryString("limit=10&offset=0&vehicle_plate=ABC123").
			Expect().
			Status(http.StatusOK).
			JSON().Array()

		firstLogistic := resp.Value(0).Object()
		firstLogistic.Value("vehicle_plate").String().IsEqual("ABC123")

		t.Log("Ending test:", t.Name())
	})

	// Logistics - Get by ID
	t.Run("Get logistic by id", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		resp := expect.GET(fmt.Sprintf("/logistics/%s", logisticsID)).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusOK).
			JSON().Object()

		resp.Value("id").String().IsEqual(logisticsID)

		t.Log("Ending test:", t.Name())
	})

	// Logistics - Modify
	t.Run("Modify logistics", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		updatedLogistic := map[string]interface{}{
			"customer_id":       customerID,
			"product_id":        productID,
			"warehouse_port_id": warehousePortID,
			"type":              domain.Maritime,
			"quantity":          20,
			"registration_date": time.Now(),
			"delivery_date":     time.Now(),
			"shipping_price":    200.00,
			"fleet_number":      "DEF1234H",
			"guide_number":      "85217963",
		}

		resp := expect.PUT(fmt.Sprintf("/logistics/%s", logisticsID)).
			WithJSON(updatedLogistic).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusOK).
			JSON().Object()

		resp.Value("quantity").Number().IsEqual(20)

		t.Log("Ending test:", t.Name())
	})

	// Logistics - Delete
	t.Run("Delete logistics", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		expect.DELETE(fmt.Sprintf("/logistics/%s", logisticsID)).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusOK)

		t.Log("Ending test:", t.Name())
	})
}
