package main
import "fmt"

type Core struct{
	tok *Tokenizer
	parser *Parser
}

func NewCore () *Core {
	core := new(Core)
	core.tok = NewTokenizer();
	return core;
}

func (c *Core) Exec(config *Config, prog string) (err error) {

	tks, err := c.tok.Tokenize(prog)

	if err != nil {
		return err
	}
	root := c.parser.Parse(tks);
	fmt.Println(root.String())
	return nil
}
