package charge

import (
	"testing"

	"github.com/stretchr/testify/assert"

	conekta "github.com/conekta/conekta-go"
	"github.com/conekta/conekta-go/order"
)

func init() {
	conekta.APIKey = conekta.TestKey
}

func CreateOrder() (*conekta.Order, *conekta.Error) {
	op := &conekta.OrderParams{}
	ord, err := order.Create(op.Mock())
	return ord, err
}
func CreateOrderWithoutCharges() (*conekta.Order, *conekta.Error) {
	op := &conekta.OrderParams{}
	ord, err := order.Create(op.MockWithoutCharges())
	return ord, err
}

func TestCreate(t *testing.T) {
	cp := &conekta.ChargeParams{}
	or, _ := CreateOrderWithoutCharges()
	res, _ := Create(or.ID, cp.Mock())

	assert.Equal(t, false, res.Livemode)
	assert.NotEqual(t, nil, res.CreatedAt)
	assert.Equal(t, "MXN", res.Currency)
	assert.Equal(t, "charge", res.Object)
	assert.Equal(t, "paid", res.Status)
	assert.Equal(t, int64(1151), res.Amount)
	assert.Equal(t, int64(328), res.Fee)
	assert.NotEqual(t, nil, res.Description)
	assert.Equal(t, "", res.CustomerID)
	assert.NotEqual(t, nil, res.OrderID)
	assert.Equal(t, or.ID, res.OrderID)
}

func TestCreateError(t *testing.T) {
	ord, _ := CreateOrderWithoutCharges()
	_, err := Create(ord.ID, &conekta.ChargeParams{})
	assert.NotNil(t, err)
	assert.Equal(t, err.ErrorType, "parameter_validation_error")
}

func TestFind(t *testing.T) {
	ord, _ := CreateOrder()
	res, _ := Find(ord.ID, ord.Charges.Data[0].ID)

	assert.Equal(t, res.Description, "Payment from order")
	assert.Equal(t, int64(1151), res.Amount)
	assert.Equal(t, int64(328), res.Fee)
	assert.Equal(t, res.OrderID, ord.ID)
	assert.Equal(t, res.Object, "charge")
	assert.Equal(t, res.PaymentMethod.Last4, "4242")
	assert.Equal(t, res.PaymentMethod.Brand, "visa")
	assert.Equal(t, res.PaymentMethod.Name, "Jorge Lopez")
	assert.Equal(t, res.PaymentMethod.ExpMonth, "12")
	assert.Equal(t, res.PaymentMethod.ExpYear, "19")
}

func TestFindError(t *testing.T) {
	ord, _ := CreateOrderWithoutCharges()
	_, err := Find(ord.ID, "123")
	assert.NotNil(t, err)
	assert.Equal(t, err.ErrorType, "resource_not_found_error")
}
