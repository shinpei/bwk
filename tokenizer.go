package main
import "fmt"

type Tokenizer struct {
}

type TokenNode struct{
	ttype TokenType
	body string
	// TODO: position for debugging
}
type TokenType int

const (
	ILLEGAL TokenType = iota
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
	DIV // '/'
	MOD // '%'

	EQL
	NEQ
	ASSIGN

	operator_end

// Special characters
	LPAREN // (
	LBRACE // {

	RPAREN // )
	RBRACE // }

	COLON // :
	SEMICOLON // ;

	COMMA
	DOT

	DOLLAR // $

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

var tokenDefs = [...]string{
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
	DIV: "/",
	MOD: "%",

	EQL: "==",
	NEQ: "!=",
	ASSIGN: "=",

	LPAREN: "(",
	LBRACE: "{",

	RPAREN: ")",
	RBRACE : "}",

	COLON: ":",
	SEMICOLON: ";",

	COMMA: ",",
	DOT: ".",

	DOLLAR : "$",

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

var keywords map[string]TokenType

func (tok TokenType) IsKeyword() bool {
	return keyword_begin < tok && tok <= keyword_end
}

func (tok TokenType) IsLiteral() bool {
	return literal_begin < tok && tok <= literal_end
}

func (tok TokenType) IsOperator() bool {
	return operator_begin < tok && tok <= operator_end
}
func (tok TokenType) String() string{
	return tokenDefs[tok]
}

const (
	LowestPrec  = 0 // non-operators
//	UnaryPrec   = 6
//	HighestPrec = 7
)

func (tok TokenType) Precedence() int {
	switch tok {
//	case LOR:
//		return 1
//	case LAND:
//		return 2
	case EQL, NEQ/*, LSS, LEQ, GTR, GEQ*/:
		return 3
	case ADD, SUB/*, OR, XOR*/:
		return 4
	case MUL/*, QUO, REM, SHL, SHR, AND, AND_NOT*/:
		return 5
	}
	return LowestPrec
}

func NewTokenizer() *Tokenizer {
	tok := new(Tokenizer)
	keywords = make(map[string]TokenType)
	for i := keyword_begin + 1; i < keyword_end; i++ {
		keywords[tokenDefs[i]] = i
	}
	return tok
}

// snip from golang
func Lookup (ident string) TokenType {
	if tok, is_keyword := keywords[ident]; is_keyword {
		return tok
	}
	return SYMBOL
}


func (t *Tokenizer) Tokenize(src string) (tokens []TokenType, err error) {
	tokens = make([]TokenType, 0)
	s := new(Scanner)
	eh := func (msg string) {
		fmt.Errorf("Error handler called (msg=%s)", msg)
	}
	srcByte := []byte(src)
	s.Init(srcByte, eh)
	for {
		_, tok, lit := s.Scan()
		if tok == EOF {
			break;
		}
		fmt.Println("[Token:", tokenDefs[tok], ":", lit, "]")
		tokens = append(tokens, tok)
	}
	return tokens, nil
}
