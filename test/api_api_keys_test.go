/*
Conekta API

Testing ApiKeysApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func Test_conekta_ApiKeysApiService(t *testing.T) {

	configuration := conekta.NewConfiguration()
	configuration.Host = _basePath
	configuration.Scheme = "http"
	apiClient := conekta.NewAPIClient(configuration)

	t.Run("Test ApiKeysApiService CreateApiKey", func(t *testing.T) {
		rq := conekta.ApiKeyRequest{
			Description: conekta.PtrString("description"),
			Role:        "admin",
		}
		resp, httpRes, err := apiClient.ApiKeysAPI.CreateApiKey(context.TODO()).
			ApiKeyRequest(rq).
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

	t.Run("Test ApiKeysApiService DeleteApiKey", func(t *testing.T) {
		resp, httpRes, err := apiClient.ApiKeysAPI.DeleteApiKey(context.TODO(), "64625cc9f3e02c00163f5e4d").
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

	t.Run("Test ApiKeysApiService GetApiKey", func(t *testing.T) {
		resp, httpRes, err := apiClient.ApiKeysAPI.GetApiKey(context.TODO(), "64625cc9f3e02c00163f5e4d").
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

	t.Run("Test ApiKeysApiService GetApiKeys", func(t *testing.T) {
		resp, httpRes, err := apiClient.ApiKeysAPI.GetApiKeys(context.TODO()).
			Limit(20).
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

	t.Run("Test ApiKeysApiService UpdateApiKey", func(t *testing.T) {
		rq := conekta.ApiKeyUpdateRequest{
			Active:      conekta.PtrBool(false),
			Description: conekta.PtrString("description"),
		}
		resp, httpRes, err := apiClient.ApiKeysAPI.UpdateApiKey(context.TODO(), "64625cc9f3e02c00163f5e4d").
			ApiKeyUpdateRequest(rq).
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
