# ChargeDataPaymentMethodCardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | Pointer to **string** |  | [optional] 
**AuthCode** | Pointer to **string** |  | [optional] 
**Brand** | Pointer to **string** |  | [optional] 
**ContractId** | Pointer to **string** | Id sent for recurrent charges. | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**ExpMonth** | Pointer to **string** |  | [optional] 
**ExpYear** | Pointer to **string** |  | [optional] 
**FraudIndicators** | Pointer to **[]interface{}** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**Last4** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewChargeDataPaymentMethodCardResponse

`func NewChargeDataPaymentMethodCardResponse() *ChargeDataPaymentMethodCardResponse`

NewChargeDataPaymentMethodCardResponse instantiates a new ChargeDataPaymentMethodCardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeDataPaymentMethodCardResponseWithDefaults

`func NewChargeDataPaymentMethodCardResponseWithDefaults() *ChargeDataPaymentMethodCardResponse`

NewChargeDataPaymentMethodCardResponseWithDefaults instantiates a new ChargeDataPaymentMethodCardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *ChargeDataPaymentMethodCardResponse) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *ChargeDataPaymentMethodCardResponse) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *ChargeDataPaymentMethodCardResponse) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *ChargeDataPaymentMethodCardResponse) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAuthCode

`func (o *ChargeDataPaymentMethodCardResponse) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *ChargeDataPaymentMethodCardResponse) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *ChargeDataPaymentMethodCardResponse) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *ChargeDataPaymentMethodCardResponse) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetBrand

`func (o *ChargeDataPaymentMethodCardResponse) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *ChargeDataPaymentMethodCardResponse) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *ChargeDataPaymentMethodCardResponse) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *ChargeDataPaymentMethodCardResponse) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetContractId

`func (o *ChargeDataPaymentMethodCardResponse) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *ChargeDataPaymentMethodCardResponse) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *ChargeDataPaymentMethodCardResponse) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *ChargeDataPaymentMethodCardResponse) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### GetCountry

`func (o *ChargeDataPaymentMethodCardResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ChargeDataPaymentMethodCardResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ChargeDataPaymentMethodCardResponse) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ChargeDataPaymentMethodCardResponse) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetExpMonth

`func (o *ChargeDataPaymentMethodCardResponse) GetExpMonth() string`

GetExpMonth returns the ExpMonth field if non-nil, zero value otherwise.

### GetExpMonthOk

`func (o *ChargeDataPaymentMethodCardResponse) GetExpMonthOk() (*string, bool)`

GetExpMonthOk returns a tuple with the ExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpMonth

`func (o *ChargeDataPaymentMethodCardResponse) SetExpMonth(v string)`

SetExpMonth sets ExpMonth field to given value.

### HasExpMonth

`func (o *ChargeDataPaymentMethodCardResponse) HasExpMonth() bool`

HasExpMonth returns a boolean if a field has been set.

### GetExpYear

`func (o *ChargeDataPaymentMethodCardResponse) GetExpYear() string`

GetExpYear returns the ExpYear field if non-nil, zero value otherwise.

### GetExpYearOk

`func (o *ChargeDataPaymentMethodCardResponse) GetExpYearOk() (*string, bool)`

GetExpYearOk returns a tuple with the ExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpYear

`func (o *ChargeDataPaymentMethodCardResponse) SetExpYear(v string)`

SetExpYear sets ExpYear field to given value.

### HasExpYear

`func (o *ChargeDataPaymentMethodCardResponse) HasExpYear() bool`

HasExpYear returns a boolean if a field has been set.

### GetFraudIndicators

`func (o *ChargeDataPaymentMethodCardResponse) GetFraudIndicators() []interface{}`

GetFraudIndicators returns the FraudIndicators field if non-nil, zero value otherwise.

### GetFraudIndicatorsOk

`func (o *ChargeDataPaymentMethodCardResponse) GetFraudIndicatorsOk() (*[]interface{}, bool)`

GetFraudIndicatorsOk returns a tuple with the FraudIndicators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudIndicators

`func (o *ChargeDataPaymentMethodCardResponse) SetFraudIndicators(v []interface{})`

SetFraudIndicators sets FraudIndicators field to given value.

### HasFraudIndicators

`func (o *ChargeDataPaymentMethodCardResponse) HasFraudIndicators() bool`

HasFraudIndicators returns a boolean if a field has been set.

### GetIssuer

`func (o *ChargeDataPaymentMethodCardResponse) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ChargeDataPaymentMethodCardResponse) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ChargeDataPaymentMethodCardResponse) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *ChargeDataPaymentMethodCardResponse) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetLast4

`func (o *ChargeDataPaymentMethodCardResponse) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *ChargeDataPaymentMethodCardResponse) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *ChargeDataPaymentMethodCardResponse) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *ChargeDataPaymentMethodCardResponse) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### GetName

`func (o *ChargeDataPaymentMethodCardResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChargeDataPaymentMethodCardResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChargeDataPaymentMethodCardResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChargeDataPaymentMethodCardResponse) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


