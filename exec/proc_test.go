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
			name:   "access result",
			cmdStr: `echo "Hello World"`,
			exec: func(cmd string) {
				proc := NewProc(cmd)
				if err := proc.Start().Err(); err != nil {
					t.Fatal(err)
				}
				if err := proc.Wait().Err(); err != nil {
					t.Fatal(err)
				}

				if strings.TrimSpace(proc.Result()) != "Hello World" {
					t.Errorf("unexpected result: %s", proc.Result())
				}
			},
		},
		{
			name:   "access proc.Out()",
			cmdStr: `echo "Hello World"`,
			exec: func(cmd string) {
				proc := NewProc(cmd)
				if err := proc.Start().Wait().Err(); err != nil {
					t.Fatal(err)
				}

				buf := new(bytes.Buffer)
				if _, err := buf.ReadFrom(proc.Out()); err != nil {
					t.Fatal(err)
				}

				if strings.TrimSpace(buf.String()) != "Hello World" {
					t.Errorf("unexpected result: %s", buf.String())
				}
			},
		},
		{
			name:   "with expansion",
			cmdStr: "echo '$MSG'",
			exec: func(cmd string) {
				proc := NewProcWithVars(cmd, vars.New().SetVar("MSG", "Hello World"))
				if err := proc.Start().Err(); err != nil {
					t.Fatal(err)
				}
				if err := proc.Wait().Err(); err != nil {
					t.Fatal(err)
				}
				if strings.TrimSpace(proc.Result()) != "Hello World" {
					t.Errorf("unexpected result: %s", proc.Result())
				}
			},
		},
		{
			name:   "custom stdout/stderr",
			cmdStr: `echo "Hello World"`,
			exec: func(cmd string) {
				proc := NewProc(cmd)
				buf := new(bytes.Buffer)
				proc.SetStdout(buf)
				proc.SetStderr(buf)

				if err := proc.Start().Wait().Err(); err != nil {
					t.Fatalf("Failed to start/wait for proc: %s", err)
				}

				if strings.TrimSpace(buf.String()) != "Hello World" {
					t.Errorf("unexpected result: %s", buf.String())
				}
			},
		},
		{
			name:   "start with error",
			cmdStr: `foo "Hello World"`,
			exec: func(cmd string) {
				proc := NewProc(cmd)
				if err := proc.Start().Err(); err == nil {
					t.Fatalf("Expecting error, but got none")
				}

				if err := proc.Wait().Err(); err == nil {
					t.Fatal("Expecting error, but got none")
				}

				if strings.TrimSpace(proc.Result()) == "Hello World" {
					t.Errorf("Expecting error but did not get it")
				}
			},
		},
		{
			name:   "run with result",
			cmdStr: `echo "Hello World"`,
			exec: func(cmd string) {
				proc := NewProc(cmd)
				if err := proc.Run().Err(); err != nil {
					t.Fatal(err)
				}
				if strings.TrimSpace(proc.Result()) != "Hello World" {
					t.Errorf("unexpected result: %s", proc.Result())
				}
			},
		},
		{
			name:   "run with proc.Out()",
			cmdStr: `echo "Hello World"`,
			exec: func(cmd string) {
				proc := NewProc(cmd)
				if err := proc.Run().Err(); err != nil {
					t.Fatal(err)
				}
				buf := new(bytes.Buffer)
				if _, err := buf.ReadFrom(proc.Out()); err != nil {
					t.Fatal(err)
				}

				if strings.TrimSpace(buf.String()) != "Hello World" {
					t.Errorf("unexpected result: %s", buf.String())
				}
			},
		},
		{
			name:   "missing command with Result",
			cmdStr: `foobar "Hello World"`,
			exec: func(cmd string) {
				proc := NewProc(cmd)

				if err := proc.Run().Err(); err == nil {
					t.Fatalf("Expecting error, but got none")
				}
				result := proc.Result()
				if strings.TrimSpace(result) == "Hello World" {
					t.Errorf("Expecting error but did not get it")
				}
				if !strings.Contains(result, "executable file not found") {
					t.Errorf("Expecting result 'command not found', got: %s", result)
				}
			},
		},
		{
			name:   "missing command with proc.Out",
			cmdStr: `foobar "Hello World"`,
			exec: func(cmd string) {
				proc := NewProc(cmd)

				if err := proc.Run().Err(); err == nil {
					t.Fatalf("Expecting error, but got none")
				}
				result := proc.Result()
				if !strings.Contains(result, "file not found") {
					t.Errorf("Expecting result 'command not found', got: %s", result)
				}
			},
		},
		{
			name:   "bad command with result",
			cmdStr: `date -xx`,
			exec: func(cmd string) {
				proc := NewProc(cmd)

				err := proc.Run().Err()
				if err == nil {
					t.Fatalf("Expecting error, but got none")
				}

				result := strings.TrimSpace(proc.Result())
				if !strings.Contains(result, "illegal option") {
					t.Errorf("Expecting error but did not get it")
				}
			},
		},
		{
			name:   "bad command with proc.Out",
			cmdStr: `date -xx`,
			exec: func(cmd string) {
				proc := NewProc(cmd)

				err := proc.Run().Err()
				if err == nil {
					t.Fatalf("Expecting error, but got none")
				}

				buf := new(bytes.Buffer)
				if _, err := buf.ReadFrom(proc.Out()); err != nil {
					t.Fatal(err)
				}

				result := strings.TrimSpace(buf.String())
				if !strings.Contains(result, "illegal option") {
					t.Errorf("Expecting 'illegal option' error, but got: %s", result)
				}
			},
		},
		{
			name:   "proc status",
			cmdStr: `echo "HELLO WORLD!"`,
			exec: func(cmd string) {
				p := NewProc(cmd).Start()

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

				if p.Result() != "HELLO WORLD!" {
					t.Errorf("Unexpected proc.Result(): %s", p.Result())
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
			name:   "access result",
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
			name:   "access proc.Out",
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
			name:   "with expansion",
			cmdStr: "echo '$MSG'",
			exec: func(cmd string) {
				proc := StartProcWithVars(cmd, vars.New().SetVar("MSG", "Hello World")).Wait()
				if err := proc.Err(); err != nil {
					t.Fatal(err)
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

				if p.Result() != "HELLO WORLD!" {
					t.Errorf("Unexpected proc.Result(): %s", p.Result())
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
			name:   "long-running",
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
		{
			name:   "missing command with result",
			cmdStr: `foobar "HELLO WORLD!"`,
			exec: func(cmd string) {
				proc := StartProc(cmd).Wait()
				if proc.Err() == nil {
					t.Error("expecting command to fail")
				}
				result := proc.Result()
				if !strings.Contains(result, "executable file not found") {
					t.Errorf("Expecting result 'command not found', got: %s", result)
				}
				t.Log(result)
			},
		},
		{
			name:   "missing command with proc.Out",
			cmdStr: `foobar "HELLO WORLD!"`,
			exec: func(cmd string) {
				proc := StartProc(cmd).Wait()
				if proc.Err() == nil {
					t.Error("expecting command to fail")
				}
				buf := new(bytes.Buffer)
				if _, err := buf.ReadFrom(proc.Out()); err != nil {
					t.Fatal(err)
				}
				if buf.Len() != 0 {
					t.Error("Unexpected output from stdout/stderr")
				}
			},
		},
		{
			name:   "bad command",
			cmdStr: `date -xx`,
			exec: func(cmd string) {
				proc := StartProc(cmd).Wait()
				if proc.Err() == nil {
					t.Error("expecting command to fail")
				}
				result := proc.Result()
				if !strings.Contains(result, "illegal option") {
					t.Errorf("Expecting result 'command not found', got: %s", result)
				}
				t.Log(result)
			},
		},
		{
			name:   "bad command with proc.Out",
			cmdStr: `date -xx`,
			exec: func(cmd string) {
				proc := StartProc(cmd).Wait()
				if proc.Err() == nil {
					t.Error("expecting command to fail")
				}

				buf := new(bytes.Buffer)
				if _, err := buf.ReadFrom(proc.Out()); err != nil {
					t.Fatal(err)
				}

				result := strings.TrimSpace(buf.String())
				if !strings.Contains(result, "illegal option") {
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

func TestRunProc(t *testing.T) {
	tests := []struct {
		name   string
		cmdStr string
		exec   func(string)
	}{
		{
			name:   "access result",
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
			name:   "access proc.Out()",
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
			name:   "with expansion",
			cmdStr: "echo $MSG",
			exec: func(cmd string) {
				v := vars.New().Vars("MSG='Hello World'")
				proc := RunProcWithVars(cmd, v)
				if err := proc.Err(); err != nil {
					t.Fatalf("failed to run proc: %s", err)
				}
				if proc.Result() != v.Val("MSG") {
					t.Fatal("Unexpected command result:", proc.Result())
				}
			},
		},
		{
			name:   "missing command with result",
			cmdStr: `foobar "HELLO WORLD!"`,
			exec: func(cmd string) {
				proc := RunProc(cmd)
				if proc.Err() == nil {
					t.Error("expecting command to fail")
				}
				result := proc.Result()
				if !strings.Contains(result, "executable file not found") {
					t.Errorf("Expecting result 'command not found', got: %s", result)
				}
				t.Log(result)
			},
		},
		{
			name:   "missing command with proc.Out",
			cmdStr: `foobar "HELLO WORLD!"`,
			exec: func(cmd string) {
				proc := RunProc(cmd)
				if proc.Err() == nil {
					t.Error("expecting command to fail")
				}

				buf := new(bytes.Buffer)
				if _, err := buf.ReadFrom(proc.Out()); err != nil {
					t.Fatal(err)
				}

				result := buf.String()
				if result != "" {
					t.Errorf("Expecting no output, but got %s", result)
				}
				t.Log(result)
			},
		},
		{
			name:   "bad command with result",
			cmdStr: `date -xx`,
			exec: func(cmd string) {
				proc := RunProc(cmd)
				if proc.Err() == nil {
					t.Error("expecting command to fail")
				}
				result := proc.Result()
				if !strings.Contains(result, "illegal") {
					t.Errorf("Expecting result 'illegal option', got: %s", result)
				}
				t.Log(result)
			},
		},
		{
			name:   "bad command option with proc.Out",
			cmdStr: `date -xx`,
			exec: func(cmd string) {
				proc := RunProc(cmd)
				if proc.Err() == nil {
					t.Error("expecting command to fail")
				}

				buf := new(bytes.Buffer)
				if _, err := buf.ReadFrom(proc.Out()); err != nil {
					t.Fatal(err)
				}

				result := buf.String()
				if !strings.Contains(result, "illegal") {
					t.Errorf("Expecting result 'illegal option', got: %s", result)
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
				result := RunWithVars(cmd, vars.New().SetVar("MSG", "Hello World"))
				if result != "Hello World" {
					t.Fatal("Unexpected command result:", result)
				}
			},
		},
		{
			name:   "missing command",
			cmdStr: `foobar "HELLO WORLD!"`,
			exec: func(cmd string) {
				result := Run(cmd)
				if !strings.Contains(result, "not found") {
					t.Errorf("Expecting result 'command not found', got: %s", result)
				}

				t.Log(result)
			},
		},
		{
			name:   "bad command",
			cmdStr: `date -xx"`,
			exec: func(cmd string) {
				result := Run(cmd)
				if !strings.Contains(result, "illegal") {
					t.Errorf("Expecting 'illegal option', got: %s", result)
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
