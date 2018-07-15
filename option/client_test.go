package option

import (
	"testing"

	finance "github.com/piquette/finance-go"
	tests "github.com/piquette/finance-go/testing"
	"github.com/stretchr/testify/assert"
)

func TestGetOption(t *testing.T) {
	tests.SetMarket(finance.MarketStateRegular)

	q, err := Get(tests.TestOptionSymbol)

	assert.Nil(t, err)
	assert.NotNil(t, q)
	assert.Equal(t, finance.MarketStateRegular, q.MarketState)
	assert.Equal(t, tests.TestOptionSymbol, q.Symbol)
}
