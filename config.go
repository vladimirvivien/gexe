package echo

// Config stores configuration
type Config struct {
	panicOnErr bool
	verbose    bool
}

// SetPanicOnErr panics program on any error
func (c *Config) SetPanicOnErr(val bool) *Config {
	c.panicOnErr = val
	return c
}

// IsPanicOnErr returns panic-on-error flag
func (c *Config) IsPanicOnErr() bool {
	return c.panicOnErr
}

// SetVerbose sets verbosity
func (c *Config) SetVerbose(val bool) *Config {
	c.verbose = val
	return c
}

// IsVerbose returns verbosity flag
func (c *Config) IsVerbose() bool {
	return c.verbose
}
