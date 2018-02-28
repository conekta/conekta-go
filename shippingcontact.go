package conekta

// ShippingContact represents a Conekta shipping contact.
// For details see https://developers.conekta.com/api#shipping-contact.
type ShippingContact struct {
	ID             string   `json:"id,omitempty"`
	CreatedAt      int64    `json:"created_at,omitempty"`
	Receiver       string   `json:"receiver,omitempty"`
	Phone          string   `json:"phone,omitempty"`
	BetweenStreets string   `json:"between_streets,omitempty"`
	Object         string   `json:"object,omitempty"`
	ParentID       string   `json:"parent_id,omitempty"`
	Address        *Address `json:"address,omitempty"`
	Deleted        bool     `json:"deleted,omitempty"`
	Default        bool     `json:"default,omitempty"`
}

// ShippingContactParams is the set of parameters that can be used when creating or updating a shipping contact.
type ShippingContactParams struct {
	Phone          string   `json:"phone,omitempty"`
	Receiver       string   `json:"receiver,omitempty"`
	BetweenStreets string   `json:"between_streets,omitempty"`
	Address        *Address `json:"address,omitempty"`
}

// ShippingContactList is a list of shipping contact.
type ShippingContactList struct {
	ListMeta
	Data []*ShippingContact `json:"data,omitempty"`
}

// Bytes converts a CustomerParams to []byte
func (p *ShippingContactParams) Bytes() []byte {
	return paramsToBytes(p)
}
