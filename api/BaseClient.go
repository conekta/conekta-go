package api

import (
)

type BaseClient struct {
    Plan  string  `json:"plan,omitempty"`
    Name  string  `json:"name,omitempty"`
    Email  string  `json:"email,omitempty"`
    Phone  string  `json:"phone,omitempty"`
    Cards  []string  `json:"cards,omitempty"`
    
}
