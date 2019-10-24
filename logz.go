package logz

import (
	"fmt"
	"io"
	"os"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)
// Logger is a basic logging interface for a structured logger.
type Logger interface {
	// Info writes Info level logs of key/value pairs
	Info(keyvals ...interface{})
	// InfoWithMessage writes Info level logs of key/value pairs. The msg is written with a key of "msg".
	InfoWithMessage(msg string, keyvals ...interface{})

	// Warn writes Warn level logs of key/value pairs
	Warn(keyvals ...interface{})
	// WarnWithMessage writes Warn level logs of key/value pairs. The msg is written with a key of "msg".
	WarnWithMessage(msg string, keyvals ...interface{})

	// Error writes Error level logs of key/value pairs. err is written with a key of "err".
	Error(err error, keyvals ...interface{})
	// ErrorWithMessage writes Error level logs of key/value pairs. err is written with a key of "err". The msg is written with a key of "msg".
	ErrorWithMessage(err error, msg string, keyvals ...interface{})

	// Debug writes Debug level logs of key/value pairs
	Debug(keyvals ...interface{})
	// DebugWithMessage writes Debug level logs of key/value pairs. The msg is written with a key of "msg".
	DebugWithMessage(msg string, keyvals ...interface{})

	// FatalError writes an Error level message and the calls os.Exit(1) if err is not nil. The "fataL" key value will be "true".
	FatalError(err error, keyvals ... interface{})
}

type logger struct {
	logger kitlog.Logger
}

func (l *logger) Log(keyvals ...interface{}) {
	_ = l.logger.Log(keyvals...)
}

func (l *logger) Info(keyvals ...interface{}) {
	_ = level.Info(l.logger).Log(keyvals...)
}

func (l *logger) InfoWithMessage(msg string, keyvals ...interface{}) {
	keyvals = append(keyvals, "msg", msg)
	_ = level.Info(l.logger).Log(keyvals...)
}

func (l *logger) Debug(keyvals ...interface{}) {
	_ = level.Debug(l.logger).Log(keyvals...)
}

func (l *logger) DebugWithMessage(msg string, keyvals ...interface{}) {
	keyvals = append(keyvals, "msg", msg)
	_ = level.Debug(l.logger).Log(keyvals...)
}

func (l *logger) Warn(keyvals ...interface{}) {
	_ = level.Warn(l.logger).Log(keyvals...)
}

func (l *logger) WarnWithMessage(msg string, keyvals ...interface{}) {
	keyvals = append(keyvals, "msg", msg)
	_ = level.Warn(l.logger).Log(keyvals...)
}

func (l *logger) Error(err error, keyvals ...interface{}) {
	if err == nil {
		return
	}

	keyvals = append(keyvals, "err", fmt.Sprintf("%+v", err))
	e := level.Error(l.logger).Log(keyvals...)

	if e != nil {
		panic(e) // Got error logging error, something seriously broken, panic and get out of here
	}
}

func (l *logger) ErrorWithMessage(err error, msg string, keyvals ...interface{}) {
	if err == nil {
		return
	}

	keyvals = append(keyvals, "err", fmt.Sprintf("%+v", err), "msg", msg)
	e := level.Error(l.logger).Log(keyvals...)

	if e != nil {
		panic(e) // Got error logging error, something seriously broken, panic and get out of here
	}
}

func (l *logger) FatalError(err error, keyvals ...interface{}) {
	if err == nil {
		return
	}

	keyvals = append(keyvals, "err", fmt.Sprintf("%+v", err), "fatal", true)
	e := level.Error(l.logger).Log(keyvals...)

	if e != nil {
		panic(e) // Got error logging error, something seriously broken, panic and get out of here
	}

	os.Exit(1)
}

// NewLogger creates a new instance of Logger using the specified Level as the lowest level that will be logged.
func NewLogger(lvl Level) Logger {

	l := new(logger)

	var writer io.Writer

	writer = kitlog.NewSyncWriter(os.Stdout)

	opt := lvl.Option()

	l.logger = kitlog.NewJSONLogger(writer)
	l.logger = level.NewFilter(l.logger, opt)
	l.logger = kitlog.With(l.logger, "ts", kitlog.DefaultTimestampUTC, "caller", kitlog.Caller(4))

	return l
}

// NewNullLogger returns a no op logger. Useful if your testing code that has a Logger dependency.
func NewNullLogger() Logger {
	return &logger {
		logger: kitlog.NewNopLogger(),
	}
}

// With wrapps theLogger with a new logger that will always include keyvals in its output.
func With(theLogger Logger, keyvals ...interface{}) Logger {

	l, ok := theLogger.(*logger)

	if !ok {
		return theLogger
	}

	newLogger := new(logger)

	newLogger.logger = kitlog.With(l.logger, keyvals...)

	return newLogger
}
