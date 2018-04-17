package api

import (
)

type Event struct {
    Data  map[string]string  `json:"data,omitempty"`
    WebhookStatus  string  `json:"webhook_status,omitempty"`
    WebhookLogs  []WebhookLog  `json:"webhook_logs,omitempty"`
    Livemode  string  `json:"livemode,omitempty"`
    Id  string  `json:"id,omitempty"`
    Object  string  `json:"object,omitempty"`
    Type_  string  `json:"type,omitempty"`
    CreatedAt  string  `json:"created_at,omitempty"`
    
}


func Events()([]Event, error){
    api := NewDefaultApi(Conekta().apiKey);
    return api.EventsGet()
}