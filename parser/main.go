package main

import (
	"fmt"
	"parser/grammar"
	"parser/parser"
)

func main() {
	g := grammar.New("./grammar.txt")
	p := parser.New()
	fmt.Println(p.Parse(*g, []grammar.Terminal{"abcd"}))
}
