# \ShippingContactsAPI

All URIs are relative to *https://api.conekta.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomerShippingContacts**](ShippingContactsAPI.md#CreateCustomerShippingContacts) | **Post** /customers/{id}/shipping_contacts | Create a shipping contacts
[**DeleteCustomerShippingContacts**](ShippingContactsAPI.md#DeleteCustomerShippingContacts) | **Delete** /customers/{id}/shipping_contacts/{shipping_contacts_id} | Delete shipping contacts
[**UpdateCustomerShippingContacts**](ShippingContactsAPI.md#UpdateCustomerShippingContacts) | **Put** /customers/{id}/shipping_contacts/{shipping_contacts_id} | Update shipping contacts



## CreateCustomerShippingContacts

> CustomerShippingContactsResponse CreateCustomerShippingContacts(ctx, id).CustomerShippingContacts(customerShippingContacts).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Create a shipping contacts



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
	customerShippingContacts := *openapiclient.NewCustomerShippingContacts(*openapiclient.NewCustomerShippingContactsAddress()) // CustomerShippingContacts | requested field for customer shippings contacts
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShippingContactsAPI.CreateCustomerShippingContacts(context.Background(), id).CustomerShippingContacts(customerShippingContacts).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShippingContactsAPI.CreateCustomerShippingContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomerShippingContacts`: CustomerShippingContactsResponse
	fmt.Fprintf(os.Stdout, "Response from `ShippingContactsAPI.CreateCustomerShippingContacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerShippingContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerShippingContacts** | [**CustomerShippingContacts**](CustomerShippingContacts.md) | requested field for customer shippings contacts | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**CustomerShippingContactsResponse**](CustomerShippingContactsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomerShippingContacts

> CustomerShippingContactsResponse DeleteCustomerShippingContacts(ctx, id, shippingContactsId).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Delete shipping contacts



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
	shippingContactsId := "6307a60c41de27127515a575" // string | identifier
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShippingContactsAPI.DeleteCustomerShippingContacts(context.Background(), id, shippingContactsId).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShippingContactsAPI.DeleteCustomerShippingContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCustomerShippingContacts`: CustomerShippingContactsResponse
	fmt.Fprintf(os.Stdout, "Response from `ShippingContactsAPI.DeleteCustomerShippingContacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 
**shippingContactsId** | **string** | identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomerShippingContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**CustomerShippingContactsResponse**](CustomerShippingContactsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomerShippingContacts

> CustomerShippingContactsResponse UpdateCustomerShippingContacts(ctx, id, shippingContactsId).CustomerUpdateShippingContacts(customerUpdateShippingContacts).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Update shipping contacts



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
	shippingContactsId := "6307a60c41de27127515a575" // string | identifier
	customerUpdateShippingContacts := *openapiclient.NewCustomerUpdateShippingContacts() // CustomerUpdateShippingContacts | requested field for customer update shippings contacts
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShippingContactsAPI.UpdateCustomerShippingContacts(context.Background(), id, shippingContactsId).CustomerUpdateShippingContacts(customerUpdateShippingContacts).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShippingContactsAPI.UpdateCustomerShippingContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomerShippingContacts`: CustomerShippingContactsResponse
	fmt.Fprintf(os.Stdout, "Response from `ShippingContactsAPI.UpdateCustomerShippingContacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 
**shippingContactsId** | **string** | identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerShippingContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **customerUpdateShippingContacts** | [**CustomerUpdateShippingContacts**](CustomerUpdateShippingContacts.md) | requested field for customer update shippings contacts | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**CustomerShippingContactsResponse**](CustomerShippingContactsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

