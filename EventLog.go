package stdlog

// EventLog struct
type EventLog struct {
	Code    int
	Message string
}

// String return message from an event log
func (event EventLog) String() string {
	msg := "unknown"

	if event.Message != "" {
		msg = event.Message
	}

	return msg
}
