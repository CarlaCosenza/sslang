package semantics

import (
	"fmt"
	"os"

	"github.com/lucbarr/sslang/scope"
	"github.com/lucbarr/sslang/syntatical"
)

type Analyser struct {
	Stack []Attribute
	scope *scope.Analyser
	f     *os.File
}

func NewAnalyser() *Analyser {
	f, _ := os.Open("out")
	return &Analyser{
		scope: &scope.Analyser{},
		Stack: []Attribute{},
		f:     f,
	}
}

func (a *Analyser) Parse(r int) {
	rule := Rule(r)

	switch rule {
	case P0:
		break
	case LDE0:
		break
	case LDE1:
		break
	case DE0:
		break
	case DE1:
		break
	case T0: // T-> 'integer'
		t := TStatic.Attribute.(T)
		t.Type = scope.PIntObj

		TStatic.Size = 1
		TStatic.Type = syntatical.T
		TStatic.Attribute = t

		a.Push(TStatic)
	case T1: // T -> 'char'
		t := TStatic.Attribute.(T)
		t.Type = scope.PCharObj

		TStatic.Size = 1
		TStatic.Type = syntatical.T
		TStatic.Attribute = t

		a.Push(TStatic)
		break
	case T2: // T -> 'boolean'
		t := TStatic.Attribute.(T)
		t.Type = scope.PBoolObj

		TStatic.Size = 1
		TStatic.Type = syntatical.T
		TStatic.Attribute = t

		a.Push(TStatic)
		break
	case T3: // T -> 'string'
		t := TStatic.Attribute.(T)
		t.Type = scope.PStringObj

		TStatic.Size = 1
		TStatic.Type = syntatical.T
		TStatic.Attribute = t

		a.Push(TStatic)
		break
	case T4: // T -> IDU
		IDUStatic = a.Top()
		id := IDUStatic.Attribute.(ID)
		p = id.Object
		a.Pop()

		if p.Kind.IsType() || p.Kind == scope.KindUniversal {
			t := TStatic.Attribute.(T)
			t.Type = p
			TStatic.Attribute = t

			if p.Kind == scope.KindAliasType {
				alias := p.T.(scope.Alias)
				TStatic.Size = alias.Size
			} else if p.Kind == scope.KindArrayType {
				arr := p.T.(scope.Alias)
				TStatic.Size = arr.Size
			} else if p.Kind == scope.KindStructType {
				strct := p.T.(scope.Alias)
				TStatic.Size = strct.Size
			}
		} else {
			t := TStatic.Attribute.(T)
			t.Type = scope.PUniversalObj

			TStatic.Attribute = t
			TStatic.Size = 0
		}
		TStatic.Type = syntatical.T
		a.Push(TStatic)
		break
	case DT0: // DT -> 'type' IDD '=' 'array' '[' NUM ']' 'of' T
		TStatic = a.Top()
		a.Pop()
		NUMStatic = a.Top()
		a.Pop()
		IDDStatic = a.Top()
		a.Pop()

		id := IDDStatic.Attribute.(ID)
		num := IDDStatic.Attribute.(NUM)
		typ := IDDStatic.Attribute.(T)

		p = id.Object
		n = num.Val
		t = typ.Type

		p.Kind = scope.KindArrayType
		p.T = scope.Array{
			NumElements: n,
			ElemType:    t,
			Size:        n * TStatic.Size,
		}
		break
	case DT1: // DT -> 'type' IDD '=' 'struct' NB '{' DC '}'
		DCStatic = a.Top()
		a.Pop()
		IDDStatic = a.Top()
		a.Pop()

		obj := IDDStatic.Attribute.(ID)
		dc := DCStatic.Attribute.(DC)
		p = obj.Object

		p.Kind = scope.KindStructType
		p.T = scope.Struct{
			Fields: dc.List,
			Size:   DCStatic.Size,
		}
		a.scope.EndBlock()
		break
	case DT2: // DT -> 'type' IDD '=' T
		TStatic = a.Top()
		a.Pop()
		IDDStatic = a.Top()
		a.Pop()

		id := IDDStatic.Attribute.(ID)
		typ := TStatic.Attribute.(T)

		p = id.Object
		t = typ.Type

		p.Kind = scope.KindAliasType
		p.T = scope.Alias{
			BaseType: t,
			Size:     TStatic.Size,
		}
		break
	case DC0: // DC -> DC ';' LI ':' T
		TStatic = a.Top()
		a.Pop()
		LIStatic = a.Top()
		a.Pop()
		DC1Static = a.Top()

		li := LIStatic.Attribute.(LI)
		typ := TStatic.Attribute.(T)

		p = li.List
		t = typ.Type
		n = DC1Static.Size

		for {
			if p == nil || p.Kind != scope.KindUndefined {
				break
			}

			p.Kind = scope.KindField

			field := p.T.(scope.Field)
			field.PType = t
			field.Index = n
			field.Size = TStatic.Size
			p.T = field

			n = n + TStatic.Size
			p = p.Next
		}

		dc0 := DC0Static.Attribute.(DC)
		dc1 := DC1Static.Attribute.(DC)
		dc0.List = dc1.List
		DC0Static.Size = n
		DC0Static.Type = syntatical.DC

		break
	case DC1: // DC -> LI ':' T
		TStatic = a.Top()
		a.Pop()
		LIStatic = a.Top()
		a.Pop()

		li := LIStatic.Attribute.(LI)
		p = li.List
		ts := TStatic.Attribute.(T)
		t = ts.Type
		n = 0

		for {
			if p == nil || p.Kind != scope.KindUndefined {
				break
			}

			p.Kind = scope.KindField
			field := p.T.(scope.Field)
			field.PType = t
			field.Size = TStatic.Size
			field.Index = n
			p.T = field

			n = n + TStatic.Size
			p = p.Next
		}

		dc := DCStatic.Attribute.(DC)
		li = LIStatic.Attribute.(LI)
		dc.List = li.List
		DCStatic.Attribute = dc

		a.Push(DCStatic)
		break
	case DF0: // DF -> 'function' IDD NF '(' LP ')' ':' T MF B
		a.scope.EndBlock()
		break
	case LP0: // LP -> LP ',' IDD ':' T
		TStatic = a.Top()
		a.Pop()
		IDDStatic = a.Top()
		a.Pop()
		LP1Static = a.Top()
		a.Pop()

		id := IDDStatic.Attribute.(ID)
		typ := TStatic.Attribute.(T)

		p = id.Object
		t = typ.Type
		n = LP1Static.Size

		p.Kind = scope.KindParam
		p.T = scope.Param{
			PType: t,
			Index: n,
			Size:  TStatic.Size,
		}

		lp1 := LP1Static.Attribute.(LP)
		lp0 := LP0Static.Attribute.(LP)
		lp0.List = lp1.List
		LP0Static.Attribute = lp0
		LP0Static.Size = n + TStatic.Size
		LP0Static.Type = syntatical.LP
		a.Push(LP0Static)
		break
	case LP1: // LP -> IDD ':' T
		TStatic = a.Top()
		a.Pop()
		IDDStatic = a.Top()
		a.Pop()

		id := IDDStatic.Attribute.(ID)
		typ := TStatic.Attribute.(T)

		p = id.Object
		t = typ.Type

		p.Kind = scope.KindParam
		p.T = scope.Param{
			PType: t,
			Index: 0,
			Size:  TStatic.Size,
		}

		lp := LPStatic.Attribute.(LP)
		lp.List = p
		LPStatic.Attribute = lp
		LPStatic.Size = TStatic.Size
		LPStatic.Type = syntatical.LP

		a.Push(LPStatic)
		break
	case B0: // B -> '{' LDV LS '}'
		pos, _ := a.f.Seek(0, os.SEEK_CUR)
		a.f.Seek(int64(functionVarPos), os.SEEK_SET)
		funct := f.T.(scope.Function)
		a.f.WriteString(fmt.Sprintf("%02d", funct.Vars))
		a.f.Seek(pos, os.SEEK_SET)
		break

	case LDV0: // LDV -> LDV DV
	case LDV1: // LDV -> DV
	case LS0: // LS -> LS S
	case LS1: // LS -> S
		break
	case DV0: // DV -> 'var' LI ':' T ';'
		TStatic = a.Top()
		typ := TStatic.Attribute.(T)
		t = typ.Type
		a.Pop()

		LIStatic = a.Top()
		a.Pop()

		li := LIStatic.Attribute.(LI)
		p = li.List
		funct := curFunction.T.(scope.Function)
		n = funct.Vars

		for {
			if p == nil || p.Kind != scope.KindUndefined {
				break
			}

			p.Kind = scope.KindVar
			v := p.T.(scope.Var)
			v.PType = t
			v.Size = TStatic.Size
			v.Index = n
			p.T = v

			n = n + TStatic.Size
			p = p.Next
		}

		funct = curFunction.T.(scope.Function)
		funct.Vars = n
		curFunction.T = funct
		break
	case LI0: // LI -> LI ',' IDD
		IDDStatic = a.Top()
		a.Pop()
		LI1Static = a.Top()
		a.Pop()

		li0 := LI0Static.Attribute.(LI)
		li1 := LI0Static.Attribute.(LI)
		li0.List = li1.List
		LI0Static.Type = syntatical.LI
		LI0Static.Attribute = li0
		a.Push(LI0Static)
		break

	case LI1: // LI -> IDD
	case S0: // S -> M
	case S1: // S -> U
	case U0: // U -> 'if' '(' E ')' MT S
	case U1: // U -> 'if' '(' E ')' MT M 'else' ME U
	case M0: // M -> 'if' '(' E ')' MT M 'else' ME M
	case M1: // M -> 'while' MW '(' E ')' MT M
	case M2: // M -> 'do' MW M 'while' '(' E ')' ';'
	case M3: // M -> NB B
	case M4: // M -> LV '=' E ';'
	case M5: // M -> 'break' ';'
	case M6: // M -> 'continue' ';'
	case E0: // E -> E '&&' L
	case E1: // E -> E '||' L
	case E2: // E -> L
	case L0: // L -> L '<' R
	case L1: // L -> L '>' R
	case L2: // L -> L '<=' R
	case L3: // L -> L '>=' R
	case L4: // L -> L '==' R
	case L5: // L -> L '!=' R
	case L6: // L -> R
	case R0: // R -> R '+' Y
	case R1: // R -> R '-' Y
	case R2: // R -> Y
	case Y0: // Y -> Y '*' F
	case Y1: // Y -> Y '/' F
	case Y2: // Y -> F
	case F0: // F -> LV
	case F1: // F -> '++' LV
	case F2: // F -> '--' LV
	case F3: // F -> LV '++'
	case F4: // F -> LV '--'
	case F5: // F -> '(' E ')'
	case F6: // F -> IDU MC '(' LE ')'
	case F7: // F -> '-' F
	case F8: // F -> '!' F
	case F9: // F -> TRUE
	case F10: // F -> FALSE
	case F11: // F -> CHR
	case F12: // F -> STR
	case F13: // F -> NUM
	case LE0: // LE -> LE ',' E
	case LE1: // LE -> E
	case LV0: // LV -> LV '.' IDU
	case LV1: // LV -> LV '[' E ']'
	case LV2: // LV -> IDU
	case IDD0: // IDD -> Id
	case IDU0: // IDU -> Id
	case ID0: // ID -> Id
	case TRUE0: // TRUE ->  'true'
	case FALSE0: // FALSE -> 'false'
	case CHR0: // CHR -> c
	case STR0: // STR -> s
	case NUM0: // NUM -> n
	case NB0: // NB -> ''
	case MF0: // MF -> ''
	case MC0: // MC -> ''
	case NF0: // NF -> ''
	case MT0: // MT -> ''
	case ME0: // ME -> ''
	case MW0: // MW -> ''
	}
}

func (a *Analyser) Push(attr Attribute) {
	a.Stack = append(a.Stack, attr)
}

func (a *Analyser) Pop() {
	a.Stack = a.Stack[:len(a.Stack)-1]
}

func (a *Analyser) Top() Attribute {
	if len(a.Stack) == 0 {
		return Attribute{}
	}
	return a.Stack[len(a.Stack)-1]
}
