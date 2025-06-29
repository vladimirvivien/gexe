package gexe

import (
	"path/filepath"
	"strings"
)

// Join uses strings.Join to join arbitrary strings with a separator
// while applying gexe variable expansion to each element
func (e *Session) Join(sep string, elem ...string) string {
	expandedElems := make([]string, len(elem))
	for i, element := range elem {
		expandedElems[i] = e.vars.Eval(element)
	}
	return strings.Join(expandedElems, sep)
}

// JoinPath uses filepath.Join to join file paths using OS-specific path separators
// while applying gexe variable expansion to each element
func (e *Session) JoinPath(elem ...string) string {
	expandedElems := make([]string, len(elem))
	for i, element := range elem {
		expandedElems[i] = e.vars.Eval(element)
	}
	return filepath.Join(expandedElems...)
}
