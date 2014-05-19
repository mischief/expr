package expr

import (
	"fmt"
)

var expop map[Op]func(*Node) *Node

func init() {
	expop = map[Op]func(*Node) *Node{
		ONAME:  oname,
		OCONST: oconst,
		OMUL:   omul,
		ODIV:   odiv,
		OMOD:   omod,
		OADD:   oadd,
		OSUB:   osub,
		ORSH:   orsh,
		OLSH:   olsh,
		OLT:    olt,
		OGT:    ogt,
		OLEQ:   oleq,
		OGEQ:   ogeq,
		OEQ:    oeq,
		ONEQ:   oneq,
		OLAND:  oland,
		OXOR:   oxor,
		OLOR:   olor,
		OCAND:  ocand,
		OCOR:   ocor,
		ONOT:   onot,
	}
}

func expr(n *Node) *Node {
	if n == nil {
		return nil
	}

	opfun, ok := expop[n.Op]
	if !ok {
		panic(fmt.Sprintf("missing op %s", n.Op))
	}
	return opfun(n)
}

func oname(n *Node) *Node {
	return nil
}

func oconst(n *Node) *Node {
	res := &Node{
		Op:    OCONST,
		Type:  n.Type,
		Store: n.Store,
	}

	return res
}

func omul(n *Node) *Node {
	l := expr(n.Left)
	r := expr(n.Right)

	res := &Node{
		Op:   OCONST,
		Type: TFLOAT,
	}

	switch l.Type {
	default:
		panic(fmt.Sprintf("bad lhs type for * %s", l.Type))
	case TINT:
		switch r.Type {
		case TINT:
			res.Type = TINT
			res.ival = l.ival * r.ival
		case TFLOAT:
			res.fval = float64(l.ival) * r.fval
		default:
			panic(fmt.Sprintf("bad rhs type for * %s", r.Type))
		}
	case TFLOAT:
		switch r.Type {
		case TINT:
			res.fval = l.fval * float64(r.ival)
		case TFLOAT:
			res.fval = l.fval * r.fval
		default:
			panic(fmt.Sprintf("bad rhs type for * %s", r.Type))
		}
	/* TSTRING * TINT is prone to abuse. */
	/*
	case TSTRING:
		res.Type = TSTRING
		switch r.Type {
		case TINT:
			res.sval = strings.Repeat(l.sval, int(r.ival))
		default:
			panic(fmt.Sprintf("bad rhs type for * %s", r.Type))
		}
	*/
	}

	return res
}

func odiv(n *Node) *Node {
	l := expr(n.Left)
	r := expr(n.Right)

	res := &Node{
		Op:   OCONST,
		Type: TFLOAT,
	}

	switch l.Type {
	default:
		panic(fmt.Sprintf("bad lhs type for / %s", l.Type))
	case TINT:
		switch r.Type {
		case TINT:
			res.Type = TINT
			if r.ival == 0 {
				panic("divide by zero")
			}
			res.ival = l.ival / r.ival
		case TFLOAT:
			if r.fval == 0 {
				panic("divide by zero")
			}
			res.fval = float64(l.fval) / r.fval
		default:
			panic(fmt.Sprintf("bad rhs type for / %s", r.Type))
		}
	case TFLOAT:
		switch r.Type {
		case TINT:
			if r.ival == 0 {
				panic("divide by zero")
			}
			res.fval = l.fval / float64(r.ival)
		case TFLOAT:
			if r.fval == 0 {
				panic("divide by zero")
			}
			res.fval = l.fval / r.fval
		default:
			panic(fmt.Sprintf("bad rhs type for / %s", r.Type))
		}
	}

	return res

}

func omod(n *Node) *Node {
	l := expr(n.Left)
	r := expr(n.Right)

	res := &Node{
		Op:   OCONST,
		Type: TINT,
	}

	if l.Type != TINT || r.Type != TINT {
		panic(fmt.Sprintf("bad expr %s %% %s", l.Type, r.Type))
	}

	res.ival = l.ival % r.ival

	return res
}

