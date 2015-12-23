package main

type FlagAcceptor interface {
	Desc() string
	Accept(string)
	PopValue() string
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
func (s *DelimiterFlag) PopValue() string {
	return s.Delimiter
}

type VarFlag struct {
}

func (s *VarFlag) Desc() string {
	return "Defines variable with specific name and value. e.g., -v myVar=3"
}

func (s *VarFlag) Accept(arg string) {
}
func (s *VarFlag) PopValue() string {
	return ""
}

type ProgFlag struct {
}

func (s *ProgFlag) Desc() string {
	return "Specify awk program file. e.g., -f mycode.awk"
}
func (s *ProgFlag) Accept(arg string) {
}

func (s *ProgFlag) PopValue() string {
	return ""
}
