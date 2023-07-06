# OrderUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Charges** | Pointer to [**[]ChargeRequest**](ChargeRequest.md) |  | [optional] 
**Checkout** | Pointer to [**CheckoutRequest**](CheckoutRequest.md) |  | [optional] 
**Currency** | Pointer to **string** | Currency with which the payment will be made. It uses the 3-letter code of the [International Standard ISO 4217.](https://es.wikipedia.org/wiki/ISO_4217) | [optional] 
**CustomerInfo** | Pointer to [**OrderUpdateRequestCustomerInfo**](OrderUpdateRequestCustomerInfo.md) |  | [optional] 
**DiscountLines** | Pointer to [**[]OrderDiscountLinesRequest**](OrderDiscountLinesRequest.md) | List of [discounts](https://developers.conekta.com/v2.1.0/reference/orderscreatediscountline) that are applied to the order. You must have at least one discount. | [optional] 
**LineItems** | Pointer to [**[]Product**](Product.md) | List of [products](https://developers.conekta.com/v2.1.0/reference/orderscreateproduct) that are sold in the order. You must have at least one product. | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**PreAuthorize** | Pointer to **bool** | Indicates whether the order charges must be preauthorized | [optional] [default to false]
**ShippingContact** | Pointer to [**CustomerShippingContacts**](CustomerShippingContacts.md) |  | [optional] 
**ShippingLines** | Pointer to [**[]ShippingRequest**](ShippingRequest.md) | List of [shipping costs](https://developers.conekta.com/v2.1.0/reference/orderscreateshipping). If the online store offers digital products. | [optional] 
**TaxLines** | Pointer to [**[]OrderTaxRequest**](OrderTaxRequest.md) |  | [optional] 

## Methods

### NewOrderUpdateRequest

`func NewOrderUpdateRequest() *OrderUpdateRequest`

NewOrderUpdateRequest instantiates a new OrderUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderUpdateRequestWithDefaults

`func NewOrderUpdateRequestWithDefaults() *OrderUpdateRequest`

NewOrderUpdateRequestWithDefaults instantiates a new OrderUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCharges

`func (o *OrderUpdateRequest) GetCharges() []ChargeRequest`

GetCharges returns the Charges field if non-nil, zero value otherwise.

### GetChargesOk

`func (o *OrderUpdateRequest) GetChargesOk() (*[]ChargeRequest, bool)`

GetChargesOk returns a tuple with the Charges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharges

`func (o *OrderUpdateRequest) SetCharges(v []ChargeRequest)`

SetCharges sets Charges field to given value.

### HasCharges

`func (o *OrderUpdateRequest) HasCharges() bool`

HasCharges returns a boolean if a field has been set.

### GetCheckout

`func (o *OrderUpdateRequest) GetCheckout() CheckoutRequest`

GetCheckout returns the Checkout field if non-nil, zero value otherwise.

### GetCheckoutOk

`func (o *OrderUpdateRequest) GetCheckoutOk() (*CheckoutRequest, bool)`

GetCheckoutOk returns a tuple with the Checkout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckout

`func (o *OrderUpdateRequest) SetCheckout(v CheckoutRequest)`

SetCheckout sets Checkout field to given value.

### HasCheckout

`func (o *OrderUpdateRequest) HasCheckout() bool`

HasCheckout returns a boolean if a field has been set.

### GetCurrency

`func (o *OrderUpdateRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OrderUpdateRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OrderUpdateRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *OrderUpdateRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerInfo

`func (o *OrderUpdateRequest) GetCustomerInfo() OrderUpdateRequestCustomerInfo`

GetCustomerInfo returns the CustomerInfo field if non-nil, zero value otherwise.

### GetCustomerInfoOk

`func (o *OrderUpdateRequest) GetCustomerInfoOk() (*OrderUpdateRequestCustomerInfo, bool)`

GetCustomerInfoOk returns a tuple with the CustomerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerInfo

`func (o *OrderUpdateRequest) SetCustomerInfo(v OrderUpdateRequestCustomerInfo)`

SetCustomerInfo sets CustomerInfo field to given value.

### HasCustomerInfo

`func (o *OrderUpdateRequest) HasCustomerInfo() bool`

HasCustomerInfo returns a boolean if a field has been set.

### GetDiscountLines

`func (o *OrderUpdateRequest) GetDiscountLines() []OrderDiscountLinesRequest`

GetDiscountLines returns the DiscountLines field if non-nil, zero value otherwise.

### GetDiscountLinesOk

`func (o *OrderUpdateRequest) GetDiscountLinesOk() (*[]OrderDiscountLinesRequest, bool)`

GetDiscountLinesOk returns a tuple with the DiscountLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountLines

`func (o *OrderUpdateRequest) SetDiscountLines(v []OrderDiscountLinesRequest)`

SetDiscountLines sets DiscountLines field to given value.

### HasDiscountLines

`func (o *OrderUpdateRequest) HasDiscountLines() bool`

HasDiscountLines returns a boolean if a field has been set.

### GetLineItems

`func (o *OrderUpdateRequest) GetLineItems() []Product`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *OrderUpdateRequest) GetLineItemsOk() (*[]Product, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *OrderUpdateRequest) SetLineItems(v []Product)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *OrderUpdateRequest) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetMetadata

`func (o *OrderUpdateRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrderUpdateRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrderUpdateRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OrderUpdateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPreAuthorize

`func (o *OrderUpdateRequest) GetPreAuthorize() bool`

GetPreAuthorize returns the PreAuthorize field if non-nil, zero value otherwise.

### GetPreAuthorizeOk

`func (o *OrderUpdateRequest) GetPreAuthorizeOk() (*bool, bool)`

GetPreAuthorizeOk returns a tuple with the PreAuthorize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthorize

`func (o *OrderUpdateRequest) SetPreAuthorize(v bool)`

SetPreAuthorize sets PreAuthorize field to given value.

### HasPreAuthorize

`func (o *OrderUpdateRequest) HasPreAuthorize() bool`

HasPreAuthorize returns a boolean if a field has been set.

### GetShippingContact

`func (o *OrderUpdateRequest) GetShippingContact() CustomerShippingContacts`

GetShippingContact returns the ShippingContact field if non-nil, zero value otherwise.

### GetShippingContactOk

`func (o *OrderUpdateRequest) GetShippingContactOk() (*CustomerShippingContacts, bool)`

GetShippingContactOk returns a tuple with the ShippingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingContact

`func (o *OrderUpdateRequest) SetShippingContact(v CustomerShippingContacts)`

SetShippingContact sets ShippingContact field to given value.

### HasShippingContact

`func (o *OrderUpdateRequest) HasShippingContact() bool`

HasShippingContact returns a boolean if a field has been set.

### GetShippingLines

`func (o *OrderUpdateRequest) GetShippingLines() []ShippingRequest`

GetShippingLines returns the ShippingLines field if non-nil, zero value otherwise.

### GetShippingLinesOk

`func (o *OrderUpdateRequest) GetShippingLinesOk() (*[]ShippingRequest, bool)`

GetShippingLinesOk returns a tuple with the ShippingLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLines

`func (o *OrderUpdateRequest) SetShippingLines(v []ShippingRequest)`

SetShippingLines sets ShippingLines field to given value.

### HasShippingLines

`func (o *OrderUpdateRequest) HasShippingLines() bool`

HasShippingLines returns a boolean if a field has been set.

### GetTaxLines

`func (o *OrderUpdateRequest) GetTaxLines() []OrderTaxRequest`

GetTaxLines returns the TaxLines field if non-nil, zero value otherwise.

### GetTaxLinesOk

`func (o *OrderUpdateRequest) GetTaxLinesOk() (*[]OrderTaxRequest, bool)`

GetTaxLinesOk returns a tuple with the TaxLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxLines

`func (o *OrderUpdateRequest) SetTaxLines(v []OrderTaxRequest)`

SetTaxLines sets TaxLines field to given value.

### HasTaxLines

`func (o *OrderUpdateRequest) HasTaxLines() bool`

HasTaxLines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


