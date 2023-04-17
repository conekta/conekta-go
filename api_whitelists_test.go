package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func TestWhitelistsApi(t *testing.T) {
	cfg := conekta.NewConfiguration()
	cfg.Host = "localhost:3000"
	cfg.Scheme = "http"
	client := conekta.NewAPIClient(cfg)
	t.Run("GetWhiteList success", func(t *testing.T) {
		whiteList, response, err := client.WhitelistsApi.GetWhiteList(context.TODO()).
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
	t.Run("CreateNewRuleWhitelist success", func(t *testing.T) {
		req := conekta.CreateRiskRulesData{
			Description: conekta.PtrString("Desc"),
			Type:        conekta.PtrString("type"),
			Value:       conekta.PtrString("rule test"),
		}
		response, err := client.WhitelistsApi.CreateNewRuleWhitelist(context.TODO()).
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
		response, err := client.WhitelistsApi.DeleteRuleWhitelist(context.TODO()).
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
	})
}
