// File provided by the K Framework Go backend. Timestamp: 2019-07-15 11:26:24.140

package koreparser

import (
	"bytes"
	"log"
)

func isWhitespace(c rune) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}

// UnescapeKString ... removes starting and ending double-quotes, unescapes double-quotes inside
func UnescapeKString(s []byte) []byte {
	if len(s) < 2 {
		log.Fatalf("K string should begin and end with '\"'. Its length cannot therefore be less than 2. Actual string: %s", s)
	}
	if s[0] != '"' {
		log.Fatalf("K string should begin with '\"'. Actual string: %s", s)
	}
	if s[len(s)-1] != '"' {
		log.Fatalf("K string should end with '\"'. Actual string: %s", s)
	}
	s = s[1 : len(s)-1]
	s = bytes.ReplaceAll(s, []byte("\\\""), []byte("\"")) // unescape double-quotes
	// TODO: unescape other chars, like \n \t etc. ??
	return s
}

// UnescapeKLabel ... removes starting and ending back-quotes, unescapes back-quotes inside
func UnescapeKLabel(s []byte) []byte {
	if len(s) < 2 {
		log.Fatalf("K label should begin and end with '`'. Its length cannot therefore be less than 2. Actual string: %s", s)
	}
	if s[0] != '`' {
		log.Fatalf("K label should begin with '`\"`'. Actual string: %s", s)
	}
	if s[len(s)-1] != '`' {
		log.Fatalf("K label should end with '`'. Actual string: %s", s)
	}
	s = s[1 : len(s)-1]
	s = bytes.ReplaceAll(s, []byte("\\`"), []byte("`")) // unescape back ticks
	// TODO: unescape other chars, like \n \t etc. ??
	return s
}
