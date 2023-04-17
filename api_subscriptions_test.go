package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func TestSubscriptionsApi(t *testing.T) {
	cfg := conekta.NewConfiguration()
	cfg.Host = "localhost:3000"
	cfg.Scheme = "http"
	client := conekta.NewAPIClient(cfg)
	t.Run("CreateSubscription success", func(t *testing.T) {
		req := conekta.SubscriptionRequest{
			PlanId:   "plan_2tZb5q8Z3PYpg6SJd",
			TrialEnd: conekta.PtrInt32(12312),
		}
		subscription, response, err := client.SubscriptionsApi.CreateSubscription(context.TODO(), "cus_2tYENskzTjjgkGQLt").
			SubscriptionRequest(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if subscription == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", subscription)
		}
	})
	t.Run("CancelSubscription success", func(t *testing.T) {
		subscription, response, err := client.SubscriptionsApi.CancelSubscription(context.TODO(), "cus_2tYENskzTjjgkGQLt").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if subscription == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", subscription)
		}
	})
	t.Run("GetSubscription success", func(t *testing.T) {
		subscription, response, err := client.SubscriptionsApi.GetSubscription(context.TODO(), "cus_2tYENskzTjjgkGQLt").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if subscription == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", subscription)
		}
	})
	t.Run("PauseSubscription success", func(t *testing.T) {
		subscription, response, err := client.SubscriptionsApi.PauseSubscription(context.TODO(), "cus_2tYENskzTjjgkGQLt").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if subscription == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", subscription)
		}
	})
	t.Run("ResumeSubscription success", func(t *testing.T) {
		subscription, response, err := client.SubscriptionsApi.ResumeSubscription(context.TODO(), "cus_2tYENskzTjjgkGQLt").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if subscription == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", subscription)
		}
	})
	t.Run("GetAllEventsFromSubscription success", func(t *testing.T) {
		events, response, err := client.SubscriptionsApi.GetAllEventsFromSubscription(context.TODO(), "cus_2tYENskzTjjgkGQLt").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if events == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", events)
		}
	})
}
