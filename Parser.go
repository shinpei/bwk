package main

type Parser struct {


}

type AST struct {

}

func (ast *AST) String() string{
return ""
}

func (p *Parser) Parse(tokens *[]Token) *AST{
	ast := new(AST)
	return ast
}

func NewParser() *Parser {
	p := new(Parser);
	return p
}

