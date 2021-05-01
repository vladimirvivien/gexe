package exec

import (
	"bytes"
	"strings"
	"testing"

	"github.com/vladimirvivien/echo/vars"
)

func TestProc(t *testing.T) {
	tests := []struct {
		name   string
		cmdStr string
		exec   func(string)
	}{
		{
			name: "start proc",
			cmdStr: `echo "HELLO WORLD!"`,
			exec: func(cmd string) {
				p := StartProc(cmd)
				if p.Err() != nil {
					t.Fatal("Unexpected error:", p.Err().Error())
				}
				if p.ExitCode() != -1 {
					t.Fatal("Expecting -1, got:", p.ExitCode())
				}
				if p.IsSuccess() {
					t.Fatal("Success should be false")
				}

				data := &bytes.Buffer{}
				if _, err := data.ReadFrom(p.StdOut()); err != nil {
					t.Fatal(err)
				}

				if err := p.Wait().Err(); err != nil {
					t.Fatal(err)
				}

				result := strings.TrimSpace(data.String())
				if result != "HELLO WORLD!" {
					t.Errorf("Unexpected proc.Result(): %s", result)
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
			name: "start proc/long-running",
			cmdStr: `/bin/sh -c "for i in {1..3}; do echo 'HELLO WORLD!\$i'; sleep 0.2; done"`,
			exec: func(cmd string) {
				p := StartProc(cmd)
				if p.Err() != nil {
					t.Fatal(p.Err())
				}

				data := &bytes.Buffer{}
				if _, err := data.ReadFrom(p.StdOut()); err != nil {
					t.Fatal(err)
				}

				if err := p.Wait().Err(); err != nil {
					t.Fatal(err)
				}

				result := strings.TrimSpace(data.String())
				lines := strings.Split(result, "\n")
				if len(lines) != 3 {
					t.Fatal("unexpected lines generated:", len(lines))
				}
				if !strings.Contains(result, "HELLO WORLD!") {
					t.Fatal("Unexpected result:", result)
				}
			},
		},
		{
			name: "run proc",
			cmdStr: `echo "HELLO WORLD!"`,
			exec: func(cmd string) {
				p := RunProc(cmd)
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
			name: "simple run",
			cmdStr: `echo "HELLO WORLD!"`,
			exec: func(cmd string) {
				result := Run(cmd)
				if result != "HELLO WORLD!" {
					t.Fatal("Unexpected command result:", result)
				}
			},
		},
		{
			name: "simple with expansion",
			cmdStr: "echo $MSG",
			exec: func(cmd string) {
				v := vars.New().Vars("MSG=Hello World")
				result := Run(v.Eval(cmd))
				if result != v.Val("MSG") {
					t.Fatal("Unexpected command result:", result)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.exec(test.cmdStr)
		})
	}
}
