package echo

import (
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
			name: "Vars with single line",
			echo: func() *Echo {
				e := New()
				e.Vars("foo=bar fuzz=buzz")
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
			name: "Vars with multilines line",
			echo: func() *Echo {
				e := New()
				e.Vars("bazz=azz foo=bar\nfuzz=buzz")
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
			name: "Vars with expansion",
			echo: func() *Echo {
				e := New()
				e.SetVar("bazz", "dazz")
				e.Vars("foo=${bazz} fuzz=buzz")
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
			name: "Vars overwrite Env",
			echo: func() *Echo {
				e := New()
				e.SetVar("foo", "bar")
				e.SetEnv("foo", "fuzz")
				return e
			},
			test: func(e *Echo) {
				if e.Val("foo") != "bar" {
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
				e.Envs("foo=bar fuzz=buzz")
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
				e.Envs("bazz=azz foo=bar\nfuzz=buzz")
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
				e.Envs("foo=${bazz} fuzz=buzz")
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
			name: "Env overwrite Vars",
			echo: func() *Echo {
				e := New()
				e.SetVar("foo", "fuzz")
				e.SetEnv("foo", "bar")
				return e
			},
			test: func(e *Echo) {
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
