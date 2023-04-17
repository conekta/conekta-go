package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func TestBlacklistApi(t *testing.T) {
	cfg := conekta.NewConfiguration()
	cfg.Host = "localhost:3000"
	cfg.Scheme = "http"
	client := conekta.NewAPIClient(cfg)
	t.Run("GetBlacklist success", func(t *testing.T) {
		blacklist, response, err := client.BlacklistApi.GetBlacklist(context.TODO()).
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if blacklist == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", blacklist)
		}
	})
	t.Run("DeleteBlacklistRule success", func(t *testing.T) {
		_, response, err := client.BlacklistApi.DeleteBlacklistRule(context.TODO()).
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
	})
	t.Run("CreateNewBlacklistRule success", func(t *testing.T) {
		req := conekta.CreateRiskRulesData{
			Description: conekta.PtrString("desc"),
			Type:        conekta.PtrString("type"),
			Value:       conekta.PtrString("value"),
		}
		_, response, err := client.BlacklistApi.CreateNewBlacklistRule(context.TODO()).
			CreateRiskRulesData(req).
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
	})
}
