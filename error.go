package finance

import "encoding/json"

// ErrorCode error code returned by api.
type ErrorCode string

// ErrorDescription detailed error description returned by api.
type ErrorDescription string

const (
	// ErrorCodeArguments is triggered when an
	// exception related to a missing query
	// parameter occurs. Check Description
	// for more information.
	ErrorCodeArguments ErrorCode = "argument-error"
	// ErrorCodeAPI covers errors that occur
	// through api response parsing problems and
	// other edge cases.
	ErrorCodeAPI ErrorCode = "api-error"

	// ErrorDescriptionSymbols describes a possible
	// error scenario where the symbols parameter was
	// missing.
	ErrorDescriptionSymbols ErrorDescription = "Missing value for the \"symbols\" argument"
	// ErrorDescriptionParsing possible response parsing problem.
	ErrorDescriptionParsing ErrorDescription = "Error occured parsing an error response."
	// ErrorDescriptionResponse base response problem.
	ErrorDescriptionResponse ErrorDescription = "Error occurred."
)

// Error represents an error returned as a response.
type Error struct {
	Code        ErrorCode        `json:"code"`
	Description ErrorDescription `json:"description"`

	// Err contains an internal error with an additional level of granularity
	// that can be used in some cases to get more detailed information about
	// what went wrong.
	Err error `json:"-"`

	// HTTPStatusCode optional status code information.
	HTTPStatusCode int `json:"-"`
}

// Error serializes the error object to JSON and returns it as a string.
func (e *Error) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}
