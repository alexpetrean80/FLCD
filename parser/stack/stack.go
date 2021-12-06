package stack

import (
	"container/list"
	"errors"
	g "parser/grammar"
)

type Stack interface {
	Push(g.Symbol)
	Pop() (g.Symbol, error)
	Top() (g.Symbol, error)
	Peek(uint) (g.Symbols, error)
	Empty() bool
}

type stack struct {
	l *list.List
}

func New() Stack {
	return &stack{
		l: list.New(),
	}
}

func (s *stack) Push(sym g.Symbol) {
	s.l.PushFront(sym)
}

func (s *stack) Pop() (sym g.Symbol, err error) {
	if s.l.Len() == 0 {
		err = errors.New("stack is empty")
		return
	}

	el := s.l.Front()
	sym = el.Value.(g.Symbol)
	s.l.Remove(el)

	return
}

func (s *stack) Top() (sym g.Symbol, err error) {
	sym, ok := s.l.Front().Value.(g.Symbol)
	if !ok {
		err = errors.New("error casting stack value to symbol")
	}

	return
}

func (s *stack) Peek(count uint) (syms g.Symbols, err error) {
	if uint(s.l.Len()) > count {
		err = errors.New("stack is too short")
		return
	}

	f := s.l.Front()
	for count > 0 {
		count--
		sym, ok := f.Value.(g.Symbol)
		if !ok {
			err = errors.New("error casting stack value to symbol")
			return
		}
		syms = append(syms, sym)
		f = f.Next()
	}

	return
}

func (s *stack) Empty() bool {
	return s.l.Len() == 0
}
