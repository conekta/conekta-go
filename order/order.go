package order

import "github.com/conekta/charge"

//DefaultParams should be a struct of the api response
type DefaultParams struct {
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
	LineItemsParams       LineItemsParams       `json:"line_items"`
	TaxLinesParams        TaxLinesParams        `json:"tax_lines"`
	ShippingLinesParams   ShippingLinesParams   `json:"shippping_lines"`
	DiscountLinesParams   DiscountLinesParams   `json:"discount_lines"`
	ShippingContactParams ShippingContactParams `json:"shipping_contact"`
	ChargeParams          charge.DefaultParams  `json:"charge"`
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

//TaxLinesParams should be nested struct of order
type TaxLinesParams struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	ParentID    string `json:"parent_id"`
	Object      string `json:"object"`
}

//ShippingLinesParams should be nested struct of order
type ShippingLinesParams struct {
	Amount         int    `json:"amount"`
	Carrier        string `json:"carrier"`
	Method         string `json:"method"`
	TrackingNumber string `json:"tracking_number"`
	Object         string `json:"object"`
	ID             string `json:"id"`
	ParentID       string `json:"parent_id"`
}

//DiscountLinesParams should be nested struct of order
type DiscountLinesParams struct {
	Code     string `json:"code"`
	Amount   int    `json:"amount"`
	Type     string `json:"type"`
	Object   string `json:"object"`
	ID       string `json:"id"`
	ParentID string `json:"parent_id"`
}

//ShippingContactParams should be nested struct of order
type ShippingContactParams struct {
	ID                     string                 `json:"id"`
	Object                 string                 `json:"object"`
	CreatedAt              int                    `json:"created_at"`
	ParentID               string                 `json:"parent_id"`
	Receiver               string                 `json:"receiver"`
	Phone                  string                 `json:"phone"`
	BetweenStreets         string                 `json:"between_streets"`
	ShippignContactAddress ShippignContactAddress `json:"address"`
}

//ShippignContactAddress should be nested struct of order
type ShippignContactAddress struct {
	Street1     string `json:"street1"`
	City        string `json:"city"`
	State       string `json:"state"`
	PostalCode  string `json:"postal_code"`
	Country     string `json:"country"`
	Residential bool   `json:"residential"`
	Object      string `json:"object"`
}
