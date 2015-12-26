package main
import (
)

type Parser struct {
	scanner Scanner

	tok TokenType
	pos Pos
	lit string
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

type ParsedThing struct {
	Decls []Decl
}

func (p *Parser) error (msg string) {
	panic(msg) // TODO: more smarter way is required.
}

func (p *Parser) expect(tok TokenType) Pos {
	pos := p.pos
	if p.tok != tok {
		p.error(tok.String() + " is expected but we got "+ p.tok.String())
	}
	p.next()
	return pos
}

/*
func (p *Parser) parseSimpleStmt() (Stmt, bool) {
	return
} */

func (p *Parser) parseStmt() (s Stmt) {
	switch p.tok {
	case SYMBOL, INT, FLOAT, STRING, LPAREN,
		ADD,SUB,MUL,DIV:
		//s, _ = p.parseSimpleStmt()
	}
	return s
}

func (p *Parser) parseStmtList()(list []Stmt) {
	for p.tok != RBRACE && p.tok != EOF {
		list = append(list, p.parseStmt())
	}
	return
}
func (p *Parser) parseBody(/*scope Scope */) *BlockStmt{
	lbrace:=p.expect(LBRACE)
	list := p.parseStmtList()
	rbrace := p.expect(RBRACE)

	return &BlockStmt{
		Lbrace: lbrace,
		List: list,
		Rbrace : rbrace,
	}
}

// it's like parseFuncDecl
// no global scope var defs
func (p *Parser) parseDecl () Decl {


/*	if p.tok == LPAREN {
		//

	}
	*/

	var body *BlockStmt
	if p.tok == LBRACE {
		body = p.parseBody(/*scope*/)
	}
	//p.expectSemi()

	decl := &FuncDecl {
		Body: body,
	}
	return   decl
}

func (p *Parser) Parse() *ParsedThing{

	var decls []Decl
	for p.tok != EOF {
		decls = append(decls, p.parseDecl())
	}

	return &ParsedThing {
		Decls: decls,
	}
}

func (p *Parser) Init(src []byte) {
	eh :=func(msg string) {//p.errors.Add(msg)
		// TODO
		println(msg)
	}
	p.scanner.Init(src, eh)
	p.next()
}

// this is corresponding to go/parse.go, next0
func (p *Parser) next() {

	p.pos, p.tok, p.lit = p.scanner.Scan()
}
