package syntatical

import (
	"fmt"
	"testing"

	"github.com/lucbarr/sslang/lexical"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	tt := map[string]struct {
		program string
		err     error
	}{
		"test sample program syntactically correct": {
			program: `
function main(arg:integer):integer
{
	var a:integer;
	var b:integer;
	var c:integer;
	b = 1;
	c = 2;
}`,
			err: nil,
		},
		"test sample program syntactically incorrect": {
			program: `
function main(arg:integer):integer
{
	var a:integer;
	var b:integer;
	var c:integer
	b = 1;
	c = 2;
}`,
			err: fmt.Errorf("Syntax error at line 6"),
		},
	}

	for name, table := range tt {
		t.Run(name, func(t *testing.T) {
			syntatical, _ := NewParser("action_table.csv")
			lexer := lexical.NewLexer([]byte(table.program))
			err := syntatical.Run(lexer)
			assert.Equal(t, table.err, err)
		})
	}
}
