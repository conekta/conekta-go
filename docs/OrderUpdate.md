# OrderUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Charges** | Pointer to [**[]ChargeRequest**](ChargeRequest.md) |  | [optional] 
**Checkout** | Pointer to [**OrderCheckoutRequest**](OrderCheckoutRequest.md) |  | [optional] 
**Currency** | Pointer to **string** | Currency with which the payment will be made. It uses the 3-letter code of the [International Standard ISO 4217.](https://es.wikipedia.org/wiki/ISO_4217) | [optional] 
**CustomerInfo** | Pointer to [**OrderUpdateCustomerInfo**](OrderUpdateCustomerInfo.md) |  | [optional] 
**DiscountLines** | Pointer to [**[]OrderDiscountLinesRequest**](OrderDiscountLinesRequest.md) | List of [discounts](https://developers.conekta.com/v2.3.0/reference/orderscreatediscountline) that are applied to the order. | [optional] 
**FiscalEntity** | Pointer to [**OrderUpdateFiscalEntityRequest**](OrderUpdateFiscalEntityRequest.md) |  | [optional] 
**LineItems** | Pointer to [**[]Product**](Product.md) | List of [products](https://developers.conekta.com/v2.3.0/reference/orderscreateproduct) that are sold in the order. You must have at least one product. | [optional] 
**Metadata** | Pointer to  |  | [optional] 
**PreAuthorize** | Pointer to **bool** | Indicates whether the order charges must be preauthorized | [optional] 
**ShippingContact** | Pointer to [**CustomerShippingContactsRequest**](CustomerShippingContactsRequest.md) |  | [optional] 
**ShippingLines** | Pointer to [**[]ShippingRequest**](ShippingRequest.md) | List of [shipping costs](https://developers.conekta.com/v2.3.0/reference/orderscreateshipping). If the online store offers digital products. | [optional] 
**TaxLines** | Pointer to [**[]OrderTaxRequest**](OrderTaxRequest.md) |  | [optional] 

## Methods

### NewOrderUpdate

`func NewOrderUpdate() *OrderUpdate`

NewOrderUpdate instantiates a new OrderUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderUpdateWithDefaults

`func NewOrderUpdateWithDefaults() *OrderUpdate`

NewOrderUpdateWithDefaults instantiates a new OrderUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCharges

`func (o *OrderUpdate) GetCharges() []ChargeRequest`

GetCharges returns the Charges field if non-nil, zero value otherwise.

### GetChargesOk

`func (o *OrderUpdate) GetChargesOk() (*[]ChargeRequest, bool)`

GetChargesOk returns a tuple with the Charges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharges

`func (o *OrderUpdate) SetCharges(v []ChargeRequest)`

SetCharges sets Charges field to given value.

### HasCharges

`func (o *OrderUpdate) HasCharges() bool`

HasCharges returns a boolean if a field has been set.

### GetCheckout

`func (o *OrderUpdate) GetCheckout() OrderCheckoutRequest`

GetCheckout returns the Checkout field if non-nil, zero value otherwise.

### GetCheckoutOk

`func (o *OrderUpdate) GetCheckoutOk() (*OrderCheckoutRequest, bool)`

GetCheckoutOk returns a tuple with the Checkout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckout

`func (o *OrderUpdate) SetCheckout(v OrderCheckoutRequest)`

SetCheckout sets Checkout field to given value.

### HasCheckout

`func (o *OrderUpdate) HasCheckout() bool`

HasCheckout returns a boolean if a field has been set.

### GetCurrency

`func (o *OrderUpdate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OrderUpdate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OrderUpdate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *OrderUpdate) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerInfo

`func (o *OrderUpdate) GetCustomerInfo() OrderUpdateCustomerInfo`

GetCustomerInfo returns the CustomerInfo field if non-nil, zero value otherwise.

### GetCustomerInfoOk

`func (o *OrderUpdate) GetCustomerInfoOk() (*OrderUpdateCustomerInfo, bool)`

GetCustomerInfoOk returns a tuple with the CustomerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerInfo

`func (o *OrderUpdate) SetCustomerInfo(v OrderUpdateCustomerInfo)`

SetCustomerInfo sets CustomerInfo field to given value.

### HasCustomerInfo

`func (o *OrderUpdate) HasCustomerInfo() bool`

HasCustomerInfo returns a boolean if a field has been set.

### GetDiscountLines

`func (o *OrderUpdate) GetDiscountLines() []OrderDiscountLinesRequest`

GetDiscountLines returns the DiscountLines field if non-nil, zero value otherwise.

### GetDiscountLinesOk

`func (o *OrderUpdate) GetDiscountLinesOk() (*[]OrderDiscountLinesRequest, bool)`

GetDiscountLinesOk returns a tuple with the DiscountLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountLines

`func (o *OrderUpdate) SetDiscountLines(v []OrderDiscountLinesRequest)`

SetDiscountLines sets DiscountLines field to given value.

### HasDiscountLines

`func (o *OrderUpdate) HasDiscountLines() bool`

HasDiscountLines returns a boolean if a field has been set.

### GetFiscalEntity

`func (o *OrderUpdate) GetFiscalEntity() OrderUpdateFiscalEntityRequest`

GetFiscalEntity returns the FiscalEntity field if non-nil, zero value otherwise.

### GetFiscalEntityOk

`func (o *OrderUpdate) GetFiscalEntityOk() (*OrderUpdateFiscalEntityRequest, bool)`

GetFiscalEntityOk returns a tuple with the FiscalEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalEntity

`func (o *OrderUpdate) SetFiscalEntity(v OrderUpdateFiscalEntityRequest)`

SetFiscalEntity sets FiscalEntity field to given value.

### HasFiscalEntity

`func (o *OrderUpdate) HasFiscalEntity() bool`

HasFiscalEntity returns a boolean if a field has been set.

### GetLineItems

`func (o *OrderUpdate) GetLineItems() []Product`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *OrderUpdate) GetLineItemsOk() (*[]Product, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *OrderUpdate) SetLineItems(v []Product)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *OrderUpdate) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetMetadata

`func (o *OrderUpdate) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrderUpdate) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrderUpdate) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OrderUpdate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *OrderUpdate) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *OrderUpdate) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetPreAuthorize

