package clog

import (
	"os"

	"github.com/op/go-logging"
)

var logLevels = map[string]logging.Level{
	"CRITICAL": logging.CRITICAL,
	"ERROR":    logging.ERROR,
	"WARNING":  logging.WARNING,
	"NOTICE":   logging.NOTICE,
	"INFO":     logging.INFO,
	"DEBUG":    logging.DEBUG,
}

var logger = logging.MustGetLogger("clog")

// Init initializes the custom logger sub-system
func Init(name, level string, debug bool) {

	logger = logging.MustGetLogger(name)
	backends := make([]logging.Backend, 0, 1)

	consoleFmt := logging.MustStringFormatter(
		`%{color} ▶ [%{level:.4s} %{id:05x}%{color:reset}] %{message}`,
	)
	console := logging.NewLogBackend(os.Stdout, "", 0)
	consoleFormat := logging.NewBackendFormatter(console, consoleFmt)
	consoleLevel := logging.AddModuleLevel(consoleFormat)
	if debug {
		consoleLevel.SetLevel(logging.DEBUG, "")
	} else {
		consoleLevel.SetLevel(logLevels[level], "")
	}
	backends = append(backends, consoleLevel)

	logging.SetBackend(backends...)
}

// Critical logs a simple message when severity is set to CRITICAL or above
func Critical(args ...interface{}) {
	logger.Critical(args...)
}

// Criticalf logs a formatted message when severity is set to CRITICAL or above
func Criticalf(format string, args ...interface{}) {
	logger.Criticalf(format, args...)
}

// Error logs a simple message when severity is set to ERROR or above
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Errorf logs a formatted message when severity is set to ERROR or above
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Warning logs a simple message when severity is set to WARNING or above
func Warning(args ...interface{}) {
	logger.Warning(args...)
}

// Warningf logs a formatted message when severity is set to WARNING or above
func Warningf(format string, args ...interface{}) {
	logger.Warningf(format, args...)
}

// Notice logs a simple message when severity is set to NOTICE or above
func Notice(args ...interface{}) {
	logger.Notice(args...)
}

// Noticef logs a formatted message when severity is set to NOTICE or above
func Noticef(format string, args ...interface{}) {
	logger.Noticef(format, args...)
}

// Info logs a simple message when severity is set to INFO or above
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof logs a formatted message when severity is set to INFO or above
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Debug logs a simple message when severity is set to DEBUG or above
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Debugf logs a formatted message when severity is set to DEBUG or above
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Fatal logs a simple message which is fatal
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Fatalf logs a formatted message which is fatal
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

// Panic logs a simple message which leads to panic
func Panic(args ...interface{}) {
	logger.Panic(args...)
}

// Panicf logs a formatted message which leads to panic
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}
