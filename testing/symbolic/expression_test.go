package symbolic_test

import (
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"testing"
)

/*
expression_test.go
Description:
	Tests for the functions mentioned in the expression.go file.
*/

/*
TestExpression_NumVariables1
Description:

	Computes the number of variables in a variable.
*/
func TestExpression_NumVariables1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	if symbolic.NumVariables(x) != 1 {
		t.Errorf(
			"The number of variables in a %T should be 1; received %v",
			x,
			symbolic.NumVariables(x),
		)
	}
}
