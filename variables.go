package gexe

import (
	"github.com/vladimirvivien/gexe/vars"
)

// Variables returns the variable mapping for echo session e
func (e *Session) Variables() *vars.Variables {
	return e.vars
}

// Envs declares environment variables using
// a multi-line space-separated list:
//
//	Envs("GOOS=linux" "GOARCH=amd64", `platform="$GOOS:$GOARCH"`)
//
// Environment vars can be used in string values
// using Eval("building for os=$GOOS")
func (e *Session) Envs(variables ...string) *Session {
	vars := e.vars.Envs(variables...)
	e.err = vars.Err()
	return e
}

// SetEnv sets a global process environment variable.
func (e *Session) SetEnv(name, value string) *Session {
	vars := e.vars.SetEnv(name, value)
	e.err = vars.Err()
	return e
}

// Vars declares multiple session-scope variables using
// string literal format:
//
// Envs("foo=bar", "platform=amd64", `"data="info ${platform}"`)
//
// Note that session vars are only available
// for the running process.
func (e *Session) Vars(variables ...string) *Session {
	vars := e.vars.Vars(variables...)
	e.err = vars.Err()
	return e
}

// SetVar declares a session variable.
func (e *Session) SetVar(name, value string) *Session {
	vars := e.vars.SetVar(name, value)
	e.err = vars.Err()
	return e
}

// UnsetVar removes a session variable.
func (e *Session) UnsetVar(name string) *Session {
	vars := e.vars.UnsetVar(name)
	e.err = vars.Err()
	return e
}

// Val retrieves a session or environment variable
func (e *Session) Val(name string) string {
	return e.vars.Val(name)
}

// Eval returns the string str with its content expanded
// with variable values i.e. Eval("I am $HOME") returns
// "I am </user/home/path>"
func (e *Session) Eval(str string) string {
	return e.vars.Eval(str)
}
