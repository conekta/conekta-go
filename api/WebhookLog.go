package api

import (
)

type WebhookLog struct {
    Id  string  `json:"id,omitempty"`
    Url  string  `json:"url,omitempty"`
    FailedAttempts  string  `json:"failed_attempts,omitempty"`
    LastHttpResponseStatus  string  `json:"last_http_response_status,omitempty"`
    Object  string  `json:"object,omitempty"`
    LastAttemptedAt  string  `json:"last_attempted_at,omitempty"`
    
}
