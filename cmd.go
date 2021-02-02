package echo

import (
	"fmt"
	"os"

	"github.com/vladimirvivien/echo/exec"
)

// RunProc executes command in cmdStr and waits for the result.
// It returns a *Proc with information about the executed process.
func (e *Echo) RunProc(cmdStr string) *exec.Proc {
	e.shouldLog(cmdStr)
	proc := e.StartProc(cmdStr)

	if proc.Err() != nil {
		return proc
	}
	proc.Wait()
	return proc
}

// StartProc executes the command in cmdStr and returns immediately
// without waiting. Information about the running process is stored in *Proc.
func (e *Echo) StartProc(cmdStr string) *exec.Proc {
	// TODO expand variables
	p := exec.StartProc(cmdStr)
	if p.Err() != nil {
		e.shouldLog(p.Err().Error())
		e.shouldPanic(p.Err().Error())
		return p
	}
	return p
}

// Run executes cmdStr, waits, and returns the result as a string.
func (e *Echo) Run(cmdStr string) string {
	p := e.RunProc(cmdStr)
	data := p.Result()
	if p.Err() != nil {
		e.shouldLog(p.Err().Error())
		e.shouldPanic(p.Err().Error())
		return ""
	}

	return os.Expand(data, e.Val)
}

// Runout executes command cmdStr and prints out the result
func (e *Echo) Runout(cmdStr string) {
	fmt.Print(e.Run(cmdStr))
}