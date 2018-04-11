package transfer

import (
	"testing"

	"github.com/conekta/conekta-go"
	"github.com/conekta/conekta-go/destination"
	"github.com/conekta/conekta-go/payee"
	"github.com/stretchr/testify/assert"
)

func init() {
	conekta.APIKey = conekta.TestKey
}

func createDestination() (*conekta.Payee, *conekta.Destination, error) {
	p := &conekta.PayeeParams{}
	payee, err := payee.Create(p.Mock())
	if err != nil {
		return nil, nil, err
	}
	dp := &conekta.DestinationParams{}
	des, err := destination.Create(payee.ID, dp.Mock())
	if err != nil {
		return nil, nil, err
	}
	return payee, des, nil
}

func createTransfer() (*conekta.Transfer, error) {
	p, _, err := createDestination()
	if err != nil {
		return nil, err
	}
	tp := &conekta.TransferParams{}
	tp.Mock()
	tp.PayeeID = p.ID
	return Create(tp)
}

func TestCreate(t *testing.T) {
	transfer, err := createTransfer()
	assert.NotZero(t, transfer.ID)
	assert.Nil(t, err)
}

func TestCreateError(t *testing.T) {
	_, err := Create(&conekta.TransferParams{})
	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "parameter_validation_error")
}

func TestFind(t *testing.T) {
	p, _, _ := createDestination()
	tp := &conekta.TransferParams{}
	tp.Mock()
	tp.PayeeID = p.ID
	transfer, _ := Create(tp)
	transfer, err := Find(transfer.ID)
	assert.Equal(t, tp.PayeeID, transfer.Destination.PayeeID)
	assert.Equal(t, tp.Amount, transfer.Amount)
	assert.Nil(t, err)
}

func TestFindError(t *testing.T) {
	_, err := Find("122")
	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "resource_not_found_error")
}
