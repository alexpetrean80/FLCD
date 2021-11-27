package stack

import g "parser/grammar"

type Stack interface {
	Push(g.Symbol)
	Pop() (g.Symbol, error)
	Top() g.Symbol
	Peek(int) g.Symbols
}
