package config

import (
	"testing"
)

func TestIncorrectSize(t *testing.T) {

	str := "10KG"

	size, err := sizeFromString(str)
	if err == nil {
		t.Errorf("No error on invalid size string. Calculated size=%d for '%s'", size, str)
	}

}

func TestCorrectSize(t *testing.T) {

	str := "42"
	waited := 42
	size, err := sizeFromString(str)
	if err != nil {
		t.Errorf("Error while converting for '%s', err=%v", str, err)
	}
	if size != waited {
		t.Errorf("Wrong conversion. Calculated size=%d for '%s', should be %d", size, str, waited)
	}

	str = "42B"
	waited = 42
	size, err = sizeFromString(str)
	if err != nil {
		t.Errorf("Error while converting for '%s', err=%v", str, err)
	}
	if size != waited {
		t.Errorf("Wrong conversion. Calculated size=%d for '%s', should be %d", size, str, waited)
	}

	str = "42KB"
	waited = 42 * 1024
	size, err = sizeFromString(str)
	if err != nil {
		t.Errorf("Error while converting for '%s', err=%v", str, err)
	}
	if size != waited {
		t.Errorf("Wrong conversion. Calculated size=%d for '%s', should be %d", size, str, waited)
	}

	str = "42MB"
	waited = 42 * 1024 * 1024
	size, err = sizeFromString(str)
	if err != nil {
		t.Errorf("Error while converting for '%s', err=%v", str, err)
	}
	if size != waited {
		t.Errorf("Wrong conversion. Calculated size=%d for '%s', should be %d", size, str, waited)
	}

	str = "42GB"
	waited = 42 * 1024 * 1024 * 1024
	size, err = sizeFromString(str)
	if err != nil {
		t.Errorf("Error while converting for '%s', err=%v", str, err)
	}
	if size != waited {
		t.Errorf("Wrong conversion. Calculated size=%d for '%s', should be %d", size, str, waited)
	}

	str = "42TB"
	waited = 42 * 1024 * 1024 * 1024 * 1024
	size, err = sizeFromString(str)
	if err != nil {
		t.Errorf("Error while converting for '%s', err=%v", str, err)
	}
	if size != waited {
		t.Errorf("Wrong conversion. Calculated size=%d for '%s', should be %d", size, str, waited)
	}

}
