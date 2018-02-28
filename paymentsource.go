package conekta

// PaymentSource represents a Conekta payment source.
// For details see https://developers.conekta.com/api#payment-source.
type PaymentSource struct {
	ID          string `json:"id,omitempty"`
	Object      string `json:"object,omitempty"`
	PaymentType string `json:"type,omitempty"`
	Last4       string `json:"last4,omitempty"`
	CreatedAt   int64  `json:"created_at,omitempty"`
	Name        string `json:"name,omitempty"`
	ExpMonth    string `json:"exp_month,omitempty"`
	ExpYear     string `json:"exp_year,omitempty"`
	Brand       string `json:"brand,omitempty"`
	ParentID    string `json:"parent_id,omitempty"`
}

// PaymentSourceParams is the set of parameters that can be used when creating a payment source.
// For details see https://developers.conekta.com/api#create-payment-source.
type PaymentSourceParams struct {
	TokenID     string `json:"token_id,omitempty"`
	PaymentType string `json:"type,omitempty"`
}

// PaymentSourceList is a list of  payment sources.
type PaymentSourceList struct {
	ListMeta
	Data []*PaymentSource `json:"data,omitempty"`
}

// Bytes converts a PaymentSourceParams to []byte
func (p *PaymentSourceParams) Bytes() []byte {
	return paramsToBytes(p)
}
