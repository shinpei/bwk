package main
import (
	"strconv"
)

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


func (v *IValue) String() string {
	return strconv.Itoa(v.Val)
}
func (v *SValue) String() string {return v.Val}
func (v *FValue) String() string {return "0.0";}


func (v *IValue) Set(x int) {
	v.Val = x
}

func (v *IValue) Get() int {
	return v.Val
}