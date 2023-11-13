package logger

type Logger interface {
	// Debug
	Debug(msg string, kvs map[string]interface{})
	Debugf(msg string, args ...any)

	// Info
	Info(msg string, kvs map[string]interface{})
	Infof(msg string, args ...any)

	// Error
	Error(msg string, err error)
	Errorf(msg string, args ...any)

	// Panic
	Panic(msg string, err error)
	Panicf(msg string, args ...any)

	Shutdown() error
}
