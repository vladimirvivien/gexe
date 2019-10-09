package echo

import (
	"fmt"
	"regexp"
)

// Echo represents a new Echo session
type Echo struct {
	vars map[string]string
	Conf *conf
}

var (
	spaceRgx = regexp.MustCompile("\\s")
	lineRgx  = regexp.MustCompile("\\n")
)

// New creates a new Echo session
func New() *Echo {
	e := &Echo{vars: make(map[string]string)}
	e.Conf = &conf{e: e}
	return e
}

func (e *Echo) shouldPanic(msg string) {
	if e.Conf.isPanicOnErr() {
		panic(msg)
	}
}

func (e *Echo) shouldLog(msg string) {
	if e.Conf.isVerbose() {
		fmt.Println(msg)
	}
}

func (e *Echo) String() string {
	return fmt.Sprintf("Vars[%#v]", e.vars)
}

type conf struct {
	e          *Echo
	panicOnErr bool
	verbose    bool
}

func (c *conf) SetPanicOnErr(val bool) *Echo {
	c.panicOnErr = val
	return c.e
}

func (c *conf) isPanicOnErr() bool {
	return c.panicOnErr
}

func (c *conf) SetVerbose(val bool) *Echo {
	c.verbose = val
	return c.e
}

func (c *conf) isVerbose() bool {
	return c.verbose
}
