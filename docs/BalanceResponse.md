# BalanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to [**[]BalanceCommonFielsResponse**](BalanceCommonFielsResponse.md) | The balance&#39;s available | [optional] 
**CashoutRetentionAmount** | Pointer to [**[]BalanceCommonFielsResponse**](BalanceCommonFielsResponse.md) | The balance&#39;s cashout retention amount | [optional] 
**ConektaRetention** | Pointer to [**[]BalanceCommonFielsResponse**](BalanceCommonFielsResponse.md) | The balance&#39;s conekta retention | [optional] 
**Gateway** | Pointer to [**[]BalanceCommonFielsResponse**](BalanceCommonFielsResponse.md) | The balance&#39;s gateway | [optional] 
**Pending** | Pointer to [**[]BalanceCommonFielsResponse**](BalanceCommonFielsResponse.md) | The balance&#39;s pending | [optional] 
**Retained** | Pointer to [**[]BalanceCommonFielsResponse**](BalanceCommonFielsResponse.md) | The balance&#39;s retained | [optional] 
**RetentionAmount** | Pointer to [**[]BalanceCommonFielsResponse**](BalanceCommonFielsResponse.md) | The balance&#39;s retention amount | [optional] 
**TargetCollateralAmount** | Pointer to **map[string]interface{}** | The balance&#39;s target collateral amount | [optional] 
**TargetRetentionAmount** | Pointer to [**[]BalanceCommonFielsResponse**](BalanceCommonFielsResponse.md) | The balance&#39;s target retention amount | [optional] 
**TemporarilyRetained** | Pointer to [**[]BalanceCommonFielsResponse**](BalanceCommonFielsResponse.md) | The balance&#39;s temporarily retained | [optional] 

## Methods

### NewBalanceResponse

`func NewBalanceResponse() *BalanceResponse`

NewBalanceResponse instantiates a new BalanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceResponseWithDefaults

`func NewBalanceResponseWithDefaults() *BalanceResponse`

NewBalanceResponseWithDefaults instantiates a new BalanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *BalanceResponse) GetAvailable() []BalanceCommonFielsResponse`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *BalanceResponse) GetAvailableOk() (*[]BalanceCommonFielsResponse, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *BalanceResponse) SetAvailable(v []BalanceCommonFielsResponse)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *BalanceResponse) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCashoutRetentionAmount

`func (o *BalanceResponse) GetCashoutRetentionAmount() []BalanceCommonFielsResponse`

GetCashoutRetentionAmount returns the CashoutRetentionAmount field if non-nil, zero value otherwise.

### GetCashoutRetentionAmountOk

`func (o *BalanceResponse) GetCashoutRetentionAmountOk() (*[]BalanceCommonFielsResponse, bool)`

GetCashoutRetentionAmountOk returns a tuple with the CashoutRetentionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashoutRetentionAmount

`func (o *BalanceResponse) SetCashoutRetentionAmount(v []BalanceCommonFielsResponse)`

SetCashoutRetentionAmount sets CashoutRetentionAmount field to given value.

### HasCashoutRetentionAmount

`func (o *BalanceResponse) HasCashoutRetentionAmount() bool`

HasCashoutRetentionAmount returns a boolean if a field has been set.

### GetConektaRetention

`func (o *BalanceResponse) GetConektaRetention() []BalanceCommonFielsResponse`

GetConektaRetention returns the ConektaRetention field if non-nil, zero value otherwise.

### GetConektaRetentionOk

`func (o *BalanceResponse) GetConektaRetentionOk() (*[]BalanceCommonFielsResponse, bool)`

GetConektaRetentionOk returns a tuple with the ConektaRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConektaRetention

`func (o *BalanceResponse) SetConektaRetention(v []BalanceCommonFielsResponse)`

SetConektaRetention sets ConektaRetention field to given value.

### HasConektaRetention

`func (o *BalanceResponse) HasConektaRetention() bool`

HasConektaRetention returns a boolean if a field has been set.

### GetGateway

`func (o *BalanceResponse) GetGateway() []BalanceCommonFielsResponse`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *BalanceResponse) GetGatewayOk() (*[]BalanceCommonFielsResponse, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *BalanceResponse) SetGateway(v []BalanceCommonFielsResponse)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *BalanceResponse) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetPending

`func (o *BalanceResponse) GetPending() []BalanceCommonFielsResponse`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *BalanceResponse) GetPendingOk() (*[]BalanceCommonFielsResponse, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *BalanceResponse) SetPending(v []BalanceCommonFielsResponse)`

