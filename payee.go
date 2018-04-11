package conekta

// PayeeParams is the set of parameters of Payee request Object
type PayeeParams struct {
	Name           string                     `json:"name,omitempty"`
	Email          string                     `json:"email,omitempty"`
	Phone          string                     `json:"phone,omitempty"`
	BillingAddress *PayeeBillingAddressParams `json:"billing_address,omitempty"`
}

// PayeeBillingAddressParams describes billing address of Payee object
type PayeeBillingAddressParams struct {
	CompanyName string `json:"company_name,omitempty"`
	TaxID       string `json:"tax_id,omitempty"`
	Street1     string `json:"street1,omitempty"`
	Street2     string `json:"street2,omitempty"`
	Street3     string `json:"street3,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	Country     string `json:"country,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
}

// Payee describes Payee response object
type Payee struct {
	ID                   string                     `json:"id,omitempty"`
	Email                string                     `json:"email,omitempty"`
	Name                 string                     `json:"name,omitempty"`
	Phone                string                     `json:"phone,omitempty"`
	LiveMode             bool                       `json:"livemode,omitempty"`
	BillingAddress       *PayeeBillingAddressParams `json:"billing_address,omitempty"`
	Object               string                     `json:"object,omitempty"`
	DefaultDestinationID string                     `json:"default_destination_id,omitempty"`
	CreatedAt            int64                      `json:"created_at,omitempty"`
	Destinations         []*Destination             `json:"destinations,omitempty"`
	Deleted              bool                       `json:"deleted,omitempty"`
}

// PayeeList is a list of payees
type PayeeList struct {
	ListMeta
	Data []*Payee `json:"data,omitempty"`
}

// Bytes converts a PayeeParams to []byte
func (c *PayeeParams) Bytes() []byte {
	return paramsToBytes(c)
}
