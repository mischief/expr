%{

package expr

import (
	"math"
	"fmt"
	"unicode"
	"unicode/utf8"
)

var (
	stack nstack
)

%}

%union {
	n Node
	v float64 /* number */
	i string /* ident */
}

%token IDENT NUMBER

%type <n> node
%type <i> IDENT
%type <v> NUMBER

%right '='
%left '+' '-'
%left '*' '/' '%'
%left '^'
%right ','
%left UMINUS

%start list

%%

list:	/* eat empty */
|	list statement
;

statement:
	node {
		//stack.Push($1)
	}
|	node ',' statement {
		//stack.Push($1)
		//stack.Push($3)
	}
;

node:
	'(' node ')' {
		$$ = $2
		stack.Pop()
		stack.Push($$)
	}
|	node '+' node {
		$$ = Expr(EAdd, $1, $3)
		stack.Pop()
		stack.Pop()
		stack.Push($$)
	}
|	node '-' node {
		$$ = Expr(ESub, $1, $3)
		stack.Pop()
		stack.Pop()
		stack.Push($$)
	}
|	node '*' node {
		$$ = Expr(EMul, $1, $3)
		stack.Pop()
		stack.Pop()
		stack.Push($$)
	}
|	node '/' node {
		$$ = Expr(EDiv, $1, $3)
		stack.Pop()
		stack.Pop()
		stack.Push($$)
	}
|	node '%' node {
		$$ = Expr(EMod, $1, $3)
		stack.Pop()
		stack.Pop()
		stack.Push($$)
	}
|	node '^' node {
		$$ = Expr(EPow, $1, $3)
		stack.Pop()
		stack.Pop()
		stack.Push($$)
	}
|	'-' node	%prec UMINUS {
		$$ = Expr(ENeg, $2, Number(0))
		stack.Pop()
		stack.Push($$)
	}
|	IDENT {
		$$ = Ident(IGet, $1, nil)
		stack.Push($$)
	}
|	IDENT '(' node ')' {
		$$ = Ident(ICall, $1, $3)
		stack.Pop()
		stack.Push($$)
	}
|	IDENT '=' node {
		$$ = Ident(ISet, $1, $3)
		stack.Pop()
		stack.Push($$)
	}
|	NUMBER {
		$$ = Number($1)
		stack.Push($$)
	}
;

%%

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

