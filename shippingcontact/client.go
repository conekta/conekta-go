package shippingcontact

import (
	conekta "github.com/conekta/conekta-go"
)

// Create creates a new shipping contact
// For details see https://developers.conekta.com/api#shipping-contact.
func Create(custID string, p *conekta.ShippingContactParams) (*conekta.ShippingContact, error) {
	sc := &conekta.ShippingContact{}
	err := conekta.MakeRequest("POST", "/customers/"+custID+"/shipping_contacts", p, sc)
	return sc, err
}

// Update updates a shipping contact
// For details see https://developers.conekta.com/api#update-shipping-contact
func Update(custID string, id string, p *conekta.ShippingContactParams) (*conekta.ShippingContact, error) {
	sc := &conekta.ShippingContact{}
	err := conekta.MakeRequest("PUT", "/customers/"+custID+"/shipping_contacts/"+id, p, sc)
	return sc, err
}

// Find gets a shipping contact by id
func Find(custID, id string) (*conekta.ShippingContact, error) {
	sc := &conekta.ShippingContact{}
	err := conekta.MakeRequest("GET", "/customers/"+custID+"/shipping_contacts/"+id, &conekta.EmptyParams{}, sc)
	return sc, err
}

// Delete deletes a shipping contact
func Delete(custID, id string) (*conekta.ShippingContact, error) {
	sc := &conekta.ShippingContact{}
	err := conekta.MakeRequest("DELETE", "/customers/"+custID+"/shipping_contacts/"+id, &conekta.EmptyParams{}, sc)
	return sc, err
}

// All gets all shipping contacts from a customer
func All(custID string) (*conekta.ShippingContactList, error) {
	scl := &conekta.ShippingContactList{}
	err := conekta.MakeRequest("GET", "/customers/"+custID+"/shipping_contacts/", &conekta.EmptyParams{}, scl)
	return scl, err
}
