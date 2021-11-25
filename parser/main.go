package main

import (
	"fmt"
	"parser/grammar"
)

func main() {
	g := grammar.New("./grammar.json")
	fmt.Println(g)
}
