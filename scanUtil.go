package nosj

import (
	"bufio"
)

func SkipSpaces(s *bufio.Scanner) {
	Skip(s, " ")
}

func Skip(s *bufio.Scanner, runes ...string) {
start:
	for {
		for _, r := range runes {
			text := s.Text()
			// if s.Text - one of skipped runes - scan and continue
			if r == text {
				s.Scan()
				continue start
			}
			// if not - return
		}
		return
	}
}

func ScanQuote(s *bufio.Scanner) string {
	res := ""

	if s.Text() != `"` {
		return res
	}

	for s.Scan() {
		if s.Text() == `"` {
			s.Scan()
			return res
		}
		res += s.Text()
	}
	return res
}

func ConsumeScanner(s *bufio.Scanner) string {
	res := s.Text()

	for s.Scan() {
		res += s.Text()
	}

	return res
}
