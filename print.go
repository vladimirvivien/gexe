package gexe

import (
	"fmt"
	"io"
)

// Print outputs the formatted string to stdout without a newline.
// The string supports both Go's fmt.Sprintf formatting and gexe variable expansion.
// Variables are expanded using ${VAR} or $VAR syntax.
// Returns the session for method chaining.
//
// Example:
//
//	gexe.Print("Processing %s in ${HOME}", filename)
func (e *Session) Print(format string, args ...interface{}) *Session {
	expanded := e.vars.Eval(applyFmt(format, args...))
	fmt.Print(expanded)
	return e
}

// Println outputs the formatted string to stdout with a trailing newline.
// The string supports both Go's fmt.Sprintf formatting and gexe variable expansion.
// Variables are expanded using ${VAR} or $VAR syntax.
// Returns the session for method chaining.
//
// Example:
//
//	gexe.Println("User ${USER} logged in at %s", timestamp)
func (e *Session) Println(format string, args ...interface{}) *Session {
	expanded := e.vars.Eval(applyFmt(format, args...))
	fmt.Println(expanded)

	return e
}

// PrintTo outputs the formatted string to the specified writer.
// The string supports both Go's fmt.Sprintf formatting and gexe variable expansion.
// Variables are expanded using ${VAR} or $VAR syntax.
// Returns the session for method chaining.
//
// Example:
//
//	var buf bytes.Buffer
//	gexe.PrintTo(&buf, "Log: ${USER} performed %s", action)
func (e *Session) PrintTo(w io.Writer, format string, args ...interface{}) *Session {
	expanded := e.vars.Eval(applyFmt(format, args...))
	fmt.Fprint(w, expanded)

	return e
}
