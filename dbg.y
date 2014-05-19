%{

package expr

import (
	"fmt"
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
%token <ival> Tconst Tfmt
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
		| Tfn Tid '(' args ')' zsemi '{' slist '}'
		{
			$2.Proc = AN(OLIST, $4, $8)
		}
		| Tfn Tid
		{
			$2.Proc = nil
		}
		;

zsemi		:
		| ';' zsemi
		;

zname		:
		{ $$ = nil; }
		| Tid
		;

slist		: stmnt
		{
			$$ = $1
		}
		| slist stmnt
		{
			$$ = AN(OLIST, $1, $2)
		}
		;

stmnt		: zexpr ';'
		{
			$$ = $1
		}
		| '{' slist '}'
		{
			$$ = $2
		}
		| Tif expr Tthen stmnt
		{
			$$ = AN(OIF, $2, $4);
		}
		| Tif expr Tthen stmnt Telse stmnt
		{
			$$ = AN(OIF, $2, AN(OELSE, $4, $6));
		}
		| Tloop expr ',' expr Tdo stmnt
		{
			$$ = AN(ODO, AN(OLIST, $2, $4), $6);
		}
		| Twhile expr Tdo stmnt
		{
			$$ = AN(OWHILE, $2, $4);
		}
		| Tret expr ';'
		{
			$$ = AN(ORET, $2, nil);
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
		| expr Tfmt
		{
			$$ = AN(OFMT, $1, Const($2));
		}
		;

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

