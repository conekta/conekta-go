package api

import (
        "strings"
        "fmt"
        "encoding/json"
        "errors"
        "encoding/base64"
        "github.com/dghubble/sling"
        "net/url"
)

type DefaultApi struct {
        basePath  string
        apiKey string
}

func NewDefaultApi(apiKey string) *DefaultApi{
        return &DefaultApi {
                basePath:   "https://api.conekta.io",
                apiKey: apiKey,
        }
}

func (a DefaultApi) NewWithAuth() *sling.Sling {
  return sling.New().Set("Authorization", "Basic "+ base64.StdEncoding.EncodeToString([]byte(a.apiKey)))
}

func NewDefaultApiWithBasePath(basePath string) *DefaultApi{
    return &DefaultApi {
        basePath:   basePath,
    }
}

/**
 * 
 * captures an payment
 * @param chargeId 
 * @return Charge
 */
//func (a DefaultApi) ChargesChargeIdCapturePost (chargeId string) (Charge, error) {
func (a DefaultApi) ChargesChargeIdCapturePost (chargeId string) (Charge, error) {

    _sling := a.NewWithAuth().Post(a.basePath)

    // create path and map variables
    path := "/charges/{charge_id}/capture"
    path = strings.Replace(path, "{" + "charge_id" + "}", fmt.Sprintf("%v", chargeId), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string {  }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }


  var successPayload = new(Charge)

  // We use this map (below) so that any arbitrary error JSON can be handled.
  // FIXME: This is in the absence of this Go generator honoring the non-2xx
  // response (error) models, which needs to be implemented at some point.
  var failurePayload map[string]interface{}

  httpResponse, err := _sling.Receive(successPayload, &failurePayload)

  if err == nil {
    // err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
    if failurePayload != nil {
      // If the failurePayload is present, there likely was some kind of non-2xx status
      // returned (and a JSON payload error present)
      var str []byte
      str, err = json.Marshal(failurePayload)
      if err == nil { // For safety, check for an error marshalling... probably superfluous
        // This will return the JSON error body as a string
        err = errors.New(string(str))
      }
  } else {
    // So, there was no network-type error, and nothing in the failure payload,
    // but we should still check the status code
    if httpResponse == nil {
      // This should never happen...
      err = errors.New("No HTTP Response received.")
    } else if code := httpResponse.StatusCode; 200 > code || code > 299 {
        err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
      }
    }
  }

  return *successPayload, err
}
/**
 * 
 * 
 * @param chargeId 
 * @return Charge
 */
//func (a DefaultApi) ChargesChargeIdGet (chargeId string) (Charge, error) {
func (a DefaultApi) ChargesChargeIdGet (chargeId string) (Charge, error) {

    _sling := a.NewWithAuth().Get(a.basePath)

    // create path and map variables
    path := "/charges/{charge_id}"
    path = strings.Replace(path, "{" + "charge_id" + "}", fmt.Sprintf("%v", chargeId), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string {  }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }


  var successPayload = new(Charge)

  // We use this map (below) so that any arbitrary error JSON can be handled.
  // FIXME: This is in the absence of this Go generator honoring the non-2xx
  // response (error) models, which needs to be implemented at some point.
  var failurePayload map[string]interface{}

  httpResponse, err := _sling.Receive(successPayload, &failurePayload)

  if err == nil {
    // err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
    if failurePayload != nil {
      // If the failurePayload is present, there likely was some kind of non-2xx status
      // returned (and a JSON payload error present)
      var str []byte
      str, err = json.Marshal(failurePayload)
      if err == nil { // For safety, check for an error marshalling... probably superfluous
        // This will return the JSON error body as a string
        err = errors.New(string(str))
      }
  } else {
    // So, there was no network-type error, and nothing in the failure payload,
    // but we should still check the status code
    if httpResponse == nil {
      // This should never happen...
      err = errors.New("No HTTP Response received.")
    } else if code := httpResponse.StatusCode; 200 > code || code > 299 {
        err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
      }
    }
  }

  return *successPayload, err
}
/**
 * 
 * 
 * @param chargeId 
 * @param amountQty 
 * @return Charge
 */
