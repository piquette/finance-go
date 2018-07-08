package main

import (
	"fmt"

	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/quote"
)

// This file lists several usage examples of this library
// and can be used to verify behavior.
func main() {

	// Basic quote example.
	// --------------------
	q, err := quote.Get("SPY")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(q)
	}

	// Basic chart example.
	// --------------------
	params := &chart.Params{
		Symbol:   "TWTR",
		Interval: chart.OneHour,
	}
	iter := chart.Get(params)

	for iter.Next() {
		b := iter.Bar()
		fmt.Println(chart.NewDatetimeU(b.Timestamp))

	}
	if iter.Err() != nil {
		fmt.Println(iter.Err())
	}
}
