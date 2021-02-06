package strings

import (
	"bytes"
	"testing"
)

func TestString_Concat(t *testing.T) {
	tests := []struct{
		name string
		items []string
		expected string
	}{
		{
			name : "single str",
			items:[]string{"Bar"},
			expected:"FooBar",
		},
		{
			name : "multiple items",
			items:[]string{"Bar","Bazz"},
			expected:"FooBarBazz",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			s := New("Foo").Concat(test.items...)
			if s.String() != test.expected {
				t.Errorf("expecting %s, got %s", test.expected, s)
			}
		})
	}
}

func TestString_CopyTo(t *testing.T) {
	var b bytes.Buffer
	New("HelloWorld").CopyTo(&b)
	if b.String() != "HelloWorld" {
		t.Fatal("CopyTo failed")
	}

}

func TestString_Eq(t *testing.T) {
	if !New("Hello, 世界").Eq("Hello, 世界") {
		t.Fatal("Eq failed")
	}
}

func TestString_IsEmpty(t *testing.T) {
	if New("Hello").IsEmpty() {
		t.Fatal("IsEmpty failed")
	}
	if !New("").IsEmpty() {
		t.Fatal("IsEmpty failed")
	}
}

func TestString_Reader(t *testing.T) {
	b := new(bytes.Buffer)
	b.ReadFrom(New("Hello").Reader())
	if b.String()  != "Hello" {
		t.Fatal("Reader failed")
	}
}

func TestString_ReplaceAll(t *testing.T) {
	s := New("lieave in oeane hourea").ReplaceAll("ea","")
	if s.String() != "live in one hour" {
		t.Fatal("ReplaceAll failed")
	}
}

func TestString_ToBool(t *testing.T) {
	if New("true").Bool() != true {
		t.Error("ToBool true unexpected failure")
	}
	if New("false").Bool() != false {
		t.Error("ToBool false unexpected failure")
	}
	s := New("foo")
	_ = s.Bool()
	if s.Err() == nil {
		t.Error("ToBool expecting error, but got none")
	}
}

func TestString_ToFloat64(t *testing.T) {

}

func TestString_ToInt(t *testing.T) {

}

func TestString_ToLower(t *testing.T) {

}

func TestString_ToTitle(t *testing.T) {

}

func TestString_ToUpper(t *testing.T) {

}

func TestString_Trim(t *testing.T) {

}

func TestString_TrimLeft(t *testing.T) {

}

func TestString_TrimRight(t *testing.T) {

}

func TestString_TrimSpaces(t *testing.T) {

}