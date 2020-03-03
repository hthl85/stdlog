package stdlog

// IStdLog interface
type IStdLog interface {
	Info(msg string)
	Infof(event ErrorLog)
	Error(msg string)
	Errorf(event ErrorLog)
	Fatal(msg string)
	Fatalf(event ErrorLog)
	Panic(msg string)
	Panicf(event ErrorLog)
}
