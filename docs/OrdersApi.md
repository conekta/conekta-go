# \OrdersAPI

All URIs are relative to *https://api.conekta.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelOrder**](OrdersAPI.md#CancelOrder) | **Post** /orders/{id}/cancel | Cancel Order
[**CreateOrder**](OrdersAPI.md#CreateOrder) | **Post** /orders | Create order
[**GetOrderById**](OrdersAPI.md#GetOrderById) | **Get** /orders/{id} | Get Order
[**GetOrders**](OrdersAPI.md#GetOrders) | **Get** /orders | Get a list of Orders
[**OrderCancelRefund**](OrdersAPI.md#OrderCancelRefund) | **Delete** /orders/{id}/refunds/{refund_id} | Cancel Refund
[**OrderRefund**](OrdersAPI.md#OrderRefund) | **Post** /orders/{id}/refunds | Refund Order
[**OrdersCreateCapture**](OrdersAPI.md#OrdersCreateCapture) | **Post** /orders/{id}/capture | Capture Order
[**UpdateOrder**](OrdersAPI.md#UpdateOrder) | **Put** /orders/{id} | Update Order



## CancelOrder

> OrderResponse CancelOrder(ctx, id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Cancel Order



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
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.CancelOrder(context.Background(), id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.CancelOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelOrder`: OrderResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.CancelOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**OrderResponse**](OrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrder

> OrderResponse CreateOrder(ctx).OrderRequest(orderRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

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
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.CreateOrder(context.Background()).OrderRequest(orderRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.CreateOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrder`: OrderResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.CreateOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderRequest** | [**OrderRequest**](OrderRequest.md) | requested field for order | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**OrderResponse**](OrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrderById

> OrderResponse GetOrderById(ctx, id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

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
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.GetOrderById(context.Background(), id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.GetOrderById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrderById`: OrderResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.GetOrderById`: %v\n", resp)
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
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**OrderResponse**](OrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrders

> GetOrdersResponse GetOrders(ctx).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Limit(limit).Search(search).Next(next).Previous(previous).PaymentStatus(paymentStatus).LastPaymentInfoStatus(lastPaymentInfoStatus).CreatedAt(createdAt).CreatedAtGte(createdAtGte).CreatedAtLte(createdAtLte).UpdatedAtGte(updatedAtGte).UpdatedAtLte(updatedAtLte).Execute()

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
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)
	limit := int32(56) // int32 | The numbers of items to return, the maximum value is 250 (optional) (default to 20)
	search := "search_example" // string | General order search, e.g. by mail, reference etc. (optional)
	next := "next_example" // string | next page (optional)
	previous := "previous_example" // string | previous page (optional)
	paymentStatus := "paid" // string | Filters by order status (optional)
	lastPaymentInfoStatus := "pending_payment" // string | Filters by last payment info status (optional)
	createdAt := int64(1612137600) // int64 | created equal to (optional)
	createdAtGte := int64(1612137600) // int64 | created at greater than or equal to (optional)
	createdAtLte := int64(1612137600) // int64 | created at less than or equal to (optional)
	updatedAtGte := int64(1612137600) // int64 | updated at greater than or equal to (optional)
	updatedAtLte := int64(1612137600) // int64 | updated at less than or equal to (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.GetOrders(context.Background()).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Limit(limit).Search(search).Next(next).Previous(previous).PaymentStatus(paymentStatus).LastPaymentInfoStatus(lastPaymentInfoStatus).CreatedAt(createdAt).CreatedAtGte(createdAtGte).CreatedAtLte(createdAtLte).UpdatedAtGte(updatedAtGte).UpdatedAtLte(updatedAtLte).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.GetOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrders`: GetOrdersResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.GetOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 
 **limit** | **int32** | The numbers of items to return, the maximum value is 250 | [default to 20]
 **search** | **string** | General order search, e.g. by mail, reference etc. | 
 **next** | **string** | next page | 
 **previous** | **string** | previous page | 
 **paymentStatus** | **string** | Filters by order status | 
 **lastPaymentInfoStatus** | **string** | Filters by last payment info status | 
 **createdAt** | **int64** | created equal to | 
 **createdAtGte** | **int64** | created at greater than or equal to | 
 **createdAtLte** | **int64** | created at less than or equal to | 
 **updatedAtGte** | **int64** | updated at greater than or equal to | 
 **updatedAtLte** | **int64** | updated at less than or equal to | 

### Return type

[**GetOrdersResponse**](GetOrdersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderCancelRefund

> OrderResponse OrderCancelRefund(ctx, id, refundId).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Cancel Refund



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
	refundId := "6407b5bee1329a000175ba11" // string | refund identifier
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.OrderCancelRefund(context.Background(), id, refundId).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.OrderCancelRefund``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderCancelRefund`: OrderResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.OrderCancelRefund`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 
**refundId** | **string** | refund identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderCancelRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**OrderResponse**](OrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderRefund

> OrderResponse OrderRefund(ctx, id).OrderRefundRequest(orderRefundRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

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
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.OrderRefund(context.Background(), id).OrderRefundRequest(orderRefundRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.OrderRefund``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrderRefund`: OrderResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.OrderRefund`: %v\n", resp)
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
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**OrderResponse**](OrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrdersCreateCapture

> OrderResponse OrdersCreateCapture(ctx, id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).OrderCaptureRequest(orderCaptureRequest).Execute()

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
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)
	orderCaptureRequest := *openapiclient.NewOrderCaptureRequest(int64(500)) // OrderCaptureRequest | requested fields for capture order (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.OrdersCreateCapture(context.Background(), id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).OrderCaptureRequest(orderCaptureRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.OrdersCreateCapture``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrdersCreateCapture`: OrderResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.OrdersCreateCapture`: %v\n", resp)
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
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 
 **orderCaptureRequest** | [**OrderCaptureRequest**](OrderCaptureRequest.md) | requested fields for capture order | 

### Return type

[**OrderResponse**](OrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.2.0+json

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
	resp, r, err := apiClient.OrdersAPI.UpdateOrder(context.Background(), id).OrderUpdateRequest(orderUpdateRequest).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.UpdateOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrder`: OrderResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.UpdateOrder`: %v\n", resp)
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
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

