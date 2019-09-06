package main

import (
	"fmt"

	"github.com/lucbarr/sslang/lexical"
)

func main() {
	program := []byte("potato 102349 potato2")
	lexer := lexical.NewLexer(program)

	tokens, err := lexer.Run()

	fmt.Println(tokens, err)
}
