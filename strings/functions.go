package strings

import (
	"bytes"
	"io"
	"regexp"
	"strconv"
	"strings"
)

var (
	notSpaceRegex = regexp.MustCompile(`\S`)
)

type String struct {
	val string
	err error
}

func New(str string) *String {
	return &String{val:str}
}

// String returns the string value
func (s *String) String() string {
	return s.val
}

// Err returns any captured error
func (s *String) Err() error {
	return s.err
}


// IsEmpty returns true if len(s) == 0
func (s *String) IsEmpty() bool {
	return s.val == ""
}

// Streq returns true if both strings are equal
func (s *String) Eq(val1 string) bool {
	return strings.EqualFold(s.val,val1)
}


// Split s.val using the sep as delimiter
func (s *String) Split(sep string) []string {
	return strings.Split(s.val, sep)
}

// SplitSpaces properly splits s.val into []elements
// separated by one or more Unicode.IsSpace characters
// i.e. SplitSpaces("ab   cd e\tf\ng") returns 5 elements
func (s *String) SplitSpaces() []string{
	return notSpaceRegex.Split(s.val, -1)
}

func (s *String) SplitRegex(exp string) []string {
	return regexp.MustCompile(exp).Split(s.val, -1)
}


// Bytes returns []byte(s.val)
func (s *String) Bytes() []byte {
	return []byte(s.val)
}

// ToBool converts s.val from string to a bool representation
// Check s.Error() for parsing errors
func (s *String) Bool() bool {
	val, err := strconv.ParseBool(s.val)
	if err != nil {
		s.err = err
	}
	return val
}

// ToInt converts s.val from string to a int representation
// Check s.Error() for parsing errors
func (s *String) Int() int {
	val, err := strconv.Atoi(s.val)
	if err != nil {
		s.err = err
	}
	return val
}

// ToFloat converts s.val from string to a float64 representation
// Check s.Error() for parsing errors
func (s *String) Float64() float64 {
	val, err := strconv.ParseFloat(s.val, 64)
	if err != nil {
		s.err = err
	}
	return val
}

func (s *String) Reader() io.Reader {
	return bytes.NewReader([]byte(s.val))
}


// Lower returns val as lower case
func (s *String) ToLower() *String {
	s.val = strings.ToLower(s.val)
	return s
}

// Upper returns val as upper case
func (s *String) ToUpper() *String {
	s.val = strings.ToUpper(s.val)
	return s
}

func(s *String) ToTitle() *String {
	s.val = strings.ToTitle(s.val)
	return s
}

// Trim removes spaces around a val
func (s *String) TrimSpaces() *String {
	s.val = strings.TrimSpace(s.val)
	return s
}

// TrimLeft removes each character in cutset at the
// start of s.val
func (s *String) TrimLeft(cutset string) *String {
	s.val = strings.TrimLeft(s.val, cutset)
	return s
}

// TrimRight removes each character in cutset removed at the
// start of s.val
func (s *String) TrimRight(cutset string) *String {
	s.val = strings.TrimRight(s.val, cutset)
	return s
}

// Trim removes each character in cutset from around s.val
func (s *String) Trim(cutset string) *String {
	s.val = strings.Trim(s.val, cutset)
	return s
}

// ReplaceAll replaces all occurrences of old with new in s.val
func (s *String) ReplaceAll(old, new string) *String {
	s.val = strings.ReplaceAll(s.val, old, new)
	return s
}

// Concat concatenates val1 to s.val
func (s *String) Concat(vals...string) *String {
	s.val = strings.Join(append([]string{s.val}, vals...), "")
	return s
}

// CopyTo copies s.val unto dest
// Check s.Error() for copy error.
func (s *String) CopyTo(dest io.Writer) * String {
	if _, err := io.Copy(dest, bytes.NewBufferString(s.val)); err != nil {
		s.err = err
	}
	return s
}
