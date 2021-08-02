package gexe

import (
	"github.com/vladimirvivien/gexe/prog"
	"github.com/vladimirvivien/gexe/vars"
)

var (
	// DefaultEcho surfaces package-level Echo session
	// used for all package functions
	DefaultEcho = New()
)

// Echo represents a new Echo session
type Echo struct {
	err  error
	vars *vars.Variables // session vars
	prog *prog.Info
	Conf *Config // session config
}

// New creates a new Echo session
func New() *Echo {
	e := &Echo{
		vars: vars.New(),
		prog: prog.Prog(),
		Conf: &Config{escapeChar: '\\'},
	}
	return e
}
