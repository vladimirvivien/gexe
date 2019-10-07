package echo

import "testing"

func TestEchoVariables(t *testing.T) {
	tests := []struct {
		name string
		echo func() *echo
		test func(*echo)
	}{
		{
			name: "set and get var",
			echo: func() *echo {
				e := New()
				e.Var("foo", "bar")
				return e
			},
			test: func(e *echo) {
				if e.Val("foo") != "bar" {
					t.Fatal("unexpected value:", e.Val("foo"))
				}
			},
		},
		{
			name: "set and get env and var",
			echo: func() *echo {
				e := New()
				e.Export("fuzz", "buzz")
				e.Var("foo", "bar")
				return e
			},
			test: func(e *echo) {
				if e.Val("foo") != "bar" {
					t.Fatal("unexpected value:", e.Val("foo"))
				}
				if e.Val("fuzz") != "buzz" {
					t.Fatal("unexpected value:", e.Val("fuzz"))
				}
			},
		},
		{
			name: "var overwrites env",
			echo: func() *echo {
				e := New()
				e.Export("fuzz", "buzz")
				e.Var("fuzz", "bar")
				return e
			},
			test: func(e *echo) {
				if e.Val("fuzz") != "bar" {
					t.Fatal("unexpected value:", e.Val("fuzz"))
				}
			},
		},
		{
			name: "export overwrites var",
			echo: func() *echo {
				e := New()
				e.Var("fuzz", "buzz")
				e.Export("fuzz", "bar")
				return e
			},
			test: func(e *echo) {
				if e.Val("fuzz") != "bar" {
					t.Fatal("unexpected value:", e.Val("fuzz"))
				}
			},
		},
		{
			name: "var exapansion",
			echo: func() *echo {
				e := New()
				e.Var("Foo", "bar!")
				e.Var("fuzz", "${Foo}")
				return e
			},
			test: func(e *echo) {
				if e.Val("fuzz") != "bar!" {
					t.Fatal("unexpected value:", e.Val("fuzz"))
				}
			},
		},
		{
			name: "export exapansion",
			echo: func() *echo {
				e := New()
				e.Var("Foo", "bar!")
				e.Export("fuzz", "${Foo}")
				return e
			},
			test: func(e *echo) {
				if e.Val("fuzz") != "bar!" {
					t.Fatal("unexpected value:", e.Val("fuzz"))
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
