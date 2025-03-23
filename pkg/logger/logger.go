package logger

import (
	"log"
	"os"

	"go.uber.org/zap"
)

type Logger interface {
	Info(msg string)
	All(msgs ...string)
	Warn(msg string)
	Error(err error)
	Fatal(msg string)
	FatalError(err error)
	ErrorIf(err error)
}

type Log struct {
	logger *zap.Logger
}

func (l *Log) Info(msg string) {
	l.logger.Info(msg)
}

func (l *Log) All(msgs ...string) {
	for _, m := range msgs {
		l.logger.Info(m)
	}
}

func (l *Log) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l *Log) Error(err error) {
	l.logger.Error(err.Error())
}

func (l *Log) Fatal(msg string) {
	l.logger.Fatal(msg)
}

func (l *Log) FatalError(err error) {
	l.logger.Fatal(err.Error())
}

func (l *Log) ErrorIf(err error) {
	if err != nil {
		l.logger.Error(err.Error())
	}
}

func NewLogger(environment string) Logger {
	var zapLogger *zap.Logger
	var err error

	if environment == "production" {
		zapLogger, err = zap.NewProduction()
	} else {
		zapLogger, err = zap.NewDevelopment()
	}

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer zapLogger.Sync()

	log := Log{
		logger: zapLogger,
	}
	return &log
}
