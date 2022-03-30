package charge

import (
	"encoding/json"

	conekta "github.com/conekta/conekta-go/v2"
)

//Create charges object sending requesto api
//For more information please see https://developers.conekta.com/api#create-charge
func Create(orderID string, p *conekta.ChargeParams) (*conekta.Charge, error) {
	ch := &conekta.Charge{}
	err := conekta.MakeRequest("POST", "/orders/"+orderID+"/charges", p, ch)
	if err != nil && err.(conekta.Error).ErrorType == "processing_error" {
		json.Unmarshal(err.(conekta.Error).Data, ch)
		return ch, nil
	}
	return ch, err
}

//Find returns a charge based on his unique ID
func Find(orderID string, id string) (*conekta.Charge, error) {
	ch := &conekta.Charge{}
	err := conekta.MakeRequest("GET", "/orders/"+orderID+"/charges/"+id, &conekta.EmptyParams{}, ch)
	return ch, err
}
