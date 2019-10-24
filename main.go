package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/lucbarr/sslang/lexical"
	"github.com/lucbarr/sslang/syntatical"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(`No SSL file provided.
usage : ./sslang example.ssl`)
		return
	}

	file := args[0]

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Could not read file %v\n", file)
	}

	lexer := lexical.NewLexer(bytes)

	parser, err := syntatical.NewParser("syntatical/action_table.csv")
	if err != nil {
		fmt.Println(err)
	}

	err = parser.Run(lexer)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v compiled successfully!!\n", file)
	}

}
