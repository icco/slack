package slack

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

type discard struct{}

func (d discard) Output(int, string) error {
	return nil
}
func (d discard) Print(...interface{})                   {}
func (d discard) Printf(string, ...interface{})          {}
func (d discard) Println(...interface{})                 {}
func (d discard) Debugf(format string, v ...interface{}) {}
func (d discard) Debugln(v ...interface{})               {}
