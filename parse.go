package nosj

import "bufio"

// Parse parses first json object in s into Node
// s must be splitted with bufio.SplitRunes!
func Parse(s *bufio.Scanner) Node {
	var res Node

	SkipSpaces(s)

	if s.Text() == "[" {
		res = NewArray()
	}
	if s.Text() == "{" {
		res = NewObject()
	}
	if s.Text() == `"` {
		res = NewString()
	}

	if res != nil {
		res.Parse(s)
	}

	return res
}
