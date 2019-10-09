package echo

import (
	"bufio"
	"strings"
	"unicode"
	"unicode/utf8"
)

// splitWords splits a string into space-seprated items
func (e *Echo) splitWords(val string) []string {
	result := []string{}
	scnr := bufio.NewScanner(strings.NewReader(val))
	scnr.Split(scanQuotedWords)
	for scnr.Scan() {
		result = append(result, scnr.Text())
	}
	if err := scnr.Err(); err != nil {
		e.shouldPanic(err.Error())
		e.shouldLog(err.Error())
	}
	return result
}

// scanQuotedWords implements bufio.SplitFunc to support quoted text when splitting words.
// Code based on https://golang.org/src/bufio/scan.go?s=13096:13174#L380
// TODO return error if unbalanced.
func scanQuotedWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Munch on empty spaces until the first non-space char or eof
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:]) // decode the first utf8 encoding in data[]
		if !unicode.IsSpace(r) {
			break
		}
	}

	// scan words
	inQuote := false
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		switch {
		// scan quoted string
		case isQuote(r):
			if !inQuote {
				inQuote = true
				start = start + width // dont include quote
				continue
			}
			return i + width, data[start:i], nil

		// scan outside of quotes
		case !isQuote(r):
			if !inQuote {
				if unicode.IsSpace(r) {
					return i + width, data[start:i], nil
				}
			}
		}
	}

	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}

	// Request more data.
	return start, nil, nil
}

func isQuote(r rune) bool {
	switch r {
	case '"', '\'':
		return true
	}
	return false
}
