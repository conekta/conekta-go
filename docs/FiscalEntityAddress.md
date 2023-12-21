# FiscalEntityAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Street1** | **string** | Street name and number | 
**Street2** | Pointer to **NullableString** | Street name and number | [optional] 
**PostalCode** | **string** | Postal code | 
**City** | **string** | City | 
**State** | Pointer to **string** | State | [optional] 
**Country** | **string** | this field follows the [ISO 3166-1 alpha-2 standard](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) | 
**ExternalNumber** | **string** | External number | 

## Methods

### NewFiscalEntityAddress

`func NewFiscalEntityAddress(street1 string, postalCode string, city string, country string, externalNumber string, ) *FiscalEntityAddress`

NewFiscalEntityAddress instantiates a new FiscalEntityAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiscalEntityAddressWithDefaults

`func NewFiscalEntityAddressWithDefaults() *FiscalEntityAddress`

NewFiscalEntityAddressWithDefaults instantiates a new FiscalEntityAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreet1

`func (o *FiscalEntityAddress) GetStreet1() string`

GetStreet1 returns the Street1 field if non-nil, zero value otherwise.

### GetStreet1Ok

`func (o *FiscalEntityAddress) GetStreet1Ok() (*string, bool)`

GetStreet1Ok returns a tuple with the Street1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet1

`func (o *FiscalEntityAddress) SetStreet1(v string)`

SetStreet1 sets Street1 field to given value.


### GetStreet2

`func (o *FiscalEntityAddress) GetStreet2() string`

GetStreet2 returns the Street2 field if non-nil, zero value otherwise.

### GetStreet2Ok

`func (o *FiscalEntityAddress) GetStreet2Ok() (*string, bool)`

GetStreet2Ok returns a tuple with the Street2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet2

`func (o *FiscalEntityAddress) SetStreet2(v string)`

SetStreet2 sets Street2 field to given value.

### HasStreet2

`func (o *FiscalEntityAddress) HasStreet2() bool`

HasStreet2 returns a boolean if a field has been set.

### SetStreet2Nil

`func (o *FiscalEntityAddress) SetStreet2Nil(b bool)`

 SetStreet2Nil sets the value for Street2 to be an explicit nil

### UnsetStreet2
`func (o *FiscalEntityAddress) UnsetStreet2()`

UnsetStreet2 ensures that no value is present for Street2, not even an explicit nil
### GetPostalCode

`func (o *FiscalEntityAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *FiscalEntityAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *FiscalEntityAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetCity

`func (o *FiscalEntityAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *FiscalEntityAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *FiscalEntityAddress) SetCity(v string)`

SetCity sets City field to given value.


### GetState

`func (o *FiscalEntityAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FiscalEntityAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FiscalEntityAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FiscalEntityAddress) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *FiscalEntityAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *FiscalEntityAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *FiscalEntityAddress) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetExternalNumber

`func (o *FiscalEntityAddress) GetExternalNumber() string`

GetExternalNumber returns the ExternalNumber field if non-nil, zero value otherwise.

### GetExternalNumberOk

`func (o *FiscalEntityAddress) GetExternalNumberOk() (*string, bool)`

GetExternalNumberOk returns a tuple with the ExternalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalNumber

`func (o *FiscalEntityAddress) SetExternalNumber(v string)`

SetExternalNumber sets ExternalNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


