package conekta

// LineItems should be nested struct of order
type LineItems struct {
	ID        string `json:"id,omitempty"`
	Object    string `json:"object,omitempty"`
	Name      string `json:"name,omitempty"`
	UnitPrice int64  `json:"unit_price,omitempty"`
	Quantity  int    `json:"quantity,omitempty"`
	Deleted   bool   `json:"deleted,omitempty"`
}

// LineItemsParams should be response wrapper for api line_items endpoint
type LineItemsParams struct {
	Name      string   `json:"name,omitempty"`
	UnitPrice int64    `json:"unit_price,omitempty"`
	Quantity  int      `json:"quantity,omitempty"`
	Metadata  struct{} `json:"metadata,omitempty"`
}

// LineItemsList is a list of shippingLines
type LineItemsList struct {
	ListMeta
	Data []*LineItems `json:"data,omitempty"`
}

// Bytes converts a ChargeParams to []byte
func (c *LineItemsParams) Bytes() []byte {
	return paramsToBytes(c)
}
