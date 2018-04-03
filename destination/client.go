package destination

import (
	conekta "github.com/conekta/conekta-go"
)

// Create creates a new Payee destination
func Create(payeeID string, p *conekta.DestinationParams) (*conekta.Destination, error) {
	d := &conekta.Destination{}
	err := conekta.MakeRequest("POST", "/payees/"+payeeID+"/destinations", p, d)
	return d, err
}
