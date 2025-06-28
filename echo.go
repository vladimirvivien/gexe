package gexe

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/vladimirvivien/gexe/prog"
	"github.com/vladimirvivien/gexe/vars"
)

type (
	// Echo is a Session alias for backward compatibility
	Echo = Session
)

var (
	// DefaultSession surfaces an gexe session used for all package functions
	DefaultSession = New()
	DefaultEcho    = DefaultSession // alias for backward compatibility
)

// Session represents a new session used for accessing
// Gexe types and methods.
type Session struct {
	err  error
	vars *vars.Variables // session vars
	prog *prog.Info
}

// New creates a new Echo session
func New() *Session {
	e := &Session{
		vars: vars.New(),
		prog: prog.Prog(),
	}
	return e
}

// AddExecPath adds an executable path to PATH
func (e *Session) AddExecPath(execPath string) {
	oldPath := os.Getenv("PATH")
	os.Setenv("PATH", fmt.Sprintf("%s%c%s", oldPath, os.PathListSeparator, e.Eval(execPath)))
}

// ProgAvail returns the full path of the program if found on exec PATH
func (e *Session) ProgAvail(progName string, args ...interface{}) string {
	progName = applyFmt(progName, args...)
	path, err := exec.LookPath(e.Eval(progName))
	if err != nil {
		return ""
	}
	return path
}
