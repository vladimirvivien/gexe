package vars

import (
	"os"
	"testing"
)

func TestVariables_Envs(t *testing.T) {
	tests := []struct {
		name   string
		envStr string
		envs   map[string]string
	}{
		{
			name:   "single line single key",
			envStr: "foo=bar",
			envs:   map[string]string{"foo": "bar"},
		},
		{
			name:   "single line multiple keys no spaces",
			envStr: "foo=bar bazz=razz dazz=jazz",
			envs:   map[string]string{"foo": "bar", "bazz": "razz", "dazz": "jazz"},
		},
		{
			name:   "single line multiple keys with spaces",
			envStr: "foo= bar bazz =razz dazz = jazz",
			envs:   map[string]string{"foo": "bar", "bazz": "razz", "dazz": "jazz"},
		},
		{
			name:   "multiple lines single key",
			envStr: "foo=bar\nbazz=razz",
			envs:   map[string]string{"foo": "bar", "bazz": "razz"},
		},
		{
			name:   "multiple lines multiple keys no spaces",
			envStr: "foo=bar bazz=razz\ndazz=jazz\nmadd=dadd sadd=fadd",
			envs:   map[string]string{"foo": "bar", "bazz": "razz", "dazz": "jazz", "madd": "dadd", "sadd": "fadd"},
		},
		{
			name:   "multiple lines multiple keys with spaces",
			envStr: "foo= bar bazz =razz\n dazz = jazz \n madd  =  dadd sadd= fadd",
			envs:   map[string]string{"foo": "bar", "bazz": "razz", "dazz": "jazz", "madd": "dadd", "sadd": "fadd"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			vars := New()
			vars.Envs(test.envStr)
			for key, val := range test.envs {
				if os.Getenv(key) != val {
					t.Errorf("unexpected env: %s = %s (needs %s)", key, val, os.Getenv(key))
				}
			}
		})
	}
}

func TestVariables_SetEnv(t *testing.T) {
	tests := []struct {
		name    string
		envName string
		envVal  string
	}{
		{
			name:    "simple value",
			envName: "foo",
			envVal:  "bar",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			New().SetEnv(test.envName, test.envVal)
			if os.Getenv(test.envName) == "" {
				t.Errorf("env %s not set", test.envName)
			}
		})
	}
}

func TestVariables_Vars(t *testing.T) {
	tests := []struct {
		name   string
		varStr string
		vars   map[string]string
	}{
		{
			name:   "single line single key",
			varStr: "foo=bar",
			vars:   map[string]string{"foo": "bar"},
		},
		{
			name:   "single line multiple keys no spaces",
			varStr: "foo=bar bazz=razz dazz=jazz",
			vars:   map[string]string{"foo": "bar", "bazz": "razz", "dazz": "jazz"},
		},
		{
			name:   "single line multiple keys with spaces",
			varStr: "foo= bar bazz =razz dazz = jazz",
			vars:   map[string]string{"foo": "bar", "bazz": "razz", "dazz": "jazz"},
		},
		{
			name:   "multiple lines single key",
			varStr: "foo=bar\nbazz=razz",
			vars:   map[string]string{"foo": "bar", "bazz": "razz"},
		},
		{
			name:   "multiple lines multiple keys no spaces",
			varStr: "foo=bar bazz=razz\ndazz=jazz\nmadd=dadd sadd=fadd",
			vars:   map[string]string{"foo": "bar", "bazz": "razz", "dazz": "jazz", "madd": "dadd", "sadd": "fadd"},
		},
		{
			name:   "multiple lines multiple keys with spaces",
			varStr: "foo= bar bazz =razz\n dazz = jazz \n madd  =  dadd sadd= fadd",
			vars:   map[string]string{"foo": "bar", "bazz": "razz", "dazz": "jazz", "madd": "dadd", "sadd": "fadd"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			vars := New()
			vars.Vars(test.varStr)
			for key, val := range test.vars {
				if v, ok := vars.vars[key]; !ok || v != val {
					t.Errorf("unexpected var: %s = %s (needs %s)", key, val, v)
				}
			}
		})
	}
}

func TestVariables_SetVar(t *testing.T) {
	tests := []struct {
		name    string
		varName string
		varVal  string
	}{
		{
			name:    "simple value",
			varName: "foo",
			varVal:  "bar",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			vars := New()
			vars.SetVar(test.varName, test.varVal)
			if v, ok := vars.vars[test.varName]; !ok || v != test.varVal {
				t.Errorf("var %s not set", test.varName)
			}
		})
	}
}

