package token

import (
	conekta "github.com/conekta/conekta-go"
)

//Create create discount line insde a order
// For details see https://developers.conekta.com/api#create-discount-line
func Create(p *conekta.TokenParams) (*conekta.Token, *conekta.Error) {
	tk := &conekta.Token{}
	err := conekta.MakeRequest("POST", "/tokens", p, tk)
	return tk, err
}
