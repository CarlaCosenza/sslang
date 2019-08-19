package lexical

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// Lexer analyse if a set of tokens is part of our language and
// parse its tokens
type Lexer struct {
	scapeRunes         []rune
	reservedWordTokens map[string]int

	program *bytes.Buffer

	isLiteralReg *regexp.Regexp

	identifiers map[string]int

	intConstants    []int
	stringConstants []string
	runeConstants   []rune
}

// NewLexer builds an analyser
func NewLexer() *Lexer {
	return &Lexer{
		reservedWordTokens: map[string]int{
			"integer":  Integer,
			"char":     Char,
			"type":     Type,
			"struct":   Struct,
			"function": Function,
			"var":      Var,
			"boolean":  Boolean,
			"of":       Of,
			"do":       Do,
			"break":    Break,
			"continue": Continue,
			"if":       If,
			"else":     Else,
			"true":     True,
			"false":    False,
		},
		isLiteralReg: regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9]+"),
		identifiers:  map[string]int{},
	}
}

func (a *Lexer) isLiteral(s string) bool {
	return a.isLiteralReg.MatchString(s)
}

// Run runs the lexical analysis
func (a *Lexer) Run(buf *bytes.Buffer) ([]int, error) {
	tokens := []int{}
	for token, err := a.nextToken(buf); err != io.EOF; {
		tokens = append(tokens, token)
		if err != nil {
			return nil, err
		}
	}
	return tokens, nil
}

func (a *Lexer) nextToken(buf *bytes.Buffer) (int, error) {
	var nextRune rune
	var err error
	token := UNKNOWN
	for nextRune, _, err = buf.ReadRune(); unicode.IsSpace(nextRune); {
		if err != nil {
			return -1, err
		}
	}

	if isAlpha(nextRune) {
		text, err := parseWord(nextRune, buf, func(r rune) bool {
			return isAlpha(r) || r == '_'
		})

		if err != nil {
			return -1, err
		}

		token, ok := a.reservedWordTokens[text]
		if !ok {
			return -1, fmt.Errorf("%s unknown identifier", text)
		}

		if token == ID {
			a.registerIdentifier(text)
		}
	} else if isDigit(nextRune) {
		text, err := parseWord(nextRune, buf, func(r rune) bool {
			return isDigit(r)
		})
		if err != nil {
			return -1, err
		}

		val, err := strconv.Atoi(text)
		if err != nil {
			return -1, err
		}

		a.intConstants = append(a.intConstants, val)
	} else {

	}

	return token, nil
}

func parseWord(nextRune rune, buf *bytes.Buffer, criteria func(rune) bool) (string, error) {
	var sb strings.Builder
	var err error

	for isAlpha(nextRune) || nextRune == '_' && err != io.EOF {
		sb.WriteRune(nextRune)

		nextRune, _, err = buf.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}

			return "", err
		}
	}

	return sb.String(), nil
}

func isAlpha(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' || r <= 'Z')
}

func isAlphaNumeric(r rune) bool {
	return isAlpha(r) || isDigit(r)
}

func isDigit(r rune) bool {
	return (r >= '0' && r <= '9')
}

func (a *Lexer) registerIdentifier(s string) {
	secondaryToken, ok := a.identifiers[s]

	if !ok {
		secondaryToken = len(a.identifiers)
		a.identifiers[s] = secondaryToken
	}
}
