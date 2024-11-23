package gexe

import (
	"context"
	"fmt"

	"github.com/vladimirvivien/gexe/exec"
)

// NewProc setups a new process with specified command cmdStr and returns immediately
// without starting. Use Proc.Wait to wait for exection and then retrieve process result.
// Information about the running process is stored in *exec.Proc.
func (e *Echo) NewProcWithContext(ctx context.Context, cmdStr string) *exec.Proc {
	return exec.NewProcWithContextVars(ctx, cmdStr, e.vars)
}

// NewProc a convenient function that calls NewProcWithContext with a default contet.
func (e *Echo) NewProc(cmdStr string) *exec.Proc {
	return exec.NewProcWithContextVars(context.Background(), cmdStr, e.vars)
}

// StartProc executes the command in cmdStr, with the specified context, and returns immediately
// without waiting. Use Proc.Wait to wait for exection and then retrieve process result.
// Information about the running process is stored in *Proc.
func (e *Echo) StartProcWithContext(ctx context.Context, cmdStr string) *exec.Proc {
	return exec.StartProcWithContextVars(ctx, cmdStr, e.vars)
}

// StartProc executes the command in cmdStr and returns immediately
// without waiting. Use Proc.Wait to wait for exection and then retrieve process result.
// Information about the running process is stored in *Proc.
func (e *Echo) StartProc(cmdStr string) *exec.Proc {
	return exec.StartProcWithContextVars(context.Background(), cmdStr, e.vars)
}

// RunProcWithContext executes command in cmdStr, with given context, and waits for the result.
// It returns a *Proc with information about the executed process.
func (e *Echo) RunProcWithContext(ctx context.Context, cmdStr string) *exec.Proc {
	return exec.RunProcWithContextVars(ctx, cmdStr, e.vars)
}

// RunProc executes command in cmdStr and waits for the result.
// It returns a *Proc with information about the executed process.
func (e *Echo) RunProc(cmdStr string) *exec.Proc {
	return exec.RunProcWithContextVars(context.Background(), cmdStr, e.vars)
}

// Run executes cmdStr, with given context, and returns the result as a string.
func (e *Echo) RunWithContext(ctx context.Context, cmdStr string) string {
	return exec.RunWithContextVars(ctx, cmdStr, e.vars)
}

// Run executes cmdStr, waits, and returns the result as a string.
func (e *Echo) Run(cmdStr string) string {
	return exec.RunWithContextVars(context.Background(), cmdStr, e.vars)
}

// Runout executes command cmdStr and prints out the result
func (e *Echo) Runout(cmdStr string) {
	fmt.Print(e.Run(cmdStr))
}

// Commands returns a *exe.CommandBuilder to build a multi-command execution flow.
func (e *Echo) Commands(cmdStrs ...string) *exec.CommandBuilder {
	return exec.CommandsWithVars(e.vars, cmdStrs...)
}

// StartAll starts the sequential execution of each command, in cmdStrs, and does not
// wait for their completion.
func (e *Echo) StartAll(cmdStrs ...string) *exec.CommandResult {
	return exec.CommandsWithVars(e.vars, cmdStrs...).Start()
}

// RunAll executes each command sequentially, in cmdStrs, and wait for their completion.
func (e *Echo) RunAll(cmdStrs ...string) *exec.CommandResult {
	return exec.CommandsWithVars(e.vars, cmdStrs...).Run()
}

// StartConcur starts the concurrent execution of each command, in cmdStrs, and does not
// wait for their completion.
func (e *Echo) StartConcur(cmdStrs ...string) *exec.CommandResult {
	return exec.CommandsWithVars(e.vars, cmdStrs...).Concurr()
}

// RunConcur executes each command concurrently, in cmdStrs, and waits
// their completion.
func (e *Echo) RunConcur(cmdStrs ...string) *exec.CommandResult {
	return exec.CommandsWithVars(e.vars, cmdStrs...).Concurr().Wait()
}

// Pipe executes each command, in cmdStrs, by piping the result
// of the previous command as input to the next command until done.
func (e *Echo) Pipe(cmdStrs ...string) *exec.PipedCommandResult {
	return exec.CommandsWithVars(e.vars, cmdStrs...).Pipe()
}

// ParseCommand parses the string into individual command tokens
func (e *Echo) ParseCommand(cmdStr string) (cmdName string, args []string) {
	result, err := exec.Parse(e.vars.Eval(cmdStr))
	if err != nil {
		e.err = err
	}
	cmdName = result[0]
	args = result[1:]
	return
}
