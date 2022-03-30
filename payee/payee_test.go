package payee

import (
	"testing"

	"github.com/conekta/conekta-go/v2"
	"github.com/stretchr/testify/assert"
)

func init() {
	conekta.APIKey = conekta.TestKey
}

func createPayee() (*conekta.Payee, error) {
	p := &conekta.PayeeParams{}
	return Create(p.Mock())
}

func TestCreate(t *testing.T) {
	payee, err := createPayee()
	assert.NotZero(t, payee.ID)
	assert.Nil(t, err)
}

func TestCreateError(t *testing.T) {
	_, err := Create(&conekta.PayeeParams{})
	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "parameter_validation_error")
}

func TestFind(t *testing.T) {
	pp := &conekta.PayeeParams{}
	payee, _ := Create(pp.Mock())
	payee, err := Find(payee.ID)
	assert.Equal(t, pp.Email, payee.Email)
	assert.Equal(t, pp.Name, payee.Name)
	assert.Nil(t, err)
}

func TestFindError(t *testing.T) {
	_, err := Find("122")
	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "resource_not_found_error")
}

func TestDelete(t *testing.T) {
	payee, _ := createPayee()
	payee, err := Delete(payee.ID)
	assert.Equal(t, payee.Deleted, true)
	assert.Nil(t, err)
}

func TestDeleteError(t *testing.T) {
	_, err := Delete("122")
	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "resource_not_found_error")
}

func TestAll(t *testing.T) {
	createPayee()
	pl, err := All()
	assert.NotZero(t, len(pl.Data))
	assert.Nil(t, err)
}
