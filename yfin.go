package finance

import "encoding/json"

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
