package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func TestTransfersApi(t *testing.T) {
	cfg := conekta.NewConfiguration()
	cfg.Host = _basePath
	cfg.Scheme = "http"
	client := conekta.NewAPIClient(cfg)
	t.Run("GetTransfers success", func(t *testing.T) {
		transfers, response, err := client.TransfersApi.GetTransfers(context.TODO()).
			AcceptLanguage("es").
			Limit(5).
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if transfers == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", transfers)
		}
	})
	t.Run("GetTransfer success", func(t *testing.T) {
		transfer, response, err := client.TransfersApi.GetTransfer(context.TODO(), "64462930651b2600017b6d43").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if transfer == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", transfer)
		}
	})
}
