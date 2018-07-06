package main

import (
	"fmt"

	"github.com/piquette/finance-go/history"
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
	params := &history.Params{
		Symbol:   "TWTR",
		Interval: history.OneHour,
	}
	chart := history.Get(params)

	for chart.Next() {
		b := chart.Bar()
		fmt.Println(history.NewDatetimeU(b.Timestamp))

	}
	if chart.Err() != nil {
		fmt.Println(chart.Err())
	}

}
