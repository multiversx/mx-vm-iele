// File provided by the K Framework Go backend. Timestamp: 2019-06-12 17:39:27.917

// This file holds the go generate command to run yacc on the grammar in koreparser.y.
// To build koreparser:
//	% go generate
//	% go build

//go:generate goyacc -o koreparser.go -p "kore" koreparser.y

package koreparser
