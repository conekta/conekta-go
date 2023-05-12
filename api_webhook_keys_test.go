package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func TestWebhookKeysApi(t *testing.T) {
	cfg := conekta.NewConfiguration()
	cfg.Host = _basePath
	cfg.Scheme = "http"
	client := conekta.NewAPIClient(cfg)
	t.Run("GetWebhookKeys success", func(t *testing.T) {
		webhookKeys, response, err := client.WebhookKeysApi.GetWebhookKeys(context.TODO()).
			AcceptLanguage("es").
			Limit(2).
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if webhookKeys == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", webhookKeys)
		}
	})
	t.Run("GetWebhookKey success", func(t *testing.T) {
		webhookKey, response, err := client.WebhookKeysApi.GetWebhookKey(context.TODO(), "645a5eb022e7da0001cad2a4").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if webhookKey == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", webhookKey)
		}
	})
	t.Run("CreateWebhookKey success", func(t *testing.T) {
		req := conekta.WebhookKeyRequest{Active: conekta.PtrBool(true)}
		webhookKey, response, err := client.WebhookKeysApi.CreateWebhookKey(context.TODO()).
			WebhookKeyRequest(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if webhookKey == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", webhookKey)
		}
	})
	t.Run("UpdateWebhookKey success", func(t *testing.T) {
		req := conekta.WebhookKeyUpdateRequest{Active: conekta.PtrBool(false)}
		webhookKey, response, err := client.WebhookKeysApi.UpdateWebhookKey(context.TODO(), "645a613622e7da0001cad882").
			WebhookKeyUpdateRequest(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if webhookKey == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", webhookKey)
		}
	})
	t.Run("DeleteWebhookKey success", func(t *testing.T) {
		webhookKey, response, err := client.WebhookKeysApi.DeleteWebhookKey(context.TODO(), "645a59da22e7da0001cad283").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if webhookKey == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", webhookKey)
		}
	})
}
