package exec

import (
	"bytes"
	"strings"
	"testing"

	"github.com/vladimirvivien/gexe/vars"
)

func TestNewProc(t *testing.T) {
	tests := []struct {
		name   string
		cmdStr string
		exec   func(string)
	}{
		{
			name:   "start proc directly",
			cmdStr: `echo "Hello World"`,
			exec: func(cmd string) {
				proc := NewProc(cmd)
				result := new(bytes.Buffer)
				proc.SetStdout(result)
				if err := proc.Start().Err(); err != nil {
					t.Fatal(err)
				}
				if err := proc.Wait().Err(); err != nil {
					t.Fatal(err)
				}
				if strings.TrimSpace(result.String()) != "Hello World" {
					t.Errorf("unexpected result: %s", result.String())
				}
			},
		},
		{
			name:   "new proc with proc.Out()",
			cmdStr: `echo "Hello World"`,
			exec: func(cmd string) {
				proc := NewProc(cmd)
				buf := new(bytes.Buffer)
				if _, err := buf.ReadFrom(proc.Out()); err != nil {
					t.Error(err)
				}
				if strings.TrimSpace(buf.String()) != "Hello World" {
					t.Errorf("unexpected result: %s", buf.String())
				}
			},
		},
		{
			name:   "new proc with proc.Run()",
			cmdStr: `echo "Hello World"`,
			exec: func(cmd string) {
				proc := NewProc(cmd)
				result := new(bytes.Buffer)
				proc.SetStdout(result)
				if err := proc.Run().Err(); err != nil {
					t.Fatal(err)
				}
				if strings.TrimSpace(result.String()) != "Hello World" {
					t.Errorf("unexpected result: %s", result.String())
				}
			},
		},
		{
			name:   "proc.Run() with error",
			cmdStr: `foobar "Hello World"`,
			exec: func(cmd string) {
				proc := NewProc(cmd)
				result := new(bytes.Buffer)
				proc.SetStderr(result)
				if len(result.String()) > 0 {
					t.Error("expecting error string from stderr")
				}
			},
		},
		{
			name:   "new proc with proc.Result()",
			cmdStr: `echo "Hello World"`,
			exec: func(cmd string) {
				proc := NewProc(cmd)
				if strings.TrimSpace(proc.Result()) != "Hello World" {
					t.Errorf("unexpected result: %s", proc.Result())
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

func TestStartProc(t *testing.T) {
	tests := []struct {
		name   string
		cmdStr string
		exec   func(string)
	}{
		{
			name:   "start proc out",
			cmdStr: `echo "Hello World"`,
			exec: func(cmd string) {
				proc := StartProc(cmd)

				if err := proc.Wait().Err(); err != nil {
					t.Fatalf("failed to wait: %s", err)
				}

				buf := new(bytes.Buffer)

				if _, err := buf.ReadFrom(proc.Out()); err != nil {
					t.Error(err)
				}
				if strings.TrimSpace(buf.String()) != "Hello World" {
					t.Errorf("unexpected result: %s", buf.String())
				}
			},
		},
		{
			name:   "start proc with proc.Result()",
			cmdStr: `echo "Hello World"`,
			exec: func(cmd string) {
				proc := StartProc(cmd).Wait()
				if err := proc.Err(); err != nil {
					t.Fatalf("failed to start/wait proc: %s", err)
				}
				if strings.TrimSpace(proc.Result()) != "Hello World" {
					t.Errorf("unexpected result: %s", proc.Result())
				}
			},
		},
		{
			name:   "test proc status",
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
				if !p.hasStarted() {
					t.Fatal("proc has not started")
				}

				// wait for proc to finish
				if err := p.Wait().Err(); err != nil {
					t.Fatalf("failed to wait for proc to finish: %s", err)
				}

				data := &bytes.Buffer{}
				out := p.Out()
				if _, err := data.ReadFrom(out); err != nil {
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
			name:   "start proc/long-running",
			cmdStr: `/bin/bash -c 'for i in {1..3}; do echo "HELLO WORLD!"; sleep 0.7; done'`,
			exec: func(cmd string) {
				p := StartProc(cmd).Wait()

				if p.Err() != nil {
					t.Fatal(p.Err())
				}

				data := &bytes.Buffer{}
				if _, err := data.ReadFrom(p.Out()); err != nil {
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
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.exec(test.cmdStr)
		})
	}
}

func TestRunProc(t *testing.T) {
	tests := []struct {
		name   string
		cmdStr string
		exec   func(string)
	}{
		{
			name:   "run proc with proc.Out()",
			cmdStr: `echo "HELLO WORLD!"`,
			exec: func(cmd string) {
				p := RunProc(cmd)
				if p.ExitCode() != 0 {
					t.Fatal("Expecting exit code 0, got:", p.ExitCode())
				}
				if !p.IsSuccess() {
					t.Fatal("Process should be success")
				}
				result := new(bytes.Buffer)
				if _, err := result.ReadFrom(p.Out()); err != nil {
					t.Fatal(err)
				}

				if strings.TrimSpace(result.String()) != "HELLO WORLD!" {
					t.Fatal("Unexpected command result:", result.String())
				}
			},
		},
		{
			name:   "run proc with proc.Result()",
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
			name:   "simple run",
			cmdStr: `echo "HELLO WORLD!"`,
			exec: func(cmd string) {
				result := Run(cmd)
				if result != "HELLO WORLD!" {
					t.Fatal("Unexpected command result:", result)
				}
			},
		},
		{
			name:   "simple with expansion",
			cmdStr: "echo $MSG",
			exec: func(cmd string) {
				v := vars.New().Vars("MSG=Hello World")
				result := Run(v.Eval(cmd))
				if result != v.Val("MSG") {
					t.Fatal("Unexpected command result:", result)
				}
			},
		},
		{
			name:   "bad command",
			cmdStr: `foobar "HELLO WORLD!"`,
			exec: func(cmd string) {
				result := RunProc(cmd)
				if result.Err() == nil {
					t.Error("expecting command to fail")
				}

				t.Log(result.Err())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.exec(test.cmdStr)
		})
	}
}

func TestRun(t *testing.T) {
	tests := []struct {
		name   string
		cmdStr string
		exec   func(string)
	}{
		{
			name:   "simple run",
			cmdStr: `echo "HELLO WORLD!"`,
			exec: func(cmd string) {
				result := Run(cmd)
				if result != "HELLO WORLD!" {
					t.Fatal("Unexpected command result:", result)
				}
			},
		},
		{
			name:   "run with expansion",
			cmdStr: "echo $MSG",
			exec: func(cmd string) {
				v := vars.New().Vars("MSG=Hello World")
				result := RunWithVars(cmd, v)
				if result != v.Val("MSG") {
					t.Fatal("Unexpected command result:", result)
				}
			},
		},
		{
			name:   "bad command",
			cmdStr: `foobar "HELLO WORLD!"`,
			exec: func(cmd string) {
				result := Run(cmd)
				if !strings.Contains(result, "executable file not found") {
					t.Errorf("Expecting result 'command not found', got: %s", result)
				}

				t.Log(result)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.exec(test.cmdStr)
		})
	}
}
