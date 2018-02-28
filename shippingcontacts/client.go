package shippingcontact

import (
	conekta "github.com/conekta/conekta-go"
)

// Create creates a new shipping contact
// For details see https://developers.conekta.com/api#shipping-contact.
func Create(custID string, p *conekta.ShippingContactParams) (*conekta.ShippingContact, *conekta.Error) {
	sc := &conekta.ShippingContact{}
	err := conekta.MakeRequest("POST", "/customers/"+custID+"/shipping_contacts", p, sc)
	return sc, err
}
