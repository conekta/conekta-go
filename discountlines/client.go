package discountlines

import conekta "github.com/conekta/conekta-go/v2"

//Create create discount line insde a order
// For details see https://developers.conekta.com/api#create-discount-line
func Create(orderID string, p *conekta.DiscountLinesParams) (*conekta.DiscountLines, error) {
	dlp := &conekta.DiscountLines{}
	err := conekta.MakeRequest("POST", "/orders/"+orderID+"/discount_lines", p, dlp)
	return dlp, err
}

// Update updates a DiscountLine inside order
// For details see https://developers.conekta.com/api#update-discount-line
func Update(orderID string, id string, p *conekta.DiscountLinesParams) (*conekta.DiscountLines, error) {
	ord := &conekta.DiscountLines{}
	err := conekta.MakeRequest("PUT", "/orders/"+orderID+"/discount_lines/"+id, p, ord)
	return ord, err
}

// Delete deletes a Order Discount Line
func Delete(orderID string, id string) (*conekta.DiscountLines, error) {
	ord := &conekta.DiscountLines{}
	err := conekta.MakeRequest("DELETE", "/orders/"+orderID+"/discount_lines/"+id, &conekta.OrderParams{}, ord)
	return ord, err
}
