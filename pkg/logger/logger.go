package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func Init() {
	Log = logrus.New()
	Log.SetFormatter(&logrus.JSONFormatter{}) // Формат JSON
	Log.SetOutput(os.Stdout)                  // Вывод в стандартный поток
	Log.SetLevel(logrus.InfoLevel)            // Уровень логирования
}
