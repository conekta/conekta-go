package conekta

// Webhook represents a Conekta webhook.
// For details see https://developers.conekta.com/api#events.
type Webhook struct {
	ID                 string   `json:"id,omitempty"`
	Object             string   `json:"object,omitempty"`
	Status             string   `json:"status,omitempty"`
	Deleted            bool     `json:"deleted,omitempty"`
	URL                string   `json:"url,omitempty"`
	ProductionEnabled  bool     `json:"production_enabled,omitempty"`
	DevelopmentEnabled bool     `json:"development_enabled,omitempty"`
	SubscribedEvents   []string `json:"subscribed_events,omitempty"`
}

// WebhookParams is the set of parameters that can be used when creating or updating a webhook.
// For details see https://developers.conekta.com/api#create-customer and https://developers.conekta.com/api#update-customer.
type WebhookParams struct {
	URL                string `json:"url,omitempty"`
	ProductionEnabled  bool   `json:"production_enabled,omitempty"`
	DevelopmentEnabled bool   `json:"development_enabled,omitempty"`
}

// WebhookList is a list of webhooks
type WebhookList struct {
	ListMeta
	Data []*Webhook `json:"data,omitempty"`
}

// Bytes converts a WebhookParams to []byte
func (p *WebhookParams) Bytes() []byte {
	return paramsToBytes(p)
}
