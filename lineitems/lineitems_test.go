package lineitems

import (
	conekta "github.com/conekta/conekta-go"
	"github.com/conekta/conekta-go/order"
	"github.com/stretchr/testify/assert"
	"testing"
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
	lip := &conekta.LineItemsParams{}
	li, _ := Create(ord.ID, lip.Mock())

	assert.Equal(t, lip.Name, li.Name)
	assert.Equal(t, "line_item", li.Object)
	assert.Equal(t, lip.Quantity, li.Quantity)
	assert.Equal(t, lip.UnitPrice, li.UnitPrice)
}

func TestCreateError(t *testing.T) {
	ord, _ := CreateOrderWithoutCharges()
	_, err := Create(ord.ID, &conekta.LineItemsParams{})
	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "parameter_validation_error")
}

func TestUpdate(t *testing.T) {
	ord, _ := CreateOrderWithoutCharges()
	lip := &conekta.LineItemsParams{}
	li, _ := Create(ord.ID, lip.Mock())
	lip.Quantity = 2
	liu, err := Update(ord.ID, li.ID, lip)
	assert.Equal(t, 2, liu.Quantity)
	assert.Nil(t, err)
}

func TestUpdateError(t *testing.T) {
	ord, _ := CreateOrderWithoutCharges()
	lip := &conekta.LineItemsParams{}
	li, _ := Create(ord.ID, lip.Mock())
	lip.Quantity = -1
	_, err := Update(ord.ID, li.ID, lip)
	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "parameter_validation_error")
}

func TestDelete(t *testing.T) {
	ord, _ := CreateOrderWithoutCharges()
	lip := &conekta.LineItemsParams{}
	li, _ := Create(ord.ID, lip.Mock())
	li, err := Delete(ord.ID, li.ID)
	assert.Nil(t, err)
	assert.Equal(t, li.Deleted, true)

}

func TestDeleteError(t *testing.T) {
	ord, _ := CreateOrderWithoutCharges()
	_, err := Delete(ord.ID, "123")
	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "resource_not_found_error")
}
