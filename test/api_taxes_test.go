/*
Conekta API

Testing TaxesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package conekta

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/conekta/conekta-go"
)

func Test_conekta_TaxesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TaxesAPIService OrdersCreateTaxes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.TaxesAPI.OrdersCreateTaxes(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaxesAPIService OrdersDeleteTaxes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var taxId string

		resp, httpRes, err := apiClient.TaxesAPI.OrdersDeleteTaxes(context.Background(), id, taxId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaxesAPIService OrdersUpdateTaxes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var taxId string

		resp, httpRes, err := apiClient.TaxesAPI.OrdersUpdateTaxes(context.Background(), id, taxId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
