package nosj

import (
	"bufio"
	"fmt"
	"strings"
)

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
	SkipSpacesWith(s)

	for {
		key := o.getKey(s)
		o.items[key] = parse(s)
		SkipSpacesWith(s)
		if s.Text() == "," {
			s.Scan()
			SkipSpacesWith(s)
			continue
		}
		break
	}

	SkipSpacesWith(s)

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
	SkipSpacesWith(s)

	res := ""

	if s.Text() != `"` {
		return res
	}

	res = ScanQuote(s)

	SkipSpacesWith(s, ":")

	return res
}

func (o *Object) String() string {
	res := "{"

	for k, v := range o.items {
		res += fmt.Sprintf(`"%s" : %s, `, k, v)
	}

	res = strings.TrimSuffix(res, ", ")

	res += " }"

	return res
}

func (o *Object) PrettyString() string {
	return o.prettyString("")
}

func (o *Object) prettyString(indent string) string {
	res := "\n" + indent + "{\n"

	newIndent := indent + defaultIndent

	for k, v := range o.items {
		res += newIndent + fmt.Sprintf("\"%s\" : %s,\n", k, v.prettyString(indent+defaultIndent))
	}

	res = strings.TrimSuffix(res, ",\n")

	res += "\n" + indent + "}"

	return res
}
