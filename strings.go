package echo

import (
	"os"
	"strings"
)

// Empty returns true if len(val) == 0
func (e *echo) Empty(val string) bool {
	return len(os.Expand(val, e.Val)) == 0
}

// Lower returns val as lower case
func (e *echo) Lower(val string) string {
	return strings.ToLower(os.Expand(val, e.Val))
}

// Upper returns val as upper case
func (e *echo) Upper(val string) string {
	return strings.ToUpper(os.Expand(val, e.Val))
}

// Streq returns true if both strings are equal
func (e *echo) Streq(val0, val1 string) bool {
	return strings.EqualFold(
		os.Expand(val0, e.Val),
		os.Expand(val1, e.Val),
	)
}

func (e *echo) Trim(val string) string {
	return strings.TrimSpace(os.Expand(val, e.Val))
}
