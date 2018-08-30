package chart

import (
	"context"

	finance "github.com/piquette/finance-go"
	"github.com/piquette/finance-go/datetime"
	form "github.com/piquette/finance-go/form"
	"github.com/piquette/finance-go/iter"
	"github.com/shopspring/decimal"
)

// Client is used to invoke chart APIs.
type Client struct {
	B finance.Backend
}

func getC() Client {
	return Client{finance.GetBackend(finance.YFinBackend)}
}

// Params carries a context and chart information.
type Params struct {
	// Context access.
	finance.Params `form:"-"`

	// Accessible fields.
	Symbol   string             `form:"-"`
	Start    *datetime.Datetime `form:"-"`
	End      *datetime.Datetime `form:"-"`
	Interval datetime.Interval  `form:"-"`

	IncludeExt bool `form:"includePrePost"`

	// Internal request fields.
	interval string `form:"interval"`
	start    int    `form:"period1"`
	end      int    `form:"period2"`
}

// Iter is a structure containing results
// and related metadata for a
// yfin chart request.
type Iter struct {
	*iter.Iter
}

// Bar returns the next Bar
// visited by a call to Next.
func (i *Iter) Bar() *finance.ChartBar {
	return i.Current().(*finance.ChartBar)
}

// Meta returns the chart metadata
// related to a chart response.
func (i *Iter) Meta() finance.ChartMeta {
	return i.Iter.Meta().(finance.ChartMeta)
}

// Get returns a historical chart.
// and requires a params
// struct as an argument.
func Get(params *Params) *Iter {
	return getC().Get(params)
}

// Get returns a historical chart.
func (c Client) Get(params *Params) *Iter {

	// Construct request from params input.
	// TODO: validate symbol..
	if params == nil || len(params.Symbol) == 0 {
		return &Iter{iter.NewE(finance.CreateArgumentError())}
	}

	if params.Context == nil {
		ctx := context.TODO()
		params.Context = &ctx
	}

	// Start and End times
	params.start = -1
	params.end = -1
	if params.Start != nil {
		params.start = params.Start.Unix()
	}
	if params.End != nil {
		params.end = params.End.Unix()
	}
	if params.start > params.end {
		return &Iter{iter.NewE(finance.CreateChartTimeError())}
	}

	// Parse interval.
	if params.Interval != "" {
		params.interval = string(params.Interval)
	}

	// Build request.
	body := &form.Values{}
	form.AppendTo(body, params)
	// Set request meta data.
	body.Set("region", "US")
	body.Set("corsDomain", "com.finance.yahoo")

	return &Iter{iter.New(body, func(b *form.Values) (m interface{}, bars []interface{}, err error) {

		resp := response{}
		err = c.B.Call("v8/finance/chart/"+params.Symbol, body, params.Context, &resp)
		if err != nil {
			return
		}

		if resp.Inner.Error != nil {
			err = resp.Inner.Error
			return
		}

		result := resp.Inner.Results[0]
		if result == nil || result.Indicators == nil {
			err = finance.CreateRemoteErrorS("no results in chart response")
			return
		}

		barQuotes := result.Indicators.Quote
		if barQuotes == nil || barQuotes[0] == nil {
			err = finance.CreateRemoteErrorS("no results in chart response")
			return
		}
		adjCloses := result.Indicators.Adjclose

		// Process chart response
		// and chart meta data.
		for i, t := range result.Timestamp {

			b := &finance.ChartBar{
				Timestamp: t,
				Open:      decimal.NewFromFloat(barQuotes[0].Open[i]),
				High:      decimal.NewFromFloat(barQuotes[0].High[i]),
				Low:       decimal.NewFromFloat(barQuotes[0].Low[i]),
				Close:     decimal.NewFromFloat(barQuotes[0].Close[i]),
				Volume:    barQuotes[0].Volume[i],
			}

			if adjCloses != nil && adjCloses[0] != nil {
				b.AdjClose = decimal.NewFromFloat(adjCloses[0].Adjclose[i])
			}

			bars = append(bars, b)
		}

		return result.Meta, bars, nil
	})}
}

// response is a yfin chart response.
type response struct {
	Inner struct {
		Results []*result          `json:"result"`
		Error   *finance.YfinError `json:"error"`
	} `json:"chart"`
}

// result is an umbrella object for chart results.
type result struct {
	Meta       finance.ChartMeta `json:"meta"`
	Timestamp  []int             `json:"timestamp"`
	Indicators *struct {
		Quote []*struct {
			Open   []float64 `json:"open"`
			Low    []float64 `json:"low"`
			High   []float64 `json:"high"`
			Close  []float64 `json:"close"`
			Volume []int     `json:"volume"`
		} `json:"quote"`
		Adjclose []*struct {
			Adjclose []float64 `json:"adjclose"`
		} `json:"adjclose"`
	} `json:"indicators"`
}
