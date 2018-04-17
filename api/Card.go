package api

import (
)

type Card struct {
    Id  string  `json:"id,omitempty"`
    Object  string  `json:"object,omitempty"`
    CreatedAt  string  `json:"created_at,omitempty"`
    Name  string  `json:"name,omitempty"`
    Last4  string  `json:"last4,omitempty"`
    ExpMonth  string  `json:"exp_month,omitempty"`
    ExpYear  string  `json:"exp_year,omitempty"`
    Active  string  `json:"active,omitempty"`
    Address  Address  `json:"address,omitempty"`
    
}
