package echo

import (
	"testing"

	"github.com/vladimirvivien/echo/exec"
)

func TestEchoProc(t *testing.T) {
	tests := []struct {
		name    string
		cmdProc func() (*Echo, *exec.Proc)
		exec    func(*Echo, *exec.Proc)
	}{
		{
			name: "start proc, no wait",
			cmdProc: func() (*Echo, *exec.Proc) {
				e := New()
				p := e.StartProc(`echo "HELLO WORLD!"`)
				return e, p
			},
			exec: func(e *Echo, p *exec.Proc) {
				if p.Err() != nil {
					t.Fatal("Unexpected error:", p.Err().Error())
				}
				if len(e.Procs) < 1 {
					t.Errorf("expecting at least 1 proc in echo, got %d", len(e.Procs))
				}
				if p.ID() <= 0 {
					t.Errorf("process id may not be valid: %d", p.ID())
				}
				if p.ExitCode() != -1 {
					t.Fatal("Expecting -1, got:", p.ExitCode())
				}
				if p.IsSuccess() {
					t.Fatal("Success should be false")
				}
			},
		},
		{
			name: "start proc, with wait",
			cmdProc: func() (*Echo, *exec.Proc) {
				e := New()
				p := e.StartProc(`echo "HELLO WORLD!"`)
				p.Wait()
				return e, p
			},
			exec: func(e *Echo, p *exec.Proc) {
				if p.Err() != nil {
					t.Fatal("Unexpected error:", p.Err().Error())
				}
				if len(e.Procs) < 1 {
					t.Errorf("expecting at least 1 proc in echo, got %d", len(e.Procs))
				}
				if p.ID() <= 0 {
					t.Errorf("process id may not be valid: %d", p.ID())
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
			cmdProc: func() (*Echo, *exec.Proc) {
				e := New()
				p := e.RunProc(`echo "HELLO WORLD!"`)
				return e, p
			},
			exec: func(e *Echo, p *exec.Proc) {
				if p.Err() != nil {
					t.Fatal("Unexpected error:", p.Err().Error())
				}
				if len(e.Procs) < 1 {
					t.Errorf("expecting at least 1 proc in echo, got %d", len(e.Procs))
				}
				if p.ID() <= 0 {
					t.Errorf("process id may not be valid: %d", p.ID())
				}
				if p.ExitCode() != 0 {
					t.Fatal("Expecting exit code 0, got:", p.ExitCode())
				}
				if !p.IsSuccess() {
					t.Fatal("Process should be success")
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.exec(test.cmdProc())
		})
	}
}
