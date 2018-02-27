package taxlines

import (
	conekta "github.com/conekta-go"
)

// New creates a new customer
// For details see https://developers.conekta.com/api#create-customer
func New(p *conekta.TaxLinesParams) (*conekta.TaxLines, *conekta.Error) {
	tx := &conekta.TaxLines{}
	err := conekta.New(tx, p, "/orders/"+p.ID+"/tax_lines")
	return tx, err
}
