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

// Exprs
type (
	BadExpr struct {
		From, To Pos
	}

	BasicLit struct {
		ValuePos Pos
		Kind TokenType
		Value string // literal string; e.g. 42, 1.2
	}

	SymbolExpr struct {
		NamePos Pos
		Name string
		// Object
	}

	UnaryExpr struct {
		OpPos Pos
		Op TokenType
		X Expr
	}
	BinaryExpr struct {
		X Expr
		OpPos Pos
		Op TokenType
		Y Expr
	}

)

func (x *BadExpr) Pos() Pos { return x.From}
func (x *BasicLit) Pos () Pos {return x.ValuePos}
func (x *SymbolExpr) Pos () Pos {return x.NamePos}
func (x *UnaryExpr) Pos() Pos      { return x.OpPos }
func (x *BinaryExpr) Pos () Pos {return x.X.Pos()}

func (x *BadExpr) End() Pos { return x.To}
func (x *BasicLit) End() Pos { return Pos{Offset:(x.ValuePos.Offset + len(x.Value))} }
func (x *SymbolExpr) End() Pos { return Pos{Offset:(x.NamePos.Offset + len(x.Name))}}
func (x *UnaryExpr) End() Pos      { return x.X.End() }
func (x *BinaryExpr) End() Pos { return x.Y.End()}

func (x *BadExpr) exprNode() {}
func (x *BasicLit) exprNode() {}
func (x *SymbolExpr) exprNode() {}
func (x *UnaryExpr) exprNode() {}
func (x *BinaryExpr) exprNode() {}

//Stmts
type (
	BadStmt struct {
		From, To Pos
	}
	DeclStmt struct {
		Decl Decl
	}

	EmptyStmt struct {
		Semicolon Pos
	}

	ExprStmt struct {
		X Expr
	}

	AssignStmt struct {
		Lhs []Expr
		TokPos Pos
		Tok TokenType
		Rhs []Expr
	}

	// { ... }
	BlockStmt struct {
		Lbrace Pos
		List []Stmt
		Rbrace Pos
	}

)


func (s *BadStmt) Pos() Pos { return s.From}
func (s *DeclStmt) Pos() Pos { return s.Decl.Pos()}
func (s *EmptyStmt) Pos() Pos { return s.Semicolon}
func (s *ExprStmt) Pos() Pos { return s.X.Pos()}
func (s *AssignStmt) Pos() Pos { return s.Lhs[0].Pos()}
func (s *BlockStmt) Pos() Pos { return s.Lbrace}

func (s *BadStmt) End() Pos {return s.To}
func (s *DeclStmt) End() Pos {return s.Decl.End()}
func (s *EmptyStmt) End() Pos {return s.Semicolon} // TODO?
func (s *ExprStmt) End() Pos {return s.X.End()}
func (s *AssignStmt) End() Pos { return s.Rhs[len(s.Rhs)-1].End()}
func (s *BlockStmt) End() Pos {return s.Rbrace/* + 1;*/} // Why +1?


func (s *BadStmt) stmtNode() {}
func (s *DeclStmt) stmtNode() {}
func (s *EmptyStmt) stmtNode() {}
func (s *ExprStmt) stmtNode () {}
func (s *AssignStmt) stmtNode() {}
func (s *BlockStmt) stmtNode() {}


type (
	Symbol struct {
		NamePos Pos
		Name string
		//Obj *Object
	}

	FuncDecl struct {
		Name *Symbol
		Body *BlockStmt
	}
)

func (d *FuncDecl) Pos() Pos { var pos Pos; return pos }
func (d *FuncDecl) End() Pos { var pos Pos; return pos}
func (d *FuncDecl) declNode() {}


