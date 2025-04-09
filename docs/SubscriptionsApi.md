# \SubscriptionsAPI

All URIs are relative to *https://api.conekta.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelSubscription**](SubscriptionsAPI.md#CancelSubscription) | **Post** /customers/{id}/subscription/cancel | Cancel Subscription [Deprecated]
[**CreateSubscription**](SubscriptionsAPI.md#CreateSubscription) | **Post** /customers/{id}/subscription | Create Subscription [Deprecated]
[**GetSubscription**](SubscriptionsAPI.md#GetSubscription) | **Get** /customers/{id}/subscription | Get Subscription [Deprecated]
[**GetSubscriptionEvents**](SubscriptionsAPI.md#GetSubscriptionEvents) | **Get** /customers/{id}/subscription/events | Get Subscription Events [Deprecated]
[**PauseSubscription**](SubscriptionsAPI.md#PauseSubscription) | **Post** /customers/{id}/subscription/pause | Pause Subscription [Deprecated]
[**ResumeSubscription**](SubscriptionsAPI.md#ResumeSubscription) | **Post** /customers/{id}/subscription/resume | Resume Subscription [Deprecated]
[**SubscriptionCancel**](SubscriptionsAPI.md#SubscriptionCancel) | **Post** /customers/{customer_id}/subscriptions/{id}/cancel | Cancel Subscription
[**SubscriptionCreate**](SubscriptionsAPI.md#SubscriptionCreate) | **Post** /customers/{customer_id}/subscriptions | Create Subscription
[**SubscriptionEvents**](SubscriptionsAPI.md#SubscriptionEvents) | **Get** /customers/{customer_id}/subscriptions/{id}/events | Get Subscription Events
[**SubscriptionList**](SubscriptionsAPI.md#SubscriptionList) | **Get** /customers/{customer_id}/subscriptions | List Subscriptions
[**SubscriptionPause**](SubscriptionsAPI.md#SubscriptionPause) | **Post** /customers/{customer_id}/subscriptions/{id}/pause | Pause Subscription
[**SubscriptionResume**](SubscriptionsAPI.md#SubscriptionResume) | **Post** /customers/{customer_id}/subscriptions/{id}/resume | Resume Subscription
[**SubscriptionUpdate**](SubscriptionsAPI.md#SubscriptionUpdate) | **Put** /customers/{customer_id}/subscriptions/{id} | Update Subscription
[**SubscriptionsGet**](SubscriptionsAPI.md#SubscriptionsGet) | **Get** /customers/{customer_id}/subscriptions/{id} | Get Subscription
[**SubscriptionsRetry**](SubscriptionsAPI.md#SubscriptionsRetry) | **Post** /customers/{customer_id}/subscriptions/{id}/retry | Retry Failed Payment
[**UpdateSubscription**](SubscriptionsAPI.md#UpdateSubscription) | **Put** /customers/{id}/subscription | Update Subscription [Deprecated]



## CancelSubscription

> SubscriptionResponse CancelSubscription(ctx, id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Cancel Subscription [Deprecated]



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
	resp, r, err := apiClient.SubscriptionsAPI.CancelSubscription(context.Background(), id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.CancelSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelSubscription`: SubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.CancelSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubscription

> SubscriptionResponse CreateSubscription(ctx, id).SubscriptionRequest(subscriptionRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Create Subscription [Deprecated]



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
	subscriptionRequest := *openapiclient.NewSubscriptionRequest("f84gdgf5g48r15fd21g8w424fd1") // SubscriptionRequest | requested field for subscriptions
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.CreateSubscription(context.Background(), id).SubscriptionRequest(subscriptionRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.CreateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubscription`: SubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.CreateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionRequest** | [**SubscriptionRequest**](SubscriptionRequest.md) | requested field for subscriptions | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscription

> SubscriptionResponse GetSubscription(ctx, id).AcceptLanguage(acceptLanguage).Execute()

Get Subscription [Deprecated]



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
	resp, r, err := apiClient.SubscriptionsAPI.GetSubscription(context.Background(), id).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.GetSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscription`: SubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.GetSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriptionEvents

> SubscriptionEventsResponse GetSubscriptionEvents(ctx, id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Get Subscription Events [Deprecated]



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
	resp, r, err := apiClient.SubscriptionsAPI.GetSubscriptionEvents(context.Background(), id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.GetSubscriptionEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubscriptionEvents`: SubscriptionEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.GetSubscriptionEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**SubscriptionEventsResponse**](SubscriptionEventsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseSubscription

> SubscriptionResponse PauseSubscription(ctx, id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Pause Subscription [Deprecated]



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
	resp, r, err := apiClient.SubscriptionsAPI.PauseSubscription(context.Background(), id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.PauseSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PauseSubscription`: SubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.PauseSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeSubscription

> SubscriptionResponse ResumeSubscription(ctx, id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Resume Subscription [Deprecated]



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
	resp, r, err := apiClient.SubscriptionsAPI.ResumeSubscription(context.Background(), id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.ResumeSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResumeSubscription`: SubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.ResumeSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionCancel

> SubscriptionResponse SubscriptionCancel(ctx, customerId, id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Cancel Subscription



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
	customerId := "cus_2tGzG1GxtDAZHEGPH" // string | Identifier of the customer resource
	id := "sub_2tGzG1GxtDAZHEGPH" // string | Identifier of the subscription resource
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionCancel(context.Background(), customerId, id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionCancel`: SubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Identifier of the customer resource | 
**id** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionCreate

> SubscriptionResponse SubscriptionCreate(ctx, customerId).SubscriptionRequest(subscriptionRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Create Subscription



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
	customerId := "cus_2tGzG1GxtDAZHEGPH" // string | Identifier of the customer resource
	subscriptionRequest := *openapiclient.NewSubscriptionRequest("f84gdgf5g48r15fd21g8w424fd1") // SubscriptionRequest | requested field for subscriptions
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionCreate(context.Background(), customerId).SubscriptionRequest(subscriptionRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionCreate`: SubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Identifier of the customer resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionRequest** | [**SubscriptionRequest**](SubscriptionRequest.md) | requested field for subscriptions | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionEvents

> SubscriptionEventsResponse SubscriptionEvents(ctx, customerId, id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Limit(limit).Search(search).Next(next).Previous(previous).Execute()

Get Subscription Events



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
	customerId := "cus_2tGzG1GxtDAZHEGPH" // string | Identifier of the customer resource
	id := "sub_2tGzG1GxtDAZHEGPH" // string | Identifier of the subscription resource
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)
	limit := int32(56) // int32 | The numbers of items to return, the maximum value is 250 (optional) (default to 20)
	search := "search_example" // string | General order search, e.g. by mail, reference etc. (optional)
	next := "next_example" // string | next page (optional)
	previous := "previous_example" // string | previous page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionEvents(context.Background(), customerId, id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Limit(limit).Search(search).Next(next).Previous(previous).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionEvents`: SubscriptionEventsResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Identifier of the customer resource | 
**id** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 
 **limit** | **int32** | The numbers of items to return, the maximum value is 250 | [default to 20]
 **search** | **string** | General order search, e.g. by mail, reference etc. | 
 **next** | **string** | next page | 
 **previous** | **string** | previous page | 

### Return type

[**SubscriptionEventsResponse**](SubscriptionEventsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionList

> SubscriptionResponse SubscriptionList(ctx, customerId).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Limit(limit).Search(search).Next(next).Previous(previous).Execute()

List Subscriptions



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
	customerId := "cus_2tGzG1GxtDAZHEGPH" // string | Identifier of the customer resource
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)
	limit := int32(56) // int32 | The numbers of items to return, the maximum value is 250 (optional) (default to 20)
	search := "search_example" // string | General order search, e.g. by mail, reference etc. (optional)
	next := "next_example" // string | next page (optional)
	previous := "previous_example" // string | previous page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionList(context.Background(), customerId).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Limit(limit).Search(search).Next(next).Previous(previous).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionList`: SubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Identifier of the customer resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 
 **limit** | **int32** | The numbers of items to return, the maximum value is 250 | [default to 20]
 **search** | **string** | General order search, e.g. by mail, reference etc. | 
 **next** | **string** | next page | 
 **previous** | **string** | previous page | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionPause

> SubscriptionResponse SubscriptionPause(ctx, customerId, id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Pause Subscription



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
	customerId := "cus_2tGzG1GxtDAZHEGPH" // string | Identifier of the customer resource
	id := "sub_2tGzG1GxtDAZHEGPH" // string | Identifier of the subscription resource
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionPause(context.Background(), customerId, id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionPause``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionPause`: SubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionPause`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Identifier of the customer resource | 
**id** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionPauseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionResume

> SubscriptionResponse SubscriptionResume(ctx, customerId, id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Resume Subscription



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
	customerId := "cus_2tGzG1GxtDAZHEGPH" // string | Identifier of the customer resource
	id := "sub_2tGzG1GxtDAZHEGPH" // string | Identifier of the subscription resource
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionResume(context.Background(), customerId, id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionResume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionResume`: SubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionResume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Identifier of the customer resource | 
**id** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionResumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionUpdate

> SubscriptionResponse SubscriptionUpdate(ctx, customerId, id).SubscriptionUpdateRequest(subscriptionUpdateRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Update Subscription



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
	customerId := "cus_2tGzG1GxtDAZHEGPH" // string | Identifier of the customer resource
	id := "sub_2tGzG1GxtDAZHEGPH" // string | Identifier of the subscription resource
	subscriptionUpdateRequest := *openapiclient.NewSubscriptionUpdateRequest() // SubscriptionUpdateRequest | requested field for update a subscription
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionUpdate(context.Background(), customerId, id).SubscriptionUpdateRequest(subscriptionUpdateRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionUpdate`: SubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Identifier of the customer resource | 
**id** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **subscriptionUpdateRequest** | [**SubscriptionUpdateRequest**](SubscriptionUpdateRequest.md) | requested field for update a subscription | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsGet

> SubscriptionResponse SubscriptionsGet(ctx, customerId, id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Get Subscription



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
	customerId := "cus_2tGzG1GxtDAZHEGPH" // string | Identifier of the customer resource
	id := "sub_2tGzG1GxtDAZHEGPH" // string | Identifier of the subscription resource
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsGet(context.Background(), customerId, id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsGet`: SubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Identifier of the customer resource | 
**id** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsRetry

> SubscriptionResponse SubscriptionsRetry(ctx, customerId, id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Retry Failed Payment



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
	customerId := "cus_2tGzG1GxtDAZHEGPH" // string | Identifier of the customer resource
	id := "sub_2tGzG1GxtDAZHEGPH" // string | Identifier of the subscription resource
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.SubscriptionsRetry(context.Background(), customerId, id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.SubscriptionsRetry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsRetry`: SubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.SubscriptionsRetry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | Identifier of the customer resource | 
**id** | **string** | Identifier of the subscription resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsRetryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSubscription

> SubscriptionResponse UpdateSubscription(ctx, id).SubscriptionUpdateRequest(subscriptionUpdateRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Update Subscription [Deprecated]



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
	subscriptionUpdateRequest := *openapiclient.NewSubscriptionUpdateRequest() // SubscriptionUpdateRequest | requested field for update a subscription
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
	xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionsAPI.UpdateSubscription(context.Background(), id).SubscriptionUpdateRequest(subscriptionUpdateRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.UpdateSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSubscription`: SubscriptionResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.UpdateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionUpdateRequest** | [**SubscriptionUpdateRequest**](SubscriptionUpdateRequest.md) | requested field for update a subscription | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**SubscriptionResponse**](SubscriptionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

