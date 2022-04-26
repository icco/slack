package main

import (
	"fmt"
	"log"

	"github.com/slack-go/slack"
)

// internalLog implements the additional methods used by our internal logging
// ontop of Output().
type internalLog struct {
	log *log.Logger

	slack.Logger
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
