package main

import (
	"github.com/lucbarr/sslang/lexical"
	"github.com/lucbarr/sslang/syntatical"
)

func main() {
	tokens := []int{lexical.Function, lexical.ID, lexical.LeftParenthesis, lexical.ID,
		lexical.Colon, lexical.Integer, lexical.RightParenthesis, lexical.Colon,
		lexical.Integer, lexical.LeftBraces, lexical.Var, lexical.ID, lexical.Colon,
		lexical.Integer, lexical.Semicolon, lexical.Var, lexical.ID, lexical.Colon,
		lexical.Integer, lexical.Semicolon, lexical.Var, lexical.ID, lexical.Colon,
		lexical.Integer, lexical.Semicolon, lexical.ID, lexical.Equals, lexical.Numeral,
		lexical.ID, lexical.Equals, lexical.Numeral, lexical.RightBraces, lexical.EOF}

	parser, err := syntatical.NewParser("syntatical/action_table.csv")

	if err != nil {
		panic(err)
	}

	parser.Run(tokens)
}
