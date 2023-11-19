package symbolic

import (
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"testing"
)

/*
constant_test.go
Description:
	Tests the constant object.
*/

/*
TestConstant_ToMonomial1
Description:

	Tests that this works correctly.
*/
func TestConstant_ToMonomial1(t *testing.T) {
	// Constants
	k1 := symbolic.K(3.14)

	// Test
	m1 := k1.ToMonomial()
	if float64(k1) == m1.Coefficient {
		t.Errorf(
			"expected monomial coefficient to be %v; received %v",
			k1,
			m1,
		)
	}

	if len(m1.VariableFactors) != 0 {
		t.Errorf(
			"expected there to be 0 variables in monomial; received %v",
			len(m1.VariableFactors),
		)
	}

	if len(m1.Degrees) != 0 {
		t.Errorf(
			"expected there to be 0 degrees in monomial; received %v",
			len(m1.Degrees),
		)
	}

}
