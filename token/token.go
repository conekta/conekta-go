package token

//DefaultParams should be a struct of tokenization response
type DefaultParams struct {
	Number       string         `json:"number"`
	Name         string         `json:"name"`
	ExpYear      string         `json:"exp_year"`
	ExpMonth     string         `json:"exp_month"`
	Cvc          string         `json:"cvc"`
	TokenAddress DefaultAddress `json:"address"`
}

//DefaultAddress token return address as submodel of TokenParams
type DefaultAddress struct {
	Street1 string `json:"street1"`
	Street2 string `json:"street2"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
	Country string `json:"country"`
}