//func (a DefaultApi) ChargesChargeIdRefundPost (chargeId string, amountQty AmountQty) (Charge, error) {
func (a DefaultApi) ChargesChargeIdRefundPost (chargeId string, amountQty AmountQty) (Charge, error) {

    _sling := a.NewWithAuth().Post(a.basePath)

    // create path and map variables
    path := "/charges/{charge_id}/refund"
    path = strings.Replace(path, "{" + "charge_id" + "}", fmt.Sprintf("%v", chargeId), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string {  }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }

// body params
    _sling = _sling.BodyJSON(amountQty)

  var successPayload = new(Charge)

  // We use this map (below) so that any arbitrary error JSON can be handled.
  // FIXME: This is in the absence of this Go generator honoring the non-2xx
  // response (error) models, which needs to be implemented at some point.
  var failurePayload map[string]interface{}

  httpResponse, err := _sling.Receive(successPayload, &failurePayload)

  if err == nil {
    // err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
    if failurePayload != nil {
      // If the failurePayload is present, there likely was some kind of non-2xx status
      // returned (and a JSON payload error present)
      var str []byte
      str, err = json.Marshal(failurePayload)
      if err == nil { // For safety, check for an error marshalling... probably superfluous
        // This will return the JSON error body as a string
        err = errors.New(string(str))
      }
  } else {
    // So, there was no network-type error, and nothing in the failure payload,
    // but we should still check the status code
    if httpResponse == nil {
      // This should never happen...
      err = errors.New("No HTTP Response received.")
    } else if code := httpResponse.StatusCode; 200 > code || code > 299 {
        err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
      }
    }
  }

  return *successPayload, err
}


/**
 *
 *
 * @return []Charge
 */
//func (a DefaultApi) ChargesGet () ([]Charge, error) {
func (a DefaultApi) ChargesGetArgs (args map[string]string) ([]Charge, error) {

        _sling := a.NewWithAuth().Get(a.basePath)

        // create path and map variables
        path := "/charges"

        _sling = _sling.Path(path)

        // accept header
        accepts := []string {  }
        for key := range accepts {
                _sling = _sling.Set("Accept", accepts[key])
                break // only use the first Accept
        }

        var successPayload = new([]Charge)

        // We use this map (below) so that any arbitrary error JSON can be handled.
        // FIXME: This is in the absence of this Go generator honoring the non-2xx
        // response (error) models, which needs to be implemented at some point.
        var failurePayload map[string]interface{}

        baseUrl := ""
        parsedUrl, _ := url.ParseQuery(baseUrl)

        for key, value := range args{
                parsedUrl.Add(key, value)
        }

        _sling.Base("https://api.conekta.io/charges?"+parsedUrl.Encode());

        fmt.Printf("%+v\n", _sling)

        httpResponse, err := _sling.Receive(successPayload, &failurePayload)

        if err == nil {
                // err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
                if failurePayload != nil {
                        // If the failurePayload is present, there likely was some kind of non-2xx status
                        // returned (and a JSON payload error present)
                        var str []byte
                        str, err = json.Marshal(failurePayload)
                        if err == nil { // For safety, check for an error marshalling... probably superfluous
                                // This will return the JSON error body as a string
                                err = errors.New(string(str))
                        }
                } else {
                        // So, there was no network-type error, and nothing in the failure payload,
                        // but we should still check the status code
                        if httpResponse == nil {
                                // This should never happen...
                                err = errors.New("No HTTP Response received.")
                        } else if code := httpResponse.StatusCode; 200 > code || code > 299 {
                                err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
                        }
                }
        }

        return *successPayload, err
}

/**
 * 
 * 
 * @return []Charge
 */
//func (a DefaultApi) ChargesGet () ([]Charge, error) {
func (a DefaultApi) ChargesGet () ([]Charge, error) {

    _sling := a.NewWithAuth().Get(a.basePath)

    // create path and map variables
    path := "/charges"

    _sling = _sling.Path(path)

    // accept header
    accepts := []string {  }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }


  var successPayload = new([]Charge)

  // We use this map (below) so that any arbitrary error JSON can be handled.
  // FIXME: This is in the absence of this Go generator honoring the non-2xx
  // response (error) models, which needs to be implemented at some point.
  var failurePayload map[string]interface{}

  httpResponse, err := _sling.Receive(successPayload, &failurePayload)

  if err == nil {
    // err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
    if failurePayload != nil {
      // If the failurePayload is present, there likely was some kind of non-2xx status
      // returned (and a JSON payload error present)
      var str []byte
      str, err = json.Marshal(failurePayload)
      if err == nil { // For safety, check for an error marshalling... probably superfluous
        // This will return the JSON error body as a string
        err = errors.New(string(str))
      }
  } else {
    // So, there was no network-type error, and nothing in the failure payload,
    // but we should still check the status code
    if httpResponse == nil {
      // This should never happen...
      err = errors.New("No HTTP Response received.")
    } else if code := httpResponse.StatusCode; 200 > code || code > 299 {
        err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
      }
    }
  }

  return *successPayload, err
}
/**
 * Create a Charge
 * 
 * @param body Charge object
 * @return Charge
 */
//func (a DefaultApi) ChargesPost (body Charge) (Charge, error) {
func (a DefaultApi) ChargesPost (body Charge) (Charge, error) {

    _sling := a.NewWithAuth().Post(a.basePath)

    // create path and map variables
    path := "/charges"

    _sling = _sling.Path(path)

    // accept header
    accepts := []string { "application/json" }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }

