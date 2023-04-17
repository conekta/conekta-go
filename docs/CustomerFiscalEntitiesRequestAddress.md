# CustomerFiscalEntitiesRequestAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Street1** | **string** |  | 
**Street2** | Pointer to **string** |  | [optional] 
**PostalCode** | **string** |  | 
**City** | **string** |  | 
**State** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** | this field follows the [ISO 3166-1 alpha-2 standard](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) | [optional] 
**Residential** | Pointer to **bool** |  | [optional] 
**ExternalNumber** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomerFiscalEntitiesRequestAddress

`func NewCustomerFiscalEntitiesRequestAddress(street1 string, postalCode string, city string, ) *CustomerFiscalEntitiesRequestAddress`

NewCustomerFiscalEntitiesRequestAddress instantiates a new CustomerFiscalEntitiesRequestAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerFiscalEntitiesRequestAddressWithDefaults

`func NewCustomerFiscalEntitiesRequestAddressWithDefaults() *CustomerFiscalEntitiesRequestAddress`

NewCustomerFiscalEntitiesRequestAddressWithDefaults instantiates a new CustomerFiscalEntitiesRequestAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreet1

`func (o *CustomerFiscalEntitiesRequestAddress) GetStreet1() string`

GetStreet1 returns the Street1 field if non-nil, zero value otherwise.

### GetStreet1Ok

`func (o *CustomerFiscalEntitiesRequestAddress) GetStreet1Ok() (*string, bool)`

GetStreet1Ok returns a tuple with the Street1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet1

`func (o *CustomerFiscalEntitiesRequestAddress) SetStreet1(v string)`

SetStreet1 sets Street1 field to given value.


### GetStreet2

`func (o *CustomerFiscalEntitiesRequestAddress) GetStreet2() string`

GetStreet2 returns the Street2 field if non-nil, zero value otherwise.

### GetStreet2Ok

`func (o *CustomerFiscalEntitiesRequestAddress) GetStreet2Ok() (*string, bool)`

GetStreet2Ok returns a tuple with the Street2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet2

`func (o *CustomerFiscalEntitiesRequestAddress) SetStreet2(v string)`

SetStreet2 sets Street2 field to given value.

### HasStreet2

`func (o *CustomerFiscalEntitiesRequestAddress) HasStreet2() bool`

HasStreet2 returns a boolean if a field has been set.

### GetPostalCode

`func (o *CustomerFiscalEntitiesRequestAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *CustomerFiscalEntitiesRequestAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *CustomerFiscalEntitiesRequestAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetCity

`func (o *CustomerFiscalEntitiesRequestAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CustomerFiscalEntitiesRequestAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CustomerFiscalEntitiesRequestAddress) SetCity(v string)`

SetCity sets City field to given value.


### GetState

`func (o *CustomerFiscalEntitiesRequestAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CustomerFiscalEntitiesRequestAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CustomerFiscalEntitiesRequestAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CustomerFiscalEntitiesRequestAddress) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *CustomerFiscalEntitiesRequestAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CustomerFiscalEntitiesRequestAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CustomerFiscalEntitiesRequestAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CustomerFiscalEntitiesRequestAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetResidential

`func (o *CustomerFiscalEntitiesRequestAddress) GetResidential() bool`

GetResidential returns the Residential field if non-nil, zero value otherwise.

### GetResidentialOk

`func (o *CustomerFiscalEntitiesRequestAddress) GetResidentialOk() (*bool, bool)`

GetResidentialOk returns a tuple with the Residential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidential

`func (o *CustomerFiscalEntitiesRequestAddress) SetResidential(v bool)`

SetResidential sets Residential field to given value.

### HasResidential

`func (o *CustomerFiscalEntitiesRequestAddress) HasResidential() bool`

HasResidential returns a boolean if a field has been set.

### GetExternalNumber

`func (o *CustomerFiscalEntitiesRequestAddress) GetExternalNumber() string`

GetExternalNumber returns the ExternalNumber field if non-nil, zero value otherwise.

### GetExternalNumberOk

`func (o *CustomerFiscalEntitiesRequestAddress) GetExternalNumberOk() (*string, bool)`

GetExternalNumberOk returns a tuple with the ExternalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalNumber

`func (o *CustomerFiscalEntitiesRequestAddress) SetExternalNumber(v string)`

SetExternalNumber sets ExternalNumber field to given value.

### HasExternalNumber

`func (o *CustomerFiscalEntitiesRequestAddress) HasExternalNumber() bool`

HasExternalNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


