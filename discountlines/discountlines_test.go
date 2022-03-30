package discountlines

import (
	"testing"

	"github.com/stretchr/testify/assert"

	conekta "github.com/conekta/conekta-go/v2"
	"github.com/conekta/conekta-go/v2/order"
)

func init() {
	conekta.APIKey = conekta.TestKey
}

func CreateOrderWithoutCharges() (*conekta.Order, error) {
	op := &conekta.OrderParams{}
	ord, err := order.Create(op.MockWithoutCharges())
	return ord, err
}
func CreateOrderWithoutDiscountLines() (*conekta.Order, error) {
	op := &conekta.OrderParams{}
	ord, err := order.Create(op.MockWithoutDiscountLines())
	return ord, err
}

func TestCreate(t *testing.T) {
	ord, _ := CreateOrderWithoutDiscountLines()
	dlp := &conekta.DiscountLinesParams{}

	assert.Nil(t, ord.DiscountLines)

	dl, _ := Create(ord.ID, dlp.Mock())

	assert.Equal(t, ord.ID, dl.ParentID)
	assert.Equal(t, dlp.Code, dl.Code)
	assert.Equal(t, dlp.Amount, dl.Amount)
	assert.Equal(t, "discount_line", dl.Object)
	assert.NotNil(t, dl.ID)
}

func TestCreateError(t *testing.T) {
	ord, _ := CreateOrderWithoutDiscountLines()
	_, err := Create(ord.ID, &conekta.DiscountLinesParams{})

	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "parameter_validation_error")
}

func TestDelete(t *testing.T) {
	ord, _ := CreateOrderWithoutCharges()
	res, _ := Delete(ord.ID, ord.DiscountLines.Data[0].ID)
	assert.Equal(t, true, res.Deleted)

}
func TestDeleteError(t *testing.T) {
	ord, _ := CreateOrderWithoutCharges()
	_, err := Delete(ord.ID, "123")
	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "resource_not_found_error")
}

func TestUpdate(t *testing.T) {
	dlp := &conekta.DiscountLinesParams{}
	ord, _ := CreateOrderWithoutCharges()

	assert.Equal(t, "Cupon de descuento", ord.DiscountLines.Data[0].Code)

	dlp.Code = "Cupon nuevo"
	dl, _ := Update(ord.ID, ord.DiscountLines.Data[0].ID, dlp)
	assert.Equal(t, "Cupon nuevo", dl.Code)
}

func TestUpdateError(t *testing.T) {
	ord, _ := CreateOrderWithoutCharges()
	dlp := &conekta.DiscountLinesParams{}
	dlp.Amount = -1
	_, err := Update(ord.ID, ord.DiscountLines.Data[0].ID, dlp)
	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "parameter_validation_error")
}
