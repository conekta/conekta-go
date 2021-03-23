package checkout

import (
	"testing"

	conekta "github.com/conekta/conekta-go"
	"github.com/stretchr/testify/assert"
)

func init() {
	conekta.APIKey = conekta.TestKey
}

type Query struct {
	Currency string `url:"currency,omitempty"`
}

func TestCreate(t *testing.T) {
	op := (&conekta.CheckoutParams{}).Mock()
	checkoutResponse, err := Create(op)

	assert.Equal(t, op.AllowedPaymentMethods, checkoutResponse.AllowedPaymentMethods)
	assert.Equal(t, op.ExpiredAt, checkoutResponse.ExpiredAt)
	assert.Equal(t, op.MonthlyInstallmentsEnabled, checkoutResponse.MonthlyInstallmentsEnabled)
	assert.Equal(t, []int64(nil), checkoutResponse.AllowedInstallmentOptions)
	assert.Equal(t, op.Name, checkoutResponse.Name)
	assert.Equal(t, op.NeedsShippingContact, checkoutResponse.NeedsShippingContact)
	assert.Equal(t, op.Recurrent, checkoutResponse.Recurrent)
	assert.Equal(t, op.Type, checkoutResponse.Type)
	assert.Equal(t, int64(0), checkoutResponse.EmailsSent)
	assert.Equal(t, int64(0), checkoutResponse.SmsSent)
	assert.Nil(t, err)
}

func TestCreateValidationError(t *testing.T) {
	op := (&conekta.CheckoutParams{}).Mock()
	op.ExpiredAt = 0
	_, err := Create(op)
	assert.NotNil(t, err)
	assert.NotNil(t, err.(conekta.Error).ErrorType, "parameter_validation_error")
}

func TestWhere(t *testing.T) {
	op := &Query{}

	res, err := Where(op)

	assert.Nil(t, err)
	assert.NotZero(t, len(res.Data))
	assert.True(t, res.HasMore)
	assert.Equal(t, "list", res.Object)

}

func TestFind(t *testing.T) {
	op := (&conekta.CheckoutParams{}).Mock()
	checkout, err := Create(op)
	res, err := Find(checkout.ID)
	assert.Equal(t, checkout.ID, res.ID)
	assert.Nil(t, err)
}

func TestCancel(t *testing.T) {
	op := (&conekta.CheckoutParams{}).Mock()
	checkout, _ := Create(op)
	res, err := Cancel(checkout.ID)

	assert.NotNil(t, res.ID)
	assert.Equal(t, res.ID, checkout.ID)
	assert.Equal(t, res.Status, "Cancelled")
	assert.Nil(t, err)

}

func TestSendEmail(t *testing.T) {
	op := (&conekta.CheckoutParams{}).Mock()
	checkout, _ := Create(op)
	res, err := SendEmail(checkout.ID, "mail@mail.com")

	assert.Equal(t, res.ID, checkout.ID)
	assert.Equal(t, res.EmailsSent, int64(1))
	assert.Nil(t, err)
}
func TestSendSms(t *testing.T) {
	op := (&conekta.CheckoutParams{}).Mock()

	checkout, _ := Create(op)
	res, err := SendSms(checkout.ID, "mail@mail.com")

	assert.Equal(t, res.ID, checkout.ID)
	assert.Equal(t, res.SmsSent, int64(1))
	assert.Nil(t, err)
}
