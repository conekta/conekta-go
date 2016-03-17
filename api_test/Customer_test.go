package api

import (
	"encoding/json"
	"testing"
	"github.com/stretchr/testify/assert"
	"conekta.io/api"
	"fmt"
)

var baseClient = `{
         "name":"James Howlett",
         "email":"james.howlett@forces.gov",
         "phone":"55-5555-5555",
         "cards": ["tok_8kZwafM8IcN23Nd9"],
         "plan": "gold-plan"
     }`


func Unmarshal(jsonObject string) api.BaseClient {
	base := api.BaseClient{}
	json.Unmarshal([]byte(jsonObject), &base)
	return base;
}

func TestCreateCustomer(t *testing.T) {
	conekta := api.Conekta()
	conekta.SetApikey("key_eYvWV7gSDkNYXsmr");

	customer, _ := api.CreateCustomer(Unmarshal(baseClient))

	assert.NotNil(t, customer, "should have charge")
	assert.NotNil(t, customer.Id, "should have id")
}

func TestDeleteCustomer(t *testing.T) {
	conekta := api.Conekta()
	conekta.SetApikey("key_eYvWV7gSDkNYXsmr");

	customer, _ := api.CreateCustomer(Unmarshal(baseClient))
	deletedCustomer, _ := api.DeleteCustomer(customer.Id)

	assert.NotNil(t, deletedCustomer.Id, "should be deleted")
	assert.NotNil(t, deletedCustomer.Deleted, "should be deleted")
}


func TestCustomerAddCard(t *testing.T) {
	conekta := api.Conekta()
	conekta.SetApikey("key_eYvWV7gSDkNYXsmr");

	customer, _ := api.CreateCustomer(Unmarshal(baseClient))
	card, _ := api.AddCustomerCard(customer, api.TokenObject{Token: "tok_8kZwafM8IcN23Nd9"})

	assert.NotNil(t, card, "should have charge")
	assert.NotNil(t, card.Id, "should have id")
}

func TestUpdateCustomerCard(t *testing.T) {
	conekta := api.Conekta()
	conekta.SetApikey("key_eYvWV7gSDkNYXsmr");

	customer, errors := api.CreateCustomer(Unmarshal(baseClient))
	card, _ := api.AddCustomerCard(customer, api.TokenObject{Token: "tok_8kZwafM8IcN23Nd9"})
	deletedCard, _ := api.DeleteCustomerCard(customer, card)

	fmt.Printf("%+v\n", errors)

	assert.NotNil(t, deletedCard, "should have charge")
	assert.NotNil(t, deletedCard.Id, "should have id")
}


func TestAddCustomerSubscrpition(t *testing.T) {
	conekta := api.Conekta()
	conekta.SetApikey("key_eYvWV7gSDkNYXsmr");

	customer, _ := api.CreateCustomer(Unmarshal(baseClient))

	subscription, _ := api.CreateCustomerSubscription(customer, api.SubscriptionReference{Plan: "opal-plan"})

	assert.NotNil(t, subscription, "should have charge")
	assert.NotNil(t, subscription.Id, "should have id")
	assert.NotNil(t, subscription.PlanId, "should have id")

}

func TestUpdateSubscriptionPlan(t *testing.T) {
	conekta := api.Conekta()
	conekta.SetApikey("key_eYvWV7gSDkNYXsmr");

	customer, _ := api.CreateCustomer(Unmarshal(baseClient))

	subscription, _ := api.UpdateCustomerSubscription(customer, api.SubscriptionReference{Plan: "opal-plan"})

	assert.NotNil(t, subscription, "should have charge")
	assert.NotNil(t, subscription.Id, "should have id")
	assert.NotNil(t, subscription.PlanId, "should have id")

}