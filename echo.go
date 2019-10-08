package echo

import (
	"regexp"
)

type echo struct {
	vars map[string]string
	Conf *conf
}

var (
	spaceRgx = regexp.MustCompile("\\s")
	lineRgx  = regexp.MustCompile("\\n")
)

// New creates a new session
func New() *echo {
	e := &echo{vars: make(map[string]string)}
	e.Conf = &conf{e: e}
	return e
}

func (e *echo) shouldPanic(msg string) {
	if e.Conf.IsPanicOnErr() {
		panic(msg)
	}
}

type conf struct {
	e          *echo
	panicOnErr bool
}

func (c *conf) SetPanicOnErr(val bool) *echo {
	c.panicOnErr = val
	return c.e
}

func (c *conf) IsPanicOnErr() bool {
	return c.panicOnErr
}
