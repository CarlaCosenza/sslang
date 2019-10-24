package semantics

import (
	"fmt"
	"os"
	"strings"

	"github.com/lucbarr/sslang/lexical"
	"github.com/lucbarr/sslang/nonterminals"
	"github.com/lucbarr/sslang/scope"
)

type Analyser struct {
	Stack []Attribute
	scope *scope.Analyser
	lexer *lexical.Lexer
	f     *os.File

	nFuncs int
	label  int
}

func (a *Analyser) Close() error {
	return a.f.Close()
}

func NewAnalyser(l *lexical.Lexer) *Analyser {
	f, _ := os.Create("out")
	return &Analyser{
		scope: &scope.Analyser{},
		Stack: []Attribute{},
		f:     f,
		lexer: l,
	}
}

func (a *Analyser) StackString() string {
	var sb strings.Builder
	sb.WriteString("[ ")
	for _, s := range a.Stack {
		sb.WriteString(fmt.Sprintf("%v ", s))
	}
	sb.WriteString("]")
	return sb.String()
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
		TStatic.Type = nonterminals.T
		TStatic.Attribute = t

		a.Push(TStatic)
		break
	case T1: // T -> 'char'
		t := TStatic.Attribute.(T)
		t.Type = scope.PCharObj

		TStatic.Size = 1
		TStatic.Type = nonterminals.T
		TStatic.Attribute = t

		a.Push(TStatic)
		break
	case T2: // T -> 'boolean'
		t := TStatic.Attribute.(T)
		t.Type = scope.PBoolObj

		TStatic.Size = 1
		TStatic.Type = nonterminals.T
		TStatic.Attribute = t

		a.Push(TStatic)
		break
	case T3: // T -> 'string'
		t := TStatic.Attribute.(T)
		t.Type = scope.PStringObj

		TStatic.Size = 1
		TStatic.Type = nonterminals.T
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
				arr := p.T.(scope.Array)
				TStatic.Size = arr.Size
			} else if p.Kind == scope.KindStructType {
				strct := p.T.(scope.Struct)
				TStatic.Size = strct.Size
			}
		} else {
			t := TStatic.Attribute.(T)
			t.Type = scope.PUniversalObj

			TStatic.Attribute = t
			TStatic.Size = 0
		}
		TStatic.Type = nonterminals.T
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
		num := NUMStatic.Attribute.(NUM)
		typ := TStatic.Attribute.(T)

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
		a.Pop()

		li := LIStatic.Attribute.(LI)
		typ := TStatic.Attribute.(T)

		p = li.List
		t = typ.Type
		n = DC1Static.Size

		for p != nil && p.Kind == scope.KindUndefined {

			p.Kind = scope.KindField

			field, ok := p.T.(scope.Field)
			if !ok {
				field = scope.Field{}
			}
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
		DC0Static.Type = nonterminals.DC

		a.Push(DC0Static)
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
			field, ok := p.T.(scope.Field)
			if !ok {
				field = scope.Field{}
			}
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
		LP0Static.Type = nonterminals.LP
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
		param, ok := p.T.(scope.Param)
		if !ok {
			param = scope.Param{}
		}
		param.PType = t
		param.Index = 0
		param.Size = TStatic.Size
		p.T = param

		lp := LPStatic.Attribute.(LP)
		lp.List = p
		LPStatic.Attribute = lp
		LPStatic.Size = TStatic.Size
		LPStatic.Type = nonterminals.LP

		a.Push(LPStatic)
		break
	case B0: // B -> '{' LDV LS '}'

		a.f.WriteString(fmt.Sprintf("END_FUNC\n"))
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

		for p != nil && p.Kind == scope.KindUndefined {

			p.Kind = scope.KindVar
			v, ok := p.T.(scope.Var)
			if !ok {
				v = scope.Var{}
			}
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
		LI0Static.Type = nonterminals.LI
		LI0Static.Attribute = li0
		a.Push(LI0Static)
		break
	case LI1: // LI -> IDD
		IDDStatic = a.Top()
		a.Pop()

		li := LIStatic.Attribute.(LI)
		li.List = IDDStatic.Attribute.(ID).Object
		LIStatic.Attribute = li
		LIStatic.Type = nonterminals.LI
		a.Push(LIStatic)
		break
	case S0: // S -> M
		break
	case S1: // S -> U
		break
	case U0: // U -> 'if' '(' E ')' MT S
		MTStatic = a.Top()
		a.Pop()
		EStatic = a.Top()
		a.Pop()

		e := EStatic.Attribute.(E)
		t = e.Type

		mt := MTStatic.Attribute.(MT)
		a.f.WriteString(fmt.Sprintf("L%d\n", mt.Label))
		break
	case U1: // U -> 'if' '(' E ')' MT M 'else' ME U
		MEStatic = a.Top()
		a.Pop()
		MTStatic = a.Top()
		a.Pop()
		EStatic = a.Top()
		a.Pop()

		me := MEStatic.Attribute.(ME)
		e := MEStatic.Attribute.(E)

		l = me.Label
		t = e.Type

		a.f.WriteString(fmt.Sprintf("L%d", l))
	case M0: // M -> 'if' '(' E ')' MT M 'else' ME M
		MEStatic = a.Top()
		a.Pop()
		MTStatic = a.Top()
		a.Pop()
		EStatic = a.Top()
		a.Pop()

		me := MEStatic.Attribute.(ME)
		e := MEStatic.Attribute.(E)

		l = me.Label
		t = e.Type

		a.f.WriteString(fmt.Sprintf("L%d", l))
		break
	case M1: // M -> 'while' MW '(' E ')' MT M
		MTStatic = a.Top()
		a.Pop()
		EStatic = a.Top()
		a.Pop()
		MWStatic = a.Top()
		a.Pop()

		mw := MWStatic.Attribute.(MW)
		mt := MTStatic.Attribute.(MT)
		es := EStatic.Attribute.(E)

		l1 = mw.Label
		l2 = mt.Label

		t = es.Type

		a.f.WriteString(fmt.Sprintf("\tTJMP_BW L%d\nL%d\n", l1, l2))
		break
	case M2: // M -> 'do' MW M 'while' '(' E ')' ';'
		EStatic = a.Top()
		a.Pop()
		MWStatic = a.Top()
		a.Pop()

		mw := MWStatic.Attribute.(MW)
		es := EStatic.Attribute.(E)

		l = mw.Label
		t = es.Type

		a.f.WriteString(fmt.Sprintf("\tNOT\n\tTJMP_BW L%d\n", l))
		break
	case M3: // M -> NB B
		a.scope.EndBlock()
		break
	case M4: // M -> LV '=' E ';'
		EStatic = a.Top()
		a.Pop()
		LVStatic = a.Top()
		a.Pop()

		lv := LVStatic.Attribute.(LV)

		t = lv.Type
		typ := lv.Type.T.(scope.Type)

		a.f.WriteString(fmt.Sprintf("\tSTORE_REF %d\n", typ.Size))
		break

	case M5: // M -> 'break' ';'
	case M6: // M -> 'continue' ';'
		break
	case E0: // E -> E '&&' L
		LStatic = a.Top()
		a.Pop()
		E1Static = a.Top()
		a.Pop()

		e := E0Static.Attribute.(E)
		e.Type = scope.PBoolObj
		E0Static.Attribute = e

		a.Push(E0Static)

		a.f.WriteString(fmt.Sprintf("\tAND\n"))
		break
	case E1: // E -> E '||' L
		LStatic = a.Top()
		a.Pop()
		E1Static = a.Top()
		a.Pop()

		e := E0Static.Attribute.(E)
		e.Type = scope.PBoolObj
		E0Static.Attribute = e

		a.Push(E0Static)

		a.f.WriteString(fmt.Sprintf("\tOR\n"))
		break
	case E2: // E -> L
		LStatic = a.Top()
		a.Pop()

		e := EStatic.Attribute.(E)
		l := LStatic.Attribute.(L)
		e.Type = l.Type

		EStatic.Attribute = e
		EStatic.Type = nonterminals.E

		a.Push(EStatic)
		break
	case L0: // L -> L '<' R
		RStatic = a.Top()
		a.Pop()
		L1Static = a.Top()
		a.Pop()

		l := L0Static.Attribute.(L)
		l.Type = scope.PBoolObj
		L0Static.Attribute = l
		L0Static.Type = nonterminals.L

		a.Push(L0Static)
		a.f.WriteString(fmt.Sprintf("\tLT\n"))
		break
	case L1: // L -> L '>' R
		RStatic = a.Top()
		a.Pop()
		L1Static = a.Top()
		a.Pop()

		l := L0Static.Attribute.(L)
		l.Type = scope.PBoolObj
		L0Static.Attribute = l
		L0Static.Type = nonterminals.L

		a.Push(L0Static)
		a.f.WriteString(fmt.Sprintf("\tGT\n"))
		break
	case L2: // L -> L '<=' R
		RStatic = a.Top()
		a.Pop()
		L1Static = a.Top()
		a.Pop()

		l := L0Static.Attribute.(L)
		l.Type = scope.PBoolObj
		L0Static.Attribute = l
		L0Static.Type = nonterminals.L

		a.Push(L0Static)
		a.f.WriteString(fmt.Sprintf("\tLE\n"))
		break
	case L3: // L -> L '>=' R
		RStatic = a.Top()
		a.Pop()
		L1Static = a.Top()
		a.Pop()

		l := L0Static.Attribute.(L)
		l.Type = scope.PBoolObj
		L0Static.Attribute = l
		L0Static.Type = nonterminals.L

		a.Push(L0Static)
		a.f.WriteString(fmt.Sprintf("\tGE\n"))
		break
	case L4: // L -> L '==' R
		RStatic = a.Top()
		a.Pop()
		L1Static = a.Top()
		a.Pop()

		l := L0Static.Attribute.(L)
		l.Type = scope.PBoolObj
		L0Static.Attribute = l
		L0Static.Type = nonterminals.L

		a.Push(L0Static)
		a.f.WriteString(fmt.Sprintf("\tEQ\n"))
		break
	case L5: // L -> L '!=' R
		RStatic = a.Top()
		a.Pop()
		L1Static = a.Top()
		a.Pop()

		l := L0Static.Attribute.(L)
		l.Type = scope.PBoolObj
		L0Static.Attribute = l
		L0Static.Type = nonterminals.L

		a.Push(L0Static)
		a.f.WriteString(fmt.Sprintf("\tNE\n"))
		break
	case L6: // L -> R
		RStatic = a.Top()
		a.Pop()

		ls := LStatic.Attribute.(L)
		rs := RStatic.Attribute.(R)

		ls.Type = rs.Type
		LStatic.Attribute = ls
		LStatic.Type = nonterminals.L

		a.Push(LStatic)
		break
	case R0: // R -> R '+' Y
		YStatic = a.Top()
		a.Pop()
		R1Static = a.Top()
		a.Pop()

		r0 := R0Static.Attribute.(R)
		r1 := R1Static.Attribute.(R)
		r0.Type = r1.Type

		R0Static.Attribute = r0
		R0Static.Type = nonterminals.R

		a.Push(R0Static)
		a.f.WriteString(fmt.Sprintf("\tADD\n"))
		break
	case R1: // R -> R '-' Y
		YStatic = a.Top()
		a.Pop()
		R1Static = a.Top()
		a.Pop()

		r0 := R0Static.Attribute.(R)
		r1 := R1Static.Attribute.(R)
		r0.Type = r1.Type

		R0Static.Attribute = r0
		R0Static.Type = nonterminals.R

		a.Push(R0Static)
		a.f.WriteString(fmt.Sprintf("\tSUB\n"))
		break
	case R2: // R -> Y
		YStatic = a.Top()
		a.Pop()
		r := RStatic.Attribute.(R)
		y := YStatic.Attribute.(Y)
		r.Type = y.Type

		RStatic.Attribute = r
		RStatic.Type = nonterminals.R

		a.Push(RStatic)
		break
	case Y0: // Y -> Y '*' F
		FStatic = a.Top()
		a.Pop()
		Y1Static = a.Top()
		a.Pop()

		y0 := Y0Static.Attribute.(Y)
		y1 := Y1Static.Attribute.(Y)
		y0.Type = y1.Type

		Y0Static.Attribute = y0
		Y0Static.Type = nonterminals.Y
		a.Push(Y0Static)

		a.f.WriteString(fmt.Sprintf("\tMUL\n"))
		break
	case Y1: // Y -> Y '/' F
		FStatic = a.Top()
		a.Pop()
		Y1Static = a.Top()
		a.Pop()

		y0 := Y0Static.Attribute.(Y)
		y1 := Y1Static.Attribute.(Y)
		y0.Type = y1.Type

		Y0Static.Attribute = y0
		Y0Static.Type = nonterminals.Y
		a.Push(Y0Static)

		a.f.WriteString(fmt.Sprintf("\tDIV\n"))
		break
	case Y2: // Y -> F
		FStatic = a.Top()
		a.Pop()

		y := YStatic.Attribute.(Y)
		f := FStatic.Attribute.(F)
		y.Type = f.Type
		Y0Static.Attribute = y
		Y0Static.Type = nonterminals.Y

		a.Push(Y0Static)
		break
	case F0: // F -> LV
		LVStatic = a.Top()
		a.Pop()

		lv := LVStatic.Attribute.(LV)
		typ := lv.Type.T.(scope.Type)

		n = typ.Size

		f := FStatic.Attribute.(F)
		f.Type = lv.Type
		FStatic.Attribute = f
		FStatic.Type = nonterminals.F

		a.Push(FStatic)
		a.f.WriteString(fmt.Sprintf("\tDE_REF %d\n", n))
		break
	case F1: // F -> '++' LV
		LVStatic = a.Top()
		a.Pop()

		lv := LVStatic.Attribute.(LV)
		t = lv.Type

		fs := FStatic.Attribute.(F)
		fs.Type = scope.PIntObj
		FStatic.Attribute = fs
		FStatic.Type = nonterminals.F

		a.Push(FStatic)
		a.f.WriteString(fmt.Sprintf("\tDUP\n\tDUP\n\tDE_REF 1\n"))
		a.f.WriteString(fmt.Sprintf("\tINC\n\tSTORE_REF 1\n\tDE_REF 1\n"))
		break
	case F2: // F -> '--' LV
		LVStatic = a.Top()
		a.Pop()

		lv := LVStatic.Attribute.(LV)
		t = lv.Type

		fs := FStatic.Attribute.(F)
		fs.Type = lv.Type
		FStatic.Attribute = fs
		FStatic.Type = nonterminals.F

		a.Push(FStatic)
		a.f.WriteString(fmt.Sprintf("\tDUP\n\tDUP\n\tDE_REF 1\n"))
		a.f.WriteString(fmt.Sprintf("\tINC\n\tSTORE_REF 1\n\tDE_REF 1\n"))
		break
	case F3: // F -> LV '++'
		LVStatic = a.Top()
		a.Pop()

		lv := LVStatic.Attribute.(LV)
		t = lv.Type

		fs := FStatic.Attribute.(F)
		fs.Type = lv.Type
		FStatic.Attribute = fs
		FStatic.Type = nonterminals.F

		a.Push(FStatic)
		a.f.WriteString(fmt.Sprintf("\tDUP\n\tDUP\n\tDE_REF 1\n"))
		a.f.WriteString(fmt.Sprintf("\tINC\n\tSTORE_REF 1\n\tDE_REF 1\n"))
		a.f.WriteString(fmt.Sprintf("\tDEC\n"))
		break
	case F4: // F -> LV '--'
		LVStatic = a.Top()
		a.Pop()

		lv := LVStatic.Attribute.(LV)
		t = lv.Type

		fs := FStatic.Attribute.(F)
		fs.Type = t
		FStatic.Attribute = fs
		FStatic.Type = nonterminals.F

		a.Push(FStatic)
		a.f.WriteString(fmt.Sprintf("\tDUP\n\tDUP\n\tDE_REF 1\n"))
		a.f.WriteString(fmt.Sprintf("\tDEC\n\tSTORE_REF 1\n\tDE_REF 1\n"))
		a.f.WriteString(fmt.Sprintf("\tINC\n"))
		break
	case F5: // F -> '(' E ')'
		EStatic = a.Top()
		a.Pop()

		fs := FStatic.Attribute.(F)
		es := EStatic.Attribute.(E)
		fs.Type = es.Type
		FStatic.Attribute = fs

		a.Push(FStatic)
		break
	case F6: // F -> IDU MC '(' LE ')'
		LEStatic = a.Top()
		a.Pop()
		MCStatic = a.Top()
		a.Pop()
		IDUStatic = a.Top()
		a.Pop()

		id := IDUStatic.Attribute.(ID)
		f = id.Object

		fs := FStatic.Attribute.(F)
		ms := MCStatic.Attribute.(MC)
		fs.Type = ms.Type
		FStatic.Attribute = fs
		FStatic.Type = nonterminals.F

		a.Push(FStatic)

		funct := f.T.(scope.Function)
		a.f.WriteString(fmt.Sprintf("\tCALL %d\n", funct.Index))
		break
	case F7: // F -> '-' F
		F1Static = a.Top()
		a.Pop()

		f1 := F1Static.Attribute.(F)
		t = f1.Type

		f0 := F0Static.Attribute.(F)
		f0.Type = t
		F0Static.Attribute = f0
		F0Static.Type = nonterminals.F

		a.Push(F0Static)
		a.f.WriteString(fmt.Sprintf("\tNEG\n"))
		break
	case F8: // F -> '!' F
		F1Static = a.Top()
		a.Pop()

		f1 := F1Static.Attribute.(F)
		t = f1.Type

		f0 := F0Static.Attribute.(F)
		f0.Type = t
		F0Static.Attribute = f0
		F0Static.Type = nonterminals.F

		a.Push(F0Static)
		a.f.WriteString(fmt.Sprintf("\tNOT\n"))
		break
	case F9: // F -> TRUE
		TRUStatic = a.Top()
		a.Pop()

		fs := FStatic.Attribute.(F)
		fs.Type = scope.PBoolObj
		FStatic.Attribute = fs
		FStatic.Type = nonterminals.F

		a.Push(FStatic)

		a.f.WriteString(fmt.Sprintf("\tLOAD_CONST %d\n", a.lexer.SecondaryToken))
		break
	case F10: // F -> FALSE
		FALSStatic = a.Top()
		a.Pop()

		fs := FStatic.Attribute.(F)
		fs.Type = scope.PBoolObj

		FStatic.Attribute = fs
		FStatic.Type = nonterminals.F

		a.Push(FStatic)

		a.f.WriteString(fmt.Sprintf("\tLOAD_CONST %d\n", a.lexer.SecondaryToken))
		break
	case F11: // F -> CHR
		CHRStatic = a.Top()
		a.Pop()

		fs := FStatic.Attribute.(F)
		fs.Type = scope.PCharObj

		FStatic.Attribute = fs
		FStatic.Type = nonterminals.F

		a.Push(FStatic)

		n = a.lexer.SecondaryToken
		a.f.WriteString(fmt.Sprintf("\tLOAD_CONST %d\n", n))
		break
	case F12: // F -> STR
		STRStatic = a.Top()
		a.Pop()

		fs := FStatic.Attribute.(F)
		fs.Type = scope.PStringObj

		FStatic.Attribute = fs
		FStatic.Type = nonterminals.F

		a.Push(FStatic)

		n = a.lexer.SecondaryToken
		a.f.WriteString(fmt.Sprintf("\tLOAD_CONST %d\n", n))
		break
	case F13: // F -> NUM
		NUMStatic = a.Top()
		a.Pop()

		fs := FStatic.Attribute.(F)
		fs.Type = scope.PIntObj

		FStatic.Attribute = fs
		FStatic.Type = nonterminals.F

		a.Push(FStatic)

		n = a.lexer.SecondaryToken
		a.f.WriteString(fmt.Sprintf("\tLOAD_CONST %d\n", n))
		break
	case LE0: // LE -> LE ',' E
		EStatic = a.Top()
		a.Pop()
		LE1Static = a.Top()
		a.Pop()

		le0 := LE0Static.Attribute.(LE)
		le1 := LE1Static.Attribute.(LE)
		le0.Param = nil
		le0.Err = le1.Err
		LE0Static.Attribute = le0

		n = le1.N
		if le1.Err != 0 {
			p = le1.Param
			if p == nil {
				le0 := LE0Static.Attribute.(LE)
				le0.Err = 1
				LE0Static.Attribute = le0
			} else {
				le0 := LE0Static.Attribute.(LE)
				le0.Param = p.Next
				le0.N = n + 1
				LE0Static.Attribute = le0
			}
		}

		LE0Static.Type = nonterminals.LE
		a.Push(LE0Static)
		break
	case LE1: // LE -> E
		EStatic = a.Top()
		a.Pop()
		MCStatic = a.Top()

		le := LEStatic.Attribute.(LE)
		mc := MCStatic.Attribute.(MC)
		le.Param = nil
		le.Err = mc.Err
		LEStatic.Attribute = le
		n = 1

		if mc.Err != 0 {
			p = mc.Param
			if p == nil {
				le := LEStatic.Attribute.(LE)
				le.Err = 1
				LEStatic.Attribute = le
			} else {
				le := LEStatic.Attribute.(LE)
				le.Param = p.Next
				le.N = n + 1
				LEStatic.Attribute = le
			}
		}

		LEStatic.Type = nonterminals.LE
		a.Push(LEStatic)
	case LV0: // LV -> LV '.' IDU
		IDStatic = a.Top()
		a.Pop()
		LV1Static = a.Top()
		a.Pop()

		lv1 := LV1Static.Attribute.(LV)
		t = lv1.Type

		if t.Kind != scope.KindStructType {
			if t.Kind != scope.KindUniversal {
				// TODO
			}
			lv0 := LV0Static.Attribute.(LV)
			lv0.Type = scope.PUniversalObj
			LV0Static.Attribute = lv0
		} else {
			st := t.T.(scope.Struct)
			p = st.Fields

			for p != nil {
				if ids := IDStatic.Attribute.(ID); p.Name == ids.Name {
					break
				}
				p = p.Next
			}

			if p == nil {
				lv0 := LV0Static.Attribute.(LV)
				lv0.Type = scope.PUniversalObj
				LV0Static.Attribute = lv0
			} else {
				lv0 := LV0Static.Attribute.(LV)
				field := p.T.(scope.Field)
				lv0.Type = field.PType

				typ := lv0.Type.T.(scope.Type)
				typ.Size = field.Size
				lv0.Type.T = typ

				LV0Static.Attribute = lv0

				a.f.WriteString(fmt.Sprintf("\tADD %d", field.Index))
			}
		}

		LV0Static.Type = nonterminals.LV
		a.Push(LV0Static)
		break
	case LV1: // LV -> LV '[' E ']'
		EStatic = a.Top()
		a.Pop()
		LV1Static = a.Top()
		a.Pop()

		t = LV1Static.Attribute.(LV).Type

		lv0 := LV0Static.Attribute.(LV)
		if a.scope.CheckTypes(t, scope.PStringObj) {
			lv0.Type = scope.PCharObj
		} else if t.Kind == scope.KindArrayType {
			if t.Kind == scope.KindUniversal {
				//err ??
			}
			lv0.Type = scope.PUniversalObj
		} else {
			elemType := t.T.(scope.Array).ElemType
			lv0.Type = elemType
			n = elemType.T.(scope.Type).Size

			a.f.WriteString(fmt.Sprintf("\tMUL %d\n\tADD\n", n))
		}

		// if a.scope.CheckTypes(EStatic.Attribute.(E).Type, scope.PIntObj)

		LV0Static.Type = nonterminals.LV
		LV0Static.Attribute = lv0
		break
	case LV2: // LV -> IDU
		IDUStatic = a.Top()
		a.Pop()

		p = IDUStatic.Attribute.(ID).Object
		lv := LVStatic.Attribute.(LV)
		vart, ok := p.T.(scope.Var)
		if !ok {
			vart = scope.Var{
				PType: &scope.Object{},
			}
		}
		if p.Kind != scope.KindVar && p.Kind != scope.KindParam {
			if p.Kind != scope.KindUniversal {
				// err ?
			}
			lv.Type = scope.PUniversalObj
		} else {
			lv.Type = vart.PType
			typ, ok := lv.Type.T.(scope.Type)
			if !ok {
				typ = scope.Type{}
			}
			typ.Size = vart.Size

			lv.Type.T = typ
			a.f.WriteString(fmt.Sprintf("\tLOAD_REF %d\n", vart.Index))
		}
		LVStatic.Attribute = lv
		LVStatic.Type = nonterminals.LV

		t = lv.Type

		a.Push(LVStatic)
		break
	case IDD0: // IDD -> Id
		name = a.lexer.SecondaryToken
		ids := IDDStatic.Attribute.(ID)
		ids.Name = name

		if p = a.scope.SearchLocalSymbol(name); p != nil {
			// err ?
		} else {
			p = a.scope.DefineSymbol(name)
		}

		ids.Object = p
		IDDStatic.Attribute = ids

		a.Push(IDDStatic)
		break
	case IDU0: // IDU -> Id
		name = a.lexer.SecondaryToken
		idu := IDUStatic.Attribute.(ID)
		idu.Name = name

		if p = a.scope.SearchGlobalSymbol(name); p == nil {
			// err ?
			p = a.scope.DefineSymbol(name)
			panic(fmt.Errorf("undeclared variable"))
		}

		idu.Object = p
		IDUStatic.Attribute = idu

		a.Push(IDUStatic)
		break
	case ID0: // ID -> Id
		name = a.lexer.SecondaryToken
		ids := IDStatic.Attribute.(ID)
		ids.Name = name
		ids.Object = nil
		IDDStatic.Attribute = ids
		a.Push(IDDStatic)
		break
	case TRUE0: // TRUE ->  'true'
		tru := TRUStatic.Attribute.(TRUE)
		tru.Type = scope.PBoolObj
		tru.Val = 1

		TRUStatic.Type = nonterminals.TRUE
		TRUStatic.Attribute = tru

		a.Push(TRUStatic)
		break
	case FALSE0: // FALSE -> 'false'
		fals := TRUStatic.Attribute.(TRUE)
		fals.Type = scope.PBoolObj
		fals.Val = 0

		FALSStatic.Type = nonterminals.FALSE
		FALSStatic.Attribute = fals

		a.Push(FALSStatic)
		break
	case CHR0: // CHR -> c
		chr := CHRStatic.Attribute.(CHR)
		chr.Type = scope.PCharObj
		chr.Pos = a.lexer.SecondaryToken
		chr.Val = a.lexer.GetRuneConstant(a.lexer.SecondaryToken)

		CHRStatic.Type = nonterminals.CHR
		CHRStatic.Attribute = chr

		a.Push(CHRStatic)
		break
	case STR0: // STR -> s
		str := STRStatic.Attribute.(STR)
		str.Type = scope.PStringObj
		str.Pos = a.lexer.SecondaryToken
		str.Val = a.lexer.GetStringConstant(a.lexer.SecondaryToken)

		STRStatic.Type = nonterminals.STR
		STRStatic.Attribute = str

		a.Push(STRStatic)
		break
	case NUM0: // NUM -> n
		num := NUMStatic.Attribute.(NUM)
		num.Type = scope.PStringObj
		num.Pos = a.lexer.SecondaryToken
		num.Val = a.lexer.GetNumeralConstant(a.lexer.SecondaryToken)

		NUMStatic.Type = nonterminals.NUM
		NUMStatic.Attribute = num

		a.Push(NUMStatic)
		break
	case NB0: // NB -> ''
		a.scope.NewBlock()
		break
	case MF0: // MF -> ''
		TStatic = a.Top()
		a.Pop()
		LPStatic = a.Top()
		a.Pop()
		IDDStatic = a.Top()
		a.Pop()

		f = IDDStatic.Attribute.(ID).Object
		ts := TStatic.Attribute.(T)
		lp := LPStatic.Attribute.(LP)

		f.Kind = scope.KindFunction
		funct := f.T.(scope.Function)
		funct.PRetType = ts.Type
		funct.PParams = lp.List
		funct.Params = LPStatic.Size
		funct.Vars = 0
		f.T = funct
		curFunction = f

		a.f.WriteString(fmt.Sprintf("BEGIN_FUNC %d, %d, %02d\n", funct.Index, funct.Params, 0))
		pos, _ := a.f.Seek(0, os.SEEK_CUR)
		functionVarPos = int(pos) - 3
		break
	case MC0: // MC -> ''
		IDUStatic = a.Top()
		f = IDUStatic.Attribute.(ID).Object

		mc := MCStatic.Attribute.(MC)
		if f.Kind != scope.KindFunction {
			mc.Type = scope.PUniversalObj
			mc.Param = nil
			mc.Err = 1
			// err ?
		} else {
			mc.Type = f.T.(scope.Function).PRetType
			mc.Param = f.T.(scope.Function).PParams
			mc.Err = 0
		}

		MCStatic.Attribute = mc
		MCStatic.Type = nonterminals.MC

		a.Push(MCStatic)
		break
	case NF0: // NF -> ''
		IDDStatic = a.Top()

		f = IDDStatic.Attribute.(ID).Object

		f.Kind = scope.KindFunction
		funct, ok := f.T.(scope.Function)
		if !ok {
			f.T = scope.Function{}
		}
		funct.Params = 0
		funct.Vars = 0
		funct.Index = a.nFuncs
		a.nFuncs++
		f.T = funct

		a.scope.NewBlock()
		break
	case MT0: // MT -> ''
		l = a.label
		a.label++

		mt := MTStatic.Attribute.(MT)
		mt.Label = l

		MTStatic.Attribute = mt
		MTStatic.Type = nonterminals.MT

		a.f.WriteString(fmt.Sprintf("\tTJMP_FW L%d\n", l))

		a.Push(MTStatic)
		break
	case ME0: // ME -> ''
		MTStatic = a.Top()
		l1 := MTStatic.Attribute.(MT).Label

		l2 := a.label
		a.label++

		me := MEStatic.Attribute.(ME)
		me.Label = l2
		MEStatic.Type = nonterminals.ME

		a.f.WriteString(fmt.Sprintf("\tTJMP_FW %d\n L%d\n", l2, l1))
		a.Push(MEStatic)
		break
	case MW0: // MW -> ''
		l = a.label
		a.label++
		mw := MWStatic.Attribute.(MW)
		mw.Label = l
		MWStatic.Attribute = mw

		a.f.WriteString(fmt.Sprintf("\tL%d", l))
		a.Push(MWStatic)
		break
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
