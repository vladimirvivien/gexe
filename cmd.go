package echo

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// RunProc parses and executes cmdStr and returns a process result.
// Use this to have access to process result informoation.
func (e *Echo) RunProc(cmdStr string) *ProcResult {
	cmdStr = lineRgx.ReplaceAllString(cmdStr, " ")
	e.shouldLog(cmdStr)

	proc := e.StartProc(cmdStr)
	proc.Wait()
	return proc
}

// StartProc concurrently starts a process
func (e *Echo) StartProc(cmdStr string) *ProcResult {
	return e.startProc(cmdStr)
}

// Run parses and executes cmdStr and returns the result as a string.
// This uses RunProc and extracts the result as a string.
func (e *Echo) Run(cmdStr string) string {
	p := e.RunProc(cmdStr)
	if p.Err() != nil {
		e.shouldLog(p.Err().Error())
		e.shouldPanic(p.Err().Error())
		return ""
	}

	data, err := ioutil.ReadAll(p.Out())
	if err != nil {
		e.shouldLog(err.Error())
		e.shouldPanic(err.Error())
		return ""
	}

	return os.Expand(strings.TrimSpace(string(data)), e.Val)
}

// Runout parses and executes command cmdStr and prints out the result
func (e *Echo) Runout(cmdStr string) {
	fmt.Print(e.Run(cmdStr))
}

func (e *Echo) startProc(cmdStr string) *ProcResult {
	words := e.splitWords(e.Eval(cmdStr))

	output := new(bytes.Buffer)
	command := exec.Command(words[0], words[1:]...)
	command.Stdout = output
	command.Stderr = output

	if err := command.Start(); err != nil {
		e.Proc = &ProcResult{cmd: command, err: err, state: command.ProcessState}
		return e.Proc
	}

	e.Proc = &ProcResult{cmd: command, state: command.ProcessState, output: output}
	return e.Proc
}
