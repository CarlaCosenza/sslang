package semantics

import "github.com/lucbarr/sslang/scope"

type Attribute struct {
	Type int
	Size int

	Attribute Obj
}

type Obj interface {
	attributeObj()
}

func (a ID) attributeObj() {}

type ID struct {
	Object *scope.Object
	Name   int
}

func (a T) attributeObj() {}

type T struct {
	Type *scope.Object
}

func (a E) attributeObj() {}

type E struct {
	Type *scope.Object
}

func (a L) attributeObj() {}

type L struct {
	Type *scope.Object
}

func (a R) attributeObj() {}

type R struct {
	Type *scope.Object
}

func (a Y) attributeObj() {}

type Y struct {
	Type *scope.Object
}

func (a F) attributeObj() {}

type F struct {
	Type *scope.Object
}

func (a LV) attributeObj() {}

type LV struct {
	Type *scope.Object
}

func (a MC) attributeObj() {}

type MC struct {
	Type  *scope.Object
	Param *scope.Object
	Err   int
}

func (a MT) attributeObj() {}

type MT struct {
	Label int
}

func (a ME) attributeObj() {}

type ME struct {
	Label int
}

func (a MW) attributeObj() {}

type MW struct {
	Label int
}

func (a MA) attributeObj() {}

type MA struct {
	Label int
}

func (a LE) attributeObj() {}

type LE struct {
	Type  *scope.Object
	Param *scope.Object
	Err   int
	N     int
}

func (a LI) attributeObj() {}

type LI struct {
	List *scope.Object
}

func (a DC) attributeObj() {}

type DC struct {
	List *scope.Object
}

func (a LP) attributeObj() {}

type LP struct {
	List *scope.Object
}

func (a TRUE) attributeObj() {}

type TRUE struct {
	Type *scope.Object
	Val  int
}

func (a FALSE) attributeObj() {}

type FALSE struct {
	Type *scope.Object
	Val  int
}

func (a CHR) attributeObj() {}

type CHR struct {
	Type *scope.Object
	Pos  int
	Val  rune
}

func (a STR) attributeObj() {}

type STR struct {
	Type *scope.Object
	Pos  int
	Val  string
}

func (a NUM) attributeObj() {}

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
