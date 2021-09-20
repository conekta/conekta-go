package conekta

type Plan struct {
	ID              string `json:"id,omitempty"`
	Livemode        bool   `json:"livemode,omitempty"`
	CreateAt        int64  `json:"create_at,omitempty"`
	Name            string `json:"name,omitempty"`
	Amount          int64  `json:"amount,omitempty"`
	Currency        int64  `json:"currency,omitempty"`
	Interval        string `json:"interval,omitempty"`
	Frequency       int64  `json:"frequency,omitempty"`
	ExpiryCount     int64  `json:"expiry_count,omitempty"`
	TrialPeriodDays int64  `json:"trial_period_days,omitempty"`
	Object          string `json:"object,omitempty"`
}

type PlanParams struct {
	ID              string `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Currency        string `json:"currency,omitempty"`
	Amount          int64  `json:"amount,omitempty"`
	Interval        string `json:"interval,omitempty"`
	Frequency       int64  `json:"frequency,omitempty"`
	TrialPeriodDays int64  `json:"trial_period_days,omitempty"`
	ExpiryCount     int64  `json:"expiry_count,omitempty"`
}

type PlanList struct {
	ListMeta
	Data []*Plan `json:"data,omitempty"`
}

func (c *PlanParams) Bytes() []byte {
	return paramsToBytes(c)
}
