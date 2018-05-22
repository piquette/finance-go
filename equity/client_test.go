package equity

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
