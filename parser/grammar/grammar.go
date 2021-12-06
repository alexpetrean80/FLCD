package grammar

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Grammar struct {
	Terminals      map[Terminal]bool          
	NonTerminals   map[NonTerminal]bool       
	StartingSymbol NonTerminal                
	Productions    map[NonTerminal]Production 
}

func New(path string) *Grammar {
	g := Grammar{}
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("couldn't open grammar file: %s", err.Error())
	}
	defer f.Close()

	lines := readFile(f)

	g.Terminals = getTerminals(lines[0])
	g.NonTerminals = getNonTerminals(lines[1])
	g.Productions = getProductions(lines[2])
	g.StartingSymbol = NonTerminal(lines[3])
	return &g
}

func (g Grammar) IsContextFree() bool {

	return len(g.Terminals) != 0 &&
	len(g.NonTerminals) != 0 &&
	checkIntersectionOfNT(g.NonTerminals, g.Terminals) && 
	g.StartingSymbol != "" &&
	g.isStartInProductions()
}

func checkIntersectionOfNT(nt map[NonTerminal]bool, t map[Terminal]bool) bool {
	for kn, vn := range nt {
		for kt, vt := range t {
			if string(kt) == string(kn) && (vn && vt)  {
				return false
			}
		}
	}

	return true
}

func (g Grammar) isStartInProductions() bool {
	productions, ok := g.Productions[g.StartingSymbol]
	if !ok { 
		return false
	}

	return !productions.Empty()
}

func readFile(f *os.File) (lines []string){
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}



func getTerminals(line string) map[Terminal]bool{
	terminalsStr := strings.Split(line,",")
	terminals := make(map[Terminal]bool)

	for _, terminalStr := range terminalsStr {
		terminals[Terminal(terminalStr)] = true
	}
	return terminals
}

func getNonTerminals(line string) map[NonTerminal]bool{
	nonTerminalsStr := strings.Split(line,",")
	nonTerminals := make(map[NonTerminal]bool)

	for _, nonTerminalStr := range nonTerminalsStr {
		nonTerminals[NonTerminal(nonTerminalStr)] = true
	}
	return nonTerminals
}


func getProductions(line string) (res map[NonTerminal]Production) {
	res = make(map[NonTerminal]Production)
	prods := strings.Split(line, ",")

	for _, prod := range prods {
		aux := strings.Split(prod, "~")
		nt := NonTerminal(aux[0])

		if res[nt] == nil {
			res[nt] = NewProduction()
		}
		p := strings.Split(aux[1], "$")
		r := []Symbol{}
		for _, i := range p {
			// TODO FIX THIS PIECE OF JUNK
			if i == strings.ToLower(i) {
				r = append(r, Terminal(i))
			} else {
				r = append(r, NonTerminal(i))
			}
		}
		res[nt].Add(r)
	}

	return
}

func appendStr(sb *strings.Builder, str string){
	if _, err := sb.WriteString(str); err != nil {
		log.Fatal(err)
	}
}
func (g Grammar) String() string {
	strBuilder := strings.Builder{}

	appendStr(&strBuilder, "Terminals: ")
	
	for nt := range g.NonTerminals {
		appendStr(&strBuilder, fmt.Sprintf("%s ", nt))
	}
	
	appendStr(&strBuilder, "\nNonTerminals: ")
	for t := range g.Terminals {
		appendStr(&strBuilder, fmt.Sprintf("%s ", t))
	}

	appendStr(&strBuilder, "\nProductions: ")
	for nt, p:= range g.Productions {
		appendStr(&strBuilder, fmt.Sprintf("%s %s, ",nt, p))
	}

	appendStr(&strBuilder, "\nStarting Symbol: ")
	appendStr(&strBuilder, string(g.StartingSymbol))

	return strBuilder.String()
}