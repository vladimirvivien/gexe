package echo

import "testing"

func TestEchoRun(t *testing.T) {
	tests := []struct {
		name   string
		cmdStr func() string
		exec   func(string)
	}{
		{
			name: "start proc",
			cmdStr: func() string {
				return `echo "HELLO WORLD!"`
			},
			exec: func(cmd string) {
				e := New()
				p := e.StartProc(cmd)
				if p.state != nil {
					t.Fatal("state should not be set yet")
				}
				if p.Err() != nil {
					t.Fatal("Unexpected error:", p.Err().Error())
				}
				if p.ExitCode() != -1 {
					t.Fatal("Expecting -1, got:", p.ExitCode())
				}
				if p.IsSuccess() {
					t.Fatal("Success should be false")
				}
				p.Wait()
				if p.state == nil {
					t.Fatal("state should be set after Wait()", p.cmd.ProcessState)
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
			name: "run proc",
			cmdStr: func() string {
				return `echo "HELLO WORLD!"`
			},
			exec: func(cmd string) {
				e := New()
				p := e.RunProc(cmd)
				if p.state == nil {
					t.Fatal("state should be set after Wait()", p.cmd.ProcessState)
				}
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
			cmdStr: func() string {
				return `echo "HELLO WORLD!"`
			},
			exec: func(cmd string) {
				e := New()
				result := e.Run(cmd)
				if result != "HELLO WORLD!" {
					t.Fatal("Unexpected command result:", result)
				}
			},
		},

		{
			name: "simple multi-line",
			cmdStr: func() string {
				return `echo 
				"HELLO WORLD!"
				`
			},
			exec: func(cmd string) {
				e := New()
				result := e.Run(cmd)
				if result != "HELLO WORLD!" {
					t.Fatal("Unexpected command result:", result)
				}
			},
		},

		{
			name: "simple with expansion",
			cmdStr: func() string {
				return "echo $MSG"
			},
			exec: func(cmd string) {
				e := New().Var("MSG=Hello World")
				result := e.Run(cmd)
				if result != e.Val("MSG") {
					t.Fatal("Unexpected command result:", result)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.exec(test.cmdStr())
		})
	}
}
