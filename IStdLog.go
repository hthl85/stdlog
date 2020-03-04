package stdlog

// IStdLog interface
type IStdLog interface {
	Infof(info string)
	Error(err error)
	Errorf(msg string)
}
