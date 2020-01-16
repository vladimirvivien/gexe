package echo

// Conf stores configuration
type Conf struct {
	panicOnErr bool
	verbose    bool
}

// SetPanicOnErr panics program on any error
func (c *Conf) SetPanicOnErr(val bool) *Conf {
	c.panicOnErr = val
	return c
}

// IsPanicOnErr returns panic-on-error flag
func (c *Conf) IsPanicOnErr() bool {
	return c.panicOnErr
}

// SetVerbose sets verbosity
func (c *Conf) SetVerbose(val bool) *Conf {
	c.verbose = val
	return c
}

// IsVerbose returns verbosity flag
func (c *Conf) IsVerbose() bool {
	return c.verbose
}
