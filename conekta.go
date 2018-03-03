package conekta

// TestKey is Conekta test key
const TestKey = "key_ZLy4aP2szht1HqzkCezDEA"

// TestPublic is Conekta test public key
const TestPublic = "key_OfWoJc2fw6oEydKspmyr76Q"

const (
	apiBase    = "https://api.conekta.io"
	apiVersion = "2.0.0"
	version    = "0.0.1"
)

// APIKey is Conekta private_key
// For details see https://developers.conekta.com/api#authentication
var APIKey string

// PubliAPIKey is Conekta private_key
// For details see https://developers.conekta.com/api#authentication
var PubliAPIKey string

// Locale is a settings for localization language
var Locale string

// ListMeta describes common metadata for lists
type ListMeta struct {
	Object      string `json:"object,omitempty"`
	HasMore     bool   `json:"has_more,omitempty"`
	Total       int    `json:"total,omitempty"`
	NextPageURL string `json:"next_page_url,omitempty"`
}

// Address represents a Conekta common Adress
type Address struct {
	Street1     string `json:"street1,omitempty"`
	Street2     string `json:"street2,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	Country     string `json:"country,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
	Object      string `json:"object,omitempty"`
	Residential bool   `json:"residential,omitempty"`
}

// EmptyParams is used in requests that don't need params
type EmptyParams struct{}

// Bytes converts a EmptyParams to []byte
func (p *EmptyParams) Bytes() []byte { return []byte{} }
