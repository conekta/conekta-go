package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func TestEventsApi(t *testing.T) {
	cfg := conekta.NewConfiguration()
	cfg.Host = _basePath
	cfg.Scheme = "http"
	client := conekta.NewAPIClient(cfg)
	t.Run("GetEvent success", func(t *testing.T) {
		event, response, err := client.EventsApi.GetEvent(context.TODO(), "63fe3d2ddf70970001cfb41a").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if event == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", event)
		}
	})
	t.Run("GetEvents success", func(t *testing.T) {
		events, response, err := client.EventsApi.GetEvents(context.TODO()).
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
