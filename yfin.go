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
	LongName                    string  `json:"longName" csv:"longName"`
	EpsTrailingTwelveMonths     float64 `json:"epsTrailingTwelveMonths" csv:"epsTrailingTwelveMonths"`
	EpsForward                  float64 `json:"epsForward" csv:"epsForward"`
	EarningsTimestamp           int     `json:"earningsTimestamp" csv:"earningsTimestamp"`
	EarningsTimestampStart      int     `json:"earningsTimestampStart" csv:"earningsTimestampStart"`
	EarningsTimestampEnd        int     `json:"earningsTimestampEnd" csv:"earningsTimestampEnd"`
	TrailingAnnualDividendRate  float64 `json:"trailingAnnualDividendRate" csv:"trailingAnnualDividendRate"`
	DividendDate                int     `json:"dividendDate" csv:"dividendDate"`
	TrailingAnnualDividendYield float64 `json:"trailingAnnualDividendYield" csv:"trailingAnnualDividendYield"`
	TrailingPE                  float64 `json:"trailingPE" csv:"trailingPE"`
	ForwardPE                   float64 `json:"forwardPE" csv:"forwardPE"`
	BookValue                   float64 `json:"bookValue" csv:"bookValue"`
	PriceToBook                 float64 `json:"priceToBook" csv:"priceToBook"`
	SharesOutstanding           int     `json:"sharesOutstanding" csv:"sharesOutstanding"`
	MarketCap                   int64   `json:"marketCap" csv:"marketCap"`
}

// ETF represents a single etf quote.
type ETF struct {
	Quote
	// MutualFund/ETF-only fields.
	YTDReturn                    float64 `json:"ytdReturn" csv:"ytdReturn"`
	TrailingThreeMonthReturns    float64 `json:"trailingThreeMonthReturns" csv:"trailingThreeMonthReturns"`
	TrailingThreeMonthNavReturns float64 `json:"trailingThreeMonthNavReturns" csv:"trailingThreeMonthNavReturns"`
}

// MutualFund represents a single mutual fund share quote.
type MutualFund struct {
	Quote
	// MutualFund/ETF-only fields.
	YTDReturn                    float64 `json:"ytdReturn" csv:"ytdReturn"`
	TrailingThreeMonthReturns    float64 `json:"trailingThreeMonthReturns" csv:"trailingThreeMonthReturns"`
	TrailingThreeMonthNavReturns float64 `json:"trailingThreeMonthNavReturns" csv:"trailingThreeMonthNavReturns"`
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
	UnderlyingSymbol         string  `json:"underlyingSymbol" csv:"underlyingSymbol"`
	OpenInterest             int     `json:"openInterest" csv:"openInterest"`
	ExpireDate               int     `json:"expireDate" csv:"expireDate"`
	Strike                   float64 `json:"strike" csv:"strike"`
	UnderlyingExchangeSymbol string  `json:"underlyingExchangeSymbol" csv:"underlyingExchangeSymbol"`
}

// Future represents a single futures contract quote
// for a specified strike and expiration.
type Future struct {
	Quote
	// Options/Futures-only fields.
	UnderlyingSymbol         string  `json:"underlyingSymbol" csv:"underlyingSymbol"`
	OpenInterest             int     `json:"openInterest" csv:"openInterest"`
	ExpireDate               int     `json:"expireDate" csv:"expireDate"`
	Strike                   float64 `json:"strike" csv:"strike"`
	UnderlyingExchangeSymbol string  `json:"underlyingExchangeSymbol" csv:"underlyingExchangeSymbol"`
	HeadSymbolAsString       string  `json:"headSymbolAsString" csv:"headSymbolAsString"`
	IsContractSymbol         bool    `json:"contractSymbol" csv:"contractSymbol"`
}

// ForexPair represents a single forex currency pair quote.
type ForexPair struct {
	Quote
}

// CryptoPair represents a single crypto currency pair quote.
type CryptoPair struct {
	Quote
	// Cryptocurrency-only fields.
	Algorithm           string `json:"algorithm" csv:"algorithm"`
	StartDate           int    `json:"startDate" csv:"startDate"`
	MaxSupply           int    `json:"maxSupply" csv:"maxSupply"`
	CirculatingSupply   int    `json:"circulatingSupply" csv:"circulatingSupply"`
	VolumeLastDay       int    `json:"volume24Hr" csv:"volume24Hr"`
	VolumeAllCurrencies int    `json:"volumeAllCurrencies" csv:"volumeAllCurrencies"`
}

