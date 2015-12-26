package main

type FlagAcceptor interface {
	Desc() string
	Accept(string)
	// FIXME: cannot fix to string, but interface{} is too much.
	PopString() string
}

type DelimiterFlag struct {
	Delimiter string

}
func NewDelimiterFlag (defaultValue string) *DelimiterFlag {
	f := new(DelimiterFlag)
	f.Delimiter = defaultValue; // default value
	return f
}

func (s *DelimiterFlag) Desc() string {
	return "Specify delimiter for the split, e.g., -F,"
}
func (s *DelimiterFlag) Accept(arg string) {
	s.Delimiter = arg
}
func (s *DelimiterFlag) PopString() string {
	return s.Delimiter
}

type VarFlag struct {
	Vars map[string]interface{}
}

func (s *VarFlag) Desc() string {
	return "Defines variable with specific name and value. e.g., -v myVar=3"
}

func (s *VarFlag) Accept(arg string) {
}
func (s *VarFlag) PopString() string {
	return ""
}

type ProgFlag struct {
}

func (s *ProgFlag) Desc() string {
	return "Specify awk program file. e.g., -f mycode.awk"
}
func (s *ProgFlag) Accept(arg string) {
}

func (s *ProgFlag) PopString() string {
	return ""
}