`func (o *OrderUpdate) GetPreAuthorize() bool`

GetPreAuthorize returns the PreAuthorize field if non-nil, zero value otherwise.

### GetPreAuthorizeOk

`func (o *OrderUpdate) GetPreAuthorizeOk() (*bool, bool)`

GetPreAuthorizeOk returns a tuple with the PreAuthorize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthorize

`func (o *OrderUpdate) SetPreAuthorize(v bool)`

SetPreAuthorize sets PreAuthorize field to given value.

### HasPreAuthorize

`func (o *OrderUpdate) HasPreAuthorize() bool`

HasPreAuthorize returns a boolean if a field has been set.

### GetShippingContact

`func (o *OrderUpdate) GetShippingContact() CustomerShippingContactsRequest`

GetShippingContact returns the ShippingContact field if non-nil, zero value otherwise.

### GetShippingContactOk

`func (o *OrderUpdate) GetShippingContactOk() (*CustomerShippingContactsRequest, bool)`

GetShippingContactOk returns a tuple with the ShippingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingContact

`func (o *OrderUpdate) SetShippingContact(v CustomerShippingContactsRequest)`

SetShippingContact sets ShippingContact field to given value.

### HasShippingContact

`func (o *OrderUpdate) HasShippingContact() bool`

HasShippingContact returns a boolean if a field has been set.

### GetShippingLines

`func (o *OrderUpdate) GetShippingLines() []ShippingRequest`

GetShippingLines returns the ShippingLines field if non-nil, zero value otherwise.

### GetShippingLinesOk

`func (o *OrderUpdate) GetShippingLinesOk() (*[]ShippingRequest, bool)`

GetShippingLinesOk returns a tuple with the ShippingLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLines

`func (o *OrderUpdate) SetShippingLines(v []ShippingRequest)`

SetShippingLines sets ShippingLines field to given value.

### HasShippingLines

`func (o *OrderUpdate) HasShippingLines() bool`

HasShippingLines returns a boolean if a field has been set.

### GetTaxLines

`func (o *OrderUpdate) GetTaxLines() []OrderTaxRequest`

GetTaxLines returns the TaxLines field if non-nil, zero value otherwise.

### GetTaxLinesOk

`func (o *OrderUpdate) GetTaxLinesOk() (*[]OrderTaxRequest, bool)`

GetTaxLinesOk returns a tuple with the TaxLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxLines

`func (o *OrderUpdate) SetTaxLines(v []OrderTaxRequest)`

SetTaxLines sets TaxLines field to given value.

### HasTaxLines

`func (o *OrderUpdate) HasTaxLines() bool`

HasTaxLines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


