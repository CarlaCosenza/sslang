package lexical

// Lexical analyser implementation, see book @ page 4

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
	"unicode"
)

// Lexer analyse if a set of tokens is part of our language and
// parse its tokens stream
type Lexer struct {
	program *bytes.Buffer

	isLiteralReg *regexp.Regexp

	identifiers map[string]int

	intConstants    []string
	stringConstants []string
	runeConstants   []rune

	line int
}

// NewLexer builds an analyser
func NewLexer(program []byte) *Lexer {
	programBuffer := bytes.NewBuffer(program)
	return &Lexer{
		isLiteralReg:    regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9]+"),
		identifiers:     map[string]int{},
		intConstants:    []string{},
		stringConstants: []string{},
		runeConstants:   []rune{},
		program:         programBuffer,
		line:            0,
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
			tokens = append(tokens, EOF)
			break
		}

		tokens = append(tokens, token)
	}
	return tokens, nil
}

func (a *Lexer) nextToken(buf *bytes.Buffer) (int, error) {
	var nextRune, nextRune2 rune
	var err error
	token := UNKNOWN

	for {
		nextRune, _, err = buf.ReadRune()
		if err != nil {
			return -1, err
		}

		if nextRune == '\n' {
			a.line++
		}

		if !unicode.IsSpace(nextRune) {
			break
		}
	}

	if isAlpha(nextRune) {
		text, err := parseWord(buf, func(r rune) bool {
			return isAlphaNumeric(r) || r == '_'
		})

		if err != nil {
			return -1, err
		}

		reservedToken, ok := ReservedWordTokens[text]
		if !ok {
			a.registerIdentifier(text)
			token = ID
		} else {
			token = reservedToken
		}
		buf.UnreadRune()

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
		switch nextRune {
		case ':':
			token = Colon
			break
		case ';':
			token = Semicolon
			break
		case ',':
			token = Comma
			break
		case '*':
			token = Times
			break
		case '/':
			token = Divide
			break
		case '.':
			token = Dot
			break
		case '[':
			token = LeftSquare
			break
		case ']':
			token = RightSquare
			break
		case '{':
			token = LeftBraces
			break
		case '}':
			token = RightBraces
			break
		case '(':
			token = LeftParenthesis
			break
		case ')':
			token = RightParenthesis
			break
		case '\'':
			runeCtt, _, err := buf.ReadRune()
			if err != nil {
				return -1, err
			}

			expectedQuotes, _, err := buf.ReadRune()
			if err != nil {
				return -1, err
			}

			if expectedQuotes != '\'' {
				return -1, fmt.Errorf("Expected quotes")
			}

			token = Character
			a.runeConstants = append(a.runeConstants, runeCtt)
			break
		case '&':
			nextRune2, _, err = buf.ReadRune()
			if err != nil {
				return -1, err
			}
			if nextRune2 != '&' {
				return -1, errors.New("Invalid character")
			}
			token = And
			break
		case '|':
			nextRune2, _, err = buf.ReadRune()
			if err != nil {
				return -1, err
			}
			if nextRune2 != '|' {
				return -1, errors.New("Invalid character")
			}
			token = Or
			break
		case '=':
			nextRune2, _, err = buf.ReadRune()
			if err != nil {
				if err != io.EOF {
					return -1, err
				}
				token = Equals
				break
			}
			if nextRune2 != '=' {
				err = buf.UnreadRune()
				if err != nil {
					return -1, err
				}
				token = Equals
			} else {
				token = EqualEqual
			}
			break
		case '<':
			nextRune2, _, err = buf.ReadRune()
			if err != nil {
				if err != io.EOF {
					return -1, err
				}
				token = LessThan
				break
			}
			if nextRune2 != '=' {
				err = buf.UnreadRune()
				if err != nil {
					return -1, err
				}
				token = LessThan
			} else {
				token = LessOrEqual
			}
		case '>':
			nextRune2, _, err = buf.ReadRune()
			if err != nil {
				if err != io.EOF {
					return -1, err
				}
				token = GreaterThan
				break
			}
			if nextRune2 != '=' {
				err = buf.UnreadRune()
				if err != nil {
					return -1, err
				}
				token = GreaterThan
			} else {
				token = GreaterOrEqual
			}
		case '!':
			nextRune2, _, err = buf.ReadRune()
			if err != nil {
				if err != io.EOF {
					return -1, err
				}

				token = Not
				break
			}
			if nextRune2 != '=' {
				err = buf.UnreadRune()
				if err != nil {
					return -1, err
				}
				token = Not
			} else {
				token = NotEqual
			}
		case '+':
			nextRune2, _, err = buf.ReadRune()
			if err != nil {
				if err != io.EOF {
					return -1, err
				}
				token = Plus
				break
			}
			if nextRune2 != '+' {
				err = buf.UnreadRune()
				if err != nil {
					return -1, err
				}
				token = Plus
			} else {
				token = PlusPlus
			}
		case '-':
			nextRune2, _, err = buf.ReadRune()
			if err != nil {
				if err != io.EOF {
					return -1, err
				}

				token = Minus
				break
			}
			if nextRune2 != '-' {
				err = buf.UnreadRune()
				if err != nil {
					return -1, err
				}
				token = Minus
			} else {
				token = MinusMinus
			}
		}
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

// Identifiers retrieves the identifiers
func (a *Lexer) Identifiers() map[string]int {
	return a.identifiers
}
