package echo

import (
	"io"
	"os"
	"os/exec"
	"time"
)

// ProcResult stores process info
type ProcResult struct {
	err    error
	state  *os.ProcessState
	output io.Reader
	cmd    *exec.Cmd
}

// Wait for associated process to complete.
// Any error can be accessed via p.Err()
func (p *ProcResult) Wait() {
	if p.cmd == nil {
		return
	}
	if err := p.cmd.Wait(); err != nil {
		p.err = err
	}
}

// ID returns process id
func (p *ProcResult) ID() int {
	if p.state == nil {
		return -1
	}
	return p.state.Pid()
}

// Exited returns true if process exits ok
func (p *ProcResult) Exited() bool {
	if p.state == nil {
		return false
	}
	return p.state.Exited()
}

// ExitCode returns process exit code
func (p *ProcResult) ExitCode() int {
	if p.state == nil {
		return -1
	}
	return p.state.ExitCode()
}

// IsSuccess returns true if proc exit ok
func (p *ProcResult) IsSuccess() bool {
	return p.state.Success()
}

// SysTime returns proc system cpu time
func (p *ProcResult) SysTime() time.Duration {
	return p.state.SystemTime()
}

// UserTime returns proc user cpu time
func (p *ProcResult) UserTime() time.Duration {
	return p.state.UserTime()
}

// Err returns any execution error
func (p *ProcResult) Err() error {
	return p.err
}

// Out surfaces an io.Reader for proc result
func (p *ProcResult) Out() io.Reader {
	return p.output
}
