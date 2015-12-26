package main

type Core struct{
	tok *Tokenizer
}

func NewCore () *Core {
	core := new(Core)
	core.tok = NewTokenizer();
	return core;
}

func (c *Core) Exec(config *Config, prog string) (err error) {
	var parser Parser
	src := []byte(prog);

	parser.Init(src);
	_ = parser.Parse()
	return nil
}
