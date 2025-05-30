# PaymentMethodCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Object** | **string** |  | 
**AccountType** | Pointer to **string** | Account type of the card | [optional] 
**AuthCode** | Pointer to **string** |  | [optional] 
**Brand** | Pointer to **string** | Brand of the card | [optional] 
**ContractId** | Pointer to **string** | Id sent for recurrent charges. | [optional] 
**Country** | Pointer to **string** | Country of the card | [optional] 
**ExpMonth** | Pointer to **string** | Expiration month of the card | [optional] 
**ExpYear** | Pointer to **string** | Expiration year of the card | [optional] 
**FraudIndicators** | Pointer to **[]interface{}** |  | [optional] 
**Issuer** | Pointer to **string** | Issuer of the card | [optional] 
**Last4** | Pointer to **string** | Last 4 digits of the card | [optional] 
**Name** | Pointer to **string** | Name of the cardholder | [optional] 
**CustomerIpAddress** | Pointer to **string** | Optional field used to capture the customer&#39;s IP address for fraud prevention and security monitoring purposes | [optional] 

## Methods

### NewPaymentMethodCard

`func NewPaymentMethodCard(object string, ) *PaymentMethodCard`

NewPaymentMethodCard instantiates a new PaymentMethodCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodCardWithDefaults

`func NewPaymentMethodCardWithDefaults() *PaymentMethodCard`

NewPaymentMethodCardWithDefaults instantiates a new PaymentMethodCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethodCard) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodCard) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodCard) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentMethodCard) HasType() bool`

HasType returns a boolean if a field has been set.

### GetObject

`func (o *PaymentMethodCard) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PaymentMethodCard) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PaymentMethodCard) SetObject(v string)`

SetObject sets Object field to given value.


### GetAccountType

`func (o *PaymentMethodCard) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *PaymentMethodCard) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *PaymentMethodCard) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *PaymentMethodCard) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAuthCode

`func (o *PaymentMethodCard) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *PaymentMethodCard) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *PaymentMethodCard) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *PaymentMethodCard) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetBrand

`func (o *PaymentMethodCard) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *PaymentMethodCard) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *PaymentMethodCard) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *PaymentMethodCard) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetContractId

`func (o *PaymentMethodCard) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *PaymentMethodCard) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *PaymentMethodCard) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *PaymentMethodCard) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### GetCountry

`func (o *PaymentMethodCard) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PaymentMethodCard) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PaymentMethodCard) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PaymentMethodCard) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetExpMonth

`func (o *PaymentMethodCard) GetExpMonth() string`

GetExpMonth returns the ExpMonth field if non-nil, zero value otherwise.

### GetExpMonthOk

`func (o *PaymentMethodCard) GetExpMonthOk() (*string, bool)`

GetExpMonthOk returns a tuple with the ExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpMonth

`func (o *PaymentMethodCard) SetExpMonth(v string)`

SetExpMonth sets ExpMonth field to given value.

### HasExpMonth

`func (o *PaymentMethodCard) HasExpMonth() bool`

HasExpMonth returns a boolean if a field has been set.

### GetExpYear

`func (o *PaymentMethodCard) GetExpYear() string`

GetExpYear returns the ExpYear field if non-nil, zero value otherwise.

### GetExpYearOk

`func (o *PaymentMethodCard) GetExpYearOk() (*string, bool)`

GetExpYearOk returns a tuple with the ExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpYear

`func (o *PaymentMethodCard) SetExpYear(v string)`

SetExpYear sets ExpYear field to given value.

### HasExpYear

`func (o *PaymentMethodCard) HasExpYear() bool`

HasExpYear returns a boolean if a field has been set.

### GetFraudIndicators

`func (o *PaymentMethodCard) GetFraudIndicators() []interface{}`

GetFraudIndicators returns the FraudIndicators field if non-nil, zero value otherwise.

### GetFraudIndicatorsOk

`func (o *PaymentMethodCard) GetFraudIndicatorsOk() (*[]interface{}, bool)`

GetFraudIndicatorsOk returns a tuple with the FraudIndicators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudIndicators

`func (o *PaymentMethodCard) SetFraudIndicators(v []interface{})`

SetFraudIndicators sets FraudIndicators field to given value.

### HasFraudIndicators

`func (o *PaymentMethodCard) HasFraudIndicators() bool`

HasFraudIndicators returns a boolean if a field has been set.

### GetIssuer

`func (o *PaymentMethodCard) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *PaymentMethodCard) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *PaymentMethodCard) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *PaymentMethodCard) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetLast4

`func (o *PaymentMethodCard) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *PaymentMethodCard) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *PaymentMethodCard) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *PaymentMethodCard) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### GetName

`func (o *PaymentMethodCard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentMethodCard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentMethodCard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaymentMethodCard) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCustomerIpAddress

`func (o *PaymentMethodCard) GetCustomerIpAddress() string`

GetCustomerIpAddress returns the CustomerIpAddress field if non-nil, zero value otherwise.

### GetCustomerIpAddressOk

`func (o *PaymentMethodCard) GetCustomerIpAddressOk() (*string, bool)`

GetCustomerIpAddressOk returns a tuple with the CustomerIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIpAddress

`func (o *PaymentMethodCard) SetCustomerIpAddress(v string)`

SetCustomerIpAddress sets CustomerIpAddress field to given value.

### HasCustomerIpAddress

`func (o *PaymentMethodCard) HasCustomerIpAddress() bool`

HasCustomerIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


