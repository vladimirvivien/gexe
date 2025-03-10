//go:build windows

package exec

import (
	"fmt"
	"testing"
)

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
			name: "no error in procs",
			commands: []string{
				`powershell.exe -c "Write-Host 'hello world'"`,
				`powershell.exe -c "Get-Date -UFormat '+%Y-%m-%d'"`,
				`powershell.exe -c "dir"`,
			},
			results: []string{
				"hello world",
				Run(`powershell.exe -c "Get-Date -UFormat '+%Y-%m-%d'"`),
				Run(`powershell.exe -c "dir"`),
			},
			expectedCmds: 3,
		},
		{
			name: "continue on 1 error",
			commands: []string{
				`foobar`,
				`powershell.exe -c "Write-Host 'hello world'"`,
				`powershell.exe -c "Get-Date -UFormat '+%Y-%m-%d'"`,
				`powershell.exe -c "dir"`,
			},
			results: []string{
				Run(`foobar`),
				"hello world",
				Run(`powershell.exe -c "Get-Date -UFormat '+%Y-%m-%d'"`),
				Run(`powershell.exe -c "dir"`),
			},
			expectedCmds: 4,
			expectedErrs: 1,
		},
		{
			name: "continue on 2 error",
			commands: []string{
				`foobar`,
				`powershell.exe -c "Write-Host 'hello world'"`,
				`daftpunk`,
				`powershell.exe -c "Get-Date -UFormat '+%Y-%m-%d'"`,
				`powershell.exe -c "dir"`,
			},
			results: []string{
				Run(`foobar`),
				"hello world",
				Run(`daftpunk`),
				Run(`powershell.exe -c "Get-Date -UFormat '+%Y-%m-%d'"`),
				Run(`powershell.exe -c "dir"`),
			},
			expectedCmds: 5,
			expectedErrs: 2,
		},
		{
			name: "exit on 1 error",
			commands: []string{
				`powershell.exe -c "Write-Host 'hello world'"`,
				`daftpunk`,
				`powershell.exe -c "Get-Date -UFormat '+%Y-%m-%d'"`,
				`powershell.exe -c "dir"`,
			},
			results: []string{
				"hello world",
				Run(`daftpunk`),
				Run(`powershell.exe -c "Get-Date -UFormat '+%Y-%m-%d'"`),
				Run(`powershell.exe -c "dir"`),
			},
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
				fmt.Printf("c.ErrProcs(): %v\n", c.ErrProcs()[0].Result())
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
			name: "no error in procs",
			commands: []string{
				`powershell.exe -c "Write-Host 'hello world'"`,
				`powershell.exe -c "Get-Date -UFormat '+%Y-%m-%d'"`,
				`powershell.exe -c "dir"`,
			},
			results: []string{
				"hello world",
				Run(`powershell.exe -c "Get-Date -UFormat '+%Y-%m-%d'"`),
				Run(`powershell.exe -c "dir"`),
			},
			expectedCmds: 3,
		},
		{
			name: "continue on 1 error",
			commands: []string{
				`foobar`,
				`powershell.exe -c "Write-Host 'hello world'"`,
				`powershell.exe -c "Get-Date -UFormat '+%Y-%m-%d'"`,
				`powershell.exe -c "dir"`,
			},
			results: []string{
				Run(`foobar`),
				"hello world",
				Run(`powershell.exe -c "Get-Date -UFormat '+%Y-%m-%d'"`),
				Run(`powershell.exe -c "dir"`),
			},
			expectedCmds: 4,
			expectedErrs: 1,
		},
		{
			name: "stop on errors",
			commands: []string{
				`powershell.exe -c "Write-Host 'hello world'"`,
				`daftpunk`,
				`powershell.exe -c "Get-Date -UFormat '+%Y-%m-%d'"`,
				`powershell.exe -c "dir"`,
			},
			results: []string{
				"hello world",
				Run(`daftpunk`),
				Run(`powershell.exe -c "Get-Date -UFormat '+%Y-%m-%d'"`),
				Run(`powershell.exe -c "dir"`),
			},
			expectedCmds: 2,
			expectedErrs: 1,
			policy:       ExitOnErrPolicy,
		},

		// concurrent
		{
			name: "concurrent no errors",
			commands: []string{
				`powershell.exe -c "Write-Host 'hello world'"`,
				`powershell.exe -c "Get-Date -UFormat '+%Y-%m-%d'"`,
				`powershell.exe -c "dir"`,
			},
			results: []string{
				"hello world",
				Run(`powershell.exe -c "Get-Date -UFormat '+%Y-%m-%d'"`),
				Run(`powershell.exe -c "dir"`),
			},
			expectedCmds: 3,
			policy:       ConcurrentExecPolicy,
		},
		{
			name: "concurrent on 1 error",
			commands: []string{
				`foobar`,
				`powershell.exe -c "Write-Host 'hello world'"`,
				`powershell.exe -c "Get-Date -UFormat '+%Y-%m-%d'"`,
				`powershell.exe -c "dir"`,
			},
			results: []string{
				Run(`foobar`),
				"hello world",
				Run(`powershell.exe -c "Get-Date -UFormat '+%Y-%m-%d'"`),
				Run(`powershell.exe -c "dir"`),
			},
			expectedCmds: 4,
			expectedErrs: 1,
			policy:       ConcurrentExecPolicy,
		},
		{
			name: "concurrent on 2 errors",
			commands: []string{
				`foobar`,
				`powershell.exe -c "Write-Host 'hello world'"`,
				`daftpunk`,
				`powershell.exe -c "Get-Date -UFormat '+%Y-%m-%d'"`,
				`powershell.exe -c "dir"`,
			},
			results: []string{
				Run(`foobar`),
				"hello world",
				Run(`daftpunk`),
				Run(`powershell.exe -c "Get-Date -UFormat '+%Y-%m-%d'"`),
				Run(`powershell.exe -c "dir"`),
			},
			expectedCmds: 5,
			expectedErrs: 2,
			policy:       ConcurrentExecPolicy,
		},
		{
			name: "concurrent on errors",
			commands: []string{
				`powershell.exe -c "Write-Host 'hello world'"`,
				`daftpunk`,
				`powershell.exe -c "Get-Date -UFormat '+%Y-%m-%d'"`,
				`powershell.exe -c "dir"`,
			},
			results: []string{
				"hello world",
				Run(`daftpunk`),
				Run(`powershell.exe -c "Get-Date -UFormat '+%Y-%m-%d'"`),
				Run(`powershell.exe -c "dir"`),
			},
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
