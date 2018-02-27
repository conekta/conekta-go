package conekta

//
type ShippingLinesParams struct {
	Description    string `json:"description"`
	Amount         int    `json:"amount"`
	TrackingNumber string `json:"tracking_number"`
	Carrier        string `json:"carrier"`
	Method         string `json:"method"`
}

//ShippingLines should be nested struct of order
type ShippingLines struct {
	Amount         int    `json:"amount"`
	Carrier        string `json:"carrier"`
	Method         string `json:"method"`
	TrackingNumber string `json:"tracking_number"`
	Object         string `json:"object"`
	ID             string `json:"id"`
	ParentID       string `json:"parent_id"`
}

// Bytes converts a ChargeParams to []byte
func (c *ShippingLines) Bytes() []byte {
	return paramsToBytes(c)
}
