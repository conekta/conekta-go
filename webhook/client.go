package webhook

import (
	conekta "github.com/conekta/conekta-go/v2"
)

// Create creates a new webhook
func Create(p *conekta.WebhookParams) (*conekta.Webhook, error) {
	wh := &conekta.Webhook{}
	err := conekta.MakeRequest("POST", "/webhooks", p, wh)
	return wh, err
}

// Update updates a webhook
func Update(id string, p *conekta.WebhookParams) (*conekta.Webhook, error) {
	wh := &conekta.Webhook{}
	err := conekta.MakeRequest("PUT", "/webhooks/"+id, p, wh)
	return wh, err
}

// Find gets a webhook by id
func Find(id string) (*conekta.Webhook, error) {
	wh := &conekta.Webhook{}
	err := conekta.MakeRequest("GET", "/webhooks/"+id, &conekta.EmptyParams{}, wh)
	return wh, err
}

// Delete deletes a webhook
func Delete(id string) (*conekta.Webhook, error) {
	wh := &conekta.Webhook{}
	err := conekta.MakeRequest("DELETE", "/webhooks/"+id, &conekta.EmptyParams{}, wh)
	return wh, err
}

// All gets all webhooks
func All() (*conekta.WebhookList, error) {
	whl := &conekta.WebhookList{}
	err := conekta.MakeRequest("GET", "/webhooks", &conekta.EmptyParams{}, whl)
	return whl, err
}
