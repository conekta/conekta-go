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
	p.ExpiresAt = int(time.Now().AddDate(0, 0, 1).Unix())
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
	p.Description = "Linea de envio"
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
