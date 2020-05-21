package nosj

import (
	"bufio"
	"fmt"
)

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

func (str *String) String() string {
	return fmt.Sprintf(`"%s"`, str.s)
}
