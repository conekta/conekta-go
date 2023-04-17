# \TaxesApi

All URIs are relative to *https://api.conekta.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrdersCreateTaxes**](TaxesApi.md#OrdersCreateTaxes) | **Post** /orders/{id}/tax_lines | Create Tax
[**OrdersDeleteTaxes**](TaxesApi.md#OrdersDeleteTaxes) | **Delete** /orders/{id}/tax_lines/{tax_id} | Delete Tax
[**OrdersUpdateTaxes**](TaxesApi.md#OrdersUpdateTaxes) | **Put** /orders/{id}/tax_lines/{tax_id} | Update Tax



## OrdersCreateTaxes

> UpdateOrderTaxResponse OrdersCreateTaxes(ctx, id).OrderTaxRequest(orderTaxRequest).AcceptLanguage(acceptLanguage).Execute()

Create Tax



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
    orderTaxRequest := *openapiclient.NewOrderTaxRequest(int64(100), "testing") // OrderTaxRequest | requested field for a taxes
    acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxesApi.OrdersCreateTaxes(context.Background(), id).OrderTaxRequest(orderTaxRequest).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxesApi.OrdersCreateTaxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrdersCreateTaxes`: UpdateOrderTaxResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxesApi.OrdersCreateTaxes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrdersCreateTaxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderTaxRequest** | [**OrderTaxRequest**](OrderTaxRequest.md) | requested field for a taxes | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]

### Return type

[**UpdateOrderTaxResponse**](UpdateOrderTaxResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrdersDeleteTaxes

> UpdateOrderTaxResponse OrdersDeleteTaxes(ctx, id, taxId).AcceptLanguage(acceptLanguage).Execute()

Delete Tax



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
    taxId := "tax_lin_2tQ974hSHcsdeSZHG" // string | identifier
    acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxesApi.OrdersDeleteTaxes(context.Background(), id, taxId).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxesApi.OrdersDeleteTaxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrdersDeleteTaxes`: UpdateOrderTaxResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxesApi.OrdersDeleteTaxes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 
**taxId** | **string** | identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrdersDeleteTaxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]

### Return type

[**UpdateOrderTaxResponse**](UpdateOrderTaxResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrdersUpdateTaxes

> UpdateOrderTaxResponse OrdersUpdateTaxes(ctx, id, taxId).UpdateOrderTaxRequest(updateOrderTaxRequest).AcceptLanguage(acceptLanguage).Execute()

Update Tax



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
    taxId := "tax_lin_2tQ974hSHcsdeSZHG" // string | identifier
    updateOrderTaxRequest := *openapiclient.NewUpdateOrderTaxRequest() // UpdateOrderTaxRequest | requested field for taxes
    acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxesApi.OrdersUpdateTaxes(context.Background(), id, taxId).UpdateOrderTaxRequest(updateOrderTaxRequest).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxesApi.OrdersUpdateTaxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrdersUpdateTaxes`: UpdateOrderTaxResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxesApi.OrdersUpdateTaxes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 
**taxId** | **string** | identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrdersUpdateTaxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrderTaxRequest** | [**UpdateOrderTaxRequest**](UpdateOrderTaxRequest.md) | requested field for taxes | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]

### Return type

[**UpdateOrderTaxResponse**](UpdateOrderTaxResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

