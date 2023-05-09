package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func TestShippingContactsApi(t *testing.T) {
	cfg := conekta.NewConfiguration()
	cfg.Host = _basePath
	cfg.Scheme = "http"
	client := conekta.NewAPIClient(cfg)
	t.Run("CreateCustomerShippingContacts success", func(t *testing.T) {
		req := conekta.CustomerShippingContacts{
			Phone: conekta.PtrString("1132312312"),
		}
		shippingContact, response, err := client.ShippingContactsApi.CreateCustomerShippingContacts(context.TODO(), "cus_2tYENskzTjjgkGQLt").
			CustomerShippingContacts(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if shippingContact == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", shippingContact)
		}
	})
	t.Run("UpdateCustomerShippingContacts success", func(t *testing.T) {
		req := conekta.CustomerUpdateShippingContacts{
			Phone: conekta.PtrString("1132312312"),
		}
		shippingContact, response, err := client.ShippingContactsApi.UpdateCustomerShippingContacts(context.TODO(), "cus_2tYENskzTjjgkGQLt", "ship_cont_2tY2ncFSBLUaohto2").
			CustomerUpdateShippingContacts(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if shippingContact == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", shippingContact)
		}
	})
	t.Run("DeleteCustomerShippingContacts success", func(t *testing.T) {
		shippingContact, response, err := client.ShippingContactsApi.DeleteCustomerShippingContacts(context.TODO(), "cus_2tYENskzTjjgkGQLt", "ship_cont_2tY2ncFSBLUaohto2").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if shippingContact == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", shippingContact)
		}
	})
}
