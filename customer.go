package conekta

// Customer represents a Conekta customer.
// For details see https://developers.conekta.com/api#customer.
type Customer struct {
	ID                       string              `json:"id,omitempty"`
	Object                   string              `json:"object,omitempty"`
	Live                     bool                `json:"livemode,omitempty"`
	CreatedAt                int64               `json:"created_at,omitempty"`
	Name                     string              `json:"name,omitempty"`
	Email                    string              `json:"email,omitempty"`
	Phone                    string              `json:"phone,omitempty"`
	Corporate                bool                `json:"corporate,omitempty"`
	DefaultShippingContactID string              `json:"default_shipping_contact_id,omitempty"`
	DefaultPaymentSourceID   string              `json:"default_payment_source_id,omitempty"`
	PaymentSources           PaymentSourceList   `json:"payment_sources,omitempty"`
	ShippingContacts         ShippingContactList `json:"shipping_contacts,omitempty"`
	Subscription             *Subscription       `json:"subscription,omitempty"`
	Deleted                  bool                `json:"deleted,omitempty"`
}

// CustomerParams is the set of parameters that can be used when creating or updating a customer.
// For details see https://developers.conekta.com/api#create-customer and https://developers.conekta.com/api#update-customer.
type CustomerParams struct {
	ID                     string                       `json:"customer_id,omitempty"`
	Name                   string                       `json:"name,omitempty"`
	Phone                  string                       `json:"phone,omitempty"`
	Email                  string                       `json:"email,omitempty"`
	PaymentSources         []*PaymentSourceCreateParams `json:"payment_sources,omitempty"`
	Corporate              bool                         `json:"corporate,omitempty"`
	DefaultPaymentSourceID string                       `json:"default_payment_source_id,omitempty"`
	ShippingContacts       []*ShippingContactParams     `json:"shipping_contacts,omitempty"`
}

// CustomerList is a list of customers
type CustomerList struct {
	ListMeta
	Data []*Customer `json:"data,omitempty"`
}

// Bytes converts a CustomerParams to []byte
func (p *CustomerParams) Bytes() []byte {
	return paramsToBytes(p)
}
