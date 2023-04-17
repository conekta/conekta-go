package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func TestProductsApi(t *testing.T) {
	cfg := conekta.NewConfiguration()
	cfg.Host = "localhost:3000"
	cfg.Scheme = "http"
	client := conekta.NewAPIClient(cfg)
	t.Run("OrdersCreateProduct success", func(t *testing.T) {
		req := conekta.Product{
			Name:     "Test Product",
			Quantity: 1,
		}
		product, response, err := client.ProductsApi.OrdersCreateProduct(context.TODO(), "ord_2tUigJ8DgBhbp6w5D").
			Product(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if product == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", product)
		}
	})
	t.Run("OrdersUpdateProduct success", func(t *testing.T) {
		req := conekta.UpdateProduct{
			Name:     conekta.PtrString("Test Product"),
			Quantity: conekta.PtrInt32(1),
		}
		product, response, err := client.ProductsApi.OrdersUpdateProduct(context.TODO(), "ord_2tUigJ8DgBhbp6w5D", "line_item_2tVz8UkyWhSxLfUd7").
			UpdateProduct(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if product == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", product)
		}
	})
	t.Run("OrdersDeleteProduct success", func(t *testing.T) {
		product, response, err := client.ProductsApi.OrdersDeleteProduct(context.TODO(), "ord_2tUigJ8DgBhbp6w5D", "line_item_2tVz8UkyWhSxLfUd7").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if product == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", product)
		}
	})
}
