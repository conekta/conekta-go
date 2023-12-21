# \ShippingsAPI

All URIs are relative to *https://api.conekta.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrdersCreateShipping**](ShippingsAPI.md#OrdersCreateShipping) | **Post** /orders/{id}/shipping_lines | Create Shipping
[**OrdersDeleteShipping**](ShippingsAPI.md#OrdersDeleteShipping) | **Delete** /orders/{id}/shipping_lines/{shipping_id} | Delete Shipping
[**OrdersUpdateShipping**](ShippingsAPI.md#OrdersUpdateShipping) | **Put** /orders/{id}/shipping_lines/{shipping_id} | Update Shipping



## OrdersCreateShipping

> ShippingOrderResponse OrdersCreateShipping(ctx, id).ShippingRequest(shippingRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Create Shipping



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
    shippingRequest := *openapiclient.NewShippingRequest(int64(100)) // ShippingRequest | requested field for a shipping
    acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
    xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingsAPI.OrdersCreateShipping(context.Background(), id).ShippingRequest(shippingRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingsAPI.OrdersCreateShipping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrdersCreateShipping`: ShippingOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `ShippingsAPI.OrdersCreateShipping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrdersCreateShippingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shippingRequest** | [**ShippingRequest**](ShippingRequest.md) | requested field for a shipping | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**ShippingOrderResponse**](ShippingOrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrdersDeleteShipping

> ShippingOrderResponse OrdersDeleteShipping(ctx, id, shippingId).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Delete Shipping



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
    shippingId := "ship_lin_2tQ974hSHcsdeSZHG" // string | identifier
    acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
    xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingsAPI.OrdersDeleteShipping(context.Background(), id, shippingId).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingsAPI.OrdersDeleteShipping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrdersDeleteShipping`: ShippingOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `ShippingsAPI.OrdersDeleteShipping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 
**shippingId** | **string** | identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrdersDeleteShippingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**ShippingOrderResponse**](ShippingOrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrdersUpdateShipping

> ShippingOrderResponse OrdersUpdateShipping(ctx, id, shippingId).ShippingRequest(shippingRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Update Shipping



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
    shippingId := "ship_lin_2tQ974hSHcsdeSZHG" // string | identifier
    shippingRequest := *openapiclient.NewShippingRequest(int64(100)) // ShippingRequest | requested field for a shipping
    acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
    xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingsAPI.OrdersUpdateShipping(context.Background(), id, shippingId).ShippingRequest(shippingRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingsAPI.OrdersUpdateShipping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrdersUpdateShipping`: ShippingOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `ShippingsAPI.OrdersUpdateShipping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 
**shippingId** | **string** | identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrdersUpdateShippingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shippingRequest** | [**ShippingRequest**](ShippingRequest.md) | requested field for a shipping | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**ShippingOrderResponse**](ShippingOrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

