package syntatical

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

// Parser parses the program
type Parser struct {
	actionTable [][]string

	stateStack []int

	out []int
}

// NewParser returns a parser from action table
func NewParser(actionTableFile string) (*Parser, error) {
	actionTable, err := buildActionTableFromFile(actionTableFile)
	if err != nil {
		return nil, err
	}

	return &Parser{
		actionTable: actionTable,
		stateStack:  []int{0},
		out:         []int{},
	}, nil
}

func buildActionTableFromFile(file string) ([][]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(f)
	reader.Comma = '\t'

	actionTable := [][]string{}

	reader.Read() // skip header
	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}

		actionTable = append(actionTable, line[1:])
	}

	return actionTable, nil
}

// Run runs the lexical analysis
func (p *Parser) Run(tokens []int) error {
	state := 0
	iToken := 0
	currentToken := tokens[iToken]

	nextToken := func() {
		iToken++
		currentToken = tokens[iToken]
	}

	action := p.actionTable[state][TokenToAction[currentToken]]

	for {

		if accept(action) {
			break
		}

		state, ok := shift(action)
		if ok {
			p.stateStack = append(p.stateStack, state)

			nextToken()
			action = p.actionTable[state][TokenToAction[currentToken]]

			continue
		}

		rule, ok := reduce(action)
		if ok {
			amountToPop := ruleNumberOfTokens[rule]
			p.stateStack = p.stateStack[:len(p.stateStack)-amountToPop]

			temporaryState := p.stateStack[len(p.stateStack)-1]

			leftToken := ruleLeftTokens[rule]
			goTo := TokenToAction[leftToken]
			stateString := p.actionTable[temporaryState][goTo]

			state, err := strconv.Atoi(stateString)
			if err != nil {
				return err
			}

			p.stateStack = append(p.stateStack, state)
		}
	}

	return nil
}

func accept(s string) bool {
	return s == "acc"
}

func reduce(s string) (int, bool) {
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

	return n, true
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
