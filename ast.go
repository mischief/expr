package expr

import (
	"fmt"
)

type Type int

const (
	TINT Type = iota
	TFLOAT
	TSTRING
	TLIST
	TCODE
)

var Typenames = map[Type]string{
	TINT:    "int",
	TFLOAT:  "float",
	TSTRING: "string",
	TLIST:   "list",
	TCODE:   "code",
}

func (t Type) String() string {
	return Typenames[t]
}

type Op int

const (
	ONAME Op = iota
	OCONST
	OMUL
	ODIV
	OMOD
	OADD
	OSUB
	ORSH
	OLSH
	OLT
	OGT
	OLEQ
	OGEQ
	OEQ
	ONEQ
	OLAND
	OXOR
	OLOR
	OCAND
	OCOR
	OASGN
	OINDM
	OEDEC
	OEINC
	OPINC
	OPDEC
	ONOT
	OIF
	ODO
	OLIST
	OCALL
	OCTRUCT
	OWHILE
	OELSE
	OHEAD
	OTAIL
	OAPPEND
	ORET
	OINDEX
	OINDC
	ODOT
	OLOCAL
	OFRAME
	OCOMPLEX
	ODELETE
	OCAST
	OFMT
	OEVAL
	OWHAT
)

var Opnames = map[Op]string{
	ONAME:  "name",
	OCONST: "const",
	OMUL:   "mul",
	ODIV:   "div",
	OMOD:   "mod",
	OADD:   "add",
	OSUB:   "sub",
	ORSH:   "rsh",
	OLSH:   "lsh",
	OLT:    "lt",
	OGT:    "gt",
	OLEQ:   "leq",
	OGEQ:   "geq",
	OEQ:    "eq",
	ONEQ:   "neq",
	OLAND:  "land",
	OXOR:   "xor",
	OLOR:   "lor",
	OCAND:  "cand",
	OCOR:   "cor",
	OASGN:  "asgn",

	ONOT:  "not",
	OIF:   "if",
	ODO:   "do",
	OLIST: "list",
	OCALL: "call",

	OEVAL: "eval",
	OWHAT: "what",
}

func (o Op) String() string {
	n, ok := Opnames[o]
	if !ok {
		return fmt.Sprintf("unknown op %d", o)
	}
	return n
}

type Store struct {
	ival int64
	fval float64
	sval string
	lval List
	cc   *Node
}

type Node struct {
	Left, Right *Node

	Op
	Type
	Store
}

func (n Node) String() string {
	switch n.Op {
	case OCONST:
		switch n.Type {
		case TINT:
			return fmt.Sprintf("%s(%d)", n.Type, n.ival)
		case TFLOAT:
			return fmt.Sprintf("%s(%f)", n.Type, n.fval)
		case TSTRING:
			return fmt.Sprintf("%s(%s)", n.Type, n.sval)
		}
	}
	return fmt.Sprintf("%s %s %s", n.Left, n.Op, n.Right)
}

func AN(op Op, l, r *Node) *Node {
	return &Node{Op: op, Left: l, Right: r}
}

func Const(v int64) *Node {
	n := AN(OCONST, nil, nil)
	n.ival = v
	n.Type = TINT
	return n
}

type Lsym struct {
	Name   string
	lexval int

	Builtin func(n, res *Node)
}

/* run some code, get a node */
func run(s string) *yyLex {
	lex := NewLexer(s + ";")
	yyParse(lex)
	return lex
}

/* convert result to string, trap panics */
func Proteval(s string) (rs string) {
	defer func() {
		if r := recover(); r != nil {
			rs = fmt.Sprintf("recover: %s\n", r)
		}
	}()

	//var res []Node

	n := run(s)
	return fmt.Sprintf("%s", expr(n.node))
	/*
		for r.Size() > 0 {
			res = append([]Node{r.Pop()}, res...)
		}
		out := " ="
		for i, n := range res {
			if i > 0 {
				out += ","
			}
			out += fmt.Sprintf(" %s", eval(n))
		}
	*/
	return ""
}
