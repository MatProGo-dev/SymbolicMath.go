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

/*
TestExpression_VariableIDs1
Description:

	Tests the variable IDs function for a single variable.
*/
func TestExpression_VariableIDs1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	ids := symbolic.VariableIDs(x)
	if len(ids) != 1 {
		t.Errorf(
			"Expected the variable IDs to be 1; received %v",
			ids,
		)
	}

	// Verify that the first element of the ids is the same as x's
	if ids[0] != x.ID {
		t.Errorf(
			"Expected the variable ID to be %v; received %v",
			x.ID,
			ids[0],
		)
	}
}

/*
TestExpression_ToExpression1
Description:

	Tests the ToExpression function panics if it is given
	an invalid expression (in this case, a string).
*/
