// Scanner for bwk.
// A lot inspired from go/scanner

package main
import (
	"unicode"
)



type Scanner struct {
	src []byte

	ch rune
	offset int
	err ErrorHandler
	rdOffset int // reading offset
	lineOffset int
	insertSemi bool // insert a semicolon before new line

	ErrorCount int
}
type ErrorHandler func(msg string)


// Read a single character
// simplified go/token/scanner.go
func (s *Scanner) next() {
	if s.rdOffset < len(s.src) {
		s.offset = s.rdOffset;
		if s.ch == '\n' {
			s.lineOffset = s.offset;
		}
		r, w := rune(s.src[s.rdOffset]), 1
		switch {
		case r == 0:
			s.error(s.offset, "illegal character NUL");
		case r >= 0x80:
			// won't handle
			s.error(s.offset, "multibyte char is not supported");
		}
		s.rdOffset += w
		s.ch =r
	} else{
		s.offset = len(s.src)
		if s.ch == '\n'{
			s.lineOffset = s.offset
		}
		s.ch = -1 // eof
	}
}
const bom = 0xFEFF // byte order mark, only permitted as very first character

func (s *Scanner ) Init (src []byte, err ErrorHandler) {
	s.src = src

	s.ch = ' '
	s.offset = 0
	s.rdOffset = 0
	s.lineOffset = 0
	s.err = err
	s.ErrorCount= 0
	s.next()
	if s.ch == bom {
		// ERROR. we don't support utf8
	}
}
func (s *Scanner) skipWhitespace() {
	for s.ch == ' ' || s.ch == '\t' || s.ch == '\n' /*&& !s.insertSemi*/ || s.ch == '\r' {
		s.next()
	}
}

func (s *Scanner) scanIdentifier() string {
	offs := s.offset
	for isLetter(s.ch) || isDigit(s.ch) {
		s.next()
	}
	return string(s.src[offs:s.offset])
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9' || ch >= 0x80 && unicode.IsDigit(ch)
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch >= 0x80 && unicode.IsLetter(ch)
}

func (s *Scanner) error (offs int, msg string) {
	if s.err != nil {
		s.err(msg)
	}
	s.ErrorCount++
}


func digitVal(ch rune) int {
	switch {
	case '0' <= ch && ch <= '9':
		return int(ch - '0')
	}
	return 10 // larger than any legal digit val
}


func (s *Scanner) scanNumber(seenDecimalPoint bool) (tok TokenType,  lit string) {

	offs := s.offset
	tok = INT// TODO: Parse float

	if seenDecimalPoint {
		offs--
		tok = FLOAT
		for digitVal(s.ch) < 10 {
			s.next()
		}
		return tok, string(s.src[offs: s.offset])
	}

	for digitVal(s.ch) < 10 {
		s.next()
	}
	return tok, string(s.src[offs: s.offset])

}

func (s *Scanner) scanString() string {
	// '"' opening already consumed
	offs := s.offset - 1

	for {
		ch := s.ch
		if ch == '\n' || ch < 0 {
			s.error(offs, "string literal not terminated")
			break
		}
		s.next()
		if ch == '"' {
			break
		}
		/*
		if ch == '\\' {
			s.scanEscape('"')
		}
		*/
	}

	return string(s.src[offs:s.offset])
}
func (s *Scanner) Scan() (pos Pos, tok TokenType, lit string) {
	//scanAgain:
	s.skipWhitespace()

	pos = Pos{Offset:s.offset}
	switch ch := s.ch; {
	case isLetter(ch):
		lit = s.scanIdentifier()
		if len(lit) > 1 {
			// keywords are longer than one letter.
			tok = Lookup(lit)
			switch tok {
			case SYMBOL, BREAK, CONTINUE, RETURN:

			}
		}else {
			tok = SYMBOL
		}
	case '0' <= ch && ch <= '9':
		tok, lit = s.scanNumber(false)
	default:
		s.next()
		switch ch  { // now, ch is a old one
		case -1:
			// TERMINATED
			if s.insertSemi{
				s.insertSemi = false
				return pos, SEMICOLON, "\n"
			}
			tok = EOF
		case '\n':
			s.insertSemi = false
			return pos, SEMICOLON, "\n"
		case '"':
			tok = STRING
			lit = s.scanString()
		case '(':
			tok = LPAREN
		case ')'  :
			tok = RPAREN
		case '{' :
			tok = LBRACE
		case '}':
			tok = RBRACE
		case '$':
			// special or symbol
			if isLetter(s.ch) {
				//
				lit = s.scanString()
				tok = SYMBOL
			}else if '0' <= s.ch && s.ch <= '9'{
				tok, lit = s.scanNumber(false)
				tok = SYMBOL
				lit = "$" + lit
			}
		case '+':
			tok = ADD
		case '-':
			tok = SUB
		case '*':
			tok = MUL
		case '/':
			tok = DIV
		case '%':
			tok = MOD
		case ',':
			tok = COMMA
		case '=':
			tok = ASSIGN
		case '.':
			if '0' <= s.ch && s.ch <= '9' {
				tok, lit = s.scanNumber(true)
			}else {
				tok = DOT
			}
		case ';':
			tok = SEMICOLON
		default:
			lit = string(ch)
		}
	}
	D("Tokenizer:", tok.String())
	return
}