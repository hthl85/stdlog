package stdlog

// IStdLog interface
type IStdLog interface {
	Info(msg string)
	Infof(event LogEvent)
	Error(msg string)
	Errorf(event LogEvent)
	Fatal(msg string)
	Fatalf(event LogEvent)
	Panic(msg string)
	Panicf(event LogEvent)
}
