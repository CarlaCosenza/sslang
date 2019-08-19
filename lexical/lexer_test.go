package lexical

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseWord(t *testing.T) {
	tt := map[string]struct {
		buf       *bytes.Buffer
		firstRune rune
		criteria  func(rune) bool

		text string
		err  error
	}{
		"test parse identifier with eof": {
			buf:       bytes.NewBufferString("otato_"),
			firstRune: 'p',
			criteria: func(r rune) bool {
				return isAlpha(r) || r == '_'
			},

			text: "potato_",
			err:  nil,
		},
		"test parse identifier no EOF": {
			buf:       bytes.NewBufferString("otato_ &*$#!@"),
			firstRune: 'p',
			criteria: func(r rune) bool {
				return isAlpha(r) || r == '_'
			},

			text: "potato_",
			err:  nil,
		},
		"test parse digit": {
			buf:       bytes.NewBufferString("123849 @#41"),
			firstRune: '1',
			criteria: func(r rune) bool {
				return isDigit(r)
			},

			text: "1123849",
			err:  nil,
		},
	}

	for name, table := range tt {
		t.Run(name, func(t *testing.T) {
			text, err := parseWord(table.firstRune, table.buf, table.criteria)

			assert.Equal(t, table.text, text)
			assert.Equal(t, table.err, err)
		})
	}
}
