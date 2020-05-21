package nosj

import "bufio"

type Array struct {
	items []Node
}

func NewArray() *Array {
	return &Array{items: make([]Node, 0)}
}

func (a *Array) Parse(s *bufio.Scanner) Node {
	return a
}
