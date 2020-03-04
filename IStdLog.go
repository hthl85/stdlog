package stdlog

// IStdLog interface
type IStdLog interface {
	Info(err string)
	Error(err string)
	Errorf(err error)
	Fatal(err string)
	Fatalf(err error)
	Panic(err string)
	Panicf(err error)
}
