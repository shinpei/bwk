package main
import "testing"

var (
tests = [...]string{
"{2+3}",
}
)

func TestParsing (t *testing.T) {
	var parser Parser
	src := []byte("{2+5+4}")
	parser.Init(src)
	parser.Parse()
}