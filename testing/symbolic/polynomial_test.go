package symbolic_test

/*
polyomial_test.go
Description:
	Tests the functions implemented in the polynomial.go file.
*/

import (
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"testing"
)

/*
TestPolynomial_VariableMonomialIndex1
Description:

	Tests that VariableMonomialIndex returns the correct index when the
	variable is in the monomial.
*/
func TestPolynomial_VariableMonomialIndex1(t *testing.T) {
	// Constants
	v1 := symbolic.NewVariable()
	v2 := symbolic.NewVariable()

	m1 := symbolic.Monomial{
		Coefficient:     1.0,
		VariableFactors: []symbolic.Variable{v1, v2},
		Degrees:         []int{1, 2},
	}

	p1 := symbolic.Polynomial{
		Monomials: []symbolic.Monomial{m1, v2.ToMonomial()},
	}

	// Test that the index is 1
	if p1.VariableMonomialIndex(v2) != 1 {
		t.Errorf(
			"expected index of %v to be 1; received %v",
			v2,
			p1.VariableMonomialIndex(v2),
		)
	}

}
