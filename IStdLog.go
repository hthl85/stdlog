package stdlog

// IStdLog interface
type IStdLog interface {
	Info(info string)
	Error(err string)
	Errorf(err error)
	Fatal(err string)
	Fatalf(err error)
	Panic(err string)
	Panicf(err error)
}
