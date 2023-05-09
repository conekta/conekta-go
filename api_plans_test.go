package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func TestPlansApi(t *testing.T) {
	cfg := conekta.NewConfiguration()
	cfg.Host = _basePath
	cfg.Scheme = "http"
	client := conekta.NewAPIClient(cfg)
	t.Run("GetPlan success", func(t *testing.T) {
		plan, response, err := client.PlansApi.GetPlan(context.TODO(), "plan_2tZb5q8Z3PYpg6SJd").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if plan == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", plan)
		}
		if *plan.Id != "plan_2tZb5q8Z3PYpg6SJd" {
			t.Errorf("assertion fail, expected=%v , actual=%v", "plan_2tZb5q8Z3PYpg6SJd", *plan.Id)
		}
	})
	t.Run("GetPlans success", func(t *testing.T) {
		plans, response, err := client.PlansApi.GetPlans(context.TODO()).
			Limit(10).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if plans == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", plans)
		}
	})
	t.Run("CreatePlan success", func(t *testing.T) {
		req := conekta.PlanRequest{
			Amount:   1000,
			Currency: conekta.PtrString("MXN"),
			Name:     "Test Plan",
		}
		plan, response, err := client.PlansApi.CreatePlan(context.TODO()).
			PlanRequest(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if plan == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", plan)
		}
	})
	t.Run("UpdatePlan success", func(t *testing.T) {
		req := conekta.PlanUpdateRequest{
			Amount:   conekta.PtrInt32(1000),
			Currency: conekta.PtrString("MXN"),
			Name:     conekta.PtrString("Test Plan"),
		}
		plan, response, err := client.PlansApi.UpdatePlan(context.TODO(), "plan_2tZb5q8Z3PYpg6SJd").
			PlanUpdateRequest(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if plan == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", plan)
		}
	})
	t.Run("DeletePlan success", func(t *testing.T) {
		plan, response, err := client.PlansApi.DeletePlan(context.TODO(), "plan_2tZb5q8Z3PYpg6SJd").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if plan == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", plan)
		}
	})
}
