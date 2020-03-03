package stdlog

import "fmt"

// ErrorLog struct
type ErrorLog struct {
	Code    int
	Message string
}

// Error get error message from log event
func (event *ErrorLog) Error() string {
	return fmt.Sprintf("[ERROR] ")
}
