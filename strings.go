package echo

import (
	"strings"
)

// Empty returns true if len(val) == 0
func (e *Echo) Empty(val string) bool {
	return len(e.Variables().Eval(val)) == 0
}

// Lower returns val as lower case
func (e *Echo) Lower(val string) string {
	return strings.ToLower(e.Variables().Eval(val))
}

// Upper returns val as upper case
func (e *Echo) Upper(val string) string {
	return strings.ToUpper(e.Variables().Eval(val))
}

// Streq returns true if both strings are equal
func (e *Echo) Streq(val0, val1 string) bool {
	return strings.EqualFold(
		e.Variables().Eval(val0),
		e.Variables().Eval(val1),
	)
}

// Trim removes spaces around a val
func (e *Echo) Trim(val string) string {
	return strings.TrimSpace(e.Variables().Eval(val))
}
