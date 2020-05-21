package nosj

import (
	"bufio"
	"fmt"
)

type Node interface {
	Parse(s *bufio.Scanner) Node
	fmt.Stringer
}
