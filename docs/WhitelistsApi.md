# \WhitelistsApi

All URIs are relative to *https://api.conekta.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewRuleWhitelist**](WhitelistsApi.md#CreateNewRuleWhitelist) | **Post** /antifraud/whitelists | Create a whitelisted rule
[**DeleteRuleWhitelist**](WhitelistsApi.md#DeleteRuleWhitelist) | **Delete** /antifraud/whitelists | Delete a whitelisted rule
[**GetWhiteList**](WhitelistsApi.md#GetWhiteList) | **Get** /antifraud/whitelists | Get a list of whitelisted rules



## CreateNewRuleWhitelist

> CreateNewRuleWhitelist(ctx).CreateRiskRulesData(createRiskRulesData).Execute()

Create a whitelisted rule

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
    createRiskRulesData := *openapiclient.NewCreateRiskRulesData() // CreateRiskRulesData |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WhitelistsApi.CreateNewRuleWhitelist(context.Background()).CreateRiskRulesData(createRiskRulesData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhitelistsApi.CreateNewRuleWhitelist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewRuleWhitelistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRiskRulesData** | [**CreateRiskRulesData**](CreateRiskRulesData.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRuleWhitelist

> DeleteRuleWhitelist(ctx).Execute()

Delete a whitelisted rule

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
    r, err := apiClient.WhitelistsApi.DeleteRuleWhitelist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhitelistsApi.DeleteRuleWhitelist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleWhitelistRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWhiteList

> RiskRulesList GetWhiteList(ctx).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WhitelistsApi.GetWhiteList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WhitelistsApi.GetWhiteList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWhiteList`: RiskRulesList
    fmt.Fprintf(os.Stdout, "Response from `WhitelistsApi.GetWhiteList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWhiteListRequest struct via the builder pattern


### Return type

[**RiskRulesList**](RiskRulesList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

