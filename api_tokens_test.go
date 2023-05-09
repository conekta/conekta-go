package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func TestTokensApi(t *testing.T) {
	cfg := conekta.NewConfiguration()
	cfg.Host = _basePath
	cfg.Scheme = "http"
	client := conekta.NewAPIClient(cfg)
	t.Run("CreateToken success", func(t *testing.T) {
		req := conekta.Token{
			Card: *conekta.NewNullableTokenCard(conekta.NewTokenCard("123", "02", "27", "Foo Foo", "5475040095304607")),
		}
		token, response, err := client.TokensApi.CreateToken(context.TODO()).
			Token(req).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if token == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", token)
		}
	})
}
