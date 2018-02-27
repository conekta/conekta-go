package charge

import "github.com/conekta/paymentmethod"

//DefaultParams should be a struct of the api response
type DefaultParams struct {
	ID                  string                      `json:"id"`
	Object              string                      `json:"object"`
	Status              string                      `json:"status"`
	Amount              string                      `json:"amount"`
	Fee                 string                      `json:"fee"`
	OrderID             string                      `json:"order_id"`
	Livemode            string                      `json:"livemode"`
	CreatedAt           string                      `json:"created_at"`
	Currency            string                      `json:"currency"`
	PaymentMethodParams paymentmethod.DefaultParams `json:"payment_method"`
}
