package grammar

type Production interface {
}

type symbols []Symbol

type prod struct {
	nt      NonTerminal
	symbols []symbols
}
