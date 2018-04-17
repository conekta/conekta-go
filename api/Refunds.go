package api

import (
)

type Refunds struct {
    CreatedAt  string  `json:"created_at,omitempty"`
    Amount  int32  `json:"amount,omitempty"`
    Currency  string  `json:"currency,omitempty"`
    Transaction  string  `json:"transaction,omitempty"`
    
}
