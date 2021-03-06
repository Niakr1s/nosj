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
	SkipSpacesWith(s)

	for {
		a.items = append(a.items, parse(s))
		SkipSpacesWith(s)
		if s.Text() == "," {
			s.Scan()
			SkipSpacesWith(s)
			continue
		}
		break
	}

	SkipSpacesWith(s)

	if s.Text() == "]" {
		s.Scan()
	}

	return a
}

func (a *Array) String() string {
	res := "[ "

	for _, v := range a.items {
		res += fmt.Sprintf("%v, ", v)
	}

	res = strings.TrimSuffix(res, ", ")

	res += " ]"

	return res
}

func (a *Array) PrettyString() string {
	return a.prettyString("")
}

func (a *Array) prettyString(indent string) string {
	res := "\n" + indent + "["

	newIndent := indent + defaultIndent

	for _, v := range a.items {
		res += indent + fmt.Sprintf("%v,", v.prettyString(newIndent))
	}

	res = strings.TrimSuffix(res, ",")

	res += "\n" + indent + "]"

	return res
}
