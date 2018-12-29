package main

import (

	chart "github.com/piquette/finance-go/chart"
	finance "github.com/piquette/finance-go"
	"github.com/piquette/finance-go/datetime"
	"fmt"
)



func GetHistoricalQuote(symbol string, month int, day int, year int) (*finance.ChartBar, error) {
	p := &chart.Params{
		Symbol:   symbol,
		Start:    &datetime.Datetime{Month: month, Day: day, Year: year},
		End:      &datetime.Datetime{Month: month, Day: day, Year: year},
		Interval: datetime.OneDay,
	  }
	iter := chart.Get(p)
	for iter.Next() {
		b := iter.Bar()
		return b, nil
	}

	return nil, nil
}

func main() {
	q, _ := GetHistoricalQuote("AAPL", 1, 11, 2018)
	fmt.Println(q)
}