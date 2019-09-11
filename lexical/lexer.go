package lexical

// Lexical analyser implementation, see book @ page 4

import (
	"bytes"
	"io"
	"regexp"
	"strings"
	"unicode"
)

// Lexer analyse if a set of tokens is part of our language and
// parse its tokens stream
type Lexer struct {
	// TODO: track lines we are reading
	scapeRunes         []rune
	reservedWordTokens map[string]int

	program *bytes.Buffer

	isLiteralReg *regexp.Regexp

	identifiers map[string]int

	intConstants    []string
	stringConstants []string
	runeConstants   []rune
}

// NewLexer builds an analyser
func NewLexer(program []byte) *Lexer {
	programBuffer := bytes.NewBuffer(program)
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
		isLiteralReg:    regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9]+"),
		identifiers:     map[string]int{},
		intConstants:    []string{},
		stringConstants: []string{},
		runeConstants:   []rune{},
		program:         programBuffer,
	}
}

func (a *Lexer) isLiteral(s string) bool {
	return a.isLiteralReg.MatchString(s)
}

// Run runs the lexical analysis
func (a *Lexer) Run() ([]int, error) {
	tokens := []int{}
	for {
		token, err := a.nextToken(a.program)
		if err != nil && err != io.EOF {
			return nil, err
		}

		if err == io.EOF {
			break
		}

		tokens = append(tokens, token)
	}
	return tokens, nil
}

func (a *Lexer) nextToken(buf *bytes.Buffer) (int, error) {
	var nextRune rune
	var err error
	token := UNKNOWN

	for {
		nextRune, _, err = buf.ReadRune()
		if err != nil {
			return -1, err
		}

		if !unicode.IsSpace(nextRune) {
			break
		}
	}

	if isAlpha(nextRune) {
		text, err := parseWord(buf, func(r rune) bool {
			return isAlpha(r) || r == '_'
		})

		if err != nil {
			return -1, err
		}

		reservedToken, ok := a.reservedWordTokens[text]
		if !ok {
			a.registerIdentifier(text)
			token = ID
		} else {
			token = reservedToken
		}

	} else if isDigit(nextRune) {
		text, err := parseWord(buf, func(r rune) bool {
			return isDigit(r)
		})
		if err != nil {
			return -1, err
		}

		a.intConstants = append(a.intConstants, text)
		token = Numeral
	} else if nextRune == '"' {
		buf.ReadRune()
		text, err := parseWord(buf, func(r rune) bool {
			return r != '"'
		})

		if err != nil {
			return -1, err
		}

		token = Stringval
		a.stringConstants = append(a.stringConstants, text)
	} else {
		// TODO: giant switch case (we could make this simpler by making a map[rune]func(rune)(int,error))
	}

	return token, nil
}

func parseWord(buf *bytes.Buffer, criteria func(rune) bool) (string, error) {
	var sb strings.Builder
	var err error

	err = buf.UnreadRune()
	if err != nil {
		return "", err
	}

	nextToken, _, err := buf.ReadRune()
	if err != nil {
		return "", err
	}

	for criteria(nextToken) && err != io.EOF {
		sb.WriteRune(nextToken)

		nextToken, _, err = buf.ReadRune()
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
	return unicode.IsLetter(r)
}

func isAlphaNumeric(r rune) bool {
	return isAlpha(r) || isDigit(r)
}

func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

func (a *Lexer) registerIdentifier(s string) {
	secondaryToken, ok := a.identifiers[s]

	if !ok {
		secondaryToken = len(a.identifiers)
		a.identifiers[s] = secondaryToken
	}
}
