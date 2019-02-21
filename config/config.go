package config

import (
	"github.com/spf13/viper"
)

// Default value
var defaults = map[string]interface{}{
	"out_file":    "/tmp/jocasta_stdout.log",
	"out_maxsize": "0",
	"out_backups": "0",
	"err_file":    "/tmp/jocasta_stderr.log",
	"err_maxsize": "0",
	"err_backups": "0",
}

// Config implements the config store of jocasta
type Config struct {
	v *viper.Viper
}

// New initialize the config store
func New(path string, filename string) (*Config, error) {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	v.SetConfigName(filename) // The file will be named [filename].json, [filename].yaml or [filename.toml]
	v.AddConfigPath(path)
	v.SetEnvPrefix("jocasta")
	v.AutomaticEnv()
	err := v.ReadInConfig()

	config := &Config{v: v}
	return config, err
}
