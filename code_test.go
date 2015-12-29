package main
import "testing"


func ExecCode(codes []Code, env *Environment) {

	for _, code := range codes {
		code.Step(env)
	}
	println(env.stack.Pop().String())
}

func TestAddi (t *testing.T) {
	env := NewEnvironment()
	var codes []Code = make([]Code, 2)
	pushi := &Pushi{x:5}
	addi := &Addi{x:3}
	codes[0] = pushi
	codes[1] = addi
	ExecCode(codes, env)
}
