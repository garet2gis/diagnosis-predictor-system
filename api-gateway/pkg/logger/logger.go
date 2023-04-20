package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

type LoggerConfig struct {
	ReportCaller    bool
	IsJSONFormatter bool
	LoggingLevel    string
}

func NewLogger(cfg LoggerConfig) *logrus.Logger {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetReportCaller(cfg.ReportCaller)

	if cfg.IsJSONFormatter {
		l.SetFormatter(&logrus.JSONFormatter{})
	} else {
		l.SetFormatter(&logrus.TextFormatter{DisableColors: true, FullTimestamp: true})
	}

	l = SetLoggerLevel(l, cfg.LoggingLevel)

	return l
}

func SetLoggerLevel(l *logrus.Logger, level string) *logrus.Logger {
	switch level {
	case "panic":
		l.SetLevel(logrus.PanicLevel)
	case "fatal":
		l.SetLevel(logrus.FatalLevel)
	case "error":
		l.SetLevel(logrus.ErrorLevel)
	case "warning":
		l.SetLevel(logrus.WarnLevel)
	case "info":
		l.SetLevel(logrus.InfoLevel)
	case "debug":
		l.SetLevel(logrus.DebugLevel)
	case "trace":
		l.SetLevel(logrus.TraceLevel)
	default:
		l.SetLevel(logrus.InfoLevel)
	}

	return l
}
