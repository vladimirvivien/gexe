package vars

import (
	"os"
	"testing"
)

func TestVariables_Envs(t *testing.T) {
	tests := []struct{
		name string
		envStr string
		envs map[string]string

	}{
		{
			name: "single line single key",
			envStr: "foo=bar",
			envs: map[string]string{"foo":"bar"},
		},
		{
			name: "single line multiple keys no spaces",
			envStr: "foo=bar bazz=razz dazz=jazz",
			envs: map[string]string{"foo":"bar", "bazz":"razz", "dazz":"jazz"},
		},
		{
			name: "single line multiple keys with spaces",
			envStr: "foo= bar bazz =razz dazz = jazz",
			envs: map[string]string{"foo":"bar", "bazz":"razz", "dazz":"jazz"},
		},
		{
			name: "multiple lines single key",
			envStr: "foo=bar\nbazz=razz",
			envs: map[string]string{"foo":"bar", "bazz":"razz"},
		},
		{
			name: "multiple lines multiple keys no spaces",
			envStr: "foo=bar bazz=razz\ndazz=jazz\nmadd=dadd sadd=fadd",
			envs: map[string]string{"foo":"bar", "bazz":"razz", "dazz":"jazz", "madd":"dadd", "sadd":"fadd"},
		},
		{
			name: "multiple lines multiple keys with spaces",
			envStr: "foo= bar bazz =razz\n dazz = jazz \n madd  =  dadd sadd= fadd",
			envs: map[string]string{"foo":"bar", "bazz":"razz", "dazz":"jazz", "madd":"dadd", "sadd":"fadd"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
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
