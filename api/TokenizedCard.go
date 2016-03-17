package api

import (
)

type TokenizedCard struct {
    Id  string  `json:"id,omitempty"`
    Name  string  `json:"name,omitempty"`
    Last4  string  `json:"last4,omitempty"`
    ExpMonth  string  `json:"exp_month,omitempty"`
    ExpYear  string  `json:"exp_year,omitempty"`
    Active  string  `json:"active,omitempty"`
    
}
