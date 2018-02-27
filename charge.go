package conekta

//
type ChargeParams struct {
	PaymentMethodParams PaymentMethodParams `json:"payment_method"`
}

//
type PaymentMethodParams struct {
	Type      string `json:"type"`
	ExpiresAt int    `json:"expires_at"`
}

//Charge should be a struct of the api response
type Charge struct {
	ID                  string        `json:"id"`
	Object              string        `json:"object"`
	Status              string        `json:"status"`
	Amount              string        `json:"amount"`
	Fee                 string        `json:"fee"`
	OrderID             string        `json:"order_id"`
	Livemode            string        `json:"livemode"`
	CreatedAt           string        `json:"created_at"`
	Currency            string        `json:"currency"`
	PaymentMethodParams PaymentMethod `json:"payment_method"`
}

//PaymentMethod should be a struct of the api response
type PaymentMethod struct {
	ID        string         `json:"id"`
	Object    string         `json:"object"`
	Type      string         `json:"type"`
	CreatedAt int            `json:"created_at"`
	Last4     string         `json:"last4"`
	Bin       string         `json:"bin"`
	ExpMonth  string         `json:"exp_month"`
	ExpYear   string         `json:"exp_year"`
	Brand     string         `json:"brand"`
	Name      string         `json:"name"`
	ParentID  string         `json:"parent_id"`
	Default   bool           `json:"default"`
	Address   DefaultAddress `json:"address"`
}

//DefaultAddress should be nested struct of PaymentMethod
type DefaultAddress struct {
	Street1    string `json:"street1"`
	Street2    string `json:"street2"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	Object     string `json:"object"`
	PostalCode string `json:"postal_code"`
}

// Bytes converts a ChargeParams to []byte
func (c *PaymentMethod) Bytes() []byte {
	return paramsToBytes(c)
}

// Bytes converts a ChargeParams to []byte
func (c *ChargeParams) Bytes() []byte {
	return paramsToBytes(c)
}
