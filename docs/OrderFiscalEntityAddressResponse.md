# OrderFiscalEntityAddressResponse

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
**Object** | Pointer to **string** |  | [optional] 

## Methods

### NewOrderFiscalEntityAddressResponse

`func NewOrderFiscalEntityAddressResponse(street1 string, postalCode string, city string, country string, externalNumber string, ) *OrderFiscalEntityAddressResponse`

NewOrderFiscalEntityAddressResponse instantiates a new OrderFiscalEntityAddressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderFiscalEntityAddressResponseWithDefaults

`func NewOrderFiscalEntityAddressResponseWithDefaults() *OrderFiscalEntityAddressResponse`

NewOrderFiscalEntityAddressResponseWithDefaults instantiates a new OrderFiscalEntityAddressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreet1

`func (o *OrderFiscalEntityAddressResponse) GetStreet1() string`

GetStreet1 returns the Street1 field if non-nil, zero value otherwise.

### GetStreet1Ok

`func (o *OrderFiscalEntityAddressResponse) GetStreet1Ok() (*string, bool)`

GetStreet1Ok returns a tuple with the Street1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet1

`func (o *OrderFiscalEntityAddressResponse) SetStreet1(v string)`

SetStreet1 sets Street1 field to given value.


### GetStreet2

`func (o *OrderFiscalEntityAddressResponse) GetStreet2() string`

GetStreet2 returns the Street2 field if non-nil, zero value otherwise.

### GetStreet2Ok

`func (o *OrderFiscalEntityAddressResponse) GetStreet2Ok() (*string, bool)`

GetStreet2Ok returns a tuple with the Street2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet2

`func (o *OrderFiscalEntityAddressResponse) SetStreet2(v string)`

SetStreet2 sets Street2 field to given value.

### HasStreet2

`func (o *OrderFiscalEntityAddressResponse) HasStreet2() bool`

HasStreet2 returns a boolean if a field has been set.

### SetStreet2Nil

`func (o *OrderFiscalEntityAddressResponse) SetStreet2Nil(b bool)`

 SetStreet2Nil sets the value for Street2 to be an explicit nil

### UnsetStreet2
`func (o *OrderFiscalEntityAddressResponse) UnsetStreet2()`

UnsetStreet2 ensures that no value is present for Street2, not even an explicit nil
### GetPostalCode

`func (o *OrderFiscalEntityAddressResponse) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *OrderFiscalEntityAddressResponse) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *OrderFiscalEntityAddressResponse) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetCity

`func (o *OrderFiscalEntityAddressResponse) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *OrderFiscalEntityAddressResponse) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *OrderFiscalEntityAddressResponse) SetCity(v string)`

SetCity sets City field to given value.


### GetState

`func (o *OrderFiscalEntityAddressResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OrderFiscalEntityAddressResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OrderFiscalEntityAddressResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *OrderFiscalEntityAddressResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *OrderFiscalEntityAddressResponse) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *OrderFiscalEntityAddressResponse) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *OrderFiscalEntityAddressResponse) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetExternalNumber

`func (o *OrderFiscalEntityAddressResponse) GetExternalNumber() string`

GetExternalNumber returns the ExternalNumber field if non-nil, zero value otherwise.

### GetExternalNumberOk

`func (o *OrderFiscalEntityAddressResponse) GetExternalNumberOk() (*string, bool)`

GetExternalNumberOk returns a tuple with the ExternalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalNumber

`func (o *OrderFiscalEntityAddressResponse) SetExternalNumber(v string)`

SetExternalNumber sets ExternalNumber field to given value.


### GetObject

`func (o *OrderFiscalEntityAddressResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *OrderFiscalEntityAddressResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *OrderFiscalEntityAddressResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *OrderFiscalEntityAddressResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


