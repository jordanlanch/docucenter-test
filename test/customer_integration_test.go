package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jordanlanch/docucenter-test/domain"
	"github.com/jordanlanch/docucenter-test/test/e2e"
)

const (
	statusOK                  = http.StatusOK
	statusBadRequest          = http.StatusBadRequest
	statusNotFound            = http.StatusNotFound
	statusCreated             = http.StatusCreated
	errorMissingIDParam       = "Missing id parameter"
	errorInvalidPayload       = "error.invalid_payload"
	errorLoginEventBaseUserID = "login.event_base.user_id"
)

func TestCustomers(t *testing.T) {
	expect, teardown := e2e.Setup(t, "fixtures/customer_test")
	defer teardown()
	var accessToken string
	var customerID string

	t.Run("Signup error ", func(t *testing.T) {
		t.Log("Iniciando prueba:", t.Name())
		in := domain.SignupRequest{
			Email: "jordan.dev93@gmail.com",
		}
		obj := expect.POST("/signup").
			WithJSON(in).
			Expect().
			Status(http.StatusBadRequest).
			JSON().Object()
		obj.Value("message").String().IsEqual("[ERROR] invalid payload [(root): password is required]")
		t.Log("Finalizando prueba:", t.Name())
	})

	t.Run("Login user not found ", func(t *testing.T) {
		t.Log("Iniciando prueba:", t.Name())
		in := domain.LoginRequest{
			Email: "jordan.dev93@gmail.com",
		}
		obj := expect.POST("/login").
			WithJSON(in).
			Expect().
			Status(http.StatusNotFound).
			JSON().Object()
		obj.Value("message").String().IsEqual("User not found with the given email")
		t.Log("Finalizando prueba:", t.Name())
	})

	t.Run("Signup success", func(t *testing.T) {
		t.Log("Iniciando prueba:", t.Name())

		in := domain.SignupRequest{
			Email:    "jordan.dev93@gmail.com",
			Password: "pass1234*",
		}

		obj := expect.POST("/signup").
			WithJSON(in).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		obj.Value("accessToken").String().NotEqual("")
		t.Log("Finalizando prueba:", t.Name())
	})

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
			"name":  "John Doe",
			"email": "john.doe@example.com",
		}
		resp := expect.POST("/customers").
			WithJSON(customer).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		customerID = resp.Value("id").String().Raw()
		resp.Value("name").String().IsEqual("John Doe")
		resp.Value("email").String().IsEqual("john.doe@example.com")

		t.Log("Ending test:", t.Name())
	})

	t.Run("Get customer by id", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		resp := expect.GET(fmt.Sprintf("/customers/%s", customerID)).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		resp.Value("id").String().IsEqual(customerID)

		t.Log("Ending test:", t.Name())
	})

	t.Run("Modify customer", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		customer := map[string]interface{}{
			"name":  "Jane Doe",
			"email": "jane.doe@example.com",
		}
		resp := expect.PUT(fmt.Sprintf("/customers/%s", customerID)).
			WithJSON(customer).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		resp.Value("name").String().IsEqual("Jane Doe")
		resp.Value("email").String().IsEqual("jane.doe@example.com")

		t.Log("Ending test:", t.Name())
	})

	t.Run("Fetch customers", func(t *testing.T) {
		t.Log("Starting test:", t.Name())
		resp := expect.GET("/customers").
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			WithQueryString("limit=10&offset=0").
			Expect().
			Status(http.StatusOK).
			JSON().Array()

		// 2 elements has been in seeds
		resp.Length().IsEqual(3)

		firstCustomer := resp.Value(2).Object()
		firstCustomer.Value("id").String().IsEqual(customerID)
		firstCustomer.Value("name").String().IsEqual("Jane Doe")
		firstCustomer.Value("email").String().IsEqual("jane.doe@example.com")
		t.Log("Ending test:", t.Name())
	})

	t.Run("Fetch customers with filter", func(t *testing.T) {
		t.Log("Starting test:", t.Name())
		resp := expect.GET("/customers").
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			WithQueryString("limit=10&offset=0&email=contacto@empresaabc.com").
			Expect().
			Status(http.StatusOK).
			JSON().Array()

		resp.Length().IsEqual(1)

		firstCustomer := resp.Value(0).Object()
		firstCustomer.Value("email").String().IsEqual("contacto@empresaabc.com")
		t.Log("Ending test:", t.Name())
	})

	t.Run("Delete customer", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		expect.DELETE(fmt.Sprintf("/customers/%s", customerID)).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusOK)

		t.Log("Ending test:", t.Name())
	})
}
