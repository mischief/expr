
//line y.y:2

package expr
import __yyfmt__ "fmt"
//line y.y:3
		
import (
	"fmt"
	"bytes"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)


//line y.y:16
type yySymType struct {
	yys int
	node *Node
	sym  *Lsym
	fval float64
	ival int64
	sval string
}

const Tfmt = 57346
const Toror = 57347
const Tandand = 57348
const Teq = 57349
const Tneq = 57350
const Tleq = 57351
const Tgeq = 57352
const Tlsh = 57353
const Trsh = 57354
const Tdec = 57355
const Tinc = 57356
const Tindir = 57357
const Tid = 57358
const Tfconst = 57359
const Tconst = 57360
const Tstring = 57361
const Tif = 57362
const Tdo = 57363
const Tthen = 57364
const Telse = 57365
const Twhile = 57366
const Tloop = 57367
const Thead = 57368
const Ttail = 57369
const Tappend = 57370
const Tfn = 57371
const Tret = 57372
const Tlocal = 57373
const Tcomplex = 57374
const Twhat = 57375
const Tdelete = 57376
const Teval = 57377
const Tbuiltin = 57378

var yyToknames = []string{
	" ;",
	" =",
	"Tfmt",
	"Toror",
	"Tandand",
	" |",
	" ^",
	" &",
	"Teq",
	"Tneq",
	" <",
	" >",
	"Tleq",
	"Tgeq",
	"Tlsh",
	"Trsh",
	" +",
	" -",
	" *",
	" /",
	" %",
	"Tdec",
	"Tinc",
	"Tindir",
	" .",
	" [",
	" (",
	"Tid",
	"Tfconst",
	"Tconst",
	"Tstring",
	"Tif",
	"Tdo",
	"Tthen",
	"Telse",
	"Twhile",
	"Tloop",
	"Thead",
	"Ttail",
	"Tappend",
	"Tfn",
	"Tret",
	"Tlocal",
	"Tcomplex",
	"Twhat",
	"Tdelete",
	"Teval",
	"Tbuiltin",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line y.y:274


const Eof = 0

var (
	keywords = map[string]int {
		"do":		Tdo,
		"if":		Tif,
		"then":		Tthen,
		"else":		Telse,
		"while":	Twhile,
		"loop":		Tloop,
		"head":		Thead,
		"tail":		Ttail,
		"append":	Tappend,
		"defn":		Tfn,
		"return":	Tret,
		"local":	Tlocal,
		"aggr":		Tcomplex,
		"union":	Tcomplex,
		"adt":		Tcomplex,
		"complex":	Tcomplex,
		"delete":	Tdelete,
		"whatis":	Twhat,
		"eval":		Teval,
		"builtin":	Tbuiltin,
	}
)

type yyLex struct {
	node *Node
	line []byte
	peek rune
}

func NewLexer(s string) *yyLex {
	lex := &yyLex{
		line: []byte(s),
		peek: Eof,
	}

	return lex
}

func (l *yyLex) Lex(lval *yySymType) int {
	lr := l.LexReal(lval)
	return lr
}

/* 1234 + 5678 */
func (l *yyLex) LexReal(lval *yySymType) int {
	var c rune
loop:
	for {
		c = l.next()
		switch c {
		case Eof:
			return Eof
		case '"':
			return l.string(lval)
		case ' ', '\t':
			goto loop

		case '/':
			c = l.next()
			if c != '/' {
				l.peek = c
				return '/'
			}
			l.eatnl()
		case '\n':
			return ';'
		case '.':
			c = l.next()
			l.peek = c
			if unicode.IsDigit(c) {
				return l.numsym('.', lval)
			}
			return '.'
		case '(', ')', '[', ']', ';', ':', ',', '~', '?', '*', '@', '^', '%':
			return int(c)

		/*
		case '{':
			stacked++;
			return c;
		case '}':
			stacked--;
			return c;

		case '\\':
			c = lexc();
			if(strchr(vfmt, c) == 0) {
				unlexc(c);
				return '\\';
			}
			yylval.ival = c;
			return Tfmt;
		*/

		case '!':
			c = l.next()
			if c == '=' {
				return Tneq
			}
			l.peek = c
			return '!'
		case '+':
			c = l.next()
			if c == '+' {
				return Tinc
			}
			l.peek = c
			return '+'
		case '\'':
			c = l.next()
			if c == '\\' {
				lval.ival = int64(c) //escchar(l.next());
			} else {
				lval.ival = int64(c)
			}
			c = l.next()
			if c != '\'' {
				panic("missing '")
				l.peek = c
			}
			return Tconst
		case '&':
			c = l.next()
			if c == '&' {
				return Tandand
			}
			l.peek = c
			return '&'

		case '=':
			c = l.next()
			if c == '=' {
				return Teq
			}
			l.peek = c
			return '='

		case '|':
			c = l.next()
			if c == '|' {
				return Toror
			}
			l.peek = c
			return '|'

		case '<':
			c = l.next()
			if c == '=' {
				return Tleq
			}
			if c == '<' {
				return Tlsh
			}
			l.peek = c
			return '<'

		case '>':
			c = l.next()
			if c == '=' {
				return Tgeq
			}
			if c == '>' {
				return Trsh
			}
			l.peek = c
			return '>'

		case '-':
			c = l.next()

			if c == '>' {
				return Tindir
			}

			if c == '-' {
				return Tdec
			}
			l.peek = c
			return '-'

		default:
			return l.numsym(c, lval)
		}
	}
}

func (l *yyLex) Error(s string) {
	fmt.Printf("%s at: %s\n", s, l.line)
}

func (l *yyLex) string(lval *yySymType) int {
	buf := new(bytes.Buffer)
	for {
		c := l.next()
		if c < 0 {
			panic("eof in string")
		}
		if c == '"' {
			break
		}
		buf.WriteRune(c)
	}

	lval.sval = buf.String()
	return Tstring
}

func (l *yyLex) eatnl() {
	for {
		c := l.next()
		switch c {
		case Eof:
			panic("eof in comment")
		case '\n':
			return
		}
	}
}

func (l *yyLex) numsym(c rune, lval *yySymType) int {
	var isbin, isfloat, ishex bool;
	var sel string

	//fmt.Printf("numsym %d %q\n", c, l.line)

	buf := new(bytes.Buffer)
	buf.WriteRune(c)

	if c == '.' {
		isfloat = true
	}

	if unicode.IsDigit(c) || isfloat {
		for {
			c = l.next()
			if c < 0 {
				panic("eof in numsym")
			}

			//if c == '\n' {}
			sel = "01234567890.xb";
			if ishex {
				sel = "01234567890abcdefABCDEF"
			} else if isbin {
				sel = "01"
			} else if isfloat {
				sel = "01234567890eE-+";
			}

			if !strings.ContainsRune(sel, c) {
				l.peek = c
				break
			}

			if c == '.' {
				isfloat = true
			}
			if !isbin && c == 'x' {
				ishex = true
			}
			if !ishex && c == 'b' {
				isbin = true
			}
			buf.WriteRune(c)
		}

		if isfloat {
			fval, err := strconv.ParseFloat(buf.String(), 64)
			if err != nil {
				panic("bad float " + buf.String() + " : " + err.Error())
			}
			lval.fval = fval
			return Tfconst
		}

		if isbin {
			ival, err := strconv.ParseInt(buf.String(), 2, 64)
			if err != nil {
				panic("bad binary " + buf.String() + " : " + err.Error())
			}
			lval.ival = ival
		} else {
			ival, err := strconv.ParseInt(buf.String(), 10, 64)
			if err != nil {
				panic("bad number " + buf.String() + " : " + err.Error())
			}
			lval.ival = ival
		}
		return Tconst
	}

	for {
		c = l.next()
		if c < 0 {
			panic("eof eating symbols")
		}
		//if c == '\n' {}
		if c != '_' && c != '$' && c <= '~' && !(unicode.IsDigit(c) || unicode.IsLetter(c)) {
			l.peek = c
			break
		}
		buf.WriteRune(c)
	}

	// s := look(buf.String())
	return Eof
}

func (l *yyLex) next() rune {
	r := l.realnext()
	//fmt.Printf("next %d(%c)\n", r, r)
	return r
}

func (l *yyLex) realnext() rune {
	if l.peek != Eof {
		r := l.peek
		l.peek = Eof
		return r
	}

	if len(l.line) == 0 {
		return Eof
	}

	r, sz := utf8.DecodeRune(l.line)
	l.line = l.line[sz:]
	if r == utf8.RuneError && sz == 1 {
		panic("invalid character")
		//return l.next()
	}

	return r
}


//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 10,
}

