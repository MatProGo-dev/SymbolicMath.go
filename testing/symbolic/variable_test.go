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

/*
TestVariable_Plus1
Description:

	Tests that the Plus() method works properly when adding a float64 to a variable.
*/
func TestVariable_Plus1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	sum := x.Plus(3.14)
	if sum.(symbolic.ScalarExpression).Constant() != 3.14 {
		t.Errorf(
			"expected %v + 3.14 to have constant component 3.14; received %v",
			x,
			x.Plus(3.14),
		)
	}

	// Test that sum is a polynomial with 2 terms
	sumAsPoly, tf := sum.(symbolic.Polynomial)
	if !tf {
		t.Errorf(
			"expected %v + 3.14 to be a polynomial; received %T",
			x,
			sum,
		)
	}

	if len(sumAsPoly.Monomials) != 2 {
		t.Errorf(
			"expected %v + 3.14 to have 2 terms; received %v",
			x,
			len(sumAsPoly.Monomials),
		)
	}

}

/*
TestVariable_Plus2
Description:

	Tests that the Plus() method works properly when adding a constant to a variable.
*/
func TestVariable_Plus2(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	sum := x.Plus(symbolic.K(3.14))
	if sum.(symbolic.ScalarExpression).Constant() != 3.14 {
		t.Errorf(
			"expected %v + 3.14 to have constant component 3.14; received %v",
			x,
			x.Plus(3.14),
		)
	}

	// Test that sum is a polynomial with 2 terms
	sumAsPoly, tf := sum.(symbolic.Polynomial)
	if !tf {
		t.Errorf(
			"expected %v + 3.14 to be a polynomial; received %T",
			x,
			sum,
		)
	}

	if len(sumAsPoly.Monomials) != 2 {
		t.Errorf(
			"expected %v + 3.14 to have 2 terms; received %v",
			x,
			len(sumAsPoly.Monomials),
		)
	}
}

/*
TestVariable_Plus3
Description:

	Tests that the Plus() method works properly when adding a variable to
	a DIFFERENT variable.
*/
func TestVariable_Plus3(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()
	y := symbolic.NewVariable()

	// Test
	sum := x.Plus(y)
	if sum.(symbolic.ScalarExpression).Constant() != 0.0 {
		t.Errorf(
			"expected %v + %v to have constant component 0.0; received %v",
			x,
			y,
			sum.(symbolic.ScalarExpression).Constant(),
		)
	}

	// Test that sum is a polynomial with 2 terms
	sumAsPoly, tf := sum.(symbolic.Polynomial)
	if !tf {
		t.Errorf(
			"expected %v + %v to be a polynomial; received %T",
			x,
			y,
			sum,
		)
	}

	if len(sumAsPoly.Monomials) != 2 {
		t.Errorf(
			"expected %v + %v to have 2 terms; received %v",
			x,
			y,
			len(sumAsPoly.Monomials),
		)
	}
}

/*
TestVariable_Plus4
Description:

	Tests that the Plus() method works properly when adding a variable to
	the same variable.
*/
func TestVariable_Plus4(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	sum := x.Plus(x)
	if sum.(symbolic.ScalarExpression).Constant() != 0.0 {
		t.Errorf(
			"expected %v + %v to have constant component 0.0; received %v",
			x,
			x,
			sum.(symbolic.ScalarExpression).Constant(),
		)
	}

	// Test that sum is a polynomial with 2 terms
	sumAsPoly, tf := sum.(symbolic.Polynomial)
	if !tf {
		t.Errorf(
			"expected %v + %v to be a polynomial; received %T",
			x,
			x,
			sum,
		)
	}

	if len(sumAsPoly.Monomials) != 1 {
		t.Errorf(
			"expected %v + %v to have 1 term; received %v",
			x,
			x,
			len(sumAsPoly.Monomials),
		)
	}
}

/*
TestVariable_Plus5
Description:


*/

/*
TestVariable_String1
Description:

	Tests that the String() method works properly.
*/
func TestVariable_String1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	if x.String() != "x_0" {
		t.Errorf(
			"expected %v to be \"x\"; received %v",
			x,
			x.String(),
		)
	}
}
