package logger

import (
	"log"
	"os"
)

type Logger interface {
	Info(msg string)
	Error(msg string)
}

type StandardLogger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func NewStandardLogger() *StandardLogger {
	return &StandardLogger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *StandardLogger) Info(msg string) {
	l.infoLogger.Println(msg)
}

func (l *StandardLogger) Error(msg string) {
	l.errorLogger.Println(msg)
}
