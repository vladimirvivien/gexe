package echo

import (
	"fmt"
	"regexp"
)

// Echo represents a new Echo session
type Echo struct {
	vars map[string]string // session var
	Conf *Config           // session config
	Proc *ProcResult       // last executed proc
}

var (
	spaceRgx = regexp.MustCompile("\\s")
	lineRgx  = regexp.MustCompile("\\n")
)

// New creates a new Echo session
func New() *Echo {
	e := &Echo{vars: make(map[string]string)}
	e.Conf = new(Config)
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
