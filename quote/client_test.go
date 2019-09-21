package quote

import (
	"testing"

	finance "github.com/piquette/finance-go"
	tests "github.com/piquette/finance-go/testing"
	"github.com/stretchr/testify/assert"
)

func TestGetRegularMarketQuote(t *testing.T) {
	tests.SetMarket(finance.MarketStateRegular)

	q, err := Get(tests.TestEquitySymbol)

	assert.Nil(t, err)
	assert.NotNil(t, q)
	assert.Equal(t, finance.MarketStateRegular, q.MarketState)
	assert.Equal(t, tests.TestEquitySymbol, q.Symbol)
}

func TestHistoricalQuote(t *testing.T) {
	TestMonth := 1
	TestDay := 11
	TestYear := 2018

	q, err := GetHistoricalQuote(tests.TestEquitySymbol, TestMonth, TestDay, TestYear)

	assert.Nil(t, err)
	assert.NotNil(t, q)
	high, _ := q.High.Float64()
	low, _ := q.Low.Float64()
	open, _ := q.Open.Float64()
	close, _ := q.Open.Float64()
	assert.Equal(t, 188.13999938964844, high)
	assert.Equal(t, 187.5500030517578, low)
	assert.Equal(t, 187.75, close)
	assert.Equal(t, 187.75, open)
}

func TestBadSymbolBar(t *testing.T) {
	chart, err := GetHistoricalQuote("BADSYMBOL", 1, 11, 2018)
	assert.Nil(t, chart)
	assert.NotNil(t, err)
}

func TestGetPostMarketQuote(t *testing.T) {
	tests.SetMarket(finance.MarketStatePost)

	q, err := Get(tests.TestEquitySymbol)

	assert.Nil(t, err)
	assert.NotNil(t, q)
	assert.Equal(t, finance.MarketStatePost, q.MarketState)
	assert.Equal(t, tests.TestEquitySymbol, q.Symbol)
}

func TestGetPreMarketQuote(t *testing.T) {
	tests.SetMarket(finance.MarketStatePre)

	q, err := Get(tests.TestEquitySymbol)

	assert.Nil(t, err)
	assert.NotNil(t, q)
	assert.Equal(t, finance.MarketStatePre, q.MarketState)
	assert.Equal(t, tests.TestEquitySymbol, q.Symbol)
}

func TestNilParamsQuote(t *testing.T) {

	iter := List(nil)

	assert.False(t, iter.Next())
	assert.Equal(t, "code: api-error, detail: missing function argument", iter.Err().Error())
}

func TestGetBadQuote(t *testing.T) {
	tests.SetMarket(finance.MarketStateRegular)

	q, err := Get("TEST")
	assert.Nil(t, q)
	assert.Nil(t, err)
}
