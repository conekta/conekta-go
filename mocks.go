package conekta

// Mock fills CustomerParams with dummy data
func (c *CustomerParams) Mock() *CustomerParams {
	c.Name = "René Daniel"
	c.Phone = "+525545345654"
	c.Email = "test@test.com"
	c.ShippingContacts = []*ShippingContactParams{
		&ShippingContactParams{
			Phone:          "5565455543",
			Receiver:       "René",
			BetweenStreets: "Calle viga y morelos",
			Address: &Address{
				Street1:    "Nuevo leon",
				City:       "CdMx",
				PostalCode: "01900",
				Country:    "México",
			},
		},
	}
	return c
}
