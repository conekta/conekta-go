package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func TestPaymentMethodsApi(t *testing.T) {
	cfg := conekta.NewConfiguration()
	cfg.Host = _basePath
	cfg.Scheme = "http"
	client := conekta.NewAPIClient(cfg)
	t.Run("GetCustomerPaymentMethods success", func(t *testing.T) {
		paymentMethods, response, err := client.PaymentMethodsApi.GetCustomerPaymentMethods(context.TODO(), "cus_2tYENskzTjjgkGQLt").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if paymentMethods == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", paymentMethods)
		}
	})
	t.Run("DeleteCustomerPaymentMethods success", func(t *testing.T) {
		paymentMethod, response, err := client.PaymentMethodsApi.DeleteCustomerPaymentMethods(context.TODO(), "cus_2tYENskzTjjgkGQLt", "src_2tbd5Bgy67RL9oycM").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if paymentMethod == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", paymentMethod)
		}
	})
	t.Run("CreateCustomerPaymentMethods success", func(t *testing.T) {
		req := conekta.CreateCustomerPaymentMethodsRequest{
			PaymentMethodCardRequest: &conekta.PaymentMethodCardRequest{
				Type:    "type",
				TokenId: "tokenID",
			},
		}
		paymentMethod, response, err := client.PaymentMethodsApi.CreateCustomerPaymentMethods(context.TODO(), "cus_2tYENskzTjjgkGQLt").
			CreateCustomerPaymentMethodsRequest(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if paymentMethod == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", paymentMethod)
		}
	})
	t.Run("UpdateCustomerPaymentMethods success", func(t *testing.T) {
		req := conekta.UpdatePaymentMethods{
			Name: conekta.PtrString("name"),
		}
		paymentMethod, response, err := client.PaymentMethodsApi.UpdateCustomerPaymentMethods(context.TODO(), "cus_2tYENskzTjjgkGQLt", "src_2tbd5Bgy67RL9oycM").
			UpdatePaymentMethods(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if paymentMethod == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", paymentMethod)
		}
	})
}
