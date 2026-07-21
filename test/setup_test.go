/*
Conekta API

Test setup helpers. These point the generated client at the Mockoon mock server
(https://raw.githubusercontent.com/conekta/openapi/main/mocks/conekta_api.json)
instead of the production API.

The mock server host is taken from the BASE_PATH environment variable
(as configured in the CI workflow), defaulting to localhost:3000 for local runs.
*/

package conekta

import (
	"os"

	conekta "github.com/conekta/conekta-go/v7"
)

// testBasePath returns the host:port of the mock server.
func testBasePath() string {
	basePath := os.Getenv("BASE_PATH")
	if basePath == "" {
		basePath = "localhost:3000"
	}
	return basePath
}

// testConfiguration builds a Configuration pointed at the mock server.
func testConfiguration() *conekta.Configuration {
	cfg := conekta.NewConfiguration()
	cfg.Servers = conekta.ServerConfigurations{
		{
			URL:         "http://" + testBasePath(),
			Description: "Conekta mock server",
		},
	}
	return cfg
}

// testClient returns an APIClient wired to the mock server.
func testClient() *conekta.APIClient {
	return conekta.NewAPIClient(testConfiguration())
}
