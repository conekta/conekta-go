package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func TestOrdersApi(t *testing.T) {
	cfg := conekta.NewConfiguration()
	cfg.Host = _basePath
	cfg.Scheme = "http"
	client := conekta.NewAPIClient(cfg)
	t.Run("GetOrderById success", func(t *testing.T) {
		orderID := "ord_2tUigJ8DgBhbp6w5D"
		order, response, err := client.OrdersApi.GetOrderById(context.TODO(), orderID).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if order == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", order)
		}
	})
	t.Run("GetOrders success", func(t *testing.T) {
		orders, response, err := client.OrdersApi.GetOrders(context.TODO()).Limit(10).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if len(orders.Data) < 0 {
			t.Errorf("assertion fail, expected=%v , actual=%v", ">0", len(orders.Data))
		}
	})
	t.Run("CreateOrder success", func(t *testing.T) {
		customerID := "cus_2tYENskzTjjgkGQLt"
		req := conekta.OrderRequest{
			Currency: "MXN",
			CustomerInfo: conekta.OrderRequestCustomerInfo{
				CustomerInfoJustCustomerId: conekta.NewCustomerInfoJustCustomerId(customerID),
			},
			LineItems: []conekta.Product{
				{
					Quantity:  1,
					UnitPrice: 1000,
				},
			},
		}
		order, response, err := client.OrdersApi.CreateOrder(context.TODO()).OrderRequest(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if order == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", order)
		}
	})
	t.Run("OrderRefund success", func(t *testing.T) {
		req := conekta.OrderRefundRequest{
			Amount: 1000,
			Reason: "reason",
		}
		orderID := "ord_2tUigJ8DgBhbp6w5D"
		order, response, err := client.OrdersApi.OrderRefund(context.TODO(), orderID).
			OrderRefundRequest(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if order == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", order)
		}
	})
	t.Run("OrdersCreateCapture success", func(t *testing.T) {
		orderID := "ord_2tUigJ8DgBhbp6w5D"
		order, response, err := client.OrdersApi.OrdersCreateCapture(context.TODO(), orderID).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if order == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", order)
		}
	})
	t.Run("UpdateOrder success", func(t *testing.T) {
		req := conekta.OrderUpdateRequest{
			Currency: conekta.PtrString("MXN"),
			CustomerInfo: &conekta.OrderRequestCustomerInfo{
				CustomerInfoJustCustomerId: conekta.NewCustomerInfoJustCustomerId("cus_2tYENskzTjjgkGQLt"),
			},
			LineItems: []conekta.Product{
				{
					Quantity:  1,
					UnitPrice: 1200,
				},
			},
		}
		orderID := "ord_2tUigJ8DgBhbp6w5D"
		order, response, err := client.OrdersApi.UpdateOrder(context.TODO(), orderID).
			OrderUpdateRequest(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if order == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", order)
		}
	})
}
