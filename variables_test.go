package gexe

import (
	"testing"
)

func TestEchoVar(t *testing.T) {
	tests := []struct {
		name string
		echo func() *Session
		test func(*testing.T, *Session)
	}{
		{
			name: "SetVar",
			echo: func() *Session {
				e := New()
				e.SetVar("foo", "bar")
				e.SetVar("fuzz", "buzz")
				return e
			},
			test: func(t *testing.T, e *Session) {
				if e.Val("foo") != "bar" {
					t.Fatal("unexpected value:", e.Val("foo"))
				}
				if e.Val("fuzz") != "buzz" {
					t.Fatal("unexpected value:", e.Val("fuzz"))
				}
			},
		},
		{
			name: "Vars multiple",
			echo: func() *Session {
				e := New()
				e.Vars("foo=bar", "fuzz=buzz")
				return e
			},
			test: func(t *testing.T, e *Session) {
				if e.Val("foo") != "bar" {
					t.Fatal("unexpected value:", e.Val("foo"))
				}
				if e.Val("fuzz") != "buzz" {
					t.Fatal("unexpected value:", e.Val("fuzz"))
				}
			},
		},
		{
			name: "Vars with quoted",
			echo: func() *Session {
				e := New()
				e.Vars("bazz=azz", `foo="bar"`, `fuzz='buzz'`)
				return e
			},
			test: func(t *testing.T, e *Session) {
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
			name: "Vars with expansion",
			echo: func() *Session {
				e := New()
				e.SetVar("bazz", "dazz")
				e.Vars("foo=${bazz}", "fuzz=buzz")
				return e
			},
			test: func(t *testing.T, e *Session) {
				if e.Val("foo") != "dazz" {
					t.Fatal("unexpected value:", e.Val("foo"))
				}
				if e.Val("fuzz") != "buzz" {
					t.Fatal("unexpected value:", e.Val("fuzz"))
				}
			},
		},
		{
			name: "Vars overwrite Env",
			echo: func() *Session {
				e := New()
				e.SetVar("foo", "bar")
				e.SetEnv("foo", "fuzz")
				return e
			},
			test: func(t *testing.T, e *Session) {
				if e.Val("foo") != "bar" {
					t.Fatal("unexpected value:", e.Val("foo"))
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.test(t, test.echo())
		})
	}
}

func TestEchoEnv(t *testing.T) {
	tests := []struct {
		name string
		echo func() *Session
		test func(*Session)
	}{
		{
			name: "SetEnv",
			echo: func() *Session {
				e := New()
				e.SetEnv("foo", "bar")
				e.SetEnv("fuzz", "buzz")
				return e
			},
			test: func(e *Session) {
				if e.Val("foo") != "bar" {
					t.Fatal("unexpected value:", e.Val("foo"))
				}
				if e.Val("fuzz") != "buzz" {
					t.Fatal("unexpected value:", e.Val("fuzz"))
				}
			},
		},
		{
			name: "Envs",
			echo: func() *Session {
				e := New()
				e.Envs("foo=bar", "fuzz=buzz")
				return e
			},
			test: func(e *Session) {
				if e.Val("foo") != "bar" {
					t.Fatal("unexpected value:", e.Val("foo"))
				}
				if e.Val("fuzz") != "buzz" {
					t.Fatal("unexpected value:", e.Val("fuzz"))
				}
			},
		},
		{
			name: "Envs quoated",
			echo: func() *Session {
				e := New()
				e.Envs("bazz=azz", `foo="bar"`, `fuzz='buzz'`)
				return e
			},
			test: func(e *Session) {
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
			echo: func() *Session {
				e := New()
				e.SetVar("bazz", "dazz")
				e.Envs(`foo=${bazz}`, `fuzz=buzz`)
				return e
			},
			test: func(e *Session) {
				if e.Val("foo") != "dazz" && e.Val("fuzz") != "buzz" {
					t.Fatal("unexpected value:", e.Val("foo"))
				}
				if e.Val("fuzz") != "buzz" {
					t.Fatal("unexpected value:", e.Val("fuzz"))
				}
			},
		},
		{
			name: "Env overwrite Vars",
			echo: func() *Session {
				e := New()
				e.SetVar("foo", "fuzz")
				e.SetEnv("foo", "bar")
				return e
			},
			test: func(e *Session) {
				if e.Val("foo") != "fuzz" {
					t.Fatal("unexpected value:", e.Val("foo"))
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
