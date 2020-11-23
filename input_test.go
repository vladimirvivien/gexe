package echo

import (
	"fmt"
	"os"
	"testing"
)

func TestInput_Read(t *testing.T) {
	origStdin := os.Stdin
	stdinReader, stdinWriter, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer func(){os.Stdin = origStdin}()
	os.Stdin = stdinReader

	prompt := "Hello"
	input := "World!"

	if _, err := stdinWriter.WriteString(fmt.Sprintf("%s\n", input)); err != nil {
		t.Fatal(err)
	}

	in := New().Input(prompt)
	inline := in.Read('\n')

	if inline.Value() != input {
		t.Errorf("expecting %s, got %s", input, inline.Value())
	}
}

func TestInput_ReadLine(t *testing.T) {
	origStdin := os.Stdin
	stdinReader, stdinWriter, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer func(){os.Stdin = origStdin}()
	os.Stdin = stdinReader

	input := "World!"

	if _, err := stdinWriter.WriteString(fmt.Sprintf("%s\n", input)); err != nil {
		t.Fatal(err)
	}

	in := New().Input()
	inline := in.ReadLine()

	if inline.Value() != input {
		t.Errorf("expecting %s, got %s", input, inline.Value())
	}
}