// body params
    _sling = _sling.BodyJSON(body)

  var successPayload = new(Charge)

  // We use this map (below) so that any arbitrary error JSON can be handled.
  // FIXME: This is in the absence of this Go generator honoring the non-2xx
  // response (error) models, which needs to be implemented at some point.
  var failurePayload map[string]interface{}

  httpResponse, err := _sling.Receive(successPayload, &failurePayload)

  if err == nil {
    // err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
    if failurePayload != nil {
      // If the failurePayload is present, there likely was some kind of non-2xx status
      // returned (and a JSON payload error present)
      var str []byte
      str, err = json.Marshal(failurePayload)
      if err == nil { // For safety, check for an error marshalling... probably superfluous
        // This will return the JSON error body as a string
        err = errors.New(string(str))
      }
  } else {
    // So, there was no network-type error, and nothing in the failure payload,
    // but we should still check the status code
    if httpResponse == nil {
      // This should never happen...
      err = errors.New("No HTTP Response received.")
    } else if code := httpResponse.StatusCode; 200 > code || code > 299 {
        err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
      }
    }
  }

  return *successPayload, err
}
/**
 * 
 * captures an payment
 * @param customerId 
 * @param cardId 
 * @return Card
 */
//func (a DefaultApi) CustomersCustomerIdCardsCardIdDelete (customerId string, cardId string) (Card, error) {
func (a DefaultApi) CustomersCustomerIdCardsCardIdDelete (customerId string, cardId string) (Card, error) {

    _sling := a.NewWithAuth().Delete(a.basePath)

    // create path and map variables
    path := "/customers/{customer_id}/cards/{card_id}"
    path = strings.Replace(path, "{" + "customer_id" + "}", fmt.Sprintf("%v", customerId), -1)
    path = strings.Replace(path, "{" + "card_id" + "}", fmt.Sprintf("%v", cardId), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string {  }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }


  var successPayload = new(Card)

  // We use this map (below) so that any arbitrary error JSON can be handled.
  // FIXME: This is in the absence of this Go generator honoring the non-2xx
  // response (error) models, which needs to be implemented at some point.
  var failurePayload map[string]interface{}

  httpResponse, err := _sling.Receive(successPayload, &failurePayload)

  if err == nil {
    // err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
    if failurePayload != nil {
      // If the failurePayload is present, there likely was some kind of non-2xx status
      // returned (and a JSON payload error present)
      var str []byte
      str, err = json.Marshal(failurePayload)
      if err == nil { // For safety, check for an error marshalling... probably superfluous
        // This will return the JSON error body as a string
        err = errors.New(string(str))
      }
  } else {
    // So, there was no network-type error, and nothing in the failure payload,
    // but we should still check the status code
    if httpResponse == nil {
      // This should never happen...
      err = errors.New("No HTTP Response received.")
    } else if code := httpResponse.StatusCode; 200 > code || code > 299 {
        err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
      }
    }
  }

  return *successPayload, err
}
/**
 * 
 * captures an payment
 * @param customerId 
 * @param cardId 
 * @param client 
 * @return Card
 */
//func (a DefaultApi) CustomersCustomerIdCardsCardIdPut (customerId string, cardId string, client TokenObject) (Card, error) {
func (a DefaultApi) CustomersCustomerIdCardsCardIdPut (customerId string, cardId string, client TokenObject) (Card, error) {

    _sling := a.NewWithAuth().Put(a.basePath)

    // create path and map variables
    path := "/customers/{customer_id}/cards/{card_id}"
    path = strings.Replace(path, "{" + "customer_id" + "}", fmt.Sprintf("%v", customerId), -1)
    path = strings.Replace(path, "{" + "card_id" + "}", fmt.Sprintf("%v", cardId), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string {  }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }

// body params
    _sling = _sling.BodyJSON(client)

  var successPayload = new(Card)

  // We use this map (below) so that any arbitrary error JSON can be handled.
  // FIXME: This is in the absence of this Go generator honoring the non-2xx
  // response (error) models, which needs to be implemented at some point.
  var failurePayload map[string]interface{}

  httpResponse, err := _sling.Receive(successPayload, &failurePayload)

  if err == nil {
    // err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
    if failurePayload != nil {
      // If the failurePayload is present, there likely was some kind of non-2xx status
      // returned (and a JSON payload error present)
      var str []byte
      str, err = json.Marshal(failurePayload)
      if err == nil { // For safety, check for an error marshalling... probably superfluous
        // This will return the JSON error body as a string
        err = errors.New(string(str))
      }
  } else {
    // So, there was no network-type error, and nothing in the failure payload,
    // but we should still check the status code
    if httpResponse == nil {
      // This should never happen...
      err = errors.New("No HTTP Response received.")
    } else if code := httpResponse.StatusCode; 200 > code || code > 299 {
        err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
      }
    }
  }

  return *successPayload, err
}
/**
 * 
 * captures an payment
 * @param customerId 
 * @param client 
 * @return Card
 */
