package conekta

// DestinationParams describes params for destination payee object request
type DestinationParams struct {
	Type              string `json:"type,omitempty"`
	AccountNumber     string `json:"account_number,omitempty"`
	AccountHolderName string `json:"account_holder_name,omitempty"`
}

// Destination describes destination payee object response
type Destination struct {
	ID                string `json:"id,omitempty"`
	PayeeID           string `json:"payee_id,omitempty"`
	Object            string `json:"object,omitempty"`
	CreatedAt         int64  `json:"created_at,omitempty"`
	Type              string `json:"type,omitempty"`
	Currency          string `json:"currency,omitempty"`
	Last4             string `json:"last4,omitempty"`
	AccountHolderName string `json:"account_holder_name,omitempty"`
	Bank              string `json:"bank,omitempty"`
}

// Bytes converts a DestinationParams to []byte
func (p *DestinationParams) Bytes() []byte {
	return paramsToBytes(p)
}
