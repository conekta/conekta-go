package api

import (
)

type Plan struct {
    Id  string  `json:"id,omitempty"`
    Object  string  `json:"object,string,omitempty"`
    Livemode  bool  `json:"livemode,omitempty"`
    CreatedAt  int32  `json:"created_at,omitempty"`
    Name  string  `json:"name,omitempty"`
    Amount  int32  `json:"amount,omitempty"`
    Currency  string  `json:"currency,omitempty"`
    Interval  int32  `json:"interval,omitempty"`
    Frequency  string  `json:"frequency,omitempty"`
    IntervalTotalCount  int32  `json:"interval_total_count,omitempty"`
    TrialPeriodDays  int32  `json:"trial_period_days,omitempty"`
    Deleted  bool  `json:"deleted,string,omitempty"`
}


func CreatePlan(plan Plan)(Plan, error){
    api := NewDefaultApi(Conekta().apiKey);
    return api.PlansPost(plan)
}

func UpdatePlan(planId string,plan Plan)(Plan, error){
    api := NewDefaultApi(Conekta().apiKey);
    return api.PlansPlanIdPut(planId, plan)
}

func DeletePlan(planId string)(Plan, error){
    api := NewDefaultApi(Conekta().apiKey);
    return api.PlansPlanIdDelete(planId);
}