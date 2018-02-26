package conekta

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
)

// RequestAPI returns conekta api response
func RequestAPI(method string, url string) []byte {
	params := GetParams()

	requestURL := BuildURL(url, params)

	client := &http.Client{}
	req, _ := http.NewRequest(method, requestURL, nil)

	setHeaders(req, params)

	res, _ := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	return body

}

// BuildURL returns base api plus endpoint passed
func BuildURL(url string, p *conektaParams) string {
	return p.APIBase + url
}

// setHeader set req object (htttp.Request) with conekta required headers to auth and identify request
func setHeaders(r *http.Request, p *conektaParams) *http.Request {
	r.Header.Set("Accept", "application/vnd.conekta-v"+p.APIVersion+"+json")
	r.Header.Set("Accept-Language", p.Locale)
	r.Header.Set("User-Agent", "Conekta/v1 GoBindings/"+p.Version)
	r.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(p.APIKey)))
	r.Header.Set("Content-Type", "application/json")
	return r
}
