package conekta

//
type OrderParams struct {
	Currency     string             `json:"currency"`
	CustomerInfo CustomerInfoParams `json:"customer_info"`
	LineItems    []LineItemsParams  `json:"line_items"`
	Charges      []ChargeParams     `json:"charges"`
}

//Order should be a struct of the api response
type Order struct {
	ID             string `json:"id"`
	Object         string `json:"object"`
	Livemode       bool   `json:"livemode"`
	Amount         int    `json:"amount"`
	AmountRefunded int    `json:"amount_refunded"`
	PaymentStatus  string `json:"payment_status"`
	Currency       string `json:"currency"`
	CreatedAt      int    `json:"created_at"`
	UpdatedAt      int    `json:"updated_at"`
	Metadata       struct {
	} `json:"metadata"`
	LineItemsParams       LineItemsParams `json:"line_items"`
	TaxLinesParams        TaxLines        `json:"tax_lines"`
	ShippingLinesParams   ShippingLines   `json:"shippping_lines"`
	DiscountLinesParams   DiscountLines   `json:"discount_lines"`
	ShippingContactParams interface{}     //FIX ME
	ChargeParams          Charge          `json:"charge"`
}

//CustomerInfoParams should be nested struct of order
//TODO: MAKE REFERENCE FROM  REAL CUSTOMER STRUCT
type CustomerInfoParams struct {
	Object     string `json:"object"`
	CustomerID string `json:"customer_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}

//LineItemsParams should be nested struct of order
type LineItemsParams struct {
	Object  string `json:"object"`
	HasMore bool   `json:"has_more"`
	Total   int    `json:"total"`
	Data    []struct {
		ID            string `json:"id"`
		Object        string `json:"object"`
		Name          string `json:"name"`
		UnitPrice     int    `json:"unit_price"`
		Quantity      int    `json:"quantity"`
		ParentID      string `json:"parent_id"`
		AntifraudInfo struct {
		} `json:"antifraud_info"`
		Metadata struct {
		} `json:"metadata"`
	}
}

// Bytes converts a ChargeParams to []byte
func (c *Order) Bytes() []byte {
	return paramsToBytes(c)
}
