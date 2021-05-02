package prog

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// ProgInfo returns information about the
// running program.
type ProgInfo struct {
	err error
}


func Prog() *ProgInfo {
	return &ProgInfo{}
}

// Args returns a slice of the program arguments
func (p *ProgInfo) Args() []string {
	return os.Args
}

// Err returns the last generated error from a method call
func (p *ProgInfo) Err() error {
	return p.err
}

// Exit prints messages and exits current program
func (p *ProgInfo) Exit(code int, msgs ...string) {
	for _, msg := range msgs {
		fmt.Print(msg)
	}
	os.Exit(code)
}

// Pid program's process id
func (p *ProgInfo) Pid() int {
	return os.Getpid()
}

// Ppid program's parent process id
func (p *ProgInfo) Ppid() int {
	return os.Getppid()
}

// Path of running program
func (p *ProgInfo) Path() string {
	path, err := os.Executable()
	if err != nil {
		p.err = err
		return ""
	}
	return path
}

// Name of executable running
func (p *ProgInfo) Name() string {
	return filepath.Base(p.Path())
}

// Avail returns full path of binary name if available
func (p *ProgInfo) Avail(progName string) string {
	path, err := exec.LookPath(progName)
	if err != nil {
		p.err = err
		return ""
	}
	return path
}

// Workdir returns the working directory
func (p *ProgInfo) Workdir() string {
	path, err := os.Getwd()
	if err != nil {
		p.err = err
		return ""
	}
	return path
}
