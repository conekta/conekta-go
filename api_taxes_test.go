package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func TestTaxesApi(t *testing.T) {
	cfg := conekta.NewConfiguration()
	cfg.Host = _basePath
	cfg.Scheme = "http"
	client := conekta.NewAPIClient(cfg)
	t.Run("OrdersCreateTaxes success", func(t *testing.T) {
		req := conekta.OrderTaxRequest{
			Amount:      100,
			Description: "Test",
		}
		tax, response, err := client.TaxesApi.OrdersCreateTaxes(context.TODO(), "ord_2tUigJ8DgBhbp6w5D").
			OrderTaxRequest(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if tax == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", tax)
		}
	})
	t.Run("OrdersUpdateTaxes success", func(t *testing.T) {
		req := conekta.UpdateOrderTaxRequest{
			Amount:      conekta.PtrInt64(100),
			Description: conekta.PtrString("Test"),
		}
		tax, response, err := client.TaxesApi.OrdersUpdateTaxes(context.TODO(), "ord_2tUigJ8DgBhbp6w5D", "tax_lin_2tVzVp6AAptCRHhgt").
			UpdateOrderTaxRequest(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if tax == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", tax)
		}
	})
	t.Run("OrdersDeleteTaxes success", func(t *testing.T) {
		tax, response, err := client.TaxesApi.OrdersDeleteTaxes(context.TODO(), "ord_2tVyWPnCPWbrV37mW", "tax_lin_2tVzVp6AAptCRHhgt").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if tax == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", tax)
		}
	})
}
