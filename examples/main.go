package main

import (
	"fmt"
	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/crypto"
	"github.com/piquette/finance-go/datetime"
	"github.com/piquette/finance-go/equity"
	"github.com/piquette/finance-go/etf"
	"github.com/piquette/finance-go/forex"
	"github.com/piquette/finance-go/mutualfund"
	"github.com/piquette/finance-go/options"
	"github.com/piquette/finance-go/quote"
)

// This file lists several usage examples of this library
// and can be used to verify behavior.
func main() {

	// Basic options example.
	// --------------------
	{
		fmt.Println("Options stradle example\n====================\n")
		iter := options.GetStraddle("AAPL")

		fmt.Println(iter.Meta())

		for iter.Next() {
			fmt.Println(iter.Straddle().Strike)
		}
		if iter.Err() != nil {
			fmt.Println(iter.Err())
		}
		fmt.Println()
	}

	// Basic quote example.
	// --------------------
	{
		fmt.Println("Quote example\n====================\n")
		q, err := quote.Get("GOOG")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(q)
		}
		fmt.Println()
	}

	// Basic chart example.
	// --------------------
	{
		fmt.Println("Chart example\n====================\n")
		params := &chart.Params{
			Symbol:   "GOOG",
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
		fmt.Println()
	}

	// Basic crypto example.
	// --------------------
	{
		fmt.Println("Crypto example\n====================\n")
		q, err := crypto.Get("BTC-USD")

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(q)
		}
		fmt.Println()
	}

	// Basic equity example.
	// --------------------
	{
		fmt.Println("Equity example\n====================\n")
		symbols := []string{"AAPL", "GOOG", "MSFT"}
		iter := equity.List(symbols)

		if iter.Err() != nil {
			fmt.Println(iter.Err())
			fmt.Println()
		} else {
			for iter.Next() {
				q := iter.Equity()
				fmt.Println(q)
				fmt.Println()
			}
		}
	}

	// Basic ETF example.
	// --------------------
	{
		fmt.Println("ETF example\n====================\n")
		q, err := etf.Get("SPY")

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(q)
		}
		fmt.Println()
	}

	// Basic forex example.
	// --------------------
	{
		fmt.Println("Forex example\n====================\n")
		q, err := forex.Get("CADUSD=X")

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(q)
		}
		fmt.Println()
	}

	// Basic future example.
	// --------------------
	{
		fmt.Println("Future example\n====================\n")
		q, err := forex.Get("CL=F")

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(q)
		}
		fmt.Println()
	}

	// Basic index example.
	// --------------------
	{
		fmt.Println("Index example\n====================\n")
		q, err := forex.Get("^DJI")

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(q)
		}
		fmt.Println()
	}

	// Basic mutual fund example.
	// --------------------
	{
		fmt.Println("Mutual fund example\n====================\n")
		q, err := mutualfund.Get("FMAGX")

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(q)
		}
		fmt.Println()
	}
}
