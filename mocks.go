package conekta

import (
	"math/rand"
	"time"
)

// Mock fills CustomerParams with dummy data
func (p *CustomerParams) Mock() *CustomerParams {
	scp := &ShippingContactParams{}
	p.Name = "René Daniel"
	p.Phone = "+525545345654"
	p.Email = "test@test.com"
	p.ShippingContacts = append(p.ShippingContacts, scp.Mock())
	return p
}

// Mock fills ShippingContactParams with dummy data
func (p *ShippingContactParams) Mock() *ShippingContactParams {
	a := &Address{}
	p.Phone = "5565455543"
	p.Receiver = "René"
	p.BetweenStreets = "Calle viga y morelos"
	p.Address = a.Mock()
	return p
}

// Mock fills Address with dummy data
func (p *Address) Mock() *Address {
	p.Street1 = "Nuevo leon"
	p.City = "CdMx"
	p.PostalCode = "01900"
	p.Country = "Mexico"
	return p
}

// Mock fills PaymentSourceCreateParams with dummy data
func (p *PaymentSourceCreateParams) Mock() *PaymentSourceCreateParams {
	p.TokenID = "tok_test_visa_4242"
	p.PaymentType = "card"
	return p
}

// MockMinimal fills CustomerParams with dummy data
func (p *CustomerParams) MockMinimal() *CustomerParams {
	p.Name = "René Daniel"
	p.Phone = "+525545345654"
	p.Email = "test@test.com"
	return p
}

func (p *CustomerParams) MockCustomerPaymentSource() *CustomerParams {
	pay := &PaymentSourceCreateParams{}
	p.Name = "René Daniel"
	p.Phone = "+525545345654"
	p.Email = "test@test.com"
	p.PaymentSources = append(p.PaymentSources, pay.Mock())
	return p
}

// Mock fills LineItemsParams with dummy data
func (li *LineItemsParams) Mock() *LineItemsParams {
	li.Name = "HotDog"
	li.UnitPrice = 50000
	li.Quantity = 1
	return li
}

// Mock fills ChargeParams with dummy data
func (p *ChargeParams) Mock() *ChargeParams {
	pm := &PaymentMethodParams{}

	p.PaymentMethod = pm.Mock()
	return p
}

// OxxoMock fills ChargeParams with dummy data
func (p *ChargeParams) OxxoMock() *ChargeParams {
	pm := &PaymentMethodParams{}

	p.PaymentMethod = pm.OxxoMock()
	return p
}

// Mock fills PaymentMethodParams with dummy data
func (p *PaymentMethodParams) Mock() *PaymentMethodParams {
	p.Type = "card"
	p.TokenID = "tok_test_visa_4242"
	return p
}

// OxxoMock fills PaymentMethodParams with dummy data
func (p *PaymentMethodParams) OxxoMock() *PaymentMethodParams {
	p.Type = "oxxo_cash"
	p.ExpiresAt = time.Now().AddDate(0, 0, 1).Unix()
	return p
}

// Mock fills TaxLinesParams with dummy data
func (p *TaxLinesParams) Mock() *TaxLinesParams {
	p.Description = "IVA"
	p.Amount = 160
	return p
}

// Mock fills Address with dummy data
func (p *DiscountLinesParams) Mock() *DiscountLinesParams {
	p.Code = "Cupon de descuento"
	p.Amount = 10
	p.Type = "loyalty"
	return p
}

// Mock fills Address with dummy data
func (p *ShippingLinesParams) Mock() *ShippingLinesParams {
	p.Amount = 1
	p.Carrier = "Carrier"
	p.Method = "Fedex"
	p.TrackingNumber = "123000000001"
	return p
}

