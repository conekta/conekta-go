package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func TestPaymentLinkApi(t *testing.T) {
	cfg := conekta.NewConfiguration()
	cfg.Host = "localhost:3000"
	cfg.Scheme = "http"
	client := conekta.NewAPIClient(cfg)
	t.Run("GetCheckout success", func(t *testing.T) {
		checkout, response, err := client.PaymentLinkApi.GetCheckout(context.TODO(), "ff6918c6-5043-43b9-a7ec-d40d407d62c1").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if checkout == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", checkout)
		}
	})
	t.Run("GetCheckouts success", func(t *testing.T) {
		checkouts, response, err := client.PaymentLinkApi.GetCheckouts(context.TODO()).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if checkouts == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", checkouts)
		}
	})
	t.Run("CancelCheckout success", func(t *testing.T) {
		checkout, response, err := client.PaymentLinkApi.CancelCheckout(context.TODO(), "ff6918c6-5043-43b9-a7ec-d40d407d62c1").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if checkout == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", checkout)
		}
	})
	t.Run("CreateCheckout success", func(t *testing.T) {
		req := conekta.Checkout{
			AllowedPaymentMethods: []string{
				"method",
			},
			Name: "FooCheckout",
		}
		checkout, response, err := client.PaymentLinkApi.CreateCheckout(context.TODO()).
			Checkout(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if checkout == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", checkout)
		}
	})
	t.Run("EmailCheckout success", func(t *testing.T) {
		req := conekta.EmailCheckoutRequest{Email: "foo@foo.com"}
		checkout, response, err := client.PaymentLinkApi.EmailCheckout(context.TODO(), "ff6918c6-5043-43b9-a7ec-d40d407d62c1").
			EmailCheckoutRequest(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if checkout == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", checkout)
		}
	})
	t.Run("SmsCheckout success", func(t *testing.T) {
		req := conekta.SmsCheckoutRequest{Phonenumber: "12312312312"}
		checkout, response, err := client.PaymentLinkApi.SmsCheckout(context.TODO(), "ff6918c6-5043-43b9-a7ec-d40d407d62c1").
			SmsCheckoutRequest(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if checkout == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", checkout)
		}
	})
}
