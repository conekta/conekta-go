package order

import (
	"testing"

	conekta "github.com/conekta/conekta-go"
	"github.com/stretchr/testify/assert"
)

func init() {
	conekta.APIKey = conekta.TestKey
}

type Query struct {
	Currency string `url:"currency,omitempty"`
}

func TestCreate(t *testing.T) {
	op := &conekta.OrderParams{}
	ord, err := Create(op.Mock())

	assert.Equal(t, len(op.DiscountLines), ord.DiscountLines.Total)
	assert.Equal(t, op.DiscountLines[0].Amount, ord.DiscountLines.Data[0].Amount)
	assert.Equal(t, op.DiscountLines[0].Code, ord.DiscountLines.Data[0].Code)
	assert.Equal(t, op.DiscountLines[0].Type, ord.DiscountLines.Data[0].Type)
	assert.NotEqual(t, nil, ord.DiscountLines.Data[0].ID)

	assert.Equal(t, len(op.TaxLines), ord.TaxLines.Total)
	assert.Equal(t, int64(op.TaxLines[0].Amount), ord.TaxLines.Data[0].Amount)
	assert.Equal(t, op.TaxLines[0].Description, ord.TaxLines.Data[0].Description)
	assert.NotEqual(t, nil, ord.TaxLines.Data[0].ID)

	assert.Equal(t, len(op.LineItems), ord.LineItems.Total)
	assert.Equal(t, op.LineItems[0].Name, ord.LineItems.Data[0].Name)
	assert.Equal(t, op.LineItems[0].Quantity, ord.LineItems.Data[0].Quantity)
	assert.Equal(t, op.LineItems[0].UnitPrice, ord.LineItems.Data[0].UnitPrice)
	assert.NotEqual(t, nil, ord.LineItems.Data[0].ID)

	assert.Equal(t, op.ShippingContact.Phone, ord.ShippingContact.Phone)
	assert.Equal(t, op.ShippingContact.Receiver, ord.ShippingContact.Receiver)
	assert.Equal(t, op.ShippingContact.BetweenStreets, ord.ShippingContact.BetweenStreets)
	assert.NotEqual(t, nil, ord.ShippingContact.ID)

	//shipping address
	assert.Equal(t, op.ShippingContact.Address.Street1, ord.ShippingContact.Address.Street1)
	assert.Equal(t, op.ShippingContact.Address.Street2, ord.ShippingContact.Address.Street2)
	assert.Equal(t, op.ShippingContact.Address.PostalCode, ord.ShippingContact.Address.PostalCode)

	assert.Equal(t, op.CustomerInfo.Name, ord.CustomerInfo.Name)
	assert.Equal(t, op.CustomerInfo.Email, ord.CustomerInfo.Email)
	assert.Equal(t, op.CustomerInfo.Phone, ord.CustomerInfo.Phone)
	assert.NotEqual(t, nil, ord.CustomerInfo.ID)

	assert.Equal(t, op.Currency, ord.Currency)
	assert.NotEqual(t, nil, ord.ID)

	//charges
	assert.Equal(t, false, ord.Charges.Data[0].Livemode)
	assert.NotEqual(t, nil, ord.Charges.Data[0].CreatedAt)
	assert.Equal(t, "MXN", ord.Charges.Data[0].Currency)
	assert.Equal(t, "charge", ord.Charges.Data[0].Object)
	assert.Equal(t, int64(1151), ord.Charges.Data[0].Amount)
	assert.Equal(t, "paid", ord.Charges.Data[0].Status)
	assert.Equal(t, int64(328), ord.Charges.Data[0].Fee)
	assert.NotEqual(t, nil, ord.Charges.Data[0].Description)
	assert.Equal(t, "", ord.Charges.Data[0].CustomerID)
	assert.NotEqual(t, nil, ord.Charges.Data[0].OrderID)
	assert.NotEqual(t, nil, ord.Charges.Data[0].PaymentMethod.Name)
	assert.NotEqual(t, nil, ord.Charges.Data[0].PaymentMethod.ExpMonth)
	assert.NotEqual(t, nil, ord.Charges.Data[0].PaymentMethod.ExpYear)
	assert.Equal(t, "card_payment", ord.Charges.Data[0].PaymentMethod.Object)
	assert.Equal(t, "credit", ord.Charges.Data[0].PaymentMethod.Type)
	assert.Equal(t, "4242", ord.Charges.Data[0].PaymentMethod.Last4)
	assert.Equal(t, "visa", ord.Charges.Data[0].PaymentMethod.Brand)

	assert.Equal(t, "shipping_line", ord.ShippingLines.Data[0].Object)
	assert.Equal(t, "discount_line", ord.DiscountLines.Data[0].Object)
	assert.Equal(t, "tax_line", ord.TaxLines.Data[0].Object)
	assert.Equal(t, "line_item", ord.LineItems.Data[0].Object)
	assert.Equal(t, "shipping_contact", ord.ShippingContact.Object)
	assert.Equal(t, "shipping_address", ord.ShippingContact.Address.Object)
	assert.Equal(t, "customer_info", ord.CustomerInfo.Object)
	assert.Equal(t, "order", ord.Object)

	assert.Nil(t, err)
}

