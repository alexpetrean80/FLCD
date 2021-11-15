package scanner

import (
	"fmt"
	"github.com/alexpetrean80/FLCD/errors"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/alexpetrean80/FLCD/pif"
	"github.com/alexpetrean80/FLCD/symtable"
)

// Scanner is the exposed interface of the scanner
type Scanner interface {
	Scan(path string) []error
	GetPIF() pif.PIF
	GetConstants() symtable.SymbolTable
	GetIdentifiers() symtable.SymbolTable
}

type scanner struct {
	identifiers symtable.SymbolTable
	constants   symtable.SymbolTable
	pif         pif.PIF
}

func New(it, ct symtable.SymbolTable, pif pif.PIF) *scanner {
	return &scanner{
		identifiers: it,
		constants:   ct,
		pif:         pif,
	}
}

// GetPIF returns the pif
func (s scanner) GetPIF() pif.PIF {
	return s.pif
}

// GetConstants returns the constants symbol table
func (s scanner) GetConstants() symtable.SymbolTable {
	return s.constants
}

// GetIdentifiers returns the identifiers symbol table
func (s scanner) GetIdentifiers() symtable.SymbolTable {
	return s.identifiers
}

// Scan parses the source file with the path given as a parameter.
func (s *scanner) Scan(inFile string) (errs []error) {
	bytes, err := ioutil.ReadFile(inFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	tokens := getTokens(bytes)
	pl := len(tokens)

	var tk string
	// loop through all tokens
	for i := 0; i < pl; i++ {
		tk = tokens[i]

		// check for operators of more than one character
		if contains([]string{":", "!", "<", ">", "="}, tk) && i < pl-1 {
			if tokens[i+1] == "=" {
				tk += tokens[i+1]
				i++
			}

		}

		// check if the strings are terminated properly
		if tk == "\"" {
			ok := false
			for i = i + 1; i < pl; i++ {
				tk += tokens[i]
				if tokens[i] == "\"" {
					ok = true
					break
				}
			}

			if !ok {
				errs = append(errs, errors.NewLexical("matching quote not found"))
			}

		}

		// classify the token
		if isReservedWord(tk) || isOperator(tk) || isSeparator(tk) {
			s.pif.Add(tk, 0)
		} else if isConstant(tk) {
			s.constants.Add(tk)
			s.pif.Add(tk, 0)
		} else if isIdentifier(tk) {
			s.identifiers.Add(tk)
			s.pif.Add(tk, 0)
		} else {
			errs = append(errs, errors.NewLexical(fmt.Sprintf("invalid token %s\n", tk)))
		}
	}
	return
}

// getTokens retrieves all the tokens from the source file passed as a slice of bytes.
func getTokens(content []byte) []string {
	var tks []string

	// regex for all separators
	rg, err := regexp.Compile(`[\+\-\*/%!(:=)(!=)\s\t==\n(<=)(>=)=<>\(\)\{\}\[\];,]`)
	if err != nil {
		log.Fatal(err.Error())
	}

	matches := rg.FindAllIndex(content, -1)

	for i, indexes := range matches {
		start := indexes[0]
		end := indexes[1]

		var prev int
		if i == 0 {
			tks = append(tks, string(content[start:end]))
			prev = 0
		} else {
			prev = matches[i-1][1]
		}

		tk := string(content[prev:start])

		if len(strings.TrimSpace(tk)) > 0 {
			tks = append(tks, tk)
		}

		sep := string(content[start:end])
		if len(sep) > 0 {
			tks = append(tks, sep)
		}

	}
	// remove all whitespace tokens to clean the source code
	tks = clearWhiteSpace(tks)
	return tks
}

// clearWhiteSpace removes all the tokens made from whitespace characters(space, tab, newline).
func clearWhiteSpace(tks []string) []string {
	var res []string
	for _, t := range tks {
		if len(strings.TrimSpace(t)) == 0 {
			continue
		}
		res = append(res, t)
	}

	return res
}

func isReservedWord(token string) bool {
	reservedWords := []string{"char", "else", "func", "for", "if", "int", "main", "nil", "string", "var", "print"}
	return contains(reservedWords, token)
}

func isOperator(token string) bool {
	operators := []string{"+", "-", "*", "/", "%", "=", ":=", "==", "!=", "<", "<=", ">", ">=", "&&", "||"}
	return contains(operators, token)
}

func isSeparator(token string) bool {
	separators := []string{"(", ")", "[", "]", "{", "}", ";", ","}
	return contains(separators, token)
}

func isIdentifier(tk string) bool {
	return checkFormat(tk, `^[_a-zA-z][a-zA-Z0-9_]*$`)
}

func isConstant(tk string) bool {
	return isInt(tk) || isString(tk) || isChar(tk)
}

func isInt(tk string) bool {
	return checkFormat(tk, `^(0|[+\-]?[1-9]\d*)$`)
}

func isString(tk string) bool {
	return checkFormat(tk, `^".*"$`)
}

func isChar(tk string) bool {
	return checkFormat(tk, `^'.'$`)
}

// utility procedure which verifies if a string matches a given format.
func checkFormat(tk string, expr string) bool {
	f, err := regexp.Compile(expr)
	if err != nil {
		log.Fatal(err)
	}

	return f.MatchString(tk)
}

// contains is an utility procedure which verifies if a given string appears in a list.
func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}

	return false
}
