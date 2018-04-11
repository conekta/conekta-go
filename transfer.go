package conekta

// TransferParams describes params for Transfer object request
type TransferParams struct {
	Amount   int64  `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
	PayeeID  string `json:"payee,omitempty"`
}

// Transfer describes Transfer object response
type Transfer struct {
	ID                   string       `json:"id,omitempty"`
	Amount               int64        `json:"amount,omitempty"`
	Currency             string       `json:"currency,omitempty"`
	Livemode             bool         `json:"livemode,omitempty"`
	Object               string       `json:"object,omitempty"`
	CreatedAt            int64        `json:"created_at,omitempty"`
	Status               string       `json:"status,omitempty"`
	StatementReference   string       `json:"statement_reference,omitempty"`
	StatementDescription string       `json:"statement_description,omitempty"`
	Destination          *Destination `json:"destination,omitempty"`
}

// Bytes converts a TransferParams to []byte
func (p *TransferParams) Bytes() []byte {
	return paramsToBytes(p)
}
