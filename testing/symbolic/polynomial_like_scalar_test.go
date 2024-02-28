package symbolic_test

import (
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"testing"
)

/*
scalar_expression_test.go
*/

/*
TestScalarExpression_IsScalarExpression1
Description:

	Tests that a float64 is identified as a scalar expression.
*/
func TestScalarExpression_IsScalarExpression1(t *testing.T) {
	// Constants
	x := 3.0

	// Test
	if !symbolic.IsScalarExpression(x) {
		t.Errorf(
			"Expected IsScalarExpression(%T) to be true; received false",
			x,
		)
	}
}

/*
TestScalarExpression_ToScalarExpression1
Description:

	Tests that the method properly returns an error if provided
	an object that is not a scalar expression (in this case, a string).
*/
func TestScalarExpression_ToScalarExpression1(t *testing.T) {
	// Constants
	x := "hello"

	// Test
	_, err := symbolic.ToPolynomialLikeScalar(x)
	if err == nil {
		t.Errorf(
			"Expected ToScalarExpression(%T) to return an error; received nil",
			x,
		)
	}
}

/*
TestScalarExpression_ToScalarExpression2
Description:

	Tests that the method properly converts a float64 to a K object.
*/
func TestScalarExpression_ToScalarExpression2(t *testing.T) {
	// Constants
	x := 3.0

	// Test
	xAsSE, err := symbolic.ToPolynomialLikeScalar(x)
	if err != nil {
		t.Errorf(
			"Expected ToScalarExpression(%T) to return nil; received %v",
			x,
			err,
		)
	}

	// Verify that xAsSE is a K object
	if _, ok := xAsSE.(symbolic.K); !ok {
		t.Errorf(
			"Expected the output to be a K; received %T",
			xAsSE,
		)
	}
}

/*
TestScalarExpression_IsLinear1
Description:

	Tests that a float64 is identified as a linear expression.
*/
func TestScalarExpression_IsLinear1(t *testing.T) {
	// Constants
	x := symbolic.K(3.0)

	// Test
	if !symbolic.IsLinear(x) {
		t.Errorf(
			"Expected IsLinear(%T) to be true; received false",
			x,
		)
	}
}

/*
TestScalarExpression_IsLinear2
Description:

	Tests that a monomial of degree 5 is NOT identified as a linear expression.
*/
func TestScalarExpression_IsLinear2(t *testing.T) {
	// Constants
	x := symbolic.Monomial{
		Coefficient:     3.0,
		VariableFactors: []symbolic.Variable{symbolic.NewVariable()},
		Exponents:       []int{5},
	}

	// Test
	if symbolic.IsLinear(x) {
		t.Errorf(
			"Expected IsLinear(%T) to be false; received true",
			x,
		)
	}
}
