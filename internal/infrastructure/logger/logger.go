package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	log  *logrus.Logger
	file *os.File
}

func New(logLevel string, logFileName string) (*Logger, error) {
	log := logrus.New()
	logrusLevel, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}
	file, err := os.Create(logFileName)
	if err != nil {
		return nil, err
	}
	logrusOutput := io.MultiWriter(os.Stdout, file)
	log.SetOutput(logrusOutput)
	log.SetLevel(logrusLevel)
	log.SetFormatter(&logrus.JSONFormatter{})

	logger := &Logger{
		log: log,
	}
	return logger, nil
}
func (l *Logger) Debug(message string) {
	l.log.Debug(message)
}
func (l *Logger) Debugf(message string, arg []interface{}) {
	l.log.Debugf(message, arg)
}
func (l *Logger) Info(message string) {
	l.log.Info(message)
}
func (l *Logger) Infof(message string, arg []interface{}) {
	l.log.Infof(message, arg)
}
func (l *Logger) Close() {
	l.file.Close()
}
