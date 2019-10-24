package semantics

import "github.com/lucbarr/sslang/scope"

type Attribute struct {
	Type int
	Size int

	Attribute Obj
}

func (a Attribute) String() string {
	switch a.Attribute.(type) {
	case ID:
		return "ID"
	case T:
		return "T"
	case E:
		return "E"
	case L:
		return "L"
	case R:
		return "R"
	case Y:
		return "Y"
	case F:
		return "F"
	case LV:
		return "LV"
	case MC:
		return "MC"
	case MT:
		return "MT"
	case ME:
		return "ME"
	case MW:
		return "MW"
	case MA:
		return "MA"
	case LE:
		return "LE"
	case LI:
		return "LI"
	case DC:
		return "DC"
	case LP:
		return "LP"
	case TRUE:
		return "TRUE"
	case FALSE:
		return "FALSE"
	case CHR:
		return "CHR"
	case STR:
		return "STR"
	case NUM:
		return "NUM"
	}
	return ""
}

type Obj interface {
	attributeObj()
}

func (a ID) attributeObj()    {}
func (a T) attributeObj()     {}
func (a E) attributeObj()     {}
func (a L) attributeObj()     {}
func (a R) attributeObj()     {}
func (a Y) attributeObj()     {}
func (a F) attributeObj()     {}
func (a LV) attributeObj()    {}
func (a MC) attributeObj()    {}
func (a MT) attributeObj()    {}
func (a ME) attributeObj()    {}
func (a MW) attributeObj()    {}
func (a MA) attributeObj()    {}
func (a LE) attributeObj()    {}
func (a LI) attributeObj()    {}
func (a DC) attributeObj()    {}
func (a LP) attributeObj()    {}
func (a TRUE) attributeObj()  {}
func (a FALSE) attributeObj() {}
func (a CHR) attributeObj()   {}
func (a STR) attributeObj()   {}
func (a NUM) attributeObj()   {}

type ID struct {
	Object *scope.Object
	Name   int
}

type T struct {
	Type *scope.Object
}

type E struct {
	Type *scope.Object
}

type L struct {
	Type *scope.Object
}

type R struct {
	Type *scope.Object
}

type Y struct {
	Type *scope.Object
}

type F struct {
	Type *scope.Object
}

type LV struct {
	Type *scope.Object
}

type MC struct {
	Type  *scope.Object
	Param *scope.Object
	Err   int
}

type MT struct {
	Label int
}

type ME struct {
	Label int
}

type MW struct {
	Label int
}

type MA struct {
	Label int
}

type LE struct {
	Type  *scope.Object
	Param *scope.Object
	Err   int
	N     int
}

type LI struct {
	List *scope.Object
}

type DC struct {
	List *scope.Object
}

type LP struct {
	List *scope.Object
}

type TRUE struct {
	Type *scope.Object
	Val  int
}

type FALSE struct {
	Type *scope.Object
	Val  int
}

type CHR struct {
	Type *scope.Object
	Pos  int
	Val  rune
}

type STR struct {
	Type *scope.Object
	Pos  int
	Val  string
}

type NUM struct {
	Type *scope.Object
	Pos  int
	Val  int
}

func init() {
	IDDStatic.Attribute = ID{}
	IDUStatic.Attribute = ID{}
	IDStatic.Attribute = ID{}
	TStatic.Attribute = T{}
	LIStatic.Attribute = LI{}
	LI0Static.Attribute = LI{}
	LI1Static.Attribute = LI{}
	TRUStatic.Attribute = TRUE{}
	FALSStatic.Attribute = FALSE{}
	STRStatic.Attribute = STR{}
	CHRStatic.Attribute = CHR{}
	NUMStatic.Attribute = NUM{}
	DCStatic.Attribute = DC{}
	DC0Static.Attribute = DC{}
	DC1Static.Attribute = DC{}
	LPStatic.Attribute = LP{}
	LP0Static.Attribute = LP{}
	LP1Static.Attribute = LP{}
	EStatic.Attribute = E{}
	E0Static.Attribute = E{}
	E1Static.Attribute = E{}
	LStatic.Attribute = L{}
	L0Static.Attribute = L{}
	L1Static.Attribute = L{}
	RStatic.Attribute = R{}
	R0Static.Attribute = R{}
	R1Static.Attribute = R{}
	YStatic.Attribute = Y{}
	Y0Static.Attribute = Y{}
	Y1Static.Attribute = Y{}
	FStatic.Attribute = F{}
	F0Static.Attribute = F{}
	F1Static.Attribute = F{}
	LVStatic.Attribute = LV{}
	LV0Static.Attribute = LV{}
	LV1Static.Attribute = LV{}
	MCStatic.Attribute = MC{}
	LEStatic.Attribute = LE{}
	LE0Static.Attribute = LE{}
	LE1Static.Attribute = LE{}
	MTStatic.Attribute = MT{}
	MEStatic.Attribute = ME{}
	MWStatic.Attribute = MW{}
}

var IDDStatic, IDUStatic, IDStatic, TStatic, LIStatic, LI0Static, LI1Static, TRUStatic, FALSStatic, STRStatic, CHRStatic, NUMStatic, DCStatic, DC0Static, DC1Static, LPStatic, LP0Static, LP1Static, EStatic, E0Static, E1Static, LStatic, L0Static, L1Static, RStatic, R0Static, R1Static, YStatic, Y0Static, Y1Static, FStatic, F0Static, F1Static, LVStatic, LV0Static, LV1Static, MCStatic, LEStatic, LE0Static, LE1Static, MTStatic, MEStatic, MWStatic, NBStatic Attribute

var p, t, f *scope.Object
var name, n, l, l1, l2 int
var functionVarPos int
var curFunction *scope.Object
