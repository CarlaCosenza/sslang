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
	lexical.Integer:          0,
	lexical.Char:             1,
	lexical.Boolean:          2,
	lexical.String:           3,
	lexical.Type:             4,
	lexical.Equals:           5,
	lexical.Array:            6,
	lexical.LeftSquare:       7,
	lexical.RightSquare:      8,
	lexical.Of:               9,
	lexical.Struct:           10,
	lexical.LeftBraces:       11,
	lexical.RightBraces:      12,
	lexical.Semicolon:        13,
	lexical.Colon:            14,
	lexical.Function:         15,
	lexical.LeftParenthesis:  16,
	lexical.RightParenthesis: 17,
	lexical.Comma:            18,
	lexical.Var:              19,
	lexical.If:               20,
	lexical.Else:             21,
	lexical.While:            22,
	lexical.Do:               23,
	lexical.Break:            24,
	lexical.Continue:         25,
	lexical.And:              26,
	lexical.Or:               27,
	lexical.LessThan:         28,
	lexical.GreaterThan:      29,
	lexical.LessOrEqual:      30,
	lexical.GreaterOrEqual:   31,
	lexical.EqualEqual:       32,
	lexical.NotEqual:         33,
	lexical.Plus:             34,
	lexical.Minus:            35,
	lexical.Times:            36,
	lexical.Divide:           37,
	lexical.PlusPlus:         38,
	lexical.MinusMinus:       39,
	lexical.Not:              40,
	lexical.Dot:              41,
	lexical.ID:               42,
	lexical.True:             43,
	lexical.False:            44,
	lexical.Character:        45,
	lexical.Stringval:        46,
	lexical.Numeral:          47,
	lexical.EOF:              48,
	PLINE:                    49,

	// Goto actions
	P:     50,
	LDE:   51,
	DE:    52,
	T:     53,
	DT:    54,
	DC:    55,
	DF:    56,
	LP:    57,
	B:     58,
	LDV:   59,
	LS:    60,
	DV:    61,
	LI:    62,
	S:     63,
	U:     64,
	M:     65,
	E:     66,
	L:     67,
	R:     68,
	Y:     69,
	F:     70,
	LE:    71,
	LV:    72,
	ID:    73,
	TRUE:  74,
	FALSE: 75,
	CHR:   76,
	STR:   77,
	NUM:   78,
}

const ruleLeftTokens = []int{P, LDE, LDE, DE, DE, T, T, T, T, T, DT, DT, DT, DC, DC, DF, LP, LP, B, LDV, LDV, LS, LS, DV, LI, LI, S, S, U, U, M, M, M, M, M, M, M, E, E, E, L, L, L, L, L, L, L, R, R, R, Y, Y, Y, F, F, F, F, F, F, F, F, F, F, F, F, F, F, LE, LE, LV, LV, LV, ID, TRUE, FALSE, CHR, STR, NUM}
const ruleNumberOfTokens = []int{1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 9, 7, 4, 5, 3, 8, 5, 3, 4, 2, 1, 2, 1, 5, 3, 1, 1, 1, 5, 7, 7, 5, 7, 1, 4, 2, 2, 3, 3, 1, 3, 3, 3, 3, 3, 3, 1, 3, 3, 1, 3, 3, 1, 1, 2, 2, 2, 2, 3, 4, 2, 2, 1, 1, 1, 1, 1, 3, 1, 3, 4, 1, 1, 1, 1, 1, 1, 1}
