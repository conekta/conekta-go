## [1.1.1](https://github.com/conekta/conekta-go/releases/tag/v1.1.1) - 2022-04-08
### bugfix
- fix for checkout tests
- project version reverted for golang issues with v2 or higher

## [2.0.0](https://github.com/conekta/conekta-go/releases/tag/v2.0.0) - 2020-03-04
### Feature
- Standardizes `PaymentMethod` as a child of `ChargeParams` to avoid nomenclature inconsistencies in the rest of the project
- Corrects Broken MSI specs
- Adds Subscription/Plan Support
- 402 ProcessingError now returns the modified object rather than throw an error, as a user you need to check the return status to determine whether or not it was successful

## [1.0.0](https://github.com/conekta/conekta-go/releases/tag/v1.0.0) - 2020-01-20
### Feature
- Initial Release of conekta-go
