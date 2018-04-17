package api

import (
)

type Customer struct {
    Id  string  `json:"id,omitempty"`
    Object  string  `json:"object,omitempty"`
    Livemode  string  `json:"livemode,omitempty"`
    CreatedAt  string  `json:"created_at,omitempty"`
    Name  string  `json:"name,omitempty"`
    Email  string  `json:"email,omitempty"`
    Phone  string  `json:"phone,omitempty"`
    DefaultCard  string  `json:"default_card,omitempty"`
    BillingAddress  Address  `json:"billing_address,omitempty"`
    ShippingAddress  Address  `json:"shipping_address,omitempty"`
    Cards  []TokenizedCard  `json:"cards,omitempty"`
    Subscription  Subscription  `json:"subscription,omitempty"`
    Deleted  bool  `json:"deleted,omitempty"`

}


func CreateCustomer(customer BaseClient)(Customer, error){
    api := NewDefaultApi(Conekta().apiKey);
    return api.CustomersPost(customer)
}

func DeleteCustomer(customerId string)(Customer, error){
    api := NewDefaultApi(Conekta().apiKey);
    return api.CustomersCustomerIdDelete(customerId)
}

func AddCustomerCard(customer Customer, card TokenObject)(Card, error){
    api := NewDefaultApi(Conekta().apiKey);
    return api.CustomersCustomerIdCardsPost(customer.Id, card);
}


func DeleteCustomerCard(customer Customer, card Card)(Card, error){
    api := NewDefaultApi(Conekta().apiKey);
    return api.CustomersCustomerIdCardsCardIdDelete(customer.Id, card.Id)
}


func CreateCustomerSubscription(customer Customer, subscription SubscriptionReference)(Subscription, error){
    api := NewDefaultApi(Conekta().apiKey);
    return api.CustomersCustomerIdSubscriptionPost(customer.Id, subscription)
}


func UpdateCustomerSubscription(customer Customer, subscription SubscriptionReference)(Subscription, error){
    api := NewDefaultApi(Conekta().apiKey);
    return api.CustomersCustomerIdSubscriptionPut(customer.Id, subscription)
}


func PauseCustomerSubscription(customer Customer)(Subscription, error){
    api := NewDefaultApi(Conekta().apiKey);
    return api.CustomersCustomerIdSubscriptionPausePost(customer.Id)
}


func ResumeCustomerSubscription(customer Customer)(Subscription, error){
    api := NewDefaultApi(Conekta().apiKey);
    return api.CustomersCustomerIdSubscriptionResumePost(customer.Id)
}

func CancelCustomerSubscription(customer Customer)(Subscription, error){
    api := NewDefaultApi(Conekta().apiKey);
    return api.CustomersCustomerIdSubscriptionCancelPost(customer.Id)
}


