package symbolic_test

import (
	"github.com/MatProGo-dev/SymbolicMath.go/smErrors"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"strings"
	"testing"
)

/*
scalar_expression_test.go
Description:

	Tests for the functions mentioned in the scalar_expression.go file.
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
TestScalarExpression_ToScalarExpression1
Description:

	Tests that if a well-defined expression (that is not a scalar)
	is provided to ToScalarExpression, then an error is returned.
*/
func TestScalarExpression_ToScalarExpression3(t *testing.T) {
	// Setup
	N := 2
	vv := symbolic.NewVariableVector(N)

	// Apply Function
	_, err := symbolic.ToScalarExpression(vv)
	if err == nil {
		t.Errorf("Expected error; received nil")
	} else {
		if !strings.Contains(
			err.Error(),
			"not recognized as a ScalarExpression",
		) {
			t.Errorf(
				"Expected error message to contain 'not recognized as a ScalarExpression'; received %v",
				err.Error(),
			)
		}
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

/*
TestScalarExpression_ScalarPowerTemplate1
Description:

	Tests that the ScalarPowerTemplate() function panics if the input scalar expression
	is not well-defined.
*/
func TestScalarExpression_ScalarPowerTemplate1(t *testing.T) {
	// Setup
	m1 := symbolic.Monomial{
		Coefficient:     2.0,
		VariableFactors: []symbolic.Variable{symbolic.NewVariable()},
		Exponents:       []int{1, 2},
	}

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Expected panic; received nil")
		}

		rAsE := r.(error)
		if !strings.Contains(
			rAsE.Error(),
			m1.Check().Error(),
		) {
			t.Errorf(
				"Expected panic message to contain 'not recognized as a ScalarExpression'; received %v",
				rAsE.Error(),
			)
		}
	}()

	// Call Function
	symbolic.ScalarPowerTemplate(m1, 2)
}

/*
TestScalarExpression_ScalarPowerTemplate2
Description:

	Tests that the ScalarPowerTemplate() function properly panics when a well-defined scalar
	expression is provided, but the power is negative.
*/
func TestScalarExpression_ScalarPowerTemplate2(t *testing.T) {
	// Setup
	x := symbolic.K(2.0)
	testExponent := -2

	// Test
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Expected panic; received nil")
		}

		rAsE := r.(error)
		if !strings.Contains(
			rAsE.Error(),
			smErrors.NegativeExponentError{Exponent: testExponent}.Error(),
		) {
			t.Errorf(
				"Expected panic message to contain 'power must be a non-negative integer'; received %v",
				rAsE.Error(),
			)
		}
	}()

	// Call Function
	symbolic.ScalarPowerTemplate(x, testExponent)
}
