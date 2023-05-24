# \ApiKeysApi

All URIs are relative to *https://api.conekta.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiKey**](ApiKeysApi.md#CreateApiKey) | **Post** /api_keys | Create Api Key
[**DeleteApiKey**](ApiKeysApi.md#DeleteApiKey) | **Delete** /api_keys/{id} | Delete Api Key
[**GetApiKey**](ApiKeysApi.md#GetApiKey) | **Get** /api_keys/{id} | Get Api Key
[**GetApiKeys**](ApiKeysApi.md#GetApiKeys) | **Get** /api_keys | Get list of Api Keys
[**UpdateApiKey**](ApiKeysApi.md#UpdateApiKey) | **Put** /api_keys/{id} | Update Api Key



## CreateApiKey

> ApiKeyCreateResponse CreateApiKey(ctx).ApiKeyRequest(apiKeyRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Create Api Key



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
    apiKeyRequest := *openapiclient.NewApiKeyRequest(true, "online store", "private") // ApiKeyRequest | requested field for a api keys
    acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")
    xChildCompanyId := "6441b6376b60c3a638da80af" // string | In the case of a holding company, the company id of the child company to which will process the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiKeysApi.CreateApiKey(context.Background()).ApiKeyRequest(apiKeyRequest).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKeysApi.CreateApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiKey`: ApiKeyCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `ApiKeysApi.CreateApiKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKeyRequest** | [**ApiKeyRequest**](ApiKeyRequest.md) | requested field for a api keys | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**ApiKeyCreateResponse**](ApiKeyCreateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiKey

> DeleteApiKeysResponse DeleteApiKey(ctx, id).AcceptLanguage(acceptLanguage).Execute()

Delete Api Key



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
    resp, r, err := apiClient.ApiKeysApi.DeleteApiKey(context.Background(), id).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKeysApi.DeleteApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteApiKey`: DeleteApiKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `ApiKeysApi.DeleteApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]

### Return type

[**DeleteApiKeysResponse**](DeleteApiKeysResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiKey

> ApiKeyResponse GetApiKey(ctx, id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()

Get Api Key



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
    resp, r, err := apiClient.ApiKeysApi.GetApiKey(context.Background(), id).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKeysApi.GetApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiKey`: ApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `ApiKeysApi.GetApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 

### Return type

[**ApiKeyResponse**](ApiKeyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiKeys

> GetApiKeysResponse GetApiKeys(ctx).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Limit(limit).Search(search).Next(next).Previous(previous).Execute()

Get list of Api Keys



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiKeysApi.GetApiKeys(context.Background()).AcceptLanguage(acceptLanguage).XChildCompanyId(xChildCompanyId).Limit(limit).Search(search).Next(next).Previous(previous).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKeysApi.GetApiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiKeys`: GetApiKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `ApiKeysApi.GetApiKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **xChildCompanyId** | **string** | In the case of a holding company, the company id of the child company to which will process the request. | 
 **limit** | **int32** | The numbers of items to return, the maximum value is 250 | [default to 20]
 **search** | **string** | General order search, e.g. by mail, reference etc. | 
 **next** | **string** | next page | 
 **previous** | **string** | previous page | 

### Return type

[**GetApiKeysResponse**](GetApiKeysResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiKey

> ApiKeyResponse UpdateApiKey(ctx, id).AcceptLanguage(acceptLanguage).ApiKeyUpdateRequest(apiKeyUpdateRequest).Execute()

Update Api Key



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
    apiKeyUpdateRequest := *openapiclient.NewApiKeyUpdateRequest() // ApiKeyUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApiKeysApi.UpdateApiKey(context.Background(), id).AcceptLanguage(acceptLanguage).ApiKeyUpdateRequest(apiKeyUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKeysApi.UpdateApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiKey`: ApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `ApiKeysApi.UpdateApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **apiKeyUpdateRequest** | [**ApiKeyUpdateRequest**](ApiKeyUpdateRequest.md) |  | 

### Return type

[**ApiKeyResponse**](ApiKeyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

