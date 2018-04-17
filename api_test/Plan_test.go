package api

import (
	"conekta.io/api"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"testing"
	"math/rand"
	"fmt"
)

var planObject = `{
       "name":"Gold Plan",
       "amount":10000,
       "currency":"MXN",
       "interval":"month",
       "frequency":1,
       "trial_period_days":15,
       "expiry_count":12
     }`

func UnmarshalPlan(jsonObject string) api.Plan {
	base := api.Plan{}
	json.Unmarshal([]byte(jsonObject), &base)
	return base;
}


func TestCreatePlan(t *testing.T) {
	conekta := api.Conekta()
	conekta.SetApikey("key_eYvWV7gSDkNYXsmr");

	planRef := UnmarshalPlan(planObject)
	planRef.Id = "ctestr-plan"+  RandStringBytes(rand.Intn(50))

	plan, _:= api.CreatePlan(planRef)


	assert.NotNil(t, plan.Id, "the object should return a plan")
}


func TestUpdatePlan(t *testing.T) {
	conekta := api.Conekta()
	conekta.SetApikey("key_eYvWV7gSDkNYXsmr");

	planRef := UnmarshalPlan(planObject)
	planRef.Id = RandStringBytes(12) +"-plan"+  RandStringBytes(10)

	fmt.Printf("%+v\n", planRef)

	plan, error:= api.CreatePlan(planRef)
	plan.Name = "TestPlan"
	updatedPlan, _ := api.UpdatePlan(plan.Id, plan)

	if(error != nil){
		fmt.Printf("%+v\n", error)
		fmt.Printf("%+v\n", updatedPlan)
	}

	assert.Equal(t, "TestPlan", updatedPlan.Name, "the object should update the plna")
}


func TestDeletePlan(t *testing.T) {
	conekta := api.Conekta()
	conekta.SetApikey("key_eYvWV7gSDkNYXsmr");

	planRef := UnmarshalPlan(planObject)
	planRef.Id = "ctestr-plan"+  RandStringBytes(rand.Intn(50))

	plan, _:= api.CreatePlan(planRef)

	deletedPlan, _ := api.DeletePlan(plan.Id)

	assert.NotEmpty(t, deletedPlan.Id)

}




const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}