//func (a DefaultApi) CustomersCustomerIdCardsPost (customerId string, client TokenObject) (Card, error) {
func (a DefaultApi) CustomersCustomerIdCardsPost (customerId string, client TokenObject) (Card, error) {

    _sling := a.NewWithAuth().Post(a.basePath)

    // create path and map variables
    path := "/customers/{customer_id}/cards"
    path = strings.Replace(path, "{" + "customer_id" + "}", fmt.Sprintf("%v", customerId), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string {  }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }

// body params
    _sling = _sling.BodyJSON(client)

  var successPayload = new(Card)

  // We use this map (below) so that any arbitrary error JSON can be handled.
  // FIXME: This is in the absence of this Go generator honoring the non-2xx
  // response (error) models, which needs to be implemented at some point.
  var failurePayload map[string]interface{}

  httpResponse, err := _sling.Receive(successPayload, &failurePayload)

  if err == nil {
    // err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
    if failurePayload != nil {
      // If the failurePayload is present, there likely was some kind of non-2xx status
      // returned (and a JSON payload error present)
      var str []byte
      str, err = json.Marshal(failurePayload)
      if err == nil { // For safety, check for an error marshalling... probably superfluous
        // This will return the JSON error body as a string
        err = errors.New(string(str))
      }
  } else {
    // So, there was no network-type error, and nothing in the failure payload,
    // but we should still check the status code
    if httpResponse == nil {
      // This should never happen...
      err = errors.New("No HTTP Response received.")
    } else if code := httpResponse.StatusCode; 200 > code || code > 299 {
        err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
      }
    }
  }

  return *successPayload, err
}
/**
 * 
 * captures an payment
 * @param customerId 
 * @return Customer
 */
//func (a DefaultApi) CustomersCustomerIdDelete (customerId string) (Customer, error) {
func (a DefaultApi) CustomersCustomerIdDelete (customerId string) (Customer, error) {

    _sling := a.NewWithAuth().Delete(a.basePath)

    // create path and map variables
    path := "/customers/{customer_id}"
    path = strings.Replace(path, "{" + "customer_id" + "}", fmt.Sprintf("%v", customerId), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string {  }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }


  var successPayload = new(Customer)

  // We use this map (below) so that any arbitrary error JSON can be handled.
  // FIXME: This is in the absence of this Go generator honoring the non-2xx
  // response (error) models, which needs to be implemented at some point.
  var failurePayload map[string]interface{}

  httpResponse, err := _sling.Receive(successPayload, &failurePayload)

  if err == nil {
    // err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
    if failurePayload != nil {
      // If the failurePayload is present, there likely was some kind of non-2xx status
      // returned (and a JSON payload error present)
      var str []byte
      str, err = json.Marshal(failurePayload)
      if err == nil { // For safety, check for an error marshalling... probably superfluous
        // This will return the JSON error body as a string
        err = errors.New(string(str))
      }
  } else {
    // So, there was no network-type error, and nothing in the failure payload,
    // but we should still check the status code
    if httpResponse == nil {
      // This should never happen...
      err = errors.New("No HTTP Response received.")
    } else if code := httpResponse.StatusCode; 200 > code || code > 299 {
        err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
      }
    }
  }

  return *successPayload, err
}
/**
 * 
 * captures an payment
 * @param customerId 
 * @return Customer
 */
//func (a DefaultApi) CustomersCustomerIdGet (customerId string) (Customer, error) {
func (a DefaultApi) CustomersCustomerIdGet (customerId string) (Customer, error) {

    _sling := a.NewWithAuth().Get(a.basePath)

    // create path and map variables
    path := "/customers/{customer_id}"
    path = strings.Replace(path, "{" + "customer_id" + "}", fmt.Sprintf("%v", customerId), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string {  }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }


  var successPayload = new(Customer)

  // We use this map (below) so that any arbitrary error JSON can be handled.
  // FIXME: This is in the absence of this Go generator honoring the non-2xx
  // response (error) models, which needs to be implemented at some point.
  var failurePayload map[string]interface{}

  httpResponse, err := _sling.Receive(successPayload, &failurePayload)

  if err == nil {
    // err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
    if failurePayload != nil {
      // If the failurePayload is present, there likely was some kind of non-2xx status
      // returned (and a JSON payload error present)
      var str []byte
      str, err = json.Marshal(failurePayload)
      if err == nil { // For safety, check for an error marshalling... probably superfluous
        // This will return the JSON error body as a string
        err = errors.New(string(str))
      }
  } else {
    // So, there was no network-type error, and nothing in the failure payload,
    // but we should still check the status code
    if httpResponse == nil {
      // This should never happen...
      err = errors.New("No HTTP Response received.")
    } else if code := httpResponse.StatusCode; 200 > code || code > 299 {
        err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
      }
    }
  }

  return *successPayload, err
}
/**
 * 
 * captures an payment
 * @param customerId 
 * @param client 
 * @return Customer
 */
