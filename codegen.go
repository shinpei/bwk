package main
import "strconv"

type StmtGen interface {
	stmtGen(*ExprStmt)  []Code
}
type ExprStmtGen struct {

}

func (e *ExprStmtGen) stmtGen(stmt *ExprStmt) []Code  {

	codes := make([]Code, 0)
	switch t := stmt.X.(type) {
	case *BinaryExpr:
		switch t.Op {
		case ADD:
			//cehck x is int or node
			switch x := t.X.(type) {
			case *BasicLit:
				//TODO: type check.
				val, _ :=strconv.Atoi(x.Value)
				codes = append(codes, &Pushi{x: val})
			default:
				D("ERROR!!!")
			}
		default:
			D("HI")
		}
	default:
		D("BYE")
	}

	return codes
}
