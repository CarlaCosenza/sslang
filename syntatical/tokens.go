package syntatical

import "github.com/lucbarr/sslang/lexical"

// Language tokens
const (
	P = lexical.EOF + iota + 1
	LDE
	DE
	T
	DT
	DC
	DF
	LP
	B
	LDV
	LS
	DV
	LI
	S
	U
	M
	E
	L
	R
	Y
	F
	LE
	LV
	NB
	NF
	MF
	MT
	ME
	MW
	MC
	IDD
	IDU
	ID
	TRUE
	FALSE
	CHR
	STR
	NUM
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

	// Goto actions
	P:     49,
	LDE:   50,
	DE:    51,
	T:     52,
	DT:    53,
	DC:    54,
	DF:    55,
	LP:    56,
	B:     57,
	LDV:   58,
	LS:    59,
	DV:    60,
	LI:    61,
	S:     62,
	U:     63,
	M:     64,
	E:     65,
	L:     66,
	R:     67,
	Y:     68,
	F:     69,
	LE:    70,
	LV:    71,
	NB:    72,
	NF:    73,
	MF:    74,
	MT:    75,
	ME:    76,
	MW:    77,
	MC:    78,
	IDD:   79,
	IDU:   80,
	ID:    81,
	TRUE:  82,
	FALSE: 83,
	CHR:   84,
	STR:   85,
	NUM:   86,
}

var ruleTable = [88][2]int{
	{1, P},
	{2, LDE},
	{1, LDE},
	{1, DE},
	{1, DE},
	{1, T},
	{1, T},
	{1, T},
	{1, T},
	{1, T},
	{9, DT},
	{8, DT},
	{4, DT},
	{5, DC},
	{3, DC},
	{10, DF},
	{5, LP},
	{3, LP},
	{4, B},
	{2, LDV},
	{1, LDV},
	{2, LS},
	{1, LS},
	{5, DV},
	{3, LI},
	{1, LI},
	{1, S},
	{1, S},
	{5, U},
	{7, U},
	{7, M},
	{5, M},
	{7, M},
	{2, M},
	{4, M},
	{2, M},
	{2, M},
	{3, E},
	{3, E},
	{1, E},
	{3, L},
	{3, L},
	{3, L},
	{3, L},
	{3, L},
	{3, L},
	{1, L},
	{3, R},
	{3, R},
	{1, R},
	{3, Y},
	{3, Y},
	{1, Y},
	{1, F},
	{2, F},
	{2, F},
	{2, F},
	{2, F},
	{3, F},
	{5, F},
	{2, F},
	{2, F},
	{1, F},
	{1, F},
	{1, F},
	{1, F},
	{1, F},
	{3, LE},
	{1, LE},
	{3, LV},
	{4, LV},
	{1, LV},
	{0, NB},
	{0, NF},
	{0, MF},
	{0, MT},
	{0, ME},
	{0, MW},
	{0, MC},
	{1, IDD},
	{1, IDU},
	{1, ID},
	{1, TRUE},
	{1, FALSE},
	{1, CHR},
	{1, STR},
	{1, NUM},
}
