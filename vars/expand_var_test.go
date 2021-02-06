package vars

import (
	"testing"
)

func TestExpandVarStack(t *testing.T) {
	tests := []struct {
		name  string
		stack func() *runeStack
		test  func(*runeStack)
	}{
		{
			name: "push/pop",
			stack: func() *runeStack {
				s := newRuneStack()
				s.push('a')
				s.push('b')
				s.pop()
				s.push('$')
				return s
			},
			test: func(s *runeStack) {
				if s.depth() != 2 {
					t.Errorf("unexpected stack depth: %d", s.depth())
				}
			},
		},
		{
			name: "push/pop/peek",
			stack: func() *runeStack {
				s := newRuneStack()
				s.push('a')
				s.push('b')
				s.push('$')
				s.push('\\')
				s.pop()
				return s
			},
			test: func(s *runeStack) {
				if s.depth() != 3 {
					t.Errorf("unexpected stack depth: %d", s.depth())
				}
				if s.peek() != '$' {
					t.Errorf("unexpected stack.peek value: %s", string(s.peek()))
				}
			},
		},
		{
			name: "push/pop/isempty",
			stack: func() *runeStack {
				s := newRuneStack()
				s.push('a')
				s.push('b')
				s.pop()
				s.pop()
				s.pop()
				return s
			},
			test: func(s *runeStack) {
				if s.depth() != 0 {
					t.Errorf("unexpected stack.depth: %d", s.depth())
				}
				if !s.isEmpty() {
					t.Errorf("unexpected stack.empty status: %t", s.isEmpty())
				}
				if s.peek() != 0 {
					t.Errorf("unexpected stack.peek value: %s", string(s.peek()))
				}
			},
		},
		{
			name: "push/pop/isempty",
			stack: func() *runeStack {
				s := newRuneStack()
				s.push('a')
				s.push('b')
				s.pop()
				s.pop()
				s.pop()
				s.push('c')
				s.push('d')
				return s
			},
			test: func(s *runeStack) {
				if s.depth() != 2 {
					t.Errorf("unexpected stack.depth: %d", s.depth())
				}
				if s.isEmpty() {
					t.Errorf("unexpected stack.empty status: %t", s.isEmpty())
				}
				if s.peek() != 'd' {
					t.Errorf("unexpected stack.peek value: %s", string(s.peek()))
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.test(test.stack())
		})
	}
}

func TestExpandVar(t *testing.T) {
	tests := []struct {
		name       string
		escapeChar rune
		str        string
		expected   string
	}{
		{
			name:     "no expansion",
			str:      " Hello, from the world!  ",
			expected: " Hello, from the world!  ",
		},
		{
			name:     `default escape chr - all`,
			str:      `\\\\\ \\\ \\\`,
			expected: `\\\\\ \\\ \\\`,
		},
		{
			name:     `default escape chr- string middle`,
			str:      `this \ that`,
			expected: `this \ that`,
		},
		{
			name:     `default escape chr- at end of word`,
			str:      `this\ that`,
			expected: `this\ that`,
		},
		{
			name:     `default escape chr- start of word`,
			str:      `this \that`,
			expected: `this \that`,
		},
		{
			name:     `default escape chr- start of string`,
			str:      `\this that`,
			expected: `\this that`,
		},
		{
			name:     `default escape chr- end of string`,
			str:      `this that\`,
			expected: `this that\`,
		},
		{
			name:     `default escape chr- inside single word`,
			str:      `this\that`,
			expected: `this\that`,
		},
		{
			name:     `default escape chr - inside multi words`,
			str:      `this w\t t\at`,
			expected: `this w\t t\at`,
		},
		{
			name:     `default escape chr - multi insde word`,
			str:      `t\\s that`,
			expected: `t\\s that`,
		},
		{
			name:     `default escape chr - inside multi words`,
			str:      `t\\s t\ha\t`,
			expected: `t\\s t\ha\t`,
		},
		{
			name:     `default escape chr - multi start of word`,
			str:      `this \\\\that`,
			expected: `this \\\\that`,
		},
		{
			name:     `default escape chr - multi string middle`,
			str:      `this \\\\ that`,
			expected: `this \\\\ that`,
		},
		{
			name:     `default escape chr  - multi end of word`,
			str:      `this\\\ that`,
			expected: `this\\\ that`,
		},
		{
			name:     `default escape chr - multi start of string`,
			str:      `\\\this that`,
			expected: `\\\this that`,
		},
		{
			name:     `default escape chr - multi end of string`,
			str:      `this that\\\`,
			expected: `this that\\\`,
		},
		{
			name:     `default escape chr - multi inside word`,
			str:      `this\\\that`,
			expected: `this\\\that`,
		},

		// Tests strings with dollar signs - different escape chars
		{
			name:       `escape with %`,
			escapeChar: '%',
			str:        `%$this that`,
			expected:   `$this that`,
		},
		{
			name:       `escape with #`,
			escapeChar: '#',
			str:        `this #$is that`,
			expected:   `this $is that`,
		},
		{
			name:       `escape with @`,
			escapeChar: '@',
			str:        `this @$that`,
			expected:   `this $that`,
		},
		{
			name:       `escape with &`,
			escapeChar: '&',
			str:        `thi\s\ &$that`,
			expected:   `thi\s\ $that`,
		},
		{
			name:       `escape with *`,
			escapeChar: '*',
			str:        `*$this th\at\`,
			expected:   `$this th\at\`,
		},
		{
			name:       `escape with ?`,
			escapeChar: '?',
			str:        `this?$isthat`,
			expected:   `this$isthat`,
		},
		{
			name:       `escape with slash`,
			escapeChar: '\\',
			str:        `this \${is} that`,
			expected:   `this ${is} that`,
		},
		{
			name:       `escape with !`,
			escapeChar: '!',
			str:        `this!${is}that or other`,
			expected:   `this${is}that or other`,
		},
		{
			name:     `dollar - all`,
			str:      `$$$$$ $$ $$$`,
			expected: `$$$$$ $$ $$$`,
		},
		{
			name:     `dollar - single middle`,
			str:      `foo $ bar`,
			expected: `foo $ bar`,
		},
		{
			name:     `dollar - single end of word`,
			str:      `foo$ bar`,
			expected: `foo$ bar`,
		},
		{
			name:     `dollar - single end of string`,
			str:      `foo$ bar$`,
			expected: `foo$ bar$`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			vars := New().WithEscapeChar(test.escapeChar)
			result := vars.ExpandVar(test.str, vars.Val)
			if result != test.expected {
				t.Errorf("expecting [%s] got [%s]", test.expected, result)
			}
		})
	}
}

func TestExpandVar_Eval(t *testing.T) {
	tests := []struct {
		name       string
		escapeChar rune
		setup      func(*Variables) *Variables
		str        string
		expected   string
	}{
		{
			name: `ExpandVar single var from start`,
			setup: func(v *Variables) *Variables {
				return v.SetVar("foo", "boo")
			},
			str:      `$foo bar`,
			expected: `boo bar`,
		},
		{
			name: `ExpandVar single var at end of str`,
			setup: func(v *Variables) *Variables {
				return v.SetVar("bar", "zaar")
			},
			str:      `foo $bar`,
			expected: `foo zaar`,
		},
		{
			name: `ExpandVar single var embedded`,
			setup: func(v *Variables) *Variables {
				return v.SetVar("bar", "zaar")
			},
			str:      `foo:$bar:cat`,
			expected: `foo:zaar:cat`,
		},
		{
			name: `ExpandVar with missing var`,
			setup: func(v *Variables) *Variables {
				return v.SetVar("bar", "zaar")
			},
			str:      `foo:$bar:cat:$tar`,
			expected: `foo:zaar:cat:`,
		},
		{
			name: `ExpandVar multiple envs`,
			setup: func(v *Variables) *Variables {
				return v.Envs("bar=zaar bazz=raaz")
			},
			str:      `foo $bar with $bazz`,
			expected: `foo zaar with raaz`,
		},
		{
			name: `ExpandVar multiple envs with curlies`,
			setup: func(v *Variables) *Variables {
				return v.Envs("bar=zaar bazz=raaz")
			},
			str:      `foo ${bar} with ${bazz}`,
			expected: `foo zaar with raaz`,
		},
		{
			name: `ExpandVar multiple envs with missing var`,
			setup: func(v *Variables) *Variables {
				return v.Envs("bar=zaar bazz=raaz")
			},
			str:      `foo ${bar} with ${bazz} at $jazz`,
			expected: `foo zaar with raaz at `,
		},
		{
			name: `ExpandVar multiple envs embedded`,
			setup: func(v *Variables) *Variables {
				return v.Envs("bar=zaar bazz=raaz")
			},
			str:      `foo${bar}with that${bazz}`,
			expected: `foozaarwith thatraaz`,
		},
		{
			name:     `ExpandVar with dollar amount`,
			setup:    func(v *Variables) *Variables { return v },
			str:      `foo $120.00`,
			expected: `foo 20.00`,
		},
		{
			name:     `ExpandVar with dollar amount escaped`,
			setup:    func(v *Variables) *Variables { return v },
			str:      `foo \$120.00`,
			expected: `foo $120.00`,
		},
		{
			name: `ExpandVar shell with default \ escape`,
			setup: func(v *Variables) *Variables {
				return v.SetEnv("DIR", "/var/logs")
			},
			str:      `/bin/bash -c 'files=\$(sudo find $DIR); for f in \$files; do cat \$f; done'`,
			expected: `/bin/bash -c 'files=$(sudo find /var/logs); for f in $files; do cat $f; done'`,
		},
		{
			name: `ExpandVar shell with default % escape`,
			escapeChar: '%',
			setup: func(v *Variables) *Variables {
				return v.SetEnv("DIR", "/var/logs")
			},
			str:      `/bin/bash -c 'files=%$(sudo find $DIR); for f in %$files; do cat %$f; done'`,
			expected: `/bin/bash -c 'files=$(sudo find /var/logs); for f in $files; do cat $f; done'`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			vars := test.setup(New().WithEscapeChar(test.escapeChar))
			result := vars.Eval(test.str)
			if result != test.expected {
				t.Errorf("expecting [%s] got [%s]", test.expected, result)
			}
		})
	}
}