const yyNprod = 56
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 295

var yyAct = []int{

	8, 4, 89, 96, 96, 64, 99, 95, 87, 86,
	27, 50, 51, 52, 53, 54, 55, 56, 57, 58,
	59, 60, 49, 91, 48, 47, 46, 45, 44, 42,
	43, 38, 39, 40, 41, 37, 36, 34, 35, 31,
	32, 33, 22, 3, 62, 28, 10, 11, 2, 30,
	1, 12, 13, 31, 32, 33, 21, 27, 25, 24,
	26, 7, 29, 9, 90, 0, 63, 14, 15, 16,
	0, 0, 0, 65, 88, 17, 20, 23, 5, 66,
	0, 18, 19, 0, 0, 0, 0, 93, 94, 0,
	10, 11, 0, 90, 97, 12, 13, 0, 98, 0,
	21, 27, 25, 24, 26, 0, 0, 92, 0, 0,
	0, 14, 15, 16, 34, 35, 31, 32, 33, 17,
	20, 23, 5, 10, 11, 18, 19, 0, 12, 13,
	0, 0, 0, 21, 27, 25, 24, 26, 6, 0,
	0, 0, 0, 0, 14, 15, 16, 0, 0, 0,
	0, 0, 17, 20, 23, 0, 0, 0, 18, 19,
	61, 37, 36, 34, 35, 31, 32, 33, 0, 0,
	67, 68, 69, 70, 71, 72, 73, 74, 75, 76,
	77, 78, 79, 80, 81, 82, 83, 84, 85, 49,
	0, 48, 47, 46, 45, 44, 42, 43, 38, 39,
	40, 41, 37, 36, 34, 35, 31, 32, 33, 47,
	46, 45, 44, 42, 43, 38, 39, 40, 41, 37,
	36, 34, 35, 31, 32, 33, 46, 45, 44, 42,
	43, 38, 39, 40, 41, 37, 36, 34, 35, 31,
	32, 33, 45, 44, 42, 43, 38, 39, 40, 41,
	37, 36, 34, 35, 31, 32, 33, 44, 42, 43,
	38, 39, 40, 41, 37, 36, 34, 35, 31, 32,
	33, 42, 43, 38, 39, 40, 41, 37, 36, 34,
	35, 31, 32, 33, 38, 39, 40, 41, 37, 36,
	34, 35, 31, 32, 33,
}
var yyPact = []int{

	-1000, 70, -1000, -1000, 41, 70, 184, -1000, -1000, -1000,
	103, 103, 103, 103, 103, 103, 103, 103, 103, 103,
	103, 103, 14, -21, -1000, -1000, -1000, -53, -1000, 26,
	-1000, 103, 103, 103, 103, 103, 103, 103, 103, 103,
	103, 103, 103, 103, 103, 103, 103, 103, 103, 103,
	-1000, -1000, -1000, -1000, -1000, -1000, -45, -46, -1000, -1000,
	-1000, 17, 103, -7, -21, -1000, -1000, -1000, -1000, -1000,
	31, 31, 94, 94, 143, 143, 143, 143, 270, 270,
	259, 246, 232, 217, 201, 184, 103, 103, -1000, -50,
	-1000, 103, -1000, -1000, -1000, -1000, 103, -51, -1000, -1000,
}
var yyPgo = []int{

	0, 138, 0, 63, 43, 42, 2, 1, 62, 61,
	61, 50, 48,
}
var yyR1 = []int{

	0, 11, 11, 12, 10, 10, 8, 8, 4, 4,
	7, 7, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 9, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 3, 3, 3, 3, 3,
	3, 3, 5, 5, 6, 6,
}
var yyR2 = []int{

	0, 0, 2, 1, 0, 1, 1, 2, 2, 3,
	0, 1, 1, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 1, 1, 2, 2, 2, 2, 2, 2,
	4, 4, 2, 2, 2, 3, 4, 5, 1, 1,
	1, 1, 1, 3, 1, 3,
}
var yyChk = []int{

	-1000, -11, -12, -4, -7, 52, -1, -9, -2, -3,
	20, 21, 25, 26, 41, 42, 43, 49, 55, 56,
	50, 30, -5, 51, 33, 32, 34, 31, 4, -8,
	-4, 22, 23, 24, 20, 21, 19, 18, 14, 15,
	16, 17, 12, 13, 11, 10, 9, 8, 7, 5,
	-2, -2, -2, -2, -2, -2, -2, -2, -2, -2,
	-2, -1, 30, -5, 58, -4, 53, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, 54, 54, 57, -6,
	-7, 30, -5, -2, -2, 57, 54, -6, -7, 57,
}
var yyDef = []int{

	1, -2, 2, 3, 0, 10, 11, 12, 32, 33,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 48, 0, 49, 50, 51, 52, 8, 10,
	6, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	34, 35, 36, 37, 38, 39, 0, 0, 42, 43,
	44, 0, 10, 0, 0, 7, 9, 13, 14, 15,
	16, 17, 18, 19, 20, 21, 22, 23, 24, 25,
	26, 27, 28, 29, 30, 31, 0, 0, 45, 0,
	54, 10, 53, 40, 41, 46, 10, 0, 55, 47,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 55, 3, 3, 3, 24, 11, 3,
	30, 57, 22, 20, 54, 21, 28, 23, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 58, 4,
	14, 5, 15, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 29, 3, 3, 10, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 52, 9, 53, 56,
}
var yyTok2 = []int{

	2, 3, 6, 7, 8, 12, 13, 16, 17, 18,
	19, 25, 26, 27, 31, 32, 33, 34, 35, 36,
	37, 38, 39, 40, 41, 42, 43, 44, 45, 46,
	47, 48, 49, 50, 51,
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
		//line y.y:57
		{
				//Execrec($1)
			yylex.(*yyLex).node = yyS[yypt-0].node
			}
	case 4:
		//line y.y:64
		{ yyVAL.sym = nil; }
	case 5:
		yyVAL.sym = yyS[yypt-0].sym
	case 6:
		yyVAL.node = yyS[yypt-0].node
	case 7:
		//line y.y:70
		{
				yyVAL.node = AN(OLIST, yyS[yypt-1].node, yyS[yypt-0].node)
			}
	case 8:
		//line y.y:75
		{
				yyVAL.node = yyS[yypt-1].node
			}
	case 9:
		//line y.y:79
		{
				yyVAL.node = yyS[yypt-1].node
			}
	case 10:
		//line y.y:85
		{
				yyVAL.node = nil
			}
	case 11:
		yyVAL.node = yyS[yypt-0].node
	case 12:
		yyVAL.node = yyS[yypt-0].node
	case 13:
		//line y.y:93
		{
				yyVAL.node = AN(OMUL, yyS[yypt-2].node, yyS[yypt-0].node); 
			}
	case 14:
		//line y.y:97
		{
				yyVAL.node = AN(ODIV, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 15:
		//line y.y:101
		{
				yyVAL.node = AN(OMOD, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 16:
		//line y.y:105
		{
				yyVAL.node = AN(OADD, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 17:
		//line y.y:109
		{
				yyVAL.node = AN(OSUB, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 18:
		//line y.y:113
		{
				yyVAL.node = AN(ORSH, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 19:
		//line y.y:117
		{
				yyVAL.node = AN(OLSH, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 20:
		//line y.y:121
		{
				yyVAL.node = AN(OLT, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 21:
		//line y.y:125
		{
				yyVAL.node = AN(OGT, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 22:
		//line y.y:129
		{
				yyVAL.node = AN(OLEQ, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 23:
		//line y.y:133
		{
				yyVAL.node = AN(OGEQ, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 24:
		//line y.y:137
		{
				yyVAL.node = AN(OEQ, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 25:
		//line y.y:141
		{
				yyVAL.node = AN(ONEQ, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 26:
		//line y.y:145
		{
				yyVAL.node = AN(OLAND, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 27:
		//line y.y:149
		{
				yyVAL.node = AN(OXOR, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 28:
		//line y.y:153
		{
				yyVAL.node = AN(OLOR, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 29:
		//line y.y:157
		{
				yyVAL.node = AN(OCAND, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 30:
		//line y.y:161
		{
				yyVAL.node = AN(OCOR, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 31:
		//line y.y:165
		{
				yyVAL.node = AN(OASGN, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 32:
		yyVAL.node = yyS[yypt-0].node
	case 33:
		yyVAL.node = yyS[yypt-0].node
	case 34:
		//line y.y:174
		{
				yyVAL.node = AN(OADD, yyS[yypt-0].node, nil);
			}
	case 35:
		//line y.y:178
		{
				yyVAL.node = Const(0);
				yyVAL.node = AN(OSUB, yyVAL.node, yyS[yypt-0].node);
			}
	case 36:
		//line y.y:183
		{
				yyVAL.node = AN(OEDEC, yyS[yypt-0].node, nil);
			}
	case 37:
		//line y.y:187
		{
				yyVAL.node = AN(OEINC, yyS[yypt-0].node, nil);
			}
	case 38:
		//line y.y:191
		{
				yyVAL.node = AN(OHEAD, yyS[yypt-0].node, nil);
			}
	case 39:
		//line y.y:195
		{
				yyVAL.node = AN(OTAIL, yyS[yypt-0].node, nil);
			}
	case 40:
		//line y.y:199
		{
				yyVAL.node = AN(OAPPEND, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 41:
		//line y.y:203
		{
				yyVAL.node = AN(ODELETE, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 42:
		//line y.y:207
		{
				yyVAL.node = AN(ONOT, yyS[yypt-0].node, nil);
			}
	case 43:
		//line y.y:211
		{
				yyVAL.node = AN(OXOR, yyS[yypt-0].node, Const(-1));
			}
	case 44:
		//line y.y:215
		{
				yyVAL.node = AN(OEVAL, yyS[yypt-0].node, nil);
			}
	case 45:
		//line y.y:221
		{
				yyVAL.node = yyS[yypt-1].node
			}
	case 46:
		//line y.y:225
		{
				yyVAL.node = AN(OCALL, yyS[yypt-3].node, yyS[yypt-1].node);
			}
	case 47:
		//line y.y:229
		{
				yyVAL.node = AN(OCALL, yyS[yypt-3].node, yyS[yypt-1].node);
				//$$->builtin = 1;
		}
	case 48:
		yyVAL.node = yyS[yypt-0].node
	case 49:
		//line y.y:235
		{
				//fmt.Printf("Tconst %d\n", $1)
			yyVAL.node = Const(yyS[yypt-0].ival);
			}
	case 50:
		//line y.y:240
		{
				yyVAL.node = AN(OCONST, nil, nil);
				yyVAL.node.Type = TFLOAT;
				//$$->fmt = 'f';
			yyVAL.node.fval = yyS[yypt-0].fval;
			}
	case 51:
		//line y.y:247
		{
				yyVAL.node = AN(OCONST, nil, nil);
				yyVAL.node.Type = TSTRING;
				yyVAL.node.sval = yyS[yypt-0].sval;
				//$$->fmt = 's';
		}
	case 52:
		//line y.y:256
		{
				yyVAL.node = AN(ONAME, nil, nil);
				//$$->sym = $1;
		}
	case 53:
		//line y.y:261
		{
				yyVAL.node = AN(OFRAME, yyS[yypt-0].node, nil);
				//$$->sym = $1;
		}
	case 54:
		yyVAL.node = yyS[yypt-0].node
	case 55:
		//line y.y:269
		{
				yyVAL.node = AN(OLIST, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	}
	goto yystack /* stack new state and value */
}