//func (a DefaultApi) CustomersCustomerIdPut (customerId string, client BaseClient) (Customer, error) {
func (a DefaultApi) CustomersCustomerIdPut (customerId string, client BaseClient) (Customer, error) {

    _sling := a.NewWithAuth().Put(a.basePath)

    // create path and map variables
    path := "/customers/{customer_id}"
    path = strings.Replace(path, "{" + "customer_id" + "}", fmt.Sprintf("%v", customerId), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string {  }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }

// body params
    _sling = _sling.BodyJSON(client)

  var successPayload = new(Customer)

  // We use this map (below) so that any arbitrary error JSON can be handled.
  // FIXME: This is in the absence of this Go generator honoring the non-2xx
  // response (error) models, which needs to be implemented at some point.
  var failurePayload map[string]interface{}

  httpResponse, err := _sling.Receive(successPayload, &failurePayload)

  if err == nil {
    // err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
    if failurePayload != nil {
      // If the failurePayload is present, there likely was some kind of non-2xx status
      // returned (and a JSON payload error present)
      var str []byte
      str, err = json.Marshal(failurePayload)
      if err == nil { // For safety, check for an error marshalling... probably superfluous
        // This will return the JSON error body as a string
        err = errors.New(string(str))
      }
  } else {
    // So, there was no network-type error, and nothing in the failure payload,
    // but we should still check the status code
    if httpResponse == nil {
      // This should never happen...
      err = errors.New("No HTTP Response received.")
    } else if code := httpResponse.StatusCode; 200 > code || code > 299 {
        err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
      }
    }
  }

  return *successPayload, err
}
/**
 * 
 * captures an payment
 * @param customerId 
 * @return Subscription
 */
//func (a DefaultApi) CustomersCustomerIdSubscriptionCancelPost (customerId string) (Subscription, error) {
func (a DefaultApi) CustomersCustomerIdSubscriptionCancelPost (customerId string) (Subscription, error) {

    _sling := a.NewWithAuth().Post(a.basePath)

    // create path and map variables
    path := "/customers/{customer_id}/subscription/cancel"
    path = strings.Replace(path, "{" + "customer_id" + "}", fmt.Sprintf("%v", customerId), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string {  }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }


  var successPayload = new(Subscription)

  // We use this map (below) so that any arbitrary error JSON can be handled.
  // FIXME: This is in the absence of this Go generator honoring the non-2xx
  // response (error) models, which needs to be implemented at some point.
  var failurePayload map[string]interface{}

  httpResponse, err := _sling.Receive(successPayload, &failurePayload)

  if err == nil {
    // err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
    if failurePayload != nil {
      // If the failurePayload is present, there likely was some kind of non-2xx status
      // returned (and a JSON payload error present)
      var str []byte
      str, err = json.Marshal(failurePayload)
      if err == nil { // For safety, check for an error marshalling... probably superfluous
        // This will return the JSON error body as a string
        err = errors.New(string(str))
      }
  } else {
    // So, there was no network-type error, and nothing in the failure payload,
    // but we should still check the status code
    if httpResponse == nil {
      // This should never happen...
      err = errors.New("No HTTP Response received.")
    } else if code := httpResponse.StatusCode; 200 > code || code > 299 {
        err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
      }
    }
  }

  return *successPayload, err
}
/**
 * 
 * captures an payment
 * @param customerId 
 * @return Subscription
 */
