package expr

import (
	"fmt"
	"math"
)

type Node interface {
	String() string
}

type ExprOp int

func (eo ExprOp) String() string {
	str := "+-*/^%"
	return string(str[eo])
}

const (
	EAdd ExprOp = iota
	ESub
	EMul
	EDiv
	EPow
	EMod
	ENeg
)

type ExprNode struct {
	op ExprOp
	l  Node
	r  Node
}

func Expr(op ExprOp, l, r Node) *ExprNode {
	return &ExprNode{op, l, r}
}

func (e ExprNode) String() string {
	return fmt.Sprintf("%s %s %s", e.l, e.op, e.r)
}

type NumberNode struct {
	v float64
}

func Number(v float64) *NumberNode {
	return &NumberNode{v}
}

func (n NumberNode) String() string {
	return fmt.Sprintf("%f", n.v)
}

type IdentOp int

const (
	IGet IdentOp = iota
	ISet
	ICall
)

type IdentNode struct {
	op   IdentOp
	name string
	val  Node
}

func Ident(op IdentOp, name string, val Node) *IdentNode {
	return &IdentNode{op, name, val}
}

func (i IdentNode) String() string {
	switch i.op {
	case IGet:
		return fmt.Sprintf("%s", i.name)
	case ISet:
		return fmt.Sprintf("%s = %s", i.name, i.val)
	case ICall:
		return fmt.Sprintf("%s(%s)", i.name, i.val)
	}

	return "unknown ident type"
}

var Idents = map[string]Node{}
var Funcs = map[string]func(float64) float64{}

func init() {
	Funcs["sqrt"] = math.Sqrt
}

func eval(n Node) (rv Node) {
	switch n.(type) {
	case *ExprNode:
		e := n.(*ExprNode)
		l := eval(e.l).(*NumberNode)
		r := eval(e.r).(*NumberNode)
		switch e.op {
		case EAdd:
			rv = Number(l.v + r.v)
		case ESub:
			rv = Number(l.v - r.v)
		case EMul:
			rv = Number(l.v * r.v)
		case EDiv:
			if r.v == 0 {
				rv = Number(0)
			} else {
				rv = Number(l.v / r.v)
			}
		case EMod:
			if r.v == 0 {
				rv = Number(0)
			} else {
				rv = Number(math.Mod(l.v, r.v))
			}
		case EPow:
			rv = Number(math.Pow(l.v, r.v))
		case ENeg:
			rv = Number(-l.v)
		}
	case *NumberNode:
		rv = n
	case *IdentNode:
		i := n.(*IdentNode)
		switch i.op {
		case IGet:
			if r, ok := Idents[i.name]; ok {
				rv = r
			} else {
				rv = Number(0)
				//panic(fmt.Sprintf("%s doesn't exist", i.name))
			}
		case ISet:
			Idents[i.name] = eval(i.val)
			rv = Idents[i.name]
		case ICall:
			if r, ok := Funcs[i.name]; ok {
				rv = Number(r(eval(i.val).(*NumberNode).v))
			} else {
			}
		}
	}

	return
}

/* run some code, get a node */
func run(s string) nstack {
	for stack.Size() > 0 {
		stack.Pop()
	}
	yyParse(NewLexer(s))
	return stack
}

/* convert result to string, trap panics */
func Proteval(s string) (rs string) {
	defer func() {
		if r := recover(); r != nil {
			rs = fmt.Sprintf("recover: %s\n", r)
		}
	}()

	var res []Node

	r := run(s)
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
	return out
}

