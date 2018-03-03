package conekta

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// RequestAPI returns Conekta API response
func RequestAPI(method string, url string, params ParamsConverter) ([]byte, *Error) {
	requestURL := BuildURL(url)

	client := &http.Client{}
	req, _ := http.NewRequest(method, requestURL, bytes.NewBuffer(params.Bytes()))

	setHeaders(req)
	//TODO: FIX ME make a more readable way to distinguish between a tokenize request from others
	if url == "/tokens" {
		setAdditionalHeaders(req)
	}

	res, err := client.Do(req)

	// Client errors
	if err != nil {
		return nil, getConectionError()
	}

	body, _ := ioutil.ReadAll(res.Body)

	// API errors
	if res.StatusCode != 200 {
		return nil, getAPIError(body)
	}

	return body, nil
}

// BuildURL returns base api plus endpoint passed
func BuildURL(url string) string {
	return apiBase + url
}

// setHeader set req object (htttp.Request) with conekta required headers to auth and identify request
func setHeaders(r *http.Request) *http.Request {
	r.Header.Set("Accept", "application/vnd.conekta-v"+apiVersion+"+json")
	r.Header.Set("Accept-Language", Locale)
	r.Header.Set("User-Agent", "Conekta/v1 GoBindings/"+version)
	r.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(APIKey)))
	r.Header.Set("Content-Type", "application/json")

	return r
}

// setAdditionalHeaders set req object (htttp.Request) with conekta required headers for CONEKTAJS
func setAdditionalHeaders(r *http.Request) *http.Request {
	r.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(PubliAPIKey)))
	r.Header.Set("Conekta-Client-User-Agent", "{\"agent\":\"Conekta JavascriptBindings/0.3.0\"}")
	r.Header.Set("RaiseHtmlError", "false")
	r.Header.Set("Accept", "application/vnd.conekta-v0.3.0+json")

	return r
}

// MakeRequest creates a new Conekta request,
// it takes a method, endpoint, parameters and a reference struct of conekta's models like
// &conekta.Customer{} and fills it with the response
// if the response has a Conekta's error it returns it and keep empty the Conekta model
func MakeRequest(method string, endpoint string, p ParamsConverter, v interface{}) *Error {
	res, err := RequestAPI(method, endpoint, p)
	if err != nil {
		return err
	}
	json.Unmarshal(res, v)
	return nil
}
