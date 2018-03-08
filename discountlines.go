package conekta

//DiscountLinesParams is the object to fill after api call
type DiscountLinesParams struct {
	Code   string `json:"code,omitempty"`
	Amount int64  `json:"amount,omitempty"`
	Type   string `json:"type,omitempty"`
}

//DiscountLines should be nested struct of order
type DiscountLines struct {
	Code     string `json:"code,omitempty"`
	Amount   int64  `json:"amount,omitempty"`
	Type     string `json:"type,omitempty"`
	Object   string `json:"object,omitempty"`
	ID       string `json:"id,omitempty"`
	ParentID string `json:"parent_id,omitempty"`
	Deleted  bool   `json:"deleted,omitempty"`
}

//DiscountLinesList is a list of DiscountLines
type DiscountLinesList struct {
	ListMeta
	Data []*DiscountLines `json:"data,omitempty"`
}

// Bytes converts a ChargeParams to []byte
func (c *DiscountLinesParams) Bytes() []byte {
	return paramsToBytes(c)
}
