package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

/*
type ConsoleLogger struct {
	logger *logrus.Logger
}*/

func New() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{}) // Формат JSON
	log.SetOutput(os.Stdout)                  // Вывод в стандартный поток
	log.SetLevel(logrus.InfoLevel)            // Уровень логирования
	return log
}
