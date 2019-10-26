package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/lucbarr/sslang/lexical"
	"github.com/lucbarr/sslang/syntatical"
)

var (
	outFile string
)

func init() {
	flag.StringVar(&outFile, "out", "out", "-out <output file name>")
	flag.Parse()
}

func main() {
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println(`No SSL file provided.
usage : ./sslang -out=outputFile.vm example.ssl `)
		return
	}

	file := args[0]

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Could not read file %v\n", file)
		return
	}

	lexer := lexical.NewLexer(bytes)

	parser, err := syntatical.NewParser()
	if err != nil {
		fmt.Println(err)
	}

	err = parser.Run(lexer, outFile)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v compiled successfully!!\n", file)
	}

}
