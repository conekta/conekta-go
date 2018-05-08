package conekta

import "time"

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

// Mock fills LineItemsParams with dummy data
func (li *LineItemsParams) Mock() *LineItemsParams {
	li.Name = "HotDog"
	li.UnitPrice = 1000
	li.Quantity = 1
	return li
}

// Mock fills ChargeParams with dummy data
func (p *ChargeParams) Mock() *ChargeParams {
	pm := &PaymentMethodParams{}

	p.PaymentMethodParams = pm.Mock()
	return p
}

// OxxoMock fills ChargeParams with dummy data
func (p *ChargeParams) OxxoMock() *ChargeParams {
	pm := &PaymentMethodParams{}

	p.PaymentMethodParams = pm.OxxoMock()
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
	p.ExpMonth = "12"
	p.ExpYear = "2020"
	p.Cvc = "123"
	return p
}

// Mock fills WebhookParams with dummy data
func (p *WebhookParams) Mock() *WebhookParams {
	p.URL = "c.testgowebhook.com"
	p.DevelopmentEnabled = true
	p.ProductionEnabled = false
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
