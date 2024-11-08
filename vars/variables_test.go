package vars

import (
	"os"
	"strings"
	"testing"
)

func TestParseVars(t *testing.T) {
	tests := []struct {
		name         string
		vars         []string
		expectedVars map[string]string
	}{
		{
			name:         "single line single key",
			vars:         []string{"foo=bar"},
			expectedVars: map[string]string{"foo": "bar"},
		},
		{
			name:         "multiple values",
			vars:         []string{"foo=bar", "bazz=razz", "dazz=jazz"},
			expectedVars: map[string]string{"foo": "bar", "bazz": "razz", "dazz": "jazz"},
		},
		{
			name:         "with var references",
			vars:         []string{"foo=bar", `bazz=${foo}`, `dazz=jazz ${foo}`},
			expectedVars: map[string]string{"foo": "bar", "bazz": "${foo}", "dazz": "jazz ${foo}"},
		},
		{
			name:         "no value error",
			vars:         []string{"foo=bar", "bazz", "dazz=jazz"},
			expectedVars: map[string]string{"foo": "bar", "dazz": "jazz"},
		},
		{
			name:         "unclosed quote error",
			vars:         []string{"foo=bar", "bazz='booz", "dazz=jazz"},
			expectedVars: map[string]string{"foo": "bar", "bazz": "booz", "dazz": "jazz"},
		},
		{
			name:         "multiple items",
			vars:         []string{"foo=bar bazz='booz'", "dazz=jazz"},
			expectedVars: map[string]string{"foo": "bar bazz=", "dazz": "jazz"},
		},
		{
			name:         "empty items",
			vars:         nil,
			expectedVars: map[string]string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			vars := New()
			varmaps := vars.parseVars(test.vars...)
			for _, parsedVar := range varmaps {
				if strings.TrimSpace(test.expectedVars[parsedVar.key]) != strings.TrimSpace(parsedVar.value) {
					t.Errorf("unexpected var: %s=%s (want %s=%s)", parsedVar.key, parsedVar.value, parsedVar.key, test.expectedVars[parsedVar.key])
				}
			}
		})
	}
}

func TestVariables_Vars(t *testing.T) {
	tests := []struct {
		name         string
		vars         []string
		expectedVars map[string]string
		shouldFail   bool
	}{
		{
			name:         "single line single key",
			vars:         []string{"foo=bar"},
			expectedVars: map[string]string{"foo": "bar"},
		},
		{
			name:         "multiple values",
			vars:         []string{"foo=bar", "bazz=razz", "dazz=jazz"},
			expectedVars: map[string]string{"foo": "bar", "bazz": "razz", "dazz": "jazz"},
		},
		{
			name:         "with var references",
			vars:         []string{"foo=bar", `bazz=${foo}`, `dazz="jazz ${foo}"`},
			expectedVars: map[string]string{"foo": "bar", "bazz": "bar", "dazz": "jazz bar"},
		},
		{
			name:         "no value error",
			vars:         []string{"foo=bar", "bazz", "dazz=jazz"},
			expectedVars: map[string]string{"foo": "bar", "dazz": "jazz"},
			shouldFail:   true,
		},
		{
			name:         "unclosed quote error",
			vars:         []string{"foo=bar", "bazz='booz", "dazz=jazz"},
			expectedVars: map[string]string{"foo": "bar", "bazz": "booz", "dazz": "jazz"},
			shouldFail:   true,
		},
		{
			name:         "multiple items",
			vars:         []string{"foo=bar bazz='booz'", "dazz=jazz"},
			expectedVars: map[string]string{"foo": "bar bazz=", "dazz": "jazz"},
			shouldFail:   true,
		},
		{
			name:         "empty items",
			vars:         nil,
			expectedVars: map[string]string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			vars := New()
			vars.Vars(test.vars...)

			for key, val := range test.expectedVars {
				if v, ok := vars.vars[key]; !ok || v != val {
					t.Errorf("unexpected var: %s=%s (want %s=%s)", key, v, key, val)
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

func TestVariables_Envs(t *testing.T) {
	tests := []struct {
		name         string
		envs         []string
		expectedEnvs map[string]string
	}{
		{
			name:         "single line single key",
			envs:         []string{"foo=bar"},
			expectedEnvs: map[string]string{"foo": "bar"},
		},
		{
			name:         "multiple values",
			envs:         []string{"foo=bar", "bazz=razz", "dazz=jazz"},
			expectedEnvs: map[string]string{"foo": "bar", "bazz": "razz", "dazz": "jazz"},
		},
		{
			name:         "no value error",
			envs:         []string{"foo=bar", "bazz", "dazz=jazz"},
			expectedEnvs: map[string]string{"foo": "bar", "dazz": "jazz"},
		},
		{
			name:         "unclosed quote error",
			envs:         []string{"foo=bar", "bazz='booz", "dazz=jazz"},
			expectedEnvs: map[string]string{"foo": "bar", "dazz": "jazz"},
		},
		{
			name:         "multiple items",
			envs:         []string{"foo=bar bazz='booz'", "dazz=jazz"},
			expectedEnvs: map[string]string{"foo": "bar bazz=", "dazz": "jazz"},
		},
		{
			name:         "empty items",
			envs:         nil,
			expectedEnvs: map[string]string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			vars := New()
			vars.Envs(test.envs...)
			for key, val := range test.expectedEnvs {
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
				vars.Envs("foo=bar", "bar=bazz")
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
				vars.Vars("foo=bar", "bar=bazz")
				return values
			},
		},
		{
			name: "mixed Envs Vars",
			setup: func(vars *Variables) map[string]string {
				values := map[string]string{"foo": "bar", "batt": "bazz", "kazz": "jazz", "razz": "dazz"}
				vars.Vars("foo=bar", "batt=bazz", "razz=dazz")
				vars.Envs("kazz=jazz", "razz=wazz")
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
				return vars.Vars("foo=bar", "dood=daad")
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
				return vars.Envs("foo=bar", "dood=daad")
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
