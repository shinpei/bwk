package main
import "testing"

func TestCore(t *testing.T) {
	core := NewCore()
	core.EvaluateString(new(Config), "{1+1}")
}
