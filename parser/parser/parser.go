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
	Parse(gr.Grammar, []gr.Terminal) (string, error)
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

func (p *parser) Parse(g gr.Grammar, w []gr.Terminal) (string, error) {
	p.Reset()
	p.is.Push(g.StartingSymbol)

	for p.state != Final && p.state != Error {
		fmt.Println(p)
		if p.index < len(w) {
			fmt.Println(w[p.index:])
		} else {
			fmt.Println(nil)
		}
		fmt.Println("--------------------------------------------")
		if p.state == Normal {

			if p.index == len(w) && p.is.Empty() {
				p.succes()
			} else {

				iTop, err := p.is.Top()
				if err != nil {
					log.Fatal(err)
				}

				if !iTop.IsTerminal() {
					p.expand(g, w)
				} else {
					iTopTerminal := iTop.(gr.Terminal)
					var currentTerminal gr.Terminal = gr.Terminal("")

					if p.index < len(w) {
						currentTerminal = w[p.index]
					}
					if iTopTerminal == currentTerminal {
						p.advance(g, w)
					} else {
						p.momentaryInsucces()
					}
				}

			}
		} else if p.state == Back {
			wTop, _ := p.ws.Top()

			if wTop.IsTerminal() {
				p.back(g, w)
			} else {
				p.anotherTry(g, w)
			}
		}
	}

	if p.state == Error {
		return "", fmt.Errorf("syntax error: %s not accepted.", w)
	}

	return p.getStringOfProductions(), nil
}

func (p *parser) Reset() {
	p.ws, p.is, p.index, p.state = st.New(), st.New(), initialIndex, Normal
}

func (p *parser) expand(g gr.Grammar, w []gr.Terminal) {
	isTop, err := p.is.Pop()
	if err != nil {
		log.Fatal()
	}
	p.ws.Push(isTop)
	p.ops.PushFront(firstProd)

	prods := g.Productions[isTop.(gr.NonTerminal)]
	prod := prods.GetProd(firstProd)

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

	if p.ops.Len() > 0 {
		p.ops.Remove(p.ops.Front())

	}

	p.is.Push(t)
}

func (p *parser) anotherTry(g gr.Grammar, w []gr.Terminal) {
	lastOp := p.ops.Front().Value.(int)
	lastNonTerminal, err := p.ws.Top()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lastOp)
	nextProd := g.Productions[lastNonTerminal.(gr.NonTerminal)].GetProd(uint(lastOp + 1))
	lastProd := g.Productions[lastNonTerminal.(gr.NonTerminal)].GetProd(uint(lastOp))

	if nextProd != nil {
		p.state = Normal
		p.ops.Remove(p.ops.Front())
		p.ops.PushFront(lastOp + 1)

		for i := len(lastProd) - 1; i >= 0; i-- {
			_, err := p.is.Pop()
			if err != nil {
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

		for i := len(lastProd) - 1; i >= 0; i-- {
			_, err := p.is.Pop()
			if err != nil {
				log.Fatal(err)
			}
		}

		p.is.Push(lastNonTerminal.(gr.NonTerminal))
	}
}

func (p *parser) succes() {
	p.state = Final
}

func (p *parser) getStringOfProductions() string {
	syms := p.ws.List()

	op := p.ops.Back()
	strBuilder := strings.Builder{}

	for _, sym := range syms {
		if !sym.IsTerminal() {
			productionStr := fmt.Sprintf("%s %d, ", sym, op.Value)
			strBuilder.WriteString(productionStr)
		}
		op = op.Prev()
	}

	return strBuilder.String()
}

func (p parser) String() string {
	strBuilder := strings.Builder{}

	state := fmt.Sprintf("State: %d", p.state)
	strBuilder.WriteString(state)
	strBuilder.WriteString("\n")

	index := fmt.Sprintf("Index: %d", p.index)
	strBuilder.WriteString(index)
	strBuilder.WriteString("\n")

	strBuilder.WriteString("W. Stack: ")
	strBuilder.WriteString(p.ws.String())
	strBuilder.WriteString("\n")

	strBuilder.WriteString("P. Index: ")
	for e := p.ops.Front(); e != nil; e = e.Next() {
		strBuilder.WriteString(fmt.Sprintf("%d ", e.Value))
	}
	strBuilder.WriteString("\n")

	strBuilder.WriteString("I. Stack: ")
	strBuilder.WriteString(p.is.String())

	return strBuilder.String()
}
