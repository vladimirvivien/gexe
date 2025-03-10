//go:build !windows

package exec

import "testing"

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
				t.Logf("Errors: %#v", c.ErrStrings())
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
