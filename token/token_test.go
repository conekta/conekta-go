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
	if _, err := Create(&conekta.TokenParams{}); err != nil {
		conektaErr := err.(*conekta.Error)
		assert.NotNil(t, conektaErr)
		assert.Equal(t, conektaErr.ErrorType, "parameter_validation_error")
		assert.Equal(t, conektaErr.Details[0].Message, "The \"card\" parameter is missing but is a required attribute for token.")
		assert.Equal(t, conektaErr.Details[0].MessageToPurchaser, "El pago no pudo ser procesado.")
		assert.Equal(t, conektaErr.Details[0].Code, "")
		assert.Equal(t, conektaErr.Details[0].Param, "card")
	}

}
