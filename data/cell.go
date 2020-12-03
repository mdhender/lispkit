package data

import "fmt"

// Cell is
type Cell struct {
	value Expr
}

func (self Cell) clean() {
	self.value = undef
}

func (self Cell) description() string {
	return fmt.Sprintf("«cell %q»", self)
}

// String implments the stringer interface
func (self Cell) String() string {
	return "${self.value}"
}
