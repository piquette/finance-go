package main

import (
	"fmt"

	"github.com/piquette/finance-go/options"
)

// This file lists several usage examples of this library
// and can be used to verify behavior.
func main() {

	iter := options.GetStraddle("TWTR")

	fmt.Println(iter.Meta())

	for iter.Next() {
		fmt.Println(iter.Straddle().Strike)
	}
	if iter.Err() != nil {
		fmt.Println(iter.Err())
	}

	// Basic quote example.
	// --------------------
	// q, err := quote.Get("SPY")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(q)
	// }

	// Basic chart example.
	// --------------------
	// params := &chart.Params{
	// 	Symbol:   "TWTR",
	// 	Interval: datetime.OneHour,
	// }
	// iter := chart.Get(params)
	//
	// for iter.Next() {
	// 	b := iter.Bar()
	// 	fmt.Println(datetime.FromUnix(b.Timestamp))
	//
	// }
	// if iter.Err() != nil {
	// 	fmt.Println(iter.Err())
	// }
}
