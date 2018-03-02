package conekta

//OrderParams returns api response object filled
type OrderParams struct {
	Currency        string                 `json:"currency,omitempty"`
	CustomerInfo    *CustomerParams        `json:"customer_info,omitempty"`
	LineItems       []*LineItemsParams     `json:"line_items,omitempty"`
	TaxLines        []*TaxLinesParams      `json:"tax_lines,omitempty"`
	ShippingContact *ShippingContactParams `json:"shipping_contact,omitempty"`
	DiscountLines   []*DiscountLinesParams `json:"discount_lines,omitempty"`
	ShippingLines   []*ShippingLinesParams `json:"shipping_lines,omitempty"`
	Charges         []*ChargeParams        `json:"charges,omitempty"`
}

//Order should be a struct of the api response
type Order struct {
	ID              string             `json:"id,omitempty"`
	Object          string             `json:"object,omitempty"`
	Livemode        bool               `json:"livemode,omitempty"`
	Amount          int                `json:"amount,omitempty"`
	AmountRefunded  int                `json:"amount_refunded,omitempty"`
	PaymentStatus   string             `json:"payment_status,omitempty"`
	Currency        string             `json:"currency,omitempty"`
	CreatedAt       int                `json:"created_at,omitempty"`
	UpdatedAt       int                `json:"updated_at,omitempty"`
	Metadata        struct{}           `json:"metadata,omitempty"`
	CustomerInfo    *Customer          `json:"customer_info,omitempty"`
	ShippingContact *ShippingContact   `json:"shipping_contact,omitempty"`
	LineItems       *LineItemsList     `json:"line_items,omitempty"`
	TaxLines        *TaxLinesList      `json:"tax_lines,omitempty"`
	ShippingLines   *ShippingLinesList `json:"shipping_lines,omitempty"`
	DiscountLines   *DiscountLinesList `json:"discount_lines,omitempty"`
	ChargeParams    *ChargesList       `json:"charge,omitempty"`
}

// Bytes converts a ChargeParams to []byte
func (c *OrderParams) Bytes() []byte {
	return paramsToBytes(c)
}
