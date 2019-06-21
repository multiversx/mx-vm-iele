// File provided by the K Framework Go backend. Timestamp: 2019-06-20 20:57:09.954

package koreparser

import (
	"log"
	"regexp"
	"unicode/utf8"
)

// The parser expects the lexer to return 0 on EOF.  Give it a name
// for clarity.
const eof = 0

type koreLexRegexPostprocessing func([]byte) []byte

type koreLexRegexRule struct {
	regexp         *regexp.Regexp             // compiled regex rule
	token          int                        // token to retrieve
	postProcessing koreLexRegexPostprocessing // cleanup to be performed after the string is obtained
}

var koreLexRegexRules []koreLexRegexRule

func init() {
	createRegexRule := func(rule string, token int, postProcessing koreLexRegexPostprocessing) koreLexRegexRule {
		// the '\A' means only match from the beginning of the string
		// this is essential to the way we process the input
		// we match the longest rule at the beginning of the slice, then we move the head of the slice ahead
		return koreLexRegexRule{
			regexp:         regexp.MustCompile("\\A" + rule),
			token:          token,
			postProcessing: postProcessing}
	}
	koreLexRegexRules = []koreLexRegexRule{
		createRegexRule("~>", KSEQ, nil),
		createRegexRule(".::K", DOTK, nil),
		createRegexRule(".K", DOTK, nil),
		createRegexRule("\\(", '(', nil),
		createRegexRule("\\)", ')', nil),
		createRegexRule(",", ',', nil),
		createRegexRule(".::KList", DOTKLIST, nil),
		createRegexRule(".KList", DOTKLIST, nil),
		createRegexRule("#token", TOKENLABEL, nil),
		createRegexRule("#klabel", KLABELLABEL, nil),
		createRegexRule("[#a-z]([a-zA-Z0-9])*", KLABEL, nil),
		createRegexRule("[A-Z]([a-zA-Z0-9'_])*", KVARIABLE, nil),
		createRegexRule("\"([^\"\\\\]|(\\\\.))*\"", STRING, UnescapeKString), // matches any string enclosed in ", allowing \"
		createRegexRule("`([^`\\\\]|(\\\\.))*`", KLABEL, UnescapeKLabel),     // matches any string enclosed in `, allowing \`
	}
}

// The parser uses the type <prefix>Lex as a lexer. It must provide
// the methods Lex(*<prefix>SymType) int and Error(string).
type koreLexerImpl struct {
	line []byte
}

// The parser calls this method to get each new token.
// The current implementation skips all whitespaces, then looks for the regex rule that gives the longest match
func (x *koreLexerImpl) Lex(yylval *koreSymType) int {
	c, size := x.firstRune()
	for isWhitespace(c) { // just skip all whitespaces
		x.line = x.line[size:]
		c, size = x.firstRune()
	}
	if c == eof {
		return eof
	}
	return x.regexLex(c, yylval)
}

func (x *koreLexerImpl) regexLex(c rune, yylval *koreSymType) int {
	maxLength := 0
	maxIndex := -1
	for i, krr := range koreLexRegexRules {
		loc := krr.regexp.FindIndex(x.line)
		if loc != nil {
			if loc[0] != 0 {
				log.Fatal("Regex should only match the start of the remaining sequence.")
			}
			if loc[1] > maxLength {
				maxLength = loc[1]
				maxIndex = i
			}
		}
	}

	if maxIndex == -1 {
		log.Fatalf("None of the rules matches the remaining string «%s».", x.line)
	}

	yylval.str = x.line[0:maxLength]

	// apply post-processing (if it's the case)
	if koreLexRegexRules[maxIndex].postProcessing != nil {
		yylval.str = koreLexRegexRules[maxIndex].postProcessing(yylval.str)
	}
	//fmt.Printf("token:%d  i:%d   len:%d    string:%s \n",
	//	koreLexRegexRules[maxIndex].token, maxIndex, maxLength, yylval.str)

	// trim line
	x.line = x.line[maxLength:]

	return koreLexRegexRules[maxIndex].token
}

// Returns the first rune in the slice
func (x *koreLexerImpl) firstRune() (rune, int) {
	if len(x.line) == 0 {
		return eof, 0
	}
	c, size := utf8.DecodeRune(x.line)
	if c == utf8.RuneError && size == 1 {
		log.Fatal("invalid utf8")
	}
	return c, size
}

// The parser calls this method on a parse error.
func (x *koreLexerImpl) Error(s string) {
	log.Printf("parse error: %s", s)
}
