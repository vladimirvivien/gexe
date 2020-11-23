package echo

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Input struct {
	prompt string
}

// Input sets up standard input with {0 or up to 1} prompt value.
func (e *Echo) Input(prompt ...string) *Input {
	var promptVal string
	if len(prompt) > 0 {
		promptVal = prompt[0]
	}
	return &Input{prompt: promptVal}
}

// Read prints prompt (if any), reads input from stdin until delim
func (i *Input) Read(delim byte) *InputLine {
	if len(i.prompt) > 0 {
		fmt.Print(i.prompt)
	}
	rdr := bufio.NewReader(os.Stdin)
	val, err := rdr.ReadString(delim)
	if err != nil && err != io.EOF{
		return &InputLine{err: err}
	}
	return &InputLine{val: strings.TrimSpace(val)}
}

// ReadLine prints prompt (if any), reads input from stdin until \n
func (i *Input) ReadLine() *InputLine {
	return i.Read	('\n')
}

// InputLine captures and stores a single val from stdin
type InputLine struct {
	err error
	val string
}

func (input *InputLine) Value() string {
	return input.val
}

func (input *InputLine) Err() error {
	return input.err
}

func (i *Input) ReadLineQuietly() *InputLine {
	return i.ReadLine()
}

func (i *Input) ReadLines() *InputLines {
	return &InputLines{}
}

func (i *Input) ReadLinesQuietly() *InputLines {
	return i.ReadLines()
}

type InputLines struct {
}

// ReadLine prints a prompt and accepts input from stdin
