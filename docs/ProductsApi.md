# \ProductsApi

All URIs are relative to *https://api.conekta.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrdersCreateProduct**](ProductsApi.md#OrdersCreateProduct) | **Post** /orders/{id}/line_items | Create Product
[**OrdersDeleteProduct**](ProductsApi.md#OrdersDeleteProduct) | **Delete** /orders/{id}/line_items/{line_item_id} | Delete Product
[**OrdersUpdateProduct**](ProductsApi.md#OrdersUpdateProduct) | **Put** /orders/{id}/line_items/{line_item_id} | Update Product



## OrdersCreateProduct

> ProductOrderResponse OrdersCreateProduct(ctx, id).Product(product).AcceptLanguage(acceptLanguage).Execute()

Create Product



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
    product := *openapiclient.NewProduct("Box of Cohiba S1s", int32(1), int32(20000)) // Product | requested field for a product
    acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsApi.OrdersCreateProduct(context.Background(), id).Product(product).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.OrdersCreateProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrdersCreateProduct`: ProductOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.OrdersCreateProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrdersCreateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **product** | [**Product**](Product.md) | requested field for a product | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]

### Return type

[**ProductOrderResponse**](ProductOrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrdersDeleteProduct

> ProductOrderResponse OrdersDeleteProduct(ctx, id, lineItemId).AcceptLanguage(acceptLanguage).Execute()

Delete Product



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
    lineItemId := "line_item_2tQ8HkkfbauaKP9Ho" // string | identifier
    acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsApi.OrdersDeleteProduct(context.Background(), id, lineItemId).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.OrdersDeleteProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrdersDeleteProduct`: ProductOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.OrdersDeleteProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 
**lineItemId** | **string** | identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrdersDeleteProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]

### Return type

[**ProductOrderResponse**](ProductOrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrdersUpdateProduct

> ProductOrderResponse OrdersUpdateProduct(ctx, id, lineItemId).UpdateProduct(updateProduct).AcceptLanguage(acceptLanguage).Execute()

Update Product



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
    lineItemId := "line_item_2tQ8HkkfbauaKP9Ho" // string | identifier
    updateProduct := *openapiclient.NewUpdateProduct() // UpdateProduct | requested field for products
    acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductsApi.OrdersUpdateProduct(context.Background(), id, lineItemId).UpdateProduct(updateProduct).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductsApi.OrdersUpdateProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrdersUpdateProduct`: ProductOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductsApi.OrdersUpdateProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 
**lineItemId** | **string** | identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrdersUpdateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateProduct** | [**UpdateProduct**](UpdateProduct.md) | requested field for products | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]

### Return type

[**ProductOrderResponse**](ProductOrderResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

