/*
Conekta API

Testing PlansApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func Test_conekta_PlansApiService(t *testing.T) {

	configuration := conekta.NewConfiguration()
	configuration.Host = _basePath
	configuration.Scheme = "http"
	apiClient := conekta.NewAPIClient(configuration)

	t.Run("Test PlansApiService CreatePlan", func(t *testing.T) {
		req := conekta.PlanRequest{
			Amount:   1000,
			Currency: conekta.PtrString("MXN"),
			Name:     "Test Plan",
		}
		resp, httpRes, err := apiClient.PlansApi.CreatePlan(context.TODO()).
			PlanRequest(req).
			AcceptLanguage("es").
			Execute()

		if err != nil {
			t.Fatalf("expected err nil but was %v", err)
		}
		if resp == nil {
			t.Fatalf("expected resp not nil but was %v", resp)
		}
		if httpRes.StatusCode != http.StatusOK {
			t.Fatalf("expected StatusCode 200 OK but was %d", httpRes.StatusCode)
		}

	})

	t.Run("Test PlansApiService DeletePlan", func(t *testing.T) {
		resp, httpRes, err := apiClient.PlansApi.DeletePlan(context.TODO(), "plan_2tZb5q8Z3PYpg6SJd").
			AcceptLanguage("es").
			Execute()

		if err != nil {
			t.Fatalf("expected err nil but was %v", err)
		}
		if resp == nil {
			t.Fatalf("expected resp not nil but was %v", resp)
		}
		if httpRes.StatusCode != http.StatusOK {
			t.Fatalf("expected StatusCode 200 OK but was %d", httpRes.StatusCode)
		}

	})

	t.Run("Test PlansApiService GetPlan", func(t *testing.T) {
		resp, httpRes, err := apiClient.PlansApi.GetPlan(context.TODO(), "plan_2tZb5q8Z3PYpg6SJd").
			AcceptLanguage("es").
			Execute()

		if err != nil {
			t.Fatalf("expected err nil but was %v", err)
		}
		if resp == nil {
			t.Fatalf("expected resp not nil but was %v", resp)
		}
		if httpRes.StatusCode != http.StatusOK {
			t.Fatalf("expected StatusCode 200 OK but was %d", httpRes.StatusCode)
		}

	})

	t.Run("Test PlansApiService GetPlans", func(t *testing.T) {
		resp, httpRes, err := apiClient.PlansApi.GetPlans(context.TODO()).
			Limit(10).
			AcceptLanguage("es").
			Execute()

		if err != nil {
			t.Fatalf("expected err nil but was %v", err)
		}
		if resp == nil {
			t.Fatalf("expected resp not nil but was %v", resp)
		}
		if httpRes.StatusCode != http.StatusOK {
			t.Fatalf("expected StatusCode 200 OK but was %d", httpRes.StatusCode)
		}

	})

	t.Run("Test PlansApiService UpdatePlan", func(t *testing.T) {
		req := conekta.PlanUpdateRequest{
			Amount:   conekta.PtrInt32(1000),
			Currency: conekta.PtrString("MXN"),
			Name:     conekta.PtrString("Test Plan"),
		}
		resp, httpRes, err := apiClient.PlansApi.UpdatePlan(context.TODO(), "plan_2tZb5q8Z3PYpg6SJd").
			PlanUpdateRequest(req).
			AcceptLanguage("es").
			Execute()

		if err != nil {
			t.Fatalf("expected err nil but was %v", err)
		}
		if resp == nil {
			t.Fatalf("expected resp not nil but was %v", resp)
		}
		if httpRes.StatusCode != http.StatusOK {
			t.Fatalf("expected StatusCode 200 OK but was %d", httpRes.StatusCode)
		}

	})

}
