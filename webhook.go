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

// WebhookNotification represents a Conekta webhook notification.
type WebhookNotification struct {
	Livemode      bool                     `json:"livemode,omitempty"`
	ID            string                   `json:"id,omitempty"`
	Object        string                   `json:"object,omitempty"`
	TypeWebhook   string                   `json:"type,omitempty"`
	WebhookStatus string                   `json:"webhook_status,omitempty"`
	CreatedAt     int64                    `json:"created_at,omitempty"`
	Data          *WebhookNotificationData `json:"data,omitempty"`
}

// WebhookNotificationData represents the data of Conekta webhook notification.
type WebhookNotificationData struct {
	Object             *WebhookNotificationObject                 `json:"object,omitempty"`
	PreviousAttributes *WebhookNotificationDataPreviousAttributes `json:"previous_attributes,omitempty"`
}

// WebhookNotificationDataPreviousAttributes represents previous attributes of Conekta webhook notification.
type WebhookNotificationDataPreviousAttributes struct {
	Status string `json:"status,omitempty"`
}

// WebhookNotificationObject describes webhook object
type WebhookNotificationObject struct {
	ID             string                     `json:"id,omitempty"`
	Status         string                     `json:"status,omitempty"`
	Currency       string                     `json:"currency,omitempty"`
	Description    string                     `json:"description,omitempty"`
	FailureCode    string                     `json:"failure_code,omitempty"`
	FailureMessage string                     `json:"failure_message,omitempty"`
	Object         string                     `json:"object,omitempty"`
	PaymentStatus  string                     `json:"payment_status,omitempty"`
	CustomerID     string                     `json:"customer_id,omitempty"`
	OrderID        string                     `json:"order_id,omitempty"`
	CreatedAt      int64                      `json:"created_at,omitempty"`
	UpdatedAt      int64                      `json:"updated_at,omitempty"`
	PaidAt         int64                      `json:"paid_at,omitempty"`
	Fee            int64                      `json:"fee,omitempty"`
	AmountRefunded int64                      `json:"amount_refunded,omitempty"`
	Amount         int64                      `json:"amount,omitempty"`
	Livemode       bool                       `json:"livemode,omitempty"`
	Method         *WebhookNotificationMethod `json:"method,omitempty"`
	PaymentMethod  *PaymentMethod             `json:"payment_method,omitempty"`
	LineItems      *LineItemsList             `json:"line_items,omitempty"`
	Charges        *ChargesList               `json:"charges,omitempty"`
	CustomerInfo   *Customer                  `json:"customer_info,omitempty"`
	WebhookLogs    []*WebhookLog              `json:"webhook_logs,omitempty"`
}

// WebhookNotificationMethod is the inheritance object for Payee webhook
type WebhookNotificationMethod struct {
	CreatedAt     int64  `json:"created_at,omitempty"`
	PayeeID       string `json:"payee_id,omitempty"`
	Object        string `json:"object,omitempty"`
	AccountNumber string `json:"account_number,omitempty"`
	AccountHolder string `json:"account_holder,omitempty"`
	Bank          string `json:"bank,omitempty"`
}

// WebhookLog describes webhook log object
type WebhookLog struct {
	ID                     string `json:"id,omitempty"`
	URL                    string `json:"url,omitempty"`
	FailedAttempts         string `json:"failed_attempts,omitempty"`
	LastHTTPResponseStatus string `json:"last_http_response_status,omitempty"`
	Object                 string `json:"object,omitempty"`
	LastAttemptedAt        string `json:"last_attempted_at,omitempty"`
}

// Bytes converts a WebhookParams to []byte
func (p *WebhookParams) Bytes() []byte {
	return paramsToBytes(p)
}