//func (a DefaultApi) CustomersCustomerIdSubscriptionPausePost (customerId string) (Subscription, error) {
func (a DefaultApi) CustomersCustomerIdSubscriptionPausePost (customerId string) (Subscription, error) {

    _sling := a.NewWithAuth().Post(a.basePath)

    // create path and map variables
    path := "/customers/{customer_id}/subscription/pause"
    path = strings.Replace(path, "{" + "customer_id" + "}", fmt.Sprintf("%v", customerId), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string {  }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }


  var successPayload = new(Subscription)

  // We use this map (below) so that any arbitrary error JSON can be handled.
  // FIXME: This is in the absence of this Go generator honoring the non-2xx
  // response (error) models, which needs to be implemented at some point.
  var failurePayload map[string]interface{}

  httpResponse, err := _sling.Receive(successPayload, &failurePayload)

  if err == nil {
    // err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
    if failurePayload != nil {
      // If the failurePayload is present, there likely was some kind of non-2xx status
      // returned (and a JSON payload error present)
      var str []byte
      str, err = json.Marshal(failurePayload)
      if err == nil { // For safety, check for an error marshalling... probably superfluous
        // This will return the JSON error body as a string
        err = errors.New(string(str))
      }
  } else {
    // So, there was no network-type error, and nothing in the failure payload,
    // but we should still check the status code
    if httpResponse == nil {
      // This should never happen...
      err = errors.New("No HTTP Response received.")
    } else if code := httpResponse.StatusCode; 200 > code || code > 299 {
        err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
      }
    }
  }

  return *successPayload, err
}
/**
 * 
 * captures an payment
 * @param customerId 
 * @param subscription 
 * @return Subscription
 */
//func (a DefaultApi) CustomersCustomerIdSubscriptionPost (customerId string, subscription SubscriptionReference) (Subscription, error) {
func (a DefaultApi) CustomersCustomerIdSubscriptionPost (customerId string, subscription SubscriptionReference) (Subscription, error) {

    _sling := a.NewWithAuth().Post(a.basePath)

    // create path and map variables
    path := "/customers/{customer_id}/subscription"
    path = strings.Replace(path, "{" + "customer_id" + "}", fmt.Sprintf("%v", customerId), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string {  }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }

// body params
    _sling = _sling.BodyJSON(subscription)

  var successPayload = new(Subscription)

  // We use this map (below) so that any arbitrary error JSON can be handled.
  // FIXME: This is in the absence of this Go generator honoring the non-2xx
  // response (error) models, which needs to be implemented at some point.
  var failurePayload map[string]interface{}

  httpResponse, err := _sling.Receive(successPayload, &failurePayload)

  if err == nil {
    // err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
    if failurePayload != nil {
      // If the failurePayload is present, there likely was some kind of non-2xx status
      // returned (and a JSON payload error present)
      var str []byte
      str, err = json.Marshal(failurePayload)
      if err == nil { // For safety, check for an error marshalling... probably superfluous
        // This will return the JSON error body as a string
        err = errors.New(string(str))
      }
  } else {
    // So, there was no network-type error, and nothing in the failure payload,
    // but we should still check the status code
    if httpResponse == nil {
      // This should never happen...
      err = errors.New("No HTTP Response received.")
    } else if code := httpResponse.StatusCode; 200 > code || code > 299 {
        err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
      }
    }
  }

  return *successPayload, err
}
/**
 * 
 * captures an payment
 * @param customerId 
 * @param subscription 
 * @return Subscription
 */
//func (a DefaultApi) CustomersCustomerIdSubscriptionPut (customerId string, subscription SubscriptionReference) (Subscription, error) {
func (a DefaultApi) CustomersCustomerIdSubscriptionPut (customerId string, subscription SubscriptionReference) (Subscription, error) {

    _sling := a.NewWithAuth().Put(a.basePath)

    // create path and map variables
    path := "/customers/{customer_id}/subscription"
    path = strings.Replace(path, "{" + "customer_id" + "}", fmt.Sprintf("%v", customerId), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string {  }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }

// body params
    _sling = _sling.BodyJSON(subscription)

  var successPayload = new(Subscription)

  // We use this map (below) so that any arbitrary error JSON can be handled.
  // FIXME: This is in the absence of this Go generator honoring the non-2xx
  // response (error) models, which needs to be implemented at some point.
  var failurePayload map[string]interface{}

  httpResponse, err := _sling.Receive(successPayload, &failurePayload)

  if err == nil {
    // err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
    if failurePayload != nil {
      // If the failurePayload is present, there likely was some kind of non-2xx status
      // returned (and a JSON payload error present)
      var str []byte
      str, err = json.Marshal(failurePayload)
      if err == nil { // For safety, check for an error marshalling... probably superfluous
        // This will return the JSON error body as a string
        err = errors.New(string(str))
      }
  } else {
    // So, there was no network-type error, and nothing in the failure payload,
    // but we should still check the status code
    if httpResponse == nil {
      // This should never happen...
      err = errors.New("No HTTP Response received.")
    } else if code := httpResponse.StatusCode; 200 > code || code > 299 {
        err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
      }
    }
  }

  return *successPayload, err
}
/**
 * 
 * captures an payment
 * @param customerId 
 * @return Subscription
 */
