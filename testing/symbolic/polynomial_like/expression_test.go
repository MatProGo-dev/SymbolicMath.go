package polynomial_like_test

import (
	getKMatrix "github.com/MatProGo-dev/SymbolicMath.go/get/KMatrix"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"testing"
)

/*
polynomial_like_test.go
Description:

	Tests for the functions mentioned in the polynomial_like.go file.
*/

/*
TestPolynomialLike_IsPolynomialLike1
Description:

	Tests to make sure that an expression object is identified as not a polynomial like object.
*/
func TestExpression_IsPolynomialLike1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	if symbolic.IsPolynomialLike(x) {
		t.Errorf(
			"Expected IsPolynomialLike(%T) to be false; received true",
			x,
		)
	}
}

/*
TestExpression_IsPolynomialLike2
Description:

	Tests to make sure that an object that is not an expression
	(in this case, a string) is identified as not a polynomial like object.
*/
func TestExpression_IsPolynomialLike2(t *testing.T) {
	// Constants
	x := "test"

	// Test
	if symbolic.IsPolynomialLike(x) {
		t.Errorf(
			"Expected IsPolynomialLike(%T) to be false; received true",
			x,
		)
	}
}

/*
TestExpression_ToPolynomialLike1
Description:

	Tests to make sure that an expression object is converted to a polynomial like object,
	when it is masked as an interface.
*/
func TestExpression_ToPolynomialLike1(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	if _, err := symbolic.ToPolynomialLike(x); err != nil {
		t.Errorf(
			"Expected ToPolynomialLike(%T) to return no error; received %v",
			x,
			err,
		)
	}
}

/*
TestExpression_ToPolynomialLike2
Description:

	Tests to make sure that a KMatrix object is properly
	identified as a PolynomialLike object.
*/
func TestExpression_ToPolynomialLike2(t *testing.T) {
	// Constants
	x := getKMatrix.From(symbolic.Identity(2))

	// Test
	if _, err := symbolic.ToPolynomialLike(x); err != nil {
		t.Errorf(
			"Expected ToPolynomialLike(%T) to return no error; received %v",
			x,
			err,
		)
	}
}

/*
TestExpression_ToPolynomialLike3
Description:

	Tests that a variable vector object is properly identified as a polynomial like object.
*/
func TestExpression_ToPolynomialLike3(t *testing.T) {
	// Constants
	x := symbolic.NewVariableVector(2)

	// Test
	if _, err := symbolic.ToPolynomialLike(x); err != nil {
		t.Errorf(
			"Expected ToPolynomialLike(%T) to return no error; received %v",
			x,
			err,
		)
	}
}

/*
TestExpression_ToPolynomialLike4
Description:

	Verifies that ToPolynomialLike panics when given an object that is not a polynomial like object.
*/
func TestExpression_ToPolynomialLike4(t *testing.T) {
	// Constants
	x := "test"

	// Recover
	defer func() {
		if r := recover(); r == nil {
			t.Errorf(
				"Expected ToPolynomialLike(%T) to panic; received no panic",
				x,
			)
		}
	}()

	// Test
	symbolic.ToPolynomialLike(x)
}
