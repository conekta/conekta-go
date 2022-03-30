package plan

import (
	conekta "github.com/conekta/conekta-go/v2"
	//"github.com/google/go-querystring/query"
)

// Create creates a new plan
func Create(p *conekta.PlanParams, customHeaders ...interface{}) (*conekta.Plan, error) {
	pla := &conekta.Plan{}
	err := conekta.MakeRequest("POST", "/plans", p, pla, customHeaders...)
	return pla, err
}

// Update updates a Plans
// For details see https://developers.conekta.com/api#update-plan
func Update(id string, p *conekta.PlanParams) (*conekta.Plan, error) {
	pla := &conekta.Plan{}
	err := conekta.MakeRequest("PUT", "/plans/"+id, p, pla)
	return pla, err
}

// Find gets a plan by id
func Find(id string) (*conekta.Plan, error) {
	pla := &conekta.Plan{}
	err := conekta.MakeRequest("GET", "/plans/"+id, &conekta.PlanParams{}, pla)
	return pla, err
}

// Delete deletes a Plan
func Delete(id string) (*conekta.Plan, error) {
	pla := &conekta.Plan{}
	err := conekta.MakeRequest("DELETE", "/plans/"+id, &conekta.PlanParams{}, pla)
	return pla, err
}
