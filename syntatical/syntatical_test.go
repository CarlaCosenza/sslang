package syntatical

import (
	"fmt"
	"testing"

	"github.com/lucbarr/sslang/lexical"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	tt := map[string]struct {
		tokens []int
		err    error
	}{
		"test sample program syntactically correct": {
			tokens: []int{lexical.Function, lexical.ID, lexical.LeftParenthesis, lexical.ID,
				lexical.Colon, lexical.Integer, lexical.RightParenthesis, lexical.Colon,
				lexical.Integer, lexical.LeftBraces, lexical.Var, lexical.ID, lexical.Colon,
				lexical.Integer, lexical.Semicolon, lexical.Var, lexical.ID, lexical.Colon,
				lexical.Integer, lexical.Semicolon, lexical.Var, lexical.ID, lexical.Colon,
				lexical.Integer, lexical.Semicolon, lexical.ID, lexical.Equals, lexical.Numeral, lexical.Semicolon,
				lexical.ID, lexical.Equals, lexical.Numeral, lexical.Semicolon, lexical.RightBraces, lexical.EOF},
			err: nil,
		},
		"test sample program syntactically incorrect": {
			tokens: []int{lexical.Function, lexical.ID, lexical.LeftParenthesis, lexical.ID,
				lexical.Colon, lexical.Integer, lexical.RightParenthesis, lexical.Colon,
				lexical.Integer, lexical.LeftBraces, lexical.Var, lexical.ID, lexical.Colon,
				lexical.Integer, lexical.Semicolon, lexical.Var, lexical.ID, lexical.Colon,
				lexical.Integer, lexical.Semicolon, lexical.Var, lexical.ID, lexical.Colon,
				lexical.Integer, lexical.Semicolon, lexical.ID, lexical.Equals, lexical.Numeral,
				lexical.ID, lexical.Equals, lexical.Numeral, lexical.RightBraces, lexical.EOF},
			err: fmt.Errorf("Syntax error"),
		},
	}

	for name, table := range tt {
		t.Run(name, func(t *testing.T) {
			syntatical, _ := NewParser("action_table.csv")
			err := syntatical.Run(table.tokens)
			assert.Equal(t, table.err, err)
		})
	}
}
