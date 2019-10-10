package echo

import (
	"os"
	"testing"
)

func TestEchoVar(t *testing.T) {
	tests := []struct {
		name string
		echo func() *Echo
		test func(*Echo)
	}{
		{
			name: "SetVar",
			echo: func() *Echo {
				e := New()
				e.SetVar("foo", "bar")
				e.SetVar("fuzz", "buzz")
				return e
			},
			test: func(e *Echo) {
				if e.Val("foo") != "bar" {
					t.Fatal("unexpected value:", e.Val("foo"))
				}
				if e.Val("fuzz") != "buzz" {
					t.Fatal("unexpected value:", e.Val("fuzz"))
				}
			},
		},
		{
			name: "Var with single line",
			echo: func() *Echo {
				e := New()
				e.Var("foo=bar fuzz=buzz")
				return e
			},
			test: func(e *Echo) {
				if e.Val("foo") != "bar" {
					t.Fatal("unexpected value:", e.Val("foo"))
				}
				if e.Val("fuzz") != "buzz" {
					t.Fatal("unexpected value:", e.Val("fuzz"))
				}
			},
		},
		{
			name: "Var with multilines line",
			echo: func() *Echo {
				e := New()
				e.Var("bazz=azz foo=bar\nfuzz=buzz")
				return e
			},
			test: func(e *Echo) {
				if e.Val("bazz") != "azz" {
					t.Fatal("unexpected value:", e.Val("bazz"))
				}
				if e.Val("foo") != "bar" {
					t.Fatal("unexpected value:", e.Val("foo"))
				}
				if e.Val("fuzz") != "buzz" {
					t.Fatal("unexpected value:", e.Val("fuzz"))
				}
			},
		},
		{
			name: "Var with expansion",
			echo: func() *Echo {
				e := New()
				e.SetVar("bazz", "dazz")
				e.Var("foo=${bazz} fuzz=buzz")
				return e
			},
			test: func(e *Echo) {
				if e.Val("foo") != "dazz" {
					t.Fatal("unexpected value:", e.Val("foo"))
				}
				if e.Val("fuzz") != "buzz" {
					t.Fatal("unexpected value:", e.Val("fuzz"))
				}
			},
		},
		{
			name: "Var overwrite Env",
			echo: func() *Echo {
				os.Setenv("foo", "fuzz")
				e := New()
				e.SetVar("foo", "bar")
				return e
			},
			test: func(e *Echo) {
				if e.Val("foo") != "bar" {
					t.Fatal("unexpected value:", e.Val("foo"))
				}
				if os.Getenv("foo") != "" {
					t.Fatal("Var overwrite failed:", os.Getenv("foo"))
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.test(test.echo())
		})
	}
}

func TestEchoEnv(t *testing.T) {
	tests := []struct {
		name string
		echo func() *Echo
		test func(*Echo)
	}{
		{
			name: "SetEnv",
			echo: func() *Echo {
				e := New()
				e.SetEnv("foo", "bar")
				e.SetEnv("fuzz", "buzz")
				return e
			},
			test: func(e *Echo) {
				if e.Val("foo") != "bar" {
					t.Fatal("unexpected value:", e.Val("foo"))
				}
				if e.Val("fuzz") != "buzz" {
					t.Fatal("unexpected value:", e.Val("fuzz"))
				}
			},
		},
		{
			name: "Env with single line",
			echo: func() *Echo {
				e := New()
				e.Env("foo=bar fuzz=buzz")
				return e
			},
			test: func(e *Echo) {
				if e.Val("foo") != "bar" {
					t.Fatal("unexpected value:", e.Val("foo"))
				}
				if e.Val("fuzz") != "buzz" {
					t.Fatal("unexpected value:", e.Val("fuzz"))
				}
			},
		},
		{
			name: "Env with multilines line",
			echo: func() *Echo {
				e := New()
				e.Env("bazz=azz foo=bar\nfuzz=buzz")
				return e
			},
			test: func(e *Echo) {
				if e.Val("bazz") != "azz" {
					t.Fatal("unexpected value:", e.Val("bazz"))
				}
				if e.Val("foo") != "bar" {
					t.Fatal("unexpected value:", e.Val("foo"))
				}
				if e.Val("fuzz") != "buzz" {
					t.Fatal("unexpected value:", e.Val("fuzz"))
				}
			},
		},
		{
			name: "Env with expansion",
			echo: func() *Echo {
				e := New()
				e.SetVar("bazz", "dazz")
				e.Env("foo=${bazz} fuzz=buzz")
				return e
			},
			test: func(e *Echo) {
				if e.Val("foo") != "dazz" && e.Val("fuzz") != "buzz" {
					t.Fatal("unexpected value:", e.Val("foo"))
				}
				if e.Val("fuzz") != "buzz" {
					t.Fatal("unexpected value:", e.Val("fuzz"))
				}
			},
		},
		{
			name: "Env overwrite Var",
			echo: func() *Echo {
				e := New()
				e.vars["foo"] = "fuzz"
				e.SetEnv("foo", "bar")
				return e
			},
			test: func(e *Echo) {
				if e.Val("foo") != "bar" {
					t.Fatal("unexpected value:", e.Val("foo"))
				}
				if _, ok := e.vars["foo"]; ok {
					t.Fatal("Env overwrite failed:", e.vars["foo"])
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.test(test.echo())
		})
	}
}
