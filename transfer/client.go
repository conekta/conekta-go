package transfer

import (
	conekta "github.com/conekta/conekta-go"
)

// Create creates a new Transfer
func Create(p *conekta.TransferParams) (*conekta.Transfer, error) {
	t := &conekta.Transfer{}
	err := conekta.MakeRequest("POST", "/transfers", p, t)
	return t, err
}

// Find gets a Transfer by id
func Find(id string) (*conekta.Transfer, error) {
	t := &conekta.Transfer{}
	err := conekta.MakeRequest("GET", "/transfers/"+id, &conekta.EmptyParams{}, t)
	return t, err
}
