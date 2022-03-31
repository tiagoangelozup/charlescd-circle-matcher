package logger

import (
	"fmt"
	"log"
)

type Local struct{}

func (l *Local) Debugf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (l *Local) Infof(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (l *Local) Warnf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (l *Local) Errorf(err error, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	msg = fmt.Sprintf("%s: %v", msg, err)
	log.Printf(msg)
}

func (l *Local) Criticalf(err error, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	msg = fmt.Sprintf("%s: %v", msg, err)
	log.Printf(msg)
}

func (l *Local) Debug(msg string) {
	log.Print(msg)
}

func (l *Local) Info(msg string) {
	log.Print(msg)
}

func (l *Local) Warn(msg string) {
	log.Print(msg)
}

func (l *Local) Error(err error, msg string) {
	msg = fmt.Sprintf("%s: %v", msg, err)
	log.Printf(msg)
}

func (l *Local) Critical(err error, msg string) {
	msg = fmt.Sprintf("%s: %v", msg, err)
	log.Printf(msg)
}
