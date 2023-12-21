# OrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | The total amount to be collected in cents | [optional] 
**AmountRefunded** | Pointer to **int32** | The total amount refunded in cents | [optional] 
**Channel** | Pointer to [**ChargeResponseChannel**](ChargeResponseChannel.md) |  | [optional] 
**Charges** | Pointer to [**OrderResponseCharges**](OrderResponseCharges.md) |  | [optional] 
**Checkout** | Pointer to [**OrderResponseCheckout**](OrderResponseCheckout.md) |  | [optional] 
**CreatedAt** | Pointer to **int64** | The time at which the object was created in seconds since the Unix epoch | [optional] 
**Currency** | Pointer to **string** | The three-letter ISO 4217 currency code. The currency of the order. | [optional] 
**CustomerInfo** | Pointer to [**OrderResponseCustomerInfo**](OrderResponseCustomerInfo.md) |  | [optional] 
**DiscountLines** | Pointer to [**OrderResponseDiscountLines**](OrderResponseDiscountLines.md) |  | [optional] 
**FiscalEntity** | Pointer to [**NullableOrderFiscalEntityResponse**](OrderFiscalEntityResponse.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsRefundable** | Pointer to **bool** |  | [optional] 
**LineItems** | Pointer to [**OrderResponseProducts**](OrderResponseProducts.md) |  | [optional] 
**Livemode** | Pointer to **bool** | Whether the object exists in live mode or test mode | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format. | [optional] 
**NextAction** | Pointer to [**OrderNextActionResponse**](OrderNextActionResponse.md) |  | [optional] 
**Object** | Pointer to **string** | String representing the object’s type. Objects of the same type share the same value. | [optional] 
**PaymentStatus** | Pointer to **string** | The payment status of the order. | [optional] 
**ProcessingMode** | Pointer to **string** | Indicates the processing mode for the order, either ecommerce, recurrent or validation. | [optional] 
**ShippingContact** | Pointer to [**OrderResponseShippingContact**](OrderResponseShippingContact.md) |  | [optional] 
**UpdatedAt** | Pointer to **int64** | The time at which the object was last updated in seconds since the Unix epoch | [optional] 

## Methods

### NewOrderResponse

`func NewOrderResponse() *OrderResponse`

NewOrderResponse instantiates a new OrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderResponseWithDefaults

`func NewOrderResponseWithDefaults() *OrderResponse`

NewOrderResponseWithDefaults instantiates a new OrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *OrderResponse) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OrderResponse) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OrderResponse) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *OrderResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAmountRefunded

`func (o *OrderResponse) GetAmountRefunded() int32`

GetAmountRefunded returns the AmountRefunded field if non-nil, zero value otherwise.

### GetAmountRefundedOk

`func (o *OrderResponse) GetAmountRefundedOk() (*int32, bool)`

GetAmountRefundedOk returns a tuple with the AmountRefunded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRefunded

`func (o *OrderResponse) SetAmountRefunded(v int32)`

SetAmountRefunded sets AmountRefunded field to given value.

### HasAmountRefunded

`func (o *OrderResponse) HasAmountRefunded() bool`

HasAmountRefunded returns a boolean if a field has been set.

### GetChannel

