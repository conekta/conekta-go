<div align="center">
  
![banner](readme_files/banner.png)

# Conekta GO

</div>

## Requeriments

conekta-go is tested against against all officially supported releases (Go 1.12 and 1.13).

## Installation

You can import conekta-go directly from github as follows:

```go
import (
  conekta "github.com/conekta/conekta-go/v2"
  "github.com/conekta/conekta-go/v2/order"
  "github.com/conekta/conekta-go/v2/customer"
  "time"
)
```

## Usage

```go
conekta.APIKey = "key_ZLy4aP2szht1HqzkCezDEA"

customerParams := &conekta.CustomerParams{
	Name:  "Cándida Eréndira",
	Email: "la.candida.erendira@gmail.com",
	Phone: "55-5555-5555",
}

lineItemParams := &conekta.LineItemsParams{
	Name:      "Naranjas Robadas",
	UnitPrice: 10000,
	Quantity:  2,
}

chargeParams := &conekta.ChargeParams{
	PaymentMethodParams: &conekta.PaymentMethodParams{
		Type:      "oxxo_cash",
		ExpiresAt: time.Now().AddDate(0, 0, 90).Unix(),
	},
}

orderParams := &conekta.OrderParams{}
orderParams.Currency = "MXN"
orderParams.CustomerInfo = customerParams
orderParams.LineItems = append(orderParams.LineItems, lineItemParams)
orderParams.Charges = append(orderParams.Charges, chargeParams)

ord, err := order.Create(orderParams)
if err != nil {
	code := err.(conekta.Error).Details[0].Code
	//do something
	fmt.Println(code)
} else {
	orderId := ord.ID
	chargeId := ord.Charges.Data[0].ID
	oxxoReference := ord.Charges.Data[0].PaymentMethod.Reference
	//do something
	fmt.Println(orderId)
	fmt.Println(chargeId)
	fmt.Println(oxxoReference)
}
```

## Run Tests

```bash
go test -v ./...
```

## Changelog

Please see [CHANGELOG](CHANGELOG.md) for more information what has changed recently.

## Contributing

1. Fork the repository
2. Clone the repository
```bash
    git clone git@github.com:yourUserName/conekta-go/v2.git
```
3. Create a branch
```bash
    git checkout develop
    git pull origin develop
    # You should choose the name of your branch
    git checkout -b <feature/my_branch>
```
4. Make necessary changes and commit those changes
```bash
    git add .
    git commit -m "my changes"
```
5. Push changes to GitHub
```bash
    git push origin <feature/my_branch>
```
6. Submit your changes for review, create a pull request

   To create a pull request, you need to have made your code changes on a separate branch. This branch should be named like this: **feature/my_feature** or **fix/my_fix**.

   Make sure that, if you add new features to our library, be sure to add the corresponding **unit tests**.

   If you go to your repository on GitHub, you’ll see a Compare & pull request button. Click on that button.

***

## License

Developed in Mexico by [Conekta](https://www.conekta.com). Available with [MIT License](LICENSE).

## We are always hiring!

If you are a comfortable working with a range of backend languages (Java, Python, Ruby, PHP, etc) and frameworks, you have solid foundation in data structures, algorithms and software design with strong analytical and debugging skills. Send us your CV and GitHub to quieroser@conekta.com
