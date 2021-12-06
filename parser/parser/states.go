package parser

type State int
const (
	Normal State = iota +1
	Back
	Final
	Error
)