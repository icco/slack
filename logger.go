package slack

import (
	"fmt"
)

// Logger is a logger interface compatible with both stdlib and some
// 3rd party loggers.
type Logger interface {
	Output(int, string) error
}

// ilogger represents the internal logging api we use.
type ilogger interface {
	Logger

	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})

	// Debugf print a formatted debug line.
	Debugf(format string, v ...interface{})

	// Debugln print a debug line.
	Debugln(v ...interface{})
}

// internalLog implements the additional methods used by our internal logging.
type internalLog struct {
	Logger
}

// Println replicates the behaviour of the standard logger.
func (t internalLog) Println(v ...interface{}) {
	t.Output(2, fmt.Sprintln(v...))
}

// Printf replicates the behaviour of the standard logger.
func (t internalLog) Printf(format string, v ...interface{}) {
	t.Output(2, fmt.Sprintf(format, v...))
}

// Print replicates the behaviour of the standard logger.
func (t internalLog) Print(v ...interface{}) {
	t.Output(2, fmt.Sprint(v...))
}

// Debugf print a formatted debug line.
func (t internalLog) Debugf(format string, v ...interface{}) {
	t.Output(2, fmt.Sprint(v...))
}

// Debugln print a debug line.
func (t internalLog) Debugln(v ...interface{}) {
	t.Output(2, fmt.Sprint(v...))
}
