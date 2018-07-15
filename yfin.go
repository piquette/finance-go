package finance

import (
	"context"
	"encoding/json"

	"github.com/shopspring/decimal"
)

//
// Contains all the structs / info needed to
// consume and parse yfin apis.
//

// YfinError represents information returned in an error response.
type YfinError struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

// Error serializes the error object to JSON and returns it as a string.
func (e *YfinError) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}

type (
	// QuoteType alias for asset classification.
	QuoteType string
	// MarketState alias for market state.
	MarketState string
)

const (
	// QuoteTypeEquity the returned quote should be an equity.
	QuoteTypeEquity QuoteType = "EQUITY"
	// QuoteTypeIndex the returned quote should be an index.
	QuoteTypeIndex QuoteType = "INDEX"
	// QuoteTypeOption the returned quote should be an option contract.
	QuoteTypeOption QuoteType = "OPTION"
	// QuoteTypeForexPair the returned quote should be a forex pair.
	QuoteTypeForexPair QuoteType = "CURRENCY"
	// QuoteTypeCryptoPair the returned quote should be a crypto pair.
	QuoteTypeCryptoPair QuoteType = "CRYPTOCURRENCY"
	// QuoteTypeFuture the returned quote should be a futures contract.
	QuoteTypeFuture QuoteType = "FUTURE"
	// QuoteTypeETF the returned quote should be an etf.
	QuoteTypeETF QuoteType = "ETF"
	// QuoteTypeMutualFund the returned quote should be an mutual fund.
	QuoteTypeMutualFund QuoteType = "MUTUALFUND"

	// MarketStatePrePre pre-pre market state.
	MarketStatePrePre MarketState = "PREPRE"
	// MarketStatePre pre market state.
	MarketStatePre MarketState = "PRE"
	// MarketStateRegular regular market state.
	MarketStateRegular MarketState = "REGULAR"
	// MarketStatePost post market state.
	MarketStatePost MarketState = "POST"
	// MarketStatePostPost post-post market state.
	MarketStatePostPost MarketState = "POSTPOST"
	// MarketStateClosed closed market state.
	MarketStateClosed MarketState = "CLOSED"
)

// Params used as a parameter to many api functions.
type Params struct {
	// Context used for request. It may carry deadlines, cancelation signals,
	// and other request-scoped values across API boundaries and between
	// processes.
	// Note that a cancelled or timed out context does not provide any
	// guarantee whether the operation was or was not completed.
	Context *context.Context `form:"-"`
}

// Equity representa a single equity quote.
type Equity struct {
	Quote
	// Equity-only fields.
	LongName                    string  `json:"longName"`
	EpsTrailingTwelveMonths     float64 `json:"epsTrailingTwelveMonths"`
	EpsForward                  float64 `json:"epsForward"`
	EarningsTimestamp           int     `json:"earningsTimestamp"`
	EarningsTimestampStart      int     `json:"earningsTimestampStart"`
	EarningsTimestampEnd        int     `json:"earningsTimestampEnd"`
	TrailingAnnualDividendRate  float64 `json:"trailingAnnualDividendRate"`
	DividendDate                int     `json:"dividendDate"`
	TrailingAnnualDividendYield float64 `json:"trailingAnnualDividendYield"`
	TrailingPE                  float64 `json:"trailingPE"`
	ForwardPE                   float64 `json:"forwardPE"`
	BookValue                   float64 `json:"bookValue"`
	PriceToBook                 float64 `json:"priceToBook"`
	SharesOutstanding           int     `json:"sharesOutstanding"`
	MarketCap                   int64   `json:"marketCap"`
}

// ETF represents a single etf quote.
type ETF struct {
	Quote
	// MutualFund/ETF-only fields.
	YTDReturn                    float64 `json:"ytdReturn"`
	TrailingThreeMonthReturns    float64 `json:"trailingThreeMonthReturns"`
	TrailingThreeMonthNavReturns float64 `json:"trailingThreeMonthNavReturns"`
}

// MutualFund represents a single mutual fund share quote.
type MutualFund struct {
	Quote
	// MutualFund/ETF-only fields.
	YTDReturn                    float64 `json:"ytdReturn"`
	TrailingThreeMonthReturns    float64 `json:"trailingThreeMonthReturns"`
	TrailingThreeMonthNavReturns float64 `json:"trailingThreeMonthNavReturns"`
}

