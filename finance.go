package finance

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/piquette/finance-go/form"
)

// Printfer is an interface to be implemented by Logger.
type Printfer interface {
	Printf(format string, v ...interface{})
}

// init sets inital logger defaults.
func init() {
	Logger = log.New(os.Stderr, "", log.LstdFlags)
}

const (
	// YFinBackend is a constant representing the yahoo service backend.
	YFinBackend SupportedBackend = "yahoo"
	// YFinURL is the URL of the yahoo service backend.
	YFinURL string = "https://query1.finance.yahoo.com"
	// BATSBackend is a constant representing the uploads service backend.
	BATSBackend SupportedBackend = "bats"
	// BATSURL is the URL of the uploads service backend.
	BATSURL string = ""

	// Private constants.
	// ------------------

	defaultHTTPTimeout = 80 * time.Second
	yFinURL            = "https://query1.finance.yahoo.com"
	batsURL            = ""

	crumbURL  = yFinURL + "/v1/test/getcrumb"
	cookieURL = "https://login.yahoo.com"
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/113.0"
)

var (
	// LogLevel is the logging level for this library.
	// 0: no logging
	// 1: errors only
	// 2: errors + informational (default)
	// 3: errors + informational + debug
	LogLevel = 0

	// Logger controls how this library performs logging at a package level. It is useful
	// to customise if you need it prefixed for your application to meet other
	// requirements
	Logger Printfer

	// Private vars.
	// -------------

	httpClient = &http.Client{Timeout: defaultHTTPTimeout}
	backends   Backends
)

// SupportedBackend is an enumeration of supported api endpoints.
type SupportedBackend string

// Backends are the currently supported endpoints.
type Backends struct {
	YFin, Bats Backend
	mu         sync.RWMutex
}

// BackendConfiguration is the internal implementation for making HTTP calls.
type BackendConfiguration struct {
	Type       SupportedBackend
	URL        string
	HTTPClient *http.Client
}

// yahooConfiguration is a specialization that includes a crumb and cookies for the yahoo API
type yahooConfiguration struct {
	BackendConfiguration
	expiry  time.Time
	cookies string
	crumb   string
}

// Backend is an interface for making calls against an api service.
// This interface exists to enable mocking for during testing if needed.
type Backend interface {
	Call(path string, body *form.Values, ctx *context.Context, v interface{}) error
}

// SetHTTPClient overrides the default HTTP client.
// This is useful if you're running in a Google AppEngine environment
// where the http.DefaultClient is not available.
func SetHTTPClient(client *http.Client) {
	httpClient = client
}

// NewBackends creates a new set of backends with the given HTTP client. You
// should only need to use this for testing purposes or on App Engine.
func NewBackends(httpClient *http.Client) *Backends {
	return &Backends{
		YFin: &yahooConfiguration{
			BackendConfiguration{YFinBackend, YFinURL, httpClient},
			time.Time{},
			"",
			"",
		},
		Bats: &BackendConfiguration{
			BATSBackend, BATSURL, httpClient,
		},
	}
}

// GetBackend returns the currently used backend in the binding.
func GetBackend(backend SupportedBackend) Backend {
	switch backend {
	case YFinBackend:
		backends.mu.RLock()
		ret := backends.YFin
		backends.mu.RUnlock()
		if ret != nil {
			return ret
		}
		backends.mu.Lock()
		defer backends.mu.Unlock()
		backends.YFin = &yahooConfiguration{
			BackendConfiguration{YFinBackend, YFinURL, httpClient},
			time.Time{},
			"",
			"",
		}
		return backends.YFin
	case BATSBackend:
		backends.mu.RLock()
		ret := backends.Bats
		backends.mu.RUnlock()
		if ret != nil {
			return ret
		}
		backends.mu.Lock()
		defer backends.mu.Unlock()
		backends.Bats = &BackendConfiguration{backend, batsURL, httpClient}
		return backends.Bats
	}

	return nil
}

// SetBackend sets the backend used in the binding.
func SetBackend(backend SupportedBackend, b Backend) {
	switch backend {
	case YFinBackend:
		backends.YFin = b
	case BATSBackend:
		backends.Bats = b
	}
}

func fetchCookies() (string, time.Time, error) {
	client := http.Client{}
	request, err := http.NewRequest("GET", cookieURL, nil)
	if err != nil {
		return "", time.Time{}, err
	}

	request.Header = http.Header{
		"Accept":                   {"*/*"},
		"Accept-Encoding":          {"gzip, deflate, br"},
		"Accept-Language":          {"en-US,en;q=0.5"},
		"Connection":               {"keep-alive"},
		"Host":                     {"login.yahoo.com"},
		"Sec-Fetch-Dest":           {"document"},
		"Sec-Fetch-Mode":           {"navigate"},
		"Sec-Fetch-Site":           {"none"},
		"Sec-Fetch-User":           {"?1"},
		"TE":                       {"trailers"},
		"Update-Insecure-Requests": {"1"},
		"User-Agent":               {userAgent},
	}

	response, err := client.Do(request)
	if err != nil {
		return "", time.Time{}, err
	}
	defer response.Body.Close()

	var result string
	// create a variable expiry that is one year in the future
	var expiry = time.Now().AddDate(1, 0, 0)

	for _, cookie := range response.Cookies() {
		var unixTime = cookie.Expires.Unix()

		// Skip invalid cookies
		if unixTime <= 0 {
			continue
		}

		if cookie.Name != "AS" {
			result += cookie.Name + "=" + cookie.Value + "; "
			// set expiry to the latest cookie expiry if smaller than the current expiry
			if cookie.Expires.Before(expiry) {
				expiry = cookie.Expires
			}
		}
	}
	result = strings.TrimSuffix(result, "; ")
	return result, expiry, nil
}

