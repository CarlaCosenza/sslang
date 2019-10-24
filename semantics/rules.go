package semantics

// Rule enumerates all rules
type Rule int

// All rules enumeration
const (
	PLINE0 Rule = iota
	P0
	LDE0
	LDE1
	DE0
	DE1
	T0
	T1
	T2
	T3
	T4
	DT0
	DT1
	DT2
	DC0
	DC1
	DF0
	LP0
	LP1
	B0
	LDV0
	LDV1
	LS0
	LS1
	DV0
	LI0
	LI1
	S0
	S1
	U0
	U1
	M0
	M1
	M2
	M3
	M4
	M5
	M6
	E0
	E1
	E2
	L0
	L1
	L2
	L3
	L4
	L5
	L6
	R0
	R1
	R2
	Y0
	Y1
	Y2
	F0
	F1
	F2
	F3
	F4
	F5
	F6
	F7
	F8
	F9
	F10
	F11
	F12
	F13
	LE0
	LE1
	LV0
	LV1
	LV2
	IDD0
	IDU0
	ID0
	TRUE0
	FALSE0
	CHR0
	STR0
	NUM0
	NB0
	MF0
	MC0
	NF0
	MT0
	ME0
	MW0
)

func (r Rule) String() string {
	switch r {
	case PLINE0:
		return "PLINE0"
	case P0:
		return "P0"
	case LDE0:
		return "LDE0"
	case LDE1:
		return "LDE1"
	case DE0:
		return "DE0"
	case DE1:
		return "DE1"
	case T0:
		return "T0"
	case T1:
		return "T1"
	case T2:
		return "T2"
	case T3:
		return "T3"
	case T4:
		return "T4"
	case DT0:
		return "DT0"
	case DT1:
		return "DT1"
	case DT2:
		return "DT2"
	case DC0:
		return "DC0"
	case DC1:
		return "DC1"
	case DF0:
		return "DF0"
	case LP0:
		return "LP0"
	case LP1:
		return "LP1"
	case B0:
		return "B0"
	case LDV0:
		return "LDV0"
	case LDV1:
		return "LDV1"
	case LS0:
		return "LS0"
	case LS1:
		return "LS1"
	case DV0:
		return "DV0"
	case LI0:
		return "LI0"
	case LI1:
		return "LI1"
	case S0:
		return "S0"
	case S1:
		return "S1"
	case U0:
		return "U0"
	case U1:
		return "U1"
	case M0:
		return "M0"
	case M1:
		return "M1"
	case M2:
		return "M2"
	case M3:
		return "M3"
	case M4:
		return "M4"
	case M5:
		return "M5"
	case M6:
		return "M6"
	case E0:
		return "E0"
	case E1:
		return "E1"
	case E2:
		return "E2"
	case L0:
		return "L0"
	case L1:
		return "L1"
	case L2:
		return "L2"
	case L3:
		return "L3"
	case L4:
		return "L4"
	case L5:
		return "L5"
	case L6:
		return "L6"
	case R0:
		return "R0"
	case R1:
		return "R1"
	case R2:
		return "R2"
	case Y0:
		return "Y0"
	case Y1:
		return "Y1"
	case Y2:
		return "Y2"
	case F0:
		return "F0"
	case F1:
		return "F1"
	case F2:
		return "F2"
	case F3:
		return "F3"
	case F4:
		return "F4"
	case F5:
		return "F5"
	case F6:
		return "F6"
	case F7:
		return "F7"
	case F8:
		return "F8"
	case F9:
		return "F9"
	case F10:
		return "F10"
	case F11:
		return "F11"
	case F12:
		return "F12"
	case F13:
		return "F13"
	case LE0:
		return "LE0"
	case LE1:
		return "LE1"
	case LV0:
		return "LV0"
	case LV1:
		return "LV1"
	case LV2:
		return "LV2"
	case IDD0:
		return "IDD0"
	case IDU0:
		return "IDU0"
	case ID0:
		return "ID0"
	case TRUE0:
		return "TRUE0"
	case FALSE0:
		return "FALSE0"
	case CHR0:
		return "CHR0"
	case STR0:
		return "STR0"
	case NUM0:
		return "NUM0"
	case NB0:
		return "NB0"
	case MF0:
		return "MF0"
	case MC0:
		return "MC0"
	case NF0:
		return "NF0"
	case MT0:
		return "MT0"
	case ME0:
		return "ME0"
	case MW0:
		return "MW0"
	}
	return ""
}
