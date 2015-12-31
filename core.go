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
