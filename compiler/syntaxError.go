package compiler

import "fmt"

// SyntaxError is
type SyntaxError int

const (
	empty SyntaxError = iota
	closingParenthesisMissing
	unexpectedClosingParenthesis
	unexpectedDot
	notAByteValue
	syntaxNotYetSupported
)

// String implements the stringer interface
func (self SyntaxError) String() string {
	switch self {
	case empty:
		return "empty input"
	case closingParenthesisMissing:
		return "closing parenthesis missing"
	case unexpectedClosingParenthesis:
		return "unexpected closing parenthesis"
	case unexpectedDot:
		return "unexpected dot"
	case notAByteValue:
		return "bytevector element not a byte"
	case syntaxNotYetSupported:
		return "syntax not yet supported"
	}
	panic(fmt.Sprintf("assert(self != %d)", int(self)))
}

func eq(lhs, rhs SyntaxError) bool {
	return lhs == rhs
}
