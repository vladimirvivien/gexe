package echo

import (
	"os"
	"strings"
)

// Empty returns true if len(val) == 0
func (e *Echo) Empty(val string) bool {
	return len(os.Expand(val, e.Val)) == 0
}

// Lower returns val as lower case
func (e *Echo) Lower(val string) string {
	return strings.ToLower(os.Expand(val, e.Val))
}

// Upper returns val as upper case
func (e *Echo) Upper(val string) string {
	return strings.ToUpper(os.Expand(val, e.Val))
}

// Streq returns true if both strings are equal
func (e *Echo) Streq(val0, val1 string) bool {
	return strings.EqualFold(
		os.Expand(val0, e.Val),
		os.Expand(val1, e.Val),
	)
}

// Trim removes spaces around a val
func (e *Echo) Trim(val string) string {
	return strings.TrimSpace(os.Expand(val, e.Val))
}
