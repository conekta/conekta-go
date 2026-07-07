# FiscalEntityRequestAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Street1** | **string** |  | 
**Street2** | Pointer to **string** |  | [optional] 
**PostalCode** | **string** |  | 
**City** | **string** |  | 
**State** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** | this field follows the [ISO 3166-1 alpha-2 standard](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) | [optional] 
**Residential** | Pointer to **bool** |  | [optional] [default to false]
**ExternalNumber** | Pointer to **string** |  | [optional] 

## Methods

### NewFiscalEntityRequestAddress

`func NewFiscalEntityRequestAddress(street1 string, postalCode string, city string, ) *FiscalEntityRequestAddress`

NewFiscalEntityRequestAddress instantiates a new FiscalEntityRequestAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiscalEntityRequestAddressWithDefaults

`func NewFiscalEntityRequestAddressWithDefaults() *FiscalEntityRequestAddress`

NewFiscalEntityRequestAddressWithDefaults instantiates a new FiscalEntityRequestAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreet1

`func (o *FiscalEntityRequestAddress) GetStreet1() string`

GetStreet1 returns the Street1 field if non-nil, zero value otherwise.

### GetStreet1Ok

`func (o *FiscalEntityRequestAddress) GetStreet1Ok() (*string, bool)`

GetStreet1Ok returns a tuple with the Street1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet1

`func (o *FiscalEntityRequestAddress) SetStreet1(v string)`

SetStreet1 sets Street1 field to given value.


### GetStreet2

`func (o *FiscalEntityRequestAddress) GetStreet2() string`

GetStreet2 returns the Street2 field if non-nil, zero value otherwise.

### GetStreet2Ok

`func (o *FiscalEntityRequestAddress) GetStreet2Ok() (*string, bool)`

GetStreet2Ok returns a tuple with the Street2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet2

`func (o *FiscalEntityRequestAddress) SetStreet2(v string)`

SetStreet2 sets Street2 field to given value.

### HasStreet2

`func (o *FiscalEntityRequestAddress) HasStreet2() bool`

HasStreet2 returns a boolean if a field has been set.

### GetPostalCode

`func (o *FiscalEntityRequestAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *FiscalEntityRequestAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *FiscalEntityRequestAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetCity

`func (o *FiscalEntityRequestAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *FiscalEntityRequestAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *FiscalEntityRequestAddress) SetCity(v string)`

SetCity sets City field to given value.


### GetState

`func (o *FiscalEntityRequestAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FiscalEntityRequestAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FiscalEntityRequestAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FiscalEntityRequestAddress) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *FiscalEntityRequestAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *FiscalEntityRequestAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *FiscalEntityRequestAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *FiscalEntityRequestAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetResidential

`func (o *FiscalEntityRequestAddress) GetResidential() bool`

GetResidential returns the Residential field if non-nil, zero value otherwise.

### GetResidentialOk

`func (o *FiscalEntityRequestAddress) GetResidentialOk() (*bool, bool)`

GetResidentialOk returns a tuple with the Residential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidential

`func (o *FiscalEntityRequestAddress) SetResidential(v bool)`

SetResidential sets Residential field to given value.

### HasResidential

`func (o *FiscalEntityRequestAddress) HasResidential() bool`

HasResidential returns a boolean if a field has been set.

### GetExternalNumber

`func (o *FiscalEntityRequestAddress) GetExternalNumber() string`

GetExternalNumber returns the ExternalNumber field if non-nil, zero value otherwise.

### GetExternalNumberOk

`func (o *FiscalEntityRequestAddress) GetExternalNumberOk() (*string, bool)`

GetExternalNumberOk returns a tuple with the ExternalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalNumber

`func (o *FiscalEntityRequestAddress) SetExternalNumber(v string)`

SetExternalNumber sets ExternalNumber field to given value.

### HasExternalNumber

`func (o *FiscalEntityRequestAddress) HasExternalNumber() bool`

HasExternalNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


