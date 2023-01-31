## [2.1.1](https://github.com/conekta/conekta-go/releases/tag/v2.1.1) - 2023-01-31
## What's Changed
* Feature/charge add reference id by @shamigor in https://github.com/conekta/conekta-go/pull/49
* Bump github.com/stretchr/testify from 1.7.0 to 1.8.1 by @dependabot in https://github.com/conekta/conekta-go/pull/54
* Bump github.com/google/go-querystring from 1.0.0 to 1.1.0 by @dependabot in https://github.com/conekta/conekta-go/pull/55
* add X-Conekta-Client-User-Agent header by @fcarrero in https://github.com/conekta/conekta-go/pull/56

## New Contributors
* @shamigor made their first contribution in https://github.com/conekta/conekta-go/pull/49
* @dependabot made their first contribution in https://github.com/conekta/conekta-go/pull/54
* @fcarrero made their first contribution in https://github.com/conekta/conekta-go/pull/56

**Full Changelog**: https://github.com/conekta/conekta-go/compare/v1.1.1...v2.1.1

## [2.1.0](https://github.com/conekta/conekta-go/releases/tag/v2.1.0)
### Feature
- This release adds new endpoints and parameters for the Checkout object and bumps to go 1.15

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
