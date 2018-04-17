package api

import (
)

type Details struct {
    Name  string  `json:"name,omitempty"`
    Phone  string  `json:"phone,omitempty"`
    Email  string  `json:"email,omitempty"`
    Customer  ChargeCustomerInfo  `json:"customer,omitempty"`
    LineItems  []LineItem  `json:"line_items,omitempty"`
    BillingAddress  Address  `json:"billing_address,omitempty"`
    
}
