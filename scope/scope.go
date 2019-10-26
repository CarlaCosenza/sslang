package scope

import (
	"fmt"
	"strings"
)

const (
	maxNestLevel = 64
)

// Kind defines a kind
type Kind int

// Kinds of Kind
const (
	KindVar Kind = iota
	KindParam
	KindFunction
	KindField

	KindArrayType
	KindStructType
	KindAliasType
	KindScalarType

	KindUniversal

	KindUndefined = -1
)

func (k Kind) IsType() bool {
	return k == KindArrayType ||
		k == KindStructType ||
		k == KindAliasType ||
		k == KindScalarType
}

var (
	IntObj  = Object{-1, nil, KindScalarType, nil}
	PIntObj = &IntObj

	CharObj  = Object{-1, nil, KindScalarType, nil}
	PCharObj = &CharObj

	BoolObj  = Object{-1, nil, KindScalarType, nil}
	PBoolObj = &BoolObj

	StringObj  = Object{-1, nil, KindScalarType, nil}
	PStringObj = &StringObj

	UniversalObj  = Object{-1, nil, KindScalarType, nil}
	PUniversalObj = &UniversalObj
)

// Object defines a scope object
type Object struct {
	Name int
	Next *Object
	Kind Kind

	// we use this to mimic the polymorphism the professor uses
	// on his compiler. shrug
	T ObjectType
}

// ObjectType object types have to implement this interface
type ObjectType interface {
	objType()
}

func (a Alias) objType()    {}
func (a Type) objType()     {}
func (a Array) objType()    {}
func (a Struct) objType()   {}
func (a Function) objType() {}
func (a Var) objType()      {}
func (a Param) objType()    {}
func (a Field) objType()    {}

// Alias defines the alias object type
type Alias struct {
	BaseType *Object
	Size     int
}

// Type defines the alias object type
type Type struct {
	BaseType *Object
	Size     int
}

// Array defines the array object type
type Array struct {
	ElemType    *Object
	NumElements int
	Size        int
}

// Struct defines the struct object type
type Struct struct {
	Fields *Object
	Size   int
}

// Function defines the function object type
type Function struct {
	PRetType *Object
	PParams  *Object
	Index    int
	Params   int
	Vars     int
}

// Var defines the var object type
type Var struct {
	PType *Object
	Index int
	Size  int
}

// Param defines the var object type
type Param struct {
	PType *Object
	Index int
	Size  int
}

// Field defines the var object type
type Field struct {
	PType *Object
	Index int
	Size  int
}

// Analyser is the scope analyser
type Analyser struct {
	symbolTable [maxNestLevel]*Object
	level       int
}

// NewBlock opens a new block
func (a *Analyser) NewBlock() int {
	a.level++
	a.symbolTable[a.level] = nil
	return a.level
}

// EndBlock ends a block
func (a *Analyser) EndBlock() int {
	a.level--
	return a.level
}

// DefineSymbol defines a symbol given its name
func (a *Analyser) DefineSymbol(name int) *Object {
	obj := &Object{}

	obj.Name = name
	obj.Kind = KindUndefined
	obj.Next = a.symbolTable[a.level]
	a.symbolTable[a.level] = obj

	return obj
}

// SearchLocalSymbol searches for a symbol locally
func (a *Analyser) SearchLocalSymbol(name int) *Object {
	obj := a.symbolTable[a.level]

	for obj != nil {

		if obj.Name == name {
			return obj
		}

		obj = obj.Next
	}

	return obj
}

// SearchGlobalSymbol searches for a symbol globally
func (a *Analyser) SearchGlobalSymbol(name int) *Object {
	var obj *Object

	for i := a.level; i >= 0; i-- {
		obj = a.symbolTable[i]

		for obj != nil {
			if obj.Name == name {
				return obj
			}

			obj = obj.Next
		}

	}

	return obj
}

// CheckTypes returns true if objects are of same type
func (a *Analyser) CheckTypes(p1, p2 *Object) bool {
	if p1 == p2 {
		return true
	} else if p1 == PUniversalObj || p2 == PUniversalObj {
		return true
	} else if p1.Kind == KindUniversal || p2.Kind == KindUniversal {
		return true
	} else if p1.Kind == KindAliasType && p2.Kind != KindAliasType {
		alias := p1.T.(Alias)
		return a.CheckTypes(alias.BaseType, p2)
	} else if p1.Kind != KindAliasType && p2.Kind == KindAliasType {
		alias := p2.T.(Alias)
		return a.CheckTypes(p1, alias.BaseType)
	} else if p1.Kind == p1.Kind {
		if p1.Kind == KindAliasType {
			a1 := p1.T.(Alias)
			a2 := p2.T.(Alias)
			return a.CheckTypes(a1.BaseType, a2.BaseType)
		} else if p1.Kind == KindArrayType {
			a1 := p1.T.(Array)
			a2 := p2.T.(Array)
			if a1.NumElements == a2.NumElements {
				return a.CheckTypes(a1.ElemType, a2.ElemType)
			}
		} else if p1.Kind == KindStructType {
			s1 := p1.T.(Struct)
			s2 := p2.T.(Struct)

			f1 := s1.Fields
			f2 := s2.Fields
			if f1 != nil && f2 != nil {
				// TODO
			}
		}
	}

	return false
}

func (o *Object) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Kind: %v Name: %v Type: ", o.Kind, o.Name))
	switch o.T.(type) {
	case Alias:
		sb.WriteString("Alias")
	case Type:
		sb.WriteString("Type")
	case Array:
		sb.WriteString("Array")
	case Struct:
		sb.WriteString("Struct")
	case Function:
		sb.WriteString("Function")
	case Var:
		sb.WriteString("Var")
	case Param:
		sb.WriteString("Param")
	case Field:
		sb.WriteString("Field")
	default:
		sb.WriteString("INVALID")
	}

	return sb.String()
}
