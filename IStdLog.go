package stdlog

// IStdLog interface
type IStdLog interface {
	Infof(event LogEvent)
	Errorf(event LogEvent)
	Fatalf(event LogEvent)
	Panicf(event LogEvent)
}
