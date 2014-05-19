package expr

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

const Eof = 0

var (
	keywords = map[string]int{
		"do":      Tdo,
		"if":      Tif,
		"then":    Tthen,
		"else":    Telse,
		"while":   Twhile,
		"loop":    Tloop,
		"head":    Thead,
		"tail":    Ttail,
		"append":  Tappend,
		"defn":    Tfn,
		"return":  Tret,
		"local":   Tlocal,
		"aggr":    Tcomplex,
		"union":   Tcomplex,
		"adt":     Tcomplex,
		"complex": Tcomplex,
		"delete":  Tdelete,
		"whatis":  Twhat,
		"eval":    Teval,
		"builtin": Tbuiltin,
	}

	cmap = map[rune]rune{
		'0':  '\x00',
		'n':  '\n',
		'r':  '\r',
		't':  '\t',
		'b':  '\b',
		'f':  '\f',
		'a':  '\a',
		'v':  '\v',
		'\\': '\\',
		'"':  '"',
	}

	syms = map[string]*Lsym{}
)

func init() {
	for n, t := range keywords {
		enter(n, t)
	}
}

func look(name string) *Lsym {
	return syms[name]
}

func enter(name string, t int) *Lsym {
	l := &Lsym{
		Name:   name,
		Lexval: t,
	}

	v := &Value{
		Type: TINT,
	}

	l.V = v

	syms[name] = l

	return l
}

func mkvar(name string) *Lsym {
	l := look(name)
	if l == nil {
		l = enter(name, Tid)
	}
	return l
}

type yyLex struct {
	node    *Node
	line    []byte
	peek    rune
	stacked int
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

		case '{':
			l.stacked++
			return int(c)
		case '}':
			l.stacked--
			return int(c)

		case '\\':
			c = l.next()
			/*
				if(strchr(vfmt, c) == 0) {
					unlexc(c);
					return '\\';
				}
			*/
			lval.ival = int64(c)
			return Tfmt

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
				lval.ival = int64(l.escchar(l.next()))
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

func (l *yyLex) escchar(c rune) rune {
	buf := new(bytes.Buffer)

	if c >= '0' && c <= '9' {
		buf.WriteRune(c)
		for {
			c = l.next()
			if c == Eof {
				panic("eof in escape sequence")
			}
			if !strings.ContainsRune("0123456789xX", c) {
				l.peek = c
				break
			}
			buf.WriteRune(c)
		}
		esc, err := strconv.ParseInt(buf.String(), 0, 32)
		if err != nil {
			panic(fmt.Sprintf("escchar: %s", err))
		}
		return rune(esc)
	}

	n := cmap[c]
	if n == 0 {
		return c
	}
	return n
}

func (l *yyLex) Error(s string) {
	fmt.Printf("%s at: %s\n", s, l.line)
}

func (l *yyLex) string(lval *yySymType) int {
	buf := new(bytes.Buffer)
loop:
	for {
		c := l.next()
		switch c {
		case Eof:
			panic("eof in string")
		case '"':
			break loop
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

func (l *yyLex) numsym(c rune, lval *yySymType) int {
	var isbin, isfloat, ishex bool
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
			sel = "01234567890.xb"
			if ishex {
				sel = "01234567890abcdefABCDEF"
			} else if isbin {
				sel = "01"
			} else if isfloat {
				//sel = "01234567890eE-+";
				sel = "01234567890eE"
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
			buf.Next(2)
			ival, err := strconv.ParseInt(buf.String(), 2, 64)
			if err != nil {
				panic("bad binary " + buf.String() + " : " + err.Error())
			}
			lval.ival = ival
		} else {
			ival, err := strconv.ParseInt(buf.String(), 0, 64)
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

	s := look(buf.String())
	if s == nil {
		s = enter(buf.String(), Tid)
	}

	lval.sym = s
	return s.Lexval
}

func (l *yyLex) Execrec(n *Node) {
	s := mkvar("_thiscmd")
	v := &Value{
		Type:  TCODE,
		Store: Store{cc: n},
		pop:   s.V,
	}

	s.V = v
	s.Proc = n

	l.Execute(n)

	s.Proc = s.V.cc
	s.V = v.pop
}

func (yy *yyLex) Execute(n *Node) {
	var res *Node

	//l := n.Left
	//r := n.Right

	switch n.Op {
	default:
		res = expr(n)
		/*
			if(ret || (res.type == TLIST && res.l == 0 && n->op != OADD))
				break;
			prnt->right = &res;
			expr(prnt, &xx);
		*/

	case OASGN:
		fallthrough
	case OCALL:
		res = expr(n)
	}
	fmt.Println(res)
}
