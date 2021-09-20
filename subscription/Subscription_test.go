package subscription

import (
	conekta "github.com/conekta/conekta-go"
	"github.com/conekta/conekta-go/customer"
	"github.com/conekta/conekta-go/plan"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	conekta.APIKey = conekta.TestKey
}

func CreateCustomerTest() (*conekta.Customer, error) {
	cus := &conekta.CustomerParams{}
	c, err := customer.Create(cus.MockCustomerPaymentSource())
	return c, err
}

func CreateDeclinedCustomerTest() (*conekta.Customer, error) {
	cus := &conekta.CustomerParams{}
	customerParams := cus.MockCustomerPaymentSource()
	customerParams.PaymentSources[0].TokenID = "tok_test_card_declined"
	c, err := customer.Create(customerParams)
	return c, err
}

func CreatePlanTest() (*conekta.Plan, error) {
	pl := &conekta.PlanParams{}
	pl.MockPlanCreate()
	pl.TrialPeriodDays = 0
	p, err := plan.Create(pl)
	return p, err
}

func TestCreate(t *testing.T) {
	c, _ := CreateCustomerTest()
	p, _ := CreatePlanTest()

	sp := &conekta.SubscriptionParams{
		PlanID: p.ID,
		CardID: c.DefaultPaymentSourceID,
	}
	sub, err := Create(c.ID, sp)

	assert.Equal(t, sp.CardID, sub.CardID)
	assert.Equal(t, sp.PlanID, sub.PlanID)

	assert.Nil(t, err)
}

func TestCreateProcessingError(t *testing.T) {
	c, _ := CreateDeclinedCustomerTest()
	p, _ := CreatePlanTest()

	sp := &conekta.SubscriptionParams{
		PlanID: p.ID,
		CardID: c.DefaultPaymentSourceID,
	}
	sub, err := Create(c.ID, sp)

	assert.Equal(t, sp.CardID, sub.CardID)
	assert.Equal(t, sp.PlanID, sub.PlanID)
	assert.Equal(t, "past_due", sub.Status)

	assert.Nil(t, err)
}

func TestCreateError(t *testing.T) {
	cus := &conekta.CustomerParams{}
	_, err := Create(cus.ID, &conekta.SubscriptionParams{})
	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "parameter_validation_error")
}

func TestResume(t *testing.T) {
	c, _ := CreateCustomerTest()
	p, _ := CreatePlanTest()

	sp := &conekta.SubscriptionParams{
		PlanID: p.ID,
		CardID: c.DefaultPaymentSourceID,
	}
	sub, _ := Create(c.ID, sp)
	Pause(sub.CustomerID)

	_, err := Resume(sub.CustomerID)
	assert.Nil(t, err)
}

func TestResumeProcessingError(t *testing.T) {
	c, _ := CreateDeclinedCustomerTest()
	p, _ := CreatePlanTest()

	sp := &conekta.SubscriptionParams{
		PlanID: p.ID,
		CardID: c.DefaultPaymentSourceID,
	}
	sub, _ := Create(c.ID, sp)
	Pause(sub.CustomerID)

	pausedSub, err := Resume(sub.CustomerID)
	assert.Nil(t, err)
	assert.Equal(t, "past_due", pausedSub.Status)
}

func TestPause(t *testing.T) {
	c, _ := CreateCustomerTest()
	p, _ := CreatePlanTest()

	sp := &conekta.SubscriptionParams{
		PlanID: p.ID,
		CardID: c.DefaultPaymentSourceID,
	}
	sub, _ := Create(c.ID, sp)

	_, err := Pause(sub.CustomerID)
	assert.Nil(t, err)
}

func TestCancel(t *testing.T) {
	c, _ := CreateCustomerTest()
	p, _ := CreatePlanTest()

	sp := &conekta.SubscriptionParams{
		PlanID: p.ID,
		CardID: c.DefaultPaymentSourceID,
	}
	sub, _ := Create(c.ID, sp)

	_, err := Cancel(sub.CustomerID)
	assert.Nil(t, err)
}
