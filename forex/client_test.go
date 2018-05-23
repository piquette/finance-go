package forex

import (
	"testing"

	finance "github.com/piquette/finance-go"
	tests "github.com/piquette/finance-go/testing"
	"github.com/stretchr/testify/assert"
)

func TestGetRegularMarketForexPair(t *testing.T) {
	tests.SetMarket(finance.MarketStateRegular)

	q, err := Get(tests.TestForexPairSymbol)

	assert.Nil(t, err)
	assert.NotNil(t, q)
	assert.Equal(t, finance.MarketStateRegular, q.MarketState)
	assert.Equal(t, tests.TestForexPairSymbol, q.Symbol)
}

func TestGetPostMarketForexPair(t *testing.T) {
	tests.SetMarket(finance.MarketStatePost)

	q, err := Get(tests.TestForexPairSymbol)

	assert.Nil(t, err)
	assert.NotNil(t, q)
	assert.Equal(t, finance.MarketStatePost, q.MarketState)
	assert.Equal(t, tests.TestForexPairSymbol, q.Symbol)
}

func TestGetPreMarketForexPair(t *testing.T) {
	tests.SetMarket(finance.MarketStatePre)

	q, err := Get(tests.TestForexPairSymbol)

	assert.Nil(t, err)
	assert.NotNil(t, q)
	assert.Equal(t, finance.MarketStatePre, q.MarketState)
	assert.Equal(t, tests.TestForexPairSymbol, q.Symbol)
}

func TestNilParamsForexPair(t *testing.T) {

	iter := List(nil)

	assert.False(t, iter.Next())
	assert.Equal(t, "code: api-error, detail: missing function argument", iter.Err().Error())
}

func TestGetBadForexPair(t *testing.T) {
	tests.SetMarket(finance.MarketStateRegular)

	q, err := Get("TEST")
	assert.Nil(t, q)
	assert.Nil(t, err)
}
