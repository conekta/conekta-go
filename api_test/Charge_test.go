package api_test

import (
	"encoding/json"
	"testing"
	"github.com/stretchr/testify/assert"
	"conekta.io/api"
)

var simpleCharge = `{
			      "description":"Stogies",
			      "amount": 20000,
			      "currency":"MXN",
			      "reference_id":"9839-wolf_pack",
			      "card": "tok_test_visa_4242",
			      "details": {
				"name": "Arnulfo Quimare",
				"phone": "403-342-0642",
				"email": "logan@x-men.org",
				"customer": {
				  "logged_in": true,
				  "successful_purchases": 14,
				  "created_at": 1379784950,
				  "updated_at": 1379784950,
				  "offline_payments": 4,
				  "score": 9
				},
				"line_items": [{
				  "name": "Box of Cohiba S1s",
				  "description": "Imported From Mex.",
				  "unit_price": 20000,
				  "quantity": 1,
				  "sku": "cohb_s1",
				  "category": "food"
				}],
				"billing_address": {
				  "street1":"77 Mystery Lane",
				  "street2": "Suite 124",
				  "city": "Darlington",
				  "state":"NJ",
				  "zip": "10192",
				  "country": "Mexico",
				  "tax_id": "xmn671212drx",
				  "company_name":"X-Men Inc.",
				  "phone": "77-777-7777",
				  "email": "purshasing@x-men.org"
				}
			      }
			    }`

var messesSinInterersis = `{
      "description":"Stogies",
      "amount": 20000,
      "currency":"MXN",
      "reference_id":"9839-wolf_pack",
      "card": "tok_test_visa_4242",
      "monthly_installments": 3,
      "details": {
        "name": "Arnulfo Quimare",
        "phone": "403-342-0642",
        "email": "logan@x-men.org",
        "customer": {
          "logged_in": true,
          "successful_purchases": 14,
          "created_at": 1379784950,
          "updated_at": 1379784950,
          "offline_payments": 4,
          "score": 9
        },
        "line_items": [{
          "name": "Box of Cohiba S1s",
          "description": "Imported From Mex.",
          "unit_price": 20000,
          "quantity": 1,
          "sku": "cohb_s1",
          "category": "food"
        }],
        "billing_address": {
          "street1":"77 Mystery Lane",
          "street2": "Suite 124",
          "street3": null,
          "city": "Darlington",
          "state":"NJ",
          "zip": "10192",
          "country": "Mexico",
          "tax_id": "xmn671212drx",
          "company_name":"X-Men Inc.",
          "phone": "77-777-7777",
          "email": "purshasing@x-men.org"
        }
      }
    }`

var oxxoCharges = `{
      "description":"Stogies",
      "amount": 20000,
      "currency":"MXN",
      "reference_id":"9839-wolf_pack",
      "cash": {
        "type":"oxxo",
        "expires_at": 1459106145
      },
      "details": {
        "name": "Arnulfo Quimare",
        "phone": "403-342-0642",
        "email": "logan@x-men.org",
        "customer": {
          "logged_in": true,
          "successful_purchases": 14,
          "created_at": 1379784950,
          "updated_at": 1379784950,
          "offline_payments": 4,
          "score": 9
        },
        "line_items": [{
          "name": "Box of Cohiba S1s",
          "description": "Imported From Mex.",
          "unit_price": 20000,
          "quantity": 1,
          "sku": "cohb_s1",
          "category": "food"
        }],
        "billing_address": {
          "street1":"77 Mystery Lane",
          "street2": "Suite 124",
          "street3": null,
          "city": "Darlington",
          "state":"NJ",
          "zip": "10192",
          "country": "Mexico",
          "tax_id": "xmn671212drx",
          "company_name":"X-Men Inc.",
          "phone": "77-777-7777",
          "email": "purshasing@x-men.org"
        }
      }
    }`

