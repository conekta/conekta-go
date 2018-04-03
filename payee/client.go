package payee

import (
	conekta "github.com/conekta/conekta-go"
)

// Create creates a new Payee
func Create(p *conekta.PayeeParams) (*conekta.Payee, error) {
	ps := &conekta.Payee{}
	err := conekta.MakeRequest("POST", "/payees", p, ps)
	return ps, err
}

// Update updates a Payee
func Update(id string, p *conekta.PayeeParams) (*conekta.Payee, error) {
	payee := &conekta.Payee{}
	err := conekta.MakeRequest("PUT", "/payees/"+id, p, payee)
	return payee, err
}

// Find gets a Payee by id
func Find(id string) (*conekta.Payee, error) {
	payee := &conekta.Payee{}
	err := conekta.MakeRequest("GET", "/payees/"+id, &conekta.EmptyParams{}, payee)
	return payee, err
}

// Delete deletes a Payee
func Delete(id string) (*conekta.Payee, error) {
	payee := &conekta.Payee{}
	err := conekta.MakeRequest("DELETE", "/payees/"+id, &conekta.EmptyParams{}, payee)
	return payee, err
}

// All gets all Payee
func All() (*conekta.PayeeList, error) {
	pl := &conekta.PayeeList{}
	err := conekta.MakeRequest("GET", "/payees", &conekta.EmptyParams{}, pl)
	return pl, err
}
