package conekta_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/conekta/conekta-go"
)

func TestLogsApi(t *testing.T) {
	cfg := conekta.NewConfiguration()
	cfg.Host = _basePath
	cfg.Scheme = "http"
	client := conekta.NewAPIClient(cfg)
	t.Run("GetLogById success", func(t *testing.T) {
		log, response, err := client.LogsApi.GetLogById(context.TODO(), "6419dd15b985080001fc280e").
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if log == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", log)
		}
	})
	t.Run("GetLogs success", func(t *testing.T) {
		logs, response, err := client.LogsApi.GetLogs(context.TODO()).
			AcceptLanguage("es").
			Execute()
		if err != nil {
			t.Error(err)
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("assertion fail, expected=%v , actual=%v", http.StatusOK, response.StatusCode)
		}
		if logs == nil {
			t.Errorf("assertion fail, expected=%v , actual=%v", "not nil", logs)
		}
	})
}
