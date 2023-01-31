package charge

import (
	"testing"

	"github.com/stretchr/testify/assert"

	conekta "github.com/conekta/conekta-go"
	"github.com/conekta/conekta-go/customer"
	"github.com/conekta/conekta-go/order"
	"github.com/conekta/conekta-go/paymentsource"
	"time"
)

func init() {
	conekta.APIKey = conekta.TestKey
}

func CreateOrder() (*conekta.Order, error) {
	op := &conekta.OrderParams{}
	ord, err := order.Create(op.Mock())
	return ord, err
}
func CreateOrderWithoutCharges() (*conekta.Order, error) {
	op := &conekta.OrderParams{}
	ord, err := order.Create(op.MockWithoutCharges())
	return ord, err
}

func TestCreate(t *testing.T) {
	cp := &conekta.ChargeParams{ReferenceID: "reference id"}
	or, _ := CreateOrderWithoutCharges()
	res, _ := Create(or.ID, cp.Mock())

	assert.Equal(t, false, res.Livemode)
	assert.NotEqual(t, nil, res.CreatedAt)
	assert.Equal(t, "MXN", res.Currency)
	assert.Equal(t, "charge", res.Object)
	assert.Equal(t, "paid", res.Status)
	assert.NotNil(t, res.PaidAt)
	assert.Equal(t, int64(50151), res.Amount)
	assert.NotEqual(t, nil, res.Description)
	assert.Equal(t, "", res.CustomerID)
	assert.NotEqual(t, nil, res.OrderID)
	assert.Equal(t, or.ID, res.OrderID)
	assert.Equal(t, cp.ReferenceID, res.ReferenceID)
}

func TestCreateWithPaymentSource(t *testing.T) {
	custp := &conekta.CustomerParams{}
	cust, _ := customer.Create(custp.Mock())

	psp := &conekta.PaymentSourceCreateParams{}
	ps, _ := paymentsource.Create(cust.ID, psp.Mock())

	cp := &conekta.ChargeParams{
		PaymentMethod: &conekta.PaymentMethodParams{
			MonthlyInstallments: 3,
			Type:                "card",
			PaymentSourceID:     ps.ID,
		},
		ReferenceID: "reference id",
	}

	op := &conekta.OrderParams{}
	op.MockWithoutCharges()
	op.Charges = append(op.Charges, cp)
	op.CustomerInfo = &conekta.CustomerParams{
		ID: cust.ID,
	}

	ord, err := order.Create(op)

	assert.Equal(t, "paid", ord.PaymentStatus)
	assert.Equal(t, cust.ID, ord.CustomerInfo.CustomerID)
	assert.Equal(t, int64(3), ord.Charges.Data[0].MonthlyInstallments)
	assert.NotEqual(t, nil, ord.CreatedAt)
	assert.Nil(t, err)
	assert.Equal(t, cp.ReferenceID, ord.Charges.Data[0].ReferenceID)
}

func TestCreateError(t *testing.T) {
	ord, _ := CreateOrderWithoutCharges()
	_, err := Create(ord.ID, &conekta.ChargeParams{})
	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "parameter_validation_error")
}

func TestProcessingError(t *testing.T) {
	ord, _ := CreateOrderWithoutCharges()
	chargeParams := (&conekta.ChargeParams{}).Mock()
	chargeParams.PaymentMethod.TokenID = "tok_test_card_declined"
	ch, err := Create(ord.ID, chargeParams)
	assert.Nil(t, err)
	assert.Equal(t, ch.Status, "declined")
}

func TestFind(t *testing.T) {
	ord, _ := CreateOrder()
	res, _ := Find(ord.ID, ord.Charges.Data[0].ID)

	assert.Equal(t, res.Description, "Payment from order")
	assert.Equal(t, int64(50151), res.Amount)
	assert.Equal(t, res.OrderID, ord.ID)
	assert.Equal(t, res.Object, "charge")
	assert.Equal(t, res.PaymentMethod.Last4, "4242")
	assert.Equal(t, res.PaymentMethod.Brand, "visa")
	assert.Equal(t, res.PaymentMethod.Name, "Jorge Lopez")
	assert.Equal(t, res.PaymentMethod.ExpMonth, time.Now().Format("01"))
	assert.Equal(t, res.PaymentMethod.ExpYear, time.Now().AddDate(1, 0, 0).Format("06"))
}

func TestFindError(t *testing.T) {
	ord, _ := CreateOrderWithoutCharges()
	_, err := Find(ord.ID, "123")
	assert.NotNil(t, err)
	assert.Equal(t, err.(conekta.Error).ErrorType, "resource_not_found_error")
}
