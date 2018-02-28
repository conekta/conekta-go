package conekta

const (
	apiBase    = "https://api.conekta.io"
	apiVersion = "2.0.0"
	version    = "0.0.1"
)

// APIKey is Conekta private_key
// For details see https://developers.conekta.com/api#authentication
var APIKey string

// Locale is a settings for localization language
var Locale string

// ListMeta describes common metadata for lists
type ListMeta struct {
	Object  string `json:"object,omitempty"`
	HasMore bool   `json:"has_more,omitempty"`
	Total   int    `json:"total,omitempty"`
}
