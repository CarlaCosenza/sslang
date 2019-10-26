package syntatical

import (
	"fmt"
	"strconv"

	"github.com/lucbarr/sslang/lexical"
	"github.com/lucbarr/sslang/nonterminals"
	"github.com/lucbarr/sslang/semantics"
)

// Parser parses the program
type Parser struct {
	actionTable [][]string

	stateStack []int
}

// NewParser returns a parser from action table
func NewParser() (*Parser, error) {
	return &Parser{
		actionTable: actionTable,
		stateStack:  []int{0},
	}, nil
}

// Run runs the lexical analysis
func (p *Parser) Run(lexer *lexical.Lexer, out string) error {
	state := 0
	currentToken, _ := lexer.NextToken()
	action := p.actionTable[state][currentToken]

	sem := semantics.NewAnalyser(lexer, out)
	defer sem.Close()

	for !accept(action) {
		state, ok := shift(action)
		if ok {
			p.stateStack = append(p.stateStack, state)

			currentToken, _ = lexer.NextToken()
			action = p.actionTable[state][currentToken]

			continue
		}

		rule, ok := reduce(action)
		if ok {
			amountToPop := nonterminals.RuleNumberOfTokens[rule-1]
			p.stateStack = p.stateStack[:len(p.stateStack)-amountToPop]

			temporaryState := p.stateStack[len(p.stateStack)-1]

			leftToken := nonterminals.RuleLeftTokens[rule-1]
			stateString := p.actionTable[temporaryState][leftToken]

			state, err := strconv.Atoi(stateString)
			if err != nil {
				return err
			}

			p.stateStack = append(p.stateStack, state)

			action = p.actionTable[state][currentToken]

			sem.Parse(rule)
			continue
		}

		return fmt.Errorf("Syntax error at line %v", lexer.Line)
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
