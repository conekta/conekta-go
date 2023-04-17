# \BlacklistApi

All URIs are relative to *https://api.conekta.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewBlacklistRule**](BlacklistApi.md#CreateNewBlacklistRule) | **Post** /antifraud/blacklists | Create a blacklisted rule
[**DeleteBlacklistRule**](BlacklistApi.md#DeleteBlacklistRule) | **Delete** /antifraud/blacklists | Delete a blacklisted rule
[**GetBlacklist**](BlacklistApi.md#GetBlacklist) | **Get** /antifraud/blacklists | Get a list of blacklisted rules



## CreateNewBlacklistRule

> BlacklistRuleResponse CreateNewBlacklistRule(ctx).CreateRiskRulesData(createRiskRulesData).Execute()

Create a blacklisted rule

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
    createRiskRulesData := *openapiclient.NewCreateRiskRulesData() // CreateRiskRulesData | requested field for blacklist rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlacklistApi.CreateNewBlacklistRule(context.Background()).CreateRiskRulesData(createRiskRulesData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlacklistApi.CreateNewBlacklistRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNewBlacklistRule`: BlacklistRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `BlacklistApi.CreateNewBlacklistRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewBlacklistRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRiskRulesData** | [**CreateRiskRulesData**](CreateRiskRulesData.md) | requested field for blacklist rule | 

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


## DeleteBlacklistRule

> DeletedBlacklistRuleResponse DeleteBlacklistRule(ctx).Execute()

Delete a blacklisted rule

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlacklistApi.DeleteBlacklistRule(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlacklistApi.DeleteBlacklistRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBlacklistRule`: DeletedBlacklistRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `BlacklistApi.DeleteBlacklistRule`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlacklistRuleRequest struct via the builder pattern


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


## GetBlacklist

> RiskRulesList GetBlacklist(ctx).Execute()

Get a list of blacklisted rules



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlacklistApi.GetBlacklist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlacklistApi.GetBlacklist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlacklist`: RiskRulesList
    fmt.Fprintf(os.Stdout, "Response from `BlacklistApi.GetBlacklist`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlacklistRequest struct via the builder pattern


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

