package equity

import (
	finance "github.com/piquette/finance-go"
	"github.com/piquette/finance-go/form"
)

// Client is used to invoke /quote APIs.
type Client struct {
	B finance.Backend
}

// Response represents a quote response.
type Response struct {
	finance.YFinResponse `json:"quoteResponse"`
}

func (r *Response) getFirst() *finance.Quote {
	if len(r.Result) == 0 {
		return nil
	}

	return r.Result[0].(*finance.Quote)
}

// Iter is an iterator for a list of quotes.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*finance.Iter
}

// Quote returns the most recent Quote
// visited by a call to Next.
func (i *Iter) Quote() *finance.Quote {
	return i.Current().(*finance.Quote)
}

func getC() Client {
	return Client{finance.GetBackend(finance.YFinBackend)}
}

// Get returns a quote that matches the parameters specified.
func Get(params *finance.QuoteParams) (*finance.Quote, error) {
	return getC().Get(params)
}

// Get returns a quote that matches the parameters specified.
func (c Client) Get(params *finance.QuoteParams) (*finance.Quote, error) {

	if params == nil {
		return nil, &finance.Error{
			Code:        finance.ErrorCodeArguments,
			Description: finance.ErrorDescriptionSymbols,
		}
	}

	//var body *form.Values
	body := &form.Values{}
	form.AppendTo(body, params)

	resp := &Response{}
	err := c.B.Call("GET", "/v7/finance/quote", body, params.Context, resp)
	if err != nil {
		return nil, err
	}

	return resp.getFirst(), resp.Error
}

// List returns several quotes.
func List(params *finance.QuoteParams) *Iter {
	return getC().List(params)
}

// List returns several quotes.
func (c Client) List(params *finance.QuoteParams) *Iter {
	return &Iter{}
}
