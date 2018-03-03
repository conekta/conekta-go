package conekta

//TaxLinesParams should return api object response filled in the struct
type TaxLinesParams struct {
	Description string `json:"description,omitempty"`
	Amount      int    `json:"amount,omitempty"`
}

//TaxLines should be nested struct of order
type TaxLines struct {
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Amount      int    `json:"amount,omitempty"`
	ParentID    string `json:"parent_id,omitempty"`
	Object      string `json:"object,omitempty"`
	Deleted     bool   `json:"deleted,omitempty"`
}

//TaxLinesList is a list of shippingLines
type TaxLinesList struct {
	ListMeta
	Data []*TaxLines `json:"data,omitempty"`
}

// Bytes converts a ChargeParams to []byte
func (c *TaxLinesParams) Bytes() []byte {
	return paramsToBytes(c)
}
