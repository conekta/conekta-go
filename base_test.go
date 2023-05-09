package conekta_test

import (
	"os"
	"testing"
)

var _basePath = "localhost:3000"

func TestMain(m *testing.M) {
	basePath := os.Getenv("BASE_PATH")
	if basePath != "" {
		_basePath = basePath
	}
	os.Exit(m.Run())
}
