package grammar

import (
	"encoding/json"
	"log"
	"os"
)

type Terminal string

type NonTerminal string

type Production struct {
}

type Grammar struct {
	Terminals      map[Terminal]bool          `json:"terminals"`
	NonTerminals   map[NonTerminal]bool       `json:"nonTerminals"`
	StartingSymbol NonTerminal                `json:"startingSymbol"`
	Productions    map[NonTerminal]Production `json:"productions"`
}

func New(path string) *Grammar {
	g := Grammar{}

	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("couldn't open grammar file: %s", err.Error())
	}
	defer f.Close()

	d := json.NewDecoder(f)
	if err := d.Decode(&g); err != nil {
		log.Fatalf("couldn't decode grammar: %s", err.Error())
	}

	if _, isDeclared := g.NonTerminals[g.StartingSymbol]; !isDeclared {
		g.NonTerminals[g.StartingSymbol] = true
	}

	return &g
}
