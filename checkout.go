package conekta

// CheckoutRefundParams returns api response object filled
type CheckoutRefundParams struct {
	Reason string `json:"reason,omitempty"`
	Amount int64  `json:"amount,omitempty"`
}

//CheckoutParams returns api response object filled
type CheckoutParams struct {
	AllowedPaymentMethods      []string     `json:"allowed_payment_methods,omitempty"`
	ExpiredAt                  int64        `json:"expired_at,omitempty"`
	FailureUrl                 string       `json:"failure_url,omitempty"`
	PaymentsLimitCount         int64        `json:"payments_limit_count"`
	MultifactorAuthentication  bool         `json:"multifactor_autentication"`
	MonthlyInstallmentsEnabled bool         `json:"monthly_installments_enabled"`
	MonthlyInstallmentsOptions []int64      `json:"monthly_installments_options,omitempty"`
	Name                       string       `json:"name,omitempty"`
	NeedsShippingContact       bool         `json:"needs_shipping_contact"`
	OrderTemplate              *OrderParams `json:"order_template,omitempty"`
	OnDemandEnabled            bool         `json:"on_demand_enabled"`
	Recurrent                  bool         `json:"recurrent"`
	SuccessUrl                 string       `json:"success_url,omitempty"`
	Type                       string       `json:"type,omitempty"`
}

// Checkout should be a struct of the api response
type Checkout struct {
	AllowedInstallmentOptions  []int64  `json:"allowed_installment_options,omitempty"`
	AllowedPaymentMethods      []string `json:"allowed_payment_methods,omitempty"`
	EmailsSent                 int64    `json:"emails_sent,omitempty"`
	ExpiredAt                  int64    `json:"expired_at,omitempty"`
	ExpiresAt                  int64    `json:"expires_at,omitempty"`
	Force3dsFlow               bool     `json:"force_3ds_flow,omitempty"`
	ID                         string   `json:"id,omitempty"`
	Livemode                   bool     `json:"livemode,omitempty"`
	MonthlyInstallmentsEnabled bool     `json:"monthl_installments_enabled,omitempty"`
	Name                       string   `json:"name,omitempty"`
	NeedsShippingContact       bool     `json:"needs_shipping_contact,omitempty"`
	Object                     string   `json:"object,omitempty"`
	PaidPaymentsCount          int64    `json:"paid_payments_count,omitempty"`
	Recurrent                  bool     `json:"recurrent,omitempty"`
	Slug                       string   `json:"slug,omitempty"`
	SmsSent                    int64    `json:"sms_sent,omitempty"`
	Status                     string   `json:"status,omitempty"`
	Type                       string   `json:"type,omitempty"`
	Url                        string   `json:"url,omitempty"`
}

// CheckoutList is a list of shippingLines
type CheckoutList struct {
	ListMeta
	Data []*Checkout `json:"data,omitempty"`
}

// Bytes converts a ChargeParams to []byte
func (c *CheckoutParams) Bytes() []byte {
	return paramsToBytes(c)
}
