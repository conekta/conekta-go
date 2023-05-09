package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func TestChargesApi(t *testing.T) {
	cfg := conekta.NewConfiguration()
	cfg.Host = _basePath
	cfg.Scheme = "http"
	client := conekta.NewAPIClient(cfg)
	t.Run("OrdersCreateCharge success", func(t *testing.T) {
		req := conekta.ChargeRequest{
			Amount: conekta.PtrInt32(1000),
		}
		chargeOrder, response, err := client.ChargesApi.OrdersCreateCharge(context.TODO(), "ord_2tUigJ8DgBhbp6w5D").
			ChargeRequest(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if chargeOrder == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", chargeOrder)
		}
	})
}
