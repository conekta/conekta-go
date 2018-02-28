package paymentsource

import (
	"testing"

	"github.com/conekta/conekta-go"
	"github.com/conekta/conekta-go/customer"

	"github.com/stretchr/testify/assert"
)

func init() {
	conekta.APIKey = conekta.TestKey
}

func createCustomer() *conekta.Customer {
	cp := &conekta.CustomerParams{}
	cust, _ := customer.Create(cp.Mock())
	return cust
}

func createPaymentSource() (*conekta.PaymentSource, *conekta.Customer, *conekta.Error) {
	cust := createCustomer()
	psp := &conekta.PaymentSourceCreateParams{}
	ps, err := Create(cust.ID, psp.Mock())
	return ps, cust, err
}

func TestCreate(t *testing.T) {
	cust := createCustomer()
	psp := &conekta.PaymentSourceCreateParams{}
	ps, err := Create(cust.ID, psp.Mock())
	assert.NotZero(t, ps.ID)
	assert.Nil(t, err)
}

func TestCreateError(t *testing.T) {
	cust := createCustomer()
	_, err := Create(cust.ID, &conekta.PaymentSourceCreateParams{})
	assert.NotNil(t, err)
	assert.Equal(t, err.ErrorType, "parameter_validation_error")
}

func TestUpdate(t *testing.T) {
	ps, cust, _ := createPaymentSource()
	psp := &conekta.PaymentSourceUpdateParams{}
	newn := "Another person"
	psp.Name = newn
	ps, err := Update(cust.ID, ps.ID, psp)
	assert.Equal(t, newn, ps.Name)
	assert.Nil(t, err)
}

func TestUpdateError(t *testing.T) {
	ps, cust, _ := createPaymentSource()
	psp := &conekta.PaymentSourceUpdateParams{}
	psp.ExpYear = "1"
	_, err := Update(cust.ID, ps.ID, psp)
	assert.NotNil(t, err)
	assert.Equal(t, err.ErrorType, "parameter_validation_error")
}

func TestFind(t *testing.T) {
	ps, cust, _ := createPaymentSource()
	ps, err := Find(cust.ID, ps.ID)
	assert.NotZero(t, ps.ID)
	assert.Nil(t, err)
}

func TestFindError(t *testing.T) {
	c := createCustomer()
	_, err := Find(c.ID, "123")
	assert.NotNil(t, err)
	assert.Equal(t, err.ErrorType, "resource_not_found_error")
}

func TestDelete(t *testing.T) {
	ps, cust, _ := createPaymentSource()
	ps, err := Delete(cust.ID, ps.ID)
	assert.Equal(t, true, ps.Deleted)
	assert.Nil(t, err)
}

func TestDeleteError(t *testing.T) {
	c := createCustomer()
	_, err := Delete(c.ID, "123")
	assert.NotNil(t, err)
	assert.Equal(t, err.ErrorType, "resource_not_found_error")
}

func TestAll(t *testing.T) {
	_, c, _ := createPaymentSource()
	psl, err := All(c.ID)
	assert.NotZero(t, psl.Total)
	assert.Nil(t, err)
}