//func (a DefaultApi) CustomersCustomerIdSubscriptionResumePost (customerId string) (Subscription, error) {
func (a DefaultApi) CustomersCustomerIdSubscriptionResumePost (customerId string) (Subscription, error) {

    _sling := a.NewWithAuth().Post(a.basePath)

    // create path and map variables
    path := "/customers/{customer_id}/subscription/resume"
    path = strings.Replace(path, "{" + "customer_id" + "}", fmt.Sprintf("%v", customerId), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string {  }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }


  var successPayload = new(Subscription)

  // We use this map (below) so that any arbitrary error JSON can be handled.
  // FIXME: This is in the absence of this Go generator honoring the non-2xx
  // response (error) models, which needs to be implemented at some point.
  var failurePayload map[string]interface{}

  httpResponse, err := _sling.Receive(successPayload, &failurePayload)

  if err == nil {
    // err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
    if failurePayload != nil {
      // If the failurePayload is present, there likely was some kind of non-2xx status
      // returned (and a JSON payload error present)
      var str []byte
      str, err = json.Marshal(failurePayload)
      if err == nil { // For safety, check for an error marshalling... probably superfluous
        // This will return the JSON error body as a string
        err = errors.New(string(str))
      }
  } else {
    // So, there was no network-type error, and nothing in the failure payload,
    // but we should still check the status code
    if httpResponse == nil {
      // This should never happen...
      err = errors.New("No HTTP Response received.")
    } else if code := httpResponse.StatusCode; 200 > code || code > 299 {
        err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
      }
    }
  }

  return *successPayload, err
}
/**
 * 
 * captures an payment
 * @param client 
 * @return Customer
 */
//func (a DefaultApi) CustomersPost (client BaseClient) (Customer, error) {
func (a DefaultApi) CustomersPost (client BaseClient) (Customer, error) {

    _sling := a.NewWithAuth().Post(a.basePath)

    // create path and map variables
    path := "/customers"

    _sling = _sling.Path(path)

    // accept header
    accepts := []string {  }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }

// body params
    _sling = _sling.BodyJSON(client)

  var successPayload = new(Customer)

  // We use this map (below) so that any arbitrary error JSON can be handled.
  // FIXME: This is in the absence of this Go generator honoring the non-2xx
  // response (error) models, which needs to be implemented at some point.
  var failurePayload map[string]interface{}

  httpResponse, err := _sling.Receive(successPayload, &failurePayload)

  if err == nil {
    // err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
    if failurePayload != nil {
      // If the failurePayload is present, there likely was some kind of non-2xx status
      // returned (and a JSON payload error present)
      var str []byte
      str, err = json.Marshal(failurePayload)
      if err == nil { // For safety, check for an error marshalling... probably superfluous
        // This will return the JSON error body as a string
        err = errors.New(string(str))
      }
  } else {
    // So, there was no network-type error, and nothing in the failure payload,
    // but we should still check the status code
    if httpResponse == nil {
      // This should never happen...
      err = errors.New("No HTTP Response received.")
    } else if code := httpResponse.StatusCode; 200 > code || code > 299 {
        err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
      }
    }
  }

  return *successPayload, err
}
/**
 * 
 * Gets all account events
 * @return []Event
 */
//func (a DefaultApi) EventsGet () ([]Event, error) {
func (a DefaultApi) EventsGet () ([]Event, error) {

    _sling := a.NewWithAuth().Get(a.basePath)

    // create path and map variables
    path := "/events"

    _sling = _sling.Path(path)

    // accept header
    accepts := []string {  }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }


  var successPayload = new([]Event)

  // We use this map (below) so that any arbitrary error JSON can be handled.
  // FIXME: This is in the absence of this Go generator honoring the non-2xx
  // response (error) models, which needs to be implemented at some point.
  var failurePayload map[string]interface{}

  httpResponse, err := _sling.Receive(successPayload, &failurePayload)

  if err == nil {
    // err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
    if failurePayload != nil {
      // If the failurePayload is present, there likely was some kind of non-2xx status
      // returned (and a JSON payload error present)
      var str []byte
      str, err = json.Marshal(failurePayload)
      if err == nil { // For safety, check for an error marshalling... probably superfluous
        // This will return the JSON error body as a string
        err = errors.New(string(str))
      }
  } else {
    // So, there was no network-type error, and nothing in the failure payload,
    // but we should still check the status code
    if httpResponse == nil {
      // This should never happen...
      err = errors.New("No HTTP Response received.")
    } else if code := httpResponse.StatusCode; 200 > code || code > 299 {
        err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
      }
    }
  }

  return *successPayload, err
}
/**
 * 
 * 
 * @param planId 
 * @return Plan
 */
