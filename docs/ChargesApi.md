# \ChargesAPI

All URIs are relative to *https://api.conekta.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCharges**](ChargesAPI.md#GetCharges) | **Get** /charges | Get A List of Charges
[**OrdersCreateCharge**](ChargesAPI.md#OrdersCreateCharge) | **Post** /orders/{id}/charges | Create charge
[**OrdersCreateCharges**](ChargesAPI.md#OrdersCreateCharges) | **Post** /orders/{id}/add_charges | Create charges
[**UpdateCharge**](ChargesAPI.md#UpdateCharge) | **Put** /charges/{id} | Update a charge



## GetCharges

> GetChargesResponse GetCharges(ctx).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Limit(limit).Search(search).Next(next).Previous(previous).Execute()

Get A List of Charges

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/conekta/conekta-go"
)

func main() {
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)
	limit := int32(56) // int32 | The numbers of items to return, the maximum value is 250 (optional) (default to 20)
	search := "search_example" // string | General order search, e.g. by mail, reference etc. (optional)
	next := "next_example" // string | next page (optional)
	previous := "previous_example" // string | previous page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChargesAPI.GetCharges(context.Background()).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Limit(limit).Search(search).Next(next).Previous(previous).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChargesAPI.GetCharges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCharges`: GetChargesResponse
	fmt.Fprintf(os.Stdout, "Response from `ChargesAPI.GetCharges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChargesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 
 **limit** | **int32** | The numbers of items to return, the maximum value is 250 | [default to 20]
 **search** | **string** | General order search, e.g. by mail, reference etc. | 
 **next** | **string** | next page | 
 **previous** | **string** | previous page | 

### Return type

[**GetChargesResponse**](GetChargesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrdersCreateCharge

> ChargeOrderResponse OrdersCreateCharge(ctx, id).ChargeRequest(chargeRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Create charge



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/conekta/conekta-go"
)

func main() {
	id := "6307a60c41de27127515a575" // string | Identifier of the resource
	chargeRequest := *openapiclient.NewChargeRequest(openapiclient.charge_request_payment_method{PaymentMethodCardRequest: openapiclient.NewPaymentMethodCardRequest("card | cash | spei", "198", "12", "2025", "John Doe", "4242424242424242")}) // ChargeRequest | requested field for a charge
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChargesAPI.OrdersCreateCharge(context.Background(), id).ChargeRequest(chargeRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChargesAPI.OrdersCreateCharge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrdersCreateCharge`: ChargeOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `ChargesAPI.OrdersCreateCharge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrdersCreateChargeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chargeRequest** | [**ChargeRequest**](ChargeRequest.md) | requested field for a charge | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**ChargeOrderResponse**](ChargeOrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrdersCreateCharges

> ChargesOrderResponse OrdersCreateCharges(ctx, id).ChargeRequest(chargeRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Create charges



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/conekta/conekta-go"
)

func main() {
	id := "6307a60c41de27127515a575" // string | Identifier of the resource
	chargeRequest := *openapiclient.NewChargeRequest(openapiclient.charge_request_payment_method{PaymentMethodCardRequest: openapiclient.NewPaymentMethodCardRequest("card | cash | spei", "198", "12", "2025", "John Doe", "4242424242424242")}) // ChargeRequest | requested field for a charge
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChargesAPI.OrdersCreateCharges(context.Background(), id).ChargeRequest(chargeRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChargesAPI.OrdersCreateCharges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrdersCreateCharges`: ChargesOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `ChargesAPI.OrdersCreateCharges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrdersCreateChargesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chargeRequest** | [**ChargeRequest**](ChargeRequest.md) | requested field for a charge | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**ChargesOrderResponse**](ChargesOrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCharge

> ChargeResponse UpdateCharge(ctx, id).ChargeUpdateRequest(chargeUpdateRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Update a charge

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/conekta/conekta-go"
)

func main() {
	id := "6307a60c41de27127515a575" // string | Identifier of the resource
	chargeUpdateRequest := *openapiclient.NewChargeUpdateRequest() // ChargeUpdateRequest | requested field for update a charge
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChargesAPI.UpdateCharge(context.Background(), id).ChargeUpdateRequest(chargeUpdateRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChargesAPI.UpdateCharge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCharge`: ChargeResponse
	fmt.Fprintf(os.Stdout, "Response from `ChargesAPI.UpdateCharge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChargeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chargeUpdateRequest** | [**ChargeUpdateRequest**](ChargeUpdateRequest.md) | requested field for update a charge | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**ChargeResponse**](ChargeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

