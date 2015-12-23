package main

type Tokenizer struct {

}

// better to define errors

func NewTokenizer() *Tokenizer{
	tok := new(Tokenizer)
	return tok
}


func (t *Tokenizer) Tokenize () error{
	return nil;
}
