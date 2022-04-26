package slack

import (
	"bytes"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// internalLog implements the additional methods used by our internal logging
// ontop of Output().
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

func TestLogging(t *testing.T) {
	buf := bytes.NewBufferString("")
	logger := internalLog{Logger: log.New(buf, "", 0|log.Lshortfile)}
	logger.Println("test line 123")
	assert.Equal(t, buf.String(), "logger_test.go:14: test line 123\n")
	buf.Truncate(0)
	logger.Print("test line 123")
	assert.Equal(t, buf.String(), "logger_test.go:17: test line 123\n")
	buf.Truncate(0)
	logger.Printf("test line 123\n")
	assert.Equal(t, buf.String(), "logger_test.go:20: test line 123\n")
	buf.Truncate(0)
	logger.Output(1, "test line 123\n")
	assert.Equal(t, buf.String(), "logger_test.go:23: test line 123\n")
	buf.Truncate(0)
}
