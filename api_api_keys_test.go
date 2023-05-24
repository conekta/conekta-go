package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func TestApiKeysApi(t *testing.T) {
	cfg := conekta.NewConfiguration()
	cfg.Host = _basePath
	cfg.Scheme = "http"
	client := conekta.NewAPIClient(cfg)
	t.Run("GetApiKeys success", func(t *testing.T) {
		apiKeys, response, err := client.ApiKeysApi.GetApiKeys(context.TODO()).
			Limit(20).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if apiKeys == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", apiKeys)
		}
	})
	t.Run("GetApiKey success", func(t *testing.T) {
		apiKey, response, err := client.ApiKeysApi.GetApiKey(context.TODO(), "64625cc9f3e02c00163f5e4d").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if apiKey == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", apiKey)
		}
	})
	t.Run("CreateApiKey success", func(t *testing.T) {
		rq := conekta.ApiKeyRequest{
			Active:      true,
			Description: "description",
			Role:        "admin",
		}
		apiKey, response, err := client.ApiKeysApi.CreateApiKey(context.TODO()).
			ApiKeyRequest(rq).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if apiKey == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", apiKey)
		}
	})
	t.Run("UpdateApiKey success", func(t *testing.T) {
		rq := conekta.ApiKeyUpdateRequest{
			Active:      conekta.PtrBool(false),
			Description: conekta.PtrString("description"),
		}
		apiKey, response, err := client.ApiKeysApi.UpdateApiKey(context.TODO(), "64625cc9f3e02c00163f5e4d").
			ApiKeyUpdateRequest(rq).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if apiKey == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", apiKey)
		}
	})
	t.Run("DeleteApiKey success", func(t *testing.T) {
		apiKey, response, err := client.ApiKeysApi.DeleteApiKey(context.TODO(), "64625cc9f3e02c00163f5e4d").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if apiKey == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", apiKey)
		}
	})
}
