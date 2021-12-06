package main

import (
	"fmt"
	"parser/grammar"
)

func main() {
	g := grammar.New("./grammar.txt")
	fmt.Println(g)
}
