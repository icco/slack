package slack

import (
	"fmt"
	"log"
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

// internalLog implements the additional methods used by our internal logging
// ontop of Output().
type internalLog struct {
	log *log.Logger

	Logger
}

func (t internalLog) Output(stackLevel int, msg string) error {
	return t.log.Output(stackLevel, msg)
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

type discard struct{}

func (d discard) Output(int, string) error {
	return nil
}
func (d discard) Print(...interface{})                   {}
func (d discard) Printf(string, ...interface{})          {}
func (d discard) Println(...interface{})                 {}
func (d discard) Debugf(format string, v ...interface{}) {}
func (d discard) Debugln(v ...interface{})               {}
