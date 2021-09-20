package conekta

type Subscription struct {
	ID                string `json:"id,omitempty"`
	Status            string `json:"status,omitempty"`
	Object            string `json:"object,omitempty"`
	ChargeID          string `json:"charge_id,omitempty"`
	CreateAt          string `json:"create_at,omitempty"`
	SubscriptionStart int64  `json:"subscription_start,omitempty"`
	BillingCycleStart int64  `json:"billing_cycle_start,omitempty"`
	BillingCycleEnd   int64  `json:"billing_cycle_end,omitempty"`
	PlanID            string `json:"plan_id,omitempty"`
	CustomerID        string `json:"customer_id,omitempty"`
	CardID            string `json:"card_id,omitempty"`
}

type SubscriptionParams struct {
	CardID     string `json:"card,omitempty"`
	CustomerID string `json:"customer_id,omitempty"`
	PlanID     string `json:"plan,omitempty"`
}

type SubscriptionList struct {
	ListMeta
	Data []*Subscription `json:"data,omitempty"`
}

func (p *SubscriptionParams) Bytes() []byte {
	return paramsToBytes(p)
}

func (p *Subscription) Bytes() []byte {
	return paramsToBytes(p)
}
