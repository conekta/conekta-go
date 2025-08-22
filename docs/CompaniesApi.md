# \CompaniesAPI

All URIs are relative to *https://api.conekta.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCompany**](CompaniesAPI.md#CreateCompany) | **Post** /companies | Create Company
[**GetCompanies**](CompaniesAPI.md#GetCompanies) | **Get** /companies | Get List of Companies
[**GetCompany**](CompaniesAPI.md#GetCompany) | **Get** /companies/{id} | Get Company
[**GetCompanyDocuments**](CompaniesAPI.md#GetCompanyDocuments) | **Get** /companies/{company_id}/documents | Get Company Documents
[**UpdateCompanyDocument**](CompaniesAPI.md#UpdateCompanyDocument) | **Patch** /companies/{company_id}/document | Update Company Document
[**UploadCompanyDocument**](CompaniesAPI.md#UploadCompanyDocument) | **Post** /companies/{company_id}/document | Upload Company Document



## CreateCompany

> CompanyResponse CreateCompany(ctx).CreateCompanyRequest(createCompanyRequest).Execute()

Create Company



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
	createCompanyRequest := *openapiclient.NewCreateCompanyRequest() // CreateCompanyRequest | Company data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompaniesAPI.CreateCompany(context.Background()).CreateCompanyRequest(createCompanyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.CreateCompany``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCompany`: CompanyResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.CreateCompany`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCompanyRequest** | [**CreateCompanyRequest**](CreateCompanyRequest.md) | Company data | 

### Return type

[**CompanyResponse**](CompanyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
- **Accept**: application/vnd.conekta-v2.2.0+json

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
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyDocuments

> []CompanyDocumentResponse GetCompanyDocuments(ctx, companyId).AcceptLanguage(acceptLanguage).Execute()

Get Company Documents



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
	companyId := "6307a60c41de27127515a575" // string | The unique identifier of the company.
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompaniesAPI.GetCompanyDocuments(context.Background(), companyId).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.GetCompanyDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyDocuments`: []CompanyDocumentResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.GetCompanyDocuments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]

### Return type

[**[]CompanyDocumentResponse**](CompanyDocumentResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCompanyDocument

> CompanyDocumentResponse UpdateCompanyDocument(ctx, companyId).CompanyDocumentRequest(companyDocumentRequest).AcceptLanguage(acceptLanguage).Execute()

Update Company Document



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
	companyId := "6827206b1ec60400015eb09a" // string | The unique identifier of the company.
	companyDocumentRequest := *openapiclient.NewCompanyDocumentRequest("id_legal_representative", "application/pdf", "example_document.pdf", string([B@51f1e931)) // CompanyDocumentRequest | Document information to update.
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompaniesAPI.UpdateCompanyDocument(context.Background(), companyId).CompanyDocumentRequest(companyDocumentRequest).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.UpdateCompanyDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCompanyDocument`: CompanyDocumentResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.UpdateCompanyDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCompanyDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyDocumentRequest** | [**CompanyDocumentRequest**](CompanyDocumentRequest.md) | Document information to update. | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]

### Return type

[**CompanyDocumentResponse**](CompanyDocumentResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadCompanyDocument

> CompanyDocumentResponse UploadCompanyDocument(ctx, companyId).CompanyDocumentRequest(companyDocumentRequest).AcceptLanguage(acceptLanguage).Execute()

Upload Company Document



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
	companyId := "6827206b1ec60400015eb09a" // string | The unique identifier of the company.
	companyDocumentRequest := *openapiclient.NewCompanyDocumentRequest("id_legal_representative", "application/pdf", "example_document.pdf", string([B@51f1e931)) // CompanyDocumentRequest | Document information to upload.
	acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CompaniesAPI.UploadCompanyDocument(context.Background(), companyId).CompanyDocumentRequest(companyDocumentRequest).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompaniesAPI.UploadCompanyDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadCompanyDocument`: CompanyDocumentResponse
	fmt.Fprintf(os.Stdout, "Response from `CompaniesAPI.UploadCompanyDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** | The unique identifier of the company. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadCompanyDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyDocumentRequest** | [**CompanyDocumentRequest**](CompanyDocumentRequest.md) | Document information to upload. | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]

### Return type

[**CompanyDocumentResponse**](CompanyDocumentResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.2.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

