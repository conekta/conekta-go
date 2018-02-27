package conekta

//
type TaxLinesParams struct {
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

//TaxLines should be nested struct of order
type TaxLines struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	ParentID    string `json:"parent_id"`
	Object      string `json:"object"`
}

// Bytes converts a ChargeParams to []byte
func (c *TaxLines) Bytes() []byte {
	return paramsToBytes(c)
}
