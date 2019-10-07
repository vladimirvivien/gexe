package echo

import "os"

func (e *echo) Export(name, value string) {
	delete(e.vars, name)
	os.Setenv(name, os.Expand(value, e.Val))
}

func (e *echo) Var(name, value string) {
	os.Unsetenv(name)
	e.vars[name] = os.Expand(value, e.Val)
}

func (e *echo) Val(name string) string {
	if val, ok := e.vars[name]; ok {
		return val
	}
	return os.Getenv(name)
}

func (e *echo) Eval(str string) string {
	return os.Expand(str, e.Val)
}
