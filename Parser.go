package main

type Parser struct {


}

type Node struct {

}

func (node *Node) String() string{
return ""
}

func (p *Parser) Parse(tokens *[]Token) *Node {
	node := new(Node)
	return node
}

func NewParser() *Parser {
	p := new(Parser);
	return p
}

