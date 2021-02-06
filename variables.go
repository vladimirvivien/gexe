package echo

import (
	"github.com/vladimirvivien/echo/vars"
)
func (e *Echo) Variables() *vars.Variables {
	return e.vars
}

// Envs declares process environment variables using
// a multi-line space-separated list:
// e.Envs(GOOS=linux GOARCH=amd64)
func (e *Echo) Envs(val string) *Echo {
	vars := e.vars.Envs(val)
	if vars.Err() != nil {
		e.shouldPanic(vars.Err().Error())
		// TODO surface *Echo errors
		// e.err = vars.Err()
	}
	return e
}

// SetEnv sets a global process environment variable.
func (e *Echo) SetEnv(name, value string) *Echo {
	vars := e.vars.SetEnv(name,value)
	if vars.Err() != nil {
		e.shouldPanic(vars.Err().Error())
	}
	return e
}

// Vars declares variables used in current *Echo session using:
// e.Vars("foo=bar fuzz=buzz")
func (e *Echo) Vars(val string) *Echo {
	vars := e.vars.Vars(val)
	if vars.Err() != nil {
		e.shouldPanic(vars.Err().Error())
	}
	return e
}

// SetVar declares a session variable.
func (e *Echo) SetVar(name, value string) *Echo {
	vars := e.vars.SetVar(name, value)
	if vars.Err() != nil {
		e.shouldPanic(vars.Err().Error())
	}
	return e
}

// Val retrieves a session or environment variable
func (e *Echo) Val(name string) string {
	return e.vars.Val(name)
}

// Eval returns the string str with its content expanded
// with variable values i.e. Eval("I am $HOME") returns
// "I am </user/home/path>"
func (e *Echo) Eval(str string) string {
	return e.vars.Eval(str)
}
