package echo

import (
	"os"
	"strings"
)

func (e *echo) Empty(val string) bool {
	return len(os.Expand(val, e.Val)) == 0
}

func (e *echo) Lower(val string) string {
	return strings.ToLower(os.Expand(val, e.Val))
}

func (e *echo) Upper(val string) string {
	return strings.ToUpper(os.Expand(val, e.Val))
}

func (e *echo) Streq(val0, val1 string) bool {
	return strings.EqualFold(
		os.Expand(val0, e.Val),
		os.Expand(val1, e.Val),
	)
}

func (e *echo) Trim(val string) string {
	return strings.TrimSpace(os.Expand(val, e.Val))
}
