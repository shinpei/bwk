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

func (c *Core) exec(codes []Code, env *Environment) Value{
	for _, code := range codes {
		code.Step(env)
	}

	return env.stack.Pop()
}


func (c *Core) compile(ps *ParsedThing) (codes []Code,err error) {
	gen := new(ExprStmtGen)
	for _, decl := range ps.Decls {
		if funcDecl, ok := decl.(*FuncDecl); ok {
			for _, stmt := range funcDecl.Body.List {
				if exprStmt, ok := stmt.(*ExprStmt); ok {
					stmtCodes := gen.stmtGen(exprStmt)
					codes = append(codes, stmtCodes...)
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
	codes, _ :=c.compile(ps)
	env := NewEnvironment()
	v := c.exec(codes, env)
	fmt.Println(v.String())
	return nil
}
