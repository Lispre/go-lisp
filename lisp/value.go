package lisp

import (
	"fmt"
	"strconv"
)

type Value struct {
	Kind Kind
	raw  interface{}
}

var Nil = NewValue(nil)

type Kind uint

const (
	InvalidKind Kind = iota
	NilKind
	SymbolKind
	NumberKind
	SexpKind
	ProcKind
)

func NewValue(from interface{}) Value {
	v := Value{raw: from}
	switch from.(type) {
	case string:
		v.Kind = SymbolKind
	case int:
		v.Kind = NumberKind
		v.raw = float64(from.(int))
	case float64:
		v.Kind = NumberKind
	case nil:
		v.Kind = NilKind
	case Sexp:
		v.Kind = SexpKind
	case Proc:
		v.Kind = ProcKind
	case bool:
		v.Kind = SymbolKind
		if from.(bool) {
			v.raw = "true"
		} else {
			v.raw = "false"
		}
	default:
		v.Kind = InvalidKind
	}
	return v
}

func (v Value) IsA(k Kind) bool {
	return k == v.Kind
}

func (v Value) String() string {
	if v.IsA(NumberKind) {
		return strconv.FormatFloat(v.raw.(float64), 'f', -1, 64)
	} else {
		return fmt.Sprintf("%v", v.raw)
	}
}

func (v Value) Sexp() Sexp {
	return v.raw.(Sexp)
}

func (v Value) Proc() Proc {
	return v.raw.(Proc)
}

func (v Value) Number() float64 {
	return v.raw.(float64)
}
