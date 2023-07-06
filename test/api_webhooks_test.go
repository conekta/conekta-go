/*
Conekta API

Testing WebhooksApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func Test_conekta_WebhooksApiService(t *testing.T) {

	configuration := conekta.NewConfiguration()
	configuration.Host = _basePath
	configuration.Scheme = "http"
	apiClient := conekta.NewAPIClient(configuration)

	t.Run("Test WebhooksApiService CreateWebhook", func(t *testing.T) {
		req := conekta.WebhookRequest{
			Url:         "https://www.fooapi.com",
			Synchronous: false,
		}
		resp, httpRes, err := apiClient.WebhooksApi.CreateWebhook(context.TODO()).
			WebhookRequest(req).
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

	t.Run("Test WebhooksApiService DeleteWebhook", func(t *testing.T) {
		resp, httpRes, err := apiClient.WebhooksApi.DeleteWebhook(context.TODO(), "641b1d5662d7e00001eaa46b").
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

	t.Run("Test WebhooksApiService GetWebhook", func(t *testing.T) {
		resp, httpRes, err := apiClient.WebhooksApi.GetWebhook(context.TODO(), "641b1d5662d7e00001eaa46b").
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

	t.Run("Test WebhooksApiService GetWebhooks", func(t *testing.T) {
		resp, httpRes, err := apiClient.WebhooksApi.GetWebhooks(context.TODO()).
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

	t.Run("Test WebhooksApiService TestWebhook", func(t *testing.T) {
		resp, httpRes, err := apiClient.WebhooksApi.TestWebhook(context.TODO(), "641b1d5662d7e00001eaa46b").
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

	t.Run("Test WebhooksApiService UpdateWebhook", func(t *testing.T) {
		req := conekta.WebhookUpdateRequest{
			Url: "https://www.fooapi.com",
		}
		resp, httpRes, err := apiClient.WebhooksApi.UpdateWebhook(context.TODO(), "641b1d5662d7e00001eaa46b").
			WebhookUpdateRequest(req).
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
