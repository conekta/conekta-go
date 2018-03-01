package taxlines

import (
	conekta "github.com/conekta/conekta-go"
)

//Create create discount line insde a order
// For details see https://developers.conekta.com/api#create-discount-line
func Create(orderID string, p *conekta.TaxLinesParams) (*conekta.TaxLines, *conekta.Error) {
	cust := &conekta.TaxLines{}
	err := conekta.MakeRequest("POST", "/orders/"+orderID+"/tax_lines", p, cust)
	return cust, err
}

//Update create discount line insde a order
// For details see https://developers.conekta.com/api#create-discount-line
func Update(orderID string, id string, p *conekta.TaxLinesParams) (*conekta.TaxLines, *conekta.Error) {
	cust := &conekta.TaxLines{}
	err := conekta.MakeRequest("PUT", "/orders/"+orderID+"/tax_lines/"+id, p, cust)
	return cust, err
}

//Delete create discount line insde a order
// For details see https://developers.conekta.com/api#create-discount-line
func Delete(orderID string, id string) (*conekta.TaxLines, *conekta.Error) {
	tl := &conekta.TaxLines{}
	err := conekta.MakeRequest("DELETE", "/orders/"+orderID+"/tax_lines/"+id, &conekta.OrderParams{}, tl)
	return tl, err
}
