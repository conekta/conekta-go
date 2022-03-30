package shippingcontact

import (
	"testing"

	"github.com/conekta/conekta-go/v2"
	"github.com/conekta/conekta-go/v2/customer"

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

func createCustomerSC() (*conekta.ShippingContact, *conekta.Customer, error) {
	cust := createCustomer()
	scp := &conekta.ShippingContactParams{}
	sc, err := Create(cust.ID, scp.Mock())
	return sc, cust, err
}

func TestCreateFromCustomer(t *testing.T) {
	cust := createCustomer()
	scp := &conekta.ShippingContactParams{}
	sc, err := Create(cust.ID, scp.Mock())
	assert.Equal(t, scp.Phone, sc.Phone)
	assert.Nil(t, err)
}

func TestCreateFromCustomerError(t *testing.T) {
	cust := createCustomer()
	_, err := Create(cust.ID, &conekta.ShippingContactParams{})
	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "parameter_validation_error")
}

func TestUpdateFromCustomer(t *testing.T) {
	sc, cust, _ := createCustomerSC()
	scp := &conekta.ShippingContactParams{}
	newr := "Another person"
	scp.Receiver = newr
	sc, err := Update(cust.ID, sc.ID, scp)
	assert.Equal(t, newr, sc.Receiver)
	assert.Nil(t, err)
}

func TestUpdateFromCustomerError(t *testing.T) {
	sc, cust, _ := createCustomerSC()
	scp := &conekta.ShippingContactParams{}
	scp.Phone = "12"
	_, err := Update(cust.ID, sc.ID, scp)
	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "parameter_validation_error")
}

func TestFindFromCustomer(t *testing.T) {
	sc, cust, _ := createCustomerSC()
	sc, err := Find(cust.ID, sc.ID)
	assert.NotZero(t, sc.ID)
	assert.Nil(t, err)
}

func TestFindFromCustomerError(t *testing.T) {
	c := createCustomer()
	_, err := Find(c.ID, "123")
	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "resource_not_found_error")
}

func TestDeleteFromCustomer(t *testing.T) {
	sc, cust, _ := createCustomerSC()
	sc, err := Delete(cust.ID, sc.ID)
	assert.Equal(t, true, sc.Deleted)
	assert.Nil(t, err)
}

func TestDeleteFromCustomerError(t *testing.T) {
	c := createCustomer()
	_, err := Delete(c.ID, "123")
	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "resource_not_found_error")
}

func TestAllFromCustomer(t *testing.T) {
	_, c, _ := createCustomerSC()
	scl, err := All(c.ID)
	assert.NotZero(t, len(scl.Data))
	assert.Nil(t, err)
}