func oadd(n *Node) *Node {
	if n.Right == nil {
		/* unary + */
		return expr(n.Left)
	}

	l := expr(n.Left)
	r := expr(n.Right)

	res := &Node{
		Op:   OCONST,
		Type: TFLOAT,
	}

	switch l.Type {
	default:
		panic(fmt.Sprintf("bad lhs type for + %s", l.Type))
	case TINT:
		switch r.Type {
		case TINT:
			res.Type = TINT
			res.ival = l.ival + r.ival
		case TFLOAT:
			res.fval = float64(l.ival) + r.fval
		default:
			panic(fmt.Sprintf("bad rhs type for + %s", r.Type))
		}
	case TFLOAT:
		switch r.Type {
		case TINT:
			res.fval = l.fval + float64(r.ival)
		case TFLOAT:
			res.fval = l.fval + r.fval
		default:
			panic(fmt.Sprintf("bad rhs type for + %s", r.Type))
		}
	case TSTRING:
		res.Type = TSTRING
		if r.Type == TSTRING {
			res.sval = l.sval + r.sval
			break
		}
		if r.Type == TINT {
			res.sval = fmt.Sprintf("%s%c", l.sval, rune(r.ival))
			break
		}
		panic(fmt.Sprintf("bad rhs type for + %s", r.Type))
	case TLIST:
		panic("list+ unimplemented")
	}

	return res
}

func osub(n *Node) *Node {
	l := expr(n.Left)
	r := expr(n.Right)

	res := &Node{
		Op:   OCONST,
		Type: TFLOAT,
	}

	switch l.Type {
	default:
		panic(fmt.Sprintf("bad lhs type for - %s", l.Type))
	case TINT:
		switch r.Type {
		case TINT:
			res.Type = TINT
			res.ival = l.ival - r.ival
		case TFLOAT:
			res.fval = float64(l.ival) - r.fval
		default:
			panic(fmt.Sprintf("bad rhs type for - %s", r.Type))
		}
	case TFLOAT:
		switch r.Type {
		case TINT:
			res.fval = l.fval - float64(r.ival)
		case TFLOAT:
			res.fval = l.fval - r.fval
		default:
			panic(fmt.Sprintf("bad rhs type for - %s", r.Type))
		}
	}

	return res
}

func orsh(n *Node) *Node {
	l := expr(n.Left)
	r := expr(n.Right)

	res := &Node{
		Op:   OCONST,
		Type: TINT,
	}

	if l.Type != TINT || r.Type != TINT {
		panic(fmt.Sprintf("bad expr %s >> %s", l.Type, r.Type))
	}

	res.ival = l.ival >> uint64(r.ival)

	return res
}

func olsh(n *Node) *Node {
	l := expr(n.Left)
	r := expr(n.Right)

	res := &Node{
		Op:   OCONST,
		Type: TINT,
	}

	if l.Type != TINT || r.Type != TINT {
		panic(fmt.Sprintf("bad expr %s << %s", l.Type, r.Type))
	}

	res.ival = l.ival << uint64(r.ival)

	return res
}

func olt(n *Node) *Node {
	l := expr(n.Left)
	r := expr(n.Right)

	res := &Node{
		Op:   OCONST,
		Type: TINT,
	}

	switch l.Type {
	default:
		panic(fmt.Sprintf("bad lhs type for < %s", l.Type))
	case TINT:
		switch r.Type {
		case TINT:
			if l.ival < r.ival {
				res.ival = 1
			} else {
				res.ival = 0
			}
		case TFLOAT:
			if float64(l.ival) < r.fval {
				res.ival = 1
			} else {
				res.ival = 0
			}
		default:
			panic(fmt.Sprintf("bad rhs type for < %s", r.Type))
		}
	case TFLOAT:
		switch r.Type {
		case TINT:
			if l.fval < float64(r.ival) {
				res.ival = 1
			} else {
				res.ival = 0
			}
		case TFLOAT:
			if l.fval < r.fval {
				res.ival = 1
			} else {
				res.ival = 0
			}
		default:
			panic(fmt.Sprintf("bad rhs type for < %s", r.Type))
		}
	}

	return res
}

