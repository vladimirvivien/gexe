package gexe

import (
	"strings"
	"testing"
)

func TestEchoRun(t *testing.T) {
	tests := []struct {
		name   string
		cmdStr string
		exec   func(*testing.T, string)
	}{
		{
			name:   "start proc",
			cmdStr: `echo "HELLO WORLD!"`,
			exec: func(t *testing.T, cmd string) {
				p := DefaultEcho.StartProc(cmd)
				if p.Err() != nil {
					t.Fatal("Unexpected error:", p.Err().Error())
				}
				if p.ExitCode() != -1 {
					t.Fatal("Expecting -1, got:", p.ExitCode())
				}
				if p.IsSuccess() {
					t.Fatal("Success should be false")
				}

				// wait for completion
				if err := p.Wait().Err(); err != nil {
					t.Fatalf("Failed to complete process: %s", err)
				}

				result := strings.TrimSpace(p.Result())
				if result != "HELLO WORLD!" {
					t.Errorf("Unexpected proc.Out(): %s", result)
				}

				if p.ExitCode() != 0 {
					t.Fatal("Expecting exit code 0, got:", p.ExitCode())
				}
				if !p.IsSuccess() {
					t.Fatal("Process should be success")
				}
			},
		},
		{
			name:   "start proc/long-running",
			cmdStr: `/bin/sh -c "for i in {1..3}; do echo 'HELLO WORLD!\$i'; sleep 0.2; done"`,
			exec: func(t *testing.T, cmd string) {
				p := DefaultEcho.StartProc(cmd)
				if p.Err() != nil {
					t.Fatal(p.Err())
				}

				if err := p.Wait().Err(); err != nil {
					t.Fatalf("proc failed to wait: %s", err)
				}

				if !strings.Contains(p.Result(), "HELLO WORLD!") {
					t.Fatal("Unexpected result:", p.Result())
				}
			},
		},
		{
			name:   "run proc",
			cmdStr: `echo "HELLO WORLD!"`,
			exec: func(t *testing.T, cmd string) {
				p := DefaultEcho.RunProc(cmd)
				if p.ExitCode() != 0 {
					t.Fatal("Expecting exit code 0, got:", p.ExitCode())
				}
				if !p.IsSuccess() {
					t.Fatal("Process should be success")
				}
				if p.Result() != "HELLO WORLD!" {
					t.Fatal("Unexpected command result:", p.Result())
				}
			},
		},
		{
			name:   "simple run",
			cmdStr: `echo "HELLO WORLD!"`,
			exec: func(t *testing.T, cmd string) {
				result := DefaultEcho.Run(cmd)
				if result != "HELLO WORLD!" {
					t.Fatal("Unexpected command result:", result)
				}
			},
		},

		{
			name:   "simple with expansion",
			cmdStr: "echo $MSG",
			exec: func(t *testing.T, cmd string) {
				DefaultEcho.Variables().Vars("MSG=Hello World")
				result := DefaultEcho.Run(cmd)
				if result != DefaultEcho.Variables().Val("MSG") {
					t.Fatal("Unexpected command result:", result)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.exec(t, test.cmdStr)
		})
	}
}
