package main

import (
	"fmt"
	"parser/grammar"
	"parser/parser"
)

func main() {
	g := grammar.New("./grammar.txt")
	p := parser.New()
	fmt.Println(g)
	fmt.Println(p.Parse(*g, []grammar.Terminal{"a", "a", "c", "b", "c"}))
	fmt.Println(p.Parse(*g, []grammar.Terminal{"a", "a", "a"}))
	//fmt.Println(p.Parse(*g, []grammar.Terminal{"0", "0", "0", "1", "1", "0", "1", "1", "1", "0"}))
	fmt.Println(p)
}
