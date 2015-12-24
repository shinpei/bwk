package main

type Parser struct {


}
type NodeType int

const (
	IF_STMT NodeType = iota
	FOR_STMT
	DO_STMT
	WHILE_STMT
	COMPOUND_STMT
	EXPR_STMT
)

type Node struct {
	left *Node
	right *Node
	ntype NodeType
}

func (node *Node) String() string{

return ""
}

func (p *Parser) Parse(tokens *[]TokenType) *Node {
	node := new(Node)
	return node
}

func NewParser() *Parser {
	p := new(Parser);
	return p
}

