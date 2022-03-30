package customer

import (
	"testing"

	"github.com/conekta/conekta-go/v2"

	"github.com/stretchr/testify/assert"
)

func init() {
	conekta.APIKey = conekta.TestKey
}

func TestCreate(t *testing.T) {
	cp := &conekta.CustomerParams{}
	cust, err := Create(cp.Mock())
	assert.Equal(t, cp.Name, cust.Name)
	assert.Equal(t, cp.Phone, cust.Phone)
	assert.Equal(t, cp.Corporate, cust.Corporate)
	assert.NotZero(t, len(cust.ShippingContacts.Data))
	assert.Nil(t, err)
}

func TestCreateWithCardAsPaymentSource(t *testing.T) {
	cp := &conekta.CustomerParams{}
	cp.Mock()
	cardParams := &conekta.CardParams{}
	ps := &conekta.PaymentSourceCreateParams{
		PaymentType: "card",
		CardParams:  cardParams.Mock(),
	}
	cp.PaymentSources = append(cp.PaymentSources, ps)

	cust, err := Create(cp)
	assert.Equal(t, cp.Name, cust.Name)
	assert.Equal(t, cp.Phone, cust.Phone)
	assert.Equal(t, cp.Corporate, cust.Corporate)
	assert.NotZero(t, len(cust.ShippingContacts.Data))
	assert.NotZero(t, len(cust.PaymentSources.Data))
	assert.NotNil(t, cust.DefaultPaymentSourceID)
	assert.Nil(t, err)
}

func TestCreateError(t *testing.T) {
	_, err := Create(&conekta.CustomerParams{})
	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "parameter_validation_error")
}

func TestUpdate(t *testing.T) {
	cp := &conekta.CustomerParams{}
	cust, _ := Create(cp.Mock())
	newName := "A great new name"
	cp.Name = newName
	cust, err := Update(cust.ID, cp)
	assert.Equal(t, newName, cust.Name)
	assert.Nil(t, err)
}

func TestUpdateError(t *testing.T) {
	cp := &conekta.CustomerParams{}
	cust, _ := Create(cp.Mock())
	email := "a"
	cp.Email = email
	_, err := Update(cust.ID, cp)
	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "parameter_validation_error")
}

func TestFind(t *testing.T) {
	cp := &conekta.CustomerParams{}
	cust, _ := Create(cp.Mock())
	cust, err := Find(cust.ID)
	assert.Equal(t, cp.Email, cust.Email)
	assert.Equal(t, cp.Name, cust.Name)
	assert.Nil(t, err)
}

func TestFindError(t *testing.T) {
	_, err := Find("122")
	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "resource_not_found_error")
}

func TestDelete(t *testing.T) {
	cp := &conekta.CustomerParams{}
	cust, _ := Create(cp.Mock())
	cust, err := Delete(cust.ID)
	assert.Equal(t, cust.Deleted, true)
	assert.Nil(t, err)
}

func TestDeleteError(t *testing.T) {
	_, err := Delete("122")
	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "resource_not_found_error")
}

func TestAll(t *testing.T) {
	cp := &conekta.CustomerParams{}
	Create(cp.Mock())
	cl, err := All()
	assert.NotZero(t, len(cl.Data))
	assert.Nil(t, err)
}
