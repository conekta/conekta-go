package conekta

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os/exec"
	"reflect"
	"runtime"
)

const (
	// UnknownPlatform is the string returned as the system name if we couldn't get
	// one from `uname`.
	UnknownPlatform string = "unknown platform"
)

var encodedConektaUserAgent string

func init() {
	initUserAgent()
}

// RequestAPI returns Conekta API response
func RequestAPI(method string, url string, params ParamsConverter, customHeaders ...interface{}) ([]byte, error) {
	requestURL := BuildURL(url)

	client := &http.Client{}
	req, _ := http.NewRequest(method, requestURL, bytes.NewBuffer(params.Bytes()))

	setHeaders(req, customHeaders...)

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
func setHeaders(r *http.Request, customHeaders ...interface{}) *http.Request {
	r.Header.Set("Accept", "application/vnd.conekta-v"+apiVersion+"+json")
	r.Header.Set("Accept-Language", Locale)
	r.Header.Set("X-Conekta-Client-User-Agent", encodedConektaUserAgent)
	r.Header.Set("User-Agent", "Conekta/v2 GoBindings/"+version)
	r.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(APIKey)))
	r.Header.Set("Content-Type", "application/json")

	for _, v := range customHeaders {
		cht := reflect.TypeOf(v).String()
		if cht == "map[string]string" {
			keys := reflect.ValueOf(v).MapKeys()
			ch := v.(map[string]string)
			for _, k := range keys {
				r.Header.Set(k.String(), ch[k.String()])
			}
		}
	}
	return r
}

func initUserAgent() {
	data := map[string]string{
		"bindings_version": version,
		"lang":             "go",
		"lang_version":     runtime.Version(),
		"publisher":        "conekta",
		"uname":            getUname(),
	}

	r, _ := json.Marshal(data)
	encodedConektaUserAgent = string(r)
}

// getUname tries to get an uname from the system, but not that hard. It tries
// to execute `uname -a`, but swallows any errors in case that didn't work
// (i.e. non-Unix non-Mac system or some other reason).
func getUname() string {
	path, err := exec.LookPath("uname")
	if err != nil {
		return UnknownPlatform
	}

	cmd := exec.Command(path, "-a")
	var out bytes.Buffer
	cmd.Stderr = nil // goes to os.DevNull
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return UnknownPlatform
	}

	return out.String()
}

// MakeRequest creates a new Conekta request,
// it takes a method, endpoint, parameters and a reference struct of conekta's models like
// &conekta.Customer{} and fills it with the response
// if the response has a Conekta's error it returns it and keep empty the Conekta model
func MakeRequest(method string, endpoint string, p ParamsConverter, v interface{}, customHeaders ...interface{}) error {
	res, err := RequestAPI(method, endpoint, p, customHeaders...)
	if err != nil {
		return err
	}
	json.Unmarshal(res, v)
	return nil
}
