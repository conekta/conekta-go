package api

type conekta struct {
	O interface{};
	apiKey string
}

var instantiated *conekta = nil

func Conekta() *conekta {
	if instantiated == nil {
		instantiated = new(conekta);
	}
	return instantiated;
}

func (c *conekta) SetApikey(apiKey string){
	c.apiKey = apiKey
}