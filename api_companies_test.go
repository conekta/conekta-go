package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func TestCompaniesApi(t *testing.T) {
	cfg := conekta.NewConfiguration()
	cfg.Host = _basePath
	cfg.Scheme = "http"
	client := conekta.NewAPIClient(cfg)
	t.Run("GetCompany success", func(t *testing.T) {
		company, response, err := client.CompaniesApi.GetCompany(context.TODO(), "5c537c4d27839876e2f18139").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if company == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", company)
		}
	})
	t.Run("GetCompanies success", func(t *testing.T) {
		companies, response, err := client.CompaniesApi.GetCompanies(context.TODO()).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if companies == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", companies)
		}
	})
}
