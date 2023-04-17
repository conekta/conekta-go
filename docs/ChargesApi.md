# \ChargesApi

All URIs are relative to *https://api.conekta.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrdersCreateCharge**](ChargesApi.md#OrdersCreateCharge) | **Post** /orders/{id}/charges | Create charge



## OrdersCreateCharge

> ChargeOrderResponse OrdersCreateCharge(ctx, id).ChargeRequest(chargeRequest).AcceptLanguage(acceptLanguage).Execute()

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
    chargeRequest := *openapiclient.NewChargeRequest(*openapiclient.NewChargeRequestPaymentMethod("card")) // ChargeRequest | requested field for a charge
    acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChargesApi.OrdersCreateCharge(context.Background(), id).ChargeRequest(chargeRequest).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChargesApi.OrdersCreateCharge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrdersCreateCharge`: ChargeOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `ChargesApi.OrdersCreateCharge`: %v\n", resp)
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

