package webhook

import (
	"testing"

	"github.com/conekta/conekta-go"

	"github.com/stretchr/testify/assert"
)

func init() {
	conekta.APIKey = conekta.TestKey
}

func TestCreate(t *testing.T) {
	whp := &conekta.WebhookParams{}
	wh, err := Create(whp.Mock())
	assert.Equal(t, whp.URL, wh.URL)
	assert.NotZero(t, wh.SubscribedEvents)
	assert.Nil(t, err)
	Delete(wh.ID)
}

func TestCreateError(t *testing.T) {
	_, err := Create(&conekta.WebhookParams{})
	assert.NotNil(t, err)
	assert.Equal(t, err.ErrorType, "parameter_validation_error")
}

func TestUpdate(t *testing.T) {
	whp := &conekta.WebhookParams{}
	wh, _ := Create(whp.Mock())
	newURL := "anotherurl.com"
	whp.URL = newURL
	wh, err := Update(wh.ID, whp)
	assert.Equal(t, newURL, wh.URL)
	assert.Nil(t, err)
	Delete(wh.ID)
}

func TestFind(t *testing.T) {
	whp := &conekta.WebhookParams{}
	wh, _ := Create(whp.Mock())
	wh, err := Find(wh.ID)
	assert.Equal(t, whp.URL, wh.URL)
	assert.Nil(t, err)
	Delete(wh.ID)
}

func TestFindError(t *testing.T) {
	_, err := Find("122")
	assert.NotNil(t, err)
	assert.Equal(t, err.ErrorType, "resource_not_found_error")
}

func TestDelete(t *testing.T) {
	whp := &conekta.WebhookParams{}
	wh, _ := Create(whp.Mock())
	wh, err := Delete(wh.ID)
	assert.Equal(t, wh.Deleted, true)
	assert.Nil(t, err)
	Delete(wh.ID)
}

func TestDeleteError(t *testing.T) {
	_, err := Delete("122")
	assert.NotNil(t, err)
	assert.Equal(t, err.ErrorType, "resource_not_found_error")
}

func TestAll(t *testing.T) {
	whp := &conekta.WebhookParams{}
	wh, _ := Create(whp.Mock())
	whl, err := All()
	assert.NotZero(t, whl.Total)
	assert.Nil(t, err)
	Delete(wh.ID)
}
