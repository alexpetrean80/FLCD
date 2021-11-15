package errors

import "fmt"

type LexicalError struct {
	msg string
}

func NewLexical(msg string) *LexicalError {
	return &LexicalError{msg: msg}
}
func (e LexicalError) Error() string {
	return fmt.Sprintf("lexical error: %s", e.msg)
}
