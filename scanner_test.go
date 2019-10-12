package echo

import "testing"

func TestEchoSplitWords(t *testing.T) {
	tests := []struct {
		name string
		str  func() string
		test func([]string)
	}{
		{
			name: "no quotes",
			str: func() string {
				return `a ab aba bbb cc`
			},
			test: func(parts []string) {
				if len(parts) != 5 {
					t.Error("unexpected number of split items:", len(parts))
				}
			},
		},
		{
			name: "all quoted items",
			str: func() string {
				return `"a" "ab" "aba" "bbb" "cc"`
			},
			test: func(parts []string) {
				if len(parts) != 5 {
					t.Error("unexpected number of split items:", len(parts))
				}
				if parts[0] != "a" &&
					parts[1] != "ab" &&
					parts[2] != "aba" &&
					parts[3] != "bbb" &&
					parts[4] != "cc" {
					t.Error("unexpected words returned by splitWords:", parts)
				}
			},
		},
		{
			name: "mix quoted items",
			str: func() string {
				return `"a ab" aba "bbb" "cc ee"`
			},
			test: func(parts []string) {
				if len(parts) != 4 {
					t.Error("unexpected number of split items:", len(parts))
				}
				if parts[0] != "a ab" &&
					parts[1] != "aba" &&
					parts[2] != "bbb" &&
					parts[3] != "cc ee" {
					t.Error("unexpected words returned by splitWords:", parts)
				}
			},
		},
		{
			name: "nested quotes",
			str: func() string {
				return `"a ab" aba 'bbb:"ccc"' "dd='ee'"`
			},
			test: func(parts []string) {
				if len(parts) != 4 {
					t.Error("unexpected number of split items:", len(parts), parts)
				}
				if parts[0] != "a ab" &&
					parts[1] != "aba" &&
					parts[2] != `bbb:"ccc"` &&
					parts[3] != `dd='ee'` {
					t.Error("unexpected words returned by splitWords:", parts)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := New()
			test.test(e.splitWords(test.str()))
		})
	}
}
