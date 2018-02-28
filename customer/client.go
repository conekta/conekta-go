package customer

import (
	conekta "github.com/conekta/conekta-go"
)

// Create creates a new customer
// For details see https://developers.conekta.com/api#create-customer
func Create(p *conekta.CustomerParams) (*conekta.Customer, *conekta.Error) {
	cust := &conekta.Customer{}
	err := conekta.MakeRequest("POST", "/customers", p, cust)
	return cust, err
}

// Update updates a customer
// For details see https://developers.conekta.com/api#update-customer
func Update(id string, p *conekta.CustomerParams) (*conekta.Customer, *conekta.Error) {
	cust := &conekta.Customer{}
	err := conekta.MakeRequest("PUT", "/customers/"+id, p, cust)
	return cust, err
}

// Find gets a customer by id
func Find(id string) (*conekta.Customer, *conekta.Error) {
	cust := &conekta.Customer{}
	err := conekta.MakeRequest("GET", "/customers/"+id, &conekta.EmptyParams{}, cust)
	return cust, err
}

// Delete deletes a customer
func Delete(id string) (*conekta.Customer, *conekta.Error) {
	cust := &conekta.Customer{}
	err := conekta.MakeRequest("DELETE", "/customers/"+id, &conekta.EmptyParams{}, cust)
	return cust, err
}

// All gets all customers
func All() (*conekta.CustomerList, *conekta.Error) {
	cl := &conekta.CustomerList{}
	err := conekta.MakeRequest("GET", "/customers", &conekta.EmptyParams{}, cl)
	return cl, err
}
