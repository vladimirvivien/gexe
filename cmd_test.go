package echo

import "testing"

func TestEchoRun(t *testing.T) {
	tests := []struct {
		name   string
		cmdStr func() string
		exec   func(string)
	}{
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
