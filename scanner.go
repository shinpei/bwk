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
	ErrorCount int
}
type ErrorHandler func(msg string)


// Read a single character
// simplified go/token/scanner.go
func (s *Scanner) next(){
	if s.rdOffset < len(s.src) {
		s.offset = s.rdOffset;
		if s.ch == '\n' {
			s.lineOffset = s.offset;
		}
		r,w := rune(s.src[s.rdOffset]), 1
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


func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch >= 0x80 && unicode.IsLetter(ch)
}

func (s *Scanner) error (offs int, msg string) {
	if s.err != nil {
		s.err(msg)
	}
	s.ErrorCount++
}

func (s *Scanner) Scan() (tok TokenType, lit string) {
	//scanAgain:
	s.skipWhitespace()

	switch ch:= s.ch; {
	case isLetter(ch):
		//lit = s.scanIdentifier()
		if len(lit) > 1 {
			// keywords are longer than one letter.
			tok = Lookup(lit)
			switch tok {
			case SYMBOL, BREAK, CONTINUE, RETURN:

			}
		}else {
			tok = SYMBOL
		}
	}
	return
}