// Index represents a single market Index quote.
// The term `quote` here doesn't really apply in
// a practical sense, as indicies themselves are
// by definition not tradable assets.
type Index struct {
	Quote
}

// Option represents a single option contract quote
// for a specified strike and expiration.
type Option struct {
	Quote
	// Options/Futures-only fields.
	UnderlyingSymbol         string  `json:"underlyingSymbol"`
	OpenInterest             int     `json:"openInterest"`
	ExpireDate               int     `json:"expireDate"`
	Strike                   float64 `json:"strike"`
	UnderlyingExchangeSymbol string  `json:"underlyingExchangeSymbol"`
}

// Future represents a single futures contract quote
// for a specified strike and expiration.
type Future struct {
	Quote
	// Options/Futures-only fields.
	UnderlyingSymbol         string  `json:"underlyingSymbol"`
	OpenInterest             int     `json:"openInterest"`
	ExpireDate               int     `json:"expireDate"`
	Strike                   float64 `json:"strike"`
	UnderlyingExchangeSymbol string  `json:"underlyingExchangeSymbol"`
	HeadSymbolAsString       string  `json:"headSymbolAsString"`
	IsContractSymbol         bool    `json:"contractSymbol"`
}

// ForexPair represents a single forex currency pair quote.
type ForexPair struct {
	Quote
}

// CryptoPair represents a single crypto currency pair quote.
type CryptoPair struct {
	Quote
	// Cryptocurrency-only fields.
	Algorithm           string `json:"algorithm"`
	StartDate           int    `json:"startDate"`
	MaxSupply           int    `json:"maxSupply"`
	CirculatingSupply   int    `json:"circulatingSupply"`
	VolumeLastDay       int    `json:"volume24Hr"`
	VolumeAllCurrencies int    `json:"volumeAllCurrencies"`
}

// Quote is the basic quote structure shared across
// asset classes.
//
// Contains most fields that are common across all
// possible assets.
type Quote struct {
	// Quote classifying fields.
	Symbol      string      `json:"symbol"`
	MarketState MarketState `json:"marketState"`
	QuoteType   QuoteType   `json:"quoteType"`
	ShortName   string      `json:"shortName"`

	// Regular session quote data.
	RegularMarketChangePercent float64 `json:"regularMarketChangePercent"`
	RegularMarketPreviousClose float64 `json:"regularMarketPreviousClose"`
	RegularMarketPrice         float64 `json:"regularMarketPrice"`
	RegularMarketTime          int     `json:"regularMarketTime"`
	RegularMarketChange        float64 `json:"regularMarketChange"`
	RegularMarketOpen          float64 `json:"regularMarketOpen"`
	RegularMarketDayHigh       float64 `json:"regularMarketDayHigh"`
	RegularMarketDayLow        float64 `json:"regularMarketDayLow"`
	RegularMarketVolume        int     `json:"regularMarketVolume"`

	// Quote depth.
	Bid     float64 `json:"bid"`
	Ask     float64 `json:"ask"`
	BidSize int     `json:"bidSize"`
	AskSize int     `json:"askSize"`

	// Pre-market quote data.
	PreMarketPrice         float64 `json:"preMarketPrice"`
	PreMarketChange        float64 `json:"preMarketChange"`
	PreMarketChangePercent float64 `json:"preMarketChangePercent"`
	PreMarketTime          int     `json:"preMarketTime"`

	// Post-market quote data.
	PostMarketPrice         float64 `json:"postMarketPrice"`
	PostMarketChange        float64 `json:"postMarketChange"`
	PostMarketChangePercent float64 `json:"postMarketChangePercent"`
	PostMarketTime          int     `json:"postMarketTime"`

	// 52wk ranges.
	FiftyTwoWeekLowChange         float64 `json:"fiftyTwoWeekLowChange"`
	FiftyTwoWeekLowChangePercent  float64 `json:"fiftyTwoWeekLowChangePercent"`
	FiftyTwoWeekHighChange        float64 `json:"fiftyTwoWeekHighChange"`
	FiftyTwoWeekHighChangePercent float64 `json:"fiftyTwoWeekHighChangePercent"`
	FiftyTwoWeekLow               float64 `json:"fiftyTwoWeekLow"`
	FiftyTwoWeekHigh              float64 `json:"fiftyTwoWeekHigh"`

	// Averages.
	FiftyDayAverage                   float64 `json:"fiftyDayAverage"`
	FiftyDayAverageChange             float64 `json:"fiftyDayAverageChange"`
	FiftyDayAverageChangePercent      float64 `json:"fiftyDayAverageChangePercent"`
	TwoHundredDayAverage              float64 `json:"twoHundredDayAverage"`
	TwoHundredDayAverageChange        float64 `json:"twoHundredDayAverageChange"`
	TwoHundredDayAverageChangePercent float64 `json:"twoHundredDayAverageChangePercent"`

	// Volume metrics.
	AverageDailyVolume3Month int `json:"averageDailyVolume3Month"`
	AverageDailyVolume10Day  int `json:"averageDailyVolume10Day"`

	// Quote meta-data.
	QuoteSource               string `json:"quoteSourceName"`
	CurrencyID                string `json:"currency"`
	IsTradeable               bool   `json:"tradeable"`
	QuoteDelay                int    `json:"exchangeDataDelayedBy"`
	FullExchangeName          string `json:"fullExchangeName"`
	SourceInterval            int    `json:"sourceInterval"`
	ExchangeTimezoneName      string `json:"exchangeTimezoneName"`
	ExchangeTimezoneShortName string `json:"exchangeTimezoneShortName"`
	GMTOffSetMilliseconds     int    `json:"gmtOffSetMilliseconds"`
	MarketID                  string `json:"market"`
	ExchangeID                string `json:"exchange"`
}

