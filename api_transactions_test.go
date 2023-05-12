package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func TestTransactionsApi(t *testing.T) {
	cfg := conekta.NewConfiguration()
	cfg.Host = _basePath
	cfg.Scheme = "http"
	client := conekta.NewAPIClient(cfg)
	t.Run("GetTransactions success", func(t *testing.T) {
		transactions, response, err := client.TransactionsApi.GetTransactions(context.TODO()).
			AcceptLanguage("es").
			Limit(2).
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if transactions == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", transactions)
		}
	})
	t.Run("GetTransaction success", func(t *testing.T) {
		transaction, response, err := client.TransactionsApi.GetTransaction(context.TODO(), "6456b6dfac0fd40001a64eb8").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if transaction == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", transaction)
		}
	})
}
