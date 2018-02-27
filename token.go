package conekta

//TokenParams should be a struct of tokenization response
type TokenParams struct {
	Number       string             `json:"number"`
	Name         string             `json:"name"`
	ExpYear      string             `json:"exp_year"`
	ExpMonth     string             `json:"exp_month"`
	Cvc          string             `json:"cvc"`
	TokenAddress TokenAddressParams `json:"address"`
}

//TokenAddressParams token return address as submodel of TokenParams
type TokenAddressParams struct {
	Street1 string `json:"street1"`
	Street2 string `json:"street2"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
	Country string `json:"country"`
}

//
type Token struct {
	ID       string `json:"id"`
	Object   string `json:"object"`
	Used     bool   `json:"used"`
	Livemode bool   `json:"livemode"`
}

// Bytes converts a ChargeParams to []byte
func (c *TokenParams) Bytes() []byte {
	return paramsToBytes(c)
}
