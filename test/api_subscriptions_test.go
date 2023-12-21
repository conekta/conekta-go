/*
Conekta API

Testing SubscriptionsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func Test_conekta_SubscriptionsApiService(t *testing.T) {

	configuration := conekta.NewConfiguration()
	configuration.Host = _basePath
	configuration.Scheme = "http"
	apiClient := conekta.NewAPIClient(configuration)

	t.Run("Test SubscriptionsApiService CancelSubscription", func(t *testing.T) {
		resp, httpRes, err := apiClient.SubscriptionsAPI.CancelSubscription(context.TODO(), "cus_2tYENskzTjjgkGQLt").
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

	t.Run("Test SubscriptionsApiService CreateSubscription", func(t *testing.T) {
		req := conekta.SubscriptionRequest{
			PlanId:   "plan_2tZb5q8Z3PYpg6SJd",
			TrialEnd: conekta.PtrInt32(12312),
		}
		resp, httpRes, err := apiClient.SubscriptionsAPI.CreateSubscription(context.TODO(), "cus_2tYENskzTjjgkGQLt").
			SubscriptionRequest(req).
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

	t.Run("Test SubscriptionsApiService GetAllEventsFromSubscription", func(t *testing.T) {
		resp, httpRes, err := apiClient.SubscriptionsAPI.GetAllEventsFromSubscription(context.TODO(), "cus_2tYENskzTjjgkGQLt").
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

	t.Run("Test SubscriptionsApiService GetSubscription", func(t *testing.T) {
		resp, httpRes, err := apiClient.SubscriptionsAPI.GetSubscription(context.TODO(), "cus_2tYENskzTjjgkGQLt").
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

	t.Run("Test SubscriptionsApiService PauseSubscription", func(t *testing.T) {
		resp, httpRes, err := apiClient.SubscriptionsAPI.PauseSubscription(context.TODO(), "cus_2tYENskzTjjgkGQLt").
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

	t.Run("Test SubscriptionsApiService ResumeSubscription", func(t *testing.T) {
		resp, httpRes, err := apiClient.SubscriptionsAPI.ResumeSubscription(context.TODO(), "cus_2tYENskzTjjgkGQLt").
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

	t.Run("Test SubscriptionsApiService UpdateSubscription", func(t *testing.T) {
		req := conekta.SubscriptionUpdateRequest{
			PlanId:   conekta.PtrString("plan_2tZb5q8Z3PYpg6SJd"),
			TrialEnd: conekta.PtrInt32(12312),
		}

		resp, httpRes, err := apiClient.SubscriptionsAPI.UpdateSubscription(context.Background(), "cus_2tYENskzTjjgkGQLt").
			SubscriptionUpdateRequest(req).
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
