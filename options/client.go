package options

import finance "github.com/piquette/finance-go"

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
	Underlier string `form:"-"`
	//Start  *Datetime `form:"-"`
}

// Iter is a structure containing results
// and related metadata for a
// yfin options request.
type Iter struct {
	*finance.Iter
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
	UnderlyingSymbol string         `json:"underlyingSymbol"`
	ExpirationDates  []int          `json:"expirationDates"`
	Strikes          []float64      `json:"strikes"`
	HasMiniOptions   bool           `json:"hasMiniOptions"`
	Quote            *finance.Quote `json:"quote"`
	Data             []byte         `json:"options"`
}

// chain is an options chain of puts/calls.
type chain struct {
	ExpirationDate int                `json:"expirationDate"`
	HasMiniOptions bool               `json:"hasMiniOptions"`
	Calls          []finance.Contract `json:"calls"`
	Puts           []finance.Contract `json:"puts"`
}

// straddles is a list of option straddles.
type straddles struct {
	ExpirationDate int                `json:"expirationDate"`
	HasMiniOptions bool               `json:"hasMiniOptions"`
	Straddles      []finance.Straddle `json:"straddles"`
}
