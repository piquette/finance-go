package testing

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	finance "github.com/piquette/finance-go"
	"github.com/piquette/finance-go/form"
)

// This file should contain any testing helpers that should be commonly
// available across all tests in the finance package.
const (
	// MockMinimumVersion is the minimum acceptable version for finance-mock.
	// It's here so that if the library depends on new endpoints or features
	// added in a more recent version of finance-mock, we can show people a
	// better error message instead of the test suite crashing with a bunch of
	// confusing 404 errors or the like.
	MockMinimumVersion = "0.0.5"
	TestServerAddr     = "localhost"

	// Symbols for testing asset classes.
	TestEquitySymbol     = "AAPL"
	TestETFSymbol        = "SPY"
	TestFutureSymbol     = "O=F"
	TestIndexSymbol      = "^GSPC"
	TestOptionSymbol     = "AMD180720C00003000"
	TestMutualFundSymbol = "INPSX"
	TestForexPairSymbol  = "USDGBP=X"
	TestCryptoPairSymbol = "BTC-USD"
	TestStraddleSymbol   = "AMD"
)

func init() {
	// Enable strict mode on form encoding so that we'll panic if any kind of
	// malformed param struct is detected
	form.Strict = true

	port := os.Getenv("FINANCE_MOCK_PORT")
	if port == "" {
		port = "12111"
	}

	resp, err := http.Get("http://" + TestServerAddr + ":" + port)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Couldn't reach finance-mock at `%s:%s`. Is "+
			"it running? Please see README for setup instructions.\n", TestServerAddr, port)
		os.Exit(1)
	}

	version := resp.Header.Get("Finance-Mock-Version")
	if version != "master" && compareVersions(version, MockMinimumVersion) > 0 {
		fmt.Fprintf(os.Stderr, "Your version of finance-mock (%s) is too old. The "+
			"minimum version to run this test suite is %s. Please see its "+
			"repository for upgrade instructions.\n", version, MockMinimumVersion)
		os.Exit(1)
	}

	finance.SetBackend(finance.YFinBackend, &finance.BackendConfiguration{
		Type:       finance.YFinBackend,
		URL:        "http://" + TestServerAddr + ":" + port,
		HTTPClient: &http.Client{},
	})
}

// SetMarket sets the test server to the state/session specified.
func SetMarket(state finance.MarketState) {
	// one of regular/post/pre
	var mktState string

	switch state {
	case finance.MarketStatePre,
		finance.MarketStatePrePre:
		mktState = "pre"
	case finance.MarketStatePost,
		finance.MarketStatePostPost:
		mktState = "post"
	default:
		mktState = "regular"
	}

	form := url.Values{}
	form.Add("state", mktState)

	// Post.
	resp, err := http.PostForm("http://localhost:12111/config/", form)
	if err != nil || resp.StatusCode != http.StatusOK {

		fmt.Println(resp.Request.Form)

		fmt.Fprintf(os.Stderr, "Couldn't change state of finance-mock. Is "+
			"it running? Please see README for setup instructions.\n")
		os.Exit(1)
	}

	// Success.
	fmt.Fprintf(os.Stdout, "Changed state of finance-mock to %s.\n", mktState)
}

// compareVersions compares two semantic version strings. We need this because
// with more complex double-digit numbers, lexical comparison breaks down.
func compareVersions(a, b string) (ret int) {
	as := strings.Split(a, ".")
	bs := strings.Split(b, ".")

	loopMax := len(bs)
	if len(as) > len(bs) {
		loopMax = len(as)
	}

	for i := 0; i < loopMax; i++ {
		var x, y string
		if len(as) > i {
			x = as[i]
		}
		if len(bs) > i {
			y = bs[i]
		}

		xi, _ := strconv.Atoi(x)
		yi, _ := strconv.Atoi(y)

		if xi > yi {
			ret = -1
		} else if xi < yi {
			ret = 1
		}
		if ret != 0 {
			break
		}
	}
	return
}