`func (o *OrderResponse) GetChannel() ChargeResponseChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *OrderResponse) GetChannelOk() (*ChargeResponseChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *OrderResponse) SetChannel(v ChargeResponseChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *OrderResponse) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCharges

`func (o *OrderResponse) GetCharges() OrderResponseCharges`

GetCharges returns the Charges field if non-nil, zero value otherwise.

### GetChargesOk

`func (o *OrderResponse) GetChargesOk() (*OrderResponseCharges, bool)`

GetChargesOk returns a tuple with the Charges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharges

`func (o *OrderResponse) SetCharges(v OrderResponseCharges)`

SetCharges sets Charges field to given value.

### HasCharges

`func (o *OrderResponse) HasCharges() bool`

HasCharges returns a boolean if a field has been set.

### GetCheckout

`func (o *OrderResponse) GetCheckout() OrderResponseCheckout`

GetCheckout returns the Checkout field if non-nil, zero value otherwise.

### GetCheckoutOk

`func (o *OrderResponse) GetCheckoutOk() (*OrderResponseCheckout, bool)`

GetCheckoutOk returns a tuple with the Checkout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckout

`func (o *OrderResponse) SetCheckout(v OrderResponseCheckout)`

SetCheckout sets Checkout field to given value.

### HasCheckout

`func (o *OrderResponse) HasCheckout() bool`

HasCheckout returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrderResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrderResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrderResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrderResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *OrderResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OrderResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OrderResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *OrderResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerInfo

`func (o *OrderResponse) GetCustomerInfo() OrderResponseCustomerInfo`

GetCustomerInfo returns the CustomerInfo field if non-nil, zero value otherwise.

### GetCustomerInfoOk

`func (o *OrderResponse) GetCustomerInfoOk() (*OrderResponseCustomerInfo, bool)`

GetCustomerInfoOk returns a tuple with the CustomerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerInfo

`func (o *OrderResponse) SetCustomerInfo(v OrderResponseCustomerInfo)`

SetCustomerInfo sets CustomerInfo field to given value.

### HasCustomerInfo

`func (o *OrderResponse) HasCustomerInfo() bool`

HasCustomerInfo returns a boolean if a field has been set.

### GetDiscountLines

`func (o *OrderResponse) GetDiscountLines() OrderResponseDiscountLines`

GetDiscountLines returns the DiscountLines field if non-nil, zero value otherwise.

### GetDiscountLinesOk

`func (o *OrderResponse) GetDiscountLinesOk() (*OrderResponseDiscountLines, bool)`

GetDiscountLinesOk returns a tuple with the DiscountLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountLines

`func (o *OrderResponse) SetDiscountLines(v OrderResponseDiscountLines)`

SetDiscountLines sets DiscountLines field to given value.

### HasDiscountLines

`func (o *OrderResponse) HasDiscountLines() bool`

HasDiscountLines returns a boolean if a field has been set.

### GetFiscalEntity

`func (o *OrderResponse) GetFiscalEntity() OrderFiscalEntityResponse`

GetFiscalEntity returns the FiscalEntity field if non-nil, zero value otherwise.

### GetFiscalEntityOk

`func (o *OrderResponse) GetFiscalEntityOk() (*OrderFiscalEntityResponse, bool)`

GetFiscalEntityOk returns a tuple with the FiscalEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalEntity

`func (o *OrderResponse) SetFiscalEntity(v OrderFiscalEntityResponse)`

SetFiscalEntity sets FiscalEntity field to given value.

### HasFiscalEntity

`func (o *OrderResponse) HasFiscalEntity() bool`

HasFiscalEntity returns a boolean if a field has been set.

### SetFiscalEntityNil

`func (o *OrderResponse) SetFiscalEntityNil(b bool)`

 SetFiscalEntityNil sets the value for FiscalEntity to be an explicit nil

### UnsetFiscalEntity
`func (o *OrderResponse) UnsetFiscalEntity()`

UnsetFiscalEntity ensures that no value is present for FiscalEntity, not even an explicit nil
### GetId

`func (o *OrderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrderResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsRefundable

`func (o *OrderResponse) GetIsRefundable() bool`

GetIsRefundable returns the IsRefundable field if non-nil, zero value otherwise.

### GetIsRefundableOk

`func (o *OrderResponse) GetIsRefundableOk() (*bool, bool)`

GetIsRefundableOk returns a tuple with the IsRefundable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRefundable

`func (o *OrderResponse) SetIsRefundable(v bool)`

SetIsRefundable sets IsRefundable field to given value.

### HasIsRefundable

`func (o *OrderResponse) HasIsRefundable() bool`

HasIsRefundable returns a boolean if a field has been set.

### GetLineItems

`func (o *OrderResponse) GetLineItems() OrderResponseProducts`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *OrderResponse) GetLineItemsOk() (*OrderResponseProducts, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *OrderResponse) SetLineItems(v OrderResponseProducts)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *OrderResponse) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetLivemode

`func (o *OrderResponse) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *OrderResponse) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *OrderResponse) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *OrderResponse) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetMetadata

`func (o *OrderResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrderResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrderResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OrderResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNextAction

`func (o *OrderResponse) GetNextAction() OrderNextActionResponse`

GetNextAction returns the NextAction field if non-nil, zero value otherwise.

### GetNextActionOk

`func (o *OrderResponse) GetNextActionOk() (*OrderNextActionResponse, bool)`

GetNextActionOk returns a tuple with the NextAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAction

`func (o *OrderResponse) SetNextAction(v OrderNextActionResponse)`

SetNextAction sets NextAction field to given value.

### HasNextAction

`func (o *OrderResponse) HasNextAction() bool`

HasNextAction returns a boolean if a field has been set.

### GetObject

`func (o *OrderResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *OrderResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *OrderResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *OrderResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *OrderResponse) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *OrderResponse) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *OrderResponse) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *OrderResponse) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetProcessingMode

`func (o *OrderResponse) GetProcessingMode() string`

GetProcessingMode returns the ProcessingMode field if non-nil, zero value otherwise.

### GetProcessingModeOk

`func (o *OrderResponse) GetProcessingModeOk() (*string, bool)`

GetProcessingModeOk returns a tuple with the ProcessingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingMode

`func (o *OrderResponse) SetProcessingMode(v string)`

SetProcessingMode sets ProcessingMode field to given value.

### HasProcessingMode

`func (o *OrderResponse) HasProcessingMode() bool`

HasProcessingMode returns a boolean if a field has been set.

### GetShippingContact

`func (o *OrderResponse) GetShippingContact() OrderResponseShippingContact`

GetShippingContact returns the ShippingContact field if non-nil, zero value otherwise.

### GetShippingContactOk

`func (o *OrderResponse) GetShippingContactOk() (*OrderResponseShippingContact, bool)`

GetShippingContactOk returns a tuple with the ShippingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingContact

`func (o *OrderResponse) SetShippingContact(v OrderResponseShippingContact)`

SetShippingContact sets ShippingContact field to given value.

### HasShippingContact

`func (o *OrderResponse) HasShippingContact() bool`

HasShippingContact returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OrderResponse) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrderResponse) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrderResponse) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrderResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


