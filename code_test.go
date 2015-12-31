package main
import "testing"


func ExecCode(codes []Code, env *Environment) Value{

	for _, code := range codes {
		code.Step(env)
	}

	return env.stack.Pop()
}


func TestAddi (t *testing.T) {
	env := NewEnvironment()
	var codes []Code = make([]Code, 2)
	pushi := &Pushi{x:5}
	addi := &Addi{x:3}
	codes[0] = pushi
	codes[1] = addi
	stackTop := ExecCode(codes, env)
	if stackTop.String() != "8" {
		t.Errorf("Actual result '%s' is not same as expected, '%s'.", stackTop.String(), "8")
	}
}
func TestSubi (t *testing.T) {
	env := NewEnvironment()
	var codes []Code = make([]Code, 2)
	pushi := &Pushi{x:5}
	subi := &Subi{x:3}
	codes[0] = pushi
	codes[1] = subi
	stackTop := ExecCode(codes, env)
	if stackTop.String() != "2" {
		t.Errorf("Actual result '%s' is not same as expected, '%s'.", stackTop.String(), "2")
	}
}
func TestMuli (t *testing.T) {
	env := NewEnvironment()
	var codes []Code = make([]Code, 2)
	pushi := &Pushi{x:5}
	muli := &Muli{x:3}
	codes[0] = pushi
	codes[1] = muli
	stackTop := ExecCode(codes, env)
	if stackTop.String() != "15" {
		t.Errorf("Actual result '%s' is not same as expected, '%s'.", stackTop.String(), "15")
	}
}
func TestDivi (t *testing.T) {
	env := NewEnvironment()
	var codes []Code = make([]Code, 2)
	pushi := &Pushi{x:15}
	divi := &Divi{x:3}
	codes[0] = pushi
	codes[1] = divi
	stackTop := ExecCode(codes, env)
	if stackTop.String() != "5" {
		t.Errorf("Actual result '%s' is not same as expected, '%s'.", stackTop.String(), "5")
	}
}


