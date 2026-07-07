# \SubscriptionsCustomerPortalAPI

All URIs are relative to *https://api.conekta.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomerPortal**](SubscriptionsCustomerPortalAPI.md#CreateCustomerPortal) | **Post** /subscriptions/{subscription_id}/customer_portal | Create customer portal
[**GetCustomerPortal**](SubscriptionsCustomerPortalAPI.md#GetCustomerPortal) | **Get** /subscriptions/{subscription_id}/customer_portal | Get customer portal



## CreateCustomerPortal

> CustomerPortalResponse CreateCustomerPortal(ctx, subscriptionId).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Create customer portal



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
	subscriptionId := "sub_2tGzG1GxtDAZHEGPH" // string | Identifier of the subscription resource
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsCustomerPortalAPI.CreateCustomerPortal(context.Background(), subscriptionId).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsCustomerPortalAPI.CreateCustomerPortal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomerPortal`: CustomerPortalResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsCustomerPortalAPI.CreateCustomerPortal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerPortalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**CustomerPortalResponse**](CustomerPortalResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerPortal

> CustomerPortalResponse GetCustomerPortal(ctx, subscriptionId).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Get customer portal



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
	subscriptionId := "sub_2tGzG1GxtDAZHEGPH" // string | Identifier of the subscription resource
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsCustomerPortalAPI.GetCustomerPortal(context.Background(), subscriptionId).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsCustomerPortalAPI.GetCustomerPortal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerPortal`: CustomerPortalResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsCustomerPortalAPI.GetCustomerPortal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerPortalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**CustomerPortalResponse**](CustomerPortalResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

