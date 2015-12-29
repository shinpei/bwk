package main

type Value interface {
	String()    string
}

type IValue struct {
	Val int
}

type SValue struct {
	Val string
}

type FValue struct {
	Val float64
}


func (v *IValue) String() string {return string(v.Val)}
func (v *SValue) String() string {return v.Val}
func (v *FValue) String() string {return "hi";}


func (v *IValue) Set(x int) {
	v.Val = x
}

func (v *IValue) Get() int {
	return v.Val
}