package main

import (
	"log"

	"github.com/niakr1s/nosj"
)

func main() {
	got := nosj.ParseString(`{"name": "Pavel"}`)

	log.Print(got)
}
