![GO api](conekta.png)
# conekta Go API Library
[![Go Reference](https://pkg.go.dev/badge/github.com/conekta/conekta-go.svg)](https://pkg.go.dev/github.com/conekta/conekta-go)

Conekta sdk

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.2.0
- Package version: 7.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://github.com/conekta/openapi/issues](https://github.com/conekta/openapi/issues)

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
docker-compose up -d
go test -v --race ./...
```

## Documentation for API Endpoints

All URIs are relative to *https://api.conekta.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AntifraudAPI* | [**CreateRuleBlacklist**](docs/AntifraudAPI.md#createruleblacklist) | **Post** /antifraud/blacklists | Create blacklisted rule
*AntifraudAPI* | [**CreateRuleWhitelist**](docs/AntifraudAPI.md#createrulewhitelist) | **Post** /antifraud/whitelists | Create whitelisted rule
*AntifraudAPI* | [**DeleteRuleBlacklist**](docs/AntifraudAPI.md#deleteruleblacklist) | **Delete** /antifraud/blacklists/{id} | Delete blacklisted rule
*AntifraudAPI* | [**DeleteRuleWhitelist**](docs/AntifraudAPI.md#deleterulewhitelist) | **Delete** /antifraud/whitelists/{id} | Delete whitelisted rule
*AntifraudAPI* | [**GetRuleBlacklist**](docs/AntifraudAPI.md#getruleblacklist) | **Get** /antifraud/blacklists | Get list of blacklisted rules
*AntifraudAPI* | [**GetRuleWhitelist**](docs/AntifraudAPI.md#getrulewhitelist) | **Get** /antifraud/whitelists | Get a list of whitelisted rules
*ApiKeysAPI* | [**CreateApiKey**](docs/ApiKeysAPI.md#createapikey) | **Post** /api_keys | Create Api Key
*ApiKeysAPI* | [**DeleteApiKey**](docs/ApiKeysAPI.md#deleteapikey) | **Delete** /api_keys/{id} | Delete Api Key
*ApiKeysAPI* | [**GetApiKey**](docs/ApiKeysAPI.md#getapikey) | **Get** /api_keys/{id} | Get Api Key
*ApiKeysAPI* | [**GetApiKeys**](docs/ApiKeysAPI.md#getapikeys) | **Get** /api_keys | Get list of Api Keys
*ApiKeysAPI* | [**UpdateApiKey**](docs/ApiKeysAPI.md#updateapikey) | **Put** /api_keys/{id} | Update Api Key
*BalancesAPI* | [**GetBalance**](docs/BalancesAPI.md#getbalance) | **Get** /balance | Get a company&#39;s balance
*ChargesAPI* | [**GetCharges**](docs/ChargesAPI.md#getcharges) | **Get** /charges | Get A List of Charges
*ChargesAPI* | [**OrdersCreateCharge**](docs/ChargesAPI.md#orderscreatecharge) | **Post** /orders/{id}/charges | Create charge
*ChargesAPI* | [**OrdersCreateCharges**](docs/ChargesAPI.md#orderscreatecharges) | **Post** /orders/{id}/add_charges | Create charges
*ChargesAPI* | [**UpdateCharge**](docs/ChargesAPI.md#updatecharge) | **Put** /charges/{id} | Update a charge
*CompaniesAPI* | [**GetCompanies**](docs/CompaniesAPI.md#getcompanies) | **Get** /companies | Get List of Companies
*CompaniesAPI* | [**GetCompany**](docs/CompaniesAPI.md#getcompany) | **Get** /companies/{id} | Get Company
*CustomersAPI* | [**CreateCustomer**](docs/CustomersAPI.md#createcustomer) | **Post** /customers | Create customer
*CustomersAPI* | [**CreateCustomerFiscalEntities**](docs/CustomersAPI.md#createcustomerfiscalentities) | **Post** /customers/{id}/fiscal_entities | Create Fiscal Entity
*CustomersAPI* | [**DeleteCustomerById**](docs/CustomersAPI.md#deletecustomerbyid) | **Delete** /customers/{id} | Delete Customer
*CustomersAPI* | [**GetCustomerById**](docs/CustomersAPI.md#getcustomerbyid) | **Get** /customers/{id} | Get Customer
*CustomersAPI* | [**GetCustomers**](docs/CustomersAPI.md#getcustomers) | **Get** /customers | Get a list of customers
*CustomersAPI* | [**UpdateCustomer**](docs/CustomersAPI.md#updatecustomer) | **Put** /customers/{id} | Update customer
*CustomersAPI* | [**UpdateCustomerFiscalEntities**](docs/CustomersAPI.md#updatecustomerfiscalentities) | **Put** /customers/{id}/fiscal_entities/{fiscal_entities_id} | Update  Fiscal Entity
*DiscountsAPI* | [**OrdersCreateDiscountLine**](docs/DiscountsAPI.md#orderscreatediscountline) | **Post** /orders/{id}/discount_lines | Create Discount
*DiscountsAPI* | [**OrdersDeleteDiscountLines**](docs/DiscountsAPI.md#ordersdeletediscountlines) | **Delete** /orders/{id}/discount_lines/{discount_lines_id} | Delete Discount
*DiscountsAPI* | [**OrdersGetDiscountLine**](docs/DiscountsAPI.md#ordersgetdiscountline) | **Get** /orders/{id}/discount_lines/{discount_lines_id} | Get Discount
*DiscountsAPI* | [**OrdersGetDiscountLines**](docs/DiscountsAPI.md#ordersgetdiscountlines) | **Get** /orders/{id}/discount_lines | Get a List of Discount
*DiscountsAPI* | [**OrdersUpdateDiscountLines**](docs/DiscountsAPI.md#ordersupdatediscountlines) | **Put** /orders/{id}/discount_lines/{discount_lines_id} | Update Discount
*EventsAPI* | [**GetEvent**](docs/EventsAPI.md#getevent) | **Get** /events/{id} | Get Event
*EventsAPI* | [**GetEvents**](docs/EventsAPI.md#getevents) | **Get** /events | Get list of Events
*EventsAPI* | [**ResendEvent**](docs/EventsAPI.md#resendevent) | **Post** /events/{event_id}/resend | Resend Event
*LogsAPI* | [**GetLogById**](docs/LogsAPI.md#getlogbyid) | **Get** /logs/{id} | Get Log
*LogsAPI* | [**GetLogs**](docs/LogsAPI.md#getlogs) | **Get** /logs | Get List Of Logs
*OrdersAPI* | [**CancelOrder**](docs/OrdersAPI.md#cancelorder) | **Post** /orders/{id}/cancel | Cancel Order
*OrdersAPI* | [**CreateOrder**](docs/OrdersAPI.md#createorder) | **Post** /orders | Create order
*OrdersAPI* | [**GetOrderById**](docs/OrdersAPI.md#getorderbyid) | **Get** /orders/{id} | Get Order
*OrdersAPI* | [**GetOrders**](docs/OrdersAPI.md#getorders) | **Get** /orders | Get a list of Orders
*OrdersAPI* | [**OrderCancelRefund**](docs/OrdersAPI.md#ordercancelrefund) | **Delete** /orders/{id}/refunds/{refund_id} | Cancel Refund
*OrdersAPI* | [**OrderRefund**](docs/OrdersAPI.md#orderrefund) | **Post** /orders/{id}/refunds | Refund Order
*OrdersAPI* | [**OrdersCreateCapture**](docs/OrdersAPI.md#orderscreatecapture) | **Post** /orders/{id}/capture | Capture Order
*OrdersAPI* | [**UpdateOrder**](docs/OrdersAPI.md#updateorder) | **Put** /orders/{id} | Update Order
*PaymentLinkAPI* | [**CancelCheckout**](docs/PaymentLinkAPI.md#cancelcheckout) | **Put** /checkouts/{id}/cancel | Cancel Payment Link
*PaymentLinkAPI* | [**CreateCheckout**](docs/PaymentLinkAPI.md#createcheckout) | **Post** /checkouts | Create Unique Payment Link
*PaymentLinkAPI* | [**EmailCheckout**](docs/PaymentLinkAPI.md#emailcheckout) | **Post** /checkouts/{id}/email | Send an email
*PaymentLinkAPI* | [**GetCheckout**](docs/PaymentLinkAPI.md#getcheckout) | **Get** /checkouts/{id} | Get a payment link by ID
*PaymentLinkAPI* | [**GetCheckouts**](docs/PaymentLinkAPI.md#getcheckouts) | **Get** /checkouts | Get a list of payment links
*PaymentLinkAPI* | [**SmsCheckout**](docs/PaymentLinkAPI.md#smscheckout) | **Post** /checkouts/{id}/sms | Send an sms
*PaymentMethodsAPI* | [**CreateCustomerPaymentMethods**](docs/PaymentMethodsAPI.md#createcustomerpaymentmethods) | **Post** /customers/{id}/payment_sources | Create Payment Method
*PaymentMethodsAPI* | [**DeleteCustomerPaymentMethods**](docs/PaymentMethodsAPI.md#deletecustomerpaymentmethods) | **Delete** /customers/{id}/payment_sources/{payment_method_id} | Delete Payment Method
*PaymentMethodsAPI* | [**GetCustomerPaymentMethods**](docs/PaymentMethodsAPI.md#getcustomerpaymentmethods) | **Get** /customers/{id}/payment_sources | Get Payment Methods
*PaymentMethodsAPI* | [**UpdateCustomerPaymentMethods**](docs/PaymentMethodsAPI.md#updatecustomerpaymentmethods) | **Put** /customers/{id}/payment_sources/{payment_method_id} | Update Payment Method
*PayoutOrdersAPI* | [**CancelPayoutOrderById**](docs/PayoutOrdersAPI.md#cancelpayoutorderbyid) | **Put** /payout_orders/{id}/cancel | Cancel Payout Order
*PayoutOrdersAPI* | [**CreatePayoutOrder**](docs/PayoutOrdersAPI.md#createpayoutorder) | **Post** /payout_orders | Create payout order
*PayoutOrdersAPI* | [**GetPayoutOrderById**](docs/PayoutOrdersAPI.md#getpayoutorderbyid) | **Get** /payout_orders/{id} | Get Payout Order
*PayoutOrdersAPI* | [**GetPayoutOrders**](docs/PayoutOrdersAPI.md#getpayoutorders) | **Get** /payout_orders | Get a list of Payout Orders
*PlansAPI* | [**CreatePlan**](docs/PlansAPI.md#createplan) | **Post** /plans | Create Plan
*PlansAPI* | [**DeletePlan**](docs/PlansAPI.md#deleteplan) | **Delete** /plans/{id} | Delete Plan
*PlansAPI* | [**GetPlan**](docs/PlansAPI.md#getplan) | **Get** /plans/{id} | Get Plan
*PlansAPI* | [**GetPlans**](docs/PlansAPI.md#getplans) | **Get** /plans | Get A List of Plans
*PlansAPI* | [**UpdatePlan**](docs/PlansAPI.md#updateplan) | **Put** /plans/{id} | Update Plan
*ProductsAPI* | [**OrdersCreateProduct**](docs/ProductsAPI.md#orderscreateproduct) | **Post** /orders/{id}/line_items | Create Product
*ProductsAPI* | [**OrdersDeleteProduct**](docs/ProductsAPI.md#ordersdeleteproduct) | **Delete** /orders/{id}/line_items/{line_item_id} | Delete Product
*ProductsAPI* | [**OrdersUpdateProduct**](docs/ProductsAPI.md#ordersupdateproduct) | **Put** /orders/{id}/line_items/{line_item_id} | Update Product
*ShippingContactsAPI* | [**CreateCustomerShippingContacts**](docs/ShippingContactsAPI.md#createcustomershippingcontacts) | **Post** /customers/{id}/shipping_contacts | Create a shipping contacts
*ShippingContactsAPI* | [**DeleteCustomerShippingContacts**](docs/ShippingContactsAPI.md#deletecustomershippingcontacts) | **Delete** /customers/{id}/shipping_contacts/{shipping_contacts_id} | Delete shipping contacts
*ShippingContactsAPI* | [**UpdateCustomerShippingContacts**](docs/ShippingContactsAPI.md#updatecustomershippingcontacts) | **Put** /customers/{id}/shipping_contacts/{shipping_contacts_id} | Update shipping contacts
*ShippingsAPI* | [**OrdersCreateShipping**](docs/ShippingsAPI.md#orderscreateshipping) | **Post** /orders/{id}/shipping_lines | Create Shipping
*ShippingsAPI* | [**OrdersDeleteShipping**](docs/ShippingsAPI.md#ordersdeleteshipping) | **Delete** /orders/{id}/shipping_lines/{shipping_id} | Delete Shipping
*ShippingsAPI* | [**OrdersUpdateShipping**](docs/ShippingsAPI.md#ordersupdateshipping) | **Put** /orders/{id}/shipping_lines/{shipping_id} | Update Shipping
*SubscriptionsAPI* | [**CancelSubscription**](docs/SubscriptionsAPI.md#cancelsubscription) | **Post** /customers/{id}/subscription/cancel | Cancel Subscription [Deprecated]
*SubscriptionsAPI* | [**CreateSubscription**](docs/SubscriptionsAPI.md#createsubscription) | **Post** /customers/{id}/subscription | Create Subscription [Deprecated]
*SubscriptionsAPI* | [**GetSubscription**](docs/SubscriptionsAPI.md#getsubscription) | **Get** /customers/{id}/subscription | Get Subscription [Deprecated]
*SubscriptionsAPI* | [**GetSubscriptionEvents**](docs/SubscriptionsAPI.md#getsubscriptionevents) | **Get** /customers/{id}/subscription/events | Get Subscription Events [Deprecated]
*SubscriptionsAPI* | [**PauseSubscription**](docs/SubscriptionsAPI.md#pausesubscription) | **Post** /customers/{id}/subscription/pause | Pause Subscription [Deprecated]
*SubscriptionsAPI* | [**ResumeSubscription**](docs/SubscriptionsAPI.md#resumesubscription) | **Post** /customers/{id}/subscription/resume | Resume Subscription [Deprecated]
*SubscriptionsAPI* | [**SubscriptionCancel**](docs/SubscriptionsAPI.md#subscriptioncancel) | **Post** /customers/{customer_id}/subscriptions/{id}/cancel | Cancel Subscription
*SubscriptionsAPI* | [**SubscriptionCreate**](docs/SubscriptionsAPI.md#subscriptioncreate) | **Post** /customers/{customer_id}/subscriptions | Create Subscription
*SubscriptionsAPI* | [**SubscriptionEvents**](docs/SubscriptionsAPI.md#subscriptionevents) | **Get** /customers/{customer_id}/subscriptions/{id}/events | Get Subscription Events
*SubscriptionsAPI* | [**SubscriptionList**](docs/SubscriptionsAPI.md#subscriptionlist) | **Get** /customers/{customer_id}/subscriptions | List Subscriptions
*SubscriptionsAPI* | [**SubscriptionPause**](docs/SubscriptionsAPI.md#subscriptionpause) | **Post** /customers/{customer_id}/subscriptions/{id}/pause | Pause Subscription
*SubscriptionsAPI* | [**SubscriptionResume**](docs/SubscriptionsAPI.md#subscriptionresume) | **Post** /customers/{customer_id}/subscriptions/{id}/resume | Resume Subscription
*SubscriptionsAPI* | [**SubscriptionUpdate**](docs/SubscriptionsAPI.md#subscriptionupdate) | **Put** /customers/{customer_id}/subscriptions/{id} | Update Subscription
*SubscriptionsAPI* | [**SubscriptionsGet**](docs/SubscriptionsAPI.md#subscriptionsget) | **Get** /customers/{customer_id}/subscriptions/{id} | Get Subscription
*SubscriptionsAPI* | [**SubscriptionsRetry**](docs/SubscriptionsAPI.md#subscriptionsretry) | **Post** /customers/{customer_id}/subscriptions/{id}/retry | Retry Failed Payment
*SubscriptionsAPI* | [**UpdateSubscription**](docs/SubscriptionsAPI.md#updatesubscription) | **Put** /customers/{id}/subscription | Update Subscription [Deprecated]
*TaxesAPI* | [**OrdersCreateTaxes**](docs/TaxesAPI.md#orderscreatetaxes) | **Post** /orders/{id}/tax_lines | Create Tax
*TaxesAPI* | [**OrdersDeleteTaxes**](docs/TaxesAPI.md#ordersdeletetaxes) | **Delete** /orders/{id}/tax_lines/{tax_id} | Delete Tax
*TaxesAPI* | [**OrdersUpdateTaxes**](docs/TaxesAPI.md#ordersupdatetaxes) | **Put** /orders/{id}/tax_lines/{tax_id} | Update Tax
*TokensAPI* | [**CreateToken**](docs/TokensAPI.md#createtoken) | **Post** /tokens | Create Token
*TransactionsAPI* | [**GetTransaction**](docs/TransactionsAPI.md#gettransaction) | **Get** /transactions/{id} | Get transaction
*TransactionsAPI* | [**GetTransactions**](docs/TransactionsAPI.md#gettransactions) | **Get** /transactions | Get List transactions
*TransfersAPI* | [**GetTransfer**](docs/TransfersAPI.md#gettransfer) | **Get** /transfers/{id} | Get Transfer
*TransfersAPI* | [**GetTransfers**](docs/TransfersAPI.md#gettransfers) | **Get** /transfers | Get a list of transfers
*WebhookKeysAPI* | [**CreateWebhookKey**](docs/WebhookKeysAPI.md#createwebhookkey) | **Post** /webhook_keys | Create Webhook Key
*WebhookKeysAPI* | [**DeleteWebhookKey**](docs/WebhookKeysAPI.md#deletewebhookkey) | **Delete** /webhook_keys/{id} | Delete Webhook key
*WebhookKeysAPI* | [**GetWebhookKey**](docs/WebhookKeysAPI.md#getwebhookkey) | **Get** /webhook_keys/{id} | Get Webhook Key
*WebhookKeysAPI* | [**GetWebhookKeys**](docs/WebhookKeysAPI.md#getwebhookkeys) | **Get** /webhook_keys | Get List of Webhook Keys
*WebhookKeysAPI* | [**UpdateWebhookKey**](docs/WebhookKeysAPI.md#updatewebhookkey) | **Put** /webhook_keys/{id} | Update Webhook Key
*WebhooksAPI* | [**CreateWebhook**](docs/WebhooksAPI.md#createwebhook) | **Post** /webhooks | Create Webhook
*WebhooksAPI* | [**DeleteWebhook**](docs/WebhooksAPI.md#deletewebhook) | **Delete** /webhooks/{id} | Delete Webhook
*WebhooksAPI* | [**GetWebhook**](docs/WebhooksAPI.md#getwebhook) | **Get** /webhooks/{id} | Get Webhook
*WebhooksAPI* | [**GetWebhooks**](docs/WebhooksAPI.md#getwebhooks) | **Get** /webhooks | Get List of Webhooks
*WebhooksAPI* | [**TestWebhook**](docs/WebhooksAPI.md#testwebhook) | **Post** /webhooks/{id}/test | Test Webhook
*WebhooksAPI* | [**UpdateWebhook**](docs/WebhooksAPI.md#updatewebhook) | **Put** /webhooks/{id} | Update Webhook


## Documentation For Models

 - [ApiKeyCreateResponse](docs/ApiKeyCreateResponse.md)
 - [ApiKeyRequest](docs/ApiKeyRequest.md)
 - [ApiKeyResponse](docs/ApiKeyResponse.md)
 - [ApiKeyResponseOnDelete](docs/ApiKeyResponseOnDelete.md)
 - [ApiKeyUpdateRequest](docs/ApiKeyUpdateRequest.md)
 - [BalanceCommonField](docs/BalanceCommonField.md)
 - [BalanceResponse](docs/BalanceResponse.md)
 - [BlacklistRuleResponse](docs/BlacklistRuleResponse.md)
 - [ChargeOrderResponse](docs/ChargeOrderResponse.md)
 - [ChargeOrderResponsePaymentMethod](docs/ChargeOrderResponsePaymentMethod.md)
 - [ChargeRequest](docs/ChargeRequest.md)
 - [ChargeRequestPaymentMethod](docs/ChargeRequestPaymentMethod.md)
 - [ChargeResponse](docs/ChargeResponse.md)
 - [ChargeResponseChannel](docs/ChargeResponseChannel.md)
 - [ChargeResponsePaymentMethod](docs/ChargeResponsePaymentMethod.md)
 - [ChargeResponseRefunds](docs/ChargeResponseRefunds.md)
 - [ChargeResponseRefundsData](docs/ChargeResponseRefundsData.md)
 - [ChargeUpdateRequest](docs/ChargeUpdateRequest.md)
 - [ChargesDataResponse](docs/ChargesDataResponse.md)
 - [ChargesOrderResponse](docs/ChargesOrderResponse.md)
 - [ChargesOrderResponseAllOfData](docs/ChargesOrderResponseAllOfData.md)
 - [Checkout](docs/Checkout.md)
 - [CheckoutOrderTemplate](docs/CheckoutOrderTemplate.md)
 - [CheckoutOrderTemplateCustomerInfo](docs/CheckoutOrderTemplateCustomerInfo.md)
 - [CheckoutRequest](docs/CheckoutRequest.md)
 - [CheckoutResponse](docs/CheckoutResponse.md)
 - [CheckoutsResponse](docs/CheckoutsResponse.md)
 - [CompanyFiscalInfoAddressResponse](docs/CompanyFiscalInfoAddressResponse.md)
 - [CompanyFiscalInfoResponse](docs/CompanyFiscalInfoResponse.md)
 - [CompanyPayoutDestinationResponse](docs/CompanyPayoutDestinationResponse.md)
 - [CompanyResponse](docs/CompanyResponse.md)
 - [CreateCustomerFiscalEntitiesResponse](docs/CreateCustomerFiscalEntitiesResponse.md)
 - [CreateCustomerPaymentMethodsRequest](docs/CreateCustomerPaymentMethodsRequest.md)
 - [CreateCustomerPaymentMethodsResponse](docs/CreateCustomerPaymentMethodsResponse.md)
 - [CreateRiskRulesData](docs/CreateRiskRulesData.md)
 - [Customer](docs/Customer.md)
 - [CustomerAddress](docs/CustomerAddress.md)
 - [CustomerAntifraudInfo](docs/CustomerAntifraudInfo.md)
 - [CustomerAntifraudInfoResponse](docs/CustomerAntifraudInfoResponse.md)
 - [CustomerFiscalEntitiesDataResponse](docs/CustomerFiscalEntitiesDataResponse.md)
 - [CustomerFiscalEntitiesRequest](docs/CustomerFiscalEntitiesRequest.md)
 - [CustomerFiscalEntitiesResponse](docs/CustomerFiscalEntitiesResponse.md)
 - [CustomerInfo](docs/CustomerInfo.md)
 - [CustomerInfoJustCustomerId](docs/CustomerInfoJustCustomerId.md)
 - [CustomerInfoJustCustomerIdResponse](docs/CustomerInfoJustCustomerIdResponse.md)
 - [CustomerPaymentMethodRequest](docs/CustomerPaymentMethodRequest.md)
 - [CustomerPaymentMethodsData](docs/CustomerPaymentMethodsData.md)
 - [CustomerPaymentMethodsRequest](docs/CustomerPaymentMethodsRequest.md)
 - [CustomerPaymentMethodsResponse](docs/CustomerPaymentMethodsResponse.md)
 - [CustomerResponse](docs/CustomerResponse.md)
 - [CustomerResponseShippingContacts](docs/CustomerResponseShippingContacts.md)
 - [CustomerShippingContacts](docs/CustomerShippingContacts.md)
 - [CustomerShippingContactsAddress](docs/CustomerShippingContactsAddress.md)
 - [CustomerShippingContactsDataResponse](docs/CustomerShippingContactsDataResponse.md)
 - [CustomerShippingContactsResponse](docs/CustomerShippingContactsResponse.md)
 - [CustomerShippingContactsResponseAddress](docs/CustomerShippingContactsResponseAddress.md)
 - [CustomerUpdateFiscalEntitiesRequest](docs/CustomerUpdateFiscalEntitiesRequest.md)
 - [CustomerUpdateShippingContacts](docs/CustomerUpdateShippingContacts.md)
 - [CustomersResponse](docs/CustomersResponse.md)
 - [DeleteApiKeysResponse](docs/DeleteApiKeysResponse.md)
 - [DeletedBlacklistRuleResponse](docs/DeletedBlacklistRuleResponse.md)
 - [DeletedWhitelistRuleResponse](docs/DeletedWhitelistRuleResponse.md)
 - [Details](docs/Details.md)
 - [DetailsError](docs/DetailsError.md)
 - [DiscountLinesDataResponse](docs/DiscountLinesDataResponse.md)
 - [DiscountLinesResponse](docs/DiscountLinesResponse.md)
 - [EmailCheckoutRequest](docs/EmailCheckoutRequest.md)
 - [EventResponse](docs/EventResponse.md)
 - [EventTypes](docs/EventTypes.md)
 - [EventsResendResponse](docs/EventsResendResponse.md)
 - [FiscalEntityAddress](docs/FiscalEntityAddress.md)
 - [GetApiKeysResponse](docs/GetApiKeysResponse.md)
 - [GetChargesResponse](docs/GetChargesResponse.md)
 - [GetCompaniesResponse](docs/GetCompaniesResponse.md)
 - [GetCustomerPaymentMethodDataResponse](docs/GetCustomerPaymentMethodDataResponse.md)
 - [GetEventsResponse](docs/GetEventsResponse.md)
 - [GetOrderDiscountLinesResponse](docs/GetOrderDiscountLinesResponse.md)
 - [GetOrdersResponse](docs/GetOrdersResponse.md)
 - [GetPaymentMethodResponse](docs/GetPaymentMethodResponse.md)
 - [GetPlansResponse](docs/GetPlansResponse.md)
 - [GetTransactionsResponse](docs/GetTransactionsResponse.md)
 - [GetTransfersResponse](docs/GetTransfersResponse.md)
 - [GetWebhookKeysResponse](docs/GetWebhookKeysResponse.md)
 - [GetWebhooksResponse](docs/GetWebhooksResponse.md)
 - [LogResponse](docs/LogResponse.md)
 - [LogsResponse](docs/LogsResponse.md)
 - [LogsResponseData](docs/LogsResponseData.md)
 - [ModelError](docs/ModelError.md)
 - [OrderCaptureRequest](docs/OrderCaptureRequest.md)
 - [OrderChannelResponse](docs/OrderChannelResponse.md)
 - [OrderChargesResponse](docs/OrderChargesResponse.md)
 - [OrderCustomerInfoResponse](docs/OrderCustomerInfoResponse.md)
 - [OrderDiscountLinesRequest](docs/OrderDiscountLinesRequest.md)
 - [OrderDiscountLinesResponse](docs/OrderDiscountLinesResponse.md)
 - [OrderFiscalEntityAddressResponse](docs/OrderFiscalEntityAddressResponse.md)
 - [OrderFiscalEntityRequest](docs/OrderFiscalEntityRequest.md)
 - [OrderFiscalEntityResponse](docs/OrderFiscalEntityResponse.md)
 - [OrderNextActionResponse](docs/OrderNextActionResponse.md)
 - [OrderNextActionResponseRedirectToUrl](docs/OrderNextActionResponseRedirectToUrl.md)
 - [OrderRefundRequest](docs/OrderRefundRequest.md)
 - [OrderRequest](docs/OrderRequest.md)
 - [OrderRequestCustomerInfo](docs/OrderRequestCustomerInfo.md)
 - [OrderResponse](docs/OrderResponse.md)
 - [OrderResponseCheckout](docs/OrderResponseCheckout.md)
 - [OrderResponseCustomerInfo](docs/OrderResponseCustomerInfo.md)
 - [OrderResponseProducts](docs/OrderResponseProducts.md)
 - [OrderResponseShippingContact](docs/OrderResponseShippingContact.md)
 - [OrderTaxRequest](docs/OrderTaxRequest.md)
 - [OrderUpdateFiscalEntityRequest](docs/OrderUpdateFiscalEntityRequest.md)
 - [OrderUpdateRequest](docs/OrderUpdateRequest.md)
 - [OrderUpdateRequestCustomerInfo](docs/OrderUpdateRequestCustomerInfo.md)
 - [OrdersResponse](docs/OrdersResponse.md)
 - [Page](docs/Page.md)
 - [Pagination](docs/Pagination.md)
 - [PaymentMethod](docs/PaymentMethod.md)
 - [PaymentMethodBankTransfer](docs/PaymentMethodBankTransfer.md)
 - [PaymentMethodBnplPayment](docs/PaymentMethodBnplPayment.md)
 - [PaymentMethodBnplRequest](docs/PaymentMethodBnplRequest.md)
 - [PaymentMethodCard](docs/PaymentMethodCard.md)
 - [PaymentMethodCardRequest](docs/PaymentMethodCardRequest.md)
 - [PaymentMethodCardResponse](docs/PaymentMethodCardResponse.md)
 - [PaymentMethodCash](docs/PaymentMethodCash.md)
 - [PaymentMethodCashRequest](docs/PaymentMethodCashRequest.md)
 - [PaymentMethodCashResponse](docs/PaymentMethodCashResponse.md)
 - [PaymentMethodCashResponseAllOfAgreements](docs/PaymentMethodCashResponseAllOfAgreements.md)
 - [PaymentMethodGeneralRequest](docs/PaymentMethodGeneralRequest.md)
 - [PaymentMethodResponse](docs/PaymentMethodResponse.md)
 - [PaymentMethodSpeiRecurrent](docs/PaymentMethodSpeiRecurrent.md)
 - [PaymentMethodSpeiRequest](docs/PaymentMethodSpeiRequest.md)
 - [PaymentMethodTokenRequest](docs/PaymentMethodTokenRequest.md)
 - [Payout](docs/Payout.md)
 - [PayoutMethod](docs/PayoutMethod.md)
 - [PayoutOrder](docs/PayoutOrder.md)
 - [PayoutOrderPayoutsItem](docs/PayoutOrderPayoutsItem.md)
 - [PayoutOrderResponse](docs/PayoutOrderResponse.md)
 - [PayoutOrderResponseCustomerInfo](docs/PayoutOrderResponseCustomerInfo.md)
 - [PayoutOrdersResponse](docs/PayoutOrdersResponse.md)
 - [PlanRequest](docs/PlanRequest.md)
 - [PlanResponse](docs/PlanResponse.md)
 - [PlanUpdateRequest](docs/PlanUpdateRequest.md)
 - [Product](docs/Product.md)
 - [ProductDataResponse](docs/ProductDataResponse.md)
 - [ProductOrderResponse](docs/ProductOrderResponse.md)
 - [ResendRequest](docs/ResendRequest.md)
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
 - [TransactionResponse](docs/TransactionResponse.md)
 - [TransferDestinationResponse](docs/TransferDestinationResponse.md)
 - [TransferMethodResponse](docs/TransferMethodResponse.md)
 - [TransferResponse](docs/TransferResponse.md)
 - [TransfersResponse](docs/TransfersResponse.md)
 - [UpdateCustomer](docs/UpdateCustomer.md)
 - [UpdateCustomerAntifraudInfo](docs/UpdateCustomerAntifraudInfo.md)
 - [UpdateCustomerFiscalEntitiesResponse](docs/UpdateCustomerFiscalEntitiesResponse.md)
 - [UpdateCustomerPaymentMethodsResponse](docs/UpdateCustomerPaymentMethodsResponse.md)
 - [UpdateOrderDiscountLinesRequest](docs/UpdateOrderDiscountLinesRequest.md)
 - [UpdateOrderTaxRequest](docs/UpdateOrderTaxRequest.md)
 - [UpdateOrderTaxResponse](docs/UpdateOrderTaxResponse.md)
 - [UpdatePaymentMethods](docs/UpdatePaymentMethods.md)
 - [UpdateProduct](docs/UpdateProduct.md)
 - [WebhookKeyCreateResponse](docs/WebhookKeyCreateResponse.md)
 - [WebhookKeyDeleteResponse](docs/WebhookKeyDeleteResponse.md)
 - [WebhookKeyRequest](docs/WebhookKeyRequest.md)
 - [WebhookKeyResponse](docs/WebhookKeyResponse.md)
 - [WebhookKeyUpdateRequest](docs/WebhookKeyUpdateRequest.md)
 - [WebhookLog](docs/WebhookLog.md)
 - [WebhookRequest](docs/WebhookRequest.md)
 - [WebhookResponse](docs/WebhookResponse.md)
 - [WebhookUpdateRequest](docs/WebhookUpdateRequest.md)
 - [WhitelistlistRuleResponse](docs/WhitelistlistRuleResponse.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### bearerAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


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

## Author

engineering@conekta.com
