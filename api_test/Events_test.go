package api

import (
	"conekta.io/api"
	"github.com/stretchr/testify/assert"

	"testing"


)


func TestChargePost(t *testing.T) {
	conekta := api.Conekta()
	conekta.SetApikey("key_eYvWV7gSDkNYXsmr");

	charge, _ := api.Events()

	assert.NotNil(t, charge, "should have charge")
	assert.NotEmpty(t, charge, "should have charge")
}