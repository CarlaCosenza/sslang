package lexical

// Tokens
const (
	// reserved words tokens
	Array = iota
	Boolean
	Break
	Char
	Continue
	Do
	Else
	Function
	If
	Integer
	Of
	String
	Struct
	True
	False
	Type
	Var
	While

	// symbols
	Colon
	Semicolon
	Comma
	Equals
	LeftSquare
	RightSquare
	LeftBraces
	RightBraces
	LeftParenthesis
	RightParenthesis
	And
	Or
	LessThan
	GreaterThan
	LessOrEqual
	GreaterOrEqual
	NotEqual
	EqualEqual
	Plus
	PlusPlus
	Minus
	MinusMinus
	Times
	Divide
	Dot
	Not

	// regular tokens
	Character
	Numeral
	Stringval
	ID

	// this is not my language bruh
	UNKNOWN
)

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
