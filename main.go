package main

import (
	"github.com/lucbarr/sslang/lexical"
	"github.com/lucbarr/sslang/syntatical"
)

func main() {
	program := []byte(`function main(arg:integer):integer
{
	var a:integer;
	var b:integer;
	var c:integer;
	b = 1;
	c = 2;
}`)
	lexer := lexical.NewLexer(program)

	lexer.Run()

	syntatical.NewParser("syntatical/action_table.csv")
}