var speiCharge = `{
      "description":"Stogies",
      "amount": 20000,
      "currency":"MXN",
      "reference_id":"9839-wolf_pack",
      "bank": {
        "type": "spei"
      },
      "details": {
        "name": "Arnulfo Quimare",
        "phone": "403-342-0642",
        "email": "logan@x-men.org",
        "customer": {
          "logged_in": true,
          "successful_purchases": 14,
          "created_at": 1379784950,
          "updated_at": 1379784950,
          "offline_payments": 4,
          "score": 9
        },
        "line_items": [{
          "name": "Box of Cohiba S1s",
          "description": "Imported From Mex.",
          "unit_price": 20000,
          "quantity": 1,
          "sku": "cohb_s1",
          "category": "food"
        }],
        "billing_address": {
          "street1":"77 Mystery Lane",
          "street2": "Suite 124",
          "street3": null,
          "city": "Darlington",
          "state":"NJ",
          "zip": "10192",
          "country": "Mexico",
          "tax_id": "xmn671212drx",
          "company_name":"X-Men Inc.",
          "phone": "77-777-7777",
          "email": "purshasing@x-men.org"
        }
      }
    }`

func ExampleCharge(jsonObject string) api.Charge {
	charge := api.Charge{}
	json.Unmarshal([]byte(jsonObject), &charge)
	return charge;
}

func TestChargePost(t *testing.T) {
	conekta := api.Conekta()
	conekta.SetApikey("key_eYvWV7gSDkNYXsmr");

	charge, _ := api.CreateCharge(ExampleCharge(simpleCharge))

	assert.NotNil(t, charge, "should have charge")
	assert.NotNil(t, charge.Id, "should have id")
}


func TestChargePostMessesSinIntereses(t *testing.T) {
	conekta := api.Conekta()
	conekta.SetApikey("key_eYvWV7gSDkNYXsmr");

	charge, _ := api.CreateCharge(ExampleCharge(messesSinInterersis))

	assert.NotNil(t, charge, "should have charge")
	assert.NotNil(t, charge.ReferenceId, "should have Reference id")
	assert.NotNil(t, charge.PaymentMethod, "should have Payment ")
}

func TestChargeOxxo(t *testing.T) {
	conekta := api.Conekta()
	conekta.SetApikey("key_eYvWV7gSDkNYXsmr");

	charge, _ := api.CreateCharge(ExampleCharge(oxxoCharges))

	assert.NotNil(t, charge, "should have charge")
	assert.NotNil(t, charge.ReferenceId, "should have Reference id")
	assert.NotNil(t, charge.PaymentMethod, "should have Payment ")
	assert.NotNil(t, charge.PaymentMethod.Barcode, "should have Barcode")
}

func TestSpeiCharge(t *testing.T) {
	conekta := api.Conekta()
	conekta.SetApikey("key_eYvWV7gSDkNYXsmr");

	charge, _ := api.CreateCharge(ExampleCharge(speiCharge))

	assert.NotNil(t, charge, "should have charge")
	assert.NotNil(t, charge.ReferenceId, "should have Reference id")
	assert.NotNil(t, charge.PaymentMethod, "should have Payment ")
	assert.NotNil(t, charge.PaymentMethod.Clabe, "should have Payment ")

}


func TestFindCharge(t *testing.T) {
	conekta := api.Conekta()
	conekta.SetApikey("key_eYvWV7gSDkNYXsmr");

	charge, _ := api.CreateCharge(ExampleCharge(speiCharge))

	foundCharged, _ := api.FindCharge(charge.Id)

	assert.NotNil(t, foundCharged, "should return a charge")

}


func TestFindChargeRefund(t *testing.T) {
	conekta := api.Conekta()
	conekta.SetApikey("key_eYvWV7gSDkNYXsmr");

	charge, _ := api.CreateCharge(ExampleCharge(speiCharge))

	foundCharged, errors := api.RefundCharge(charge.Id, api.AmountQty{Amount: charge.Amount})

	if (errors != nil) {
		assert.Error(t, errors, "should not contain errors")
	}

	assert.NotNil(t, foundCharged, "should return a charge")

}



func TestFindChargeArgs(t *testing.T) {
	conekta := api.Conekta()
	conekta.SetApikey("key_eYvWV7gSDkNYXsmr");

	charges, _ := api.FindCharges(map[string]string{
		"status.ne":"paid",
	})

	assert.NotEmpty(t, charges, "should return a charge")
}


