package syntatical

import (
	"fmt"
	"testing"

	"github.com/lucbarr/sslang/lexical"
	"github.com/stretchr/testify/assert"
)

func TestShift(t *testing.T) {
	tt := map[string]struct {
		s string

		n  int
		ok bool
	}{
		"test invalid rule": {
			s:  "1324",
			n:  -1,
			ok: false,
		},
		"test valid rule": {
			s:  "s123",
			n:  123,
			ok: true,
		},
		"test empty rule": {
			s:  "",
			n:  -1,
			ok: false,
		},
	}

	for name, table := range tt {
		t.Run(name, func(t *testing.T) {
			n, ok := shift(table.s)

			assert.Equal(t, table.n, n)
			assert.Equal(t, table.ok, ok)
		})
	}
}

func TestReduce(t *testing.T) {
	tt := map[string]struct {
		s string

		n  int
		ok bool
	}{
		"test invalid rule": {
			s:  "1324",
			n:  -1,
			ok: false,
		},
		"test valid rule": {
			s:  "r123",
			n:  123,
			ok: true,
		},
		"test empty rule": {
			s:  "",
			n:  -1,
			ok: false,
		},
	}

	for name, table := range tt {
		t.Run(name, func(t *testing.T) {
			n, ok := reduce(table.s)

			assert.Equal(t, table.n, n)
			assert.Equal(t, table.ok, ok)
		})
	}
}

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
		"test sample big program syntactically correct": {
			program: `
type int = integer
type intArray = array[10] of int

type Pessoa = struct {
  nome : string;
  id : int;
  telefones : intArray
}

type pessoaArray = array[20] of Pessoa

function funcaoRecursiva(n : int) : integer {
  var p : integer;
  var x : intArray;
  if (n == 0)
    p = funcaoRecursiva(n-1);
}

function addPessoa(posicao : integer, pessoa : Pessoa,arr : pessoaArray) : integer {
  var a,g,i,n : int;
  var b : int;
  var c : boolean;
  var str : string;
  var carac1,carac2 : char;

  if (a == 0)
    g = funcaoRecursiva(i-1);

  a = posicao;
  c = (a>b) && (a<b) || (a <= b) || (b>= a);

  if (a > b) {
    var tmp : integer;
    tmp = funcaoRecursiva(tmp-1);
  }
}

function fibonacci(n : integer) : integer {
  var ret : integer;
  if (n >= 2)
    ret = fibonacci(n-1) + fibonacci(n-2);
	}`,
			err: nil,
		},
	}

	for name, table := range tt {
		t.Run(name, func(t *testing.T) {
			syntatical, _ := NewParser()
			lexer := lexical.NewLexer([]byte(table.program))
			err := syntatical.Run(lexer, "out")
			assert.Equal(t, table.err, err)
		})
	}
}
