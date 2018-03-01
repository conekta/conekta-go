package conekta

//ChargeParams is the object to fill after api call
type ChargeParams struct {
	PaymentMethodParams *PaymentMethodParams `json:"payment_method,omitempty"`
}

//PaymentMethodParams is the object to fill after api call
type PaymentMethodParams struct {
	Type      string `json:"type,omitempty"`
	TokenID   string `json:"token_id,omitempty"`
	ExpiresAt int    `json:"expires_at,omitempty"`
}

//Charge should be a struct of the api response
type Charge struct {
	ID                  string         `json:"id,omitempty"`
	Object              string         `json:"object,omitempty"`
	Status              string         `json:"status,omitempty"`
	Amount              string         `json:"amount,omitempty"`
	Fee                 string         `json:"fee,omitempty"`
	OrderID             string         `json:"order_id,omitempty"`
	Livemode            string         `json:"livemode,omitempty"`
	CreatedAt           string         `json:"created_at,omitempty"`
	Currency            string         `json:"currency,omitempty"`
	PaymentMethodParams *PaymentMethod `json:"payment_method,omitempty"`
}

//PaymentMethod should be a struct of the api response
type PaymentMethod struct {
	ID        string          `json:"id,omitempty"`
	Object    string          `json:"object,omitempty"`
	Type      string          `json:"type,omitempty"`
	CreatedAt int             `json:"created_at,omitempty"`
	Last4     string          `json:"last4,omitempty"`
	Bin       string          `json:"bin,omitempty"`
	ExpMonth  string          `json:"exp_month,omitempty"`
	ExpYear   string          `json:"exp_year,omitempty"`
	Brand     string          `json:"brand,omitempty"`
	Name      string          `json:"name,omitempty"`
	ParentID  string          `json:"parent_id,omitempty"`
	Default   bool            `json:"default,omitempty"`
	Address   *DefaultAddress `json:"address,omitempty"`
}

//DefaultAddress should be nested struct of PaymentMethod
type DefaultAddress struct {
	Street1    string `json:"street1,omitempty"`
	Street2    string `json:"street2,omitempty"`
	City       string `json:"city,omitempty"`
	State      string `json:"state,omitempty"`
	Country    string `json:"country,omitempty"`
	Object     string `json:"object,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
}

//ChargesList is a list of shippingLines
type ChargesList struct {
	ListMeta
	Data []*Charge `json:"data,omitempty"`
}

// Bytes converts a ChargeParams to []byte
func (c *ChargeParams) Bytes() []byte {
	return paramsToBytes(c)
}
