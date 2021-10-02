package exec

import (
	"testing"
)

func TestCommandBuilder(t *testing.T) {
	tests := []struct {
		name           string
		initCmds       []string
		additionalCmds []string
	}{
		{name: "no procs"},
		{
			name:     "initial procs only",
			initCmds: []string{"echo 'hello world'", "date", "ls -al"},
		},
		{
			name:           "initial and one additional",
			initCmds:       []string{"echo 'hello world'", "date", "ls -al"},
			additionalCmds: []string{"git commit --signoff"},
		},
		{
			name:           "initial and multiple additional",
			initCmds:       []string{"echo 'hello world'", "date", "ls -al"},
			additionalCmds: []string{"git commit --signoff", "history", "man time", "man man"},
		},
		{
			name:           "no initial multiple additional",
			additionalCmds: []string{"git commit --signoff", "history", "man time", "man man"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := Commands(test.initCmds...)
			if len(test.initCmds) != len(c.procs) {
				t.Error("unexpected command count in CommandBuilder")
			}
			c.Add(test.additionalCmds...)
			if (len(test.initCmds) + len(test.additionalCmds)) != len(c.procs) {
				t.Error("procs are not added to builder properly")
			}
		})
	}
}

func TestCommandBuilder_Run(t *testing.T) {
	tests := []struct {
		name         string
		commands     []string
		expectedCmds int
	}{
		{
			name: "zero procs",
		},
		{
			name:         "no error in procs",
			commands:     []string{"echo 'hello world'", "date", "ls -al"},
			expectedCmds: 3,
		},
		{
			name:         "continue on 1 error",
			commands:     []string{"foobar", "echo 'hello world'", "date", "ls -al"},
			expectedCmds: 4,
		},
		{
			name:         "continue on 2 errors",
			commands:     []string{"foobar", "echo 'hello world'", "daftpunk", "date", "ls -al"},
			expectedCmds: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := Commands(test.commands...).Run()
			if len(c.procs) != test.expectedCmds {
				t.Errorf("expecting %d procs to run, got %d", test.expectedCmds, len(c.procs))
			}
		})
	}
}

func TestCommandBuilder_ConcurRun(t *testing.T) {
	tests := []struct {
		name         string
		commands     []string
		expectedCmds int
	}{
		{
			name: "zero procs",
		},
		{
			name:         "no error in procs",
			commands:     []string{"echo 'hello world'", "date", "ls -al"},
			expectedCmds: 3,
		},
		{
			name:         "continue on 1 error",
			commands:     []string{"foobar", "echo 'hello world'", "date", "ls -al"},
			expectedCmds: 4,
		},
		{
			name:         "continue on 2 errors",
			commands:     []string{"foobar", "echo 'hello world'", "daftpunk", "date", "ls -al"},
			expectedCmds: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := Commands(test.commands...).ConcurRun()
			if len(c.procs) != test.expectedCmds {
				t.Errorf("expecting %d procs to run, got %d", test.expectedCmds, len(c.procs))
			}
		})
	}
}

func TestCommandBuilder_StartWait(t *testing.T) {
	tests := []struct {
		name         string
		commands     []string
		policy       CommandPolicy
		expectedCmds int
	}{
		{
			name: "zero procs",
		},
		{
			name:         "no error in procs",
			commands:     []string{"echo 'hello world'", "date", "ls -al"},
			policy:       CmdOnErrContinue,
			expectedCmds: 3,
		},
		{
			name:         "continue on 1 error",
			commands:     []string{"foobar", "echo 'hello world'", "date", "ls -al"},
			policy:       CmdOnErrContinue,
			expectedCmds: 4,
		},
		{
			name:         "break on 1 error",
			commands:     []string{"foobar", "echo 'hello world'", "date", "ls -al"},
			policy:       CmdOnErrExit,
			expectedCmds: 1,
		},
		{
			name:         "continue on 2 errors",
			commands:     []string{"foobar", "echo 'hello world'", "daftpunk", "date", "ls -al"},
			policy:       CmdOnErrContinue,
			expectedCmds: 5,
		},
		{
			name:         "break on 2 errors",
			commands:     []string{"foobar", "echo 'hello world'", "daftpunk", "date", "ls -al"},
			policy:       CmdOnErrExit,
			expectedCmds: 1,
		},
		{
			name:         "concurrently no errors",
			commands:     []string{"echo 'hello world'", "date", "ls -al"},
			policy:       CmdExecConcurrent,
			expectedCmds: 3,
		},
		{
			name:         "concurrent 1 error",
			commands:     []string{"foobar", "echo 'hello world'", "date", "ls -al"},
			policy:       CmdExecConcurrent,
			expectedCmds: 4,
		},
		{
			name:         "continue on 2 errors",
			commands:     []string{"foobar", "echo 'hello world'", "daftpunk", "date", "ls -al"},
			policy:       CmdExecConcurrent,
			expectedCmds: 5,
		},
		{
			name:         "Concurr|Continue with 1 err",
			commands:     []string{"man cat", "echo 'hello world'", "foo", "ls -al"},
			policy:       CmdOnErrContinue | CmdExecConcurrent,
			expectedCmds: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := Commands(test.commands...).WithPolicy(test.policy).Start().Wait()
			if len(c.procs) != test.expectedCmds {
				t.Errorf("expecting %d procs to run, got %d", test.expectedCmds, len(c.procs))
			}
		})
	}
}
