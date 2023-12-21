# OrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Charges** | Pointer to [**[]ChargeRequest**](ChargeRequest.md) | List of [charges](https://developers.conekta.com/v2.1.0/reference/orderscreatecharge) that are applied to the order | [optional] 
**Checkout** | Pointer to [**CheckoutRequest**](CheckoutRequest.md) |  | [optional] 
**Currency** | **string** | Currency with which the payment will be made. It uses the 3-letter code of the [International Standard ISO 4217.](https://es.wikipedia.org/wiki/ISO_4217) | 
**CustomerInfo** | [**OrderRequestCustomerInfo**](OrderRequestCustomerInfo.md) |  | 
**DiscountLines** | Pointer to [**[]OrderDiscountLinesRequest**](OrderDiscountLinesRequest.md) | List of [discounts](https://developers.conekta.com/v2.1.0/reference/orderscreatediscountline) that are applied to the order. You must have at least one discount. | [optional] 
**FiscalEntity** | Pointer to [**OrderFiscalEntityRequest**](OrderFiscalEntityRequest.md) |  | [optional] 
**LineItems** | [**[]Product**](Product.md) | List of [products](https://developers.conekta.com/v2.1.0/reference/orderscreateproduct) that are sold in the order. You must have at least one product. | 
**Metadata** | Pointer to **map[string]interface{}** | Metadata associated with the order | [optional] 
**NeedsShippingContact** | Pointer to **bool** | Allows you to fill out the shipping information at checkout | [optional] 
**PreAuthorize** | Pointer to **bool** | Indicates whether the order charges must be preauthorized | [optional] [default to false]
**ProcessingMode** | Pointer to **string** | Indicates the processing mode for the order, either ecommerce, recurrent or validation. | [optional] 
**ReturnUrl** | Pointer to **string** | Indicates the redirection callback upon completion of the 3DS2 flow. | [optional] 
**ShippingContact** | Pointer to [**CustomerShippingContacts**](CustomerShippingContacts.md) |  | [optional] 
**ShippingLines** | Pointer to [**[]ShippingRequest**](ShippingRequest.md) | List of [shipping costs](https://developers.conekta.com/v2.1.0/reference/orderscreateshipping). If the online store offers digital products. | [optional] 
**TaxLines** | Pointer to [**[]OrderTaxRequest**](OrderTaxRequest.md) | List of [taxes](https://developers.conekta.com/v2.1.0/reference/orderscreatetaxes) that are applied to the order. | [optional] 
**ThreeDsMode** | Pointer to **string** | Indicates the 3DS2 mode for the order, either smart or strict. | [optional] 

## Methods

### NewOrderRequest

`func NewOrderRequest(currency string, customerInfo OrderRequestCustomerInfo, lineItems []Product, ) *OrderRequest`

NewOrderRequest instantiates a new OrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderRequestWithDefaults

`func NewOrderRequestWithDefaults() *OrderRequest`

NewOrderRequestWithDefaults instantiates a new OrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCharges

`func (o *OrderRequest) GetCharges() []ChargeRequest`

GetCharges returns the Charges field if non-nil, zero value otherwise.

### GetChargesOk

`func (o *OrderRequest) GetChargesOk() (*[]ChargeRequest, bool)`

GetChargesOk returns a tuple with the Charges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharges

`func (o *OrderRequest) SetCharges(v []ChargeRequest)`

SetCharges sets Charges field to given value.

### HasCharges

`func (o *OrderRequest) HasCharges() bool`

HasCharges returns a boolean if a field has been set.

### GetCheckout

`func (o *OrderRequest) GetCheckout() CheckoutRequest`

GetCheckout returns the Checkout field if non-nil, zero value otherwise.

### GetCheckoutOk

`func (o *OrderRequest) GetCheckoutOk() (*CheckoutRequest, bool)`

GetCheckoutOk returns a tuple with the Checkout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckout

`func (o *OrderRequest) SetCheckout(v CheckoutRequest)`

SetCheckout sets Checkout field to given value.

### HasCheckout

`func (o *OrderRequest) HasCheckout() bool`

HasCheckout returns a boolean if a field has been set.

### GetCurrency

`func (o *OrderRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OrderRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OrderRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerInfo

`func (o *OrderRequest) GetCustomerInfo() OrderRequestCustomerInfo`

GetCustomerInfo returns the CustomerInfo field if non-nil, zero value otherwise.

### GetCustomerInfoOk

`func (o *OrderRequest) GetCustomerInfoOk() (*OrderRequestCustomerInfo, bool)`

GetCustomerInfoOk returns a tuple with the CustomerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerInfo

`func (o *OrderRequest) SetCustomerInfo(v OrderRequestCustomerInfo)`

SetCustomerInfo sets CustomerInfo field to given value.


### GetDiscountLines

`func (o *OrderRequest) GetDiscountLines() []OrderDiscountLinesRequest`

GetDiscountLines returns the DiscountLines field if non-nil, zero value otherwise.

### GetDiscountLinesOk

`func (o *OrderRequest) GetDiscountLinesOk() (*[]OrderDiscountLinesRequest, bool)`

GetDiscountLinesOk returns a tuple with the DiscountLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountLines

`func (o *OrderRequest) SetDiscountLines(v []OrderDiscountLinesRequest)`

SetDiscountLines sets DiscountLines field to given value.

### HasDiscountLines

`func (o *OrderRequest) HasDiscountLines() bool`

HasDiscountLines returns a boolean if a field has been set.

### GetFiscalEntity

`func (o *OrderRequest) GetFiscalEntity() OrderFiscalEntityRequest`

GetFiscalEntity returns the FiscalEntity field if non-nil, zero value otherwise.

### GetFiscalEntityOk

`func (o *OrderRequest) GetFiscalEntityOk() (*OrderFiscalEntityRequest, bool)`

GetFiscalEntityOk returns a tuple with the FiscalEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalEntity

`func (o *OrderRequest) SetFiscalEntity(v OrderFiscalEntityRequest)`

SetFiscalEntity sets FiscalEntity field to given value.

### HasFiscalEntity

`func (o *OrderRequest) HasFiscalEntity() bool`

HasFiscalEntity returns a boolean if a field has been set.

### GetLineItems

`func (o *OrderRequest) GetLineItems() []Product`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *OrderRequest) GetLineItemsOk() (*[]Product, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *OrderRequest) SetLineItems(v []Product)`

SetLineItems sets LineItems field to given value.


### GetMetadata

`func (o *OrderRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrderRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrderRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OrderRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNeedsShippingContact

`func (o *OrderRequest) GetNeedsShippingContact() bool`

GetNeedsShippingContact returns the NeedsShippingContact field if non-nil, zero value otherwise.

### GetNeedsShippingContactOk

`func (o *OrderRequest) GetNeedsShippingContactOk() (*bool, bool)`

GetNeedsShippingContactOk returns a tuple with the NeedsShippingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsShippingContact

`func (o *OrderRequest) SetNeedsShippingContact(v bool)`

SetNeedsShippingContact sets NeedsShippingContact field to given value.

### HasNeedsShippingContact

`func (o *OrderRequest) HasNeedsShippingContact() bool`

HasNeedsShippingContact returns a boolean if a field has been set.

### GetPreAuthorize

`func (o *OrderRequest) GetPreAuthorize() bool`

GetPreAuthorize returns the PreAuthorize field if non-nil, zero value otherwise.

### GetPreAuthorizeOk

`func (o *OrderRequest) GetPreAuthorizeOk() (*bool, bool)`

GetPreAuthorizeOk returns a tuple with the PreAuthorize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthorize

`func (o *OrderRequest) SetPreAuthorize(v bool)`

SetPreAuthorize sets PreAuthorize field to given value.

### HasPreAuthorize

`func (o *OrderRequest) HasPreAuthorize() bool`

HasPreAuthorize returns a boolean if a field has been set.

### GetProcessingMode

`func (o *OrderRequest) GetProcessingMode() string`

GetProcessingMode returns the ProcessingMode field if non-nil, zero value otherwise.

### GetProcessingModeOk

`func (o *OrderRequest) GetProcessingModeOk() (*string, bool)`

GetProcessingModeOk returns a tuple with the ProcessingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingMode

`func (o *OrderRequest) SetProcessingMode(v string)`

SetProcessingMode sets ProcessingMode field to given value.

### HasProcessingMode

`func (o *OrderRequest) HasProcessingMode() bool`

HasProcessingMode returns a boolean if a field has been set.

### GetReturnUrl

`func (o *OrderRequest) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *OrderRequest) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *OrderRequest) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *OrderRequest) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetShippingContact

`func (o *OrderRequest) GetShippingContact() CustomerShippingContacts`

GetShippingContact returns the ShippingContact field if non-nil, zero value otherwise.

### GetShippingContactOk

`func (o *OrderRequest) GetShippingContactOk() (*CustomerShippingContacts, bool)`

GetShippingContactOk returns a tuple with the ShippingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingContact

`func (o *OrderRequest) SetShippingContact(v CustomerShippingContacts)`

SetShippingContact sets ShippingContact field to given value.

### HasShippingContact

`func (o *OrderRequest) HasShippingContact() bool`

HasShippingContact returns a boolean if a field has been set.

### GetShippingLines

`func (o *OrderRequest) GetShippingLines() []ShippingRequest`

GetShippingLines returns the ShippingLines field if non-nil, zero value otherwise.

### GetShippingLinesOk

`func (o *OrderRequest) GetShippingLinesOk() (*[]ShippingRequest, bool)`

GetShippingLinesOk returns a tuple with the ShippingLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLines

`func (o *OrderRequest) SetShippingLines(v []ShippingRequest)`

SetShippingLines sets ShippingLines field to given value.

### HasShippingLines

`func (o *OrderRequest) HasShippingLines() bool`

HasShippingLines returns a boolean if a field has been set.

### GetTaxLines

`func (o *OrderRequest) GetTaxLines() []OrderTaxRequest`

GetTaxLines returns the TaxLines field if non-nil, zero value otherwise.

### GetTaxLinesOk

`func (o *OrderRequest) GetTaxLinesOk() (*[]OrderTaxRequest, bool)`

GetTaxLinesOk returns a tuple with the TaxLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxLines

`func (o *OrderRequest) SetTaxLines(v []OrderTaxRequest)`

SetTaxLines sets TaxLines field to given value.

### HasTaxLines

`func (o *OrderRequest) HasTaxLines() bool`

HasTaxLines returns a boolean if a field has been set.

### GetThreeDsMode

`func (o *OrderRequest) GetThreeDsMode() string`

GetThreeDsMode returns the ThreeDsMode field if non-nil, zero value otherwise.

### GetThreeDsModeOk

`func (o *OrderRequest) GetThreeDsModeOk() (*string, bool)`

GetThreeDsModeOk returns a tuple with the ThreeDsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDsMode

`func (o *OrderRequest) SetThreeDsMode(v string)`

SetThreeDsMode sets ThreeDsMode field to given value.

### HasThreeDsMode

`func (o *OrderRequest) HasThreeDsMode() bool`

HasThreeDsMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


