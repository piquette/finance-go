package main

import (
	"fmt"

	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
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
		Interval: datetime.OneHour,
	}
	iter := chart.Get(params)

	for iter.Next() {
		b := iter.Bar()
		fmt.Println(datetime.FromUnix(b.Timestamp))

	}
	if iter.Err() != nil {
		fmt.Println(iter.Err())
	}
}
