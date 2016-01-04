package main
import "fmt"

type Core struct{
	tok *Tokenizer
}

func NewCore () *Core {
	core := new(Core)
	core.tok = NewTokenizer();
	return core;
}

func (c *Core) exec(codes []Code)    {

}


func (c *Core) compile(ps *ParsedThing) (codes []Code,err error) {

	for _, decl := range ps.Decls {
		if funcDecl, ok := decl.(*FuncDecl); ok {
			for _, stmt := range funcDecl.Body.List {
				if exprStmt, ok := stmt.(*ExprStmt); ok {
					if anotherStmt, ok := exprStmt.X.(*BinaryExpr); ok {
						D("BinExpr:", anotherStmt.Op.String())
					}
				}
			}
		}
	}
	return
}

func (c *Core) EvaluateString(config *Config, prog string) (err error) {
	var parser Parser
	src := []byte(prog);

	parser.Init(src);
	ps, err := parser.Parse()
	if err != nil {
		fmt.Errorf("Parsing error", err)
	}
	code, _ :=c.compile(ps)
	c.exec(code)

	return nil
}