SetPending sets Pending field to given value.

### HasPending

`func (o *BalanceResponse) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetRetained

`func (o *BalanceResponse) GetRetained() []BalanceCommonFielsResponse`

GetRetained returns the Retained field if non-nil, zero value otherwise.

### GetRetainedOk

`func (o *BalanceResponse) GetRetainedOk() (*[]BalanceCommonFielsResponse, bool)`

GetRetainedOk returns a tuple with the Retained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetained

`func (o *BalanceResponse) SetRetained(v []BalanceCommonFielsResponse)`

SetRetained sets Retained field to given value.

### HasRetained

`func (o *BalanceResponse) HasRetained() bool`

HasRetained returns a boolean if a field has been set.

### GetRetentionAmount

`func (o *BalanceResponse) GetRetentionAmount() []BalanceCommonFielsResponse`

GetRetentionAmount returns the RetentionAmount field if non-nil, zero value otherwise.

### GetRetentionAmountOk

`func (o *BalanceResponse) GetRetentionAmountOk() (*[]BalanceCommonFielsResponse, bool)`

GetRetentionAmountOk returns a tuple with the RetentionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionAmount

`func (o *BalanceResponse) SetRetentionAmount(v []BalanceCommonFielsResponse)`

SetRetentionAmount sets RetentionAmount field to given value.

### HasRetentionAmount

`func (o *BalanceResponse) HasRetentionAmount() bool`

HasRetentionAmount returns a boolean if a field has been set.

### GetTargetCollateralAmount

`func (o *BalanceResponse) GetTargetCollateralAmount() map[string]interface{}`

GetTargetCollateralAmount returns the TargetCollateralAmount field if non-nil, zero value otherwise.

### GetTargetCollateralAmountOk

`func (o *BalanceResponse) GetTargetCollateralAmountOk() (*map[string]interface{}, bool)`

GetTargetCollateralAmountOk returns a tuple with the TargetCollateralAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCollateralAmount

`func (o *BalanceResponse) SetTargetCollateralAmount(v map[string]interface{})`

SetTargetCollateralAmount sets TargetCollateralAmount field to given value.

### HasTargetCollateralAmount

`func (o *BalanceResponse) HasTargetCollateralAmount() bool`

HasTargetCollateralAmount returns a boolean if a field has been set.

### GetTargetRetentionAmount

`func (o *BalanceResponse) GetTargetRetentionAmount() []BalanceCommonFielsResponse`

GetTargetRetentionAmount returns the TargetRetentionAmount field if non-nil, zero value otherwise.

### GetTargetRetentionAmountOk

`func (o *BalanceResponse) GetTargetRetentionAmountOk() (*[]BalanceCommonFielsResponse, bool)`

GetTargetRetentionAmountOk returns a tuple with the TargetRetentionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRetentionAmount

`func (o *BalanceResponse) SetTargetRetentionAmount(v []BalanceCommonFielsResponse)`

SetTargetRetentionAmount sets TargetRetentionAmount field to given value.

### HasTargetRetentionAmount

`func (o *BalanceResponse) HasTargetRetentionAmount() bool`

HasTargetRetentionAmount returns a boolean if a field has been set.

### GetTemporarilyRetained

`func (o *BalanceResponse) GetTemporarilyRetained() []BalanceCommonFielsResponse`

GetTemporarilyRetained returns the TemporarilyRetained field if non-nil, zero value otherwise.

### GetTemporarilyRetainedOk

`func (o *BalanceResponse) GetTemporarilyRetainedOk() (*[]BalanceCommonFielsResponse, bool)`

GetTemporarilyRetainedOk returns a tuple with the TemporarilyRetained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporarilyRetained

`func (o *BalanceResponse) SetTemporarilyRetained(v []BalanceCommonFielsResponse)`

SetTemporarilyRetained sets TemporarilyRetained field to given value.

### HasTemporarilyRetained

`func (o *BalanceResponse) HasTemporarilyRetained() bool`

HasTemporarilyRetained returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


