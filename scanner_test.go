package main

import (
	"testing"
	"fmt"
)

var (
	eh = func(msg string) {
		fmt.Errorf("Error handler called (msg=%s)", msg)
	}
)

func expect(t *testing.T, s Scanner, epos Pos, etk TokenType, elit string){
	pos,tok,lit := s.Scan()
	if tok != etk {
		t.Errorf("Expected %v token, but %v", etk.String(), tok.String())
	}
	if pos.Offset != epos.Offset {
		t.Errorf("Expected %v pos, but %v", epos.Offset, pos.Offset)
	}
	if lit != elit {
		t.Errorf("Expected '%s' lit , but '%s'", elit, lit)
	}
}

func TestScanInt(t *testing.T) {
	source := []byte("123 ")
	var s Scanner
	s.Init(source, eh)
	expect(t, s, Pos{Offset:0}, INT, "123")
}

func TestScanInts(t *testing.T) {
	src := []byte("123, 456")
	var s Scanner
	s.Init(src, eh)
	expect(t, s, Pos{Offset:0}, INT, "123")
	expect(t, s, Pos{Offset:4}, COMMA, "")
}

func TestScanFloatWithoutLeadingZero(t *testing.T) {
	source := []byte(".3")
	var s Scanner
	s.Init(source, eh)
	expect(t, s, Pos{Offset:0}, FLOAT, ".3")
}

func TestScanSymbol(t *testing.T) {
	source := []byte("symbol")
	var s Scanner
	s.Init(source, eh)
	expect(t, s, Pos{Offset:0}, SYMBOL, "symbol")

}

func TestScanSpecialSymbol(t *testing.T) {
	source := []byte("$1")
	var s Scanner
	s.Init(source, eh)
	expect(t, s, Pos{Offset:0}, SYMBOL, "$1")

}


func TestScanString(t *testing.T) {
	source := []byte("\"string\"")
	var s Scanner
	s.Init(source, eh)
	expect(t, s, Pos{Offset:0}, STRING, "\"string\"")

}

func TestScanKeywordExit(t *testing.T) {
	src := []byte("exit \"exit\"");
	var s Scanner
	s.Init(src, eh)
	_,tok, _ := s.Scan()
	if tok != EXIT{
		t.Errorf("sould be exit keyword, but it's ", tok.String())
	}
}