func (p *CheckoutParams) Mock() *CheckoutParams {
	p.Name = "Payment Link Name"
	p.Type = "PaymentLink"
	p.Recurrent = false
	p.ExpiredAt = time.Now().Unix() + int64(259200) + int64(rand.Float64()*3600)
	p.AllowedPaymentMethods = []string{"cash", "card", "bank_transfer"}
	p.NeedsShippingContact = true
	p.MonthlyInstallmentsEnabled = false
	p.MonthlyInstallmentsOptions = []int64{3, 6, 9, 12}
	p.OrderTemplate = &OrderParams{}
	p.OrderTemplate.Currency = "MXN"
	p.OrderTemplate.LineItems = []*LineItemsParams{
		{
			Name:      "Red Wine",
			UnitPrice: 1000,
			Quantity:  10,
		},
	}
	p.OrderTemplate.CustomerInfo = &CustomerParams{
		Name:  "Juan Perez",
		Email: "test@conekta.com",
		Phone: "5566982090",
	}

	return p
}

// Mock fills OrderParams with dummy data
func (p *OrderParams) Mock() *OrderParams {
	cp := &CustomerParams{}
	li := &LineItemsParams{}
	ch := &ChargeParams{}
	tl := &TaxLinesParams{}
	sc := &ShippingContactParams{}
	dl := &DiscountLinesParams{}
	sl := &ShippingLinesParams{}

	p.Currency = "MXN"
	p.CustomerInfo = cp.MockMinimal()
	p.LineItems = append(p.LineItems, li.Mock())
	p.Charges = append(p.Charges, ch.Mock())
	p.TaxLines = append(p.TaxLines, tl.Mock())
	p.ShippingContact = sc.Mock()
	p.DiscountLines = append(p.DiscountLines, dl.Mock())
	p.ShippingLines = append(p.ShippingLines, sl.Mock())
	return p
}

// Mock fills OrderParams with dummy data
func (p *OrderParams) MockWithCheckout(customerID string) *OrderParams {
	p.Currency = "MXN"
	p.CustomerInfo = &CustomerParams{
		ID: customerID,
	}
	p.LineItems = []*LineItemsParams{
		{
			Name:      "Box of Cohiba S1s",
			UnitPrice: 35000,
			Quantity:  1,
		},
	}
	p.Checkout = &OrderCheckoutParams{
		ExpiresAt:             time.Now().Unix() + int64(259200) + int64(rand.Float64()*3600),
		AllowedPaymentMethods: []string{"cash", "card", "bank_transfer"},
	}
	return p
}

// MockWithoutCharges fills OrderParams with dummy data
func (p *OrderParams) MockWithoutCharges() *OrderParams {
	cp := &CustomerParams{}
	li := &LineItemsParams{}
	tl := &TaxLinesParams{}
	sc := &ShippingContactParams{}
	dl := &DiscountLinesParams{}
	sl := &ShippingLinesParams{}

	p.Currency = "MXN"
	p.CustomerInfo = cp.MockMinimal()
	p.LineItems = append(p.LineItems, li.Mock())
	p.TaxLines = append(p.TaxLines, tl.Mock())
	p.ShippingContact = sc.Mock()
	p.DiscountLines = append(p.DiscountLines, dl.Mock())
	p.ShippingLines = append(p.ShippingLines, sl.Mock())
	return p
}

// OxxoMock fills OrderParams with dummy data
func (p *OrderParams) OxxoMock() *OrderParams {
	cp := &CustomerParams{}
	li := &LineItemsParams{}
	ch := &ChargeParams{}
	tl := &TaxLinesParams{}
	sc := &ShippingContactParams{}
	dl := &DiscountLinesParams{}
	sl := &ShippingLinesParams{}

	p.Currency = "MXN"
	p.CustomerInfo = cp.MockMinimal()
	p.LineItems = append(p.LineItems, li.Mock())
	p.Charges = append(p.Charges, ch.OxxoMock())
	p.TaxLines = append(p.TaxLines, tl.Mock())
	p.ShippingContact = sc.Mock()
	p.DiscountLines = append(p.DiscountLines, dl.Mock())
	p.ShippingLines = append(p.ShippingLines, sl.Mock())
	return p
}

