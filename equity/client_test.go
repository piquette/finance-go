package equity

import (
	"testing"

	finance "github.com/piquette/finance-go"
	tests "github.com/piquette/finance-go/testing"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	tests.SetMarket(finance.MarketStateRegular)

	q, err := Get(tests.TestEquitySymbol)
	assert.Nil(t, err)
	assert.NotNil(t, q)
}
