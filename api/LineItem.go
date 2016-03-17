package api

import (
)

type LineItem struct {
    Name  string  `json:"name,omitempty"`
    Description  string  `json:"description,omitempty"`
    UnitPrice  int32  `json:"unit_price,omitempty"`
    Quantity  int32  `json:"quantity,omitempty"`
    Sku  string  `json:"sku,omitempty"`
    Category  string  `json:"category,omitempty"`
    
}
