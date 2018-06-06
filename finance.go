package finance

import (
	"context"
	"encoding/json"
	"io/ioutil"
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
	YFinURL string = "https://query2.finance.yahoo.com"
	// BATSBackend is a constant representing the uploads service backend.
	BATSBackend SupportedBackend = "bats"
	// BATSURL is the URL of the uploads service backend.
	BATSURL string = ""

	// Private constants.
	// ------------------

	defaultHTTPTimeout = 80 * time.Second
	yFinURL            = "https://query2.finance.yahoo.com"
	batsURL            = ""
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
		YFin: &BackendConfiguration{
			YFinBackend, YFinURL, httpClient,
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
		backends.YFin = &BackendConfiguration{backend, yFinURL, httpClient}
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

// Call is the Backend.Call implementation for invoking market data APIs.
func (s *BackendConfiguration) Call(path string, form *form.Values, ctx *context.Context, v interface{}) error {

	if form != nil && !form.Empty() {
		path += "?" + form.Encode()
	}

	req, err := s.NewRequest("GET", path, ctx)
	if err != nil {
		return err
	}

	if err := s.Do(req, v); err != nil {
		return err
	}

	return nil
}

// NewRequest is used by Call to generate an http.Request.
func (s *BackendConfiguration) NewRequest(method, path string, ctx *context.Context) (*http.Request, error) {
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

// Do is used by Call to execute an API request and parse the response. It uses
// the backend's HTTP client to execute the request and unmarshals the response
// into v. It also handles unmarshaling errors returned by the API.
func (s *BackendConfiguration) Do(req *http.Request, v interface{}) error {
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

	resBody, err := ioutil.ReadAll(res.Body)
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
