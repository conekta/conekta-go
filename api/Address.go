package api

import (
)

type Address struct {
    Street1  string  `json:"street1,omitempty"`
    Street2  string  `json:"street2,omitempty"`
    Street3  string  `json:"street3,omitempty"`
    City  string  `json:"city,omitempty"`
    State  string  `json:"state,omitempty"`
    Zip  string  `json:"zip,omitempty"`
    Country  string  `json:"country,omitempty"`
    TaxId  string  `json:"tax_id,omitempty"`
    CompanyName  string  `json:"company_name,omitempty"`
    Phone  string  `json:"phone,omitempty"`
    Email  string  `json:"email,omitempty"`
    
}
