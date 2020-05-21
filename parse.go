package nosj

import (
	"bufio"
	"strings"
)

func ParseString(str string) Node {
	s := prepareScanner(strings.NewReader(str))
	return parse(s)
}

func prepareScanner(r *strings.Reader) *bufio.Scanner {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes)
	s.Scan()
	return s
}

// parse parses first json object in s into Node
// s must be splitted with bufio.SplitRunes!
func parse(s *bufio.Scanner) Node {
	var res Node

	SkipSpacesWith(s)

	r := s.Text()

	if r == "[" {
		res = NewArray()
	}
	if r == "{" {
		res = NewObject()
	}
	if r == `"` {
		res = NewString()
	}

	if res != nil {
		res.Parse(s)
	}

	return res
}
