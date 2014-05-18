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
		`5 > 2`,
		`2 < 5`,
		`1 <= 1`,
		`0 >= 0`,
		`2 == 2`,
		`2 != 3`,
		`1 & 1`,
		`1 ^ 0`,
		`1 | 1`,
		`1 && 1`,
		`1 || 1`,
		`!0`,
	}

	for _, s := range exprs {
		r := Proteval(s)
		t.Logf("%s -> %s", s, r)
	}
}
