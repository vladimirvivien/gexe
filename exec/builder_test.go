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
		results      []string
		expectedCmds int
		expectedErrs int
		policy       CommandPolicy
	}{
		{
			name: "zero procs",
		},
		{
			name:         "no error in procs",
			commands:     []string{"echo 'hello world'", `date "+%Y-%m-%d"`, "ls -al"},
			results:      []string{"hello world", Run(`date "+%Y-%m-%d"`), Run(`ls -al`)},
			expectedCmds: 3,
		},
		{
			name:         "continue on 1 error",
			commands:     []string{"foobar", "echo 'hello world'", `date "+%Y-%m-%d"`, "ls -al"},
			results:      []string{Run(`foobar`), "hello world", Run(`date "+%Y-%m-%d"`), Run(`ls -al`)},
			expectedCmds: 4,
			expectedErrs: 1,
		},
		{
			name:         "continue on 2 errors",
			commands:     []string{"foobar", "echo 'hello world'", "daftpunk", `date "+%Y-%m-%d"`, "ls -al"},
			results:      []string{Run(`foobar`), "hello world", Run(`daftpunk`), Run(`date "+%Y-%m-%d"`), Run(`ls -al`)},
			expectedCmds: 5,
			expectedErrs: 2,
		},
		{
			name:         "stop on errors",
			commands:     []string{"echo 'hello world'", "daftpunk", `date "+%Y-%m-%d"`, "ls -al"},
			results:      []string{"hello world", Run(`daftpunk`), Run(`date "+%Y-%m-%d"`), Run(`ls -al`)},
			expectedCmds: 2,
			expectedErrs: 1,
			policy:       ExitOnErrPolicy,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := Commands(test.commands...).WithPolicy(test.policy).Run()
			if len(c.Procs()) != test.expectedCmds {
				t.Errorf("expecting %d procs to run, got %d", test.expectedCmds, len(c.Procs()))
			}

			if len(c.ErrProcs()) != test.expectedErrs {
				t.Errorf("expecting %d procs errors, got %d", test.expectedErrs, len(c.ErrProcs()))
			}

			for i, p := range c.Procs() {
				if p.Result() != test.results[i] {
					t.Errorf("unexpected proc result: %s", p.Result())
				}
			}
		})
	}
}

