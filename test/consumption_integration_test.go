package test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/jordanlanch/docucenter-test/domain"

	"github.com/jordanlanch/docucenter-test/test/e2e"
)

const (
	statusOK                  = http.StatusOK
	statusUnprocessableEntity = http.StatusUnprocessableEntity
	errorInvalidPayload       = "error.invalid_payload"
	errorLoginEventBaseUserID = "login.event_base.user_id"
	errorDescription          = "[ERROR] invalid payload [" + errorLoginEventBaseUserID + ": String length must be greater than or equal to 1]"
)

func TestConsumption(t *testing.T) {
	expect, teardown := e2e.Setup(t, "fixtures/login_test")
	defer teardown()
	var accessToken string

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
		obj.Value("message").String().Equal("[ERROR] invalid payload [(root): password is required]")
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
		obj.Value("message").String().Equal("User not found with the given email")
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

	t.Run("Consumption 2 meters Daily", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		resp := expect.GET("/consumption").
			WithQueryString("meters_ids=[1,2]&start_date=2023-01-01&end_date=2023-06-26&kind_period=daily&limit=5&offset=0").
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusOK).
			JSON().Object()

		resp.Keys().Contains("period", "data_graph")

		periods := resp.Value("period").Array().Iter()
		expectedPeriods := []string{"Jan 01", "Jan 02", "Jan 03", "Jan 04", "Jan 05"}
		for idx, periodVal := range periods {
			if !strings.Contains(periodVal.String().Raw(), expectedPeriods[idx]) {
				t.Errorf("Expected period %s, but got %s", expectedPeriods[idx], periodVal.String().Raw())
			}
		}

		dataGraph := resp.Value("data_graph").Array()
		dataGraph.Length().Equal(2) // Se esperan 2 objetos

		t.Log("Ending test:", t.Name())
	})
	t.Run("Consumption 2 meters Weekly", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		resp := expect.GET("/consumption").
			WithQueryString("meters_ids=[1,2]&start_date=2023-01-01&end_date=2023-06-26&kind_period=weekly&limit=5&offset=0").
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusOK).
			JSON().Object()

		resp.Keys().Contains("period", "data_graph")

		periods := resp.Value("period").Array().Iter()
		expectedPeriods := []string{"Jan 02 - Jan 08", "Jan 09 - Jan 15", "Dec 26 - Jan 01"}
		for idx, periodVal := range periods {
			if !strings.Contains(periodVal.String().Raw(), expectedPeriods[idx]) {
				t.Errorf("Expected period %s, but got %s", expectedPeriods[idx], periodVal.String().Raw())
			}
		}

		dataGraph := resp.Value("data_graph").Array()
		dataGraph.Length().Equal(2) // Se esperan 2 objetos

		t.Log("Ending test:", t.Name())
	})

	t.Run("Consumption 2 meters Monthly", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		resp := expect.GET("/consumption").
			WithQueryString("meters_ids=[1,2]&start_date=2023-01-01&end_date=2023-06-26&kind_period=monthly&limit=5&offset=0").
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusOK).
			JSON().Object()

		resp.Keys().Contains("period", "data_graph")

		periods := resp.Value("period").Array().Iter()
		expectedPeriods := []string{"Jan 2023", "Feb 2023", "Mar 2023"}
		for idx, periodVal := range periods {
			if !strings.Contains(periodVal.String().Raw(), expectedPeriods[idx]) {
				t.Errorf("Expected period %s, but got %s", expectedPeriods[idx], periodVal.String().Raw())
			}
		}

		dataGraph := resp.Value("data_graph").Array()
		dataGraph.Length().Equal(2) // Se esperan 2 objetos

		t.Log("Ending test:", t.Name())
	})

	t.Run("Consumption  meters Monthly", func(t *testing.T) {
		t.Log("Starting test:", t.Name())

		resp := expect.GET("/consumption").
			WithQueryString("meters_ids=1&start_date=2023-01-01&end_date=2023-06-26&kind_period=monthly&limit=5&offset=0").
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusOK).
			JSON().Object()

		resp.Keys().Contains("period", "data_graph")

		periods := resp.Value("period").Array().Iter()
		expectedPeriods := []string{"Jan 2023", "Feb 2023", "Mar 2023", "Apr 2023", "May 2023"}
		for idx, periodVal := range periods {
			if !strings.Contains(periodVal.String().Raw(), expectedPeriods[idx]) {
				t.Errorf("Expected period %s, but got %s", expectedPeriods[idx], periodVal.String().Raw())
			}
		}

		dataGraph := resp.Value("data_graph").Array()
		dataGraph.Length().Equal(1) // Se esperan 1 objetos

		t.Log("Ending test:", t.Name())
	})

}
