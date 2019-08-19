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
