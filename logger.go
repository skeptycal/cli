package cli

import (
	logrus "github.com/sirupsen/logrus"
)

// type Logger = logrus.Logger

var DefaultLogLevel logrus.Level = logrus.DebugLevel

var Log *logrus.Logger = NewLogger()

// NewLogger creates a new logrus logger that is compatible with
// the go log package and has terminal stderr defaults set.
//
// Configuration should be set by changing `Formatter`, `Out` and
// `Hooks` directly on the default logger instance. You can also
// just instantiate your own directly:
//
//    import logrus "github.com/sirupsen/logrus"
//
//    var log = &logrus.Logger{
//      Out: os.Stderr,
//      Formatter: new(logrus.TextFormatter),
//      Hooks: make(logrus.LevelHooks),
//      Level: logrus.DebugLevel,
//    }
//
// It's recommended to make this a global instance called `log`.
func NewLogger() *logrus.Logger {
	return &logrus.Logger{
		Out:       NewStderr(),
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     DefaultLogLevel,
	}
}
