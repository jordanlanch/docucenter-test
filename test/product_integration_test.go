package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/jordanlanch/docucenter-test/domain"
	"github.com/jordanlanch/docucenter-test/test/e2e"
)

func TestProducts(t *testing.T) {
	expect, teardown := e2e.Setup(t, "fixtures/product_test")
	defer teardown()
	var accessToken string
	var productID uuid.UUID

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

	t.Run("Create product", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		product := map[string]interface{}{
			"name": "SampleProduct",
			"type": "land",
		}
		resp := expect.POST("/products").
			WithJSON(product).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		productID = uuid.MustParse(resp.Value("id").String().Raw())
		resp.Value("name").String().IsEqual("SampleProduct")
		t.Log("Ending test:", t.Name())
	})

	t.Run("Fetch products", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		resp := expect.GET("/products").
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			WithQueryString("limit=10&offset=0").
			Expect().
			Status(http.StatusOK).
			JSON().Array()

		firstProduct := resp.Value(3).Object()
		firstProduct.Value("id").String().IsEqual(productID.String())

		t.Log("Ending test:", t.Name())
	})

	t.Run("Fetch products with Filters", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		resp := expect.GET("/products").
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			WithQueryString("limit=10&offset=0&type=land").
			Expect().
			Status(http.StatusOK).
			JSON().Array()

		firstProduct := resp.Value(0).Object()
		firstProduct.Value("type").String().IsEqual("land")

		t.Log("Ending test:", t.Name())
	})

	t.Run("Get product by id", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		resp := expect.GET(fmt.Sprintf("/products/%s", productID)).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusOK).
			JSON().Object()

		resp.Value("id").String().IsEqual(productID.String())
		resp.Value("name").String().IsEqual("SampleProduct")

		t.Log("Ending test:", t.Name())
	})

	t.Run("Modify product", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		updatedProduct := map[string]interface{}{
			"name": "UpdatedProduct",
			"type": "maritime",
		}
		resp := expect.PUT(fmt.Sprintf("/products/%s", productID)).
			WithJSON(updatedProduct).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusOK).
			JSON().Object()

		resp.Value("name").String().IsEqual("UpdatedProduct")

		t.Log("Ending test:", t.Name())
	})

	t.Run("Delete product", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		expect.DELETE(fmt.Sprintf("/products/%s", productID)).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusOK)

		t.Log("Ending test:", t.Name())
	})
}
