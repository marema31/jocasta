package logwriter

import (
	"bytes"
	"os"

	"github.com/marema31/jocasta/config"
)

// LogSaver object used to save log to a file with limits and automatic rotation
type LogWriter struct {
	params      *config.Params
	currentsize int
	file        *os.File
}

// New instanciate a LogWriter object
func New(stream string, c *config.Config) (*LogWriter, error) {
	p, err := c.GetParams(stream)
	if err != nil {
		return nil, err
	}

	f, err := os.Create(p.File)
	if err != nil {
		return nil, err
	}

	l := &LogWriter{
		params:      p,
		currentsize: 0,
		file:        f,
	}

	return l, nil
}

// Write implements the io.Writer interface
// Read the buffer line by line and save it to the file opened in New
func (l *LogWriter) Write(p []byte) (int, error) {

	n := 0
	for n < len(p) {
		i := bytes.IndexByte(p[n:], byte('\n'))
		if i == -1 {
			l.file.Write(p[n:])
			n = len(p)
			l.currentsize += len(p) - n
		} else {
			l.file.Write(p[n : n+i+1])
			n += i + 1
			l.currentsize += i + 1
		}
	}
	return n, nil
}

// Close implements the io.Writer interface
// Close the file opened in New
func (l *LogWriter) Close() error {
	return l.file.Close()
}
