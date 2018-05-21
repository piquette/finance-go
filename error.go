package finance

import (
	"fmt"
)

const (
	// apiErrorCode denotes an error caught by finance-go
	// stemming from invalid user inputs.
	apiErrorCode = "api-error"

	// remoteErrorCode denotes an error
	// communicated in a reponse from a
	// remote api source.
	remoteErrorCode = "remote-error"
)

// CreateArgumentError returns an error
// with a message about missing arguments.
func CreateArgumentError() error {
	return fmt.Errorf("code: %s, detail: %s", apiErrorCode, "missing function argument")
}

// CreateRemoteError returns an error
// with a message about a remote api problem.
func CreateRemoteError(e error) error {
	return fmt.Errorf("code: %s, detail: %s", remoteErrorCode, e.Error())
}
