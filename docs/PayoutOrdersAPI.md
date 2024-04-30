# \PayoutOrdersAPI

All URIs are relative to *https://api.conekta.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePayoutOrder**](PayoutOrdersAPI.md#CreatePayoutOrder) | **Post** /payout_orders | Create payout order
[**GetPayoutOrderById**](PayoutOrdersAPI.md#GetPayoutOrderById) | **Get** /payout_orders/{id} | Get Payout Order
[**GetPayoutOrders**](PayoutOrdersAPI.md#GetPayoutOrders) | **Get** /payout_orders | Get a list of Payout Orders



## CreatePayoutOrder

> PayoutOrderResponse CreatePayoutOrder(ctx).PayoutOrder(payoutOrder).AcceptLanguage(acceptLanguage).Execute()

Create payout order



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
	payoutOrder := *openapiclient.NewPayoutOrder([]string{"AllowedPayoutMethods_example"}, int32(100), "MXN", *openapiclient.NewCustomerInfoJustCustomerId("cus_23874283647"), *openapiclient.NewPayout(*openapiclient.NewPayoutMethod("cashout")), "Payout order for the customer") // PayoutOrder | requested field for payout order
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PayoutOrdersAPI.CreatePayoutOrder(context.Background()).PayoutOrder(payoutOrder).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PayoutOrdersAPI.CreatePayoutOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePayoutOrder`: PayoutOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `PayoutOrdersAPI.CreatePayoutOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePayoutOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payoutOrder** | [**PayoutOrder**](PayoutOrder.md) | requested field for payout order | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]

### Return type

[**PayoutOrderResponse**](PayoutOrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayoutOrderById

> PayoutOrderResponse GetPayoutOrderById(ctx, id).AcceptLanguage(acceptLanguage).Execute()

Get Payout Order



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
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PayoutOrdersAPI.GetPayoutOrderById(context.Background(), id).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PayoutOrdersAPI.GetPayoutOrderById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPayoutOrderById`: PayoutOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `PayoutOrdersAPI.GetPayoutOrderById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPayoutOrderByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]

### Return type

[**PayoutOrderResponse**](PayoutOrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayoutOrders

> PayoutOrdersResponse GetPayoutOrders(ctx).AcceptLanguage(acceptLanguage).Limit(limit).Search(search).Next(next).Previous(previous).Execute()

Get a list of Payout Orders



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
	limit := int32(56) // int32 | The numbers of items to return, the maximum value is 250 (optional) (default to 20)
	search := "search_example" // string | General order search, e.g. by mail, reference etc. (optional)
	next := "next_example" // string | next page (optional)
	previous := "previous_example" // string | previous page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PayoutOrdersAPI.GetPayoutOrders(context.Background()).AcceptLanguage(acceptLanguage).Limit(limit).Search(search).Next(next).Previous(previous).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PayoutOrdersAPI.GetPayoutOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPayoutOrders`: PayoutOrdersResponse
	fmt.Fprintf(os.Stdout, "Response from `PayoutOrdersAPI.GetPayoutOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPayoutOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **limit** | **int32** | The numbers of items to return, the maximum value is 250 | [default to 20]
 **search** | **string** | General order search, e.g. by mail, reference etc. | 
 **next** | **string** | next page | 
 **previous** | **string** | previous page | 

### Return type

[**PayoutOrdersResponse**](PayoutOrdersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