func fetchCrumb(cookies string) (string, error) {
	client := http.Client{}
	request, err := http.NewRequest("GET", crumbURL, nil)
	if err != nil {
		return "", err
	}

	request.Header = http.Header{
		"Accept":          {"*/*"},
		"Accept-Encoding": {"gzip, deflate, br"},
		"Accept-Language": {"en-US,en;q=0.5"},
		"Connection":      {"keep-alive"},
		"Content-Type":    {"text/plain"},
		"Cookie":          {cookies},
		"Host":            {"query1.finance.yahoo.com"},
		"Sec-Fetch-Dest":  {"empty"},
		"Sec-Fetch-Mode":  {"cors"},
		"Sec-Fetch-Site":  {"same-site"},
		"TE":              {"trailers"},
		"User-Agent":      {userAgent},
	}

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body[:]), nil
}

func (s *yahooConfiguration) refreshCrumb() error {
	cookies, expiry, err := fetchCookies()
	if err != nil {
		return err
	}

	crumb, err := fetchCrumb(cookies)
	if err != nil {
		return err
	}

	s.crumb = crumb
	s.expiry = expiry
	s.cookies = cookies
	return nil
}

// Call is the Backend.Call implementation for invoking market data APIs, using the Yahoo specialization
func (s *yahooConfiguration) Call(path string, form *form.Values, ctx *context.Context, v interface{}) error {
	// Check if the cookies have expired.
	if s.expiry.Before(time.Now()) {
		// Refresh the cookies and crumb.
		err := s.refreshCrumb()
		if err != nil {
			return err
		}
	}

	if s.crumb != "" {
		form.Add("crumb", s.crumb)
	}

	if form != nil && !form.Empty() {
		path += "?" + form.Encode()
	}

	req, err := s.newRequest("GET", path, ctx)
	if err != nil {
		return err
	}

	if err := s.do(req, v); err != nil {
		return err
	}

	return nil
}

// Call is the Backend.Call implementation for invoking market data APIs.
func (s *BackendConfiguration) Call(path string, form *form.Values, ctx *context.Context, v interface{}) error {

	if form != nil && !form.Empty() {
		path += "?" + form.Encode()
	}

	req, err := s.newRequest("GET", path, ctx)
	if err != nil {
		return err
	}

	if err := s.do(req, v); err != nil {
		return err
	}

	return nil
}

func (s *yahooConfiguration) newRequest(method, path string, ctx *context.Context) (*http.Request, error) {
	req, err := s.BackendConfiguration.newRequest(method, path, ctx)

	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"Accept":          {"*/*"},
		"Accept-Language": {"en-US,en;q=0.5"},
		"Connection":      {"keep-alive"},
		"Content-Type":    {"application/json"},
		"Cookie":          {s.cookies},
		"Host":            {"query1.finance.yahoo.com"},
		"Origin":          {"https://finance.yahoo.com"},
		"Referer":         {"https://finance.yahoo.com"},
		"Sec-Fetch-Dest":  {"empty"},
		"Sec-Fetch-Mode":  {"cors"},
		"Sec-Fetch-Site":  {"same-site"},
		"TE":              {"trailers"},
		"User-Agent":      {userAgent},
	}

	return req, nil
}

func (s *BackendConfiguration) newRequest(method, path string, ctx *context.Context) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = s.URL + path

	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		if LogLevel > 0 {
			Logger.Printf("Cannot create api request: %v\n", err)
		}
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(*ctx)
	}

	return req, nil
}

// do is used by Call to execute an API request and parse the response. It uses
// the backend's HTTP client to execute the request and unmarshals the response
// into v. It also handles unmarshaling errors returned by the API.
func (s *BackendConfiguration) do(req *http.Request, v interface{}) error {
	if LogLevel > 1 {
		Logger.Printf("Requesting %v %v%v\n", req.Method, req.URL.Host, req.URL.Path)
	}

	start := time.Now()

	res, err := s.HTTPClient.Do(req)

	if LogLevel > 2 {
		Logger.Printf("Completed in %v\n", time.Since(start))
	}

	if err != nil {
		if LogLevel > 0 {
			Logger.Printf("Request to api failed: %v\n", err)
		}
		return err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		if LogLevel > 0 {
			Logger.Printf("Cannot parse response: %v\n", err)
		}
		return err
	}

	if res.StatusCode >= 400 {
		if LogLevel > 0 {
			Logger.Printf("API error: %q\n", resBody)
		}
		return CreateRemoteErrorS("error response recieved from upstream api")
	}

	if LogLevel > 2 {
		Logger.Printf("API response: %q\n", resBody)
	}

	if v != nil {
		return json.Unmarshal(resBody, v)
	}

	return nil
}
