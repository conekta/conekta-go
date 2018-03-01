package conekta

//TokenParams should be a struct of tokenization response
type TokenParams struct {
	Number       string              `json:"number,omitempty"`
	Name         string              `json:"name,omitempty"`
	ExpYear      string              `json:"exp_year,omitempty"`
	ExpMonth     string              `json:"exp_month,omitempty"`
	Cvc          string              `json:"cvc,omitempty"`
	TokenAddress *TokenAddressParams `json:"address,omitempty"`
}

//TokenAddressParams token return address as submodel of TokenParams
type TokenAddressParams struct {
	Street1 string `json:"street1,omitempty"`
	Street2 string `json:"street2,omitempty"`
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	Zip     string `json:"zip,omitempty"`
	Country string `json:"country,omitempty"`
}

//Token is used when user tries to tokenize
type Token struct {
	ID       string `json:"id,omitempty"`
	Object   string `json:"object,omitempty"`
	Used     bool   `json:"used,omitempty"`
	Livemode bool   `json:"livemode,omitempty"`
}

//TokenList is a list of Token's
type TokenList struct {
	ListMeta
	Data []*Token `json:"data,omitempty"`
}

// Bytes converts a ChargeParams to []byte
func (c *TokenParams) Bytes() []byte {
	return paramsToBytes(c)
}
