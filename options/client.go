package options

import (
	"context"
	"encoding/json"

	finance "github.com/piquette/finance-go"
	"github.com/piquette/finance-go/datetime"
	form "github.com/piquette/finance-go/form"
	"github.com/piquette/finance-go/iter"
)

// Client is used to invoke options APIs.
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
	UnderlyingSymbol string             `form:"-"`
	Expiration       *datetime.Datetime `form:"-"`
	date             int                `form:"date"`
	straddle         bool               `form:"straddle"`
}

// StraddleIter is a structure containing results
// and related metadata for a
// yfin option straddles request.
type StraddleIter struct {
	*iter.Iter
}

// Straddle returns the current straddle in the iter.
func (si *StraddleIter) Straddle() *finance.Straddle {
	return si.Current().(*finance.Straddle)
}

// Meta returns the metadata associated with the options response.
func (si *StraddleIter) Meta() *finance.OptionsMeta {
	return si.Iter.Meta().(*finance.OptionsMeta)
}

// GetStraddle returns options straddles.
// and requires a underlier symbol as an argument.
func GetStraddle(underlier string) *StraddleIter {
	return GetStraddleP(&Params{UnderlyingSymbol: underlier})
}

// GetStraddleP returns options straddles.
// and requires a params struct as an argument.
func GetStraddleP(params *Params) *StraddleIter {
	return getC().GetStraddleP(params)
}

// GetStraddleP returns options straddles.
// and requires a params struct as an argument.
func (c Client) GetStraddleP(params *Params) *StraddleIter {

	// Construct request from params input.
	// TODO: validate symbol..
	if params == nil || len(params.UnderlyingSymbol) == 0 {
		return &StraddleIter{iter.NewE(finance.CreateArgumentError())}
	}

	if params.Context == nil {
		ctx := context.TODO()
		params.Context = &ctx
	}

	params.straddle = true
	params.date = -1
	if params.Expiration != nil {
		params.date = params.Expiration.Unix()
	}

	body := &form.Values{}
	form.AppendTo(body, params)

	return &StraddleIter{iter.New(body, func(b *form.Values) (meta interface{}, values []interface{}, err error) {

		resp := response{}
		err = c.B.Call("v7/finance/options/"+params.UnderlyingSymbol, body, params.Context, &resp)
		if err != nil {
			return
		}

		if resp.Inner.Error != nil {
			err = resp.Inner.Error
			return
		}

		result := resp.Inner.Results[0]
		if result == nil {
			err = finance.CreateRemoteErrorS("no results in option straddle response")
			return
		}

		var list []straddleOptions
		err = json.Unmarshal(result.Options, &list)
		if err != nil || len(list) < 1 {
			err = finance.CreateRemoteErrorS("no results in option straddle response")
			return
		}
		ls := list[0]

		meta = &finance.OptionsMeta{
			UnderlyingSymbol:   result.UnderlyingSymbol,
			ExpirationDate:     ls.ExpirationDate,
			AllExpirationDates: result.ExpirationDates,
			Strikes:            result.Strikes,
			HasMiniOptions:     ls.HasMiniOptions,
			Quote:              result.Quote,
		}
		for _, straddle := range ls.Straddles {
			values = append(values, straddle)
		}

		return
	})}
}

// response is a yfin option response.
type response struct {
	Inner struct {
		Results []*result          `json:"result"`
		Error   *finance.YfinError `json:"error"`
	} `json:"optionChain"`
}

// result is an umbrella struct for options information for a specified underlying symbol.
type result struct {
	UnderlyingSymbol string          `json:"underlyingSymbol"`
	ExpirationDates  []int           `json:"expirationDates"`
	Strikes          []float64       `json:"strikes"`
	HasMiniOptions   bool            `json:"hasMiniOptions"`
	Quote            *finance.Quote  `json:"quote"`
	Options          json.RawMessage `json:"options"`
}

// chain is an options chain of puts/calls.
type chainOptions struct {
	ExpirationDate int                 `json:"expirationDate"`
	HasMiniOptions bool                `json:"hasMiniOptions"`
	Calls          []*finance.Contract `json:"calls"`
	Puts           []*finance.Contract `json:"puts"`
}

// straddles is a list of option straddles.
type straddleOptions struct {
	ExpirationDate int                 `json:"expirationDate"`
	HasMiniOptions bool                `json:"hasMiniOptions"`
	Straddles      []*finance.Straddle `json:"straddles"`
}
