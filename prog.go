package echo

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type prog struct {
	err error
}

func (p *prog) Args() []string {
	return os.Args
}

// Err returns last generated error for Prog method call
func (p *prog) Err() error {
	return p.err
}

// Exit prints optional message and exits current program
func (p *prog) Exit(code int, msgs ...string) {
	for _, msg := range msgs {
		fmt.Print(msg)
	}
	os.Exit(code)
}

// Pid program's process id
func (p *prog) Pid() int {
	return os.Getpid()
}

// Ppid program's parent process id
func (p *prog) Ppid() int {
	return os.Getppid()
}

// Path of running program
func (p *prog) Path() string {
	path, err := os.Executable()
	if err != nil {
		p.err = err
		return ""
	}
	return path
}

// Name of executable running
func (p *prog) Name() string {
	return filepath.Base(p.Path())
}

// Avail returns full path of binary name if available
func (p *prog) Avail(progName string) string {
	path, err := exec.LookPath(progName)
	if err != nil {
		p.err = err
		return ""
	}
	return path
}

// Workdir returns the working directory
func (p *prog) Workdir() string {
	path, err := os.Getwd()
	if err != nil {
		p.err = err
		return ""
	}
	return path
}
