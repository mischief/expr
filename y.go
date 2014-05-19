
//line dbg.y:2

package expr
import __yyfmt__ "fmt"
//line dbg.y:3
		
import (
	//"fmt"
)


//line dbg.y:11
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

//line dbg.y:310



//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 19,
}

const yyNprod = 66
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 533

var yyAct = []int{

	36, 35, 3, 123, 13, 11, 105, 5, 75, 27,
	103, 102, 127, 37, 38, 39, 40, 32, 119, 33,
	61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
	71, 122, 72, 107, 116, 76, 77, 118, 74, 115,
	116, 73, 116, 41, 42, 43, 124, 83, 84, 85,
	86, 87, 88, 89, 90, 91, 92, 93, 94, 95,
	96, 97, 98, 99, 100, 101, 54, 52, 53, 48,
	49, 50, 51, 47, 46, 44, 45, 41, 42, 43,
	110, 106, 112, 109, 106, 108, 111, 59, 60, 58,
	57, 56, 55, 54, 52, 53, 48, 49, 50, 51,
	47, 46, 44, 45, 41, 42, 43, 113, 114, 44,
	45, 41, 42, 43, 117, 106, 34, 2, 1, 12,
	125, 126, 14, 0, 121, 0, 15, 16, 128, 129,
	77, 17, 18, 0, 104, 0, 26, 32, 30, 29,
	31, 7, 0, 0, 0, 9, 8, 19, 20, 21,
	0, 10, 0, 0, 0, 22, 25, 28, 0, 6,
	130, 0, 23, 24, 15, 16, 0, 0, 0, 17,
	18, 0, 0, 0, 26, 32, 30, 29, 31, 7,
	0, 0, 0, 9, 8, 19, 20, 21, 0, 10,
	0, 0, 0, 22, 25, 28, 0, 6, 78, 0,
	23, 24, 15, 16, 0, 0, 0, 17, 18, 0,
	0, 0, 26, 32, 30, 29, 31, 7, 0, 0,
	0, 9, 8, 19, 20, 21, 4, 10, 0, 0,
	0, 22, 25, 28, 0, 6, 15, 16, 23, 24,
	0, 17, 18, 0, 0, 0, 26, 32, 30, 29,
	31, 7, 0, 0, 0, 9, 8, 19, 20, 21,
	0, 10, 0, 0, 0, 22, 25, 28, 0, 6,
	0, 0, 23, 24, 59, 60, 58, 57, 56, 55,
	54, 52, 53, 48, 49, 50, 51, 47, 46, 44,
	45, 41, 42, 43, 0, 0, 15, 16, 0, 0,
	0, 17, 18, 0, 0, 0, 26, 32, 30, 29,
	31, 0, 0, 0, 0, 0, 0, 19, 20, 21,
	0, 0, 0, 0, 80, 22, 25, 28, 0, 0,
	0, 0, 23, 24, 59, 60, 58, 57, 56, 55,
	54, 52, 53, 48, 49, 50, 51, 47, 46, 44,
	45, 41, 42, 43, 48, 49, 50, 51, 47, 46,
	44, 45, 41, 42, 43, 0, 79, 59, 60, 58,
	57, 56, 55, 54, 52, 53, 48, 49, 50, 51,
	47, 46, 44, 45, 41, 42, 43, 47, 46, 44,
	45, 41, 42, 43, 0, 0, 0, 0, 120, 59,
	60, 58, 57, 56, 55, 54, 52, 53, 48, 49,
	50, 51, 47, 46, 44, 45, 41, 42, 43, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	81, 82, 59, 60, 58, 57, 56, 55, 54, 52,
	53, 48, 49, 50, 51, 47, 46, 44, 45, 41,
	42, 43, 59, 60, 58, 57, 56, 55, 54, 52,
	53, 48, 49, 50, 51, 47, 46, 44, 45, 41,
	42, 43, 57, 56, 55, 54, 52, 53, 48, 49,
	50, 51, 47, 46, 44, 45, 41, 42, 43, 56,
	55, 54, 52, 53, 48, 49, 50, 51, 47, 46,
	44, 45, 41, 42, 43, 55, 54, 52, 53, 48,
	49, 50, 51, 47, 46, 44, 45, 41, 42, 43,
	52, 53, 48, 49, 50, 51, 47, 46, 44, 45,
	41, 42, 43,
}
var yyPact = []int{

	-1000, 182, -1000, -1000, -12, 112, 216, 276, 276, 276,
	276, 447, -1000, -1000, -1000, 276, 276, 276, 276, 276,
	276, 276, 276, 276, 276, 276, 276, 11, -14, -1000,
	-1000, -1000, -50, 5, -1000, 144, -1000, 329, 269, 394,
	427, 276, 276, 276, 276, 276, 276, 276, 276, 276,
	276, 276, 276, 276, 276, 276, 276, 276, 276, 276,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -44, -45, -1000,
	-1000, -1000, 82, 276, 3, -14, 276, -1000, -1000, 216,
	276, 216, -1000, -1000, -1000, -1000, 21, 21, 89, 89,
	369, 369, 369, 369, 340, 340, 508, 55, 495, 480,
	464, 447, 276, 276, -1000, -13, -1000, 276, -1000, -15,
	-20, 362, -1000, -1000, -1000, -1000, 276, -21, 42, 216,
	216, -1000, -1000, -41, 42, -1000, -1000, 216, -1000, 106,
	-1000,
}
var yyPgo = []int{

	0, 5, 4, 122, 0, 9, 6, 7, 1, 119,
	119, 118, 117, 3,
}
var yyR1 = []int{

	0, 11, 11, 12, 12, 12, 13, 13, 10, 10,
	8, 8, 4, 4, 4, 4, 4, 4, 4, 7,
	7, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 9, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 3, 3, 3, 3, 3,
	3, 3, 5, 5, 6, 6,
}
var yyR2 = []int{

	0, 0, 2, 1, 9, 2, 0, 2, 0, 1,
	1, 2, 2, 3, 4, 6, 6, 4, 3, 0,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 2, 1, 1, 2, 2, 2, 2, 2, 2,
	4, 4, 2, 2, 2, 3, 4, 5, 1, 1,
	1, 1, 1, 3, 1, 3,
}
var yyChk = []int{

	-1000, -11, -12, -4, 44, -7, 53, 35, 40, 39,
	45, -1, -9, -2, -3, 20, 21, 25, 26, 41,
	42, 43, 49, 56, 57, 50, 30, -5, 51, 33,
	32, 34, 31, 31, 4, -8, -4, -1, -1, -1,
	-1, 22, 23, 24, 20, 21, 19, 18, 14, 15,
	16, 17, 12, 13, 11, 10, 9, 8, 7, 5,
	6, -2, -2, -2, -2, -2, -2, -2, -2, -2,
	-2, -2, -1, 30, -5, 58, 30, -4, 54, 37,
	55, 36, 4, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, 55, 55, 52, -6, -7, 30, -5, -6,
	-4, -1, -4, -2, -2, 52, 55, -6, 52, 38,
	36, -7, 52, -13, 4, -4, -4, 53, -13, -8,
	54,
}
var yyDef = []int{

	1, -2, 2, 3, 0, 0, 19, 0, 0, 0,
	0, 20, 21, 42, 43, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 58, 0, 59,
	60, 61, 62, 5, 12, 19, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	41, 44, 45, 46, 47, 48, 49, 0, 0, 52,
	53, 54, 0, 19, 0, 0, 19, 11, 13, 19,
	0, 19, 18, 22, 23, 24, 25, 26, 27, 28,
	29, 30, 31, 32, 33, 34, 35, 36, 37, 38,
	39, 40, 0, 0, 55, 0, 64, 19, 63, 0,
	14, 0, 17, 50, 51, 56, 19, 0, 6, 19,
	19, 65, 57, 0, 6, 15, 16, 19, 7, 19,
	4,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 56, 3, 3, 3, 24, 11, 3,
	30, 52, 22, 20, 55, 21, 28, 23, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 58, 4,
	14, 5, 15, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 29, 3, 3, 10, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 53, 9, 54, 57,
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
		//line dbg.y:52
		{
				//Execrec($1)
			yylex.(*yyLex).node = yyS[yypt-0].node
			}
	case 4:
		//line dbg.y:57
		{
				yyS[yypt-7].sym.Proc = AN(OLIST, yyS[yypt-5].node, yyS[yypt-1].node)
			}
	case 5:
		//line dbg.y:61
		{
				yyS[yypt-0].sym.Proc = nil
			}
	case 8:
		//line dbg.y:71
		{ yyVAL.sym = nil; }
	case 9:
		yyVAL.sym = yyS[yypt-0].sym
	case 10:
		//line dbg.y:76
		{
				yyVAL.node = yyS[yypt-0].node
			}
	case 11:
		//line dbg.y:80
		{
				yyVAL.node = AN(OLIST, yyS[yypt-1].node, yyS[yypt-0].node)
			}
	case 12:
		//line dbg.y:86
		{
				yyVAL.node = yyS[yypt-1].node
			}
	case 13:
		//line dbg.y:90
		{
				yyVAL.node = yyS[yypt-1].node
			}
	case 14:
		//line dbg.y:94
		{
				yyVAL.node = AN(OIF, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 15:
		//line dbg.y:98
		{
				yyVAL.node = AN(OIF, yyS[yypt-4].node, AN(OELSE, yyS[yypt-2].node, yyS[yypt-0].node));
			}
	case 16:
		//line dbg.y:102
		{
				yyVAL.node = AN(ODO, AN(OLIST, yyS[yypt-4].node, yyS[yypt-2].node), yyS[yypt-0].node);
			}
	case 17:
		//line dbg.y:106
		{
				yyVAL.node = AN(OWHILE, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 18:
		//line dbg.y:110
		{
				yyVAL.node = AN(ORET, yyS[yypt-1].node, nil);
			}
	case 19:
		//line dbg.y:116
		{
				yyVAL.node = nil
			}
	case 20:
		yyVAL.node = yyS[yypt-0].node
	case 21:
		yyVAL.node = yyS[yypt-0].node
	case 22:
		//line dbg.y:124
		{
				yyVAL.node = AN(OMUL, yyS[yypt-2].node, yyS[yypt-0].node); 
			}
	case 23:
		//line dbg.y:128
		{
				yyVAL.node = AN(ODIV, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 24:
		//line dbg.y:132
		{
				yyVAL.node = AN(OMOD, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 25:
		//line dbg.y:136
		{
				yyVAL.node = AN(OADD, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 26:
		//line dbg.y:140
		{
				yyVAL.node = AN(OSUB, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 27:
		//line dbg.y:144
		{
				yyVAL.node = AN(ORSH, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 28:
		//line dbg.y:148
		{
				yyVAL.node = AN(OLSH, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 29:
		//line dbg.y:152
		{
				yyVAL.node = AN(OLT, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 30:
		//line dbg.y:156
		{
				yyVAL.node = AN(OGT, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 31:
		//line dbg.y:160
		{
				yyVAL.node = AN(OLEQ, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 32:
		//line dbg.y:164
		{
				yyVAL.node = AN(OGEQ, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 33:
		//line dbg.y:168
		{
				yyVAL.node = AN(OEQ, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 34:
		//line dbg.y:172
		{
				yyVAL.node = AN(ONEQ, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 35:
		//line dbg.y:176
		{
				yyVAL.node = AN(OLAND, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 36:
		//line dbg.y:180
		{
				yyVAL.node = AN(OXOR, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 37:
		//line dbg.y:184
		{
				yyVAL.node = AN(OLOR, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 38:
		//line dbg.y:188
		{
				yyVAL.node = AN(OCAND, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 39:
		//line dbg.y:192
		{
				yyVAL.node = AN(OCOR, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 40:
		//line dbg.y:196
		{
				yyVAL.node = AN(OASGN, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 41:
		//line dbg.y:200
		{
				yyVAL.node = AN(OFMT, yyS[yypt-1].node, Const(yyS[yypt-0].ival));
			}
	case 42:
		yyVAL.node = yyS[yypt-0].node
	case 43:
		yyVAL.node = yyS[yypt-0].node
	case 44:
		//line dbg.y:210
		{
				yyVAL.node = AN(OADD, yyS[yypt-0].node, nil);
			}
	case 45:
		//line dbg.y:214
		{
				yyVAL.node = Const(0);
				yyVAL.node = AN(OSUB, yyVAL.node, yyS[yypt-0].node);
			}
	case 46:
		//line dbg.y:219
		{
				yyVAL.node = AN(OEDEC, yyS[yypt-0].node, nil);
			}
	case 47:
		//line dbg.y:223
		{
				yyVAL.node = AN(OEINC, yyS[yypt-0].node, nil);
			}
	case 48:
		//line dbg.y:227
		{
				yyVAL.node = AN(OHEAD, yyS[yypt-0].node, nil);
			}
	case 49:
		//line dbg.y:231
		{
				yyVAL.node = AN(OTAIL, yyS[yypt-0].node, nil);
			}
	case 50:
		//line dbg.y:235
		{
				yyVAL.node = AN(OAPPEND, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 51:
		//line dbg.y:239
		{
				yyVAL.node = AN(ODELETE, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	case 52:
		//line dbg.y:243
		{
				yyVAL.node = AN(ONOT, yyS[yypt-0].node, nil);
			}
	case 53:
		//line dbg.y:247
		{
				yyVAL.node = AN(OXOR, yyS[yypt-0].node, Const(-1));
			}
	case 54:
		//line dbg.y:251
		{
				yyVAL.node = AN(OEVAL, yyS[yypt-0].node, nil);
			}
	case 55:
		//line dbg.y:257
		{
				yyVAL.node = yyS[yypt-1].node
			}
	case 56:
		//line dbg.y:261
		{
				yyVAL.node = AN(OCALL, yyS[yypt-3].node, yyS[yypt-1].node);
			}
	case 57:
		//line dbg.y:265
		{
				yyVAL.node = AN(OCALL, yyS[yypt-3].node, yyS[yypt-1].node);
				//$$->builtin = 1;
		}
	case 58:
		yyVAL.node = yyS[yypt-0].node
	case 59:
		//line dbg.y:271
		{
				//fmt.Printf("Tconst %d\n", $1)
			yyVAL.node = Const(yyS[yypt-0].ival);
			}
	case 60:
		//line dbg.y:276
		{
				yyVAL.node = AN(OCONST, nil, nil);
				yyVAL.node.Type = TFLOAT;
				//$$->fmt = 'f';
			yyVAL.node.fval = yyS[yypt-0].fval;
			}
	case 61:
		//line dbg.y:283
		{
				yyVAL.node = AN(OCONST, nil, nil);
				yyVAL.node.Type = TSTRING;
				yyVAL.node.sval = yyS[yypt-0].sval;
				//$$->fmt = 's';
		}
	case 62:
		//line dbg.y:292
		{
				yyVAL.node = AN(ONAME, nil, nil);
				//$$->sym = $1;
		}
	case 63:
		//line dbg.y:297
		{
				yyVAL.node = AN(OFRAME, yyS[yypt-0].node, nil);
				//$$->sym = $1;
		}
	case 64:
		yyVAL.node = yyS[yypt-0].node
	case 65:
		//line dbg.y:305
		{
				yyVAL.node = AN(OLIST, yyS[yypt-2].node, yyS[yypt-0].node);
			}
	}
	goto yystack /* stack new state and value */
}
