package main

import (
	"bufio"
	"strings"

	"github.com/niakr1s/nosj"
)

func main() {
	r := strings.NewReader(`{"name":"pavel"}`)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes)

	node := nosj.Parse(s)

	if node == nil {
		panic("no")
	}
}
