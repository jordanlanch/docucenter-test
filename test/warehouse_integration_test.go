package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jordanlanch/docucenter-test/domain"
	"github.com/jordanlanch/docucenter-test/test/e2e"
)

func TestWarehousePorts(t *testing.T) {
	expect, teardown := e2e.Setup(t, "fixtures/warehouse_port_test")
	defer teardown()
	var accessToken string
	var warehousePortID string

	t.Run("Login success", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

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
		t.Log("Ending test:", t.Name())
	})

	t.Run("Create warehouse port", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		warehousePort := map[string]interface{}{
			"name":     "Test Warehouse",
			"type":     "land",
			"location": "Test Location",
		}
		resp := expect.POST("/warehouse_ports").
			WithJSON(warehousePort).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		warehousePortID = resp.Value("id").String().Raw()
		resp.Value("name").String().IsEqual("Test Warehouse")

		t.Log("Ending test:", t.Name())
	})

	t.Run("Fetch warehouse ports", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		resp := expect.GET("/warehouse_ports").
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			WithQueryString("limit=10&offset=0").
			Expect().
			Status(http.StatusOK).
			JSON().Array()

		firstWarehousePort := resp.Value(3).Object()
		firstWarehousePort.Value("id").String().IsEqual(warehousePortID)

		t.Log("Ending test:", t.Name())
	})

	t.Run("Fetch warehouse ports with filters", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		resp := expect.GET("/warehouse_ports").
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			WithQueryString("limit=10&offset=0?type=land").
			Expect().
			Status(http.StatusOK).
			JSON().Array()

		firstWarehousePort := resp.Value(0).Object()
		firstWarehousePort.Value("type").String().IsEqual("land")

		t.Log("Ending test:", t.Name())
	})

	t.Run("Get warehouse port by id", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		resp := expect.GET(fmt.Sprintf("/warehouse_ports/%s", warehousePortID)).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusOK).
			JSON().Object()

		resp.Value("id").String().IsEqual(warehousePortID)
		resp.Value("name").String().IsEqual("Test Warehouse")

		t.Log("Ending test:", t.Name())
	})

	t.Run("Modify warehouse port", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		updatedWarehousePort := map[string]interface{}{
			"name":     "Updated Warehouse",
			"type":     "maritime",
			"location": "Updated Location",
		}
		resp := expect.PUT(fmt.Sprintf("/warehouse_ports/%s", warehousePortID)).
			WithJSON(updatedWarehousePort).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusOK).
			JSON().Object()

		resp.Value("name").String().IsEqual("Updated Warehouse")

		t.Log("Ending test:", t.Name())
	})

	t.Run("Delete warehouse port", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		expect.DELETE(fmt.Sprintf("/warehouse_ports/%s", warehousePortID)).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusOK)

		t.Log("Ending test:", t.Name())
	})
}
