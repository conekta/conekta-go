package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func TestCustomersApi(t *testing.T) {
	cfg := conekta.NewConfiguration()
	cfg.Host = _basePath
	cfg.Scheme = "http"
	client := conekta.NewAPIClient(cfg)
	t.Run("GetCustomerById success", func(t *testing.T) {
		customer, response, err := client.CustomersApi.GetCustomerById(context.TODO(), "cus_2tYENskzTjjgkGQLt").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if customer == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", customer)
		}
	})
	t.Run("GetCustomers success", func(t *testing.T) {
		customers, response, err := client.CustomersApi.GetCustomers(context.TODO()).
			Limit(10).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if customers == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", customers)
		}
	})
	t.Run("CreateCustomer success", func(t *testing.T) {
		req := conekta.Customer{
			Email: "foo@foo.com",
			Name:  "Foo Foo",
		}
		customer, response, err := client.CustomersApi.CreateCustomer(context.TODO()).
			Customer(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if customer == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", customer)
		}
	})
	t.Run("DeleteCustomerById success", func(t *testing.T) {
		customer, response, err := client.CustomersApi.DeleteCustomerById(context.TODO(), "cus_2tYENskzTjjgkGQLt").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if customer == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", customer)
		}
	})
	t.Run("UpdateCustomer success", func(t *testing.T) {
		req := conekta.UpdateCustomer{
			Email: conekta.PtrString("foo@foo.com"),
			Name:  conekta.PtrString("Foo Foo"),
		}
		customer, response, err := client.CustomersApi.UpdateCustomer(context.TODO(), "cus_2tYENskzTjjgkGQLt").
			UpdateCustomer(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if customer == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", customer)
		}
	})
	t.Run("CreateCustomerFiscalEntities success", func(t *testing.T) {
		req := conekta.CustomerFiscalEntitiesRequest{
			Email:       conekta.PtrString("foo@foo.com"),
			CompanyName: conekta.PtrString("FooCompany"),
		}
		customer, response, err := client.CustomersApi.CreateCustomerFiscalEntities(context.TODO(), "cus_2tYENskzTjjgkGQLt").
			CustomerFiscalEntitiesRequest(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if customer == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", customer)
		}
	})
	t.Run("UpdateCustomerFiscalEntities success", func(t *testing.T) {
		req := conekta.CustomerUpdateFiscalEntitiesRequest{
			Email:       conekta.PtrString("foo@foo.com"),
			CompanyName: conekta.PtrString("FooCompany"),
		}
		customer, response, err := client.CustomersApi.UpdateCustomerFiscalEntities(context.TODO(), "cus_2tYENskzTjjgkGQLt", "fis_ent_2tYENskzTjjgkGQLr").
			CustomerUpdateFiscalEntitiesRequest(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if customer == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", customer)
		}
	})
}
