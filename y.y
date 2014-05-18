%{

package expr

import (
	"fmt"
	"bytes"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

%}

%union {
	node *Node
	sym  *Lsym
	fval float64
	ival int64
	sval string
}

%type <node> expr monexpr term stmnt name args zexpr slist
%type <node> castexpr
%type <sym> zname

%left	';'
%right	'='
%left	Tfmt
%left	Toror
%left	Tandand
%left	'|'
%left	'^'
%left	'&'
%left	Teq Tneq
%left	'<' '>' Tleq Tgeq
%left	Tlsh Trsh
%left	'+' '-'
%left	'*' '/' '%'
%right	Tdec Tinc Tindir '.' '[' '('

%token <sym> Tid
%token <fval> Tfconst
%token <ival> Tconst
%token <sval> Tstring
%token Tif Tdo Tthen Telse Twhile Tloop Thead Ttail Tappend Tfn Tret Tlocal
%token Tcomplex Twhat Tdelete Teval Tbuiltin

%%

prog		: 
		| prog bigstmnt
		;

bigstmnt	: stmnt
		{
			//Execrec($1)
			yylex.(*yyLex).node = $1
		}
		;

zname		: 
		{ $$ = nil; }
		| Tid
		;

slist		: stmnt
		| slist stmnt
		{
			$$ = AN(OLIST, $1, $2)
		}

stmnt		: zexpr ';'
		{
			$$ = $1
		}
		| '{' slist '}'
		{
			$$ = $2
		}
		;

zexpr		:
		{
			$$ = nil
		}
		| expr
		;

expr		: castexpr
		| expr '*' expr
		{
			$$ = AN(OMUL, $1, $3); 
		}
		| expr '/' expr
		{
			$$ = AN(ODIV, $1, $3);
		}
		| expr '%' expr
		{
			$$ = AN(OMOD, $1, $3);
		}
		| expr '+' expr
		{
			$$ = AN(OADD, $1, $3);
		}
		| expr '-' expr
		{
			$$ = AN(OSUB, $1, $3);
		}
		| expr Trsh expr
		{
			$$ = AN(ORSH, $1, $3);
		}
		| expr Tlsh expr
		{
			$$ = AN(OLSH, $1, $3);
		}
		| expr '<' expr
		{
			$$ = AN(OLT, $1, $3);
		}
		| expr '>' expr
		{
			$$ = AN(OGT, $1, $3);
		}
		| expr Tleq expr
		{
			$$ = AN(OLEQ, $1, $3);
		}
		| expr Tgeq expr
		{
			$$ = AN(OGEQ, $1, $3);
		}
		| expr Teq expr
		{
			$$ = AN(OEQ, $1, $3);
		}
		| expr Tneq expr
		{
			$$ = AN(ONEQ, $1, $3);
		}
		| expr '&' expr
		{
			$$ = AN(OLAND, $1, $3);
		}
		| expr '^' expr
		{
			$$ = AN(OXOR, $1, $3);
		}
		| expr '|' expr
		{
			$$ = AN(OLOR, $1, $3);
		}
		| expr Tandand expr
		{
			$$ = AN(OCAND, $1, $3);
		}
		| expr Toror expr
		{
			$$ = AN(OCOR, $1, $3);
		}
		| expr '=' expr
		{
			$$ = AN(OASGN, $1, $3);
		}

castexpr	: monexpr
		;

monexpr		: term
		| '+' monexpr
		{
			$$ = AN(OADD, $2, nil);
		}
		| '-' monexpr
		{
			$$ = Const(0);
			$$ = AN(OSUB, $$, $2);
		}
		| Tdec monexpr
		{
			$$ = AN(OEDEC, $2, nil);
		}
		| Tinc monexpr
		{
			$$ = AN(OEINC, $2, nil);
		}
		| Thead monexpr
		{
			$$ = AN(OHEAD, $2, nil);
		}
		| Ttail monexpr
		{
			$$ = AN(OTAIL, $2, nil);
		}
		| Tappend monexpr ',' monexpr
		{
			$$ = AN(OAPPEND, $2, $4);
		}
		| Tdelete monexpr ',' monexpr
		{
			$$ = AN(ODELETE, $2, $4);
		}
		| '!' monexpr
		{
			$$ = AN(ONOT, $2, nil);
		}
		| '~' monexpr
		{
			$$ = AN(OXOR, $2, Const(-1));
		}
		| Teval monexpr
		{
			$$ = AN(OEVAL, $2, nil);
		}
		;

term		: '(' expr ')'
		{
			$$ = $2
		}
		| name '(' args ')'
		{
			$$ = AN(OCALL, $1, $3);
		}
		| Tbuiltin name '(' args ')'
		{
			$$ = AN(OCALL, $2, $4);
			//$$->builtin = 1;
		}
		| name
		| Tconst
		{
			//fmt.Printf("Tconst %d\n", $1)
			$$ = Const($1);
		}
		| Tfconst
		{
			$$ = AN(OCONST, nil, nil);
			$$.Type = TFLOAT;
			//$$->fmt = 'f';
			$$.fval = $1;
		}
		| Tstring
		{
			$$ = AN(OCONST, nil, nil);
			$$.Type = TSTRING;
			$$.sval = $1;
			//$$->fmt = 's';
		}
		;

name		: Tid
		{
			$$ = AN(ONAME, nil, nil);
			//$$->sym = $1;
		}
		| Tid ':' name
		{
			$$ = AN(OFRAME, $3, nil);
			//$$->sym = $1;
		}
		;

args		: zexpr
		| args ','  zexpr
		{
			$$ = AN(OLIST, $1, $3);
		}
		;

%%

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
				//sel = "01234567890eE-+";
				sel = "01234567890eE";
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

