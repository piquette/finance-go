# finance-go

## Introduction

Welcome to the greatest best new financial data api library implemented in go :sparkles:

This go package aims to provide a go application with access to financial markets data in streamlined, well-formatted structures.  The real benchmark for success will be method signatures that a programmer from any background can understand on sight- put parameters in, get an error-resistant data structure as a result.

Accomplishing this goal across several data sources (yfin, morningstar, FRED) etc, while also maximizing code structure flexibility for data source changes in the future is an undertaking that requires some consideration beforehand. So, this README will serve as a planned feature list and development roadmap until a v1 release is stable. Thanks for your patience!

### v1.0 features

Description | Source | Input | Output
--- | --- | --- | ---
Current quote | Yahoo finance | symbol, asset type  | Instrument quote
Multiple quotes | Yahoo finance | Slice of params (symbol, asset type)  | Slice of instrument quotes
Historical quotes | Yahoo finance | Symbol string, time interval, aggregation period | Slice of OHLC quotes
Option quotes | Yahoo finance | N/A | N/A
Symbols list | Bats | N/A | Slice of available equity symbols



### Eventual features




## Development

Pull requests from the community are welcome. If you submit one, please keep
the following guidelines in mind:

1. Code must be `go fmt` compliant.
2. All types, structs and funcs should be documented.
3. Ensure that `make test` succeeds.

## Test

The test suite needs testify's `require` package to run:

    github.com/stretchr/testify/require

It also depends on [finance-mock], so make sure to fetch and run it from a
background terminal ([finance-mock's README][finance-mock] also contains
instructions for installing via Homebrew and other methods):

    go get -u github.com/piquette/finance-mock
    finance-mock

Run all tests:

    go test ./...

Run tests for one package:

    go test ./quote

Run a single test:

    go test ./quote -run TestQuoteGet

For any requests, bug or comments, please [open an issue][issues] or [submit a
pull request][pulls].