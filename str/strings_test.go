package str

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
			s := String("Foo").Concat(test.items...)
			if s.String() != test.expected {
				t.Errorf("expecting %s, got %s", test.expected, s)
			}
		})
	}
}

func TestString_CopyTo(t *testing.T) {
	var b bytes.Buffer
	String("HelloWorld").CopyTo(&b)
	if b.String() != "HelloWorld" {
		t.Fatal("CopyTo failed")
	}

}

func TestString_Eq(t *testing.T) {
	if !String("Hello, 世界").Eq("Hello, 世界") {
		t.Fatal("Eq failed")
	}
}

func TestString_IsEmpty(t *testing.T) {
	if String("Hello").IsEmpty() {
		t.Fatal("IsEmpty failed")
	}
	if !String("").IsEmpty() {
		t.Fatal("IsEmpty failed")
	}
}

func TestString_Reader(t *testing.T) {
	b := new(bytes.Buffer)
	b.ReadFrom(String("Hello").Reader())
	if b.String()  != "Hello" {
		t.Fatal("Reader failed")
	}
}

func TestString_ReplaceAll(t *testing.T) {
	s := String("lieave in oeane hourea").ReplaceAll("ea","")
	if s.String() != "live in one hour" {
		t.Fatal("ReplaceAll failed")
	}
}

func TestString_ToBool(t *testing.T) {
	if String("true").Bool() != true {
		t.Error("ToBool true unexpected failure")
	}
	if String("false").Bool() != false {
		t.Error("ToBool false unexpected failure")
	}
	s := String("foo")
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