package main

import (
	"fmt"

	"github.com/lucbarr/sslang/lexical"
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

	tokens, err := lexer.Run()

	fmt.Println(tokens, err)

	for _, token := range tokens {
		fmt.Printf("%v ", lexical.TokenToString[token])
	}

	fmt.Println(lexer.Identifiers())

}
