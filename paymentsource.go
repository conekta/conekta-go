package conekta

// PaymentSource represents a Conekta payment source.
// For details see https://developers.conekta.com/api#payment-source.
type PaymentSource struct {
	ID          string   `json:"id,omitempty"`
	Object      string   `json:"object,omitempty"`
	PaymentType string   `json:"type,omitempty"`
	Last4       string   `json:"last4,omitempty"`
	CreatedAt   int64    `json:"created_at,omitempty"`
	Name        string   `json:"name,omitempty"`
	ExpMonth    string   `json:"exp_month,omitempty"`
	ExpYear     string   `json:"exp_year,omitempty"`
	Brand       string   `json:"brand,omitempty"`
	ParentID    string   `json:"parent_id,omitempty"`
	Address     *Address `json:"address,omitempty"`
	Deleted     bool     `json:"deleted,omitempty"`
	Default     bool     `json:"default,omitempty"`
}

// PaymentSourceCreateParams is the set of parameters that can be used when creating a payment source.
// For details see https://developers.conekta.com/api#create-payment-source.
type PaymentSourceCreateParams struct {
	TokenID     string `json:"token_id,omitempty"`
	PaymentType string `json:"type,omitempty"`
}

// PaymentSourceUpdateParams is the set of parameters that can be used when update a payment source.
// For details see https://developers.conekta.com/api#update-payment-source.
type PaymentSourceUpdateParams struct {
	Name     string   `json:"name,omitempty"`
	ExpMonth string   `json:"exp_month,omitempty"`
	ExpYear  string   `json:"exp_year,omitempty"`
	Address  *Address `json:"address,omitempty"`
}

// PaymentSourceList is a list of  payment sources.
type PaymentSourceList struct {
	ListMeta
	Data []*PaymentSource `json:"data,omitempty"`
}

// Bytes converts a PaymentSourceParams to []byte
func (p *PaymentSourceCreateParams) Bytes() []byte {
	return paramsToBytes(p)
}

// Bytes converts a PaymentSourceParams to []byte
func (p *PaymentSourceUpdateParams) Bytes() []byte {
	return paramsToBytes(p)
}
