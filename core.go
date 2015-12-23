package main

type Core struct{
	tok *Tokenizer
}

func NewCore () *Core {
	core := new(Core)
	core.tok = NewTokenizer();
	return core;
}
func (c *Core) Exec(config *Config, prog string) error {


	return nil
}
