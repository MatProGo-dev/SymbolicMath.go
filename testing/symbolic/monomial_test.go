package symbolic

/*
monomial_test.go
Description:
	Tests the monomial object.
*/

import (
	"fmt"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"testing"
)

/*
TestMonomial_Check1
Description:

	Tests that the Check() method properly catches an improperly initialized
	monomial.
*/
func TestMonomial_Check1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()

	m1 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1},
		Degrees:         []int{1, 2},
	}

	// Test
	err := m1.Check()
	if err.Error() != fmt.Errorf(
		"the number of degrees (%v) does not match the number of variables (%v)",
		len(m1.Degrees),
		len(m1.VariableFactors),
	).Error() {
		t.Errorf(
			"expected Check() to return false; received %v",
			err,
		)
	}
}
