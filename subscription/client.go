package subscription

import (
	"encoding/json"
	conekta "github.com/conekta/conekta-go"
	//"github.com/google/go-querystring/query"
)

// Create creates a new subscription
func Create(id string, p *conekta.SubscriptionParams, customHeaders ...interface{}) (*conekta.Subscription, error) {
	sub := &conekta.Subscription{}
	err := conekta.MakeRequest("POST", "/customers/"+id+"/subscription", p, sub, customHeaders...)
	if err != nil && err.(conekta.Error).ErrorType == "processing_error" {
		json.Unmarshal(err.(conekta.Error).Data, sub)
		return sub, nil
	}
	return sub, err
}

func Update(id string, p *conekta.SubscriptionParams) (*conekta.Subscription, error) {
	sub := &conekta.Subscription{}
	err := conekta.MakeRequest("POST", "/customers/"+id+"/subscription", p, sub)
	if err != nil && err.(conekta.Error).ErrorType == "processing_error" {
		json.Unmarshal(err.(conekta.Error).Data, sub)
		return sub, nil
	}
	return sub, err
}

func Resume(id string) (*conekta.Subscription, error) {
	sub := &conekta.Subscription{}
	err := conekta.MakeRequest("POST", "/customers/"+id+"/subscription/resume", &conekta.SubscriptionParams{}, sub)
	if err != nil && err.(conekta.Error).ErrorType == "processing_error" {
		json.Unmarshal(err.(conekta.Error).Data, sub)
		return sub, nil
	}
	return sub, err
}

func Pause(id string) (*conekta.Subscription, error) {
	sub := &conekta.Subscription{}
	err := conekta.MakeRequest("POST", "/customers/"+id+"/subscription/pause", &conekta.SubscriptionParams{}, sub)
	return sub, err
}

func Cancel(id string) (*conekta.Subscription, error) {
	sub := &conekta.Subscription{}
	err := conekta.MakeRequest("POST", "/customers/"+id+"/subscription/cancel", &conekta.SubscriptionParams{}, sub)
	return sub, err
}

func Find(id string) (*conekta.Subscription, error) {
	subs := &conekta.Subscription{}
	err := conekta.MakeRequest("GET", "/customers/"+id+"/subscription", &conekta.EmptyParams{}, subs)
	return subs, err
}
