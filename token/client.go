package token

import (
	conekta "github.com/conekta-go"
)

// New creates a new customer
// For details see https://developers.conekta.com/api#create-customer
func New(p *conekta.TokenParams) (*conekta.Token, *conekta.Error) {
	tk := &conekta.Token{}
	err := conekta.New(tk, p, "/tokens")
	return tk, err
}
