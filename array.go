package nosj

import (
	"bufio"
	"fmt"
	"strings"
)

type Array struct {
	items []Node
}

func NewArray() *Array {
	return &Array{items: make([]Node, 0)}
}

func (a *Array) Parse(s *bufio.Scanner) Node {
	if s.Text() != `[` {
		return nil
	}
	s.Scan()
	SkipSpaces(s)

	for {
		a.items = append(a.items, parse(s))
		SkipSpaces(s)
		if s.Text() == "," {
			s.Scan()
			continue
		}
		break
	}

	SkipSpaces(s)

	if s.Text() == "]" {
		s.Scan()
	}

	return a
}

func (a *Array) String() string {
	res := "["

	for _, v := range a.items {
		res += fmt.Sprintf("%v,", v)
	}

	res = strings.TrimSuffix(res, ",")

	res += "]"

	return res
}
