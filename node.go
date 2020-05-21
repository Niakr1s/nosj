package nosj

import (
	"bufio"
	"fmt"
)

type Node interface {
	fmt.Stringer
	Parse(s *bufio.Scanner) Node

	PrettyString() string
	prettyString(indent string) string
}

const defaultIndent = "  "
