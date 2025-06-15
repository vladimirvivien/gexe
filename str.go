package gexe

import (
	"github.com/vladimirvivien/gexe/str"
)

// String creates a new str.Str value with string manipulation methods
func (e *Session) String(s string, args ...interface{}) *str.Str {
	s = applyFmt(s, args...)
	return str.StringWithVars(s, e.vars)
}
