package stdlog

// IStdLog interface
type IStdLog interface {
	LogInfo(info string)
	LogError(err error) error
	LogErrorMessage(msg string) error
}
