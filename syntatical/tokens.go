package syntatical

import "github.com/lucbarr/sslang/lexical"

// Language tokens
const (
	P = lexical.UNKNOWN + iota + 1
	LDE
	DE
	DF
	DT
	T
	DC
	LI
	LP
	B
	LDV
	DV
	LS
	S
	E
	LV
	L
	R
	Y
	F
	LE
	ID
	TRUE
	FALSE
	CHR
	STR
	NUM
	PLINE
	M
	U
)

// TokenToAction is a table to help hop through the action table
var TokenToAction = map[int]int{
	// Token actions
	lexical.Integer:          1,
	lexical.Char:             2,
	lexical.Boolean:          3,
	lexical.String:           4,
	lexical.Type:             5,
	lexical.Equals:           6,
	lexical.Array:            7,
	lexical.LeftSquare:       8,
	lexical.RightSquare:      9,
	lexical.Of:               10,
	lexical.Struct:           11,
	lexical.LeftBraces:       12,
	lexical.RightBraces:      13,
	lexical.Semicolon:        14,
	lexical.Colon:            15,
	lexical.Function:         16,
	lexical.LeftParenthesis:  17,
	lexical.RightParenthesis: 18,
	lexical.Comma:            19,
	lexical.Var:              20,
	lexical.If:               21,
	lexical.Else:             22,
	lexical.While:            23,
	lexical.Do:               24,
	lexical.Break:            25,
	lexical.Continue:         26,
	lexical.And:              27,
	lexical.Or:               28,
	lexical.LessThan:         29,
	lexical.GreaterThan:      30,
	lexical.LessOrEqual:      31,
	lexical.GreaterOrEqual:   32,
	lexical.EqualEqual:       33,
	lexical.NotEqual:         34,
	lexical.Plus:             35,
	lexical.Minus:            36,
	lexical.Times:            37,
	lexical.Divide:           38,
	lexical.PlusPlus:         39,
	lexical.MinusMinus:       40,
	lexical.Not:              41,
	lexical.Dot:              42,
	lexical.ID:               43,
	lexical.True:             44,
	lexical.False:            45,
	lexical.Character:        46,
	lexical.Stringval:        47,
	lexical.Numeral:          48,
	lexical.EOF:              49,

	// Goto actions
	PLINE: 50,
	P:     51,
	LDE:   52,
	DE:    53,
	T:     54,
	DT:    55,
	DC:    56,
	DF:    57,
	LP:    58,
	B:     59,
	LDV:   60,
	LS:    61,
	DV:    62,
	LI:    63,
	S:     64,
	U:     65,
	M:     66,
	E:     67,
	L:     68,
	R:     69,
	Y:     70,
	F:     71,
	LE:    72,
	LV:    73,
	ID:    74,
	TRUE:  75,
	FALSE: 76,
	CHR:   77,
	STR:   78,
	NUM:   79,
}
