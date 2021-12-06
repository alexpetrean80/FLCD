package grammar

import (
	"fmt"
	"log"
	"strings"
)

type Production interface {
	GetProd(index uint) Symbols
	Add(syms Symbols)
	Empty() bool
}

type prod struct {
	p []Symbols
}

func NewProduction() Production {
	return &prod{}
}

func (p prod) GetProd(index uint) Symbols {
	if int(index) >= len(p.p) {
		return nil
	}
	return p.p[index]
}

func (p *prod) Add(syms Symbols) {
	p.p = append(p.p, syms)
}
func (p prod) Empty() bool {
	return p.p == nil || len(p.p) > 0
}
func (p *prod) String() string {
	strBuilder := strings.Builder{}

	for _, results := range p.p {
		if _, err := strBuilder.WriteString(fmt.Sprintf("%v", results)); err != nil {
			log.Fatal(err)
		}

		if _, err := strBuilder.WriteString("|"); err != nil {
			log.Fatal(err)
		}
	}

	return strBuilder.String()

}
