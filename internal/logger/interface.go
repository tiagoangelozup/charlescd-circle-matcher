package logger

type Interface interface {
	Critical(err error, msg string)
	Criticalf(err error, format string, args ...interface{})
	Debug(msg string)
	Debugf(format string, args ...interface{})
	Error(err error, msg string)
	Errorf(err error, format string, args ...interface{})
	Info(msg string)
	Infof(format string, args ...interface{})
	Warn(msg string)
	Warnf(format string, args ...interface{})
}
