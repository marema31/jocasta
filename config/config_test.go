package config

import (
	"os"
	"testing"
)

func TestOutDefaultValue(t *testing.T) {
	defaultFile := "/tmp/jocasta_dummy_stdout.log"
	var defaultSize uint
	defaultBackups := 0

	c, err := New("testdata", "empty", "dummy")
	if err != nil {
		t.Errorf("Can not read the configuration file. err=%v", err)
	}

	maxsize, err := c.MaxSize("out")
	if err != nil {
		t.Errorf("Error while retrieving out_maxsize. err=%v", err)
	}
	if maxsize != defaultSize {
		t.Errorf("Default value for out.maxsize is not correct, should be %d and I receive %d", defaultSize, maxsize)
	}

	backups, err := c.Backups("out")
	if err != nil {
		t.Errorf("Error while retrieving out_backups. err=%v", err)
	}
	if backups != defaultBackups {
		t.Errorf("Default value for out.backups is not correct, should be %d and I receive %d", defaultBackups, backups)
	}

	file, err := c.File("out")
	if err != nil {
		t.Errorf("Error while retrieving out_file. err=%v", err)
	}
	if file != defaultFile {
		t.Errorf("Default value for out.file is not correct, should be %s and I receive %s", defaultFile, file)
	}
}

func TestOutRetrievedFromFile(t *testing.T) {
	defaultFile := "/tmp/correct.log"
	var defaultSize uint = 1024 * 1024
	defaultBackups := 2

	c, err := New("testdata", "correct", "dummy")
	if err != nil {
		t.Errorf("Can not read the configuration file. err=%v", err)
	}

	maxsize, err := c.MaxSize("out")
	if err != nil {
		t.Errorf("Error while retrieving out_maxsize. err=%v", err)
	}
	if maxsize != defaultSize {
		t.Errorf("Retrieved value for out.maxsize is not correct, should be %d and I receive %d", defaultSize, maxsize)
	}

	backups, err := c.Backups("out")
	if err != nil {
		t.Errorf("Error while retrieving out_backups. err=%v", err)
	}
	if backups != defaultBackups {
		t.Errorf("Retrieved value for out.backups is not correct, should be %d and I receive %d", defaultBackups, backups)
	}

	file, err := c.File("out")
	if err != nil {
		t.Errorf("Error while retrieving out_file. err=%v", err)
	}
	if file != defaultFile {
		t.Errorf("Retrieved value for out.file is not correct, should be %s and I receive %s", defaultFile, file)
	}
}

func TestErrDefaultValue(t *testing.T) {
	defaultFile := "/tmp/jocasta_dummy_stderr.log"
	var defaultSize uint
	defaultBackups := 0

	c, err := New("testdata", "empty", "dummy")
	if err != nil {
		t.Errorf("Can not read the configuration file. err=%v", err)
	}

	maxsize, err := c.MaxSize("err")
	if err != nil {
		t.Errorf("Error while retrieving err_maxsize. err=%v", err)
	}
	if maxsize != defaultSize {
		t.Errorf("Default value for err.maxsize is not correct, should be %d and I receive %d", defaultSize, maxsize)
	}

	backups, err := c.Backups("err")
	if err != nil {
		t.Errorf("Error while retrieving err_backups. err=%v", err)
	}
	if backups != defaultBackups {
		t.Errorf("Default value for err.backups is not correct, should be %d and I receive %d", defaultBackups, backups)
	}

	file, err := c.File("err")
	if err != nil {
		t.Errorf("Error while retrieving err_file. err=%v", err)
	}
	if file != defaultFile {
		t.Errorf("Default value for err.file is not correct, should be %s and I receive %s", defaultFile, file)
	}
}

