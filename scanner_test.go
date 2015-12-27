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
	if pos != epos {
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

func TestScanFloat(t *testing.T) {
	source := []byte(".3")
	var s Scanner
	s.Init(source, eh)
	_,tok, _ := s.Scan()
	if tok != FLOAT {
		t.Errorf("expected FLOAT token is not found at 1st place")
	}
}

func TestScanSymbol(t *testing.T) {
	source := []byte("symbol")
	var s Scanner
	s.Init(source, eh)
	_, tok, _ := s.Scan()
	if tok != SYMBOL {
		t.Errorf("expected SYMBOL token is not found at 2nd place")
	}
}

func TestScanDollarSymbol(t *testing.T) {
	source := []byte("$alsoSymbol")
	var s Scanner
	s.Init(source, eh)
	_, tok, _ := s.Scan()
	if tok != SYMBOL {
		t.Errorf("expected SYMBOL token is not found at 2nd place")
	}
}

func TestScanSpecialSymbol(t *testing.T) {
	source := []byte("$1")
	var s Scanner
	s.Init(source, eh)
	_,tok, lit := s.Scan()
	if tok != SYMBOL {
		t.Errorf("expected SYMBOL token is not found at 2nd place")
	}
	if lit != "$1" {
		t.Errorf("expected SYMBOL literal is $1 but '" + lit + "'")
	}
}


func TestScanString(t *testing.T) {
	source := []byte("\"string\"")
	var s Scanner
	s.Init(source, eh)
	_,tok, _ := s.Scan()
	if tok != STRING {
		t.Errorf("expected STRING token is not found at 2nd place")
	}
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
