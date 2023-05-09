package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func TestWebhooksApi(t *testing.T) {
	cfg := conekta.NewConfiguration()
	cfg.Host = _basePath
	cfg.Scheme = "http"
	client := conekta.NewAPIClient(cfg)
	t.Run("CreateWebhook success", func(t *testing.T) {
		req := conekta.WebhookRequest{
			Url:         "https://www.fooapi.com",
			Synchronous: false,
		}
		webhook, response, err := client.WebhooksApi.CreateWebhook(context.TODO()).
			WebhookRequest(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if webhook == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", webhook)
		}
	})
	t.Run("UpdateWebhook success", func(t *testing.T) {
		req := conekta.WebhookUpdateRequest{
			Url: "https://www.fooapi.com",
		}
		webhook, response, err := client.WebhooksApi.UpdateWebhook(context.TODO(), "641b1d5662d7e00001eaa46b").
			WebhookUpdateRequest(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if webhook == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", webhook)
		}
	})
	t.Run("DeleteWebhook success", func(t *testing.T) {
		webhook, response, err := client.WebhooksApi.DeleteWebhook(context.TODO(), "641b1d5662d7e00001eaa46b").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if webhook == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", webhook)
		}
	})
	t.Run("GetWebhook success", func(t *testing.T) {
		webhook, response, err := client.WebhooksApi.GetWebhook(context.TODO(), "641b1d5662d7e00001eaa46b").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if webhook == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", webhook)
		}
	})
	t.Run("GetWebhooks success", func(t *testing.T) {
		webhooks, response, err := client.WebhooksApi.GetWebhooks(context.TODO()).
			Limit(10).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if webhooks == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", webhooks)
		}
	})
}
