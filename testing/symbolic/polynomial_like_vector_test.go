package symbolic_test

import (
	getKVector "github.com/MatProGo-dev/SymbolicMath.go/get/KVector"
	"github.com/MatProGo-dev/SymbolicMath.go/symbolic"
	"testing"
)

/*
polynomial_like_vector_test.go
Description:

	Tests for the functions mentioned in the polynomial_like_vector.go file.
*/

/*
TestPolynomialLikeVector_IsPolynomialLikeVector1
Description:

	Tests to make sure that a KVector object is
	identified as a polynomial like vector object.
*/
func TestPolynomialLikeVector_IsPolynomialLikeVector1(t *testing.T) {
	// Constants
	x := getKVector.From([]int{3, 4, 5})

	// Test
	if !symbolic.IsPolynomialLikeVector(x) {
		t.Errorf(
			"Expected IsPolynomialLikeVector(%T) to be true; received false",
			x,
		)
	}
}

/*
TestPolynomialLikeVector_IsPolynomialLikeVector2
Description:

	Tests to make sure that an object that is a MonomialVector
	is identified as a polynomial like vector object.
*/
func TestPolynomialLikeVector_IsPolynomialLikeVector2(t *testing.T) {
	// Constants
	x := symbolic.NewVariableVector(3)

	// Test
	if !symbolic.IsPolynomialLikeVector(x) {
		t.Errorf(
			"Expected IsPolynomialLikeVector(%T) to be true; received false",
			x,
		)
	}
}

/*
TestPolynomialLikeVector_IsPolynomialLikeVector3
Description:

	Tests that a PolynomialVector is identified as a polynomial like vector object.
*/
func TestPolynomialLikeVector_IsPolynomialLikeVector3(t *testing.T) {
	// Constants
	x := symbolic.NewVariableVector(3).ToPolynomialVector()

	// Test
	if !symbolic.IsPolynomialLikeVector(x) {
		t.Errorf(
			"Expected IsPolynomialLikeVector(%T) to be true; received false",
			x,
		)
	}
}

/*
TestPolynomialLikeVector_IsPolynomialLikeVector4
Description:

	Tests that a mat.VecDense object is
	identified as a polynomial like vector object.
*/
func TestPolynomialLikeVector_IsPolynomialLikeVector4(t *testing.T) {
	// Constants
	x := symbolic.OnesVector(3)

	// Test
	if !symbolic.IsPolynomialLikeVector(x) {
		t.Errorf(
			"Expected IsPolynomialLikeVector(%T) to be true; received false",
			x,
		)
	}
}

/*
TestPolynomialLikeVector_ToPolynomialLikeVector1
Description:

	Tests to make sure that a KVector object is converted to a polynomial like vector object.
*/
func TestPolynomialLikeVector_ToPolynomialLikeVector1(t *testing.T) {
	// Constants
	x := getKVector.From([]int{3, 4, 5})

	// Test
	if _, err := symbolic.ToPolynomialLikeVector(x); err != nil {
		t.Errorf(
			"Expected ToPolynomialLikeVector(%T) to be successful; received error: %v",
			x,
			err,
		)
	}
}

/*
TestPolynomialLikeVector_ToPolynomialLikeVector2
Description:

	Tests to make sure that a MonomialVector object is converted to a polynomial like vector object.
*/
func TestPolynomialLikeVector_ToPolynomialLikeVector2(t *testing.T) {
	// Constants
	x := symbolic.NewVariableVector(3).ToMonomialVector()

	// Test
	if _, err := symbolic.ToPolynomialLikeVector(x); err != nil {
		t.Errorf(
			"Expected ToPolynomialLikeVector(%T) to be successful; received error: %v",
			x,
			err,
		)
	}
}

/*
TestPolynomialLikeVector_ToPolynomialLikeVector3
Description:

	Verifies that the ToPolynomialLikeVector() method
	returns an error when given a non-polynomial like object.
*/
func TestPolynomialLikeVector_ToPolynomialLikeVector3(t *testing.T) {
	// Constants
	x := symbolic.NewVariable()

	// Test
	if _, err := symbolic.ToPolynomialLikeVector(x); err == nil {
		t.Errorf(
			"Expected ToPolynomialLikeVector(%T) to return an error; received nil",
			x,
		)
	}
}

/*
TestPolynomialLikeVector_ToPolynomialLikeVector4
Description:

	Tests to make sure that a PolynomialVector object is converted to a polynomial like vector object.
*/
func TestPolynomialLikeVector_ToPolynomialLikeVector4(t *testing.T) {
	// Constants
	x := symbolic.NewVariableVector(3).ToPolynomialVector()

	// Test
	if _, err := symbolic.ToPolynomialLikeVector(x); err != nil {
		t.Errorf(
			"Expected ToPolynomialLikeVector(%T) to be successful; received error: %v",
			x,
			err,
		)
	}
}

/*
TestPolynomialLikeVector_ToPolynomialLikeVector5
Description:

	This test verifies that a mat.VecDense object is converted to a polynomial like vector object.
*/
func TestPolynomialLikeVector_ToPolynomialLikeVector5(t *testing.T) {
	// Constants
	x := symbolic.OnesVector(3)

	// Test
	xOut, err := symbolic.ToPolynomialLikeVector(x)
	if err != nil {
		t.Errorf(
			"Expected ToPolynomialLikeVector(%T) to be successful; received error: %v",
			x,
			err,
		)
	}

	// Check that xOut is a KVector object
	if _, ok := xOut.(symbolic.KVector); !ok {
		t.Errorf(
			"Expected the output to be a KVector; received %T",
			xOut,
		)
	}
}
