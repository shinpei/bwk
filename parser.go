package main

import (
	"strconv"
)

type Parser struct {
	scanner Scanner

	tok   TokenType
	pos   Pos
	lit   string
	inRhs bool
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

func (p *Parser) error(msg string) {
	panic(msg) // TODO: more smarter way is required.
}

func (p *Parser) expect(tok TokenType) Pos {
	pos := p.pos
	if p.tok != tok {
		p.error(tok.String() + " is expected but we got " + p.tok.String())
	}
	p.next()
	return pos
}

func (p *Parser) parseOperand(lhs bool) Expr {
	switch p.tok {
	case INT, FLOAT, STRING:
		x := &BasicLit{ValuePos: p.pos, Kind: p.tok, Value: p.lit}
		p.next()
		return x
	}
	return &BadExpr{}
}

func (p *Parser) parsePrimaryExpr(lhs bool) Expr {
	x := p.parseOperand(lhs)

	return x
}

func (p *Parser) parseUnaryExpr(lhs bool) Expr {

	switch p.tok {
	case ADD, MUL, SUB, DIV:
		pos, op := p.pos, p.tok
		p.next()
		x := p.parseUnaryExpr(false)
		return &UnaryExpr{OpPos: pos, Op: op, X: x}

	}
	return p.parsePrimaryExpr(lhs)
}

func (p *Parser) tokPrec() (TokenType, int) {
	tok := p.tok
	if p.inRhs && tok == ASSIGN {
		tok = EQL
	}
	return tok, tok.Precedence()
}

func (p *Parser) parseBinaryExpr(lhs bool, prec1 int) Expr {
	x := p.parseUnaryExpr(lhs)

	for _, prec := p.tokPrec(); prec >= prec1; prec-- {
		for {
			op, oprec := p.tokPrec()
			if oprec != prec {
				break
			}
			pos := p.expect(op)
			if lhs {
				lhs = false
			}

			y := p.parseBinaryExpr(false, prec+1)
			x = &BinaryExpr{X: x, OpPos: pos, Op: op, Y: y}
		}
	}

	return x
}

func (p *Parser) parseExpr(lhs bool) Expr {
	return p.parseBinaryExpr(lhs, LowestPrec + 1)
}

func (p *Parser) parseExprList(lhs bool) (list []Expr) {

	list = append(list, p.parseExpr(lhs))
	D("List size=", len(list), ", and tok=", p.tok.String())
	for p.tok == COMMA {
		p.next()
		list = append(list, p.parseExpr(lhs))
	}
	return
}

func (p *Parser) parseLhsList() []Expr {
	old := p.inRhs
	p.inRhs = false
	list := p.parseExprList(true)
	switch p.tok {
	default:
		//for _,  := range list {
		//p.resolve(x) //TODO!
		//}
	}
	p.inRhs = old
	return list
}

func (p *Parser) parseSimpleStmt() (Stmt, bool) {
	lhs := p.parseLhsList()



	D("Lhs has been parsed, cur=", p.tok.String())
	if len(lhs) > 1 {
		p.error("Lhs should be length=1, but"+ strconv.Itoa(len(lhs)))

	}
	return &ExprStmt{X: lhs[0]}, false
}

func (p *Parser) parseStmt() (s Stmt) {
	switch p.tok {
	case SYMBOL, INT, FLOAT, STRING,
		ADD, SUB, MUL, DIV, LPAREN:
		D("at parseStmt: ", p.tok.String())
		s, _ = p.parseSimpleStmt()
	case SEMICOLON:
		s = &EmptyStmt{Semicolon: p.pos}
		p.next()
	default:
		pos := p.pos
		D("Bad Stmt! at pos=" + strconv.Itoa(pos.Offset))
	}
	return
}

func (p *Parser) parseStmtList() (list []Stmt) {
	for p.tok != RBRACE && p.tok != EOF {
		D("processing:", p.tok.String())
		list = append(list, p.parseStmt())
	}
	D("Terminated Stmtlist", p.tok.String())
	return
}

func (p *Parser) parseBody( /*scope Scope */ ) *BlockStmt {
	lbrace := p.expect(LBRACE)
	list := p.parseStmtList()
	rbrace := p.expect(RBRACE)

	return &BlockStmt{
		Lbrace: lbrace,
		List:   list,
		Rbrace: rbrace,
	}
}

func (p *Parser) parseFuncDecl() Decl {

	/*	if p.tok == LPAREN {
			//

		}
	*/

	var body *BlockStmt
	//if p.tok == LBRACE {
	body = p.parseBody( /*scope*/ )
	//}
	//p.expectSemi()

	decl := &FuncDecl{
		Body: body,
	}
	return decl
}

func (p *Parser) parseDecl() Decl {
	return p.parseFuncDecl()

}

func (p *Parser) Parse() (pt *ParsedThing, err error) {

	var decls []Decl
	for p.tok != EOF {
		decls = append(decls, p.parseDecl())
	}

	return &ParsedThing{
		Decls: decls,
	}, nil
}

func (p *Parser) Init(src []byte) {
	eh := func(msg string) { //p.errors.Add(msg)
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
