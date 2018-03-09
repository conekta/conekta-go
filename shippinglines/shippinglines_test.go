package shippinglines

import (
	"testing"

	"github.com/conekta/conekta-go"
	"github.com/conekta/conekta-go/order"

	"github.com/stretchr/testify/assert"
)

func init() {
	conekta.APIKey = conekta.TestKey
}
func CreateOrderWithoutCharges() (*conekta.Order, error) {
	op := &conekta.OrderParams{}
	ord, err := order.Create(op.MockWithoutCharges())
	return ord, err
}

func TestCreate(t *testing.T) {
	ord, _ := CreateOrderWithoutCharges()
	slp := &conekta.ShippingLinesParams{}
	sl, err := Create(ord.ID, slp.Mock())

	assert.Equal(t, slp.Amount, sl.Amount)
	assert.Equal(t, slp.Method, sl.Method)
	assert.Equal(t, slp.TrackingNumber, sl.TrackingNumber)
	assert.Equal(t, slp.Carrier, sl.Carrier)
	assert.Nil(t, err)
}

func TestCreateError(t *testing.T) {
	ord, _ := CreateOrderWithoutCharges()
	slp := &conekta.ShippingLinesParams{}
	slp.Amount = -1
	_, err := Create(ord.ID, slp)

	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "parameter_validation_error")
}

func TestUpdate(t *testing.T) {
	ord, _ := CreateOrderWithoutCharges()
	slp := &conekta.ShippingLinesParams{}
	assert.Equal(t, ord.ShippingLines.Data[0].Method, "Fedex")
	slp.Method = "otro"
	slpu, _ := Update(ord.ID, ord.ShippingLines.Data[0].ID, slp)
	assert.Equal(t, slpu.Method, "otro")

}

func TestUpdateError(t *testing.T) {
	ord, _ := CreateOrderWithoutCharges()
	slp := &conekta.ShippingLinesParams{}
	slp.Amount = -1
	_, err := Update(ord.ID, ord.ShippingLines.Data[0].ID, slp)

	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "parameter_validation_error")

}

func TestDelete(t *testing.T) {
	ord, _ := CreateOrderWithoutCharges()
	sld, _ := Delete(ord.ID, ord.ShippingLines.Data[0].ID)
	assert.Equal(t, true, sld.Deleted)
}
func TestDeleteError(t *testing.T) {
	ord, _ := CreateOrderWithoutCharges()
	_, err := Delete(ord.ID, "123")
	assert.Equal(t, err.(conekta.Error).ErrorType, "resource_not_found_error")
}
