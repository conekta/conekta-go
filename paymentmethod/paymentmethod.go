package paymentmethod

//DefaultParams should be a struct of the api response
type DefaultParams struct {
	ID        string         `json:"id"`
	Object    string         `json:"object"`
	Type      string         `json:"type"`
	CreatedAt int            `json:"created_at"`
	Last4     string         `json:"last4"`
	Bin       string         `json:"bin"`
	ExpMonth  string         `json:"exp_month"`
	ExpYear   string         `json:"exp_year"`
	Brand     string         `json:"brand"`
	Name      string         `json:"name"`
	ParentID  string         `json:"parent_id"`
	Default   bool           `json:"default"`
	Address   DefaultAddress `json:"address"`
}

//DefaultAddress should be nested struct of PaymentMethod
type DefaultAddress struct {
	Street1    string `json:"street1"`
	Street2    string `json:"street2"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	Object     string `json:"object"`
	PostalCode string `json:"postal_code"`
}
