package slack

import (
	"fmt"
)

// Logger is a logger interface compatible with both stdlib and some 3rd party
// loggers.
type Logger interface {
	Output(int, string) error
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})
	Debugf(format string, v ...interface{})
	Debugln(v ...interface{})
}

// InternalLog implements the additional methods used by our internal logging
// ontop of Output().
type InternalLog struct {
	Logger
}

// Println replicates the behaviour of the standard logger.
func (t InternalLog) Println(v ...interface{}) {
	t.Output(2, fmt.Sprintln(v...))
}

// Printf replicates the behaviour of the standard logger.
func (t InternalLog) Printf(format string, v ...interface{}) {
	t.Output(2, fmt.Sprintf(format, v...))
}

// Print replicates the behaviour of the standard logger.
func (t InternalLog) Print(v ...interface{}) {
	t.Output(2, fmt.Sprint(v...))
}

// Debugf print a formatted debug line.
func (t InternalLog) Debugf(format string, v ...interface{}) {
	t.Output(2, fmt.Sprint(v...))
}

// Debugln print a debug line.
func (t InternalLog) Debugln(v ...interface{}) {
	t.Output(2, fmt.Sprint(v...))
}

type discard struct{}

func (d discard) Output(int, string) error {
	return nil
}
func (d discard) Print(...interface{})                   {}
func (d discard) Printf(string, ...interface{})          {}
func (d discard) Println(...interface{})                 {}
func (d discard) Debugf(format string, v ...interface{}) {}
func (d discard) Debugln(v ...interface{})               {}
