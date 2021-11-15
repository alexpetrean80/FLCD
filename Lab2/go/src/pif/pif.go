package pif

import (
	"fmt"
	"github.com/alexpetrean80/FLCD/utils"
	"log"
	"os"
)

// PIF is the exposed interface of the program internal form
type PIF interface {
	Add(token string, index int)
	utils.Outputter
	fmt.Stringer
}

type pif struct {
	elems []element
}

// New is the constructor of the pif struct
func New() *pif {
	return &pif{
		elems: []element{},
	}
}

// add creates a new element and appends it to the end of pif
func (p *pif) Add(t string, i int) {
	e := element{
		token: t,
		index: i,
	}

	p.elems = append(p.elems, e)
}

func (p pif) Output(outFile string) {
	if err := os.WriteFile(outFile, []byte(p.String()), 0666); err != nil {
		log.Fatalf("cannot write pif to file: %s", err.Error())
	}
}

func (p pif) String() string {
	s := ""
	for i, e := range p.elems {
		e.index = i
		s += e.String()
		s += "\n"
	}

	return s
}

type element struct {
	token string
	index int
}

func (e element) String() string {
	return fmt.Sprintf("token: %s;index: %d", e.token, e.index)
}
