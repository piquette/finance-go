# finance-go

[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/piquette/finance-go) [![Build Status](https://travis-ci.org/piquette/finance-go.svg?branch=master)](https://travis-ci.org/piquette/finance-go) [![Coverage Status](https://coveralls.io/repos/github/piquette/finance-go/badge.svg?branch=master)](https://coveralls.io/github/piquette/finance-go?branch=master)

## Summary

Welcome to the greatest best new financial data api library implemented in go :sparkles:

Not production ready! This go package aims to provide a go application with access to financial markets data in streamlined, well-formatted structures.  The real benchmark for success will be method signatures that a programmer from any background can understand on sight- put parameters in, get an error-resistant data structure as a result.

Accomplishing this goal across several data sources (yfin, morningstar, FRED) etc, while also maximizing code structure flexibility for data source changes in the future is an undertaking that requires some consideration beforehand. So, this README will serve as a planned feature list and development roadmap until a v1 release is stable. Thanks for your patience!

### Planned v1.0 features

Replication of the current features of FlashBoys/go-finance.

Status | Description | Source
--- | --- | ---
[x] | Equity quote(s) | Yahoo finance
[x] | Index quote(s) | Yahoo finance
[x] | Option quote(s) | Yahoo finance
[x] | Forex pair quote(s) | Yahoo finance
[x] | Futures quote(s) | Yahoo finance
[x] | ETF quote(s) | Yahoo finance
[x] | Mutual fund quote(s) | Yahoo finance
[x] | Historical quotes | Yahoo finance
[ ] | Options chains | Yahoo finance
[ ] | Symbols list | BATS

## Planned v1.0 documentation

A neatly formatted detailed list of implementation instructions and examples will be coming to the [piquette website][api-docs].

For details on all the functionality in this library, see the [GoDoc][godoc] documentation.

## Installation

```sh
go get github.com/piquette/finance-go
```

## Usage examples

Library usage is meant to be very specific about the user's intentions.

### Quote
```go
quote, err := equity.Get("AAPL")
if err != nil {
  // Uh-oh.  
  panic(err)
}

// Success!
fmt.Println(quote)
```

### Historical data
```go
params := &history.Params{
  Symbol:   "TWTR",
  Interval: history.OneHour,
}
chart := history.Get(params)

for chart.Next() {
  fmt.Println(chart.Bar())
}
if err := chart.Err(); err != nil {
  fmt.Println(err)
}
```

## Development

Pull requests from the community are welcome. If you submit one, please keep
the following guidelines in mind:

1. All types, structs and funcs should be documented.
2. Ensure that `make test` succeeds.

## Test

The test suite needs testify's `require` package to run:

    github.com/stretchr/testify/require

It also depends on a running instance of a test server [finance-mock], so make sure to fetch that project and run the application from another terminal session ([finance-mock's README][finance-mock] contains more information).

### Docker
```sh
  docker run -p 12111:12111 piquette/finance-mock:latest
```
### Brew

    brew tap piquette/finance-mock
    brew install finance-mock
    finance-mock

### Go

    go get -u github.com/piquette/finance-mock
    finance-mock

Run all tests:

    go test ./...

Run tests for one package:

    go test ./equity

Run a single test:

    go test ./equity -run TestGet

For any requests, bug or comments, please [open an issue][issues] or [submit a
pull request][pulls]. Also please email or tweet me as needed.

## Notes
- Yahoo changes their finance APIs without warning, which is their right to do so. However, its annoying and leads to some instability in this project..
- Big shoutout to Stripe and the team working on the [stripe-go][stripe] project, I took a lot of library design / implementation hints from them.

[godoc]: http://godoc.org/github.com/piquette/finance-go
[issues]: https://github.com/piquette/finance-go/issues/new
[pulls]: https://github.com/piquette/finance-go/pulls
[finance-mock]: https://github.com/piquette/finance-mock
[stripe]: https://github.com/stripe/stripe-go
[api-docs]: https://piquette.io/projects/finance-go/
