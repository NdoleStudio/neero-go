# neero-go

[![Build](https://github.com/NdoleStudio/neero-go/actions/workflows/main.yml/badge.svg)](https://github.com/NdoleStudio/neero-go/actions/workflows/main.yml)
[![codecov](https://codecov.io/gh/NdoleStudio/neero-go/branch/main/graph/badge.svg)](https://codecov.io/gh/NdoleStudio/neero-go)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/NdoleStudio/neero-go/badges/quality-score.png?b=main)](https://scrutinizer-ci.com/g/NdoleStudio/neero-go/?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/NdoleStudio/neero-go)](https://goreportcard.com/report/github.com/NdoleStudio/neero-go)
[![GitHub contributors](https://img.shields.io/github/contributors/NdoleStudio/neero-go)](https://github.com/NdoleStudio/neero-go/graphs/contributors)
[![GitHub license](https://img.shields.io/github/license/NdoleStudio/neero-go?color=brightgreen)](https://github.com/NdoleStudio/neero-go/blob/master/LICENSE)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/NdoleStudio/neero-go)](https://pkg.go.dev/github.com/NdoleStudio/neero-go)


This package provides a go API client for the Neero Payment Gateway API

## Installation

`neero-go` is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/NdoleStudio/neero-go
```

Alternatively the same can be achieved if you use `import` in a package:

```go
import "github.com/NdoleStudio/neero-go"
```


## Implemented

- Payment Methods
  - `POST /api/v1/payment-methods`: Create Payment Method
  - `POST /api/v1/payment-methods/resolve-details`: Resolve Payment Method Details
- Balances
  - `GET /api/v1/balances/payment-method/{paymentMethodId}`: Get Balance for a Payment Method
- Transaction Intents
  - `POST /api/v1/transaction-intents/cash-in`: Create Cash In Payment Intent
  - `POST /api/v1/transaction-intents/cash-out`: Create Cash Out Payment Intent
  - `GET /api/v1/transaction-intents/{transactionId}`: Find Transaction Intent By Id

## Usage

### Initializing the Client

An instance of the client can be created using `New()`.

```go
package main

import (
	"github.com/NdoleStudio/neero-go"
)

func main()  {
	client := neero.New(neero.WithSecretKey("" /*Your Neero merchant account secret key*/))
}
```

### Error handling

All API calls return an `error` as the last return object. All successful calls will return a `nil` error.

```go
status, response, err := statusClient.Status.Ok(context.Background())
if err != nil {
    //handle error
}
```

## Testing

You can run the unit tests for this client from the root directory using the command below:

```bash
go test -v
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