// MockWithoutDiscountLines fills OrderParams with dummy data
func (p *OrderParams) MockWithoutDiscountLines() *OrderParams {
	cp := &CustomerParams{}
	li := &LineItemsParams{}
	tl := &TaxLinesParams{}
	sc := &ShippingContactParams{}
	sl := &ShippingLinesParams{}

	p.Currency = "MXN"
	p.CustomerInfo = cp.MockMinimal()
	p.LineItems = append(p.LineItems, li.Mock())
	p.TaxLines = append(p.TaxLines, tl.Mock())
	p.ShippingContact = sc.Mock()
	p.ShippingLines = append(p.ShippingLines, sl.Mock())
	return p
}

// Mock Fills token api response
func (p *TokenParams) Mock() *TokenParams {
	cp := &CardParams{}
	p.Card = cp.Mock()
	return p
}

// Mock fills TokenParams with dummy data
func (p *CardParams) Mock() *CardParams {
	p.Number = "4242424242424242"
	p.Name = "Eduardo Enriquez"
	p.ExpMonth = time.Now().Format("01")
	p.ExpYear = time.Now().AddDate(1, 0, 0).Format("2006")
	p.Cvc = "123"
	return p
}

// Mock fills WebhookParams with dummy data
func (p *WebhookParams) Mock() *WebhookParams {
	p.URL = "https://c.testgowebhook.com"
	return p
}

// Mock fills PayeeParams
func (p *PayeeParams) Mock() *PayeeParams {
	pba := &PayeeBillingAddressParams{}
	p.Name = "Eduardo Enriquez"
	p.Email = "interfaces@conekta.com"
	p.Phone = "+5255555555"
	p.BillingAddress = pba.Mock()
	return p
}

// Mock fills PayeeBillingAdressParams
func (p *PayeeBillingAddressParams) Mock() *PayeeBillingAddressParams {
	p.CompanyName = "Bandai"
	p.TaxID = "tax123"
	p.Street1 = "calle 1"
	p.Street2 = "calle 2"
	p.Street3 = "calle 3"
	p.City = "Cuauhtemoc"
	p.State = "DF"
	p.Country = "MX"
	p.PostalCode = "06100"
	return p
}

// Mock fills DestinationParams
func (p *DestinationParams) Mock() *DestinationParams {
	p.Type = "bank_account"
	p.AccountNumber = "072225008217746674"
	p.AccountHolderName = "J D - Radcorp"
	return p
}

// Mock fills TransferParams
func (p *TransferParams) Mock() *TransferParams {
	p.Amount = 5000
	p.Currency = "MXN"
	return p
}

// Mock fills OrderRefundParams
func (p *OrderRefundParams) Mock() *OrderRefundParams {
	p.Reason = OrderRefundRequestedByClient
	return p
}

// Mock fills PlanParams
func (p *PlanParams) MockPlanCreate() *PlanParams {
	p.Name = "bronze"
	p.Amount = 10000
	p.Interval = "month"
	p.Frequency = 1
	p.TrialPeriodDays = 15
	p.ExpiryCount = 6
	return p
}

func (p *PlanParams) MockPlan() *PlanParams {
	p.ID = "bronze"
	p.Name = "bronze"
	p.Amount = 10000
	p.Interval = "month"
	p.Frequency = 1
	p.TrialPeriodDays = 15
	p.ExpiryCount = 6
	return p
}

//Mock fills Update PlanParams
func (p *PlanParams) MockPlanUpdate() *PlanParams {
	p.ID = "bronze"
	p.Amount = 9900
	return p
}

//Mock fills SubscriptionParams
/*
func (p *SubscriptionParams) MockSubCreate() *SubscriptionParams {
	cu := &CustomerParams{}
	pl := &PlanParams{}

	pla := pl.MockPlan()
	c := &customer.Create(cu.MockCustomerPaymentSource())



	d, _ := json.Marshal(c)
	e := string(d)
	fmt.Println("--Es el customer de mocks: ",e)

	//p.CustomerID = c.ID
	p.PlanID = pla.ID
	p.CardID = c.DefaultPaymentSourceID
	return p
}
*/
