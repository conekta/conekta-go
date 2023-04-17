package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func TestDiscountsApi(t *testing.T) {
	cfg := conekta.NewConfiguration()
	cfg.Host = "localhost:3000"
	cfg.Scheme = "http"
	client := conekta.NewAPIClient(cfg)
	t.Run("OrdersCreateCharge success", func(t *testing.T) {
		req := conekta.OrderDiscountLinesRequest{
			Amount: 1000,
		}
		discount, response, err := client.DiscountsApi.OrdersCreateDiscountLine(context.TODO(), "ord_2tUigJ8DgBhbp6w5D").
			OrderDiscountLinesRequest(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if discount == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", discount)
		}
	})
	t.Run("OrdersDeleteDiscountLines success", func(t *testing.T) {
		discount, response, err := client.DiscountsApi.OrdersDeleteDiscountLines(context.TODO(), "ord_2tPAmKCEJqh8RE6nY", "dis_lin_2tQQ58HPgPw7StE8z").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if discount == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", discount)
		}
	})
	t.Run("OrdersDeleteDiscountLines success", func(t *testing.T) {
		req := conekta.UpdateOrderDiscountLinesRequest{
			Amount: conekta.PtrInt64(2000),
		}
		discount, response, err := client.DiscountsApi.OrdersUpdateDiscountLines(context.TODO(), "ord_2tPAmKCEJqh8RE6nY", "dis_lin_2tQQ58HPgPw7StE8z").
			UpdateOrderDiscountLinesRequest(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if discount == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", discount)
		}
	})
}
