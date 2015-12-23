package main

type Tokenizer struct {
}

type Token int

const (
	ILLEGAL Token = iota
	EOF
	COMMENT


	literal_begin
	SYMBOL // foo
	INT    // 1
	FLOAT  // 0.112
	STRING // "abc"
	literal_end

	operator_begin
	ADD // '+'
	SUB // '-'
	MUL // '*'
	QUO // '/'
	REM // '%'

	EQL
	NEQ

	operator_end

// Special characters
	LPAREN // (
	LBRACE // {

	RPAREN // )
	RBRACE // }

	COLON // :
	SEMICOLON // ;

	// Keywords
	keyword_begin
	IF
	WHILE
	FOR
	DO
	BREAK
	CONTINUE
	PRINT
	PRINTF

	RETURN
	NEXT
	DELETE
	EXIT
	keyword_end
)

// better to define errors

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	COMMENT: "COMMENT",

	SYMBOL: "SYMBOL",
	INT:     "INT",
	FLOAT:   "FLOAT",
	STRING:  "STRING",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	QUO: "/",
	REM: "%",

	EQL: "==",
	NEQ: "!=",

	LPAREN: "(",
	LBRACE: "{",

	RPAREN: ")",
	RBRACE : "}",

	COLON: ":",
	SEMICOLON: ";",

	IF:       "if",
	WHILE:    "while",
	FOR:      "for",
	DO:       "do",
	BREAK:    "break",
	CONTINUE: "continue",
	PRINT:    "print",
	PRINTF:   "printf",

	RETURN: "return",
	NEXT:   "next",
	DELETE: "delete",
	EXIT:   "exit",
}

var keywords map[string]Token

func (tok Token) IsKeyword() bool {
	return keyword_begin < tok && tok <= keyword_end
}

func (tok Token) IsLiteral() bool {
	return literal_begin < tok && tok <= literal_end
}

func (tok Token) IsOperator() bool {
	return operator_begin < tok && tok <= operator_end
}

func NewTokenizer() *Tokenizer {
	tok := new(Tokenizer)
	keywords = make(map[string]Token)
	for i := keyword_begin + 1; i < keyword_end; i++ {
		keywords[tokens[i]] = i
	}
	return tok
}

// snip from golang
func Lookup (ident string) Token {
	if tok, is_keyword := keywords[ident]; is_keyword {
		return tok
	}
	return SYMBOL
}
func (t *Tokenizer) Tokenize(line string) (tokens *[]Token, err error) {
	tokens = new([]Token)

	return tokens, nil
}
