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

func TestScanInt(t *testing.T) {
	source := []byte("123 ")
	var s Scanner
	s.Init(source, eh)
	tok, _ := s.Scan()
	if tok != INT {
		t.Errorf("expected INT token is not found at 1st place")
	}
}

func TestScanFloatWithoutLeadingZero(t *testing.T) {
	source := []byte(".3")
	var s Scanner
	s.Init(source, eh)
	tok, _ := s.Scan()
	if tok != FLOAT {
		t.Errorf("expected FLOAT token is not found at 1st place")
	}
}
func TestScanFloat(t *testing.T) {
	source := []byte(".3")
	var s Scanner
	s.Init(source, eh)
	tok, _ := s.Scan()
	if tok != FLOAT {
		t.Errorf("expected FLOAT token is not found at 1st place")
	}
}

func TestScanSymbol(t *testing.T) {
	source := []byte("symbol")
	var s Scanner
	s.Init(source, eh)
	tok, _ := s.Scan()
	if tok != SYMBOL {
		t.Errorf("expected SYMBOL token is not found at 2nd place")
	}
}

func TestScanDollarSymbol(t *testing.T) {
	source := []byte("$alsoSymbol")
	var s Scanner
	s.Init(source, eh)
	tok, _ := s.Scan()
	if tok != SYMBOL {
		t.Errorf("expected SYMBOL token is not found at 2nd place")
	}
}

func TestScanSpecialSymbol(t *testing.T) {
	source := []byte("$1")
	var s Scanner
	s.Init(source, eh)
	tok, lit := s.Scan()
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
	tok, _ := s.Scan()
	if tok != STRING {
		t.Errorf("expected STRING token is not found at 2nd place")
	}
}

func TestScanKeywordExit(t *testing.T) {
	src := []byte("exit \"exit\"");
	var s Scanner
	s.Init(src, eh)
	tok, _ := s.Scan()
	if tok != EXIT{
		t.Errorf("sould be exit keyword, but it's ", tok.String())
	}
}
