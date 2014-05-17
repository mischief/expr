package expr

import (
	"testing"
)

func TestYacc(t *testing.T) {
	yyDebug = 1

	exprs := []nstack{
		run("a = 3"),
		run("b = 3"),
		run("c = 3"),
	}

	for _, s := range exprs {
		e := s.Pop()
		t.Logf("%s -> %s", e, eval(e))
	}

	res := run("a * b * c")

	t.Logf("result: %f", eval(res.Pop()).(*NumberNode).v)

	for stack.Size() > 0 {
		t.Logf("%v", stack.Pop())
	}
}