func TestVariables_UnsetVar(t *testing.T) {
	tests := []struct {
		name    string
		varName string
		varVal  string
	}{
		{
			name:    "simple value",
			varName: "foo",
			varVal:  "bar",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			vars := New()
			vars.SetVar(test.varName, test.varVal)
			vars.UnsetVar(test.varName)

			if _, ok := vars.vars[test.varName]; ok {
				t.Errorf("var %s: still set after Unset", test.varName)
			}
		})
	}
}

func TestVariables_Val(t *testing.T) {
	tests := []struct {
		name  string
		setup func(*Variables) map[string]string
	}{
		{
			name: "single value with SetEnv",
			setup: func(vars *Variables) map[string]string {
				values := map[string]string{"foo": "bar"}
				for k, v := range values {
					vars.SetEnv(k, v)
				}
				return values
			},
		},
		{
			name: "multiple values with Envs",
			setup: func(vars *Variables) map[string]string {
				values := map[string]string{"foo": "bar", "bar": "bazz"}
				vars.Envs("foo=bar bar=bazz")
				return values
			},
		},
		{
			name: "single value with SetVar",
			setup: func(vars *Variables) map[string]string {
				values := map[string]string{"foo": "bar"}
				for k, v := range values {
					vars.SetVar(k, v)
				}
				return values
			},
		},
		{
			name: "multiple values with Vars",
			setup: func(vars *Variables) map[string]string {
				values := map[string]string{"foo": "bar", "bar": "bazz"}
				vars.Vars("foo=bar bar=bazz")
				return values
			},
		},
		{
			name: "mixed Envs Vars",
			setup: func(vars *Variables) map[string]string {
				values := map[string]string{"foo": "bar", "batt": "bazz", "kazz": "jazz", "razz": "dazz"}
				vars.Vars("foo=bar batt=bazz razz=dazz")
				vars.Envs("kazz=jazz razz=wazz")
				return values
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			vars := New()
			for k, v := range test.setup(vars) {
				if vars.Val(k) != v {
					t.Errorf("unexpected Val(%s) = %s", k, vars.Val(k))
				}
			}
		})
	}
}

func TestVariables_Eval(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(*Variables) *Variables
		str      string
		expected string
	}{
		{
			name: "no exapansion",
			setup: func(vars *Variables) *Variables {
				return vars.SetVar("foo", "bar")
			},
			str:      "you are bar",
			expected: "you are bar",
		},
		{
			name: "expansion single SetVar",
			setup: func(vars *Variables) *Variables {
				return vars.SetVar("foo", "bar")
			},
			str:      "you are $foo",
			expected: "you are bar",
		},
		{
			name: "expansion multiple SetVars",
			setup: func(vars *Variables) *Variables {
				return vars.SetVar("foo", "bar").SetVar("dood", "daad")
			},
			str:      "is $foo a $dood?",
			expected: "is bar a daad?",
		},
		{
			name: "expansion multiple Vars",
			setup: func(vars *Variables) *Variables {
				return vars.Vars("foo=bar dood=daad")
			},
			str:      "is $foo a $dood?",
			expected: "is bar a daad?",
		},
		{
			name: "expansion single SetEnv",
			setup: func(vars *Variables) *Variables {
				return vars.SetEnv("foo", "bar")
			},
			str:      "you are $foo",
			expected: "you are bar",
		},
		{
			name: "expansion multiple SetEnvs",
			setup: func(vars *Variables) *Variables {
				return vars.SetEnv("foo", "bar").SetEnv("dood", "daad")
			},
			str:      "is $foo a $dood?",
			expected: "is bar a daad?",
		},
		{
			name: "expansion multiple Envs",
			setup: func(vars *Variables) *Variables {
				return vars.Envs("foo=bar dood=daad")
			},
			str:      "is $foo a $dood?",
			expected: "is bar a daad?",
		},
		{
			name: "expansion mix Vars Envs",
			setup: func(vars *Variables) *Variables {
				return vars.Envs("foo=bar").Vars("dood=daad")
			},
			str:      "is $foo a $dood?",
			expected: "is bar a daad?",
		},
		{
			name: "expansion mix overwrite",
			setup: func(vars *Variables) *Variables {
				return vars.Vars("foo=bar").Envs("foo=daad")
			},
			str:      "you are a $foo",
			expected: "you are a bar",
		},
		{
			name: "expansion with escape",
			setup: func(vars *Variables) *Variables {
				return vars.Envs("foo=bar")
			},
			str:      "$foo a \\$dood?",
			expected: "bar a $dood?",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			evalStr := test.setup(New().WithEscapeChar('\\')).Eval(test.str)
			if evalStr != test.expected {
				t.Errorf("expecting Eval(%s) = %s, got %s", test.str, test.expected, evalStr)
			}
		})
	}
}
