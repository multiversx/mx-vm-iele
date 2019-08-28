// File provided by the K Framework Go backend. Timestamp: 2019-08-28 22:25:14.706

// This is the source for the generation of the kore parser.
// To build it:
// goyacc -o koreparser.go -p "kore" koreparser.y (produces koreparser.go)
// go build
//

%{

package koreparser

%}

%union {
	str []byte
	k K
	kseq KSequence
	klist []K
}

%type	<k>	k kitem
%type	<klist>	klist
%type	<kseq>	ksequence

%token KSEQ DOTK '(' ')' ',' DOTKLIST TOKENLABEL KLABELLABEL KVARIABLE KLABEL STRING

%token	<str> KLABEL STRING KVARIABLE

%%

top:
	k
	{
		lastResult = $1
	}

ksequence:
	kitem KSEQ kitem
	{
		$$ = KSequence([]K{$1, $3})
	}
|	ksequence KSEQ kitem
	{
		$$ = append($1, $3)
	}
|	DOTK
	{
		$$ = nil
	}

k:
	kitem
	{
		$$ = $1
	}
|	ksequence
	{
		$$ = $1
	}

kitem:
	KLABEL '(' klist ')'
	{
		$$ = KApply { Label:string($1), List:$3 }
	}
|   KLABELLABEL '(' KLABEL ')'
    {
        $$ = InjectedKLabel { Label:string($3) }
    }
| 	TOKENLABEL '(' STRING ',' STRING ')'
	{
		$$ = KToken { Value: string($3), Sort: string($5) }
	}
|   KVARIABLE
    {
        $$ = KVariable { Name: string($1) }
    }

klist:
	k
	{
		$$ = []K{$1}
	}
|	klist ',' k
	{
		$$ = append($1, $3)
	}
|	klist ',' ',' k
	{
		$$ = append($1, $4)
	}
|	DOTKLIST
	{
		$$ = nil
	}

%%

var lastResult K

func Parse(kast []byte) K {
	lexer := koreLexerImpl{line: kast}
	koreParse(&lexer)
	return lastResult
}

