package echo

import (
	"bufio"
	"os"
	"strings"
)

// Env declares process environment variables using
// a multi-line space-separated list of KEY=VAL format:
// i.e. GOOS=linux GOARCH=amd64
func (e *echo) Env(val string) *echo {
	vars := e.declareVars(val)
	for k, v := range vars {
		if err := os.Setenv(k, v); err != nil {
			e.shouldPanic(err.Error())
		} else {
			delete(e.vars, k) // overwrite local var
		}
	}
	return e
}

// SetEnv declares a global process environment variable.
func (e *echo) SetEnv(name, value string) *echo {
	if err := os.Setenv(name, os.Expand(value, e.Val)); err != nil {
		e.shouldPanic(err.Error())
	} else {
		delete(e.vars, name)
	}
	return e
}

// Var declares variables used during current echo session using
// a multi-line space-separated list of KEY=VAL format:
// i.e. foo=bar fuzz=buzz
func (e *echo) Var(val string) *echo {
	vars := e.declareVars(val)
	for k, v := range vars {
		os.Unsetenv(k)
		e.vars[k] = v
	}
	return e
}

// SetVar declares an echo session local variable.
func (e *echo) SetVar(name, value string) *echo {
	os.Unsetenv(name)
	e.vars[name] = os.Expand(value, e.Val)
	return e
}

// Val retrieves a session or environment variable
func (e *echo) Val(name string) string {
	if val, ok := e.vars[name]; ok {
		return val
	}
	return os.Getenv(name)
}

// Eval returns the string str with its content expanded
// with variable values i.e. Eval("I am $HOME") returns
// "I am <user home path>"
func (e *echo) Eval(str string) string {
	return os.Expand(str, e.Val)
}

func (e *echo) declareVars(val string) map[string]string {
	evaled := e.Eval(val)

	// parse lines into envs = []{"KEY0=VAL0", "KEY1=VAL1",...}
	var envs []string
	scnr := bufio.NewScanner(strings.NewReader(evaled))
	for scnr.Scan() {
		envs = append(envs, spaceRgx.Split(scnr.Text(), -1)...)
	}
	if err := scnr.Err(); err != nil {
		e.shouldPanic(err.Error())
	}

	result := make(map[string]string)
	for _, env := range envs {
		kv := strings.Split(env, "=")
		if len(kv) == 2 {
			result[kv[0]] = kv[1]
		}
	}

	return result
}
