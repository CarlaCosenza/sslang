package syntatical

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/lucbarr/sslang/lexical"
	"github.com/lucbarr/sslang/semantics"
)

// Parser parses the program
type Parser struct {
	actionTable [][]string
	header      []string

	stateStack []int

	out []int
}

// NewParser returns a parser from action table
func NewParser(actionTableFile string) (*Parser, error) {
	actionTable, header, err := buildActionTableFromFile(actionTableFile)
	if err != nil {
		return nil, err
	}

	for i := lexical.Array; i <= lexical.EOF; i++ {
		fmt.Printf("%v ", header[i])
	}

	for i := P; i <= NUM; i++ {
		fmt.Printf("%v ", header[i])
	}

	return &Parser{
		actionTable: actionTable,
		header:      header,
		stateStack:  []int{0},
		out:         []int{},
	}, nil
}

func buildActionTableFromFile(file string) ([][]string, []string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, nil, err
	}

	reader := csv.NewReader(f)
	reader.Comma = '\t'

	actionTable := [][]string{}

	header, _ := reader.Read() // skip header
	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}

		actionTable = append(actionTable, line[1:])
	}

	return actionTable, header, nil
}

// Run runs the lexical analysis
func (p *Parser) Run(lexer *lexical.Lexer) error {
	state := 0
	currentToken, err := lexer.NextToken()
	action := p.actionTable[state][TokenToAction[currentToken]]

	semantics := semantics.Analyser{}

	for {
		state = p.Top()
		action = p.actionTable[state][TokenToAction[currentToken]]

		s, isShift := shift(action)
		r, isReduction := reduce(action)
		if isShift {
			p.stateStack = append(p.stateStack, s)
			currentToken, err = lexer.NextToken()
		} else if isReduction {
			p.stateStack = p.stateStack[:len(p.stateStack)-p.Len(int(r))]

			top := p.Top()
			action = p.actionTable[top][p.Left(int(r))] // TODO

			state, err := strconv.Atoi(action)
			fmt.Println("reduce", action, top, p.Left(int(r)))
			fmt.Println("row, column", top, p.header[p.Left(int(r))])
			if err != nil {
				return err
			}

			p.stateStack = append(p.stateStack, state)

			semantics.Parse(r)
		} else {
			return fmt.Errorf("Syntax error at line %v", lexer.Line)
		}

		if err != nil && err != io.EOF {
			return err
		}

		if err == io.EOF {
			break
		}

		if accept(action) {
			break
		}
	}

	return nil
}

func accept(s string) bool {
	return s == "acc"
}

func reduce(s string) (Rule, bool) {
	if len(s) == 0 {
		return -1, false
	}

	if s[0] != 'r' {
		return -1, false
	}

	n, err := strconv.Atoi(s[1:])

	if err != nil {
		return -1, false
	}

	return Rule(n - 1), true
}

func shift(s string) (int, bool) {
	if len(s) == 0 {
		return -1, false
	}

	if s[0] != 's' {
		return -1, false
	}

	n, err := strconv.Atoi(s[1:])

	if err != nil {
		return -1, false
	}

	return n, true
}

func (p *Parser) Len(r int) int {
	return ruleTable[r-1][0]
}

func (p *Parser) Left(r int) int {
	return ruleTable[r-1][1]
}

func (p *Parser) Top() int {
	return p.stateStack[len(p.stateStack)-1]
}
