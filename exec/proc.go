package exec

import (
	"bytes"
	"fmt"
	"io"
	"os"
	osexec "os/exec"
	"strings"
	"time"
)

// Proc stores process info
type Proc struct {
	id         int
	err        error
	state      *os.ProcessState
	result     *bytes.Buffer
	output     io.Reader
	stdoutPipe io.ReadCloser
	stderrPipe io.ReadCloser
	cmd        *osexec.Cmd
	process    *os.Process
}

// StartProc creates a *Proc and starts an OS process but does not wait for
// it to complete.
func StartProc(cmdStr string) *Proc {
	words, err := parse(cmdStr)
	if err != nil {
		return &Proc{err: err}
	}

	command := osexec.Command(words[0], words[1:]...)
	pipeout, outerr := command.StdoutPipe()
	pipeerr, errerr := command.StderrPipe()
	output := io.MultiReader(pipeout, pipeerr)

	if outerr != nil || errerr != nil {
		return &Proc{err: fmt.Errorf("%s; %s", outerr, errerr)}
	}

	if err := command.Start(); err != nil {
		return &Proc{id: command.Process.Pid, cmd: command, state: command.ProcessState, err: err}
	}

	return &Proc{
		id:         command.Process.Pid,
		cmd:        command,
		process:    command.Process,
		state:      command.ProcessState,
		stdoutPipe: pipeout,
		stderrPipe: pipeerr,
		output:     output,
	}
}

// RunProc creates, runs, and waits for a process to complete and
// return *Proc with result info.  This can be followed by
// Proc.Result() to access to string value of the returned by process.
func RunProc(cmdStr string) *Proc {
	proc := StartProc(cmdStr)
	if proc.Err() != nil {
		proc.err = fmt.Errorf("proc: runproc: %s", proc.Err())
		return proc
	}

	proc.result = new(bytes.Buffer)

	if _, err := proc.result.ReadFrom(proc.output); err != nil {
		proc.err = err
		return proc
	}

	if err := proc.Wait().Err(); err != nil {
		proc.err = fmt.Errorf("proc: error: %s", proc.Err())
	}

	return proc
}

// Run creates and runs a process and returns the
// result as a string.
// Equivalent to: Proc.StartProc() -> Proc.Result()
func Run(cmdStr string) (result string) {
	proc := StartProc(cmdStr)
	if proc.Err() != nil {
		proc.err = fmt.Errorf("proc: run: %s", proc.Err())
		return
	}
	return proc.Result()
}

// Command returns the os/exec.Cmd that started the process
func (p *Proc) Command() *osexec.Cmd {
	return p.cmd
}

// Peek attempts to read process state information
func (p *Proc) Peek() *Proc {
	p.state = p.cmd.ProcessState
	return p
}

// Wait waits for a process started with Proc.StartProc to complete (in a separate goroutine).
// Once process completes, Wait cleans up resources.
// Must be called after Proc.StartProc()
func (p *Proc) Wait() *Proc {
	if p.cmd == nil {
		p.err = fmt.Errorf("command is nill")
		return p
	}
	if err := p.cmd.Wait(); err != nil {
		p.err = err
		// use return below to get proc info
	}
	return p.Peek()
}

// ID returns process id
func (p *Proc) ID() int {
	return p.id
}

// Exited returns true if process exits ok
func (p *Proc) Exited() bool {
	if p.state == nil {
		return false
	}
	return p.state.Exited()
}

// ExitCode returns process exit code
func (p *Proc) ExitCode() int {
	if p.state == nil {
		return -1
	}
	return p.state.ExitCode()
}

// IsSuccess returns true if proc exit ok
func (p *Proc) IsSuccess() bool {
	if p.state == nil {
		return false
	}
	return p.state.Success()
}

// SysTime returns proc system cpu time
func (p *Proc) SysTime() time.Duration {
	if p.state == nil {
		return -1
	}
	return p.state.SystemTime()
}

// UserTime returns proc user cpu time
func (p *Proc) UserTime() time.Duration {
	if p.state == nil {
		return -1
	}
	return p.state.UserTime()
}

// Err returns any execution error
func (p *Proc) Err() error {
	return p.err
}

// Kill halts the process
func (p *Proc) Kill() *Proc {
	if err := p.cmd.Process.Kill(); err != nil {
		p.err = err
	}
	return p
}

// StdOut is a io.Reader pipe for standard out
// Must be streamed before Proc.Wait()
func (p *Proc) StdOut() io.Reader {
	return p.stdoutPipe
}

// StdErr is an io.Reader pipe for standard error
// Must be streamed before Proc.Wait().
func (p *Proc) StdErr() io.Reader {
	return p.stdoutPipe
}

// Out waits for cmd result and surfaces the result from both stdout and stderr
// in an io.Reader that can be streamed.
// Must be called after after Proc.StartProc, Proc.RunProc
func (p *Proc) Out() io.Reader {
	if p.output != nil {
		return p.output
	}
	p.output = io.MultiReader(p.stdoutPipe, p.stderrPipe)

	if err := p.Wait().Err(); err != nil {
		p.err = fmt.Errorf("proc: out: failed to wait: %s", p.Err())
		return nil
	}

	return p.output
}

// Result waits and copies the result of the combined stdout and stderr
// and returns its result as a string.
// Must be called after Proc.StartProc Proc.RunProc() or Proc.Run().
func (p *Proc) Result() string {
	if p.result != nil {
		return strings.TrimSpace(p.result.String())
	}

	p.result = new(bytes.Buffer)

	if _, err := p.result.ReadFrom(p.output); err != nil {
		p.err = err
		return ""
	}

	if err := p.Wait().Err(); err != nil {
		p.err = fmt.Errorf("proc: result: %s", p.Err())
		return ""
	}

	return strings.TrimSpace(p.result.String())
}
