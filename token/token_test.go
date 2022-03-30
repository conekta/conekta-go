package token

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/conekta/conekta-go/v2"
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
	_, err := Create(&conekta.TokenParams{})
	if err, ok := err.(conekta.Error); ok {

		assert.NotNil(t, err)
		assert.Equal(t, err.ErrorType, "parameter_validation_error")
		assert.Equal(t, err.Details[0].Message, "The \"card\" parameter is missing but is a required attribute for token.")
		assert.Equal(t, err.Details[0].MessageToPurchaser, "El pago no pudo ser procesado.")
		assert.Equal(t, err.Details[0].Code, "")
		assert.Equal(t, err.Details[0].Param, "card")
	}
}
