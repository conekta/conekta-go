package charge

import (
	conekta "github.com/conekta/conekta-go"
)

//Create charges object sending requesto api
//For more information please see https://developers.conekta.com/api#create-charge
func Create(id string, p *conekta.ChargeParams) (*conekta.Charge, *conekta.Error) {
	ch := &conekta.Charge{}
	err := conekta.MakeRequest("POST", "/orders/"+id+":/charges", p, ch)
	return ch, err
}

//Find returns a charge based on his unique ID
func Find(orderID string, id string, p *conekta.ChargeParams) (*conekta.Charge, *conekta.Error) {
	ch := &conekta.Charge{}
	err := conekta.MakeRequest("GET", "/orders/"+id+":/charges/"+id, p, ch)
	return ch, err
}