func TestOxxoCreate(t *testing.T) {
	op := &conekta.OrderParams{}
	ord, err := Create(op.OxxoMock())

	//root order
	assert.Equal(t, len(op.DiscountLines), ord.DiscountLines.Total)
	assert.Equal(t, op.DiscountLines[0].Amount, ord.DiscountLines.Data[0].Amount)
	assert.Equal(t, op.DiscountLines[0].Code, ord.DiscountLines.Data[0].Code)
	assert.Equal(t, op.DiscountLines[0].Type, ord.DiscountLines.Data[0].Type)
	assert.NotEqual(t, nil, ord.DiscountLines.Data[0].ID)
	//tax lines
	assert.Equal(t, len(op.TaxLines), ord.TaxLines.Total)
	assert.Equal(t, int64(op.TaxLines[0].Amount), ord.TaxLines.Data[0].Amount)
	assert.Equal(t, op.TaxLines[0].Description, ord.TaxLines.Data[0].Description)
	assert.NotEqual(t, nil, ord.TaxLines.Data[0].ID)

	//line items
	assert.Equal(t, len(op.LineItems), ord.LineItems.Total)
	assert.Equal(t, op.LineItems[0].Name, ord.LineItems.Data[0].Name)
	assert.Equal(t, op.LineItems[0].Quantity, ord.LineItems.Data[0].Quantity)
	assert.Equal(t, op.LineItems[0].UnitPrice, ord.LineItems.Data[0].UnitPrice)
	assert.NotEqual(t, nil, ord.LineItems.Data[0].ID)

	//shippping contact
	assert.Equal(t, op.ShippingContact.Phone, ord.ShippingContact.Phone)
	assert.Equal(t, op.ShippingContact.Receiver, ord.ShippingContact.Receiver)
	assert.Equal(t, op.ShippingContact.BetweenStreets, ord.ShippingContact.BetweenStreets)
	assert.NotEqual(t, nil, ord.ShippingContact.ID)

	//shipping address
	assert.Equal(t, op.ShippingContact.Address.Street1, ord.ShippingContact.Address.Street1)
	assert.Equal(t, op.ShippingContact.Address.Street2, ord.ShippingContact.Address.Street2)
	assert.Equal(t, op.ShippingContact.Address.PostalCode, ord.ShippingContact.Address.PostalCode)

	//customer info
	assert.Equal(t, op.CustomerInfo.Name, ord.CustomerInfo.Name)
	assert.Equal(t, op.CustomerInfo.Email, ord.CustomerInfo.Email)
	assert.Equal(t, op.CustomerInfo.Phone, ord.CustomerInfo.Phone)
	assert.NotEqual(t, nil, ord.CustomerInfo.ID)
	assert.Equal(t, op.Currency, ord.Currency)
	assert.NotEqual(t, nil, ord.ID)

	//charges
	assert.Equal(t, false, ord.Charges.Data[0].Livemode)
	assert.NotEqual(t, nil, ord.Charges.Data[0].CreatedAt)
	assert.Equal(t, "MXN", ord.Charges.Data[0].Currency)
	assert.Equal(t, "charge", ord.Charges.Data[0].Object)
	assert.Equal(t, "", ord.Charges.Data[0].CustomerID)
	assert.Equal(t, int64(1151), ord.Charges.Data[0].Amount)
	assert.Equal(t, "pending_payment", ord.Charges.Data[0].Status)
	assert.Equal(t, int64(46), ord.Charges.Data[0].Fee)
	assert.NotEqual(t, nil, ord.Charges.Data[0].Description)
	assert.Equal(t, "oxxo", ord.Charges.Data[0].PaymentMethod.Type)
	assert.NotEqual(t, nil, ord.Charges.Data[0].PaymentMethod.Reference)
	assert.NotEqual(t, nil, ord.Charges.Data[0].OrderID)

	//object type assertion
	assert.Equal(t, "shipping_line", ord.ShippingLines.Data[0].Object)
	assert.Equal(t, "discount_line", ord.DiscountLines.Data[0].Object)
	assert.Equal(t, "tax_line", ord.TaxLines.Data[0].Object)
	assert.Equal(t, "line_item", ord.LineItems.Data[0].Object)
	assert.Equal(t, "shipping_contact", ord.ShippingContact.Object)
	assert.Equal(t, "shipping_address", ord.ShippingContact.Address.Object)
	assert.Equal(t, "customer_info", ord.CustomerInfo.Object)
	assert.Equal(t, "order", ord.Object)
	assert.NotEqual(t, nil, ord.Charges.Data[0].ID)
	//This should be empty when passing customer on the go
	assert.Equal(t, "", ord.Charges.Data[0].CustomerID)
	assert.Nil(t, err)
}

