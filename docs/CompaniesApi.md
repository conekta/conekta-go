# \CompaniesAPI

All URIs are relative to *https://api.conekta.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompanies**](CompaniesAPI.md#GetCompanies) | **Get** /companies | Get List of Companies
[**GetCompany**](CompaniesAPI.md#GetCompany) | **Get** /companies/{id} | Get Company



## GetCompanies

> GetCompaniesResponse GetCompanies(ctx).AcceptLanguage(acceptLanguage).Limit(limit).Search(search).Next(next).Previous(previous).Execute()

Get List of Companies



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
    resp, r, err := apiClient.CompaniesAPI.GetCompanies(context.Background()).AcceptLanguage(acceptLanguage).Limit(limit).Search(search).Next(next).Previous(previous).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.GetCompanies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompanies`: GetCompaniesResponse
    fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.GetCompanies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **limit** | **int32** | The numbers of items to return, the maximum value is 250 | [default to 20]
 **search** | **string** | General order search, e.g. by mail, reference etc. | 
 **next** | **string** | next page | 
 **previous** | **string** | previous page | 

### Return type

[**GetCompaniesResponse**](GetCompaniesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompany

> CompanyResponse GetCompany(ctx, id).AcceptLanguage(acceptLanguage).Execute()

Get Company

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
    resp, r, err := apiClient.CompaniesAPI.GetCompany(context.Background(), id).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.GetCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompany`: CompanyResponse
    fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.GetCompany`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]

### Return type

[**CompanyResponse**](CompanyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

