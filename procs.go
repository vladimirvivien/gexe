package echo

import (
	"fmt"

	"github.com/vladimirvivien/echo/exec"
)

// StartProc executes the command in cmdStr and returns immediately
// without waiting. Information about the running process is stored in *Proc.
func (e *Echo) StartProc(cmdStr string) *exec.Proc {
	p := exec.StartProc(cmdStr)
	if p.Err() != nil {
		e.shouldLog(p.Err().Error())
		e.shouldPanic(p.Err().Error())
		return p
	}
	return p
}

// RunProc executes command in cmdStr and waits for the result.
// It returns a *Proc with information about the executed process.
func (e *Echo) RunProc(cmdStr string) *exec.Proc {
	e.shouldLog(cmdStr)
	proc :=  exec.RunProc (cmdStr)

	if proc.Err() != nil{
		e.shouldLog(proc.Err().Error())
		e.shouldPanic(proc.Err().Error())
		return proc
	}

	return proc
}

// Run executes cmdStr, waits, and returns the result as a string.
func (e *Echo) Run(cmdStr string) string {
	proc := exec.StartProc(cmdStr)
	if proc.Err() != nil {
		e.shouldLog(proc.Err().Error())
		e.shouldPanic(proc.Err().Error())
		return ""
	}
	return e.Variables().Eval(proc.Result())
}

// Runout executes command cmdStr and prints out the result
func (e *Echo) Runout(cmdStr string) {
	fmt.Print(e.Run(cmdStr))
}