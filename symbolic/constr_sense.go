package symbolic

import "fmt"

// ConstrSense represents if the constraint x <= y, x >= y, or x == y. For easy
// integration with Gurobi, the senses have been encoding using a byte in
// the same way Gurobi encodes the constraint senses.
type ConstrSense byte

// Different constraint senses conforming to Gurobi's encoding.
const (
	SenseEqual            ConstrSense = '='
	SenseLessThanEqual                = '<'
	SenseGreaterThanEqual             = '>'
)

func (cs ConstrSense) String() string {
	switch cs {
	case SenseEqual:
		return "="
	case SenseLessThanEqual:
		return "<="
	case SenseGreaterThanEqual:
		return ">="
	default:
		panic(fmt.Errorf("unexpected constraint sense!"))
	}
}

/*
Check
Description:

	This method checks if the receiver is one of the allowed types of sense.
*/
func (cs ConstrSense) Check() error {
	switch cs {
	case SenseEqual:
		return nil
	case SenseLessThanEqual:
		return nil
	case SenseGreaterThanEqual:
		return nil
	default:
		return fmt.Errorf("unexpected constraint sense: %v!", cs)
	}
}