func TestCommandBuilder_Start(t *testing.T) {
	tests := []struct {
		name         string
		commands     []string
		results      []string
		expectedCmds int
		expectedErrs int
		policy       CommandPolicy
	}{
		{
			name: "zero procs",
		},
		{
			name:         "no error in procs",
			commands:     []string{"echo 'hello world'", `date "+%Y-%m-%d"`, "ls -al"},
			results:      []string{"hello world", Run(`date "+%Y-%m-%d"`), Run(`ls -al`)},
			expectedCmds: 3,
		},
		{
			name:         "continue on 1 error",
			commands:     []string{"foobar", "echo 'hello world'", `date "+%Y-%m-%d"`, "ls -al"},
			results:      []string{Run(`foobar`), "hello world", Run(`date "+%Y-%m-%d"`), Run(`ls -al`)},
			expectedCmds: 4,
			expectedErrs: 1,
		},
		{
			name:         "continue on 2 errors",
			commands:     []string{"foobar", "echo 'hello world'", "daftpunk", `date "+%Y-%m-%d"`, "ls -al"},
			results:      []string{Run(`foobar`), "hello world", Run(`daftpunk`), Run(`date "+%Y-%m-%d"`), Run(`ls -al`)},
			expectedCmds: 5,
			expectedErrs: 2,
		},
		{
			name:         "stop on errors",
			commands:     []string{"echo 'hello world'", "daftpunk", `date "+%Y-%m-%d"`, "ls -al"},
			results:      []string{"hello world", Run(`daftpunk`), Run(`date "+%Y-%m-%d"`), Run(`ls -al`)},
			expectedCmds: 2,
			expectedErrs: 1,
			policy:       ExitOnErrPolicy,
		},

		// concurrent
		{
			name:         "concurrent no errors",
			commands:     []string{"echo 'hello world'", `date "+%Y-%m-%d"`, "ls -al"},
			results:      []string{"hello world", Run(`date "+%Y-%m-%d"`), Run(`ls -al`)},
			expectedCmds: 3,
			policy:       ConcurrentExecPolicy,
		},
		{
			name:         "concurrent on 1 error",
			commands:     []string{"foobar", "echo 'hello world'", `date "+%Y-%m-%d"`, "ls -al"},
			results:      []string{Run(`foobar`), "hello world", Run(`date "+%Y-%m-%d"`), Run(`ls -al`)},
			expectedCmds: 4,
			expectedErrs: 1,
			policy:       ConcurrentExecPolicy,
		},
		{
			name:         "concurrent on 2 errors",
			commands:     []string{"foobar", "echo 'hello world'", "daftpunk", `date "+%Y-%m-%d"`, "ls -al"},
			results:      []string{Run(`foobar`), "hello world", Run(`daftpunk`), Run(`date "+%Y-%m-%d"`), Run(`ls -al`)},
			expectedCmds: 5,
			expectedErrs: 2,
			policy:       ConcurrentExecPolicy,
		},
		{
			name:         "concurrent on errors",
			commands:     []string{"echo 'hello world'", "daftpunk", `date "+%Y-%m-%d"`, "ls -al"},
			results:      []string{"hello world", Run(`daftpunk`), Run(`date "+%Y-%m-%d"`), Run(`ls -al`)},
			expectedCmds: 4,
			expectedErrs: 1,
			policy:       ExitOnErrPolicy | ConcurrentExecPolicy, // ExitOnErr is ignored when concurrent
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Commands(test.commands...).WithPolicy(test.policy).Start()
			if len(result.Procs()) != 0 {
				t.Fatal("expecting 0 completed processes at this point")
			}
			result = result.Wait()
			if len(result.Procs()) != test.expectedCmds {
				t.Errorf("expecting %d procs to run, got %d", test.expectedCmds, len(result.Procs()))
			}

			if len(result.ErrProcs()) != test.expectedErrs {
				t.Errorf("expecting %d procs errors, got %d", test.expectedErrs, len(result.ErrProcs()))
			}

			for i, p := range result.Procs() {
				if p.Result() != test.results[i] {
					t.Errorf("unexpected proc result: %s", p.Result())
				}
			}
		})
	}
}

func TestCommandBuilder_Pipe(t *testing.T) {
	tests := []struct {
		name         string
		commands     []string
		results      []string
		expectedCmds int
		expectedErrs int
	}{
		{
			name:         "one command",
			commands:     []string{"echo 'hello world'"},
			results:      []string{"hello world"},
			expectedCmds: 1,
		},
		{
			name:         "two commands",
			commands:     []string{"echo -n 'hello world'", "wc -m"},
			results:      []string{"11"},
			expectedCmds: 2,
		},
		{
			name:         "three commands",
			commands:     []string{"echo -n 'hello world'", "grep world", "wc -w"},
			results:      []string{"2"},
			expectedCmds: 3,
		},
		{
			name:         "three commands with 1 err",
			commands:     []string{"echo -n 'hello world'", "foo", "wc -w"},
			results:      []string{"2"},
			expectedCmds: 2,
			expectedErrs: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := Commands(test.commands...).Pipe()
			if len(c.Procs()) != test.expectedCmds {
				t.Errorf("expecting %d procs to run, got %d", test.expectedCmds, len(c.Procs()))
			}

			if test.expectedErrs != len(c.ErrProcs()) {
				t.Fatalf("expecting %d errors, but got %d", test.expectedErrs, len(c.ErrProcs()))
			}
			// for pipe, only check last result
			p := c.LastProc()
			if p.Err() != nil && test.expectedErrs == 0 {
				t.Fatalf("last proc in pipe failed: %s", p.Err())
			}
			if p.Err() == nil && p.Result() != test.results[0] {
				t.Errorf("unexpected proc result: %s", p.Result())
			}
		})
	}
}
