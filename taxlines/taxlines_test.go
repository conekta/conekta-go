package taxlines

import (
	"testing"

	"github.com/conekta/conekta-go/v2"
	"github.com/conekta/conekta-go/v2/order"

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
	tlp := &conekta.TaxLinesParams{}
	tl, err := Create(ord.ID, tlp.Mock())

	assert.Equal(t, tl.Amount, int64(160))
	assert.Equal(t, tl.Description, "IVA")
	assert.Nil(t, err)
}

func TestCreateError(t *testing.T) {
	ord, _ := CreateOrderWithoutCharges()
	_, err := Create(ord.ID, &conekta.TaxLinesParams{})

	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "parameter_validation_error")
}

func TestUpdate(t *testing.T) {
	ord, _ := CreateOrderWithoutCharges()
	tlp := &conekta.TaxLinesParams{}
	assert.Equal(t, ord.TaxLines.Data[0].Description, "IVA")
	tlp.Description = "otro"
	tl, err := Update(ord.ID, ord.TaxLines.Data[0].ID, tlp)

	assert.Equal(t, tl.Description, "otro")
	assert.Nil(t, err)
}

func TestUpdateError(t *testing.T) {
	ord, _ := CreateOrderWithoutCharges()
	_, err := Create(ord.ID, &conekta.TaxLinesParams{})

	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "parameter_validation_error")
}

func TestDelete(t *testing.T) {
	ord, _ := CreateOrderWithoutCharges()
	tl, err := Delete(ord.ID, ord.TaxLines.Data[0].ID)

	assert.Equal(t, tl.Deleted, true)
	assert.Nil(t, err)
}

func TestDeleteError(t *testing.T) {
	ord, _ := CreateOrderWithoutCharges()
	_, err := Delete(ord.ID, "123")

	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "resource_not_found_error")
}
