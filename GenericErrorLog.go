package stdlog

import "fmt"

// GenericErrorLog struct
type GenericErrorLog struct {
	Code    int
	Message string
}

// Error get error message from log event
func (err *GenericErrorLog) Error() string {
	errorMessage := "unknown error"

	if err.Message != "" {
		errorMessage = fmt.Sprintf("%s", err.Message)
	}

	if err.Code > 0 {
		errorMessage = fmt.Sprintf("error code %d : %s", err.Code, errorMessage)
	}

	return errorMessage
}
