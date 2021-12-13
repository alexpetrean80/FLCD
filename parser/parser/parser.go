package parser

import (
	"container/list"
	"fmt"
	"log"
	gr "parser/grammar"
	st "parser/stack"
	"strings"
)

const initialIndex = 0

const noProd = -1
const firstProd = 0

type Parser interface {
	Parse(gr.Grammar, []gr.Terminal) string
	Reset()
}

type parser struct {
	ws    st.Stack
	ops   *list.List
	is    st.Stack
	index int
	state State
}

func New() *parser {
	return &parser{
		ws:    st.New(),
		is:    st.New(),
		ops:   list.New(),
		index: initialIndex,
		state: Normal,
	}
}

func (p *parser) Parse(g gr.Grammar, w []gr.Terminal) string {
	p.is.Push(g.StartingSymbol)
	for p.state != Final && p.state != Error {
		inputTop, err := p.is.Top()
		if err != nil {
			log.Fatal(err)
		}

		workingTop, err := p.ws.Top()
		if err != nil {
			log.Fatal(err)
		}

		if !inputTop.IsTerminal() {
			p.expand(g, w)
		} else {
			if inputTop == w[p.index] {
				p.advance(g, w)
			} else {
				p.momentaryInsucces()
			}
		}

		if workingTop == nil || workingTop.IsTerminal() {
			p.back(g, w)
		} else if p.state == Back {
			p.anotherTry(g, w)
		}

		if p.index == len(w) && p.is.Empty() {
			p.succes()
		}

	}

	return p.getStringOfProductions()
}

func (p *parser) Reset() {
	p.ws, p.is, p.index, p.state = st.New(), st.New(), initialIndex, Normal
}

func (p *parser) expand(g gr.Grammar, w []gr.Terminal) {
	isTop, err := p.is.Top()
	if err != nil {
		log.Fatal()
	}
	p.ws.Push(isTop)
	p.ops.PushFront(firstProd)

	prod := []gr.Symbol(g.Productions[isTop.(gr.NonTerminal)].GetProd(firstProd))

	for i := len(prod) - 1; i >= 0; i-- {
		p.is.Push(prod[i])
	}
}

func (p *parser) advance(g gr.Grammar, w []gr.Terminal) {
	p.index++
	t, err := p.is.Pop()
	if err != nil {
		log.Fatal(err)
	}

	p.ws.Push(t)
	p.ops.PushFront(noProd)

}

func (p *parser) momentaryInsucces() {
	p.state = Back
}

func (p *parser) back(g gr.Grammar, w []gr.Terminal) {
	if p.index == 0 {
		p.state = Error
		return
	}
	p.index--
	t, err := p.ws.Pop()
	if err != nil {
		log.Fatal(err)
	}
	p.is.Push(t)
}

func (p *parser) anotherTry(g gr.Grammar, w []gr.Terminal) {
	lastOp := p.ops.Front().Value.(int)
	lastNonTerminal, err := p.ws.Top()
	if err != nil {
		log.Fatal(err)
	}
	nextProd := g.Productions[lastNonTerminal.(gr.NonTerminal)].GetProd(uint(lastOp + 1))
	lastProd := g.Productions[lastNonTerminal.(gr.NonTerminal)].GetProd(uint(lastOp))

	if nextProd != nil {
		p.state = Normal
		p.ops.Remove(p.ops.Front())
		p.ops.PushFront(lastOp + 1)

		for i := len(lastProd); i > 0; i-- {
			if _, err := p.is.Pop(); err != nil {
				log.Fatal(err)
			}
		}

		for i := len(nextProd) - 1; i >= 0; i-- {
			p.is.Push(nextProd[i])
		}
	} else {
		if _, err := p.ws.Pop(); err != nil {
			p.state = Error
			log.Fatal(err.Error() + fmt.Sprintf(" - Syntax error: word %s is not accepted by the grammar.", w))
		}

		p.ops.Remove(p.ops.Front())
		p.is.Push(lastNonTerminal.(gr.NonTerminal))
	}
}

func (p *parser) succes() {
	p.state = Final
}

func (p *parser) getStringOfProductions() string {
	syms := p.ws.List()

	op := p.ops.Front()
	strBuilder := strings.Builder{}

	for _, sym := range syms {
		if !sym.IsTerminal() {
			productionStr := fmt.Sprintf("%s %s, ", sym, op.Value)
			strBuilder.WriteString(productionStr)
		}
		op = op.Next()
	}

	return strBuilder.String()
}
