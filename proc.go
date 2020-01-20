package echo

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Proc stores process info
type Proc struct {
	id     int
	err    error
	state  *os.ProcessState
	output io.Reader
	cmd    *exec.Cmd
}

// Peek attempts to read process state information
func (p *Proc) Peek() *Proc {
	p.state = p.cmd.ProcessState
	return p
}

// Wait for associated process to complete.
// Any error can be accessed via p.Err()
func (p *Proc) Wait() *Proc {
	if p.cmd == nil {
		return p
	}
	if err := p.cmd.Wait(); err != nil {
		p.err = err
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

// Out surfaces an io.Reader for proc result
func (p *Proc) Out() io.Reader {
	return p.output
}

// Result streams output into a string
func (p *Proc) Result() string {
	data, err := ioutil.ReadAll(p.Out())
	if err != nil {
		p.err = err
		return ""
	}

	return strings.TrimSpace(string(data))
}
