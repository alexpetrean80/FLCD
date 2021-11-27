package linkedlist

import g "parser/grammar"

type node struct {
	sym  g.Symbol
	next *node
}

type LinkedList interface {
	Add(g.Symbol)
	First() g.Symbol
	Empty() bool
	End() bool
}

type list struct {
	first *node
	curr  uint
}

func (l *list) Add(sym g.Symbol) {
	n := node{
		sym:  sym,
		next: l.first,
	}

	l.first = &n
}

func (l *list) First() (sym g.Symbol) {
	return l.first.sym
}

func (l *list) Empty() bool {
	return l.first == nil
}

func (l *list) End() bool {
	return l.first != nil && l.first.next == nil
}
