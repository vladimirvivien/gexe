package echo

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Run parses and executes command cmdStr and returns the result
// from stdout or stderr
func (e *Echo) Run(cmdStr string) string {
	cmdStr = lineRgx.ReplaceAllString(cmdStr, " ")
	if e.Conf.isVerbose() {
		fmt.Println(cmdStr)
	}
	words := e.splitWords(e.Eval(cmdStr))
	return cmdRun(words[0], words[1:]...)
}

// Runout parses and executes command cmdStr and prints out the result
func (e *Echo) Runout(cmdStr string) {
	fmt.Print(os.Expand(e.Run(cmdStr), e.Val))
}

func parseCmdStr(cmdStr string) (cmdName string, args []string) {
	args = []string{}
	parts := spaceRgx.Split(cmdStr, -1)
	if len(parts) == 0 {
		return
	}
	if len(parts) == 1 {
		cmdName = parts[0]
		return
	}
	cmdName = parts[0]
	args = parts[1:]
	return
}

func prepCmd(cmd string, args ...string) (*exec.Cmd, *bytes.Buffer) {
	output := new(bytes.Buffer)
	command := exec.Command(cmd, args...)
	command.Stdout = output
	command.Stderr = output
	return command, output
}

func cmdRun(cmd string, args ...string) string {
	command, output := prepCmd(cmd, args...)
	if err := command.Run(); err != nil {
		return fmt.Sprintf("command error: %s", err)
	}
	return strings.TrimSpace(output.String())
}
