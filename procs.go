package gexe

import (
	"context"
	"fmt"

	"github.com/vladimirvivien/gexe/exec"
)

// NewProc setups a new process with specified command cmdStr and returns immediately
// without starting. Use Proc.Wait to wait for exection and then retrieve process result.
// Information about the running process is stored in *exec.Proc.
func (e *Session) NewProcWithContext(ctx context.Context, cmdStr string, args ...interface{}) *exec.Proc {
	cmdStr = applyFmt(cmdStr, args...)
	return exec.NewProcWithContextVars(ctx, cmdStr, e.vars)
}

// NewProc a convenient function that calls NewProcWithContext with a default contet.
func (e *Session) NewProc(cmdStr string, args ...interface{}) *exec.Proc {
	cmdStr = applyFmt(cmdStr, args...)
	return exec.NewProcWithContextVars(context.Background(), cmdStr, e.vars)
}

// StartProc executes the command in cmdStr, with the specified context, and returns immediately
// without waiting. Use Proc.Wait to wait for exection and then retrieve process result.
// Information about the running process is stored in *Proc.
func (e *Session) StartProcWithContext(ctx context.Context, cmdStr string, args ...interface{}) *exec.Proc {
	cmdStr = applyFmt(cmdStr, args...)
	return exec.StartProcWithContextVars(ctx, cmdStr, e.vars)
}

// StartProc executes the command in cmdStr and returns immediately
// without waiting. Use Proc.Wait to wait for exection and then retrieve process result.
// Information about the running process is stored in *Proc.
func (e *Session) StartProc(cmdStr string, args ...interface{}) *exec.Proc {
	cmdStr = applyFmt(cmdStr, args...)
	return exec.StartProcWithContextVars(context.Background(), cmdStr, e.vars)
}

// RunProcWithContext executes command in cmdStr, with given context, and waits for the result.
// It returns a *Proc with information about the executed process.
func (e *Session) RunProcWithContext(ctx context.Context, cmdStr string, args ...interface{}) *exec.Proc {
	cmdStr = applyFmt(cmdStr, args...)
	return exec.RunProcWithContextVars(ctx, cmdStr, e.vars)
}

// RunProc executes command in cmdStr and waits for the result.
// It returns a *Proc with information about the executed process.
func (e *Session) RunProc(cmdStr string, args ...interface{}) *exec.Proc {
	cmdStr = applyFmt(cmdStr, args...)
	return exec.RunProcWithContextVars(context.Background(), cmdStr, e.vars)
}

// Run executes cmdStr, with given context, and returns the result as a string.
func (e *Session) RunWithContext(ctx context.Context, cmdStr string, args ...interface{}) string {
	cmdStr = applyFmt(cmdStr, args...)
	return exec.RunWithContextVars(ctx, cmdStr, e.vars)
}

// Run executes cmdStr, waits, and returns the result as a string.
func (e *Session) Run(cmdStr string, args ...interface{}) string {
	cmdStr = applyFmt(cmdStr, args...)
	return exec.RunWithContextVars(context.Background(), cmdStr, e.vars)
}

// Runout executes command cmdStr and prints out the result
func (e *Session) Runout(cmdStr string, args ...interface{}) {
	fmt.Print(e.Run(cmdStr, args...))
}

// Commands creates a *exe.CommandBuilder, with the specified context, to build a multi-command execution flow.
func (e *Session) CommandsWithContext(ctx context.Context, cmdStrs ...string) *exec.CommandBuilder {
	return exec.CommandsWithContextVars(ctx, e.vars, cmdStrs...)
}

// Commands returns a *exe.CommandBuilder to build a multi-command execution flow.
func (e *Session) Commands(cmdStrs ...string) *exec.CommandBuilder {
	return exec.CommandsWithContextVars(context.Background(), e.vars, cmdStrs...)
}

// StartAllWithContext uses the specified ctx to start sequential execution of each command, in cmdStrs, and does not
// wait for their completion.
func (e *Session) StartAllWithContext(ctx context.Context, cmdStrs ...string) *exec.CommandResult {
	return exec.CommandsWithContextVars(ctx, e.vars, cmdStrs...).Start()
}

// StartAll starts the sequential execution of each command, in cmdStrs, and does not
// wait for their completion.
func (e *Session) StartAll(cmdStrs ...string) *exec.CommandResult {
	return exec.CommandsWithContextVars(context.Background(), e.vars, cmdStrs...).Start()
}

// RunAllWithContext executes each command sequentially, in cmdStrs, and wait for their completion.
func (e *Session) RunAllWithContext(ctx context.Context, cmdStrs ...string) *exec.CommandResult {
	return exec.CommandsWithContextVars(ctx, e.vars, cmdStrs...).Run()
}

// RunAll executes each command sequentially, in cmdStrs, and wait for their completion.
func (e *Session) RunAll(cmdStrs ...string) *exec.CommandResult {
	return exec.CommandsWithContextVars(context.Background(), e.vars, cmdStrs...).Run()
}

// StartConcurWithContext uses specified context to start the concurrent execution of each command, in cmdStrs, and does not
// wait for their completion.
func (e *Session) StartConcurWithContext(ctx context.Context, cmdStrs ...string) *exec.CommandResult {
	return exec.CommandsWithContextVars(ctx, e.vars, cmdStrs...).Concurr()
}

// StartConcur starts the concurrent execution of each command, in cmdStrs, and does not
// wait for their completion.
func (e *Session) StartConcur(cmdStrs ...string) *exec.CommandResult {
	return exec.CommandsWithContextVars(context.Background(), e.vars, cmdStrs...).Concurr()
}

// RunConcurWithContext uses context to execute each command concurrently, in cmdStrs, and waits
// their completion.
func (e *Session) RunConcurWithContext(ctx context.Context, cmdStrs ...string) *exec.CommandResult {
	return exec.CommandsWithContextVars(ctx, e.vars, cmdStrs...).Concurr().Wait()
}

// RunConcur executes each command concurrently, in cmdStrs, and waits
// their completion.
func (e *Session) RunConcur(cmdStrs ...string) *exec.CommandResult {
	return exec.CommandsWithContextVars(context.Background(), e.vars, cmdStrs...).Concurr().Wait()
}

// Pipe uses specified context to execute each command, in cmdStrs, by piping the result
// of the previous command as input to the next command until done.
func (e *Session) PipeWithContext(ctx context.Context, cmdStrs ...string) *exec.PipedCommandResult {
	return exec.CommandsWithContextVars(ctx, e.vars, cmdStrs...).Pipe()
}

// Pipe executes each command, in cmdStrs, by piping the result
// of the previous command as input to the next command until done.
func (e *Session) Pipe(cmdStrs ...string) *exec.PipedCommandResult {
	return exec.CommandsWithContextVars(context.Background(), e.vars, cmdStrs...).Pipe()
}

// ParseCommand parses the string into individual command tokens
func (e *Session) ParseCommand(cmdStr string, args ...interface{}) (cmdName string, argsList []string) {
	cmdStr = applyFmt(cmdStr, args...)
	result, err := exec.Parse(e.vars.Eval(cmdStr))
	if err != nil {
		e.err = err
	}
	cmdName = result[0]
	argsList = result[1:]
	return
}
