package nosj

import (
	"bufio"
)

type Node interface {
	Parse(s *bufio.Scanner) Node
}
