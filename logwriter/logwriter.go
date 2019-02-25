package logwriter

import (
	"bytes"
	"fmt"
	"os"

	"github.com/marema31/jocasta/config"
)

// LogWriter object used to save log to a file with limits and automatic rotation
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
	fmt.Printf("Will log std%s on %s with limit:%d and backups:%d\n", stream, p.File, p.Maxsize, p.Backups)
	return l, nil
}

// Write implements the io.Writer interface
// Read the buffer line by line and save it to the file opened in New
func (l *LogWriter) Write(p []byte) (int, error) {

	n := 0
	for n < len(p) {

		// Manage the limitation
		if l.params.Backups > 0 && l.currentsize > l.params.Maxsize {
			err := l.rotation()
			if err != nil {
				return n, err
			}
			l.currentsize = 0
		} else if l.params.Backups < 1 {
			l.currentsize = 0
		}

		// Write line by line
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

func (l *LogWriter) rotation() error {
	err := l.file.Close()
	if err != nil {
		return err
	}

	for copy := l.params.Backups - 1; copy > 0; copy-- {
		src := fmt.Sprintf("%s.%d", l.params.File, copy)
		dst := fmt.Sprintf("%s.%d", l.params.File, copy+1)
		if _, err := os.Stat(src); err == nil {
			os.Rename(src, dst)
			if err != nil {
				return err
			}
		}
	}
	err = os.Rename(l.params.File, l.params.File+".1")
	if err != nil {
		return err
	}

	l.file, err = os.Create(l.params.File)
	if err != nil {
		return err
	}
	return nil
}
