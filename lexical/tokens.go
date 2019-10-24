package lexical

// Tokens
const (
	// this is not my language bruh
	UNKNOWN = -1
	// reserved words tokens
	Integer = iota
	Char
	Boolean
	String
	Type
	Equals
	Array
	LeftSquare
	RightSquare
	Of
	Struct
	LeftBraces
	RightBraces
	Semicolon
	Colon
	Function
	LeftParenthesis
	RightParenthesis
	Comma
	Var
	If
	Else
	While
	Do
	Break
	Continue
	And
	Or
	LessThan
	GreaterThan
	LessOrEqual
	GreaterOrEqual
	EqualEqual
	NotEqual
	Plus
	Minus
	Times
	Divide
	PlusPlus
	MinusMinus
	Not
	Dot
	ID
	True
	False

	// symbols

	// regular tokens
	Character
	Stringval
	Numeral

	// end of file help token
	EOF
)

const UNKNOWN = -1

//TokenToString is a toStr equivalent utility map
var TokenToString = map[int]string{
	Array:    "Array",
	Boolean:  "Boolean",
	Break:    "Break",
	Char:     "Char",
	Continue: "Continue",
	Do:       "Do",
	Else:     "Else",
	Function: "Function",
	If:       "If",
	Integer:  "Integer",
	Of:       "Of",
	String:   "String",
	Struct:   "Struct",
	True:     "True",
	False:    "False",
	Type:     "Type",
	Var:      "Var",
	While:    "While",

	Colon:            "Colon",
	Semicolon:        "Semicolon",
	Comma:            "Comma",
	Equals:           "Equals",
	LeftSquare:       "LeftSquare",
	RightSquare:      "RightSquare",
	LeftBraces:       "LeftBraces",
	RightBraces:      "RightBraces",
	LeftParenthesis:  "LeftParenthesis",
	RightParenthesis: "RightParenthesis",
	And:              "And",
	Or:               "Or",
	LessThan:         "LessThan",
	GreaterThan:      "GreaterThan",
	LessOrEqual:      "LessOrEqual",
	GreaterOrEqual:   "GreaterOrEqual",
	NotEqual:         "NotEqual",
	EqualEqual:       "EqualEqual",
	Plus:             "Plus",
	PlusPlus:         "PlusPlus",
	Minus:            "Minus",
	MinusMinus:       "MinusMinus",
	Times:            "Times",
	Divide:           "Divide",
	Dot:              "Dot",
	Not:              "Not",
	// regular tokens : "//",
	Character: "Character",
	Numeral:   "Numeral",
	Stringval: "Stringval",
	ID:        "ID",
	// this is not my language bruh : "//",
	UNKNOWN: "UNKNOWN",
}

// ReservedWordTokens maps reserved words strings into its tokens
var ReservedWordTokens = map[string]int{
	"array":    Array,
	"boolean":  Boolean,
	"break":    Break,
	"char":     Char,
	"continue": Continue,
	"do":       Do,
	"else":     Else,
	"function": Function,
	"if":       If,
	"integer":  Integer,
	"of":       Of,
	"string":   String,
	"struct":   Struct,
	"true":     True,
	"false":    False,
	"type":     Type,
	"var":      Var,
	"while":    While,
}
