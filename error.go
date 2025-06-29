package gexe

import "errors"

// Error creates and returns an error with the formatted string.
// The string supports both Go's fmt.Sprintf formatting and gexe variable expansion.
// Variables are expanded using ${VAR} or $VAR syntax.
//
// Example:
//
//	err := session.Error("Failed to process %s in ${HOME}", filename)
//	return err
func (e *Session) Error(format string, args ...interface{}) error {
	expanded := e.vars.Eval(applyFmt(format, args...))
	return errors.New(expanded)
}
