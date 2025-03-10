//go:build windows

package exec

import "testing"

func TestCommandBuilder_Pipe(t *testing.T) {
	tests := []struct {
		name         string
		commands     []string
		results      []string
		expectedErrs int
		shellCmd     string
	}{
		{
			name:     "single command",
			commands: []string{`powershell.exe -c "Write-Host -NoNewline 'hello world'"`},
			results:  []string{"hello world"},
		},
		{
			name:     "single command, set shell",
			shellCmd: "powershell.exe -c",
			commands: []string{`Write-Host -NoNewline 'hello world'`},
			results:  []string{"hello world"},
		},
		{
			name:     "count characters",
			shellCmd: "powershell.exe -c",
			commands: []string{`Write-Output 'Hello World!'`, `Measure-Object -Character -Line`, `Select-Object -ExpandProperty Characters`},
			results:  []string{"11"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmds := Commands(test.commands...)
			if test.shellCmd != "" {
				cmds.WithShell(test.shellCmd)
			}

			c := cmds.Pipe()
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
