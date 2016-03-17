package api

import (
)

type ChargeCustomerInfo struct {
    LoggedIn  string  `json:"logged_in,omitempty"`
    SuccessfulPurchases  int32  `json:"successful_purchases,omitempty"`
    CreatedAt  int32  `json:"created_at,omitempty"`
    UpdatedAt  int32  `json:"updated_at,omitempty"`
    OfflinePayments  int32  `json:"offline_payments,omitempty"`
    Score  int32  `json:"score,omitempty"`
    
}
