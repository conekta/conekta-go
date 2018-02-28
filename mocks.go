package conekta

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
	p.Country = "México"
	return p
}

// Mock fills PaymentSourceCreateParams with dummy data
func (p *PaymentSourceCreateParams) Mock() *PaymentSourceCreateParams {
	p.TokenID = "tok_test_visa_4242"
	p.PaymentType = "card"
	return p
}
