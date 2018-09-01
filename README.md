# finance-go

[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/piquette/finance-go) [![Build Status](https://travis-ci.org/piquette/finance-go.svg?branch=master)](https://travis-ci.org/piquette/finance-go) [![Coverage Status](https://coveralls.io/repos/github/piquette/finance-go/badge.svg?branch=master)](https://coveralls.io/github/piquette/finance-go?branch=master)

## Summary

This go package aims to provide a go application with access to current and historical financial markets data in streamlined, well-formatted structures.

Check out the [qtrn cli application][qtrn], which is intended as a living example of this package. It prints quotes/options info in your favorite command-line in a few keystrokes!

### Features

Description | Source
--- | ---
Quote(s) | Yahoo finance
Equity quote(s) | Yahoo finance
Index quote(s) | Yahoo finance
Option quote(s) | Yahoo finance
Forex pair quote(s) | Yahoo finance
Cryptocurrency pair quote(s) | Yahoo finance
Futures quote(s) | Yahoo finance
ETF quote(s) | Yahoo finance
Mutual fund quote(s) | Yahoo finance
Historical quotes | Yahoo finance
Options straddles | Yahoo finance

## Documentation

A neatly formatted detailed list of implementation instructions and examples will be available on the [piquette website][api-docs].

For now, for details on all the functionality in this library, see the [GoDoc][godoc] documentation.

## Installation

It is best to use a dependency management tool, but if you want to retrieve it manually, use -

```sh
go get github.com/piquette/finance-go
```

## Usage example

Library usage is meant to be very specific about the user's intentions.

### Quote
```go
q, err := quote.Get("AAPL")
if err != nil {
  // Uh-oh.  
  panic(err)
}

// Success!
fmt.Println(q)
```

### Equity quote (more fields)
```go
q, err := equity.Get("AAPL")
if err != nil {
  // Uh-oh.  
  panic(err)
}

// Success!
fmt.Println(q)
```

### Historical quotes (OHLCV)
```go
params := &chart.Params{
  Symbol:   "TWTR",
  Interval: datetime.OneHour,
}
iter := chart.Get(params)

for iter.Next() {
  fmt.Println(iter.Bar())
}
if err := iter.Err(); err != nil {
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
[qtrn]: https://github.com/piquette/qtrn
[pulls]: https://github.com/piquette/finance-go/pulls
[finance-mock]: https://github.com/piquette/finance-mock
[stripe]: https://github.com/stripe/stripe-go
[api-docs]: https://piquette.io/projects/finance-go/
