package token

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/conekta/conekta-go"
)

func init() {
	conekta.PubliAPIKey = conekta.TestPublic
}

func TestCreate(t *testing.T) {
	tp := &conekta.TokenParams{}
	tok, _ := Create(tp.Mock())
	assert.NotNil(t, tok.ID)
	assert.Equal(t, tok.Livemode, false)
	assert.Equal(t, tok.Object, "token")
	assert.Equal(t, tok.Used, false)
}

func TestCreateError(t *testing.T) {
	tok, _ := Create(&conekta.TokenParams{})

	assert.NotNil(t, tok)
	assert.Equal(t, tok.Type, "parameter_validation_error")
	assert.Equal(t, tok.Message, "The \"card\" parameter is missing but is a required attribute for token.")
	assert.Equal(t, tok.MessageToPurchaser, "El pago no pudo ser procesado.")
	assert.Equal(t, tok.ErrorCode, "")
	assert.Equal(t, tok.Param, "card")
}
