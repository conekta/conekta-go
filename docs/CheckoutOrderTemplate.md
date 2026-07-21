# CheckoutOrderTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | It is the currency in which the order will be created. It must be a valid ISO 4217 currency code. | 
**CustomerInfo** | Pointer to [**OrderRequestCustomerInfo**](OrderRequestCustomerInfo.md) |  | [optional] 
**LineItems** | [**[]Product**](Product.md) | They are the products to buy. Each contains the \&quot;unit price\&quot; and \&quot;quantity\&quot; parameters that are used to calculate the total amount of the order. | 
**Metadata** | Pointer to **map[string]interface{}** | It is a set of key-value pairs that you can attach to the order. It can be used to store additional information about the order in a structured format. | [optional] 
**TaxLines** | Pointer to [**[]OrderTaxRequest**](OrderTaxRequest.md) | List of [taxes](https://developers.conekta.com/v2.3.0/reference/orderscreatetaxes) that are applied to the order. | [optional] 
**DiscountLines** | Pointer to [**[]OrderDiscountLinesRequest**](OrderDiscountLinesRequest.md) | List of [discounts](https://developers.conekta.com/v2.3.0/reference/orderscreatediscountline) that are applied to the order. | [optional] 

## Methods

### NewCheckoutOrderTemplate

`func NewCheckoutOrderTemplate(currency string, lineItems []Product, ) *CheckoutOrderTemplate`

NewCheckoutOrderTemplate instantiates a new CheckoutOrderTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutOrderTemplateWithDefaults

`func NewCheckoutOrderTemplateWithDefaults() *CheckoutOrderTemplate`

NewCheckoutOrderTemplateWithDefaults instantiates a new CheckoutOrderTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *CheckoutOrderTemplate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CheckoutOrderTemplate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CheckoutOrderTemplate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCustomerInfo

`func (o *CheckoutOrderTemplate) GetCustomerInfo() OrderRequestCustomerInfo`

GetCustomerInfo returns the CustomerInfo field if non-nil, zero value otherwise.

### GetCustomerInfoOk

`func (o *CheckoutOrderTemplate) GetCustomerInfoOk() (*OrderRequestCustomerInfo, bool)`

GetCustomerInfoOk returns a tuple with the CustomerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerInfo

`func (o *CheckoutOrderTemplate) SetCustomerInfo(v OrderRequestCustomerInfo)`

SetCustomerInfo sets CustomerInfo field to given value.

### HasCustomerInfo

`func (o *CheckoutOrderTemplate) HasCustomerInfo() bool`

HasCustomerInfo returns a boolean if a field has been set.

### GetLineItems

`func (o *CheckoutOrderTemplate) GetLineItems() []Product`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *CheckoutOrderTemplate) GetLineItemsOk() (*[]Product, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *CheckoutOrderTemplate) SetLineItems(v []Product)`

SetLineItems sets LineItems field to given value.


### GetMetadata

`func (o *CheckoutOrderTemplate) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CheckoutOrderTemplate) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CheckoutOrderTemplate) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CheckoutOrderTemplate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTaxLines

`func (o *CheckoutOrderTemplate) GetTaxLines() []OrderTaxRequest`

GetTaxLines returns the TaxLines field if non-nil, zero value otherwise.

### GetTaxLinesOk

`func (o *CheckoutOrderTemplate) GetTaxLinesOk() (*[]OrderTaxRequest, bool)`

GetTaxLinesOk returns a tuple with the TaxLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxLines

`func (o *CheckoutOrderTemplate) SetTaxLines(v []OrderTaxRequest)`

SetTaxLines sets TaxLines field to given value.

### HasTaxLines

`func (o *CheckoutOrderTemplate) HasTaxLines() bool`

HasTaxLines returns a boolean if a field has been set.

### GetDiscountLines

`func (o *CheckoutOrderTemplate) GetDiscountLines() []OrderDiscountLinesRequest`

GetDiscountLines returns the DiscountLines field if non-nil, zero value otherwise.

### GetDiscountLinesOk

`func (o *CheckoutOrderTemplate) GetDiscountLinesOk() (*[]OrderDiscountLinesRequest, bool)`

GetDiscountLinesOk returns a tuple with the DiscountLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountLines

`func (o *CheckoutOrderTemplate) SetDiscountLines(v []OrderDiscountLinesRequest)`

SetDiscountLines sets DiscountLines field to given value.

### HasDiscountLines

`func (o *CheckoutOrderTemplate) HasDiscountLines() bool`

HasDiscountLines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


