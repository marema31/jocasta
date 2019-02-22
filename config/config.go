package config

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/spf13/viper"
)

// Default value
var defaults = map[string]interface{}{
	"out_file":    "/tmp/jocasta_{{.App}}_stdout.log",
	"out_maxsize": "0",
	"out_backups": "0",
	"err_file":    "/tmp/jocasta_{{.App}}_stderr.log",
	"err_maxsize": "0",
	"err_backups": "0",
}

// Config implements the config store of jocasta
type Config struct {
	v   *viper.Viper
	App string
}

// Params type for characteristics of a stream
type Params struct {
	File    string
	Maxsize uint
	Backups int
}

// New initialize the config store
func New(path string, filename string, app string) (*Config, error) {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	v.SetConfigName(filename) // The file will be named [filename].json, [filename].yaml or [filename.toml]
	v.AddConfigPath(path)
	v.SetEnvPrefix("jocasta")
	v.AutomaticEnv()
	err := v.ReadInConfig()

	config := &Config{v: v, App: app}
	return config, err
}

func keyName(key, subkey string) (string, error) {
	switch key {
	case "out", "err":
		return fmt.Sprintf("%s_%s", key, subkey), nil
	default:
		return "", fmt.Errorf("don't know anything about %s", key)
	}

}

// File return the filename for logs for given stream
func (c *Config) File(stream string) (string, error) {
	key, err := keyName(stream, "file")
	if err != nil {
		return "", err
	}

	t, err := template.New("filename").Parse(c.v.GetString(key))
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, c); err != nil {
		return "", err
	}

	return tpl.String(), nil
}

// MaxSize return the max size of log file before rotation for given stream
func (c *Config) MaxSize(stream string) (uint, error) {
	key, err := keyName(stream, "maxsize")
	if err != nil {
		return 0, err
	}

	return c.v.GetSizeInBytes(key), nil
}

// Backups return the number of historical files for logs for given stream
func (c *Config) Backups(stream string) (int, error) {
	key, err := keyName(stream, "backups")
	if err != nil {
		return 0, err
	}
	return c.v.GetInt(key), nil
}

// GetParams return the whole logs info for given stream in Params type
func (c *Config) GetParams(stream string) (*Params, error) {

	maxsize, err := c.MaxSize(stream)
	if err != nil {
		return nil, err
	}
	backups, err := c.Backups(stream)
	if err != nil {
		return nil, err
	}

	file, err := c.File(stream)
	if err != nil {
		return nil, err
	}
	p := &Params{
		Maxsize: maxsize,
		Backups: backups,
		File:    file,
	}
	return p, nil
}
