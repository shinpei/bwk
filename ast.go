package main

type Node interface {
	Pos() Pos
	End() Pos
}

type Expr interface {
	Node
	exprNode()
}


type Stmt interface {
	Node
	stmtNode()
}

type Decl interface {
	Node
	declNode()
}
type (
	Symbol struct {
		NamePos Pos
		Name string
		//Obj *Object
	}
	BlockStmt struct {
		Lbrace Pos
		List []Stmt
		Rbrace Pos
	}
	FuncDecl struct {
		Name *Symbol
		Body *BlockStmt
	}
)

func (d *FuncDecl) Pos() Pos { var pos Pos; return pos }
func (*FuncDecl) End() Pos { var pos Pos; return pos}
func (*FuncDecl) declNode() {}


