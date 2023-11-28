package symbolic_test

import (
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"testing"
)

/*
variable_test.go
Description:
	Testing functions relevant to the Var() object. (Scalar Variable)
*/

/*
TestVariable_Constant1
Description:

	Tests whether or not NumVars returns 0 as the constant included in the a single variable.
*/
func TestVariable_Constant1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	if x.Constant() != 0.0 {
		t.Errorf(
			"The constant component of a variable (%T) should be 0.0; received %v",
			x,
			x.Constant(),
		)
	}

}
