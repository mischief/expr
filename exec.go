package expr

func (n *Node) Bool() bool {
	if n.Op != OCONST {
		panic("not const")
	}

	switch n.Type {
	case TINT:
		if n.ival != 0 {
			return true
		}
	case TFLOAT:
		if n.fval != 0.0 {
			return true
		}
	case TSTRING:
		if len(n.sval) > 0 {
			return true
		}
	case TLIST:
		// TODO: lists
	}

	return false
}
