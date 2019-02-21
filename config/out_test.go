package config

import (
	"os"
	"testing"
)

func TestOutDefaultValue(t *testing.T) {
	defaultFile := "/tmp/jocasta_stdout.log"
	defaultSize := 0
	defaultBackups := 0

	c, err := New("testdata", "empty")
	if err != nil {
		t.Errorf("Can not read the configuration file. err=%v", err)
	}

	maxsize, err := c.OutMaxSize()
	if err != nil {
		t.Errorf("Can not retrieve out.maxsize. err=%v", err)
	}
	if maxsize != defaultSize {
		t.Errorf("Default value for out.maxsize is not correct, should be %d and I receive %d", defaultSize, maxsize)
	}

	backups := c.OutBackups()
	if backups != defaultBackups {
		t.Errorf("Default value for out.backups is not correct, should be %d and I receive %d", defaultBackups, backups)
	}

	file := c.OutFile()
	if file != defaultFile {
		t.Errorf("Default value for out.file is not correct, should be %s and I receive %s", defaultFile, file)
	}
}

func TestOutRetrievedFromFile(t *testing.T) {
	defaultFile := "/tmp/correct.log"
	defaultSize := 1024 * 1024
	defaultBackups := 2

	c, err := New("testdata", "correct")
	if err != nil {
		t.Errorf("Can not read the configuration file. err=%v", err)
	}

	maxsize, err := c.OutMaxSize()
	if err != nil {
		t.Errorf("Can not retrieve out.maxsize. err=%v", err)
	}
	if maxsize != defaultSize {
		t.Errorf("Retrieved value for out.maxsize is not correct, should be %d and I receive %d", defaultSize, maxsize)
	}

	backups := c.OutBackups()
	if backups != defaultBackups {
		t.Errorf("Retrieved value for out.backups is not correct, should be %d and I receive %d", defaultBackups, backups)
	}

	file := c.OutFile()
	if file != defaultFile {
		t.Errorf("Retrieved value for out.file is not correct, should be %s and I receive %s", defaultFile, file)
	}
}

func TestOutEnv(t *testing.T) {

	str := "dummy.log"

	os.Setenv("JOCASTA_OUT_FILE", str)

	c, err := New("testdata", "empty")
	if err != nil {
		t.Errorf("Can not read the configuration file. err=%v", err)
	}

	file := c.OutFile()
	if file != str {
		t.Errorf("Not able to retrieve out.file from environment variable . Found=%s for '%s'", file, str)
	}

}

func TestIncorrectOutSize(t *testing.T) {

	str := "10KG"

	os.Setenv("JOCASTA_OUT_MAXSIZE", str)
	c, err := New("testdata", "empty")
	if err != nil {
		t.Errorf("Can not read the configuration file. err=%v", err)
	}

	maxsize, err := c.OutMaxSize()
	if err == nil {
		t.Errorf("No error on invalid out.maxsize string. Calculated size=%d for '%s'", maxsize, str)
	}

}
