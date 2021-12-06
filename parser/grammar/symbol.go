package grammar

type Symbol interface {
	IsTerminal() bool
}

type Symbols []Symbol

type NonTerminal string

func (nt NonTerminal) IsTerminal() bool {
	return false
}

type Terminal string

func (t Terminal) IsTerminal() bool {
	return true
}
