package gexe

import "fmt"

// hasFormatVerbs checks if a string contains format verbs like %s, %d, etc.
func hasFormatVerbs(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '%' && s[i+1] != '%' { // %% is escaped %
			return true
		}
	}
	return false
}

// applySprintfIfNeeded applies fmt.Sprintf only if there are format verbs and args
func applySprintfIfNeeded(format string, args ...interface{}) string {
	if len(args) > 0 && hasFormatVerbs(format) {
		return fmt.Sprintf(format, args...)
	}
	return format
}
