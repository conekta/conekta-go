package conekta

//
type DiscountLinesParams struct {
	Code   string `json:"code"`
	Amount int    `json:"amount"`
	Type   string `json:"type"`
}

//DiscountLines should be nested struct of order
type DiscountLines struct {
	Code     string `json:"code"`
	Amount   int    `json:"amount"`
	Type     string `json:"type"`
	Object   string `json:"object"`
	ID       string `json:"id"`
	ParentID string `json:"parent_id"`
}

// Bytes converts a ChargeParams to []byte
func (c *DiscountLines) Bytes() []byte {
	return paramsToBytes(c)
}
