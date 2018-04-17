package api

import (
)

type PaymentMethod struct {
    Object  string  `json:"object,omitempty"`
    Name  string  `json:"name,omitempty"`
    ExpMonth  int32  `json:"exp_month,omitempty"`
    Clabe  string  `json:"clabe,omitempty"`
    ExpYear  int32  `json:"exp_year,omitempty"`
    AuthCode  string  `json:"auth_code,omitempty"`
    Last4  string  `json:"last4,omitempty"`
    Brand  string  `json:"brand,omitempty"`
    Type_  string  `json:"type,omitempty"`
    BarcodeUrl  string  `json:"barcode_url,omitempty"`
    Barcode  string  `json:"barcode,omitempty"`
    Address  Address  `json:"address,omitempty"`
    
}
