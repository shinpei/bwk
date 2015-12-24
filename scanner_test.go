package main
import "testing"


func TestScan(t *testing.T) {
	eh := func (msg string) {
		t.Errorf("Error handler called (msg=%s)", msg)
	}

	source :=[]byte("123  222\n")
	var s Scanner
	s.Init(source, eh);
	for {
		tok, _ := s.Scan()
		print (tok)
	}
}