package scanner

import (
	"bufio"
	"log"
	"os"

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

	return nil
}

func getToken() string {
	return ""
}

func isReservedWord(token string) bool {
	reservedWords := []string{"char", "else", "func", "for", "if", "int", "main", "nil", "string", "var"}
	return contains(reservedWords, token)
}

func isOperator(token string) bool {
	operators := []string{"+", "-", "*", "/", "=", "==", "<", "<=", ">", ">=", "&&", "||"}
	return contains(operators, token)
}

func isSeparator(token string) bool {
	separators := []string{"(", ")", "[", "]", "{", "}", ";"}
	return contains(separators, token)
}

func genPIF() {

}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}

	return false
}
