package echo

import (
	"os"
	"path/filepath"
	"regexp"
)

// Split splits a space-separated (default) string into []string.
// An optional separator sep maybe provided as well.
func (e *echo) Split(list string, sep ...string) []string {
	sepRegx := spaceRgx
	if len(sep) > 0 {
		sepRegx = regexp.MustCompile(sep[0])
	}
	return sepRegx.Split(os.Expand(list, e.Val), -1)
}

// Glob uses shell file path pattern (i.e. /usr/*/*a) to return
// a []string of file names
func (e *echo) Glob(pathPattern string) []string {
	matches, err := filepath.Glob(os.Expand(pathPattern, e.Val))
	if err != nil {
		e.shouldPanic(err.Error())
		return []string{}
	}
	return matches
}
