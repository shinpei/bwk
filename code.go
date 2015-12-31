package main

type Stack struct {
	values []Value
	sp     int // stack pointer
}

const (
	DEFAULT_STACK_SIZE = 1
)

func NewStack(size int) *Stack{
	stack := new(Stack)
	stack.sp = 0
	if size==0 {
		size = DEFAULT_STACK_SIZE
	}
	stack.values = make([]Value, size)
	return stack
}

func (s *Stack) Push(v Value) {
	//FIXME: append is a bad way to expand stack
	s.values = append(s.values, v)
	s.sp++
}

func (s *Stack) Pop() Value{
	val := s.values[s.sp]
	s.sp--
	return val
}

type Environment struct{
	stack *Stack
}

func NewEnvironment () *Environment {
	env := new(Environment)
	env.stack = NewStack(0)
	return env
}


type Code interface{
	Step(e *Environment)
}


// for example,
// popi
// addi 1
//
type (
	Pushi struct {x int}
	Popi struct {x int}
	Addi struct { x int}
	Subi struct {x int}
	Muli struct {x int}
	Divi struct {x int}
)

func (c *Pushi) Step (e *Environment) {
	// push x to the stack
	e.stack.Push(&IValue{Val: c.x})
}

func (c *Popi) Step (e *Environment) {
	// push x to the stack
	e.stack.Pop()
}

func (c *Addi) Step (e *Environment) {
	// add integer to the stack
	v := e.stack.Pop()
	ival, ok := v.(*IValue)
	if ok {
		ival.Set(ival.Get() + c.x)
	}
	e.stack.Push(ival)
}

func (c *Subi) Step (e *Environment) {
	// add integer to the stack
	v := e.stack.Pop()
	ival, ok := v.(*IValue)
	if ok {
		ival.Set(ival.Get() - c.x)
	}
	e.stack.Push(ival)
}

func (c *Muli) Step (e *Environment) {
	// add integer to the stack
	v := e.stack.Pop()
	ival, ok := v.(*IValue)
	if ok {
		ival.Set(ival.Get() * c.x)
	}
	e.stack.Push(ival)
}
func (c *Divi) Step (e *Environment) {
	// add integer to the stack
	v := e.stack.Pop()
	ival, ok := v.(*IValue)
	if ok {
		// TODO: c.x shouldn't be 0
		ival.Set(ival.Get() / c.x)
	}
	e.stack.Push(ival)
}