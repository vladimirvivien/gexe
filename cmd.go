package echo

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func (e *echo) Run(cmdStr string) string {
	cmdStr = lineRgx.ReplaceAllString(cmdStr, " ")
	cmdName, args := parseCmdStr(os.Expand(cmdStr, e.Val))
	return cmdRun(cmdName, args...)
}

func (e *echo) Runout(cmdStr string) {
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
