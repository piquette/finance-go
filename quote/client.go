package quote

import (
	"context"
	"strings"

	finance "github.com/piquette/finance-go"
	"github.com/piquette/finance-go/form"
	"github.com/piquette/finance-go/iter"
)

// Client is used to invoke quote APIs.
type Client struct {
	B finance.Backend
}

func getC() Client {
	return Client{finance.GetBackend(finance.YFinBackend)}
}

// Params carries a context and symbols information.
type Params struct {
	finance.Params `form:"-"`
	// Symbols are the symbols for which a
	// quote is requested.
	Symbols []string `form:"-"`
	sym     string   `form:"symbols"`
}

// Iter is an iterator for a list of quotes.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*iter.Iter
}

// Quote returns the most recent Quote
// visited by a call to Next.
func (i *Iter) Quote() *finance.Quote {
	return i.Current().(*finance.Quote)
}

// Get returns an Quote quote that matches the parameters specified.
func Get(symbol string) (*finance.Quote, error) {
	i := List([]string{symbol})

	if !i.Next() {
		return nil, i.Err()
	}

	return i.Quote(), nil
}

// List returns several quotes.
func List(symbols []string) *Iter {
	return ListP(&Params{Symbols: symbols})
}

// ListP returns a quote iterator and requires a params
// struct as an argument.
func ListP(params *Params) *Iter {
	return getC().ListP(params)
}

// ListP returns a quote iterator.
func (c Client) ListP(params *Params) *Iter {

	if params.Context == nil {
		ctx := context.TODO()
		params.Context = &ctx
	}

	// Validate input.
	// TODO: validate symbols..
	if params == nil || len(params.Symbols) == 0 {
		return &Iter{iter.NewE(finance.CreateArgumentError())}
	}
	params.sym = strings.Join(params.Symbols, ",")

	body := &form.Values{}
	form.AppendTo(body, params)

	return &Iter{iter.New(body, func(b *form.Values) (interface{}, []interface{}, error) {

		resp := response{}
		err := c.B.Call("/v7/finance/quote", body, params.Context, &resp)
		if err != nil {
			err = finance.CreateRemoteError(err)
		}

		ret := make([]interface{}, len(resp.Inner.Result))
		for i, v := range resp.Inner.Result {
			ret[i] = v
		}
		if resp.Inner.Error != nil {
			err = finance.CreateRemoteError(resp.Inner.Error)
		}

		return nil, ret, err
	})}
}

// response is a yfin quote response.
type response struct {
	Inner struct {
		Result []*finance.Quote   `json:"result"`
		Error  *finance.YfinError `json:"error"`
	} `json:"quoteResponse"`
}