//func (a DefaultApi) PlansPlanIdDelete (planId string) (Plan, error) {
func (a DefaultApi) PlansPlanIdDelete (planId string) (Plan, error) {

    _sling := a.NewWithAuth().Delete(a.basePath)

    // create path and map variables
    path := "/plans/{plan_id}"
    path = strings.Replace(path, "{" + "plan_id" + "}", fmt.Sprintf("%v", planId), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string {  }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }


  var successPayload = new(Plan)

  // We use this map (below) so that any arbitrary error JSON can be handled.
  // FIXME: This is in the absence of this Go generator honoring the non-2xx
  // response (error) models, which needs to be implemented at some point.
  var failurePayload map[string]interface{}

  httpResponse, err := _sling.Receive(successPayload, &failurePayload)

  if err == nil {
    // err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
    if failurePayload != nil {
      // If the failurePayload is present, there likely was some kind of non-2xx status
      // returned (and a JSON payload error present)
      var str []byte
      str, err = json.Marshal(failurePayload)
      if err == nil { // For safety, check for an error marshalling... probably superfluous
        // This will return the JSON error body as a string
        err = errors.New(string(str))
      }
  } else {
    // So, there was no network-type error, and nothing in the failure payload,
    // but we should still check the status code
    if httpResponse == nil {
      // This should never happen...
      err = errors.New("No HTTP Response received.")
    } else if code := httpResponse.StatusCode; 200 > code || code > 299 {
        err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
      }
    }
  }

  return *successPayload, err
}
/**
 * 
 * captures an payment
 * @param planId 
 * @param plan 
 * @return Plan
 */
//func (a DefaultApi) PlansPlanIdPut (planId string, plan Plan) (Plan, error) {
func (a DefaultApi) PlansPlanIdPut (planId string, plan Plan) (Plan, error) {

    _sling := a.NewWithAuth().Put(a.basePath)

    // create path and map variables
    path := "/plans/{plan_id}"
    path = strings.Replace(path, "{" + "plan_id" + "}", fmt.Sprintf("%v", planId), -1)

    _sling = _sling.Path(path)

    // accept header
    accepts := []string {  }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }

// body params
    _sling = _sling.BodyJSON(plan)

  var successPayload = new(Plan)

  // We use this map (below) so that any arbitrary error JSON can be handled.
  // FIXME: This is in the absence of this Go generator honoring the non-2xx
  // response (error) models, which needs to be implemented at some point.
  var failurePayload map[string]interface{}

  httpResponse, err := _sling.Receive(successPayload, &failurePayload)

  if err == nil {
    // err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
    if failurePayload != nil {
      // If the failurePayload is present, there likely was some kind of non-2xx status
      // returned (and a JSON payload error present)
      var str []byte
      str, err = json.Marshal(failurePayload)
      if err == nil { // For safety, check for an error marshalling... probably superfluous
        // This will return the JSON error body as a string
        err = errors.New(string(str))
      }
  } else {
    // So, there was no network-type error, and nothing in the failure payload,
    // but we should still check the status code
    if httpResponse == nil {
      // This should never happen...
      err = errors.New("No HTTP Response received.")
    } else if code := httpResponse.StatusCode; 200 > code || code > 299 {
        err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
      }
    }
  }

  return *successPayload, err
}
/**
 * 
 * captures an payment
 * @param subscription 
 * @return Plan
 */
//func (a DefaultApi) PlansPost (subscription Plan) (Plan, error) {
func (a DefaultApi) PlansPost (subscription Plan) (Plan, error) {

    _sling := a.NewWithAuth().Post(a.basePath)

    // create path and map variables
    path := "/plans"

    _sling = _sling.Path(path)

    // accept header
    accepts := []string {  }
    for key := range accepts {
        _sling = _sling.Set("Accept", accepts[key])
        break // only use the first Accept
    }

// body params
    _sling = _sling.BodyJSON(subscription)

  var successPayload = new(Plan)

  // We use this map (below) so that any arbitrary error JSON can be handled.
  // FIXME: This is in the absence of this Go generator honoring the non-2xx
  // response (error) models, which needs to be implemented at some point.
  var failurePayload map[string]interface{}

  httpResponse, err := _sling.Receive(successPayload, &failurePayload)

  if err == nil {
    // err == nil only means that there wasn't a sub-application-layer error (e.g. no network error)
    if failurePayload != nil {
      // If the failurePayload is present, there likely was some kind of non-2xx status
      // returned (and a JSON payload error present)
      var str []byte
      str, err = json.Marshal(failurePayload)
      if err == nil { // For safety, check for an error marshalling... probably superfluous
        // This will return the JSON error body as a string
        err = errors.New(string(str))
      }
  } else {
    // So, there was no network-type error, and nothing in the failure payload,
    // but we should still check the status code
    if httpResponse == nil {
      // This should never happen...
      err = errors.New("No HTTP Response received.")
    } else if code := httpResponse.StatusCode; 200 > code || code > 299 {
        err = errors.New("HTTP Error: " + string(httpResponse.StatusCode))
      }
    }
  }

  return *successPayload, err
}
