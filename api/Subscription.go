package api

import (
)

type Subscription struct {
    Id  string  `json:"id,omitempty"`
    Card  string  `json:"card,omitempty"`
    PlanId  string  `json:"plan_id,omitempty"`
    Status  string  `json:"status,omitempty"`
    Start  string  `json:"start,omitempty"`
    BillingCycleStart  string  `json:"billing_cycle_start,omitempty"`
    BillingCycleEnd  string  `json:"billing_cycle_end,omitempty"`
    
}


