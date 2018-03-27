package finance

import "encoding/json"

// YFinResponse wraps a generic yfin api response.
type YFinResponse struct {
	Result []interface{} `json:"result"`
	Error  *YFinError    `json:"error"`
}

// YFinError is the response returned when a yfin call is unsuccessful.
type YFinError struct {
	Err *Error `json:"error"`
}

// Error serializes fields.
func (y *YFinError) Error() string {
	ret, _ := json.Marshal(y)
	return string(ret)
}
