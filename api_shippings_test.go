package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func TestShippingsApi(t *testing.T) {
	cfg := conekta.NewConfiguration()
	cfg.Host = "localhost:3000"
	cfg.Scheme = "http"
	client := conekta.NewAPIClient(cfg)
	t.Run("OrdersCreateShipping success", func(t *testing.T) {
		req := conekta.ShippingRequest{
			Amount:         100,
			Carrier:        conekta.PtrString("FedEx"),
			TrackingNumber: conekta.PtrString("ASDA12312"),
		}
		shipping, response, err := client.ShippingsApi.OrdersCreateShipping(context.TODO(), "ord_2tUigJ8DgBhbp6w5D").
			ShippingRequest(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if shipping == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", shipping)
		}
	})
	t.Run("OrdersUpdateShipping success", func(t *testing.T) {
		req := conekta.ShippingRequest{
			Amount:         100,
			Carrier:        conekta.PtrString("FedEx"),
			TrackingNumber: conekta.PtrString("ASDA12312"),
		}
		shipping, response, err := client.ShippingsApi.OrdersUpdateShipping(context.TODO(), "ord_2tUigJ8DgBhbp6w5D", "ship_lin_2tVzNuDGSaDwreMg6").
			ShippingRequest(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if shipping == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", shipping)
		}
	})
	t.Run("OrdersDeleteShipping success", func(t *testing.T) {
		shipping, response, err := client.ShippingsApi.OrdersDeleteShipping(context.TODO(), "ord_2tUigJ8DgBhbp6w5D", "ship_lin_2tVzNuDGSaDwreMg6").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if shipping == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", shipping)
		}
	})
}