func TestErrFromFile(t *testing.T) {
	defaultFile := "/tmp/correct_err.log"
	var defaultSize uint = 2 * 1024 * 1024
	defaultBackups := 3

	c, err := New("testdata", "correct", "dummy")
	if err != nil {
		t.Errorf("Can not read the configuration file. err=%v", err)
	}

	maxsize, err := c.MaxSize("err")
	if err != nil {
		t.Errorf("Error while retrieving err_maxsize. err=%v", err)
	}
	if maxsize != defaultSize {
		t.Errorf("Retrieved value for err.maxsize is not correct, should be %d and I receive %d", defaultSize, maxsize)
	}

	backups, err := c.Backups("err")
	if err != nil {
		t.Errorf("Error while retrieving err_backups. err=%v", err)
	}
	if backups != defaultBackups {
		t.Errorf("Retrieved value for err.backups is not correct, should be %d and I receive %d", defaultBackups, backups)
	}

	file, err := c.File("err")
	if err != nil {
		t.Errorf("Error while retrieving err_file. err=%v", err)
	}
	if file != defaultFile {
		t.Errorf("Retrieved value for err.file is not correct, should be %s and I receive %s", defaultFile, file)
	}
}

func TestParams(t *testing.T) {
	defaultFile := "/tmp/correct_err.log"
	defaultSize := 2 * 1024 * 1024
	defaultBackups := 3

	c, err := New("testdata", "correct", "dummy")
	if err != nil {
		t.Errorf("Can not read the configuration file. err=%v", err)
	}

	p, err := c.GetParams("err")
	if err != nil {
		t.Errorf("Error while retrieving err_maxsize. err=%v", err)
	}
	if p.Maxsize != defaultSize {
		t.Errorf("Retrieved value for err.maxsize is not correct, should be %d and I receive %d", defaultSize, p.Maxsize)
	}

	if p.Backups != defaultBackups {
		t.Errorf("Retrieved value for err.backups is not correct, should be %d and I receive %d", defaultBackups, p.Backups)
	}

	if p.File != defaultFile {
		t.Errorf("Retrieved value for err.file is not correct, should be %s and I receive %s", defaultFile, p.File)
	}
}

func TestPIncorrectStreamParams(t *testing.T) {
	c, err := New("testdata", "correct", "dummy")
	if err != nil {
		t.Errorf("Can not read the configuration file. err=%v", err)
	}

	_, err = c.GetParams("dummy")
	if err == nil {
		t.Errorf("Should not be able to retrieve params for dummy.")
	}

}
func TestEnv(t *testing.T) {

	str := "dummy.log"

	os.Setenv("JOCASTA_OUT_FILE", str)

	c, err := New("testdata", "empty", "dummy")
	if err != nil {
		t.Errorf("Can not read the configuration file. err=%v", err)
	}

	file, err := c.File("out")
	if err != nil {
		t.Errorf("Error while retrieving out_file. err=%v", err)
	}
	if file != str {
		t.Errorf("Not able to retrieve out_file from environment variable . Found=%s for '%s'", file, str)
	}

}

func TestIncorrectTemplateInFile(t *testing.T) {

	str := "dummy{{.toto}}.log"

	os.Setenv("JOCASTA_OUT_FILE", str)

	c, err := New("testdata", "empty", "dummy")
	if err != nil {
		t.Errorf("Can not read the configuration file. err=%v", err)
	}

	_, err = c.File("out")
	if err == nil {
		t.Errorf("Should not be able to parse the template")
	}

	_, err = c.GetParams("out")
	if err == nil {
		t.Errorf("Should not be able to parse the template")
	}

}

func TestWrongStream(t *testing.T) {

	c, err := New("testdata", "empty", "dummy")
	if err != nil {
		t.Errorf("Can not read the configuration file. err=%v", err)
	}

	_, err = c.MaxSize("dummy")
	if err == nil {
		t.Errorf("No error while retrieving dummy_file. err=%v", err)
	}

	_, err = c.Backups("dummy")
	if err == nil {
		t.Errorf("No error while retrieving dummy_file. err=%v", err)
	}

	_, err = c.File("dummy")
	if err == nil {
		t.Errorf("No error while retrieving dummy_file. err=%v", err)
	}

	_, err = c.GetParams("dummy")
	if err == nil {
		t.Errorf("No error while retrieving dummy params. err=%v", err)
	}
}
