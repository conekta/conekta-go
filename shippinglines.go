package conekta

//ShippingLinesParams is the set of parameters that can be used when creating or updating a shipping contact.
type ShippingLinesParams struct {
	Amount         int64  `json:"amount,omitempty"`
	TrackingNumber string `json:"tracking_number,omitempty"`
	Carrier        string `json:"carrier,omitempty"`
	Method         string `json:"method,omitempty"`
}

//ShippingLines should be nested struct of order
type ShippingLines struct {
	Amount         int64  `json:"amount,omitempty"`
	Carrier        string `json:"carrier,omitempty"`
	Method         string `json:"method,omitempty"`
	TrackingNumber string `json:"tracking_number,omitempty"`
	Object         string `json:"object,omitempty"`
	ID             string `json:"id,omitempty"`
	ParentID       string `json:"parent_id,omitempty"`
	Deleted        bool   `json:"deleted,omitempty"`
}

//ShippingLinesList is a list of shippingLines
type ShippingLinesList struct {
	ListMeta
	Data []*ShippingLines `json:"data,omitempty"`
}

// Bytes converts a ChargeParams to []byte
func (c *ShippingLinesParams) Bytes() []byte {
	return paramsToBytes(c)
}
