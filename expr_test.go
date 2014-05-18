package expr

import (
	"testing"
)

func TestYacc(t *testing.T) {
	yyDebug = 4

	exprs := []string{
		`1024 * 1024`,
		`"foo" + "bar"`,
		`5 * (1000 + 100) * 5`,
		`8<<8`,
	}

	for _, s := range exprs {
		r := Proteval(s)
		t.Logf("%s -> %s", s, r)
	}
}
