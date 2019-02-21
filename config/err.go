package config

// Errfile return the filename for stdErr logs
func (c *Config) ErrFile() string {
	return c.v.GetString("err_file")
}

// ErrMaxSize return the max size of stdstderr log file before rotation
func (c *Config) ErrMaxSize() (int, error) {
	size, err := sizeFromString(c.v.GetString("err_maxsize"))
	if err != nil {
		return 0, err
	}

	return size, nil
}

// ErrBackups return the number of historical files for stdstderr logs
func (c *Config) ErrBackups() int {
	return c.v.GetInt("err_backups")
}
