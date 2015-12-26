package main
import "testing"


func TestScan(t *testing.T) {
	eh := func (msg string) {
		t.Errorf("Error handler called (msg=%s)", msg)
	}

	source :=[]byte("123  symbol \"string\"\n")
	var s Scanner
	s.Init(source, eh);
	tok, _ := s.Scan()
	if tok != INT {
		t.Errorf("expected INT token is not found at 1st place");
	}
	tok, _ = s.Scan()
	if tok != SYMBOL {
		t.Errorf("expected SYMBOL token is not found at 2nd place")
	}
	tok, _ = s.Scan()
	if tok != STRING {
		t.Errorf("expected STRING token is not found at 2nd place")
	}

}