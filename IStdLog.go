package stdlog

// IStdLog interface
type IStdLog interface {
	Infof(msg string, args ...interface{})
	Error(err error)
	Errorf(msg string, args ...interface{})
}
