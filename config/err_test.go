package config

import (
	"os"
	"testing"
)

func TestErrDefaultValue(t *testing.T) {
	defaultFile := "/tmp/jocasta_stderr.log"
	defaultSize := 0
	defaultBackups := 0

	c, err := New("testdata", "empty")
	if err != nil {
		t.Errorf("Can not read the configuration file. err=%v", err)
	}

	maxsize, err := c.ErrMaxSize()
	if err != nil {
		t.Errorf("Can not retrieve err.maxsize. err=%v", err)
	}
	if maxsize != defaultSize {
		t.Errorf("Default value for err.maxsize is not correct, should be %d and I receive %d", defaultSize, maxsize)
	}

	backups := c.ErrBackups()
	if backups != defaultBackups {
		t.Errorf("Default value for err.backups is not correct, should be %d and I receive %d", defaultBackups, backups)
	}

	file := c.ErrFile()
	if file != defaultFile {
		t.Errorf("Default value for err.file is not correct, should be %s and I receive %s", defaultFile, file)
	}
}

func TestErrFromFile(t *testing.T) {
	defaultFile := "/tmp/correct_err.log"
	defaultSize := 2 * 1024 * 1024
	defaultBackups := 3

	c, err := New("testdata", "correct")
	if err != nil {
		t.Errorf("Can not read the configuration file. err=%v", err)
	}

	maxsize, err := c.ErrMaxSize()
	if err != nil {
		t.Errorf("Can not retrieve err.maxsize. err=%v", err)
	}
	if maxsize != defaultSize {
		t.Errorf("Retrieved value for err.maxsize is not correct, should be %d and I receive %d", defaultSize, maxsize)
	}

	backups := c.ErrBackups()
	if backups != defaultBackups {
		t.Errorf("Retrieved value for err.backups is not correct, should be %d and I receive %d", defaultBackups, backups)
	}

	file := c.ErrFile()
	if file != defaultFile {
		t.Errorf("Retrieved value for err.file is not correct, should be %s and I receive %s", defaultFile, file)
	}
}

func TestIncorrectErrSize(t *testing.T) {

	str := "10KG"

	os.Setenv("JOCASTA_ERR_MAXSIZE", str)
	c, err := New("testdata", "empty")
	if err != nil {
		t.Errorf("Can not read the configuration file. err=%v", err)
	}

	maxsize, err := c.ErrMaxSize()
	if err == nil {
		t.Errorf("No error on invalid err.maxsize string. Calculated size=%d for '%s'", maxsize, str)
	}

}
