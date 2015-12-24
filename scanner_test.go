package main
import "testing"


func TestScan(t *testing.T) {
	eh := func (msg string) {
		t.Errorf("Error handler called (msg=%s)", msg)
	}

	source :="hogehoge"
	var s Scanner
	s.Init(source, eh);
	for {
		tok, lit = s.Scan()

	}
}