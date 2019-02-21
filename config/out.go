package config

// OutFile return the filename for stdout logs
func (c *Config) OutFile() string {
	return c.v.GetString("out_file")
}

// OutMaxSize return the max size of stdout log file before rotation
func (c *Config) OutMaxSize() (int, error) {
	size, err := sizeFromString(c.v.GetString("out_maxsize"))
	if err != nil {
		return 0, err
	}

	return size, nil
}

// OutBackups return the number of historical files for stdout logs
func (c *Config) OutBackups() int {
	return c.v.GetInt("out_backups")
}
