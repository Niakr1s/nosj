package nosj

import "bufio"

type Object struct {
	items map[string]Node
}

func NewObject() *Object {
	return &Object{items: make(map[string]Node)}
}

func (o *Object) Parse(s *bufio.Scanner) Node {
	if s.Text() != `{` {
		return nil
	}
	s.Scan()

	for {
		key := o.getKey(s)
		o.items[key] = Parse(s)
		SkipSpaces(s)
		if s.Text() == "," {
			s.Scan()
			continue
		}
		break
	}

	SkipSpaces(s)

	if s.Text() == "}" {
		s.Scan()
	}

	return o
}

// scans a key returns it
// consumes s until value
//
// for example input `"name" : "pavel"}`
// will be consumed till `"pavel"}`
func (o *Object) getKey(s *bufio.Scanner) string {
	SkipSpaces(s)

	res := ""

	if s.Text() != `"` {
		return res
	}

	res = ScanQuote(s)

	Skip(s, " ", ":")

	return res
}
