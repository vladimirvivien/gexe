package echo

import (
	"os"
	"path/filepath"
	"regexp"
)

func (e *echo) Arr(list string, seps ...string) []string {
	sepRegx := regexp.MustCompile("\\s")
	if len(seps) > 0 {
		sepRegx = regexp.MustCompile(seps[0])
	}
	return sepRegx.Split(os.Expand(list, e.Val), -1)
}

func (e *echo) Glob(pathPattern string) []string {
	matches, err := filepath.Glob(os.Expand(pathPattern, e.Val))
	if err != nil {
		return nil
	}
	return matches
}
