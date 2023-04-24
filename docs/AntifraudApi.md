# \AntifraudApi

All URIs are relative to *https://api.conekta.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRuleBlacklist**](AntifraudApi.md#CreateRuleBlacklist) | **Post** /antifraud/blacklists | Create blacklisted rule
[**CreateRuleWhitelist**](AntifraudApi.md#CreateRuleWhitelist) | **Post** /antifraud/whitelists | Create whitelisted rule
[**DeleteRuleBlacklist**](AntifraudApi.md#DeleteRuleBlacklist) | **Delete** /antifraud/blacklists/{id} | Delete blacklisted rule
[**DeleteRuleWhitelist**](AntifraudApi.md#DeleteRuleWhitelist) | **Delete** /antifraud/whitelists/{id} | Delete whitelisted rule
[**GetRuleBlacklist**](AntifraudApi.md#GetRuleBlacklist) | **Get** /antifraud/blacklists | Get list of blacklisted rules
[**GetRuleWhitelist**](AntifraudApi.md#GetRuleWhitelist) | **Get** /antifraud/whitelists | Get a list of whitelisted rules



## CreateRuleBlacklist

> BlacklistRuleResponse CreateRuleBlacklist(ctx).CreateRiskRulesData(createRiskRulesData).AcceptLanguage(acceptLanguage).Execute()

Create blacklisted rule

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
    createRiskRulesData := *openapiclient.NewCreateRiskRulesData("this client's email was verified at 20/09/22 by internal process", "email | phone | card_token", "email@example.com | 818081808180 | src_2qUCNd5AyQqfPMBuV") // CreateRiskRulesData | requested field for blacklist rule
    acceptLanguage := "es" // string | Use for knowing which language to use (optional) (default to "es")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AntifraudApi.CreateRuleBlacklist(context.Background()).CreateRiskRulesData(createRiskRulesData).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AntifraudApi.CreateRuleBlacklist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRuleBlacklist`: BlacklistRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `AntifraudApi.CreateRuleBlacklist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRuleBlacklistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRiskRulesData** | [**CreateRiskRulesData**](CreateRiskRulesData.md) | requested field for blacklist rule | 
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]

### Return type

[**BlacklistRuleResponse**](BlacklistRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRuleWhitelist

> WhitelistlistRuleResponse CreateRuleWhitelist(ctx).AcceptLanguage(acceptLanguage).CreateRiskRulesData(createRiskRulesData).Execute()

Create whitelisted rule

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
    createRiskRulesData := *openapiclient.NewCreateRiskRulesData("this client's email was verified at 20/09/22 by internal process", "email | phone | card_token", "email@example.com | 818081808180 | src_2qUCNd5AyQqfPMBuV") // CreateRiskRulesData |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AntifraudApi.CreateRuleWhitelist(context.Background()).AcceptLanguage(acceptLanguage).CreateRiskRulesData(createRiskRulesData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AntifraudApi.CreateRuleWhitelist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRuleWhitelist`: WhitelistlistRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `AntifraudApi.CreateRuleWhitelist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRuleWhitelistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]
 **createRiskRulesData** | [**CreateRiskRulesData**](CreateRiskRulesData.md) |  | 

### Return type

[**WhitelistlistRuleResponse**](WhitelistlistRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRuleBlacklist

> DeletedBlacklistRuleResponse DeleteRuleBlacklist(ctx, id).AcceptLanguage(acceptLanguage).Execute()

Delete blacklisted rule

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
    resp, r, err := apiClient.AntifraudApi.DeleteRuleBlacklist(context.Background(), id).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AntifraudApi.DeleteRuleBlacklist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRuleBlacklist`: DeletedBlacklistRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `AntifraudApi.DeleteRuleBlacklist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleBlacklistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]

### Return type

[**DeletedBlacklistRuleResponse**](DeletedBlacklistRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRuleWhitelist

> DeletedWhitelistRuleResponse DeleteRuleWhitelist(ctx, id).AcceptLanguage(acceptLanguage).Execute()

Delete whitelisted rule

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
    resp, r, err := apiClient.AntifraudApi.DeleteRuleWhitelist(context.Background(), id).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AntifraudApi.DeleteRuleWhitelist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRuleWhitelist`: DeletedWhitelistRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `AntifraudApi.DeleteRuleWhitelist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleWhitelistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]

### Return type

[**DeletedWhitelistRuleResponse**](DeletedWhitelistRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleBlacklist

> RiskRulesList GetRuleBlacklist(ctx).AcceptLanguage(acceptLanguage).Execute()

Get list of blacklisted rules



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AntifraudApi.GetRuleBlacklist(context.Background()).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AntifraudApi.GetRuleBlacklist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRuleBlacklist`: RiskRulesList
    fmt.Fprintf(os.Stdout, "Response from `AntifraudApi.GetRuleBlacklist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleBlacklistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]

### Return type

[**RiskRulesList**](RiskRulesList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleWhitelist

> RiskRulesList GetRuleWhitelist(ctx).AcceptLanguage(acceptLanguage).Execute()

Get a list of whitelisted rules



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AntifraudApi.GetRuleWhitelist(context.Background()).AcceptLanguage(acceptLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AntifraudApi.GetRuleWhitelist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRuleWhitelist`: RiskRulesList
    fmt.Fprintf(os.Stdout, "Response from `AntifraudApi.GetRuleWhitelist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleWhitelistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **string** | Use for knowing which language to use | [default to &quot;es&quot;]

### Return type

[**RiskRulesList**](RiskRulesList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.conekta-v2.1.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