func TestUpdate(t *testing.T) {
	op := &conekta.OrderParams{}
	ord, _ := Create(op.MockWithoutCharges())
	assert.Equal(t, "MXN", ord.Currency)
	op.Currency = "USD"
	res, err := Update(ord.ID, op)
	assert.Nil(t, err)
	assert.Equal(t, "USD", res.Currency)
}

func TestCapture(t *testing.T) {
	op := &conekta.OrderParams{}
	op.PreAuth = true
	ord, _ := Create(op.Mock())
	assert.NotEqual(t, op.PreAuth, ord.PreAuth)
	assert.Equal(t, "pre_authorized", ord.PaymentStatus)
	res, err := Capture(ord.ID)
	assert.Equal(t, false, res.PreAuth)
	assert.Equal(t, "paid", res.PaymentStatus)
	assert.Nil(t, err)
}

func TestWhere(t *testing.T) {
	op := &Query{Currency: "MXN"}

	res, _ := Where(op)

	assert.NotNil(t, res.Total)
	assert.True(t, res.HasMore)
	assert.Equal(t, "list", res.Object)

}

func TestFind(t *testing.T) {
	op := &conekta.OrderParams{}
	ord, _ := Create(op.Mock())
	res, err := Find(ord.ID)
	assert.Equal(t, ord.ID, res.ID)
	assert.Nil(t, err)
}

func TestCancel(t *testing.T) {
	op := &conekta.OrderParams{}
	ord, _ := Create(op.OxxoMock())
	res, err := Cancel(ord.ID)

	assert.NotNil(t, res.ID)
	assert.Equal(t, ord.ID, res.ID)
	assert.Equal(t, "canceled", res.PaymentStatus)
	assert.Equal(t, "canceled", res.Charges.Data[0].Status)
	assert.Nil(t, err)

}

func TestRefund(t *testing.T) {
	op := &conekta.OrderParams{}
	ord, _ := Create(op.Mock())

	rp := &conekta.OrderRefundParams{}
	res, err := Refund(ord.ID, rp.Mock())

	assert.Equal(t, "refunded", res.PaymentStatus)
	assert.Nil(t, err)
}