func ogt(n *Node) *Node {
	l := expr(n.Left)
	r := expr(n.Right)

	res := &Node{
		Op:   OCONST,
		Type: TINT,
	}

	switch l.Type {
	default:
		panic(fmt.Sprintf("bad lhs type for > %s", l.Type))
	case TINT:
		switch r.Type {
		case TINT:
			if l.ival > r.ival {
				res.ival = 1
			} else {
				res.ival = 0
			}
		case TFLOAT:
			if float64(l.ival) > r.fval {
				res.ival = 1
			} else {
				res.ival = 0
			}
		default:
			panic(fmt.Sprintf("bad rhs type for > %s", r.Type))
		}
	case TFLOAT:
		switch r.Type {
		case TINT:
			if l.fval > float64(r.ival) {
				res.ival = 1
			} else {
				res.ival = 0
			}
		case TFLOAT:
			if l.fval > r.fval {
				res.ival = 1
			} else {
				res.ival = 0
			}
		default:
			panic(fmt.Sprintf("bad rhs type for > %s", r.Type))
		}
	}

	return res
}

func oleq(n *Node) *Node {
	l := expr(n.Left)
	r := expr(n.Right)

	res := &Node{
		Op:   OCONST,
		Type: TINT,
	}

	switch l.Type {
	default:
		panic(fmt.Sprintf("bad lhs type for <= %s", l.Type))
	case TINT:
		switch r.Type {
		case TINT:
			if l.ival <= r.ival {
				res.ival = 1
			} else {
				res.ival = 0
			}
		case TFLOAT:
			if float64(l.ival) <= r.fval {
				res.ival = 1
			} else {
				res.ival = 0
			}
		default:
			panic(fmt.Sprintf("bad rhs type for <= %s", r.Type))
		}
	case TFLOAT:
		switch r.Type {
		case TINT:
			if l.fval <= float64(r.ival) {
				res.ival = 1
			} else {
				res.ival = 0
			}
		case TFLOAT:
			if l.fval <= r.fval {
				res.ival = 1
			} else {
				res.ival = 0
			}
		default:
			panic(fmt.Sprintf("bad rhs type for <= %s", r.Type))
		}
	}

	return res
}

func ogeq(n *Node) *Node {
	l := expr(n.Left)
	r := expr(n.Right)

	res := &Node{
		Op:   OCONST,
		Type: TINT,
	}

	switch l.Type {
	default:
		panic(fmt.Sprintf("bad lhs type for >= %s", l.Type))
	case TINT:
		switch r.Type {
		case TINT:
			if l.ival >= r.ival {
				res.ival = 1
			} else {
				res.ival = 0
			}
		case TFLOAT:
			if float64(l.ival) >= r.fval {
				res.ival = 1
			} else {
				res.ival = 0
			}
		default:
			panic(fmt.Sprintf("bad rhs type for >= %s", r.Type))
		}
	case TFLOAT:
		switch r.Type {
		case TINT:
			if l.fval >= float64(r.ival) {
				res.ival = 1
			} else {
				res.ival = 0
			}
		case TFLOAT:
			if l.fval >= r.fval {
				res.ival = 1
			} else {
				res.ival = 0
			}
		default:
			panic(fmt.Sprintf("bad rhs type for >= %s", r.Type))
		}
	}

	return res
}

