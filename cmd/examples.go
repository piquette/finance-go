package main

import (
	"fmt"

	"github.com/piquette/finance-go/history"
)

// This file lists several usage examples of this library
// and can be used to verify behavior.
func main() {

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
	if err := chart.Err(); err != nil {
		fmt.Println(err)
	}

}
