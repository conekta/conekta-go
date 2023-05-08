![GO api](conekta.png)
# Conekta Go API Library
[![Go Reference](https://pkg.go.dev/badge/github.com/conekta/conekta-go.svg)](https://pkg.go.dev/github.com/conekta/conekta-go)

This is the officially supported .NET library for using Conekta's APIs.
## Supported API versions
The library supports all APIs under the following services:

| API                                                                                         | Description | Service Name | Supported version |
|---------------------------------------------------------------------------------------------| ----------- |-------|-------------------|
| [Payments API](https://developers.conekta.com/reference)                  | Our classic integration for online payments. Current supported version | Payments API | **v2.1.0**        |

For more information, refer to our [documentation](https://developers.conekta.com/docs).

## Prerequisites
- [Conekta account](https://panel.conekta.com/)
- [API key](https://developers.conekta.com/docs/como-obtener-tus-api-keys).  your API credential .

## Installation
Download conekta-go package:
```
go get -u github.com/conekta/conekta-go
```
## Using the library

In order to submit http request to Conekta API you need to initialize the client. The following example makes a order request:
```go
package main

import (
    "context"
	"net/http"
	
    "github.com/conekta/conekta-go"
)

func main() {
	// Create a OrderRequest
	const acceptLanguage = "es"
	const accessToken = "Your merchant XAPI key"

	// create the http client
	config := conekta.NewConfiguration()
	client := conekta.NewAPIClient(config)

	ctx := context.WithValue(context.TODO(), conekta.ContextAccessToken, accessToken)

	// create customer
	customer := conekta.Customer{
		Name:  "test go",
		Phone: "+573143159063",
		Email: "test@conekta.com",
	}
	customerResponse, httpResponse, err := client.CustomersApi.CreateCustomer(ctx).
		Customer(customer).
		AcceptLanguage(acceptLanguage).
		Execute()
	if err != nil {
		panic(err)
	}
	if httpResponse.StatusCode != http.StatusCreated {
		panic("invalid response statusCode")
	}

	// Create OrderRequest
	chargeRequest := conekta.ChargeRequest{
		Amount:        conekta.PtrInt32(1555),
		PaymentMethod: *conekta.NewChargeRequestPaymentMethod("cash"),
	}
	productLine := conekta.Product{
		Name:      "toshiba",
		Quantity:  1,
		UnitPrice: 1555,
	}
	orderRequest := conekta.OrderRequest{
		Charges: []conekta.ChargeRequest{
			chargeRequest,
		},
		Currency: "MXN",
		CustomerInfo: conekta.OrderRequestCustomerInfo{
			CustomerInfoJustCustomerId: conekta.NewCustomerInfoJustCustomerId(customerResponse.Id),
		},
		LineItems: []conekta.Product{
			productLine,
		},
	}

	//Make the call to the service. This example code makes a call to /orders
	orderResponse, httpResponse, err := client.OrdersApi.CreateOrder(ctx).
		OrderRequest(orderRequest).
		AcceptLanguage(acceptLanguage).
		Execute()
	if err != nil {
		panic(err)
	}
	if httpResponse.StatusCode != http.StatusCreated {
		panic("invalid response statusCode")
	}
	println(*orderResponse)   
}
```

## Running tests
Navigate to conekta-go folder and run the following commands.
```
docker-compose up
go test ./...
```

## Documentation for API Endpoints

All URIs are relative to *https://api.conekta.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AntifraudApi* | [**CreateRuleBlacklist**](docs/AntifraudApi.md#createruleblacklist) | **Post** /antifraud/blacklists | Create blacklisted rule
*AntifraudApi* | [**CreateRuleWhitelist**](docs/AntifraudApi.md#createrulewhitelist) | **Post** /antifraud/whitelists | Create whitelisted rule
*AntifraudApi* | [**DeleteRuleBlacklist**](docs/AntifraudApi.md#deleteruleblacklist) | **Delete** /antifraud/blacklists/{id} | Delete blacklisted rule
*AntifraudApi* | [**DeleteRuleWhitelist**](docs/AntifraudApi.md#deleterulewhitelist) | **Delete** /antifraud/whitelists/{id} | Delete whitelisted rule
*AntifraudApi* | [**GetRuleBlacklist**](docs/AntifraudApi.md#getruleblacklist) | **Get** /antifraud/blacklists | Get list of blacklisted rules
*AntifraudApi* | [**GetRuleWhitelist**](docs/AntifraudApi.md#getrulewhitelist) | **Get** /antifraud/whitelists | Get a list of whitelisted rules
*ChargesApi* | [**OrdersCreateCharge**](docs/ChargesApi.md#orderscreatecharge) | **Post** /orders/{id}/charges | Create charge
*CompaniesApi* | [**GetCompanies**](docs/CompaniesApi.md#getcompanies) | **Get** /companies | Get List of Companies
*CompaniesApi* | [**GetCompany**](docs/CompaniesApi.md#getcompany) | **Get** /companies/{id} | Get Company
*CustomersApi* | [**CreateCustomer**](docs/CustomersApi.md#createcustomer) | **Post** /customers | Create customer
*CustomersApi* | [**CreateCustomerFiscalEntities**](docs/CustomersApi.md#createcustomerfiscalentities) | **Post** /customers/{id}/fiscal_entities | Create Fiscal Entity
*CustomersApi* | [**DeleteCustomerById**](docs/CustomersApi.md#deletecustomerbyid) | **Delete** /customers/{id} | Delete Customer
*CustomersApi* | [**GetCustomerById**](docs/CustomersApi.md#getcustomerbyid) | **Get** /customers/{id} | Get Customer
*CustomersApi* | [**GetCustomers**](docs/CustomersApi.md#getcustomers) | **Get** /customers | Get a list of customers
*CustomersApi* | [**UpdateCustomer**](docs/CustomersApi.md#updatecustomer) | **Put** /customers/{id} | Update customer
*CustomersApi* | [**UpdateCustomerFiscalEntities**](docs/CustomersApi.md#updatecustomerfiscalentities) | **Put** /customers/{id}/fiscal_entities/{fiscal_entities_id} | Update  Fiscal Entity
*DiscountsApi* | [**OrdersCreateDiscountLine**](docs/DiscountsApi.md#orderscreatediscountline) | **Post** /orders/{id}/discount_lines | Create Discount
*DiscountsApi* | [**OrdersDeleteDiscountLines**](docs/DiscountsApi.md#ordersdeletediscountlines) | **Delete** /orders/{id}/discount_lines/{discount_lines_id} | Delete Discount
*DiscountsApi* | [**OrdersUpdateDiscountLines**](docs/DiscountsApi.md#ordersupdatediscountlines) | **Put** /orders/{id}/discount_lines/{discount_lines_id} | Update Discount
*EventsApi* | [**GetEvent**](docs/EventsApi.md#getevent) | **Get** /events/{id} | Get Event
*EventsApi* | [**GetEvents**](docs/EventsApi.md#getevents) | **Get** /events | Get list of Events
*LogsApi* | [**GetLogById**](docs/LogsApi.md#getlogbyid) | **Get** /logs/{id} | Get Log
*LogsApi* | [**GetLogs**](docs/LogsApi.md#getlogs) | **Get** /logs | Get List Of Logs
*OrdersApi* | [**CreateOrder**](docs/OrdersApi.md#createorder) | **Post** /orders | Create order
*OrdersApi* | [**GetOrderById**](docs/OrdersApi.md#getorderbyid) | **Get** /orders/{id} | Get Order
*OrdersApi* | [**GetOrders**](docs/OrdersApi.md#getorders) | **Get** /orders | Get a list of Orders
*OrdersApi* | [**OrderRefund**](docs/OrdersApi.md#orderrefund) | **Post** /orders/{id}/refunds | Refund Order
*OrdersApi* | [**OrdersCreateCapture**](docs/OrdersApi.md#orderscreatecapture) | **Post** /orders/{id}/capture | Capture Order
*OrdersApi* | [**UpdateOrder**](docs/OrdersApi.md#updateorder) | **Put** /orders/{id} | Update Order
*PaymentLinkApi* | [**CancelCheckout**](docs/PaymentLinkApi.md#cancelcheckout) | **Put** /checkouts/{id}/cancel | Cancel Payment Link
*PaymentLinkApi* | [**CreateCheckout**](docs/PaymentLinkApi.md#createcheckout) | **Post** /checkouts | Create Unique Payment Link
*PaymentLinkApi* | [**EmailCheckout**](docs/PaymentLinkApi.md#emailcheckout) | **Post** /checkouts/{id}/email | Send an email
*PaymentLinkApi* | [**GetCheckout**](docs/PaymentLinkApi.md#getcheckout) | **Get** /checkouts/{id} | Get a payment link by ID
*PaymentLinkApi* | [**GetCheckouts**](docs/PaymentLinkApi.md#getcheckouts) | **Get** /checkouts | Get a list of payment links
*PaymentLinkApi* | [**SmsCheckout**](docs/PaymentLinkApi.md#smscheckout) | **Post** /checkouts/{id}/sms | Send an sms
*PaymentMethodsApi* | [**CreateCustomerPaymentMethods**](docs/PaymentMethodsApi.md#createcustomerpaymentmethods) | **Post** /customers/{id}/payment_sources | Create Payment Method
*PaymentMethodsApi* | [**DeleteCustomerPaymentMethods**](docs/PaymentMethodsApi.md#deletecustomerpaymentmethods) | **Delete** /customers/{id}/payment_sources/{payment_method_id} | Delete Payment Method
*PaymentMethodsApi* | [**GetCustomerPaymentMethods**](docs/PaymentMethodsApi.md#getcustomerpaymentmethods) | **Get** /customers/{id}/payment_sources | Get Payment Methods
*PaymentMethodsApi* | [**UpdateCustomerPaymentMethods**](docs/PaymentMethodsApi.md#updatecustomerpaymentmethods) | **Put** /customers/{id}/payment_sources/{payment_method_id} | Update Payment Method
*PlansApi* | [**CreatePlan**](docs/PlansApi.md#createplan) | **Post** /plans | Create Plan
*PlansApi* | [**DeletePlan**](docs/PlansApi.md#deleteplan) | **Delete** /plans/{id} | Delete Plan
*PlansApi* | [**GetPlan**](docs/PlansApi.md#getplan) | **Get** /plans/{id} | Get Plan
*PlansApi* | [**GetPlans**](docs/PlansApi.md#getplans) | **Get** /plans | Get A List of Plans
*PlansApi* | [**UpdatePlan**](docs/PlansApi.md#updateplan) | **Put** /plans/{id} | Update Plan
*ProductsApi* | [**OrdersCreateProduct**](docs/ProductsApi.md#orderscreateproduct) | **Post** /orders/{id}/line_items | Create Product
*ProductsApi* | [**OrdersDeleteProduct**](docs/ProductsApi.md#ordersdeleteproduct) | **Delete** /orders/{id}/line_items/{line_item_id} | Delete Product
*ProductsApi* | [**OrdersUpdateProduct**](docs/ProductsApi.md#ordersupdateproduct) | **Put** /orders/{id}/line_items/{line_item_id} | Update Product
*ShippingContactsApi* | [**CreateCustomerShippingContacts**](docs/ShippingContactsApi.md#createcustomershippingcontacts) | **Post** /customers/{id}/shipping_contacts | Create a shipping contacts
*ShippingContactsApi* | [**DeleteCustomerShippingContacts**](docs/ShippingContactsApi.md#deletecustomershippingcontacts) | **Delete** /customers/{id}/shipping_contacts/{shipping_contacts_id} | Delete shipping contacts
*ShippingContactsApi* | [**UpdateCustomerShippingContacts**](docs/ShippingContactsApi.md#updatecustomershippingcontacts) | **Put** /customers/{id}/shipping_contacts/{shipping_contacts_id} | Update shipping contacts
*ShippingsApi* | [**OrdersCreateShipping**](docs/ShippingsApi.md#orderscreateshipping) | **Post** /orders/{id}/shipping_lines | Create Shipping
*ShippingsApi* | [**OrdersDeleteShipping**](docs/ShippingsApi.md#ordersdeleteshipping) | **Delete** /orders/{id}/shipping_lines/{shipping_id} | Delete Shipping
*ShippingsApi* | [**OrdersUpdateShipping**](docs/ShippingsApi.md#ordersupdateshipping) | **Put** /orders/{id}/shipping_lines/{shipping_id} | Update Shipping
*SubscriptionsApi* | [**CancelSubscription**](docs/SubscriptionsApi.md#cancelsubscription) | **Post** /customers/{id}/subscription/cancel | Cancel Subscription
*SubscriptionsApi* | [**CreateSubscription**](docs/SubscriptionsApi.md#createsubscription) | **Post** /customers/{id}/subscription | Create Subscription
*SubscriptionsApi* | [**GetAllEventsFromSubscription**](docs/SubscriptionsApi.md#getalleventsfromsubscription) | **Get** /customers/{id}/subscription/events | Get Events By Subscription
*SubscriptionsApi* | [**GetSubscription**](docs/SubscriptionsApi.md#getsubscription) | **Get** /customers/{id}/subscription | Get Subscription
*SubscriptionsApi* | [**PauseSubscription**](docs/SubscriptionsApi.md#pausesubscription) | **Post** /customers/{id}/subscription/pause | Pause Subscription
*SubscriptionsApi* | [**ResumeSubscription**](docs/SubscriptionsApi.md#resumesubscription) | **Post** /customers/{id}/subscription/resume | Resume Subscription
*SubscriptionsApi* | [**UpdateSubscription**](docs/SubscriptionsApi.md#updatesubscription) | **Put** /customers/{id}/subscription | Update Subscription
*TaxesApi* | [**OrdersCreateTaxes**](docs/TaxesApi.md#orderscreatetaxes) | **Post** /orders/{id}/tax_lines | Create Tax
*TaxesApi* | [**OrdersDeleteTaxes**](docs/TaxesApi.md#ordersdeletetaxes) | **Delete** /orders/{id}/tax_lines/{tax_id} | Delete Tax
*TaxesApi* | [**OrdersUpdateTaxes**](docs/TaxesApi.md#ordersupdatetaxes) | **Put** /orders/{id}/tax_lines/{tax_id} | Update Tax
*TokensApi* | [**CreateToken**](docs/TokensApi.md#createtoken) | **Post** /tokens | Create Token
*WebhooksApi* | [**CreateWebhook**](docs/WebhooksApi.md#createwebhook) | **Post** /webhooks | Create Webhook
*WebhooksApi* | [**DeleteWebhook**](docs/WebhooksApi.md#deletewebhook) | **Delete** /webhooks/{id} | Delete Webhook
*WebhooksApi* | [**GetWebhook**](docs/WebhooksApi.md#getwebhook) | **Get** /webhooks/{id} | Get Webhook
*WebhooksApi* | [**GetWebhooks**](docs/WebhooksApi.md#getwebhooks) | **Get** /webhooks | Get List of Webhooks
*WebhooksApi* | [**UpdateWebhook**](docs/WebhooksApi.md#updatewebhook) | **Put** /webhooks/{id} | Update Webhook


## Documentation For Models

- [BlacklistRuleResponse](docs/BlacklistRuleResponse.md)
- [ChargeDataPaymentMethodBankTransferResponse](docs/ChargeDataPaymentMethodBankTransferResponse.md)
- [ChargeDataPaymentMethodCardResponse](docs/ChargeDataPaymentMethodCardResponse.md)
- [ChargeDataPaymentMethodCashResponse](docs/ChargeDataPaymentMethodCashResponse.md)
- [ChargeOrderResponse](docs/ChargeOrderResponse.md)
- [ChargeOrderResponseChannel](docs/ChargeOrderResponseChannel.md)
- [ChargeOrderResponsePaymentMethod](docs/ChargeOrderResponsePaymentMethod.md)
- [ChargeRequest](docs/ChargeRequest.md)
- [ChargeRequestPaymentMethod](docs/ChargeRequestPaymentMethod.md)
- [ChargeResponse](docs/ChargeResponse.md)
- [ChargeResponseRefunds](docs/ChargeResponseRefunds.md)
- [ChargeResponseRefundsAllOf](docs/ChargeResponseRefundsAllOf.md)
- [ChargeResponseRefundsData](docs/ChargeResponseRefundsData.md)
- [ChargesDataResponse](docs/ChargesDataResponse.md)
- [Checkout](docs/Checkout.md)
- [CheckoutOrderTemplate](docs/CheckoutOrderTemplate.md)
- [CheckoutRequest](docs/CheckoutRequest.md)
- [CheckoutResponse](docs/CheckoutResponse.md)
- [CheckoutsResponse](docs/CheckoutsResponse.md)
- [CheckoutsResponseAllOf](docs/CheckoutsResponseAllOf.md)
- [CompanyFiscalInfoAddressResponse](docs/CompanyFiscalInfoAddressResponse.md)
- [CompanyFiscalInfoResponse](docs/CompanyFiscalInfoResponse.md)
- [CompanyPayoutDestinationResponse](docs/CompanyPayoutDestinationResponse.md)
- [CompanyResponse](docs/CompanyResponse.md)
- [CreateCustomerFiscalEntitiesResponse](docs/CreateCustomerFiscalEntitiesResponse.md)
- [CreateCustomerFiscalEntitiesResponseAllOf](docs/CreateCustomerFiscalEntitiesResponseAllOf.md)
- [CreateCustomerPaymentMethodsRequest](docs/CreateCustomerPaymentMethodsRequest.md)
- [CreateCustomerPaymentMethodsResponse](docs/CreateCustomerPaymentMethodsResponse.md)
- [CreateRiskRulesData](docs/CreateRiskRulesData.md)
- [Customer](docs/Customer.md)
- [CustomerAddress](docs/CustomerAddress.md)
- [CustomerAntifraudInfo](docs/CustomerAntifraudInfo.md)
- [CustomerAntifraudInfoResponse](docs/CustomerAntifraudInfoResponse.md)
- [CustomerFiscalEntitiesDataResponse](docs/CustomerFiscalEntitiesDataResponse.md)
- [CustomerFiscalEntitiesRequest](docs/CustomerFiscalEntitiesRequest.md)
- [CustomerFiscalEntitiesRequestAddress](docs/CustomerFiscalEntitiesRequestAddress.md)
- [CustomerFiscalEntitiesResponse](docs/CustomerFiscalEntitiesResponse.md)
- [CustomerFiscalEntitiesResponseAllOf](docs/CustomerFiscalEntitiesResponseAllOf.md)
- [CustomerInfo](docs/CustomerInfo.md)
- [CustomerInfoJustCustomerId](docs/CustomerInfoJustCustomerId.md)
- [CustomerInfoJustCustomerIdResponse](docs/CustomerInfoJustCustomerIdResponse.md)
- [CustomerInfoResponse](docs/CustomerInfoResponse.md)
- [CustomerPaymentMethodRequest](docs/CustomerPaymentMethodRequest.md)
- [CustomerPaymentMethods](docs/CustomerPaymentMethods.md)
- [CustomerPaymentMethodsData](docs/CustomerPaymentMethodsData.md)
- [CustomerPaymentMethodsRequest](docs/CustomerPaymentMethodsRequest.md)
- [CustomerPaymentMethodsResponse](docs/CustomerPaymentMethodsResponse.md)
- [CustomerResponse](docs/CustomerResponse.md)
- [CustomerResponseShippingContacts](docs/CustomerResponseShippingContacts.md)
- [CustomerResponseShippingContactsAllOf](docs/CustomerResponseShippingContactsAllOf.md)
- [CustomerShippingContacts](docs/CustomerShippingContacts.md)
- [CustomerShippingContactsAddress](docs/CustomerShippingContactsAddress.md)
- [CustomerShippingContactsDataResponse](docs/CustomerShippingContactsDataResponse.md)
- [CustomerShippingContactsDataResponseAllOf](docs/CustomerShippingContactsDataResponseAllOf.md)
- [CustomerShippingContactsResponse](docs/CustomerShippingContactsResponse.md)
- [CustomerShippingContactsResponseAddress](docs/CustomerShippingContactsResponseAddress.md)
- [CustomerUpdateFiscalEntitiesRequest](docs/CustomerUpdateFiscalEntitiesRequest.md)
- [CustomerUpdateShippingContacts](docs/CustomerUpdateShippingContacts.md)
- [CustomersResponse](docs/CustomersResponse.md)
- [CustomersResponseAllOf](docs/CustomersResponseAllOf.md)
- [DeletedBlacklistRuleResponse](docs/DeletedBlacklistRuleResponse.md)
- [DeletedWhitelistRuleResponse](docs/DeletedWhitelistRuleResponse.md)
- [Details](docs/Details.md)
- [DetailsError](docs/DetailsError.md)
- [DiscountLinesDataResponse](docs/DiscountLinesDataResponse.md)
- [DiscountLinesResponse](docs/DiscountLinesResponse.md)
- [DiscountLinesResponseAllOf](docs/DiscountLinesResponseAllOf.md)
- [EmailCheckoutRequest](docs/EmailCheckoutRequest.md)
- [ErrorAllOf](docs/ErrorAllOf.md)
- [EventResponse](docs/EventResponse.md)
- [GetCompaniesResponse](docs/GetCompaniesResponse.md)
- [GetCompaniesResponseAllOf](docs/GetCompaniesResponseAllOf.md)
- [GetCustomerPaymentMethodDataResponse](docs/GetCustomerPaymentMethodDataResponse.md)
- [GetEventsResponse](docs/GetEventsResponse.md)
- [GetEventsResponseAllOf](docs/GetEventsResponseAllOf.md)
- [GetOrdersResponse](docs/GetOrdersResponse.md)
- [GetPaymentMethodResponse](docs/GetPaymentMethodResponse.md)
- [GetPaymentMethodResponseAllOf](docs/GetPaymentMethodResponseAllOf.md)
- [GetPlansResponse](docs/GetPlansResponse.md)
- [GetPlansResponseAllOf](docs/GetPlansResponseAllOf.md)
- [GetWebhooksResponse](docs/GetWebhooksResponse.md)
- [GetWebhooksResponseAllOf](docs/GetWebhooksResponseAllOf.md)
- [LogResponse](docs/LogResponse.md)
- [LogsResponse](docs/LogsResponse.md)
- [LogsResponseData](docs/LogsResponseData.md)
- [ModelError](docs/ModelError.md)
- [OrderCaptureRequest](docs/OrderCaptureRequest.md)
- [OrderDiscountLinesRequest](docs/OrderDiscountLinesRequest.md)
- [OrderRefundRequest](docs/OrderRefundRequest.md)
- [OrderRequest](docs/OrderRequest.md)
- [OrderRequestCustomerInfo](docs/OrderRequestCustomerInfo.md)
- [OrderResponse](docs/OrderResponse.md)
- [OrderResponseCharges](docs/OrderResponseCharges.md)
- [OrderResponseChargesAllOf](docs/OrderResponseChargesAllOf.md)
- [OrderResponseCheckout](docs/OrderResponseCheckout.md)
- [OrderResponseCustomerInfo](docs/OrderResponseCustomerInfo.md)
- [OrderResponseCustomerInfoAllOf](docs/OrderResponseCustomerInfoAllOf.md)
- [OrderResponseDiscountLines](docs/OrderResponseDiscountLines.md)
- [OrderResponseDiscountLinesAllOf](docs/OrderResponseDiscountLinesAllOf.md)
- [OrderResponseFiscalEntity](docs/OrderResponseFiscalEntity.md)
- [OrderResponseFiscalEntityAddress](docs/OrderResponseFiscalEntityAddress.md)
- [OrderResponseFiscalEntityAddressAllOf](docs/OrderResponseFiscalEntityAddressAllOf.md)
- [OrderResponseProducts](docs/OrderResponseProducts.md)
- [OrderResponseProductsAllOf](docs/OrderResponseProductsAllOf.md)
- [OrderResponseShippingContact](docs/OrderResponseShippingContact.md)
- [OrderResponseShippingContactAllOf](docs/OrderResponseShippingContactAllOf.md)
- [OrderTaxRequest](docs/OrderTaxRequest.md)
- [OrderUpdateRequest](docs/OrderUpdateRequest.md)
- [OrdersResponse](docs/OrdersResponse.md)
- [Page](docs/Page.md)
- [Pagination](docs/Pagination.md)
- [PaymentMethod](docs/PaymentMethod.md)
- [PaymentMethodBankTransfer](docs/PaymentMethodBankTransfer.md)
- [PaymentMethodCard](docs/PaymentMethodCard.md)
- [PaymentMethodCardRequest](docs/PaymentMethodCardRequest.md)
- [PaymentMethodCardRequestAllOf](docs/PaymentMethodCardRequestAllOf.md)
- [PaymentMethodCardResponse](docs/PaymentMethodCardResponse.md)
- [PaymentMethodCardResponseAllOf](docs/PaymentMethodCardResponseAllOf.md)
- [PaymentMethodCash](docs/PaymentMethodCash.md)
- [PaymentMethodCashRequest](docs/PaymentMethodCashRequest.md)
- [PaymentMethodCashRequestAllOf](docs/PaymentMethodCashRequestAllOf.md)
- [PaymentMethodCashResponse](docs/PaymentMethodCashResponse.md)
- [PaymentMethodCashResponseAllOf](docs/PaymentMethodCashResponseAllOf.md)
- [PaymentMethodResponse](docs/PaymentMethodResponse.md)
- [PaymentMethodSpeiRecurrent](docs/PaymentMethodSpeiRecurrent.md)
- [PaymentMethodSpeiRecurrentAllOf](docs/PaymentMethodSpeiRecurrentAllOf.md)
- [PaymentMethodSpeiRequest](docs/PaymentMethodSpeiRequest.md)
- [PlanRequest](docs/PlanRequest.md)
- [PlanResponse](docs/PlanResponse.md)
- [PlanUpdateRequest](docs/PlanUpdateRequest.md)
- [Product](docs/Product.md)
- [ProductDataResponse](docs/ProductDataResponse.md)
- [ProductDataResponseAllOf](docs/ProductDataResponseAllOf.md)
- [ProductOrderResponse](docs/ProductOrderResponse.md)
- [ProductOrderResponseAllOf](docs/ProductOrderResponseAllOf.md)
- [RiskRules](docs/RiskRules.md)
- [RiskRulesData](docs/RiskRulesData.md)
- [RiskRulesList](docs/RiskRulesList.md)
- [ShippingOrderResponse](docs/ShippingOrderResponse.md)
- [ShippingRequest](docs/ShippingRequest.md)
- [SmsCheckoutRequest](docs/SmsCheckoutRequest.md)
- [SubscriptionEventsResponse](docs/SubscriptionEventsResponse.md)
- [SubscriptionRequest](docs/SubscriptionRequest.md)
- [SubscriptionResponse](docs/SubscriptionResponse.md)
- [SubscriptionUpdateRequest](docs/SubscriptionUpdateRequest.md)
- [Token](docs/Token.md)
- [TokenCard](docs/TokenCard.md)
- [TokenCheckout](docs/TokenCheckout.md)
- [TokenResponse](docs/TokenResponse.md)
- [TokenResponseCheckout](docs/TokenResponseCheckout.md)
- [UpdateCustomer](docs/UpdateCustomer.md)
- [UpdateCustomerAntifraudInfo](docs/UpdateCustomerAntifraudInfo.md)
- [UpdateCustomerFiscalEntitiesResponse](docs/UpdateCustomerFiscalEntitiesResponse.md)
- [UpdateCustomerFiscalEntitiesResponseAllOf](docs/UpdateCustomerFiscalEntitiesResponseAllOf.md)
- [UpdateCustomerPaymentMethodsResponse](docs/UpdateCustomerPaymentMethodsResponse.md)
- [UpdateOrderDiscountLinesRequest](docs/UpdateOrderDiscountLinesRequest.md)
- [UpdateOrderTaxRequest](docs/UpdateOrderTaxRequest.md)
- [UpdateOrderTaxResponse](docs/UpdateOrderTaxResponse.md)
- [UpdateOrderTaxResponseAllOf](docs/UpdateOrderTaxResponseAllOf.md)
- [UpdatePaymentMethods](docs/UpdatePaymentMethods.md)
- [UpdateProduct](docs/UpdateProduct.md)
- [WebhookLog](docs/WebhookLog.md)
- [WebhookRequest](docs/WebhookRequest.md)
- [WebhookResponse](docs/WebhookResponse.md)
- [WebhookUpdateRequest](docs/WebhookUpdateRequest.md)
- [WhitelistlistRuleResponse](docs/WhitelistlistRuleResponse.md)

## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Contributing
We encourage you to contribute to this repository, so everyone can benefit from new features, bug fixes, and any other improvements.
Have a look at our [contributing guidelines](https://github.com/conekta/conekta-go/blob/main/CONTRIBUTING.md) to find out how to raise a pull request.

## Support
If you have a feature request, or spotted a bug or a technical problem, [create an issue here](https://github.com/conekta/conekta-go/issues/choose).

For other questions, [contact our Support Team](https://developers.conekta.com/discuss).

## Licence
This repository is available under the [MIT license](https://github.com/conekta/conekta-go/blob/master/LICENSE).

## See also
* [Conekta docs](https://developers.conekta.com/docs)