// Quote is the basic quote structure shared across
// asset classes.
//
// Contains most fields that are common across all
// possible assets.
type Quote struct {
	// Quote classifying fields.
	Symbol      string      `json:"symbol" csv:"symbol"`
	MarketState MarketState `json:"marketState" csv:"marketState"`
	QuoteType   QuoteType   `json:"quoteType" csv:"quoteType"`
	ShortName   string      `json:"shortName" csv:"shortName"`

	// Regular session quote data.
	RegularMarketChangePercent float64 `json:"regularMarketChangePercent" csv:"regularMarketChangePercent"`
	RegularMarketPreviousClose float64 `json:"regularMarketPreviousClose" csv:"regularMarketPreviousClose"`
	RegularMarketPrice         float64 `json:"regularMarketPrice" csv:"regularMarketPrice"`
	RegularMarketTime          int     `json:"regularMarketTime" csv:"regularMarketTime"`
	RegularMarketChange        float64 `json:"regularMarketChange" csv:"regularMarketChange"`
	RegularMarketOpen          float64 `json:"regularMarketOpen" csv:"regularMarketOpen"`
	RegularMarketDayHigh       float64 `json:"regularMarketDayHigh" csv:"regularMarketDayHigh"`
	RegularMarketDayLow        float64 `json:"regularMarketDayLow" csv:"regularMarketDayLow"`
	RegularMarketVolume        int     `json:"regularMarketVolume" csv:"regularMarketVolume"`

	// Quote depth.
	Bid     float64 `json:"bid" csv:"bid"`
	Ask     float64 `json:"ask" csv:"ask"`
	BidSize int     `json:"bidSize" csv:"bidSize"`
	AskSize int     `json:"askSize" csv:"askSize"`

	// Pre-market quote data.
	PreMarketPrice         float64 `json:"preMarketPrice" csv:"preMarketPrice"`
	PreMarketChange        float64 `json:"preMarketChange" csv:"preMarketChange"`
	PreMarketChangePercent float64 `json:"preMarketChangePercent" csv:"preMarketChangePercent"`
	PreMarketTime          int     `json:"preMarketTime" csv:"preMarketTime"`

	// Post-market quote data.
	PostMarketPrice         float64 `json:"postMarketPrice" csv:"postMarketPrice"`
	PostMarketChange        float64 `json:"postMarketChange" csv:"postMarketChange"`
	PostMarketChangePercent float64 `json:"postMarketChangePercent" csv:"postMarketChangePercent"`
	PostMarketTime          int     `json:"postMarketTime" csv:"postMarketTime"`

	// 52wk ranges.
	FiftyTwoWeekLowChange         float64 `json:"fiftyTwoWeekLowChange" csv:"fiftyTwoWeekLowChange"`
	FiftyTwoWeekLowChangePercent  float64 `json:"fiftyTwoWeekLowChangePercent" csv:"fiftyTwoWeekLowChangePercent"`
	FiftyTwoWeekHighChange        float64 `json:"fiftyTwoWeekHighChange" csv:"fiftyTwoWeekHighChange"`
	FiftyTwoWeekHighChangePercent float64 `json:"fiftyTwoWeekHighChangePercent" csv:"fiftyTwoWeekHighChangePercent"`
	FiftyTwoWeekLow               float64 `json:"fiftyTwoWeekLow" csv:"fiftyTwoWeekLow"`
	FiftyTwoWeekHigh              float64 `json:"fiftyTwoWeekHigh" csv:"fiftyTwoWeekHigh"`

	// Averages.
	FiftyDayAverage                   float64 `json:"fiftyDayAverage" csv:"fiftyDayAverage"`
	FiftyDayAverageChange             float64 `json:"fiftyDayAverageChange" csv:"fiftyDayAverageChange"`
	FiftyDayAverageChangePercent      float64 `json:"fiftyDayAverageChangePercent" csv:"fiftyDayAverageChangePercent"`
	TwoHundredDayAverage              float64 `json:"twoHundredDayAverage" csv:"twoHundredDayAverage"`
	TwoHundredDayAverageChange        float64 `json:"twoHundredDayAverageChange" csv:"twoHundredDayAverageChange"`
	TwoHundredDayAverageChangePercent float64 `json:"twoHundredDayAverageChangePercent" csv:"twoHundredDayAverageChangePercent"`

	// Volume metrics.
	AverageDailyVolume3Month int `json:"averageDailyVolume3Month" csv:"averageDailyVolume3Month"`
	AverageDailyVolume10Day  int `json:"averageDailyVolume10Day" csv:"averageDailyVolume10Day"`

	// Quote meta-data.
	QuoteSource               string `json:"quoteSourceName" csv:"quoteSourceName"`
	CurrencyID                string `json:"currency" csv:"currency"`
	IsTradeable               bool   `json:"tradeable" csv:"tradeable"`
	QuoteDelay                int    `json:"exchangeDataDelayedBy" csv:"exchangeDataDelayedBy"`
	FullExchangeName          string `json:"fullExchangeName" csv:"fullExchangeName"`
	SourceInterval            int    `json:"sourceInterval" csv:"sourceInterval"`
	ExchangeTimezoneName      string `json:"exchangeTimezoneName" csv:"exchangeTimezoneName"`
	ExchangeTimezoneShortName string `json:"exchangeTimezoneShortName" csv:"exchangeTimezoneShortName"`
	GMTOffSetMilliseconds     int    `json:"gmtOffSetMilliseconds" csv:"gmtOffSetMilliseconds"`
	MarketID                  string `json:"market" csv:"market"`
	ExchangeID                string `json:"exchange" csv:"exchange"`
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

// OHLCHistoric is a historical quotation.
type OHLCHistoric struct {
	Open      float64
	Low       float64
	High      float64
	Close     float64
	AdjClose  float64
	Volume    int
	Timestamp int
}

// ChartMeta is meta data associated with a chart response.
type ChartMeta struct {
	Currency             string    `json:"currency" csv:"currency"`
	Symbol               string    `json:"symbol" csv:"symbol"`
	ExchangeName         string    `json:"exchangeName" csv:"exchangeName"`
	QuoteType            QuoteType `json:"instrumentType" csv:"instrumentType"`
	FirstTradeDate       int       `json:"firstTradeDate" csv:"firstTradeDate"`
	Gmtoffset            int       `json:"gmtoffset" csv:"gmtoffset"`
	Timezone             string    `json:"timezone" csv:"timezone"`
	ExchangeTimezoneName string    `json:"exchangeTimezoneName" csv:"exchangeTimezoneName"`
	ChartPreviousClose   float64   `json:"chartPreviousClose" csv:"chartPreviousClose"`
	CurrentTradingPeriod struct {
		Pre struct {
			Timezone  string `json:"timezone" csv:"timezone"`
			Start     int    `json:"start" csv:"start"`
			End       int    `json:"end" csv:"end"`
			Gmtoffset int    `json:"gmtoffset" csv:"gmtoffset"`
		} `json:"pre" csv:"pre_,inline"`
		Regular struct {
			Timezone  string `json:"timezone" csv:"timezone"`
			Start     int    `json:"start" csv:"start"`
			End       int    `json:"end" csv:"end"`
			Gmtoffset int    `json:"gmtoffset" csv:"gmtoffset"`
		} `json:"regular" csv:"regular_,inline"`
		Post struct {
			Timezone  string `json:"timezone" csv:"timezone"`
			Start     int    `json:"start" csv:"start"`
			End       int    `json:"end" csv:"end"`
			Gmtoffset int    `json:"gmtoffset" csv:"gmtoffset"`
		} `json:"post" csv:"post_,inline"`
	} `json:"currentTradingPeriod" csv:"currentTradingPeriod_,inline"`
	DataGranularity string   `json:"dataGranularity" csv:"dataGranularity"`
	ValidRanges     []string `json:"validRanges" csv:"-"`
}

// OptionsMeta is meta data associated with an options response.
type OptionsMeta struct {
	UnderlyingSymbol   string    `json:"underlyingSymbol" csv:"underlyingSymbol"`
	ExpirationDate     int       `json:"expirationDate" csv:"expirationDate"`
	AllExpirationDates []int     `json:"allExpirationDates" csv:"-"`
	Strikes            []float64 `json:"strikes" csv:"-"`
	HasMiniOptions     bool      `json:"hasMiniOptions"`
	Quote              *Quote    `json:"quote,omitempty" csv:"quote_,inline"`
}

// Straddle is a put/call straddle for a particular strike.
type Straddle struct {
	Strike float64   `json:"strike" csv:"strike"`
	Call   *Contract `json:"call,omitempty" csv:"call_,inline"`
	Put    *Contract `json:"put,omitempty" csv:"put_,inline"`
}

// Contract is a struct containing a single option contract, usually part of a chain.
type Contract struct {
	Symbol            string  `json:"contractSymbol" csv:"contractSymbol"`
	Strike            float64 `json:"strike" csv:"strike"`
	Currency          string  `json:"currency" csv:"currency"`
	LastPrice         float64 `json:"lastPrice" csv:"lastPrice"`
	Change            float64 `json:"change" csv:"change"`
	PercentChange     float64 `json:"percentChange" csv:"percentChange"`
	Volume            int     `json:"volume" csv:"volume"`
	OpenInterest      int     `json:"openInterest" csv:"openInterest"`
	Bid               float64 `json:"bid" csv:"bid"`
	Ask               float64 `json:"ask" csv:"ask"`
	Size              string  `json:"contractSize" csv:"contractSize"`
	Expiration        int     `json:"expiration" csv:"expiration"`
	LastTradeDate     int     `json:"lastTradeDate" csv:"lastTradeDate"`
	ImpliedVolatility float64 `json:"impliedVolatility" csv:"impliedVolatility"`
	InTheMoney        bool    `json:"inTheMoney" csv:"inTheMoney"`
}
