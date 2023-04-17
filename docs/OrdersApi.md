# \OrdersApi

All URIs are relative to *https://api.conekta.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrder**](OrdersApi.md#CreateOrder) | **Post** /orders | Create order
[**GetOrderById**](OrdersApi.md#GetOrderById) | **Get** /orders/{id} | Get Order
[**GetOrders**](OrdersApi.md#GetOrders) | **Get** /orders | Get a list of Orders
[**OrderRefund**](OrdersApi.md#OrderRefund) | **Post** /orders/{id}/refunds | Refund Order
[**OrdersCreateCapture**](OrdersApi.md#OrdersCreateCapture) | **Post** /orders/{id}/capture | Capture Order
[**UpdateOrder**](OrdersApi.md#UpdateOrder) | **Put** /orders/{id} | Update Order



## CreateOrder

> OrderResponse CreateOrder(ctx).OrderRequest(orderRequest).AcceptLanguage(acceptLanguage).Execute()

Create order



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
    orderRequest := *openapiclient.NewOrderRequest("MXN", openapiclient.order_request_customer_info{CustomerInfo: openapiclient.NewCustomerInfo("DevTest", "test@conekta.com", "5522997233")}, []openapiclient.Product{*openapiclient.NewProduct("Box of Cohiba S1s", int32(1), int32(20000))}) // OrderRequest | requested field for order
    acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.CreateOrder(context.Background()).OrderRequest(orderRequest).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrder`: OrderResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderRequest** | [**OrderRequest**](OrderRequest.md) | requested field for order | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]

### Return type

[**OrderResponse**](OrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderById

> OrderResponse GetOrderById(ctx, id).AcceptLanguage(acceptLanguage).Execute()

Get Order



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
    resp, r, err := apiClient.OrdersApi.GetOrderById(context.Background(), id).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrderById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderById`: OrderResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrderById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]

### Return type

[**OrderResponse**](OrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrders

> GetOrdersResponse GetOrders(ctx).AcceptLanguage(acceptLanguage).Limit(limit).Search(search).Next(next).Previous(previous).Execute()

Get a list of Orders



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
    resp, r, err := apiClient.OrdersApi.GetOrders(context.Background()).AcceptLanguage(acceptLanguage).Limit(limit).Search(search).Next(next).Previous(previous).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrders`: GetOrdersResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **limit** | **int32** | The numbers of items to return, the maximum value is 250 | [default to 20]
 **search** | **string** | General order search, e.g. by mail, reference etc. | 
 **next** | **string** | next page | 
 **previous** | **string** | previous page | 

### Return type

[**GetOrdersResponse**](GetOrdersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderRefund

> OrderResponse OrderRefund(ctx, id).OrderRefundRequest(orderRefundRequest).AcceptLanguage(acceptLanguage).Execute()

Refund Order



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
    orderRefundRequest := *openapiclient.NewOrderRefundRequest(int32(500), "suspected_fraud") // OrderRefundRequest | requested field for a refund
    acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.OrderRefund(context.Background(), id).OrderRefundRequest(orderRefundRequest).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.OrderRefund``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderRefund`: OrderResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.OrderRefund`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderRefundRequest** | [**OrderRefundRequest**](OrderRefundRequest.md) | requested field for a refund | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]

### Return type

[**OrderResponse**](OrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrdersCreateCapture

> OrderResponse OrdersCreateCapture(ctx, id).AcceptLanguage(acceptLanguage).OrderCaptureRequest(orderCaptureRequest).Execute()

Capture Order



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
    orderCaptureRequest := *openapiclient.NewOrderCaptureRequest(int64(500)) // OrderCaptureRequest | requested fields for capture order (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.OrdersCreateCapture(context.Background(), id).AcceptLanguage(acceptLanguage).OrderCaptureRequest(orderCaptureRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.OrdersCreateCapture``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrdersCreateCapture`: OrderResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.OrdersCreateCapture`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrdersCreateCaptureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **orderCaptureRequest** | [**OrderCaptureRequest**](OrderCaptureRequest.md) | requested fields for capture order | 

### Return type

[**OrderResponse**](OrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrder

> OrderResponse UpdateOrder(ctx, id).OrderUpdateRequest(orderUpdateRequest).AcceptLanguage(acceptLanguage).Execute()

Update Order



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
    orderUpdateRequest := *openapiclient.NewOrderUpdateRequest() // OrderUpdateRequest | requested field for an order
    acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.UpdateOrder(context.Background(), id).OrderUpdateRequest(orderUpdateRequest).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.UpdateOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrder`: OrderResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.UpdateOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderUpdateRequest** | [**OrderUpdateRequest**](OrderUpdateRequest.md) | requested field for an order | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]

### Return type

[**OrderResponse**](OrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