func oeq(n *Node) *Node {
	l := expr(n.Left)
	r := expr(n.Right)

	res := &Node{
		Op:   OCONST,
		Type: TINT,
	}

	switch l.Type {
	default:
		panic(fmt.Sprintf("bad lhs type for == %s", l.Type))
	case TINT:
		switch r.Type {
		case TINT:
			if l.ival == r.ival {
				res.ival = 1
			} else {
				res.ival = 0
			}
		case TFLOAT:
			if float64(l.ival) == r.fval {
				res.ival = 1
			} else {
				res.ival = 0
			}
		default:
			panic(fmt.Sprintf("bad rhs type for == %s", r.Type))
		}
	case TFLOAT:
		switch r.Type {
		case TINT:
			if l.fval == float64(r.ival) {
				res.ival = 1
			} else {
				res.ival = 0
			}
		case TFLOAT:
			if l.fval == r.fval {
				res.ival = 1
			} else {
				res.ival = 0
			}
		default:
			panic(fmt.Sprintf("bad rhs type for == %s", r.Type))
		}
	}

	return res
}

func oneq(n *Node) *Node {
	l := expr(n.Left)
	r := expr(n.Right)

	res := &Node{
		Op:   OCONST,
		Type: TINT,
	}

	switch l.Type {
	default:
		panic(fmt.Sprintf("bad lhs type for != %s", l.Type))
	case TINT:
		switch r.Type {
		case TINT:
			if l.ival != r.ival {
				res.ival = 1
			} else {
				res.ival = 0
			}
		case TFLOAT:
			if float64(l.ival) != r.fval {
				res.ival = 1
			} else {
				res.ival = 0
			}
		default:
			panic(fmt.Sprintf("bad rhs type for != %s", r.Type))
		}
	case TFLOAT:
		switch r.Type {
		case TINT:
			if l.fval != float64(r.ival) {
				res.ival = 1
			} else {
				res.ival = 0
			}
		case TFLOAT:
			if l.fval != r.fval {
				res.ival = 1
			} else {
				res.ival = 0
			}
		default:
			panic(fmt.Sprintf("bad rhs type for != %s", r.Type))
		}
	}

	return res
}

func oland(n *Node) *Node {
	l := expr(n.Left)
	r := expr(n.Right)

	res := &Node{
		Op:   OCONST,
		Type: TINT,
	}

	if l.Type != TINT || r.Type != TINT {
		panic(fmt.Sprintf("bad expr %s ^ %s", l.Type, r.Type))
	}

	res.ival = l.ival & r.ival

	return res
}

func oxor(n *Node) *Node {
	l := expr(n.Left)
	r := expr(n.Right)

	res := &Node{
		Op:   OCONST,
		Type: TINT,
	}

	if l.Type != TINT || r.Type != TINT {
		panic(fmt.Sprintf("bad expr %s ^ %s", l.Type, r.Type))
	}

	res.ival = l.ival ^ r.ival

	return res
}

func olor(n *Node) *Node {
	l := expr(n.Left)
	r := expr(n.Right)

	res := &Node{
		Op:   OCONST,
		Type: TINT,
	}

	if l.Type != TINT || r.Type != TINT {
		panic(fmt.Sprintf("bad expr %s ^ %s", l.Type, r.Type))
	}

	res.ival = l.ival | r.ival

	return res
}

func ocand(n *Node) *Node {
	res := &Node{
		Op:   OCONST,
		Type: TINT,
	}

	res.ival = 0

	if l := expr(n.Left); l.Bool() == false {
		return res
	}
	if r := expr(n.Left); r.Bool() == false {
		return res
	}

	res.ival = 1
	return res
}

func ocor(n *Node) *Node {
	res := &Node{
		Op:   OCONST,
		Type: TINT,
	}

	res.ival = 0

	if l := expr(n.Left); l.Bool() == true {
		res.ival = 1
	}

	if r := expr(n.Left); r.Bool() == true {
		res.ival = 1
	}

	return res
}

func onot(n *Node) *Node {
	res := &Node{
		Op:   OCONST,
		Type: TINT,
	}

	res.ival = 0

	if l := expr(n.Left); l.Bool() == false {
		res.ival = 1
	}

	return res
}
