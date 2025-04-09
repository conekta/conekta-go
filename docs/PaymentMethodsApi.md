# \PaymentMethodsAPI

All URIs are relative to *https://api.conekta.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomerPaymentMethods**](PaymentMethodsAPI.md#CreateCustomerPaymentMethods) | **Post** /customers/{id}/payment_sources | Create Payment Method
[**DeleteCustomerPaymentMethods**](PaymentMethodsAPI.md#DeleteCustomerPaymentMethods) | **Delete** /customers/{id}/payment_sources/{payment_method_id} | Delete Payment Method
[**GetCustomerPaymentMethods**](PaymentMethodsAPI.md#GetCustomerPaymentMethods) | **Get** /customers/{id}/payment_sources | Get Payment Methods
[**UpdateCustomerPaymentMethods**](PaymentMethodsAPI.md#UpdateCustomerPaymentMethods) | **Put** /customers/{id}/payment_sources/{payment_method_id} | Update Payment Method



## CreateCustomerPaymentMethods

> CreateCustomerPaymentMethodsResponse CreateCustomerPaymentMethods(ctx, id).CreateCustomerPaymentMethodsRequest(createCustomerPaymentMethodsRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Create Payment Method



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
	createCustomerPaymentMethodsRequest := openapiclient.CreateCustomerPaymentMethods_request{PaymentMethodCashRequest: openapiclient.NewPaymentMethodCashRequest("card | cash | spei | bnpl")} // CreateCustomerPaymentMethodsRequest | requested field for customer payment methods
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentMethodsAPI.CreateCustomerPaymentMethods(context.Background(), id).CreateCustomerPaymentMethodsRequest(createCustomerPaymentMethodsRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsAPI.CreateCustomerPaymentMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomerPaymentMethods`: CreateCustomerPaymentMethodsResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsAPI.CreateCustomerPaymentMethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerPaymentMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createCustomerPaymentMethodsRequest** | [**CreateCustomerPaymentMethodsRequest**](CreateCustomerPaymentMethodsRequest.md) | requested field for customer payment methods | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**CreateCustomerPaymentMethodsResponse**](CreateCustomerPaymentMethodsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomerPaymentMethods

> UpdateCustomerPaymentMethodsResponse DeleteCustomerPaymentMethods(ctx, id, paymentMethodId).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Delete Payment Method



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
	paymentMethodId := "src_2tQ974hSHcsdeSZHG" // string | Identifier of the payment method
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentMethodsAPI.DeleteCustomerPaymentMethods(context.Background(), id, paymentMethodId).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsAPI.DeleteCustomerPaymentMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCustomerPaymentMethods`: UpdateCustomerPaymentMethodsResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsAPI.DeleteCustomerPaymentMethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 
**paymentMethodId** | **string** | Identifier of the payment method | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomerPaymentMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**UpdateCustomerPaymentMethodsResponse**](UpdateCustomerPaymentMethodsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerPaymentMethods

> GetPaymentMethodResponse GetCustomerPaymentMethods(ctx, id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Limit(limit).Next(next).Previous(previous).Search(search).Execute()

Get Payment Methods



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
	limit := int32(56) // int32 | The numbers of items to return, the maximum value is 250 (optional) (default to 20)
	next := "next_example" // string | next page (optional)
	previous := "previous_example" // string | previous page (optional)
	search := "search_example" // string | General order search, e.g. by mail, reference etc. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentMethodsAPI.GetCustomerPaymentMethods(context.Background(), id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Limit(limit).Next(next).Previous(previous).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsAPI.GetCustomerPaymentMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerPaymentMethods`: GetPaymentMethodResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsAPI.GetCustomerPaymentMethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerPaymentMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 
 **limit** | **int32** | The numbers of items to return, the maximum value is 250 | [default to 20]
 **next** | **string** | next page | 
 **previous** | **string** | previous page | 
 **search** | **string** | General order search, e.g. by mail, reference etc. | 

### Return type

[**GetPaymentMethodResponse**](GetPaymentMethodResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomerPaymentMethods

> UpdateCustomerPaymentMethodsResponse UpdateCustomerPaymentMethods(ctx, id, paymentMethodId).UpdatePaymentMethods(updatePaymentMethods).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Update Payment Method



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
	paymentMethodId := "src_2tQ974hSHcsdeSZHG" // string | Identifier of the payment method
	updatePaymentMethods := *openapiclient.NewUpdatePaymentMethods() // UpdatePaymentMethods | requested field for customer payment methods
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentMethodsAPI.UpdateCustomerPaymentMethods(context.Background(), id, paymentMethodId).UpdatePaymentMethods(updatePaymentMethods).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentMethodsAPI.UpdateCustomerPaymentMethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomerPaymentMethods`: UpdateCustomerPaymentMethodsResponse
	fmt.Fprintf(os.Stdout, "Response from `PaymentMethodsAPI.UpdateCustomerPaymentMethods`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 
**paymentMethodId** | **string** | Identifier of the payment method | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerPaymentMethodsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updatePaymentMethods** | [**UpdatePaymentMethods**](UpdatePaymentMethods.md) | requested field for customer payment methods | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**UpdateCustomerPaymentMethodsResponse**](UpdateCustomerPaymentMethodsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

