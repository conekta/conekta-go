package conekta

const (
	// OrderRefundRequestedByClient is a reason for refund
	OrderRefundRequestedByClient = "requested_by_client"
	// OrderRefundCannotBeFulfilled is a reason for refund
	OrderRefundCannotBeFulfilled = "cannot_be_fulfilled"
	// OrderRefundDuplicatedTransaction is a reason for refund
	OrderRefundDuplicatedTransaction = "duplicated_transaction"
	// OrderRefundSuspectedFraud is a reason for refund
	OrderRefundSuspectedFraud = "suspected_fraud"
	// OrderRefundOther is a reason for refund
	OrderRefundOther = "other"
)

// OrderRefundParams returns api response object filled
type OrderRefundParams struct {
	Reason string `json:"reason,omitempty"`
	Amount int64  `json:"amount,omitempty"`
}

//OrderParams returns api response object filled
type OrderParams struct {
	Currency        string                 `json:"currency,omitempty"`
	PreAuth         bool                   `json:"pre_authorize,omitempty"`
	CustomerInfo    *CustomerParams        `json:"customer_info,omitempty"`
	LineItems       []*LineItemsParams     `json:"line_items,omitempty"`
	TaxLines        []*TaxLinesParams      `json:"tax_lines,omitempty"`
	ShippingContact *ShippingContactParams `json:"shipping_contact,omitempty"`
	DiscountLines   []*DiscountLinesParams `json:"discount_lines,omitempty"`
	ShippingLines   []*ShippingLinesParams `json:"shipping_lines,omitempty"`
	Metadata        struct{}               `json:"metadata,omitempty"`
	Charges         []*ChargeParams        `json:"charges,omitempty"`
	Checkout        *OrderCheckoutParams   `json:"checkout,omitempty"`
}

// Order should be a struct of the api response
type Order struct {
	ID              string             `json:"id,omitempty"`
	Object          string             `json:"object,omitempty"`
	Livemode        bool               `json:"livemode,omitempty"`
	Amount          int64              `json:"amount,omitempty"`
	AmountRefunded  int64              `json:"amount_refunded,omitempty"`
	PaymentStatus   string             `json:"payment_status,omitempty"`
	Currency        string             `json:"currency,omitempty"`
	CreatedAt       int64              `json:"created_at,omitempty"`
	UpdatedAt       int64              `json:"updated_at,omitempty"`
	PreAuth         bool               `json:"pre_authorize,omitempty"`
	Metadata        struct{}           `json:"metadata,omitempty"`
	CustomerInfo    *CustomerInfo      `json:"customer_info,omitempty"`
	ShippingContact *ShippingContact   `json:"shipping_contact,omitempty"`
	LineItems       *LineItemsList     `json:"line_items,omitempty"`
	TaxLines        *TaxLinesList      `json:"tax_lines,omitempty"`
	ShippingLines   *ShippingLinesList `json:"shipping_lines,omitempty"`
	DiscountLines   *DiscountLinesList `json:"discount_lines,omitempty"`
	Charges         *ChargesList       `json:"charges,omitempty"`
	Checkout        *Checkout          `json:"checkout,omitempty"`
}

type OrderCheckoutParams struct {
	Name                       string   `json:"name,omitempty"`
	Type                       string   `json:"type,omitempty"`
	Recurrent                  bool     `json:"recurrent,omitempty"`
	ExpiresAt                  int64    `json:"expires_at,omitempty"`
	AllowedPaymentMethods      []string `json:"allowed_payment_methods,omitempty"`
	NeedsShippingContact       bool     `json:"needs_shipping_contact,omitempty"`
	MonthlyInstallmentsEnabled bool     `json:"monthly_installments_enabled,omitempty"`
	MonthlyInstallmentsOptions []int64  `json:"monthly_installments_options,omitempty"`
}

// CustomerInfo describes customer info
type OrderCustomerInfoParams struct {
	CustomerID string `json:"customer_id,omitempty"`
	Name       string `json:"name,omitempty"`
	Phone      string `json:"phone,omitempty"`
	Email      string `json:"email,omitempty"`
	Corporate  bool   `json:"corporate,omitempty"`
}

// CustomerInfo describes customer info
type CustomerInfo struct {
	CustomerID string `json:"customer_id,omitempty"`
	*Customer
}

// OrderList is a list of shippingLines
type OrderList struct {
	ListMeta
	Data []*Order `json:"data,omitempty"`
}

// Bytes converts a ChargeParams to []byte
func (c *OrderParams) Bytes() []byte {
	return paramsToBytes(c)
}

// Bytes converts a OrderRefundParams to []byte
func (c *OrderRefundParams) Bytes() []byte {
	return paramsToBytes(c)
}
