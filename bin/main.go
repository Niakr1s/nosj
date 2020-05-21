package main

import (
	"log"

	"github.com/niakr1s/nosj"
)

func main() {
	cases := []string{
		// `{"name": "Pavel"}`,
		// `{"text": "hello!", "name": "Pavel"}`,
		// `{"text": "hello!", "user": {"name": "Pavel", "token": "12345"}}`,

		`[{"text": "hello!", "user": {"name": "Pavel", "token": "12345"}},
		{"text": "hello!", "user": {"name": "Pavel", "token": "12345"}}]`,
	}

	for _, c := range cases {
		n := nosj.ParseString(c)
		log.Print(n)
	}
}
