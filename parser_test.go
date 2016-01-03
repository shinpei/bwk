package main
import "testing"

func TestParsing (t *testing.T) {
	var parser Parser
	src := []byte("{5+4}")
	parser.Init(src)
	parser.Parse()
}