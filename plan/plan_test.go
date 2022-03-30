package plan

import (
	"testing"

	"github.com/conekta/conekta-go/v2"
	"github.com/stretchr/testify/assert"
)

func init() {
	conekta.APIKey = conekta.TestKey
}

func TestCreate(t *testing.T) {
	pp := &conekta.PlanParams{}
	pl, err := Create(pp.MockPlanCreate())
	assert.Equal(t, pp.Name, pl.Name)
	assert.Equal(t, pp.Amount, pl.Amount)
	assert.Equal(t, pp.Interval, pl.Interval)
	assert.Equal(t, pp.Frequency, pl.Frequency)
	assert.Equal(t, pp.TrialPeriodDays, pl.TrialPeriodDays)
	assert.Equal(t, pp.ExpiryCount, pl.ExpiryCount)
	assert.Nil(t, err)
}

func TestCreateError(t *testing.T) {
	_, err := Create(&conekta.PlanParams{})
	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "parameter_validation_error")
}

func TestUpdate(t *testing.T) {
	pp := &conekta.PlanParams{}
	pl, _ := Update(pp.ID, pp.MockPlanUpdate())
	assert.Equal(t, pp.Amount, pl.Amount)
}

func TestFind(t *testing.T) {
	pp := &conekta.PlanParams{}
	pl, _ := Find(pp.ID)
	assert.Equal(t, pp.ID, pl.ID)
}

func TestDelete(t *testing.T) {
	pp := &conekta.PlanParams{}
	pl, _ := Delete(pp.ID)
	assert.Equal(t, pp.ID, pl.ID)
}
