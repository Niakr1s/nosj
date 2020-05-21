package nosj

import "bufio"

type String struct {
	s string
}

func NewString() *String {
	return &String{}
}

func (str *String) Parse(s *bufio.Scanner) Node {
	str.s = ScanQuote(s)
	return str
}
