package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jordanlanch/docucenter-test/domain"
	"github.com/jordanlanch/docucenter-test/test/e2e"
)

func TestDiscounts(t *testing.T) {
	expect, teardown := e2e.Setup(t, "fixtures/discount_test")
	defer teardown()
	var accessToken string
	var discountID string

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

	t.Run("Create discount", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		discount := map[string]interface{}{
			"type":          domain.Land,
			"quantity_from": 80,
			"quantity_to":   100,
			"percentage":    20,
		}
		resp := expect.POST("/discounts").
			WithJSON(discount).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(statusCreated).
			JSON().Object()

		discountID = resp.Value("id").String().Raw()
		resp.Value("type").String().IsEqual(string(domain.Land))

		t.Log("Ending test:", t.Name())
	})

	t.Run("Fetch discounts", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		resp := expect.GET("/discounts").
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			WithQueryString("limit=10&offset=0").
			Expect().
			Status(statusOK).
			JSON().Array()

		firstDiscount := resp.Value(2).Object()
		firstDiscount.Value("id").String().IsEqual(discountID)

		t.Log("Ending test:", t.Name())
	})

	t.Run("Get discount by id", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		resp := expect.GET(fmt.Sprintf("/discounts/%s/%d", domain.Land, 90)).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(statusOK).
			JSON().Object()

		resp.Value("id").String().IsEqual(discountID)
		resp.Value("type").String().IsEqual(string(domain.Land))

		t.Log("Ending test:", t.Name())
	})

	t.Run("Modify discount", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		updatedDiscount := map[string]interface{}{
			"type":          domain.Maritime,
			"quantity_from": 20,
			"quantity_to":   200,
			"percentage":    15,
		}
		resp := expect.PUT(fmt.Sprintf("/discounts/%s", discountID)).
			WithJSON(updatedDiscount).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(statusOK).
			JSON().Object()

		resp.Value("type").String().IsEqual(string(domain.Maritime))

		t.Log("Ending test:", t.Name())
	})

	t.Run("Delete discount", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		expect.DELETE(fmt.Sprintf("/discounts/%s", discountID)).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(statusOK)

		t.Log("Ending test:", t.Name())
	})
}
