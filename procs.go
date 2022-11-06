package gexe

import (
	"fmt"

	"github.com/vladimirvivien/gexe/exec"
)

// NewProc setups a new process with specified command cmdStr and returns immediately
// without starting. Use Proc.Wait to wait for exection and then retrieve process result.
// Information about the running process is stored in *exec.Proc.
func (e *Echo) NewProc(cmdStr string) *exec.Proc {
	return exec.NewProc(cmdStr)
}

// StartProc executes the command in cmdStr and returns immediately
// without waiting. Use Proc.Wait to wait for exection and then retrieve process result.
// Information about the running process is stored in *Proc.
func (e *Echo) StartProc(cmdStr string) *exec.Proc {
	return exec.StartProc(e.Eval(cmdStr))
}

// RunProc executes command in cmdStr and waits for the result.
// It returns a *Proc with information about the executed process.
func (e *Echo) RunProc(cmdStr string) *exec.Proc {
	return exec.RunProc(e.Eval(cmdStr))
}

// Run executes cmdStr, waits, and returns the result as a string.
func (e *Echo) Run(cmdStr string) string {
	return exec.Run(e.Eval(cmdStr))
}

// Runout executes command cmdStr and prints out the result
func (e *Echo) Runout(cmdStr string) {
	fmt.Print(e.Run(cmdStr))
}

// RunAll executes each command, in cmdStrs, successively.
func (e *Echo) RunAll(cmdStrs ...string) *exec.CommandResult {
	for i, cmd := range cmdStrs {
		cmdStrs[i] = e.Eval(cmd)
	}
	return exec.Commands(cmdStrs...).Run()
}

// RunConcur executes each command, in cmdStrs, concurrently and waits
// their completion.
func (e *Echo) RunConcur(cmdStrs ...string) *exec.CommandResult {
	for i, cmd := range cmdStrs {
		cmdStrs[i] = e.Eval(cmd)
	}
	return exec.Commands(cmdStrs...).Concurr()
}

// Pipe executes each command, in cmdStrs, by piping the result
// of the previous command as input to the next command until done.
func (e *Echo) Pipe(cmdStrs ...string) *exec.PipedCommandResult {
	for i, cmd := range cmdStrs {
		cmdStrs[i] = e.Eval(cmd)
	}
	return exec.Commands(cmdStrs...).Pipe()
}
