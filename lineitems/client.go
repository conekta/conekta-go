package lineitems

import (
	conekta "github.com/conekta/conekta-go"
)

//Create function creates a line_item object
//For more information please see https://developers.conekta.com/api#create-line-item
func Create(id string, p *conekta.LineItemsParams) (*conekta.LineItems, *conekta.Error) {
	li := &conekta.LineItems{}
	err := conekta.MakeRequest("POST", "/orders/"+id+"/line_items", p, li)
	return li, err
}

// Update updates a Order
// For details see https://developers.conekta.com/api#update-Order
func Update(orderID string, id string, p *conekta.LineItemsParams) (*conekta.LineItems, *conekta.Error) {
	li := &conekta.LineItems{}
	err := conekta.MakeRequest("PUT", "/orders/"+id+"/line_items/"+id, p, li)
	return li, err
}

// Delete deletes a Order
func Delete(orderID string, id string) (*conekta.LineItems, *conekta.Error) {
	li := &conekta.LineItems{}
	err := conekta.MakeRequest("DELETE", "/orders/"+orderID+"/line_items/"+id, &conekta.LineItemsParams{}, li)
	return li, err
}
