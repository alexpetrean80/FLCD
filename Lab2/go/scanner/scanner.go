package scanner

import (
	"bufio"
	"log"
	"os"
	"regexp"

	"github.com/alexpetrean80/FLCD/symtable"
)

type Scanner struct {
	IT *symtable.SymbolTable
	CT *symtable.SymbolTable
}

func (s *Scanner) Scan(path string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	fr := bufio.NewScanner(file)

	for fr.Scan() {
		tk := getToken()
		if isReservedWord(tk) || isOperator(tk) || isSeparator(tk) {
			// genPIF(tk, 0)
		}
	}

	if err = fr.Err(); err != nil {
		return err
	}
}

func getToken() string {
	return ""
}

func isReservedWord(token string) bool {
	// char else func for if int main nil switch string var
	re, err := regexp.Compile(`(char|else|func|for|if|int|main|nil|switch|string|var){1}`)
}

func isOperator(token string) bool {
	// + - * / = == < <= > >= && ||
	re, err := regexp.Compile(`[\+-*/<>(={1,2})(<=)(>=)(&&)(||)]`)
	if err != nil {
		log.Fatal(err)
	}

	return re.MatchString(token)
}

func isSeparator(token string) bool {
	// () [] {} ;
	re, err := regexp.Compile(`\(\)\[\]\{\}\;`)
	if err != nil {
		log.Fatal(err)
	}

	return re.MatchString(token)
}

func genPIF() {

}
