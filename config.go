package main
import (
	"text/template"
	"os"
)

const display = `
Delimiter: {{.Delimiter}}
Defined variables:
ProgFile : {{.ProgFile}}
`

type Config struct {
	template *template.Template
	Delimiter string
	Variables map[string]interface{}
	ProgFile string
}

func NewConfig() *Config {
	c:= new(Config)
	c.template = template.Must(template.New("hi").Parse(display))
	return c;
}

func (c *Config) Print() {
	c.template.Execute(os.Stdout, c);
}