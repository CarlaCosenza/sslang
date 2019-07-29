package lexical

import "regexp"

type Analyser struct {
	scapeRunes    []rune
	reservedWords map[string]int
}

func NewAnalyser() *Analyser {
	return &Analyser{
		scapeRunes: []rune{' ', '\t', '\v', '\n', '\r', '\f'},
		reservedWords: map[string]int{
			"integer":  INTEGER,
			"char":     CHAR,
			"type":     TYPE,
			"struct":   STRUCT,
			"function": FUNCTION,
			"var":      VARIABLE,
			"boolean":  BOOLEAN,
			"of":       OF,
			"do":       DO,
			"break":    BREAK,
			"continue": CONTINUE,
			"if":       IF,
			"else":     ELSE,
			"true":     TRUE,
			"false":    FALSE,
		},
	}
}

const (
	START = iota
)

const (
	INTEGER = iota
	CHAR
	BOOLEAN
	STRING
	IDENTIFIER
	TYPE
	ARRAY
	NUMERAL
	STRUCT
	WHILE
	FUNCTION
	VARIABLE
	OF
	CONTINUE
	BREAK
	DO
	IF
	ELSE
	TRUE
	FALSE
)

type Automata struct {
	// a transition maps a state to a function that, given an input,
	// returns the next state
	transitions map[int]func(...rune) int

	// final tells if state is final
	final map[int]bool

	// token returns the token for a given final state
	token map[int]int

	current int
}

func isLiteral(s string) bool {
	reg := regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9]+")
	return reg.MatchString(s)
}

func NewAutomata() *Automata {
	return &Automata{}
}
