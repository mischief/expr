
//line y.y:2

package expr
import __yyfmt__ "fmt"
//line y.y:3
		
import (
	"math"
	"fmt"
	"unicode"
	"unicode/utf8"
)

var (
	stack nstack
)


//line y.y:18
type yySymType struct {
	yys int
	n Node
	v float64 /* number */
	i string /* ident */
}

const IDENT = 57346
const NUMBER = 57347
const UMINUS = 57348

var yyToknames = []string{
	"IDENT",
	"NUMBER",
	" =",
	" +",
	" -",
	" *",
	" /",
	" %",
	" ^",
	" ,",
	"UMINUS",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line y.y:122


const eof = 0

type yyLex struct {
	line []byte
	peek rune
}

func NewLexer(s string) *yyLex {
	lex := &yyLex{
		line: []byte(s),
	}

	return lex
}

/* 1234 + 5678 */
func (l *yyLex) Lex(lval *yySymType) int {
	var c rune
	for {
		c = l.next()
		switch c {
		case '+', '-', '/', '*', '^', '%', '=', '(', ')', ',':
			return int(c)
		case ' ', '\t', '\n', '\r':
			/* Eat whitespace */
		case '\x00':
			/* eof */
			return eof
		default:
			switch {
			case unicode.IsDigit(c):
				/* number */
				return l.num(c, lval)
			case unicode.IsLetter(c):
				/* identifier */
				return l.ident(c, lval)
			default:
				fmt.Printf("unexpected %q\n", c)
				return eof
			}
		}
	}

	fmt.Printf("unexpected %q\n", c)
	return eof
}

func (l *yyLex) Error(s string) {
	fmt.Printf("%s: at: %s\n", s, l.line)
}

func (l *yyLex) num(c rune, lval *yySymType) int {
	high := float64(int(c - '0'))
	low := float64(0)
	lowdigit := 0
	dot := false
loop:
	for {
		c = l.next()
		switch {
		case unicode.IsDigit(c):
			if dot == false {
				high = high * 10 + float64(int(c - '0'))
			} else {
				low = low * 10 + float64(int(c - '0'))
				lowdigit++
			}
		case c == '.':
			dot = true
		default:
			break loop
		}
	}

	if c != 0 {
		l.peek = c
	}

	lval.v = high
	lval.v += low * math.Pow(0.1, float64(lowdigit))

	return NUMBER
}

func (l *yyLex) ident(c rune, lval *yySymType) int {
	str := string(c)
loop:
	for {
		c = l.next()
		switch {
		case unicode.IsLetter(c):
			str += string(c)
		default:
		break loop
		}
	}

	if c != 0 {
		l.peek = c
	}

/*
	switch str {
	case "x", "y":
		return int(str[0])
	case "new":
		lval.s = str
		return NEW
	case "log", "sin", "cos", "sqrt":
		lval.s = str
		return FN
	}
*/

	lval.i = str
	return IDENT
}

func (l *yyLex) next() rune {
	if l.peek != eof {
		r := l.peek
		l.peek = eof
		return r
	}

	if len(l.line) == 0 {
		return eof
	}

	r, sz := utf8.DecodeRune(l.line)
	l.line = l.line[sz:]
	if r == utf8.RuneError && sz == 1 {
		fmt.Printf("invalid character\n")
		return l.next()
	}

	return r
}

func (l *yyLex) follow(x rune, y int, z rune) int {
	r := l.next()

	if r == x {
		return y
	}

	l.peek = r
	return int(z)
}


//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 17
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 59

var yyAct = []int{

	3, 11, 12, 13, 14, 15, 16, 18, 14, 1,
	20, 21, 22, 23, 24, 25, 17, 2, 27, 28,
	9, 10, 11, 12, 13, 14, 19, 0, 0, 29,
	9, 10, 11, 12, 13, 14, 0, 0, 0, 26,
	9, 10, 11, 12, 13, 14, 8, 6, 7, 0,
	0, 5, 9, 10, 11, 12, 13, 14, 4,
}
var yyPact = []int{

	-1000, 43, -1000, 33, 43, 43, 1, -1000, 43, 43,
	43, 43, 43, 43, 43, 23, -1000, 43, 43, -1000,
	-8, -8, -4, -4, -4, -1000, -1000, 13, 45, -1000,
}
var yyPgo = []int{

	0, 0, 9, 17,
}
var yyR1 = []int{

	0, 2, 2, 3, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1,
}
var yyR2 = []int{

	0, 0, 2, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 1, 4, 3, 1,
}
var yyChk = []int{

	-1000, -2, -3, -1, 15, 8, 4, 5, 13, 7,
	8, 9, 10, 11, 12, -1, -1, 15, 6, -3,
	-1, -1, -1, -1, -1, -1, 16, -1, -1, 16,
}
var yyDef = []int{

	1, -2, 2, 3, 0, 0, 13, 16, 0, 0,
	0, 0, 0, 0, 0, 0, 12, 0, 0, 4,
	6, 7, 8, 9, 10, 11, 5, 0, 15, 14,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 11, 3, 3,
	15, 16, 9, 7, 13, 8, 3, 10, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 6, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 12,
}
var yyTok2 = []int{

	2, 3, 4, 5, 14,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 3:
		//line y.y:46
		{
			//stack.Push($1)
	}
	case 4:
		//line y.y:49
		{
			//stack.Push($1)
		//stack.Push($3)
	}
	case 5:
		//line y.y:56
		{
			yyVAL.n = yyS[yypt-1].n
			stack.Pop()
			stack.Push(yyVAL.n)
		}
	case 6:
		//line y.y:61
		{
			yyVAL.n = Expr(EAdd, yyS[yypt-2].n, yyS[yypt-0].n)
			stack.Pop()
			stack.Pop()
			stack.Push(yyVAL.n)
		}
	case 7:
		//line y.y:67
		{
			yyVAL.n = Expr(ESub, yyS[yypt-2].n, yyS[yypt-0].n)
			stack.Pop()
			stack.Pop()
			stack.Push(yyVAL.n)
		}
	case 8:
		//line y.y:73
		{
			yyVAL.n = Expr(EMul, yyS[yypt-2].n, yyS[yypt-0].n)
			stack.Pop()
			stack.Pop()
			stack.Push(yyVAL.n)
		}
	case 9:
		//line y.y:79
		{
			yyVAL.n = Expr(EDiv, yyS[yypt-2].n, yyS[yypt-0].n)
			stack.Pop()
			stack.Pop()
			stack.Push(yyVAL.n)
		}
	case 10:
		//line y.y:85
		{
			yyVAL.n = Expr(EMod, yyS[yypt-2].n, yyS[yypt-0].n)
			stack.Pop()
			stack.Pop()
			stack.Push(yyVAL.n)
		}
	case 11:
		//line y.y:91
		{
			yyVAL.n = Expr(EPow, yyS[yypt-2].n, yyS[yypt-0].n)
			stack.Pop()
			stack.Pop()
			stack.Push(yyVAL.n)
		}
	case 12:
		//line y.y:97
		{
			yyVAL.n = Expr(ENeg, yyS[yypt-0].n, Number(0))
			stack.Pop()
			stack.Push(yyVAL.n)
		}
	case 13:
		//line y.y:102
		{
			yyVAL.n = Ident(IGet, yyS[yypt-0].i, nil)
			stack.Push(yyVAL.n)
		}
	case 14:
		//line y.y:106
		{
			yyVAL.n = Ident(ICall, yyS[yypt-3].i, yyS[yypt-1].n)
			stack.Pop()
			stack.Push(yyVAL.n)
		}
	case 15:
		//line y.y:111
		{
			yyVAL.n = Ident(ISet, yyS[yypt-2].i, yyS[yypt-0].n)
			stack.Pop()
			stack.Push(yyVAL.n)
		}
	case 16:
		//line y.y:116
		{
			yyVAL.n = Number(yyS[yypt-0].v)
			stack.Push(yyVAL.n)
		}
	}
	goto yystack /* stack new state and value */
}
