package logwriter

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/marema31/jocasta/config"
)

func TestWriteWithoutRotation(t *testing.T) {
	f, err := os.Create("testdata/test.log")
	if err != nil {
		t.Errorf("Can not open log file. err=%v", err)
	}

	l := LogWriter{
		params: &config.Params{
			Backups: 0,
			Maxsize: 0,
			File:    "testdata/test.log",
		},
		currentsize: 0,
		file:        f,
	}
	defer l.Close()

	err = writeTestData(&l)
	if err != nil {
		t.Errorf("Can't create test log file. err=%v", err)
	}

	// Verify the file content
	b, err := ioutil.ReadFile("testdata/test.log")
	if err != nil {
		t.Errorf("Can read log file. err=%v", err)
	}
	if !bytes.Equal(b, []byte("12\n34\n56\n78\n90\nab\ncd\n")) {
		t.Errorf("The content of log file not correct. b=%v", b)
	}
	clean()
}

func TestWriteWithRotation(t *testing.T) {
	f, err := os.Create("testdata/test.log")
	if err != nil {
		t.Errorf("Can not open log file. err=%v", err)
	}

	l := LogWriter{
		params: &config.Params{
			Backups: 2,
			Maxsize: 4,
			File:    "testdata/test.log",
		},
		currentsize: 0,
		file:        f,
	}
	defer l.Close()

	err = writeTestData(&l)
	if err != nil {
		t.Errorf("Can't create test log file. err=%v", err)
	}

	// Verify the file content
	b, err := ioutil.ReadFile("testdata/test.log")
	if err != nil {
		t.Errorf("Can read first log file. err=%v", err)
	}
	if !bytes.Equal(b, []byte("cd\n")) {
		t.Errorf("The content of log file not correct. b=%v", b)
	}
	// Verify the file content
	b, err = ioutil.ReadFile("testdata/test.log.1")
	if err != nil {
		t.Errorf("Can read first backup log file. err=%v", err)
	}
	if !bytes.Equal(b, []byte("90\nab\n")) {
		t.Errorf("The content of first backup log file not correct. b=%v", b)
	}

	// Verify the file content
	b, err = ioutil.ReadFile("testdata/test.log.2")
	if err != nil {
		t.Errorf("Can read second backup log file. err=%v", err)
	}
	if !bytes.Equal(b, []byte("56\n78\n")) {
		t.Errorf("The content of second backup log file not correct. b=%v", b)
	}

	// Verify the file content
	if _, err := os.Stat("testdata/test.log.3"); !os.IsNotExist(err) {
		t.Errorf("A third backup file should not been created.")
	}
	clean()
}

func writeTestData(l *LogWriter) error {
	n, err := l.Write([]byte("12\n"))
	if err != nil {
		return fmt.Errorf("Can not write the first line. err=%v", err)
	}
	if n != 3 {
		return fmt.Errorf("Wrong number of byte written for first line. n=%d", n)
	}

	n, err = l.Write([]byte("34\n56\n"))
	if err != nil {
		return fmt.Errorf("Can not write the second and third lines. err=%v", err)
	}
	if n != 6 {
		return fmt.Errorf("Wrong number of byte written for second and third lines. n=%d", n)
	}

	n, err = l.Write([]byte("78\n90\nab\ncd\n"))
	if err != nil {
		return fmt.Errorf("Can not write the remaining lines. err=%v", err)
	}
	if n != 12 {
		return fmt.Errorf("Wrong number of byte written for remaining lines. n=%d", n)
	}

	return nil
}

func clean() {
	for _, ext := range []string{"log", "log.1", "log.2", "log.3"} {
		if _, err := os.Stat("testdata/test." + ext); !os.IsNotExist(err) {
			os.Remove("testdata/test." + ext)
		}
	}
}
