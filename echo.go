package echo

import (
	"fmt"
	"regexp"

	"github.com/vladimirvivien/echo/prog"
	"github.com/vladimirvivien/echo/vars"
)

// Echo represents a new Echo session
type Echo struct {
	vars  *vars.Variables // session vars
	Conf  *Config         // session config
	Prog  *prog.prog      // Program info
}

var (
	DefaultEcho = New()
	spaceRgx = regexp.MustCompile("\\s")
	lineRgx  = regexp.MustCompile("\\n")
)

// New creates a new Echo session
func New() *Echo {
	e := &Echo{
		vars: vars.New(),
		Conf: &Config{escapeChar: '\\'},
		Prog: new(prog.prog),
	}
	return e
}

func (e *Echo) shouldPanic(msg string) {
	if e.Conf.IsPanicOnErr() {
		panic(msg)
	}
}

func (e *Echo) shouldLog(msg string) {
	if e.Conf.IsVerbose() {
		fmt.Println(msg)
	}
}

func (e *Echo) String() string {
	return fmt.Sprintf("Vars[%#v]", e.vars)
}
