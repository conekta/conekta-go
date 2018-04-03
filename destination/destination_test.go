package destination

import (
	"testing"

	conekta "github.com/conekta/conekta-go"
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
	des, err := Create(payee.ID, dp.Mock())
	if err != nil {
		return nil, nil, err
	}
	return payee, des, nil
}

func TestCreate(t *testing.T) {
	_, des, err := createDestination()
	assert.NotZero(t, des.ID)
	assert.Nil(t, err)
}
