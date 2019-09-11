package lexical

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseWord(t *testing.T) {
	tt := map[string]struct {
		buf      *bytes.Buffer
		criteria func(rune) bool

		text string
		err  error
	}{
		"test parse identifier with eof": {
			buf: bytes.NewBufferString("potato_"),
			criteria: func(r rune) bool {
				return isAlpha(r) || r == '_'
			},

			text: "potato_",
			err:  nil,
		},
		"test parse identifier no EOF": {
			buf: bytes.NewBufferString("potato_ &*$#!@"),
			criteria: func(r rune) bool {
				return isAlpha(r) || r == '_'
			},

			text: "potato_",
			err:  nil,
		},
		"test parse digit": {
			buf: bytes.NewBufferString("123849 @#41"),
			criteria: func(r rune) bool {
				return isDigit(r)
			},

			text: "123849",
			err:  nil,
		},
	}

	for name, table := range tt {
		t.Run(name, func(t *testing.T) {
			table.buf.ReadRune()
			text, err := parseWord(table.buf, table.criteria)

			assert.Equal(t, table.text, text)
			assert.Equal(t, table.err, err)
		})
	}
}

func TestNextToken(t *testing.T) {
	tt := map[string]struct {
		buf *bytes.Buffer

		token int
		err   error
	}{
		"test parse identifier": {
			buf: bytes.NewBufferString("foo_"),

			token: ID,
			err:   nil,
		},
		"test parse identifier after a hell lot of whitespace": {
			buf: bytes.NewBufferString("    foo_"),

			token: ID,
			err:   nil,
		},
		"test parse numeral": {
			buf: bytes.NewBufferString("1023498"),

			token: Numeral,
			err:   nil,
		},
	}

	lexer := NewLexer([]byte{})

	for name, table := range tt {
		t.Run(name, func(t *testing.T) {
			token, err := lexer.nextToken(table.buf)
			assert.Equal(t, table.err, err)
			assert.Equal(t, table.token, token)
		})
	}
}

func TestRun(t *testing.T) {
	tt := map[string]struct {
		program string

		tokens []int
		err    error
	}{
		"test id and numeral": {
			program: "potato 102349 potato2",
			tokens:  []int{ID, Numeral, ID},
			err:     nil,
		},
		"test id reserved words": {
			program: "var potato integer",
			tokens:  []int{Var, ID, Integer},
			err:     nil,
		},
	}

	for name, table := range tt {
		t.Run(name, func(t *testing.T) {
			program := []byte(table.program)
			lexer := NewLexer(program)

			tokens, err := lexer.Run()

			assert.Equal(t, table.tokens, tokens)
			assert.Equal(t, table.err, err)
		})
	}
}
