package token

import (
	"encoding/base64"

	conekta "github.com/conekta/conekta-go/v2"
)

//Create create discount line insde a order
// For details see https://developers.conekta.com/api#create-discount-line
func Create(p *conekta.TokenParams) (*conekta.Token, error) {
	tk := &conekta.Token{}
	h := map[string]string{
		"Authorization":             "Basic " + base64.StdEncoding.EncodeToString([]byte(conekta.PubliAPIKey)),
		"Conekta-Client-User-Agent": "{\"agent\":\"Conekta JavascriptBindings/0.3.0\"}",
		"RaiseHtmlError":            "false",
		"Accept":                    "application/vnd.conekta-v0.3.0+json",
	}
	conekta.MakeRequest("POST", "/tokens", p, tk, h)
	if tk.Object == "error" {
		return nil, &conekta.Error{
			ErrorType: tk.Type,
			Details: []conekta.ErrorDetails{
				{
					Message:            tk.Message,
					MessageToPurchaser: tk.MessageToPurchaser,
					Param:              tk.Param,
					Code:               tk.Code,
				},
			},
		}
	}
	return tk, nil
}
