package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func TestAntifraudApi(t *testing.T) {
	cfg := conekta.NewConfiguration()
	cfg.Host = "localhost:3000"
	cfg.Scheme = "http"
	client := conekta.NewAPIClient(cfg)
	t.Run("GetRuleBlacklist success", func(t *testing.T) {
		blacklist, response, err := client.AntifraudApi.GetRuleBlacklist(context.TODO()).
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
	t.Run("DeleteRuleBlacklist success", func(t *testing.T) {
		_, response, err := client.AntifraudApi.DeleteRuleBlacklist(context.TODO(), "618c3f30db8b8da9be376b1e").
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
			Description: "desc",
			Field:       "field",
			Value:       "value",
		}
		_, response, err := client.AntifraudApi.CreateRuleWhitelist(context.TODO()).
			CreateRiskRulesData(req).
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
	})
	t.Run("GetRuleWhitelist success", func(t *testing.T) {
		whiteList, response, err := client.AntifraudApi.GetRuleWhitelist(context.TODO()).
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if whiteList == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", whiteList)
		}
	})
	t.Run("CreateRuleWhitelist success", func(t *testing.T) {
		req := conekta.CreateRiskRulesData{
			Description: "desc",
			Field:       "field",
			Value:       "value",
		}
		_, response, err := client.AntifraudApi.CreateRuleWhitelist(context.TODO()).
			CreateRiskRulesData(req).
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
	})
	t.Run("DeleteRuleWhitelist success", func(t *testing.T) {
		_, response, err := client.AntifraudApi.DeleteRuleWhitelist(context.TODO(), "618c3f2fdb8b8da9be376afe").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
	})
}