// ChartBar is a single instance of a chart bar.
type ChartBar struct {
	Open      decimal.Decimal
	Low       decimal.Decimal
	High      decimal.Decimal
	Close     decimal.Decimal
	AdjClose  decimal.Decimal
	Volume    int
	Timestamp int
}

// ChartMeta is meta data associated with a chart response.
type ChartMeta struct {
	Currency             string    `json:"currency"`
	Symbol               string    `json:"symbol"`
	ExchangeName         string    `json:"exchangeName"`
	QuoteType            QuoteType `json:"instrumentType"`
	FirstTradeDate       int       `json:"firstTradeDate"`
	Gmtoffset            int       `json:"gmtoffset"`
	Timezone             string    `json:"timezone"`
	ExchangeTimezoneName string    `json:"exchangeTimezoneName"`
	ChartPreviousClose   float64   `json:"chartPreviousClose"`
	CurrentTradingPeriod struct {
		Pre struct {
			Timezone  string `json:"timezone"`
			Start     int    `json:"start"`
			End       int    `json:"end"`
			Gmtoffset int    `json:"gmtoffset"`
		} `json:"pre"`
		Regular struct {
			Timezone  string `json:"timezone"`
			Start     int    `json:"start"`
			End       int    `json:"end"`
			Gmtoffset int    `json:"gmtoffset"`
		} `json:"regular"`
		Post struct {
			Timezone  string `json:"timezone"`
			Start     int    `json:"start"`
			End       int    `json:"end"`
			Gmtoffset int    `json:"gmtoffset"`
		} `json:"post"`
	} `json:"currentTradingPeriod"`
	DataGranularity string   `json:"dataGranularity"`
	ValidRanges     []string `json:"validRanges"`
}

// OptionsMeta is meta data associated with an options response.
type OptionsMeta struct {
	UnderlyingSymbol   string
	ExpirationDate     int
	AllExpirationDates []int
	Strikes            []float64
	HasMiniOptions     bool
	Quote              *Quote
}

// Straddle is a put/call straddle for a particular strike.
type Straddle struct {
	Strike float64   `json:"strike"`
	Call   *Contract `json:"call,omitempty"`
	Put    *Contract `json:"put,omitempty"`
}

// Contract is a struct containing a single option contract, usually part of a chain.
type Contract struct {
	Symbol            string  `json:"contractSymbol"`
	Strike            float64 `json:"strike"`
	Currency          string  `json:"currency"`
	LastPrice         float64 `json:"lastPrice"`
	Change            float64 `json:"change"`
	PercentChange     float64 `json:"percentChange"`
	Volume            int     `json:"volume"`
	OpenInterest      int     `json:"openInterest"`
	Bid               float64 `json:"bid"`
	Ask               float64 `json:"ask"`
	Size              string  `json:"contractSize"`
	Expiration        int     `json:"expiration"`
	LastTradeDate     int     `json:"lastTradeDate"`
	ImpliedVolatility float64 `json:"impliedVolatility"`
	InTheMoney        bool    `json:"inTheMoney"`
}
