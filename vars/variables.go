package vars

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

var (
	varsKeyValRgx = regexp.MustCompile("\\s*=\\s*")
	varsLineRegx  = regexp.MustCompile("\\w+\\s*=\\s*\\w+")
)

type Variables struct {
	err error
	vars map[string]string
	escapeChar rune
}

func New() *Variables {
	return &Variables{vars: make(map[string]string), escapeChar: '\\'}
}

func (v *Variables) WithEscapeChar(r rune) *Variables {
	v.escapeChar = r
	return v
}

// Envs declares process environment variables using
// a multi-line space-separated list of KEY=VAL format:
// i.e. GOOS=linux GOARCH=amd64
func (v *Variables) Envs(val string) *Variables {
	vars, err := parseVars(val)
	if err != nil {
		v.err = err
		return v
	}
	for key, value := range vars {
		if err := os.Setenv(key, v.ExpandVar(value, v.Val)); err != nil {
			v.err = err
			return v
		}
	}
	return v
}

// SetEnv sets a process environment variable.
func (v *Variables) SetEnv(name, value string) *Variables {
	if err := os.Setenv(name, v.ExpandVar(value, v.Val)); err != nil {
		v.err = err
		return v
	}
	return v
}

// Var declares an internal variable used during current echo session.
// It uses a multi-line, space-separated list of KEY=VAL format:
// i.e. foo=bar fuzz=buzz
func (v *Variables) Var(val string) *Variables {
	vars, err := parseVars(val)
	if err != nil {
		v.err = err
		return v
	}

	for key, val := range vars {
		v.vars[key] = v.ExpandVar(val, v.Val)
	}

	return v
}

// SetVar declares an in-process local variable.
func (v *Variables) SetVar(name, value string) *Variables {
	v.vars[name] = v.ExpandVar(value, v.Val)
	return v
}

// Val retrieves an in-process variable if found
// or process environment variable
func (v *Variables) Val(name string) string {
	if val, ok := v.vars[name]; ok {
		return val
	}
	return os.Getenv(name)
}

// Eval returns the string str with its content expanded
// with variable values i.e. Eval("I am $HOME") returns
// "I am </user/home/path>"
func (v *Variables) Eval(str string) string {
	return v.ExpandVar(str, v.Val)
}

// parseVars parses multi-line, space-separated key=value pairs
// into map[string]string
func parseVars(lines string) (map[string]string, error) {
	// parse lines into envs = []{"KEY0=VAL0", "KEY1=VAL1",...}
	var envs []string
	scnr := bufio.NewScanner(strings.NewReader(lines))

	for scnr.Scan() {
		envs = append(envs, varsLineRegx.FindAllString(scnr.Text(), -1)...)
	}
	if err := scnr.Err(); err != nil {
		return nil, err
	}

	// parse each item in []string{"key=value",...} item into key=value
	result := make(map[string]string)
	for _, env := range envs {
		kv := varsKeyValRgx.Split(env, 2)
		if len(kv) == 2 {
			result[kv[0]] = kv[1]
		}
	}

	return result, nil
}

