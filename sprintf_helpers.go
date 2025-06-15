package gexe

import "fmt"

// applyFmt applies fmt.Sprintf only if there are format verbs and args
func applyFmt(format string, args ...any) string {
	if len(args) > 0 {
		return fmt.Sprintf(format, args...)
	}
	return format
}
